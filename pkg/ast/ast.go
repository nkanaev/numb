package ast

import (
	"fmt"

	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/value"
)

type Node interface {
	Eval(map[string]value.Value) value.Value
	String() string
}

type BinOP struct {
	Lhs, Rhs Node
	Op token.Token
}

func (n *BinOP) String() string {
	return fmt.Sprintf("%s %s %s", n.Lhs.String(), n.Op.String(), n.Rhs.String())
}

func (n *BinOP) Eval(env map[string]value.Value) value.Value {
	switch n.Op {
	case token.ADD:
		return n.Lhs.Eval(env).Add(n.Rhs.Eval(env))
	case token.SUB:
		return n.Lhs.Eval(env).Sub(n.Rhs.Eval(env))
	case token.MUL:
		return n.Lhs.Eval(env).Mul(n.Rhs.Eval(env))
	case token.QUO:
		return n.Lhs.Eval(env).Quo(n.Rhs.Eval(env))
	case token.SHL:
		return n.Lhs.Eval(env).Lsh(n.Rhs.Eval(env))
	case token.SHR:
		return n.Lhs.Eval(env).Rsh(n.Rhs.Eval(env))
	}
	return value.NewInt(0)
}

type ParenExpr struct {
	Expr Node	
}

func (n *ParenExpr) Eval(env map[string]value.Value) value.Value {
	return n.Expr.Eval(env)
}

func (n *ParenExpr) String() string {
	return "(" + n.Expr.String() + ")"
}
