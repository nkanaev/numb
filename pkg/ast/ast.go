package ast

import (
	"fmt"
	"strings"

	"github.com/nkanaev/numb/pkg/funcs"
	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/unit"
	"github.com/nkanaev/numb/pkg/value"
)

type Node interface {
	Eval(map[string]value.Value) value.Value
	String() string
}

type BinOP struct {
	Lhs, Rhs Node
	Op       token.Token
	Implicit bool
}

func (n *BinOP) String() string {
	if n.Implicit {
		return fmt.Sprintf("%s %s", n.Lhs.String(), n.Rhs.String())
	}
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

type Assign struct {
	Name string
	Expr Node
}

func (n *Assign) Eval(env map[string]value.Value) value.Value {
	if _, ok := value.Consts[n.Name]; ok {
		panic("cannot assign to const " + n.Name)
	}
	val := n.Expr.Eval(env)
	env[n.Name] = val
	return val
}

func (n *Assign) String() string {
	return n.Name + " = " + n.Expr.String()
}

type Var struct {
	Name string
}

func (n *Var) Eval(env map[string]value.Value) value.Value {
	if val, ok := value.Consts[n.Name]; ok {
		return val
	}
	if unit := unit.Get(n.Name); unit != nil {
		return value.NewInt(1).WithUnit(unit)
	}
	val, ok := env[n.Name]
	if !ok {
		panic(n.Name + " not defined")
	}
	return val
}

func (n *Var) String() string {
	return n.Name
}

type Format struct {
	Expr Node
	Fmt  value.NumeralSystem
}

func (n *Format) Eval(env map[string]value.Value) value.Value {
	return n.Expr.Eval(env).As(n.Fmt)
}

func (n *Format) String() string {
	return n.Expr.String() + " as " + n.Fmt.String()
}

type Unit struct {
	Expr Node
	Unit *unit.Unit
}

func (n *Unit) Eval(env map[string]value.Value) value.Value {
	expr := n.Expr.Eval(env)
	if expr.Unit != nil {
		panic(fmt.Sprintf("cannot set unit `%s` to `%s`", n.Unit, n.Expr))
	}
	return expr.WithUnit(n.Unit)
}

func (n *Unit) String() string {
	return n.Expr.String() + " " + n.Unit.String()
}

type Convert struct {
	Expr, Unit Node
}

func (n *Convert) Eval(env map[string]value.Value) value.Value {
	u := n.Unit.Eval(env)
	if u.Unit == nil {
		panic(n.Unit.String() + " is not a unit")
	}
	if u.Num.IsInt() && u.Num.Num().Int64() != 1 {
		panic("cannot convert to a unit with a value: " + n.Unit.String())
	}
	return n.Expr.Eval(env).To(u.Unit)
}

func (n *Convert) String() string {
	return n.Expr.String() + " to " + n.Unit.String()
}

type FunCall struct {
	Name string
	Args []Node
}

func (n *FunCall) Eval(env map[string]value.Value) value.Value {
	fn := funcs.Get(n.Name)
	if fn == nil {
		panic("no such function: " + n.Name)
	}
	args := make([]value.Value, len(n.Args))
	for i, arg := range n.Args {
		args[i] = arg.Eval(env)
	}
	return (*fn)(args...)
}

func (n *FunCall) String() string {
	args := make([]string, len(n.Args))
	for i, arg := range n.Args {
		args[i] = arg.String()
	}
	return n.Name + "(" + strings.Join(args, ", ") + ")"
}
