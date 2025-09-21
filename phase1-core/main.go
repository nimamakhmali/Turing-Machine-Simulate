package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/nimamakhmali/turing-machine-go/turing"
)

func main() {
	// Parse command line arguments
	inputFile := flag.String("input", "input/complex_calculator.json", "Input JSON file for Turing machine")
	flag.Parse()

	// Open and read the JSON file
	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", *inputFile, err)
		os.Exit(1)
	}
	defer file.Close()

	// Parse the JSON
	var def turing.TuringMachineDefinition
	if err := json.NewDecoder(file).Decode(&def); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Create and run the Turing machine
	tm := turing.NewTuringMachine(def)

	fmt.Println("=== Turing Machine Simulator ===")
	fmt.Printf("Input file: %s\n", *inputFile)
	fmt.Printf("Initial tape: %s\n", def.Tape)
	fmt.Printf("Initial head position: %d\n", def.HeadPosition)
	fmt.Printf("Start state: %s\n", def.StartState)
	fmt.Printf("Accept state: %s\n", def.AcceptState)
	fmt.Printf("Reject state: %s\n", def.RejectState)
	fmt.Println("---")

	// Display initial configuration
	fmt.Println("=== Initial Configuration ===")
	displayTape(tm, 0)
	fmt.Println()

	// Run step by step with display
	fmt.Println("=== Step-by-Step Execution ===")
	step := 1
	for tm.State != tm.Definition.AcceptState &&
		tm.State != tm.Definition.RejectState &&
		step <= 1000 {

		fmt.Printf("--- Step %d ---\n", step)
		fmt.Printf("Current State: %s\n", tm.State)
		fmt.Printf("Head Position: %d\n", tm.Head)

		// Show current tape before step
		displayTape(tm, step-1)

		// Execute one step
		tm.Step()

		// Show tape after step
		fmt.Printf("After Step %d:\n", step)
		displayTape(tm, step)
		fmt.Println()

		step++
	}

	result := "Rejected"
	if tm.State == tm.Definition.AcceptState {
		result = "Accepted"
	}

	fmt.Println("=== Final Results ===")
	fmt.Printf("Result: %s\n", result)
	fmt.Printf("Final tape: %v\n", tm.Tape)
	fmt.Printf("Total steps taken: %d\n", step-1)
}

// Helper function to display tape with head position
func displayTape(tm *turing.TuringMachine, stepNum int) {
	fmt.Printf("Tape [Step %d]: ", stepNum)
	for i, symbol := range tm.Tape {
		if i == tm.Head {
			fmt.Printf("[%s]", symbol) // Highlight head position
		} else {
			fmt.Printf(" %s ", symbol)
		}
	}
	fmt.Println()
}
