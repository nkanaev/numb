package ast

import (
	"fmt"
	"strconv"

	"github.com/nkanaev/numb/pkg/scanner"
)

type Node interface {
	Eval(map[string]float64) float64
	String() string
}

type Const struct {
	Val float64
}

func (n *Const) Eval(map[string]float64) float64 {
	return n.Val
}

func (n *Const) String() string {
	return strconv.FormatFloat(n.Val, 'f', -1, 64)
}

type BinOP struct {
	Lhs, Rhs Node
	Op scanner.Token	
}

func (n *BinOP) String() string {
	return fmt.Sprintf("%s %s %s", n.Lhs.String(), n.Op.String(), n.Rhs.String())
}

func (n *BinOP) Eval(map[string]float64) float64 {
	return 0
}
