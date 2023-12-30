build:
	@go build -o bin/one_liner cmd/app/main.go

run: build
	@./bin/one_liner

test:
	@go test -v ./...
