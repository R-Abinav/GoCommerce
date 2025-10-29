build: 
	@go build -o bin/GoCommerce cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/GoCommerce