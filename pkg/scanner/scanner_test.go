package scanner

import (
	"reflect"
	"testing"

	"github.com/nkanaev/numb/pkg/token"
)

func TestParseNumber(t *testing.T) {
	cases := []struct{
		tok token.Token
		want, have string	
	}{
		{token.NUM_DEC, "123", "123"},
		{token.NUM_BIN, "0b101", "0b101"},
		{token.NUM_HEX, "0x123", "0x123"},
		{token.NUM_DEC, "123", "123"},
		{token.NUM_OCT, "0o123", "0o123"},
		{token.NUM_DEC, "123.456", "123.456"},
		{token.NUM_DEC, "100200300", "100200300"},
		{token.NUM_DEC, "100200300", "100,200,300"},
		{token.NUM_DEC, "100.200300", "100.200,300"},
		{token.NUM_DEC, "0", "0"},
		{token.NUM_DEC, "123", "0123"},
		{token.NUM_DEC, "1.23", "01.23"},
		{token.NUM_DEC, "0.123", "0.123"},
		{token.NUM_SCI, "1.2e3", "1.2e3"},
		{token.NUM_SCI, "1.2e3", "1.2e+3"},
		{token.NUM_SCI, "1.2e-3", "1.2e-3"},
	}

	text := ""
	for _, c := range cases {
		text += " " + c.have
	}

	s := New(text)
	for i := 0; i < len(cases); i++ {
		if !s.Scan() {
			t.Fatal("finished too early")
		}
		if s.Token != cases[i].tok {
			t.Log(s.Value)
			t.Fatalf("expected %s, got %s", cases[i].tok, s.Token)
		}
		if s.Value != cases[i].want {
			t.Fatalf("\nwant: %s\nhave: %s", cases[i].want, s.Value)
		}
	}

	if s.Scan(); s.Token != token.END {
		t.Fatalf("did not finish properly, last token: %s (%s)", s.Token, s.Value)
	}
}

func TestParseToken(t *testing.T) {
	want := []token.Token{
		token.ADD, token.SUB, token.MUL, token.QUO,
		token.LPAREN, token.RPAREN,
		token.SHL, token.SHR,
		token.AND, token.OR, token.XOR, token.REM, token.EXP,
		token.COLON, token.ASSIGN, token.COMMA,
	}
	have := make([]token.Token, 0)
	text := `
		+ - * /
		( )
		<< >>
		and or xor mod ^
		: = ,
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
	want := []string{"foobar", "foo_bar"}
	have := make([]string, 0)
	text := ` foobar foo_bar `
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
	want := []token.Token{token.TO, token.IN, token.NAME}
	have := make([]token.Token, 0)
	text := " to in hex"

	s := New(text)
	for s.Scan() {
		have = append(have, s.Token)
	}

	if !reflect.DeepEqual(want, have) {
		t.Fatalf("\nwant: %s\nhave: %s", want, have)
	}
}

// TODO: syntax error markers
