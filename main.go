package main

import (
	"encoding/json"
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> e8d0360 (fixed main.go)
	"flag"
	"fmt"
	"os"

<<<<<<< HEAD
=======
	"fmt"
	"os"
>>>>>>> dd8b44d (first try)
=======
>>>>>>> e8d0360 (fixed main.go)
	"github.com/nimamakhmali/turing-machine-go/turing"
)

func main() {
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> e8d0360 (fixed main.go)
	// Parse command line arguments
	inputFile := flag.String("input", "input/example_tm.json", "Input JSON file for Turing machine")
	flag.Parse()

	// Open and read the JSON file
	file, err := os.Open(*inputFile)
<<<<<<< HEAD
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

	result, tape, steps := tm.Run(1000)

	fmt.Println("=== Results ===")
	fmt.Printf("Result: %s\n", result)
	fmt.Printf("Final tape: %v\n", tape)
	fmt.Printf("Steps taken: %d\n", steps)
=======
	file, err := os.Open("input/example_tm.json")
=======
>>>>>>> e8d0360 (fixed main.go)
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

	result, tape, steps := tm.Run(1000)

<<<<<<< HEAD
	fmt.Println("Result:", result)
	fmt.Println("Final Tape:", tape)
	fmt.Println("Steps:", steps)
>>>>>>> dd8b44d (first try)
=======
	fmt.Println("=== Results ===")
	fmt.Printf("Result: %s\n", result)
	fmt.Printf("Final tape: %v\n", tape)
	fmt.Printf("Steps taken: %d\n", steps)
>>>>>>> e8d0360 (fixed main.go)
}
