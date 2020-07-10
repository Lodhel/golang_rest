.PHONY: build
build:
	go build -v /.cmd/book

.DEFAULT_GOAL := build