#!/bin/sh

go run hi.go "$@" && exit 0

echo "nope; pwd is $PWD"
env

exec go run ./hi.go -- "$@"

