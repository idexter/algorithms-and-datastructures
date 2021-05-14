
all: lint test

lint:
	golint ./...

test:
	go test -race -v ./...
