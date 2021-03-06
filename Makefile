.PHONY: build start test test-integration



cover:
	go tool cover -html=cover.out

start:
	go run ./main.go

test:
	go test -coverprofile=cover.out -short ./...

test-integration:
	go test -coverprofile=cover.out -p 1 ./...