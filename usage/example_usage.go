package main

import (
	"fmt"

	"github.com/looplab/fsm"
	"github.com/mrmagooey/fsm-graphviz"
)

//go:generate bash -c "go build && ./usage | dot -Tpng -Gdpi=300 -o myFsm.png"

func main() {
	event1 := "Install FSM Graphviz"
	event2 := "Start New Project"
	state1 := "No Graphviz Diagrams"
	state2 := "Have Graphviz Diagrams"
	dot, _ := fsmGraphviz.CreateGraphvizString("No Graphviz Diagrams", fsm.Events{
		{
			Name: event1,
			Src:  []string{state1},
			Dst:  state2,
		},
		{
			Name: event2,
			Src:  []string{state2},
			Dst:  state1,
		},
	},
		"My FSM diagram",
		"LR",
	)
	fmt.Println(dot)
}
