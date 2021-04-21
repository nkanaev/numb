package scanner

import (
	"reflect"
	"testing"
)

func TestParseNumber(t *testing.T) {
	p := New(" 123")
	p.Scan()
	if p.Token != NUM || p.Value != "123" {
		t.Errorf("\nwant: %s (%#v)\nhave: %s (%#v)", NUM, "123", p.Token, p.Value)
	}
}

func TestParseToken(t *testing.T) {
	want := []Token{ADD, SUB, MUL, QUO}
	have := make([]Token, 0)
	text := "+ - * /"

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
