build:
	@go build -o bin/transforms

run: build
	@./bin/transforms

test:
	@go test -v ./...
