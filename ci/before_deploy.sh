#!/bin/bash

set -ex

mkdir -p releases
GOOS=linux GOARCH=amd64 go build -o ./releases/ci-info.linux-amd64 .
tar czfv ./releases/ci-info.linux-amd64.tar.gz ./releases/ci-info.linux-amd64

shasum -a 256 ./releases/* > ./releases/sha256sums.txt
cat ./releases/sha256sums.txt
