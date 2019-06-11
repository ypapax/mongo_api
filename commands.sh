#!/usr/bin/env bash
set -ex

build(){
	go get ./...
	go install
}

run(){
	go run main.go
}

test(){
	curl localhost:80/auditlog
}
$@
