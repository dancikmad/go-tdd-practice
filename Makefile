# List of known test directories
DIRS := helloworld integers iterations arrays

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

test-benchmark:
	@if [ -z "$(DIR)" ]; then \
		go test -bench=. ./... -v; \
	elif echo "$(DIRS)" | grep -q "$(DIR)"; then \
		go test -bench=. ./$(DIR) -v; \
	else \
		echo "Unknown directory: $(DIR)"; \
		exit 1; \
	fi

# List available test directories
list:
	@echo "Available test directories:"
	@for dir in $(DIRS); do echo "  - $$dir"; done

clean:
	@go clean 
	@rm -f coverage.out
	@rm -rf *.test 
	@find . -name "*.coverprofile" -delete
