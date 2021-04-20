package ast

import (
	"testing"

	"github.com/nkanaev/numb/pkg/scanner"
)

func TestASTConst(t *testing.T) {
	root := &Const{Val: 123}
	have := root.Eval(nil)
	want := float64(123)
	if have != want {
		t.Fatalf("\nwant: %f\nhave: %f", want, have)
	}
}

func TestASTBinOPEval(t *testing.T) {
	root := &BinOP{Lhs: &Const{Val: 2}, Rhs: &Const{Val: 3}, Op: scanner.ADD}	
	have := root.Eval(nil)
	want := float64(5)
	if have != want {
		t.Fatalf("\nwant: %f\nhave: %f", want, have)
	}
}
