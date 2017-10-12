#!/usr/bin/env bash
set -e
echo "building linux executable"
docker build -t roycevanle/zipclient .
docker push roycevanle/zipclient
