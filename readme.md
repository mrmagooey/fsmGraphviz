# fsmGraphviz

A go library for creating graphviz diagrams using data structures from <https://github.com/looplab/fsm>.

You need graphviz installed to actually create the diagrams, this just creates the graphviz information.

<img src="myFsm.png" alt="FSM Example" width="500px" height="200px"/>

## Installation

    go get github.com/mrmagooey/fsmGraphviz
    
Or,

    dep ensure -add github.com/mrmagooey/fsmGraphviz

## Usage

[Docs at godoc](https://godoc.org/github.com/mrmagooey/fsmGraphviz).

The suggestion for usage is to add a new package to your program similar to the `usage/` folder in this repository. This package will generate a small, standalone binary that prints the graphviz "dot-file" information to stdout. You can then pipe this output directly to graphviz. `go generate` can be used to streamline this process:

    //go:generate bash -c "go build && ./usage | dot -Tpng -Gdpi=300 -o myFsm.png"

The `usage/` folder is an example of this approach.

