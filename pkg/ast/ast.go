package ast

import (
	"fmt"
	"strings"

	"github.com/nkanaev/numb/pkg/funcs"
	"github.com/nkanaev/numb/pkg/ratutils"
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
	lhs := n.Lhs.Eval(env)
	rhs := n.Rhs.Eval(env)
	val, err := lhs.BinOP(n.Op, rhs)
	if err != nil {
		panic(err)
	}
	return val
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
	val, err := n.Expr.Eval(env).UnOP(n.Op)
	if err != nil {
		panic(err)
	}
	return val
}

func (n *Unary) String() string {
	return n.Op.String() + n.Expr.String()
}

type Assign struct {
	Name string
	Expr Node
	Unit bool
}

func (n *Assign) Eval(env map[string]value.Value) value.Value {
	val := n.Expr.Eval(env)
	if n.Unit {
		if val, isnum := val.(value.Number); isnum {
			unit.Add(n.Name, val.Num, unit.Units{})
			return val
		}
		if val, isunit := val.(value.Unit); isunit {
			unit.Add(n.Name, val.Num, val.Units)
			return val
		}
		panic(fmt.Sprintf("cannot create unit from %s (%s type)", val, value.Type(val)))
	} else {
		env[n.Name] = val
	}
	return val
}

func (n *Assign) String() string {
	if n.Unit {
		return n.Name + " : " + n.Expr.String()
	}
	return n.Name + " = " + n.Expr.String()
}

type Var struct {
	Name string
}

func (n *Var) Eval(env map[string]value.Value) value.Value {
	if val, ok := env[n.Name]; ok {
		return val
	}
	if time := value.GetNamedTime(n.Name); time != nil {
		return time
	}
	if unit, ok := unit.Get(n.Name); ok {
		return value.Unit{Num: ratutils.ONE, Units: unit}
	}
	panic(n.Name + " not defined")
}

func (n *Var) String() string {
	return n.Name
}

type Format struct {
	Expr Node
	Fmt  string
}

func (n *Format) Eval(env map[string]value.Value) value.Value {
	val, err := n.Expr.Eval(env).In(n.Fmt)
	if err != nil {
		panic(err)
	}
	return val
}

func (n *Format) String() string {
	return fmt.Sprintf("%s %s %s", n.Expr.String(), token.IN, n.Fmt)
}

/*
type Convert struct {
	Expr, Unit Node
}

func (n *Convert) Eval(env map[string]value.Value) value.Value {
	l := n.Expr.Eval(env)
	u := n.Unit.Eval(env)
	if len(l.Unit) == 0 {
		panic(l.String() + " is a unitless value")
	}
	if len(u.Unit) == 0 {
		panic(n.Unit.String() + " is not a unit")
	}
	if u.Num.Cmp(ratutils.ONE) != 0 {
		panic("cannot convert to a unit with a value: " + n.Unit.String())
	}
	return l.To(u.Unit)
}

func (n *Convert) String() string {
	return n.Expr.String() + " to " + n.Unit.String()
}
*/

type FunCall struct {
	Name string
	Args []Node
}

func (n *FunCall) Eval(env map[string]value.Value) value.Value {
	fn := funcs.Get(n.Name)
	if fn == nil {
		panic("unknown function: " + n.Name)
	}
	args := make([]value.Value, len(n.Args))
	for i, arg := range n.Args {
		args[i] = arg.Eval(env)
	}
	val, err := (*fn)(args...)
	if err != nil {
		panic(n.Name + ": " + err.Error())
	}
	return val
}

func (n *FunCall) String() string {
	args := make([]string, len(n.Args))
	for i, arg := range n.Args {
		args[i] = arg.String()
	}
	return n.Name + "(" + strings.Join(args, ", ") + ")"
}

type Literal struct {
	Val value.Value
}

func (n *Literal) Eval(env map[string]value.Value) value.Value {
	return n.Val
}

func (n *Literal) String() string {
	return n.Val.String()
}
