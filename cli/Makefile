# CLI Makefile
BINARY_NAME = cli
BUILD_DIR = build
CMD_DIR = ./cmd

# Colors
BLUE = \033[34m
GREEN = \033[32m
NC = \033[0m # No Color

.PHONY: all build clean help test

all: build

build:
	@printf "$(BLUE)→ Building $(BINARY_NAME)...$(NC)\n"
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)
	@printf "$(GREEN)✓ Build completed: $(BUILD_DIR)/$(BINARY_NAME)$(NC)\n"

clean:
	@printf "$(BLUE)→ Cleaning build directory...$(NC)\n"
	rm -rf $(BUILD_DIR)
	@printf "$(GREEN)✓ Clean completed$(NC)\n"

test:
	@printf "$(BLUE)→ Running tests...$(NC)\n"
	go test -v ./...
	@printf "$(GREEN)✓ Tests completed$(NC)\n"

install: build
	@printf "$(BLUE)→ Installing $(BINARY_NAME)...$(NC)\n"
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@printf "$(GREEN)✓ Installation completed$(NC)\n"

run: build
	@printf "$(BLUE)→ Running $(BINARY_NAME)...$(NC)\n"
	./$(BUILD_DIR)/$(BINARY_NAME)

help:
	@printf "$(BLUE)→ Available targets:$(NC)\n"
	@echo "  build    - Build the CLI binary"
	@echo "  clean    - Remove build artifacts"
	@echo "  test     - Run tests"
	@echo "  install  - Install binary to /usr/local/bin (requires sudo)"
	@echo "  run      - Build and run the CLI"
	@echo "  help     - Show this help message"