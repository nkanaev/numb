package parser

import (
	"reflect"
	"testing"

	"github.com/nkanaev/numb/pkg/ast"
	"github.com/nkanaev/numb/pkg/scanner"
	"github.com/nkanaev/numb/pkg/value"
)

func TestParserBinOP(t *testing.T) {
	ops := []struct {
		str string
		val scanner.Token
	}{
		{"+", scanner.ADD},
		{"-", scanner.SUB},
		{"*", scanner.MUL},
		{"/", scanner.QUO},
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
