.PHONY: lint test install

lint:
	golangci-lint run -v ./...

test: lint
	go test -v ./...

install:
	go install
