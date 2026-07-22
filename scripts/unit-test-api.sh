#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

cd "$ROOT_DIR"

echo "==========================================="
echo "Running unit tests from internal/api"
echo "==========================================="

packages=$(go list ./internal/api/...)

failed=0

for pkg in $packages; do
    echo
    echo ">>> Testing: $pkg"

    if ! go test -count=1 -race -v "$pkg"; then
        echo
        echo "FAILED: $pkg"
        failed=1
    fi
done

echo

if [ "$failed" -ne 0 ]; then
    echo "==========================================="
    echo "Some tests failed."
    echo "==========================================="
    exit 1
fi

echo "==========================================="
echo "All API tests passed!"
echo "==========================================="
