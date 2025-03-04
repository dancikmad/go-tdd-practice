# List of known test directories
DIRS := helloworld integers

# Test specific directory or all
test:
	@if [ -z "$(DIR)" ]; then \
		go test ./... -v; \
	elif echo "$(DIRS)" | grep -q "$(DIR)"; then \
		go test ./$(DIR) -v; \
	else \
		echo "Unknown directory: $(DIR)"; \
		exit 1; \
	fi

# List available test directories
list:
	@echo "Available test directories:"
	@for dir in $(DIRS); do echo "  - $$dir"; done
