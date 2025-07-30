package test

import (
	"testing"

	"github.com/nimamakhmali/turing-machine-go/turing"
)

func TestTuringMachineBasic(t *testing.T) {
	// Simple TM that accepts strings ending with 1
	def := turing.TuringMachineDefinition{
		States:        []string{"q0", "q1", "q_accept", "q_reject"},
		InputAlphabet: []string{"0", "1"},
		TapeAlphabet:  []string{"0", "1", "_"},
		StartState:    "q0",
		AcceptState:   "q_accept",
		RejectState:   "q_reject",
		Transitions: map[string]turing.Transition{
			"q0,0": {NewState: "q0", Write: "0", Direction: "R"},
			"q0,1": {NewState: "q1", Write: "1", Direction: "R"},
			"q1,0": {NewState: "q_reject", Write: "0", Direction: "N"},
			"q1,1": {NewState: "q1", Write: "1", Direction: "R"},
			"q1,_": {NewState: "q_accept", Write: "_", Direction: "N"},
		},
		Tape:         "001",
		HeadPosition: 0,
	}

	tm := turing.NewTuringMachine(def)
	result, _, steps := tm.Run(100)

	if result != "Accepted" {
		t.Errorf("Expected Accepted, got %s", result)
	}

	if steps == 0 {
		t.Error("Expected steps > 0")
	}
}

func TestTuringMachineReject(t *testing.T) {
	// Simple TM that rejects strings ending with 0
	def := turing.TuringMachineDefinition{
		States:        []string{"q0", "q1", "q_accept", "q_reject"},
		InputAlphabet: []string{"0", "1"},
		TapeAlphabet:  []string{"0", "1", "_"},
		StartState:    "q0",
		AcceptState:   "q_accept",
		RejectState:   "q_reject",
		Transitions: map[string]turing.Transition{
			"q0,0": {NewState: "q0", Write: "0", Direction: "R"},
			"q0,1": {NewState: "q1", Write: "1", Direction: "R"},
			"q1,0": {NewState: "q_reject", Write: "0", Direction: "N"},
			"q1,1": {NewState: "q1", Write: "1", Direction: "R"},
			"q1,_": {NewState: "q_accept", Write: "_", Direction: "N"},
		},
		Tape:         "000",
		HeadPosition: 0,
	}

	tm := turing.NewTuringMachine(def)
	result, _, _ := tm.Run(100)

	if result != "Rejected" {
		t.Errorf("Expected Rejected, got %s", result)
	}
}

func TestTuringMachineInfiniteLoop(t *testing.T) {
	// TM that goes into infinite loop
	def := turing.TuringMachineDefinition{
		States:        []string{"q0", "q_accept", "q_reject"},
		InputAlphabet: []string{"0", "1"},
		TapeAlphabet:  []string{"0", "1", "_"},
		StartState:    "q0",
		AcceptState:   "q_accept",
		RejectState:   "q_reject",
		Transitions: map[string]turing.Transition{
			"q0,0": {NewState: "q0", Write: "0", Direction: "R"},
			"q0,1": {NewState: "q0", Write: "1", Direction: "R"},
		},
		Tape:         "000",
		HeadPosition: 0,
	}

	tm := turing.NewTuringMachine(def)
	result, _, steps := tm.Run(10) // Limited steps

	// Check that it either reaches max steps or gets rejected
	if steps < 3 {
		t.Errorf("Expected at least 3 steps, got %d", steps)
	}

	if result != "Rejected" {
		t.Errorf("Expected Rejected (due to max steps or no transition), got %s", result)
	}
}

func TestTuringMachineTapeExpansion(t *testing.T) {
	// TM that moves left and right to test tape expansion
	def := turing.TuringMachineDefinition{
		States:        []string{"q0", "q1", "q_accept", "q_reject"},
		InputAlphabet: []string{"0", "1"},
		TapeAlphabet:  []string{"0", "1", "_"},
		StartState:    "q0",
		AcceptState:   "q_accept",
		RejectState:   "q_reject",
		Transitions: map[string]turing.Transition{
			"q0,0": {NewState: "q1", Write: "X", Direction: "R"},
			"q1,0": {NewState: "q1", Write: "0", Direction: "R"},
			"q1,1": {NewState: "q1", Write: "1", Direction: "R"},
			"q1,_": {NewState: "q_accept", Write: "_", Direction: "N"},
		},
		Tape:         "0",
		HeadPosition: 0,
	}

	tm := turing.NewTuringMachine(def)
	result, tape, _ := tm.Run(100)

	if result != "Accepted" {
		t.Errorf("Expected Accepted, got %s", result)
	}

	// Check that tape was properly expanded
	if len(tape) < 2 {
		t.Errorf("Expected tape to be expanded, got length %d", len(tape))
	}
}
