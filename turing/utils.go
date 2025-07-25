package turing

import (
	"fmt"
	"strings"
)

func PrintTape(tape []string, head int) {
	var builder strings.Builder
	for i, symbol := range tape {
		if i == head {
			builder.WriteString("[" + symbol + "]")
		} else {
			builder.WriteString(" " + symbol + " ")
		}
	}
	fmt.Println("Tape:", builder.String())
}
