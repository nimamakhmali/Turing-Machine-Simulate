 # Turing Machine Simulator in Go

A complete Turing machine simulator written in Go for the Theory of Languages and Automata course.

## Overview

This project implements a Turing machine simulator that can:
- Accept any Turing machine description in JSON format
- Simulate the computation step by step
- Handle infinite tape expansion (both left and right)
- Support all standard Turing machine operations (read, write, move left/right/stay)
- Provide detailed output of the computation process

## Features

- **Flexible Input**: Accepts Turing machine definitions in JSON format
- **Infinite Tape**: Automatically expands the tape as needed
- **Step-by-step Simulation**: Shows the computation process
- **Multiple Examples**: Includes various Turing machine examples
- **Comprehensive Testing**: Includes unit tests for various scenarios
- **Command Line Interface**: Easy to use with command line arguments

## Project Structure

```
Turing-Machine-Go/
├── main.go                 # Main application entry point
├── turing/
│   ├── machine.go         # Turing machine implementation
│   ├── types.go           # Data structures and types
│   └── utils.go           # Utility functions
├── input/
│   ├── example_tm.json    # Simple example Turing machine
│   └── palindrome_tm.json # Palindrome recognition TM
├── test/
│   └── machine_test.go    # Unit tests
└── output/                # Output directory for results
```

## Usage

### Basic Usage

```bash
# Run with default example
go run main.go

# Run with specific input file
go run main.go -input input/palindrome_tm.json
```

### Command Line Options

- `-input`: Specify the input JSON file (default: `input/example_tm.json`)

## JSON Format

The simulator accepts Turing machine definitions in the following JSON format:

```json
{
  "states": ["q0", "q1", "q_accept", "q_reject"],
  "input_alphabet": ["0", "1"],
  "tape_alphabet": ["0", "1", "X", "_"],
  "start_state": "q0",
  "accept_state": "q_accept",
  "reject_state": "q_reject",
  "transitions": {
    "q0,0": {"new_state": "q1", "write": "X", "direction": "R"},
    "q1,0": {"new_state": "q1", "write": "0", "direction": "R"},
    "q1,1": {"new_state": "q1", "write": "1", "direction": "R"},
    "q1,_": {"new_state": "q_accept", "write": "_", "direction": "N"}
  },
  "tape": "0011",
  "head_position": 0
}
```

### Transition Format

Transitions are defined as `"current_state,symbol"` keys with the following values:
- `new_state`: The next state
- `write`: The symbol to write on the tape
- `direction`: Movement direction ("L" for left, "R" for right, "N" for no movement)

## Examples

### 1. Simple Example (example_tm.json)
This Turing machine recognizes strings that contain at least one '0' followed by any number of symbols.

### 2. Palindrome Recognition (palindrome_tm.json)
This Turing machine recognizes palindromes over the alphabet {0, 1}.

## Running Tests

```bash
go test ./test/
```

## Implementation Details

### Core Components

1. **TuringMachine**: The main simulator class
2. **TuringMachineDefinition**: Data structure for TM description
3. **Transition**: Represents a single transition rule

### Key Features

- **Infinite Tape**: The tape automatically expands in both directions
- **State Management**: Proper handling of accept/reject states
- **Step Limiting**: Prevents infinite loops with maximum step limit
- **Error Handling**: Robust error handling for invalid inputs

### Algorithm

1. **Initialization**: Load TM definition and set initial state
2. **Execution Loop**: 
   - Read current symbol under head
   - Find applicable transition
   - Write new symbol
   - Move head
   - Update state
3. **Termination**: Stop when reaching accept/reject state or max steps

## Academic Context

This project demonstrates key concepts from the Theory of Languages and Automata course:

- **Turing Machine Model**: Complete implementation of the TM model
- **Computational Theory**: Understanding of computation and decidability
- **Formal Languages**: Recognition of formal languages
- **Algorithm Design**: Step-by-step algorithm implementation

## Requirements

- Go 1.16 or higher
- Standard library only (no external dependencies)

## Building and Running

```bash
# Build the project
go build -o turing-simulator main.go

# Run the executable
./turing-simulator -input input/example_tm.json
```

## Output Format

The simulator provides detailed output including:
- Initial configuration
- Final result (Accepted/Rejected)
- Final tape state
- Number of steps taken

## Extensions and Improvements

Potential improvements for future versions:
- Visual simulation with step-by-step display
- Support for multiple tapes
- Non-deterministic Turing machines
- Performance optimizations
- GUI interface

## License

This project is created for educational purposes as part of the Theory of Languages and Automata course.