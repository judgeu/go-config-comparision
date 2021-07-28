#!/usr/bin/env bash


for cmd in "go run -mod vendor viper/base1.go" "VP_SINCE=600s VP_IDK=yes VP_DATABASE_PORT=5555 VP_GAMES_4=game4 go run -mod vendor viper/base1.go";
do
    echo "=== $cmd ==="
    docker run --rm -it -v $PWD:/app -w /app golang:alpine sh -c "$cmd"
    #eval $cmd
    echo
done

