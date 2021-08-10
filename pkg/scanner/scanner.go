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

func (s *Scanner) Pos() int {
	return s.cur
}

func (s *Scanner) char() rune {
	if s.cur >= len(s.src) {
		return 0
	}
	return s.src[s.cur]
}

func (s *Scanner) next() {
	s.cur += 1
}

func isDecimal(ch rune) bool { return '0' <= ch && ch <= '9' }

func (s *Scanner) scan() {
	for ; unicode.IsSpace(s.char()); s.next() {
	}

	if s.cur >= len(s.src) {
		s.Token = token.END
		return
	}

	ch := s.char()
	switch {
	case isDecimal(ch):
		s.Token, s.Value = s.scanNumber()
	case ch == '=':
		s.Token = token.ASSIGN
		s.next()
	case ch == ':':
		s.Token = token.COLON
		s.next()
	case ch == '(':
		s.Token = token.LPAREN
		s.next()
	case ch == ')':
		s.Token = token.RPAREN
		s.next()
	case ch == ',':
		s.Token = token.COMMA
		s.next()
	case ch == '^':
		s.Token = token.EXP
		s.next()
	case ch == '*' || ch == '/' || ch == '+' || ch == '-':
		s.Token = token.StringToOperator[string(ch)]
		s.next()
	case ch == '<' || ch == '>':
		s.next()
		if s.char() != ch {
			s.cur -= 1
			s.Token = token.Illegal
			return
		}
		s.next()
		switch ch {
		case '<':
			s.Token = token.SHL
		case '>':
			s.Token = token.SHR
		}
	case ch == '{':
		letters := make([]rune, 0)
		for {
			if s.char() == 0 {
				panic("unexpected end of word")
			}
			letters = append(letters, s.char())
			if s.char() == '}' {
				s.next()
				break
			}
			s.next()
		}
		s.Token = token.WORD
		s.Value = string(letters)
	default:
		letters := make([]rune, 0)
		for ch != 0 && !unicode.IsSpace(ch) && !strings.Contains("^*/+-()=:", string(ch)) {
			letters = append(letters, ch)
			s.next()
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
		} else if unicode.IsLetter(letters[0]) {
			s.Token = token.WORD
			s.Value = word
		} else {
			s.Token = token.Illegal
		}
	}
}

func (s *Scanner) digits(base int) string {
	separators := ",_"
	digits := make([]rune, 0)
	accept := "0123456789abcdef"[:base]
	if base == 16 {
		accept += "ABCDEF"
	}
	loop: for  {
		ch := s.char()
		switch {
		case strings.ContainsRune(accept, ch):
			digits = append(digits, s.char())
			s.next()
		case strings.ContainsRune(separators, ch):
			s.next()
		default:
			break loop
		}
	}
	return string(digits)
}

func (s *Scanner) scanNumber() (token.Token, string) {
	value := ""
	base := 10

	// integer
	if s.char() == '0' {
		s.next()
		switch s.char() {
		case 'x', 'X':
			s.next()
			base = 16
			value = "0x"
		case 'o', 'O':
			s.next()
			base = 8
			value = "0o"
		case 'b', 'B':
			s.next()
			base = 2
			value = "0b"
		default:
			value = "0"
		}
	}
	value += s.digits(base)

	// HACK: remove leading 0 from decimal numbers to avoid confusion with octal
	// 0123 = 123 (!= 0o123)
	if len(value) > 1 && value[0] == '0' && base == 10 {
		value = value[1:]
	}

	// fraction
	if base == 10 && s.char() == '.' {
		s.next()
		value += "." + s.digits(base)
	}

	// exponent
	if base == 10 && s.char() == 'e' {
		s.next()
		switch s.char() {
		case '-':
			s.next()
			value += "e-" + s.digits(base)
		case '+':
			s.next()
			fallthrough
		default:
			value += "e" + s.digits(base)
		}
	}
	return token.NUM, value
}

func (s *Scanner) Scan() bool {
	s.Value = ""
	s.scan()
	return s.Token != token.Illegal && s.Token != token.END
}
