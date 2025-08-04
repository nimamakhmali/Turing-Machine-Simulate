package turing

<<<<<<< HEAD
<<<<<<< HEAD
import "strings"

type TuringMachine struct {
	Definition TuringMachineDefinition
	Tape       []string
	Head       int
	State      string
=======
//import "strings"

type TuringMachine struct {
	Definition TuringMachineDefinition
	Tape      []string
	Head      int
	State     string
>>>>>>> dd8b44d (first try)
=======
import "strings"

type TuringMachine struct {
	Definition TuringMachineDefinition
	Tape       []string
	Head       int
	State      string
>>>>>>> c18322f (edit functions)
}

func NewTuringMachine(def TuringMachineDefinition) *TuringMachine {
	return &TuringMachine{
		Definition: def,
<<<<<<< HEAD
<<<<<<< HEAD
		Tape:       strings.Split(def.Tape, ""),
		Head:       def.HeadPosition,
		State:      def.StartState,
=======
		Tape:      make([]string, 0),
		Head:      0,
		State:     "initial",
>>>>>>> dd8b44d (first try)
=======
		Tape:       strings.Split(def.Tape, ""),
		Head:       def.HeadPosition,
		State:      def.StartState,
>>>>>>> c18322f (edit functions)
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

<<<<<<< HEAD
<<<<<<< HEAD
	key := tm.State + "," + tm.Tape[tm.Head]
=======
	currentSymbol := tm.Tape[tm.Head]
	key := tm.State + "," + currentSymbol
>>>>>>> dd8b44d (first try)
=======
	key := tm.State + "," + tm.Tape[tm.Head]
>>>>>>> c18322f (edit functions)

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

<<<<<<< HEAD
<<<<<<< HEAD
=======

>>>>>>> dd8b44d (first try)
=======
>>>>>>> c18322f (edit functions)
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
