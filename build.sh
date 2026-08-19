#!/usr/bin/env bash

set -euo pipefail

VERSION="${1:-dev}"
OS="${2:-$(go env GOOS)}"
ARCH="${3:-$(go env GOARCH)}"

export GOOS="$OS"
export GOARCH="$ARCH"
export CGO_ENABLED=0

EXE=""
if [[ "$GOOS" == "windows" ]]; then
  EXE=".exe"
fi

mkdir -p outputs

LDFLAGS="-X github.com/zavocc/youtube-watcher-cli/internal/shared.Version=$VERSION"

go build \
  -ldflags "$LDFLAGS" \
  -o "outputs/youtube-watcher-cli-${GOOS}-${GOARCH}-${VERSION}${EXE}" \
  ./cli/youtube-watcher-cli/

go build \
  -ldflags "$LDFLAGS" \
  -o "outputs/youtube-search-cli-${GOOS}-${GOARCH}-${VERSION}${EXE}" \
  ./cli/youtube-search-cli/
