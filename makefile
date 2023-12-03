# Simple Makefile for a Go project

# Build the application
all: build

build:
	@go build main.go

new-year:
	@echo "Creating Year"
	@cobra-cli add 20$(YEAR)
	@mkdir cmd/20$(YEAR)
	@mkdir input/20$(YEAR)
	@touch input/20$(YEAR)/.keep

new-day:
	@cobra-cli add day$(DAY) -p 20$(YEAR)
	@mv cmd/day$(DAY).go cmd/20$(YEAR)/


# Run the application
run:
	@go run main.go 20$(YEAR) day$(DAY)

.PHONY: all build run 
