package ast

import (
	"testing"

	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/value"
)

func TestASTConst(t *testing.T) {
	root := value.NewInt(123)
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
	}
	for _, testcase := range testcases {
		root := &BinOP{Lhs: value.NewInt(6), Rhs: value.NewInt(2), Op: testcase.tok}
		have := root.Eval(nil).String()
		want := testcase.str
		if have != want {
			t.Errorf("\nexpr: %s\nwant: %s\nhave: %s", root, want, have)
		}
	}
}
