.PHONY: run test

run:
	go run cmd/tag/main.go

test:
	go test -v ./...
