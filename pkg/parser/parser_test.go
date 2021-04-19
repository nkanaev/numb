package parser

import (
	"reflect"
	"testing"

	"github.com/nkanaev/numb/pkg/ast"
	"github.com/nkanaev/numb/pkg/scanner"
)

func TestParser(t *testing.T) {
	have := Parse("1 + 2")
	want := &ast.BinOP{
		Lhs: &ast.Const{Val: 1},
		Rhs: &ast.Const{Val: 2},
		Op: scanner.ADD,
	}

	if !reflect.DeepEqual(want, have) {
		t.Fatalf("\nwant: %s\nhave: %s", want, have)
	}
}
