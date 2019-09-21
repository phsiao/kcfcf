
all: fmt vet

fmt:
	@echo "==== go fmt ==="
	@go fmt ./...

vet:
	@echo "==== go vet ==="
	@go vet ./...
