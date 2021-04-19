package scanner

import "testing"

func TestParseNumber(t *testing.T) {
	p := New(" 123")
	p.next()
	if p.tok != NUM || p.val != "123" {
		t.Errorf("\nwant: %s (%#v)\nhave: %s (%#v)", NUM, "123", p.tok, p.val)
	}
}

func TestParseToken(t *testing.T) {
	want := []token{ADD, SUB, LEQ, GEQ, LSS, GTR}
	have := "+ - <= >= < >"

	p := New(have)
	for _, tok := range want {
		p.next()
		if p.tok != tok {
			t.Fatalf("\nwant: %s\nhave: %s", tok, p.tok)
		}
	}
}
