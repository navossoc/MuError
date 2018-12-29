#!/bin/bash
# Reference: https://github.com/containous/traefik/blob/master/script/crossbinary-default

set -e
mkdir -p dist

GO_BUILD_CMD="go build"

OS_PLATFORM_ARG=(linux windows darwin)
OS_ARCH_ARG=(amd64 386)
for OS in ${OS_PLATFORM_ARG[@]}; do
  BIN_EXT=''
  if [ "$OS" == "windows" ]; then
    BIN_EXT='.exe'
  fi
  for ARCH in ${OS_ARCH_ARG[@]}; do
    echo "Building binary for ${OS}/${ARCH}..."
    GOARCH=${ARCH} GOOS=${OS} CGO_ENABLED=0 ${GO_BUILD_CMD} -o "dist/MuError_${OS}-${ARCH}${BIN_EXT}" .
  done
done
