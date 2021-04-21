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
	Op       token.Token
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
	case token.AND:
		return n.Lhs.Eval(env).And(n.Rhs.Eval(env))
	case token.OR:
		return n.Lhs.Eval(env).Or(n.Rhs.Eval(env))
	case token.XOR:
		return n.Lhs.Eval(env).Xor(n.Rhs.Eval(env))
	case token.REM:
		return n.Lhs.Eval(env).Rem(n.Rhs.Eval(env))
	case token.EXP:
		return n.Lhs.Eval(env).Exp(n.Rhs.Eval(env))
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

type Unary struct {
	Op   token.Token
	Expr Node
}

func (n *Unary) Eval(env map[string]value.Value) value.Value {
	if n.Op == token.SUB {
		return n.Expr.Eval(env).Neg()
	}
	return n.Expr.Eval(env)
}

func (n *Unary) String() string {
	return n.Op.String() + n.Expr.String()
}
