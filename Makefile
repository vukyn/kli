PRJ=

build:
	@echo "Building $(PRJ)..."
	@go build -o bin/ ./$(PRJ)
	@echo "Build complete"

install:
	@echo "Installing $(PRJ)..."
	@go install ./$(PRJ)
	@echo "Install complete"