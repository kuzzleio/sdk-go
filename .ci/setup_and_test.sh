#!/usr/bin/env bash
set -e

mkdir -p /root/go/src/github.com/kuzzleio/sdk-go
cp -fr /mnt/* /root/go/src/github.com/kuzzleio/sdk-go/
cd /root/go/src/github.com/kuzzleio/sdk-go/
go get ./...
./test.sh
cp -fr /root/go/src/github.com/kuzzleio/sdk-go/.cover /mnt/
cd /mnt
