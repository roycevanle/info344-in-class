#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t roycevanle/testserver .
docker push roycevanle/testserver
