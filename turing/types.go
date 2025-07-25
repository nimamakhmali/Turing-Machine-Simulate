#Nima
package types

type Transition struct {
    NewState   string `json:"new_state"`
    Write      string `json:"write"`
    Direction  string `json:"direction"`
}

type TuringMachine struct {
    States         []string              `json:"states"`
    InputAlphabet  []string              `json:"input_alphabet"`
    TapeAlphabet   []string              `json:"tape_alphabet"`
	StartState     string                `json:"start_state"`
	AcceptStates   []string              `json:"accept_states"`
	RejectStates   []string              `json:"reject_states"`
	Transitions    map[string]Transition `json:"transitions"`
    Tape           []string              `json:"tape"`
	HeadPosition   int                   `json:"head_position"`
}