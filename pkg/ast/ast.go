package ast

import (
	"fmt"

	"github.com/nkanaev/numb/pkg/scanner"
	"github.com/nkanaev/numb/pkg/value"
)

type Node interface {
	Eval(map[string]value.Value) value.Value
	String() string
}

type Const struct {
	Val value.Value
}

type BinOP struct {
	Lhs, Rhs Node
	Op scanner.Token	
}

func (n *BinOP) String() string {
	return fmt.Sprintf("%s %s %s", n.Lhs.String(), n.Op.String(), n.Rhs.String())
}

func (n *BinOP) Eval(env map[string]value.Value) value.Value {
	switch n.Op {
	case scanner.ADD:
		return n.Lhs.Eval(env).Add(n.Rhs.Eval(env))
	}
	return value.NewInt(0)
}
