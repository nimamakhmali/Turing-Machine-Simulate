# Phase 1: Core Turing Machine Simulator

This is the core implementation of the Turing Machine Simulator in Go.

## Structure

```
phase1-core/
├── main.go              # CLI application entry point
├── turing/              # Core Turing machine logic
│   ├── machine.go       # Machine implementation
│   ├── types.go         # Data structures
│   └── utils.go         # Utility functions
├── input/               # Example configurations
│   ├── example_tm.json  # Simple example
│   └── palindrome_tm.json # Palindrome checker
├── test/                # Unit tests
│   └── machine_test.go  # Machine tests
├── go.mod               # Go module file
├── go.sum               # Go dependencies
├── README.md            # This file
└── LICENSE              # MIT License
```

## Features

- **Pure Go implementation** - No external dependencies
- **JSON configuration** - Easy to define machines
- **CLI interface** - Command line usage
- **Unit tests** - Comprehensive testing
- **Example machines** - Ready-to-use configurations

## Usage

```bash
# Run with example configuration
go run main.go -config input/example_tm.json -tape "101"

# Run palindrome checker
go run main.go -config input/palindrome_tm.json -tape "1001"

# Run tests
go test ./test
```

## Requirements

- Go 1.16 or higher

## Installation

```bash
go mod tidy
```

This phase contains the fundamental Turing machine logic that is used by both web and desktop versions.