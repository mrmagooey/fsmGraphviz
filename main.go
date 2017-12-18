package fsmGraphviz

import (
	"fmt"

	"github.com/looplab/fsm"
)

// CreateGraphvizString create a graphviz string from events
func CreateGraphvizString(startState string, events fsm.Events, fsmName string) (outputString string, err error) {
	// get all the states from the events
	states := map[string]bool{}
	for _, event := range events {
		for _, srcState := range event.Src {
			states[srcState] = true
		}
		states[event.Dst] = true
	}
	// start making string
	preamble := "digraph finite_state_machine {\n" +
		"labelloc=\"t\";\n" +
		fmt.Sprintf("label=\"%s\";\n", fsmName) +
		"rankdir=LR;\n" +
		"size=\"8,5\";\n\n" +
		"node [shape = point] qi;\n"
	outputString += preamble
	for state := range states {
		outputString += fmt.Sprintf("node [shape = circle, label = \"%s\", fontsize = 12] %s;\n", state, state)
	}

	// add the starting event
	outputString += fmt.Sprintf("qi -> %s; \n", startState)
	// add each event
	for _, event := range events {
		for _, src := range event.Src {
			outputString += fmt.Sprintf("%s -> %s [ label = \"%s\" ];\n", src, event.Dst, event.Name)
		}
	}

	postamble := "\n}"
	outputString += postamble
	return
}
