package parser

import (
	"reflect"
	"testing"

	"github.com/nkanaev/numb/pkg/ast"
	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/unit"
	"github.com/nkanaev/numb/pkg/value"
)

func TestParserBinOP(t *testing.T) {
	ops := []struct {
		str string
		val token.Token
	}{
		{"+", token.ADD},
		{"-", token.SUB},
		{"*", token.MUL},
		{"/", token.QUO},
	}
	for _, op := range ops {
		expr := "1 " + op.str + " 2"
		have := Parse(expr)
		want := &ast.BinOP{
			Lhs: value.NewInt(1),
			Rhs: value.NewInt(2),
			Op:  op.val,
		}

		if !reflect.DeepEqual(want, have) {
			t.Errorf("\nexpr: %s\nwant: %s\nhave: %s", expr, want, have)
		}
	}
}

func TestParserParen(t *testing.T) {
	expr := "(1 + 2) * 3"
	have := Parse(expr)
	want := &ast.BinOP{
		Lhs: &ast.ParenExpr{Expr: &ast.BinOP{Lhs: value.NewInt(1), Rhs: value.NewInt(2), Op: token.ADD}},
		Rhs: value.NewInt(3),
		Op:  token.MUL,
	}
	if !reflect.DeepEqual(want, have) {
		t.Errorf("\nexpr: %s\nwant: %s\nhave: %s", expr, want, have)
	}
}

func TestParserUnary(t *testing.T) {
	expr := "-100"
	have := Parse(expr)
	want := &ast.Unary{Op: token.SUB, Expr: value.NewInt(100)}
	if !reflect.DeepEqual(want, have) {
		t.Errorf("\nexpr: %s\nwant: %s\nhave: %s", expr, want, have)
	}
}

func TestParseBitOps(t *testing.T) {
	expr := "0b101 and 0b111"
	have := Parse(expr)
	want := &ast.BinOP{
		Lhs: value.NewInt(5).As(value.BIN),
		Rhs: value.NewInt(7).As(value.BIN),
		Op:  token.AND,
	}
	if !reflect.DeepEqual(want, have) {
		t.Errorf("\nexpr: %s\nwant: %s\nhave: %s", expr, want, have)
	}
}

func TestParseAssign(t *testing.T) {
	expr := "foo = 123"
	have := Parse(expr)
	want := &ast.Assign{
		Name: "foo",
		Expr: value.NewInt(123),
	}
	if !reflect.DeepEqual(want, have) {
		t.Errorf("\nexpr: %s\nwant: %s\nhave: %s", expr, want, have)
	}
}

func TestParseVar(t *testing.T) {
	expr := "foo + 123"
	have := Parse(expr)
	want := &ast.BinOP{
		Lhs: &ast.Var{Name: "foo"},
		Rhs: value.NewInt(123),
		Op:  token.ADD,
	}
	if !reflect.DeepEqual(want, have) {
		t.Errorf("\nexpr: %s\nwant: %s\nhave: %s", expr, want, have)
	}
}

func TestParseFormat(t *testing.T) {
	expr := "10 + 1 as hex"
	have := Parse(expr)
	want := &ast.Format{
		Expr: &ast.BinOP{
			Lhs: value.NewInt(10),
			Rhs: value.NewInt(1),
			Op:  token.ADD,
		},
		Fmt: value.HEX,
	}
	if !reflect.DeepEqual(want, have) {
		t.Errorf("\nexpr: %s\nwant: %#v\nhave: %#v", expr, want, have)
	}
}

func TestParseFunCall(t *testing.T) {
	expr := "sin(2 radian)"
	have := Parse(expr)
	want := &ast.FunCall{
		Name: "sin",
		Args: []ast.Node{&ast.Unit{Expr: value.NewInt(2), Unit: unit.Get("rad")}},
	}
	if !reflect.DeepEqual(want, have) {
		t.Errorf("\nexpr: %s\nwant: %#v\nhave: %#v", expr, want, have)
	}
}
