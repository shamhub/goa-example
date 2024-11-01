#! /usr/bin/make

all: clean generate

clean:
	@rm -rf ./gen

generate:
	goa gen github.com/shamhub/goa-example/api

example:
	goa example github.com/shamhub/goa-example/api

buildserver:
	go build -o server-binary ./cmd/calc

buildclient:
	go build -o client-binary ./cmd/calc-cli