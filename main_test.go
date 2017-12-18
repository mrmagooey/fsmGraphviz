package fsmGraphviz

import (
	"testing"

	"github.com/looplab/fsm"
)

func TestBasic(t *testing.T) {
	_, err := CreateGraphvizString(
		"2",
		fsm.Events{
			fsm.EventDesc{
				Name: "a",
				Src:  []string{"1", "2"},
				Dst:  "1",
			},
			fsm.EventDesc{
				Name: "b",
				Src:  []string{"3"},
				Dst:  "2",
			},
		},
		"Blah")
	if err != nil {
		t.Error(err)
		return
	}
}
