# Thank you ChatGPT

# Define the list of all days
DAYS := $(shell seq 1 25)

# Default target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make dayX    - Run the Go program for day X (e.g., day1, day2, etc.)"
	@echo "  make all     - Run the Go programs for all days"

# Rule to run a specific day's Go program
.PHONY: $(foreach day,$(DAYS),day$(day))
$(foreach day,$(DAYS),day$(day)):
	@cd ./day$(@:day%=%) && go run .

# Run all days sequentially
.PHONY: all
all:
	@for day in $(DAYS); do \
		if [ -d ./day$$day ]; then \
			echo "Running Day $$day"; \
			(cd ./day$$day && go run .); \
		fi \
	done

