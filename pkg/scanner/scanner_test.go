package scanner

import (
	"reflect"
	"testing"

	"github.com/nkanaev/numb/pkg/token"
)

func TestParseNumber(t *testing.T) {
	text := " 123 0b101 0x123 0123 0o123"
	want := []string{"123", "0b101", "0x123", "123", "0o123"}
	have := make([]string, 0)

	s := New(text)
	for s.Scan() {
		if s.Token != token.NUM {
			t.Log(s.Value)
			t.Fatalf("expected %s, got %s", token.NUM, s.Token)
		}
		have = append(have, s.Value)
	}

	if !reflect.DeepEqual(want, have) {
		t.Fatalf("\nwant: %s\nhave: %s", want, have)
	}
}

func TestParseToken(t *testing.T) {
	want := []token.Token{
		token.ADD, token.SUB, token.MUL, token.QUO,
		token.LPAREN, token.RPAREN,
		token.SHL, token.SHR,
		token.AND, token.OR, token.XOR, token.REM, token.EXP,
		token.VAR,
	}
	have := make([]token.Token, 0)
	text := `
		+ - * /
		( )
		<< >>
		and or xor mod pow
		foobar
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

func TestParseKeywords(t *testing.T) {
	want := []token.Token{token.AS, token.TO}
	have := make([]token.Token, 0)
	text := " as to"
	
	s := New(text)
	for s.Scan() {
		have = append(have, s.Token)
	}

	if !reflect.DeepEqual(want, have) {
		t.Fatalf("\nwant: %#v\nhave: %#v", want, have)
	}
}
