#! /usr/bin/make

all: clean generate

clean:
	@rm -rf ./gen

generate:
	goa gen github.com/shamhub/goa-example/api