package scanner

import (
	"reflect"
	"testing"

	"github.com/nkanaev/numb/pkg/token"
)

func TestParseNumber(t *testing.T) {
	p := New(" 123")
	p.Scan()
	if p.Token != token.NUM || p.Value != "123" {
		t.Errorf("\nwant: %s (%#v)\nhave: %s (%#v)", token.NUM, "123", p.Token, p.Value)
	}
}

func TestParseToken(t *testing.T) {
	want := []token.Token{
		token.ADD, token.SUB, token.MUL, token.QUO,
		token.LPAREN, token.RPAREN,
		token.SHL, token.SHR,
		token.AND, token.OR, token.XOR, token.REM, token.VAR,
	}
	have := make([]token.Token, 0)
	text := `
		+ - * /
		( )
		<< >>
		and or xor mod foobar
	`
	s := New(text)
	for s.Scan() {
		have = append(have, s.Token)
	}

	if !reflect.DeepEqual(want, have) {
		t.Fatalf("\nwant: %s\nhave: %s", want, have)
	}

	// final call
	if s.Scan() {
		t.Fatal("expected to return false")
	}
}
