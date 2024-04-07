all: fmt lint test

lint:
	golangci-lint run
fmt:
	go fmt ./...
test:
	go test ./...