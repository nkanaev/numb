package scanner

import (
	"reflect"
	"testing"

	"github.com/nkanaev/numb/pkg/token"
)

func TestParseNumber(t *testing.T) {
	text := " 123 0b101 0x123 0123 0o123 123.456, 100,200,300 100_200_300 0 0.123 1.2e3 1.2e+3 1.2e-3"
	want := []string{"123", "0b101", "0x123", "123", "0o123", "123.456", "100200300", "100200300", "0", "0.123", "1.2e3", "1.2e3", "1.2e-3"}
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
	}
	have := make([]token.Token, 0)
	text := `
		+ - * /
		( )
		<< >>
		and or xor mod ^
	`
	s := New(text)
	for i := 0; i < len(want); i++ {
		s.Scan()
		have = append(have, s.Token)
	}

	if !reflect.DeepEqual(want, have) {
		t.Fatalf("\nwant: %s\nhave: %s", want, have)
	}

	// final call
	if s.Scan() {
		t.Fatal("expected to return false")
	}
	if s.Token != token.END {
		t.Fatal("expected end token")
	}
}

func TestParseWord(t *testing.T) {
	want := []string{"foobar", "{foo bar}"}
	have := make([]string, 0)
	text := ` foobar {foo bar} `
	s := New(text)
	for i := 0; i < len(want); i++ {
		s.Scan()
		if s.Token != token.WORD {
			t.Fatalf("expected %s, got %s", token.WORD, s.Token)
		}
		have = append(have, s.Value)
	}
	if !reflect.DeepEqual(want, have) {
		t.Fatalf("\nwant: %s\nhave: %s", want, have)
	}
	// final call
	if s.Scan() {
		t.Fatal("expected to return false")
	}
	if s.Token != token.END {
		t.Fatalf("expected %v, got %v", token.END, s.Token)
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
