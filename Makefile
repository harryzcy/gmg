.PHONY: install lint test

COMMIT := $(shell git rev-parse --short HEAD)
FLAGS := -ldflags "-X go.zcy.dev/gmg/cmd.commit=$(COMMIT)"

install:
	go install $(FLAGS)

lint:
	golangci-lint run -v ./...

test: lint
	go test -v ./...
