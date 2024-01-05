build:
	@go build -o bin/snippet cmd/app/main.go

run: build
	@./bin/snippet

test:
	@go test -v ./...
