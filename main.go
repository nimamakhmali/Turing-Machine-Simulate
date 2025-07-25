package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/nimamakhmali/turing-machine-go/turing"
)

func main() {
	file, err := os.Open("input/example_tm.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var def turing.TuringMachineDefinition
	if err := json.NewDecoder(file).Decode(&def); err != nil {
		panic(err)
	}

	tm := turing.NewTuringMachine(def)
	result, tape, steps := tm.Run(1000)

	fmt.Println("Result:", result)
	fmt.Println("Final Tape:", tape)
	fmt.Println("Steps:", steps)
}
