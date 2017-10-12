#!/usr/bin/env bash
set -e
echo "building go server for Linux..."
GOOS=linux go build
docker build -t roycevanle/zipsvr .
docker push roycevanle/zipsvr
go clean
