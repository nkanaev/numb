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
	Eval(map[string]value.Value) (value.Value, error)
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

func (n *BinOP) Eval(env map[string]value.Value) (value.Value, error) {
	lhs, err := n.Lhs.Eval(env)
	if err != nil {
		return nil, err
	}
	rhs, err := n.Rhs.Eval(env)
	if err != nil {
		return nil, err
	}

	val, err := lhs.BinOP(n.Op, rhs)
	if err != nil {
		return nil, err
	}
	return val, nil
}

type ParenExpr struct {
	Expr Node
}

func (n *ParenExpr) Eval(env map[string]value.Value) (value.Value, error) {
	return n.Expr.Eval(env)
}

func (n *ParenExpr) String() string {
	return "(" + n.Expr.String() + ")"
}

type Unary struct {
	Op   token.Token
	Expr Node
}

func (n *Unary) Eval(env map[string]value.Value) (value.Value, error) {
	val, err := n.Expr.Eval(env)
	if err != nil {
		return nil, err
	}
	val, err = val.UnOP(n.Op)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (n *Unary) String() string {
	return n.Op.String() + n.Expr.String()
}

type Assign struct {
	Name string
	Expr Node
	Unit bool
}

func (n *Assign) Eval(env map[string]value.Value) (value.Value, error) {
	val, err := n.Expr.Eval(env)
	if err != nil {
		return nil, err
	}
	if n.Unit {
		if val, isnum := val.(value.Number); isnum {
			unit.Add(n.Name, val.Num, unit.Units{})
			return val, nil
		}
		if val, isunit := val.(value.Unit); isunit {
			unit.Add(n.Name, val.Num, val.Units)
			return val, nil
		}
		return nil, fmt.Errorf("cannot create unit from %s (%s type)", val, value.Type(val))
	} else {
		env[n.Name] = val
	}
	return val, nil
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

func (n *Var) Eval(env map[string]value.Value) (value.Value, error) {
	if val, ok := env[n.Name]; ok {
		return val, nil
	}
	if time := value.GetNamedTime(n.Name); time != nil {
		return time, nil
	}
	if unit, ok := unit.Get(n.Name); ok {
		return value.Unit{Num: ratutils.ONE, Units: unit}, nil
	}
	return nil, fmt.Errorf("%s not defined", n.Name)
}

func (n *Var) String() string {
	return n.Name
}

type Format struct {
	Expr Node
	Fmt  string
}

func (n *Format) Eval(env map[string]value.Value) (value.Value, error) {
	val, err := n.Expr.Eval(env)
	if err != nil {
		return nil, err
	}
	val, err = val.In(n.Fmt)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (n *Format) String() string {
	return fmt.Sprintf("%s %s %s", n.Expr.String(), token.IN, n.Fmt)
}

type FunCall struct {
	Name string
	Args []Node
}

func (n *FunCall) Eval(env map[string]value.Value) (value.Value, error) {
	fn := funcs.Get(n.Name)
	if fn == nil {
		return nil, fmt.Errorf("unknown function: %s", n.Name)
	}
	args := make([]value.Value, len(n.Args))
	var err error
	for i, arg := range n.Args {
		args[i], err = arg.Eval(env)
		if err != nil {
			return nil, err
		}
	}
	val, err := (*fn)(args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", n.Name, err)
	}
	return val, nil
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

func (n *Literal) Eval(env map[string]value.Value) (value.Value, error) {
	return n.Val, nil
}

func (n *Literal) String() string {
	return n.Val.String()
}
