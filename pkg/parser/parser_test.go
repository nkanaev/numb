package parser

import (
	"reflect"
	"testing"

	"github.com/nkanaev/numb/pkg/ast"
	"github.com/nkanaev/numb/pkg/ratutils"
	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/value"
)

var tests = map[string]struct {
	name   string
	expr string
	want ast.Node
}{
	"BinOp": {
		expr: "1 + 2",
		want: &ast.BinOP{
			Lhs: &ast.Literal{value.Int64(1)},
			Rhs: &ast.Literal{value.Int64(2)},
			Op:  token.ADD,
		},
	},
	"Paren": {
		expr: "(1 + 2) * 3",
		want: &ast.BinOP{
			Lhs: &ast.ParenExpr{
				Expr: &ast.BinOP{
					Lhs: &ast.Literal{value.Int64(1)},
					Rhs: &ast.Literal{value.Int64(2)},
					Op:  token.ADD,
				},
			},
			Rhs: &ast.Literal{value.Int64(3)},
			Op:  token.MUL,
		},
	},
	"Unary": {
		expr: "-100",
		want: &ast.Unary{Op: token.SUB, Expr: &ast.Literal{value.Int64(100)}},
	},
	"BitOps": {
		expr: "0b101 and 0b111",
		want: &ast.BinOP{
			Lhs: &ast.Literal{value.Number{Num: ratutils.Must("5"), Fmt: value.BIN}},
			Rhs: &ast.Literal{value.Number{Num: ratutils.Must("7"), Fmt: value.BIN}},
			Op:  token.AND,
		},
	},
	"Assignment": {
		expr: "foo = 123",
		want: &ast.Assign{
			Name: "foo",
			Expr: &ast.Literal{value.Int64(123)},
		},
	},
	"Var": {
		expr: "foo + 123",
		want: &ast.BinOP{
			Lhs: &ast.Var{Name: "foo"},
			Rhs: &ast.Literal{value.Int64(123)},
			Op:  token.ADD,
		},
	},
	"Format": {
		expr: "10 + 1 in hex",
		want: &ast.Format{
			Expr: &ast.BinOP{
				Lhs: &ast.Literal{value.Int64(10)},
				Rhs: &ast.Literal{value.Int64(1)},
				Op:  token.ADD,
			},
			Fmt: "hex",
		},
	},
	"FunCall": {
		expr: "sin(2 radian)",
		want: &ast.FunCall{
			Name: "sin",
			Args: []ast.Node{
				&ast.BinOP{
					Lhs:      &ast.Literal{value.Int64(2)},
					Rhs:      &ast.Var{Name: "radian"},
					Op:       token.MUL,
					Implicit: true,
				},
			},
		},
	},
}

func TestParser(t *testing.T) {
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			expr := tc.expr
			want := tc.want
			have, err := Parse(expr)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(want, have) {
				t.Errorf("\nexpr: %s\nwant: %#v\nhave: %#v", expr, want, have)
			}
		})
	}
}
