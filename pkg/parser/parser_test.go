package parser

import (
	"reflect"
	"testing"

	"github.com/nkanaev/numb/pkg/ast"
	"github.com/nkanaev/numb/pkg/token"
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
		Op: token.MUL,
	}
	if !reflect.DeepEqual(want, have) {
		t.Errorf("\nexpr: %s\nwant: %s\nhave: %s", expr, want, have)
	}
}
