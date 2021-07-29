#!/usr/bin/env bash


docker run --rm -it -v $PWD:/app -w /app golang:alpine sh -c "go run table_header/generate.go"
for lib in viper configor
do
    cmd="run-${lib}.sh"
    bash $cmd
done

