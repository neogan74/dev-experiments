#!/usr/bin/env python3
"""Parallel etcd data loader powered by Faker."""

import argparse
import json
import math
import random
import threading
import time
from typing import Iterable, List, Sequence, Tuple

import etcd3
from faker import Faker


ValueType = str


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--endpoints",
        default="localhost:2379",
        help="Comma-separated list of etcd endpoints (host:port).",
    )
    parser.add_argument(
        "-n",
        "--count",
        type=int,
        default=10_000,
        help="Total number of keys to write.",
    )
    parser.add_argument(
        "-c",
        "--concurrency",
        type=int,
        default=16,
        help="Number of concurrent writer threads.",
    )
    parser.add_argument(
        "--prefix",
        default="datasets/python",
        help="Key prefix for generated entries.",
    )
    parser.add_argument(
        "--types",
        default="text,json,binary",
        help="Comma-separated value types to generate (choices: text,json,binary).",
    )
    parser.add_argument(
        "--binary-min",
        type=int,
        default=32,
        help="Minimum size in bytes for binary payloads.",
    )
    parser.add_argument(
        "--binary-max",
        type=int,
        default=256,
        help="Maximum size in bytes for binary payloads.",
    )
    parser.add_argument(
        "--progress",
        type=int,
        default=5_000,
        help="Emit progress every N successful writes (0 disables).",
    )
    parser.add_argument(
        "--seed",
        type=int,
        default=int(time.time_ns() % (1 << 31)),
        help="Seed for deterministic data generation.",
    )
    parser.add_argument(
        "--user",
        default=None,
        help="Optional etcd username for authenticated clusters.",
    )
    parser.add_argument(
        "--password",
        default=None,
        help="Optional etcd password for authenticated clusters.",
    )
    return parser.parse_args()


def parse_endpoints(raw: str) -> Sequence[Tuple[str, int]]:
    endpoints: List[Tuple[str, int]] = []
    for part in raw.split(","):
        trimmed = part.strip()
        if not trimmed:
            continue
        if ":" not in trimmed:
            raise ValueError(f"endpoint '{trimmed}' must be in host:port format")
        host, port_str = trimmed.rsplit(":", 1)
        try:
            port = int(port_str)
        except ValueError as exc:
            raise ValueError(f"endpoint '{trimmed}' has invalid port") from exc
        endpoints.append((host, port))
    if not endpoints:
        raise ValueError("no valid endpoints provided")
    return endpoints


def choose_value_types(raw: str) -> Sequence[ValueType]:
    allowed = {"text", "json", "binary"}
    result: List[ValueType] = []
    for part in raw.split(","):
        t = part.strip().lower()
        if not t:
            continue
        if t not in allowed:
            raise ValueError(f"unknown value type '{t}', allowed: {sorted(allowed)}")
        if t not in result:
            result.append(t)
    if not result:
        raise ValueError("no value types selected")
    return result


def chunked_indices(total: int, workers: int) -> Iterable[Tuple[int, range]]:
    for worker_id in range(workers):
        yield worker_id, range(worker_id, total, workers)


def generate_entry(
    faker: Faker,
    rng: random.Random,
    prefix: str,
    idx: int,
    value_types: Sequence[ValueType],
    binary_min: int,
    binary_max: int,
) -> Tuple[str, bytes]:
    value_type = rng.choice(value_types)
    suffix = f"{rng.getrandbits(64):016x}"
    base_key = f"{prefix}/{value_type}/{idx:09d}-{suffix}"

    if value_type == "json":
        payload = {
            "order_id": f"py-{idx}",
            "customer": faker.name(),
            "email": faker.email(),
            "product": faker.catch_phrase(),
            "price": round(rng.uniform(10.0, 5000.0), 2),
            "purchased_at": faker.date_time_between(start_date="-30d").isoformat(),
            "tags": [faker.bs(), faker.company_suffix()],
        }
        return base_key, json.dumps(payload).encode("utf-8")

    if value_type == "binary":
        size = binary_min
        if binary_max > binary_min:
            size += rng.randint(0, binary_max - binary_min)
        buf = bytearray(size)
        for i in range(size):
            buf[i] = rng.randrange(0, 256)
        return base_key, bytes(buf)

    # default to plain text
    sentence = faker.sentence(nb_words=16)
    return base_key, sentence.encode("utf-8")


def writer(
    worker_id: int,
    indices: Iterable[int],
    endpoint: Tuple[str, int],
    creds: Tuple[str, str],
    args: argparse.Namespace,
    progress: dict,
    lock: threading.Lock,
) -> Tuple[int, int]:
    host, port = endpoint
    user, password = creds
    client = etcd3.client(host=host, port=port, user=user, password=password)
    faker = Faker()
    faker.seed_instance(args.seed + worker_id)
    rng = random.Random(args.seed + worker_id)

    successes = 0
    failures = 0
    try:
        for idx in indices:
            key, value = generate_entry(
                faker,
                rng,
                args.prefix,
                idx,
                args.value_types,
                args.binary_min,
                args.binary_max,
            )
            try:
                client.put(key, value)
                successes += 1
                if args.progress > 0:
                    with lock:
                        progress["completed"] += 1
                        if progress["completed"] % args.progress == 0:
                            rate = 0.0
                            elapsed = time.time() - progress["start"]
                            if elapsed > 0:
                                rate = progress["completed"] / elapsed
                            print(
                                f"[python-faker] progress: {progress['completed']}/{args.count} "
                                f"writes ({rate:,.0f} ops/s)"
                            )
            except Exception as exc:
                failures += 1
                print(f"[python-faker] worker {worker_id} failed to write {key}: {exc}")
    finally:
        client.close()
    return successes, failures


def main() -> None:
    args = parse_args()
    try:
        args.endpoints = parse_endpoints(args.endpoints)
        args.value_types = choose_value_types(args.types)
    except ValueError as exc:
        raise SystemExit(f"argument error: {exc}")

    if args.count <= 0:
        raise SystemExit("argument error: --count must be > 0")
    if args.concurrency <= 0:
        raise SystemExit("argument error: --concurrency must be > 0")
    if args.binary_min <= 0:
        raise SystemExit("argument error: --binary-min must be > 0")
    if args.binary_max < args.binary_min:
        print(
            f"[python-faker] adjusting binary-max ({args.binary_max}) "
            f"to binary-min ({args.binary_min})"
        )
        args.binary_max = args.binary_min

    creds = (args.user, args.password)
    start = time.time()

    progress = {"completed": 0, "start": start}
    lock = threading.Lock()

    total_success = 0
    total_failures = 0

    # Spread workers across endpoints in round-robin fashion.
    endpoint_cycle = list(args.endpoints)

    from concurrent.futures import ThreadPoolExecutor

    with ThreadPoolExecutor(max_workers=args.concurrency) as executor:
        futures = []
        for worker_id, idx_range in chunked_indices(args.count, args.concurrency):
            endpoint = endpoint_cycle[worker_id % len(endpoint_cycle)]
            futures.append(
                executor.submit(
                    writer,
                    worker_id,
                    idx_range,
                    endpoint,
                    creds,
                    args,
                    progress,
                    lock,
                )
            )

        for future in futures:
            success, failure = future.result()
            total_success += success
            total_failures += failure

    elapsed = time.time() - start
    ops = total_success / elapsed if elapsed > 0 else math.inf
    print(
        f"[python-faker] completed in {elapsed:.2f}s "
        f"({total_success} successes, {total_failures} failures, {ops:,.0f} ops/s)"
    )


if __name__ == "__main__":
    main()
