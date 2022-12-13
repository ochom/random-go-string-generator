build:
	@echo "Building..."
	@go build -o ./cmd/main .
	
test:
	@echo "Testing..."
	@go test ./...
