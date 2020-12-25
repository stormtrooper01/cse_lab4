package commands

import (
	"strings"
	"github.com/stormtrooper01/cse_lab4/engine"
)

type printcCommand struct {
	Count  int
	Symbol string
}

func (p *printcCommand) Execute(loop engine.Handler) {
	repeated := strings.Repeat(p.Symbol, p.Count)
	loop.Post(&printCommand{Arg: repeated})
}
