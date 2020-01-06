#!/bin/bash

set -ex

mkdir -p deployment
GOOS=linux GOARCH=amd64 go build -o ./deployment/ci-info.linux-amd64 .
tar czfv ./deployment/ci-info.linux-amd64.tar.gz ./deployment/ci-info.linux-amd64

shasum -a 256 ./deployment/* > ./deployment/sha256sums.txt
cat ./deployment/sha256sums.txt
