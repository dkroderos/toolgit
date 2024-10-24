build:
	@go build -o bin/tg main.go

test:
	@go test -v ./...

run: build
	@./bin/tg