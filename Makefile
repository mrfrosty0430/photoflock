.PHONY: build run test clean

build:
	go build -o bin/api cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test ./...

lint:
	golangci-lint run

clean:
	rm -rf bin/
