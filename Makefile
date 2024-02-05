build:
	@echo "Building..."
	@go build -o bin/ ./
	@echo "Build complete"

install:
	@echo "Installing..."
	@go install
	@echo "Install complete"