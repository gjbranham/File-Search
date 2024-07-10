.PHONY: build clean test

makefile_directory_path := $(realpath $(dir $(realpath $(lastword $(MAKEFILE_LIST)))))

build:
	cd $(makefile_directory_path)/cmd/text-finder && go build -o $(makefile_directory_path)/bin/text-finder

test:
	go test -v ./... -coverprofile=.cover.out && go tool cover -func=.cover.out