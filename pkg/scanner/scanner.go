package scanner

import (
	"strings"
	"unicode"

	"github.com/nkanaev/numb/pkg/token"
)

type Scanner struct {
	src []rune
	cur int

	Token token.Token
	Value string
}

func New(line string) *Scanner {
	return &Scanner{
		src:   []rune(line),
		Token: token.Illegal,
	}
}

func (s *Scanner) char() rune {
	if s.cur >= len(s.src) {
		return 0
	}
	return s.src[s.cur]
}

func (s *Scanner) nextChar() {
	s.cur += 1
}

func (s *Scanner) next() {
	for ; unicode.IsSpace(s.char()); s.nextChar() {
	}
	ch := s.char()
	switch {
	case unicode.IsDigit(ch):
		accept := "0123456789"
		prefix := ""
		digits := make([]rune, 0)
		if ch == '0' {
			s.nextChar()
			ch = s.char()
			if ch == 'b' {
				prefix = "0b"
				accept = "01"
				s.nextChar()
			} else if ch == 'x' {
				prefix = "0x"
				accept = "0123456789abcdefABCDEF"
				s.nextChar()
			} else if ch == 'o' {
				prefix = "0o"
				accept = "01234567"
				s.nextChar()
			}
		}
		for ; strings.ContainsRune(accept, s.char()); s.nextChar() {
			digits = append(digits, s.char())
		}
		if prefix == "" && s.char() == '.' {
			digits = append(digits, s.char())
			s.nextChar()
			for ; strings.ContainsRune(accept, s.char()); s.nextChar() {
				digits = append(digits, s.char())
			}
		}
		s.Token = token.NUM
		s.Value = prefix + string(digits)
	case ch == '=':
		s.Token = token.ASSIGN
		s.nextChar()
	case ch == '(':
		s.Token = token.LPAREN
		s.nextChar()
	case ch == ')':
		s.Token = token.RPAREN
		s.nextChar()
	case ch == ',':
		s.Token = token.COMMA
		s.nextChar()
	case ch == '*' || ch == '/' || ch == '+' || ch == '-':
		s.Token = token.StringToOperator[string(ch)]
		s.nextChar()
	case ch == '<' || ch == '>':
		s.nextChar()
		if s.char() != ch {
			s.cur -= 1
			s.Token = token.Illegal
			return
		}
		s.nextChar()
		switch ch {
		case '<':
			s.Token = token.SHL
		case '>':
			s.Token = token.SHR
		}
	case unicode.IsLetter(ch):
		letters := make([]rune, 0)
		for unicode.IsLetter(ch) || unicode.IsNumber(ch) {
			letters = append(letters, ch)
			s.nextChar()
			ch = s.char()
		}
		word := string(letters)
		if tok, ok := token.StringToOperator[word]; ok {
			s.Token = tok
			s.Value = word
		} else if word == "as" {
			s.Token = token.AS
		} else if word == "to" {
			s.Token = token.TO
		} else {
			s.Token = token.VAR
			s.Value = word
		}
	default:
		s.Token = token.Illegal
	}
}

func (s *Scanner) Scan() bool {
	s.Value = ""
	s.next()
	return s.Token != token.Illegal
}
