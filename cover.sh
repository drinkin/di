#!/usr/bin/env bash

set -eE -o pipefail

PROJECT=$(basename "$(git rev-parse --show-toplevel)")
COVER_TMP=$(mktemp -t "cover.$PROJECT.XXX")

echo "mode: atomic" > "$COVER_TMP"
for cp in $(find . -type f -name "*.coverprofile" | grep "${1:-""}");
do
  tail -n +2 "$cp" >> "$COVER_TMP"
done


cat "$COVER_TMP" > coverage.txt
