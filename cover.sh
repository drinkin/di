#!/usr/bin/env bash

set -eE -o pipefail
shopt -s globstar

cd "$(git rev-parse --show-toplevel)"


rm -rf /tmp/cover
mkdir -p /tmp/cover


for coverFile in ./**/*.coverprofile; do
  out="/tmp/cover/$(basename "$coverFile" .coverprofile).html"
  go tool cover -html "$coverFile" -o "$out"
  command -v open >/dev/null 2>&1 && open "$out"
done
