#!/usr/bin/env bash

docker run --rm -it -v $PWD:/app -w /app golang bash -c "go mod $1"