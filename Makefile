build:
	@go build -o bin/quiz

run: build
	@./bin/quiz

test:
	@go test -v ./...
