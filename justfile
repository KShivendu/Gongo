#!/usr/bin/env just --justfile

set dotenv-load := true

HOST := "0.0.0.0"
PORT := "8000"

# display all commands
default:
    @just --list --unsorted

stop:
    fuser -k 5000/tcp &

dev: stop
    air {{HOST}}:{{PORT}}

build:
    go build

start: build stop
    ./gongo