.DEFAULT_GOAL := build

.PHONY: fmt vet build

test:
	go test ./...

fmt:
	go fmt ./...

vet: fmt 
	go vet ./...

clean: vet 
	go clean

build: clean 
	go build

repl-build: clean 
	go build -o goextrack ./cmd/goextrack-repl/main.go 

repl-run: clean
	go run ./cmd/goextrack-repl/main.go
