#! /usr/bin/make

all: clean generate

clean:
	@rm -rf ./gen

generate:
	goa gen github.com/shamhub/goa-example/api

stubs:
	goa example github.com/shamhub/goa-example/api

build:
	go build -o server-binary ./cmd/server && go build -o client-binary ./cmd/server-cli

test:
	go test ./...