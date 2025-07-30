package turing

import "strings"

type TuringMachine struct {
	Definition TuringMachineDefinition
	Tape       []string
	Head       int
	State      string
}

func NewTuringMachine(def TuringMachineDefinition) *TuringMachine {
	return &TuringMachine{
		Definition: def,
		Tape:       strings.Split(def.Tape, ""),
		Head:       def.HeadPosition,
		State:      def.StartState,
	}
}

func (tm *TuringMachine) Step() bool {
	if tm.Head < 0 {
		tm.Tape = append([]string{"_"}, tm.Tape...)
		tm.Head = 0
	}
	if tm.Head >= len(tm.Tape) {
		tm.Tape = append(tm.Tape, "_")
	}

	key := tm.State + "," + tm.Tape[tm.Head]

	transition, ok := tm.Definition.Transitions[key]
	if !ok {
		tm.State = tm.Definition.RejectState
		return false
	}

	tm.Tape[tm.Head] = transition.Write
	tm.State = transition.NewState

	switch transition.Direction {
	case "R":
		tm.Head++
	case "L":
		tm.Head--
	case "N":
		// no move
	}

	return true
}

func (tm *TuringMachine) Run(maxSteps int) (string, []string, int) {
	steps := 0
	for tm.State != tm.Definition.AcceptState &&
		tm.State != tm.Definition.RejectState &&
		steps < maxSteps {
		tm.Step()
		steps++
	}

	result := "Rejected"
	if tm.State == tm.Definition.AcceptState {
		result = "Accepted"
	}
	return result, tm.Tape, steps
}
