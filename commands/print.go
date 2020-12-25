package commands

import (
	"fmt"

	"github.com/stormtrooper01/cse_lab4/engine"
)

type printCommand struct {
	Arg string
}

func (p *printCommand) Execute(loop engine.Handler) {
	fmt.Println(p.Arg)
}
