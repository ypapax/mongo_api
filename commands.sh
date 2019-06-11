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
	curl -X POST -H "Content-Type: application/json" -d @./testInput.json http://localhost:80/auditlog
}

test2(){
	curl -X POST -H "Content-Type: application/json" -d @./testInput2.json http://localhost:80/auditlog
}

startMongo(){
	 mongod --config /usr/local/etc/mongod.conf
}

$@
