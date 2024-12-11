.PHONY: all clean run-%

# Go commands
GO = go

# Directory structure
SRC_DIR = src
INPUT_DIR = inputs

# Find all Go files in src directory
SOLUTIONS := $(wildcard $(SRC_DIR)/day*.go)

# Default target runs all days
all: $(SOLUTIONS)
	@for day in $(wildcard $(SRC_DIR)/day*.go); do \
		echo "\nRunning $${day}..."; \
		$(GO) run $(SRC_DIR)/*.go -day=$$(basename $${day} .go); \
	done

# Run a specific day (usage: make run-01 or make run-1 for day 1)
run-%: 
	@DAY=day$(shell printf %02d $(patsubst run-%,%,$@)); \
	if [ -f "$(SRC_DIR)/$$DAY.go" ]; then \
		echo "Running $$DAY..."; \
		$(GO) run $(SRC_DIR)/*.go -day=$$DAY; \
	else \
		echo "Error: Source file $$DAY.go does not exist"; \
		exit 1; \
	fi

# Clean any generated files
clean:
	$(GO) clean
	rm -f $(SRC_DIR)/*.test

# Help target
help:
	@echo "Available targets:"
	@echo "  make all       - Run all days"
	@echo "  make run-N     - Run specific day (e.g., make run-1 or make run-01)"
	@echo "  make clean     - Clean generated files" 