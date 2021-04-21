package ast

import (
	"testing"

	"github.com/nkanaev/numb/pkg/scanner"
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
	root := &BinOP{Lhs: value.NewInt(2), Rhs: value.NewInt(3), Op: scanner.ADD}	
	have := root.Eval(nil).String()
	want := "5"
	if have != want {
		t.Fatalf("\nwant: %s\nhave: %s", want, have)
	}
}
