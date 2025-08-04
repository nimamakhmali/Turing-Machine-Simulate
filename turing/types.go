package turing

type Transition struct {
	NewState  string `json:"new_state"`
	Write     string `json:"write"`
	Direction string `json:"direction"`
}

type TuringMachineDefinition struct {
	States        []string              `json:"states"`
	InputAlphabet []string              `json:"input_alphabet"`
	TapeAlphabet  []string              `json:"tape_alphabet"`
	StartState    string                `json:"start_state"`
	AcceptState   string                `json:"accept_state"`
	RejectState   string                `json:"reject_state"`
	Transitions   map[string]Transition `json:"transitions"`
	Tape          string                `json:"tape"`
	HeadPosition  int                   `json:"head_position"`
}
