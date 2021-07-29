#!/usr/bin/env bash


for cmd in "go run -mod vendor configor/base1.go" "CONFIGOR_SINCE=600s CONFIGOR_IDK=yes CONFIGOR_DATABASE_PORT=5555 CONFIGOR_DB_PORT=6666 CONFIGOR_GAMES_4=game4 go run -mod vendor configor/base1.go";
do
    echo -n "| ${cmd} | "
    docker run --rm -it -v $PWD:/app -w /app golang:alpine sh -c "$cmd"
    echo
done

