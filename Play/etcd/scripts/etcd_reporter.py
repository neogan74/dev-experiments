#!/usr/bin/env python3
"""Summarize etcd keyspaces with basic counts and JSON price statistics."""

import argparse
import json
import math
import sys
import time
from dataclasses import dataclass, field
from typing import Dict, Iterable, List, Optional, Sequence, Tuple

import etcd3


ValueType = str


@dataclass
class TypeStats:
    count: int = 0
    total_bytes: int = 0
    price_sum: float = 0.0
    price_count: int = 0
    price_min: float = field(default_factory=lambda: math.inf)
    price_max: float = field(default_factory=lambda: -math.inf)


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--endpoints",
        default="localhost:2379",
        help="Comma-separated list of etcd endpoints (host:port).",
    )
    parser.add_argument(
        "--prefix",
        default="datasets/loadtest",
        help="Key prefix to inspect.",
    )
    parser.add_argument(
        "--batch-size",
        type=int,
        default=0,
        help="Optional page size for the underlying range request (0 uses server default).",
    )
    parser.add_argument(
        "--serializable",
        action="store_true",
        help="Use serializable reads instead of linearizable (may be faster for read-only workloads).",
    )
    parser.add_argument(
        "--user",
        default=None,
        help="Optional etcd username.",
    )
    parser.add_argument(
        "--password",
        default=None,
        help="Optional etcd password.",
    )
    parser.add_argument(
        "--dial-timeout",
        type=float,
        default=5.0,
        help="Dial timeout in seconds when connecting to etcd.",
    )
    parser.add_argument(
        "--limit",
        type=int,
        default=0,
        help="Optional hard cap on the number of keys to process (0 = no limit).",
    )
    return parser.parse_args()


def parse_endpoints(raw: str) -> Sequence[Tuple[str, int]]:
    endpoints: List[Tuple[str, int]] = []
    for part in raw.split(","):
        item = part.strip()
        if not item:
            continue
        if ":" not in item:
            raise ValueError(f"endpoint '{item}' must be host:port")
        host, port_str = item.rsplit(":", 1)
        try:
            port = int(port_str)
        except ValueError as exc:
            raise ValueError(f"endpoint '{item}' has invalid port") from exc
        endpoints.append((host, port))
    if not endpoints:
        raise ValueError("no endpoints provided")
    return endpoints


def connect(endpoints: Sequence[Tuple[str, int]], user: Optional[str], password: Optional[str], timeout: float) -> etcd3.Etcd3Client:
    last_err: Optional[Exception] = None
    for host, port in endpoints:
        try:
            client = etcd3.client(host=host, port=port, user=user, password=password, timeout=timeout)
            # status() forces a round-trip; fail fast if endpoint unavailable.
            client.status()
            return client
        except Exception as exc:
            last_err = exc
            continue
    raise RuntimeError(f"unable to connect to any endpoint: {last_err}")  # type: ignore[arg-type]


def detect_type(key: str) -> ValueType:
    if "/json/" in key:
        return "json"
    if "/blob/" in key:
        return "binary"
    if "/text/" in key:
        return "text"
    return "other"


def update_stats(stats: Dict[ValueType, TypeStats], key: str, value: bytes) -> None:
    vtype = detect_type(key)
    bucket = stats.setdefault(vtype, TypeStats())
    bucket.count += 1
    bucket.total_bytes += len(value)

    if vtype == "json":
        try:
            obj = json.loads(value.decode("utf-8"))
        except Exception:
            return
        price = extract_price(obj)
        if price is None:
            return
        bucket.price_sum += price
        bucket.price_count += 1
        if price < bucket.price_min:
            bucket.price_min = price
        if price > bucket.price_max:
            bucket.price_max = price


def extract_price(obj) -> Optional[float]:
    if isinstance(obj, dict) and "price" in obj:
        return to_float(obj["price"])
    return None


def to_float(value) -> Optional[float]:
    if isinstance(value, (int, float)):
        return float(value)
    if isinstance(value, str):
        try:
            return float(value)
        except ValueError:
            return None
    if isinstance(value, json.Number):  # type: ignore[attr-defined]
        try:
            return float(value)
        except ValueError:
            return None
    return None


def iterate_keys(
    client: etcd3.Etcd3Client,
    prefix: str,
    batch_size: int,
    serializable: bool,
) -> Iterable[Tuple[bytes, bytes]]:
    kwargs: Dict[str, object] = {
        "sort_order": "ascend",
        "sort_target": "key",
        "serializable": serializable,
    }
    if batch_size > 0:
        kwargs["limit"] = batch_size

    for value, meta in client.get_prefix(prefix, **kwargs):
        if value is None or meta is None:
            continue
        yield meta.key, value


def format_bytes(num_bytes: int) -> str:
    if num_bytes == 0:
        return "0 B"
    units = ["B", "KB", "MB", "GB", "TB"]
    idx = min(int(math.log(num_bytes, 1024)), len(units) - 1)
    scaled = num_bytes / (1024 ** idx)
    return f"{scaled:.2f} {units[idx]}"


def main() -> None:
    args = parse_args()
    try:
        endpoints = parse_endpoints(args.endpoints)
    except ValueError as exc:
        print(f"argument error: {exc}", file=sys.stderr)
        raise SystemExit(2)

    try:
        client = connect(endpoints, args.user, args.password, args.dial_timeout)
    except Exception as exc:
        print(f"[etcd-reporter] failed to establish connection: {exc}", file=sys.stderr)
        raise SystemExit(1)

    stats: Dict[ValueType, TypeStats] = {}
    total_keys = 0

    start = time.time()
    try:
        for idx, (key, value) in enumerate(iterate_keys(client, args.prefix, args.batch_size, args.serializable), start=1):
            update_stats(stats, key.decode("utf-8"), value)
            total_keys += 1
            if args.limit and total_keys >= args.limit:
                break
    finally:
        client.close()

    elapsed = time.time() - start
    print("=== etcd dataset report (python) ===")
    print(f"Prefix: {args.prefix}")
    print(f"Total keys: {total_keys}")
    print(f"Elapsed: {elapsed:.2f}s")

    if total_keys == 0:
        return

    print("\nValue types:")
    for vtype in ("text", "json", "binary", "other"):
        bucket = stats.get(vtype, TypeStats())
        print(f"  - {vtype}: {bucket.count} keys ({format_bytes(bucket.total_bytes)})")
        if vtype == "json" and bucket.price_count > 0:
            avg_price = bucket.price_sum / bucket.price_count
            min_price = bucket.price_min if bucket.price_min != math.inf else 0.0
            max_price = bucket.price_max if bucket.price_max != -math.inf else 0.0
            print(
                f"      prices -> avg: {avg_price:.2f}, min: {min_price:.2f}, max: {max_price:.2f} "
                f"({bucket.price_count} records)"
            )


if __name__ == "__main__":
    main()
