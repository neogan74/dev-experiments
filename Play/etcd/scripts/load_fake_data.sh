#!/usr/bin/env bash
set -euo pipefail

COMPOSE_CMD=${COMPOSE_CMD:-docker compose}
ETCD_SERVICE=${ETCD_SERVICE:-etcd}
ETCD_ENDPOINT=${ETCD_ENDPOINT:-http://localhost:2379}

STRING_COUNT=${STRING_COUNT:-50}
JSON_COUNT=${JSON_COUNT:-20}
BLOB_COUNT=${BLOB_COUNT:-10}
BLOB_BYTES=${BLOB_BYTES:-32}

log() {
  printf '[load-etcd] %s\n' "$*" >&2
}

compose_exec() {
  "$COMPOSE_CMD" exec -T "$ETCD_SERVICE" "$@"
}

run_etcdctl() {
  compose_exec etcdctl --endpoints="$ETCD_ENDPOINT" "$@"
}

ensure_ready() {
  if ! run_etcdctl endpoint health >/dev/null 2>&1; then
    log "etcd endpoint $ETCD_ENDPOINT is not reachable via service '$ETCD_SERVICE'. Is the compose stack up?"
    exit 1
  fi
}

load_strings() {
  log "Loading $STRING_COUNT simple string keys"
  for i in $(seq 1 "$STRING_COUNT"); do
    key=$(printf 'datasets/simple/%03d' "$i")
    value=$(printf 'sample-value-%03d' "$i")
    run_etcdctl put "$key" "$value" >/dev/null
  done
}

load_json() {
  log "Loading $JSON_COUNT JSON documents"
  for i in $(seq 1 "$JSON_COUNT"); do
    key=$(printf 'datasets/json/%03d' "$i")
    activated=$([ $((i % 2)) -eq 0 ] && echo "true" || echo "false")
    payload=$(printf '{"id":%d,"email":"user%03d@example.com","active":%s,"roles":["viewer","tester"]}' "$i" "$i" "$activated")
    run_etcdctl put "$key" "$payload" >/dev/null
  done
}

load_binary() {
  log "Loading $BLOB_COUNT binary blobs ($BLOB_BYTES bytes each)"
  for i in $(seq 1 "$BLOB_COUNT"); do
    key=$(printf 'datasets/blob/%03d' "$i")
    compose_exec sh -s "$key" "$ETCD_ENDPOINT" "$BLOB_BYTES" <<'EOF'
set -euo pipefail
key=$1
endpoint=$2
size=$3
head -c "$size" /dev/urandom | etcdctl --endpoints="$endpoint" put "$key" >/dev/null
EOF
  done
}

emit_summary() {
  log "Loaded keys under prefix 'datasets/'."
  log "Inspect with: $COMPOSE_CMD exec $ETCD_SERVICE etcdctl --endpoints=$ETCD_ENDPOINT get datasets/ --prefix --keys-only"
}

main() {
  ensure_ready
  load_strings
  load_json
  load_binary
  emit_summary
}

main "$@"
