#!/usr/bin/env bash

set -eu -o pipefail

if ! command -v co-author &>/dev/null; then
  echo "co-author not installed or available in the PATH" >&2
  echo "run command below to install it" >&2
  echo "go install github.com/erkanzileli/co-author@latest" >&2
  echo "or check documentations https://github.com/erkanzileli/co-author" >&2
  exit 1
fi

exec co-author commit "$@" </dev/tty >/dev/tty
