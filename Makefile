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
	@for solution in $(SOLUTIONS); do \
		echo "\nRunning $$solution..."; \
		$(GO) run $$solution; \
	done

# Run a specific day (usage: make run-01 or make run-1 for day 1)
run-%: 
	@DAY_FILE=$(SRC_DIR)/day$(shell printf %02d $(patsubst run-%,%,$@)).go; \
	echo "Looking for file: $$DAY_FILE"; \
	if [ -f "$$DAY_FILE" ]; then \
		echo "Running Day $(patsubst run-%,%,$@)..."; \
		$(GO) run "$$DAY_FILE"; \
	else \
		echo "Error: Source file $$DAY_FILE does not exist"; \
		ls -l $(SRC_DIR)/; \
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