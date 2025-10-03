#!/usr/bin/env bash
set -euo pipefail

HOST="${HOST:-localhost:8080}"

echo "Health:"
curl -sf "http://$HOST/healthz" && echo || (echo "health failed"; exit 1)

echo "\nTools list:"
curl -s \
  -H 'Content-Type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}' \
  "http://$HOST/ws" || echo "Use a WS client to test interactively"

echo "Done"


