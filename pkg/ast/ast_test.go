package ast

import (
	"testing"

	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/value"
)

func TestASTLiteral(t *testing.T) {
	root := &Literal{value.Int64(123)}
	have := root.Eval(nil).String()
	want := "123"
	if have != want {
		t.Fatalf("\nwant: %s\nhave: %s", want, have)
	}
}

func TestASTBinOPEval(t *testing.T) {
	testcases := []struct {
		str string
		tok token.Token
	}{
		{"8", token.ADD},
		{"4", token.SUB},
		{"12", token.MUL},
		{"3", token.QUO},
		{"24", token.SHL},
		{"1", token.SHR},
		{"2", token.AND},
		{"6", token.OR},
		{"4", token.XOR},
		{"0", token.REM},
		{"36", token.EXP},
	}
	for _, testcase := range testcases {
		root := &BinOP{
			Lhs: &Literal{value.Int64(6)},
			Rhs: &Literal{value.Int64(2)},
			Op: testcase.tok,
		}
		have := root.Eval(nil).String()
		want := testcase.str
		if have != want {
			t.Errorf("\nexpr: %s\nwant: %s\nhave: %s", root, want, have)
		}
	}
}
