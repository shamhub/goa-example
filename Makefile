#! /usr/bin/make

all: clean generate

clean:
	@rm -rf ./gen

keys:
	rm -f /tmp/test_rsa*
	ssh-keygen -q -t rsa -m PEM -f /tmp/test_rsa -q -N ""
	openssl rsa -in /tmp/test_rsa -pubout -out /tmp/test_rsa.pub

generate:
	goa gen github.com/shamhub/goa-example/api

stubs:
	goa example github.com/shamhub/goa-example/api

build:
	go build -o server-binary ./cmd/server

build-cli:	
	go build -o client-binary ./cmd/server-cli

test:
	go test ./...

run-local-http: keys build
	./server-binary -debug -private-key /tmp/test_rsa -public-key /tmp/test_rsa.pub