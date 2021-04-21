package parser

import (
	"reflect"
	"testing"

	"github.com/nkanaev/numb/pkg/ast"
	"github.com/nkanaev/numb/pkg/scanner"
	"github.com/nkanaev/numb/pkg/value"
)

func TestParser(t *testing.T) {
	have := Parse("1 + 2")
	want := &ast.BinOP{
		Lhs: value.NewInt(1),
		Rhs: value.NewInt(2),
		Op: scanner.ADD,
	}

	if !reflect.DeepEqual(want, have) {
		t.Fatalf("\nwant: %s\nhave: %s", want, have)
	}
}
