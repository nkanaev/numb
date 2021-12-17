package scanner

import (
	"strings"
	"unicode"

	"github.com/nkanaev/numb/pkg/token"
)

type Scanner struct {
	src []rune
	cur int
	ch  rune

	namemode bool

	Token token.Token
	Value string
}

func New(line string) *Scanner {
	s := &Scanner{
		src:   []rune(line),
		cur:   -1,
		Token: token.Illegal,
	}
	s.next()
	return s
}

func (s *Scanner) Pos() int {
	return s.cur
}

func (s *Scanner) next() {
	s.cur += 1
	if s.cur < len(s.src) {
		s.ch = s.src[s.cur]
	} else {
		s.ch = 0
	}
}

func isDecimal(ch rune) bool { return '0' <= ch && ch <= '9' }

func (s *Scanner) skipWhitespace() {
	for ; unicode.IsSpace(s.ch); s.next() {
	}
}

func (s *Scanner) scan() {
	if s.cur >= len(s.src) {
		s.Token = token.END
		return
	}

	ch := s.ch
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
		if s.ch != ch {
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
	default:
		letters := make([]rune, 0)
		for isWordChar(ch) {
			letters = append(letters, ch)
			s.next()
			ch = s.ch
		}
		if len(letters) == 0 {
			s.Token = token.Illegal
		} else {
			word := string(letters)
			if tok, ok := token.StringToOperator[word]; ok {
				s.Token = tok
				s.Value = word
				if tok == token.IN {
					s.namemode = true
				}
			} else {
				s.Token = token.WORD
				s.Value = word
			}
		}
	}
}

func isWordChar(ch rune) bool {
	// L/N/Sc/Pc = letters/numbers/currency symbols/connector punctuation
	return !strings.Contains(token.SpecialChars, string(ch)) && unicode.In(ch, unicode.L, unicode.N, unicode.Sc, unicode.Pc)
}

func (s *Scanner) digits(base int) string {
	separators := ",_"
	digits := make([]rune, 0)
	accept := "0123456789abcdef"[:base]
	if base == 16 {
		accept += "ABCDEF"
	}
	loop: for  {
		ch := s.ch
		switch {
		case strings.ContainsRune(accept, ch):
			digits = append(digits, s.ch)
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
	tok := token.NUM_DEC
	val := ""
	base := 10

	// integer
	if s.ch == '0' {
		s.next()
		switch s.ch {
		case 'x', 'X':
			s.next()
			base = 16
			val = "0x"
			tok = token.NUM_HEX
		case 'o', 'O':
			s.next()
			base = 8
			val = "0o"
			tok = token.NUM_OCT
		case 'b', 'B':
			s.next()
			base = 2
			val = "0b"
			tok = token.NUM_BIN
		default:
			val = "0"
		}
	}
	val += s.digits(base)

	// HACK: remove leading 0 from decimal numbers to avoid confusion with octal
	// 0123 = 123 (!= 0o123)
	if len(val) > 1 && val[0] == '0' && base == 10 {
		val = val[1:]
	}

	// fraction
	if base == 10 && s.ch == '.' {
		s.next()
		val += "." + s.digits(base)
	}

	// exponent
	if base == 10 && s.ch == 'e' {
		tok = token.NUM_SCI
		s.next()
		switch s.ch {
		case '-':
			s.next()
			val += "e-" + s.digits(base)
		case '+':
			s.next()
			fallthrough
		default:
			val += "e" + s.digits(base)
		}
	}
	return tok, val
}

func (s *Scanner) Scan() bool {
	s.skipWhitespace()
	s.Value = ""

	if s.namemode {
		s.scanname()
	} else {
		s.scan()
	}
	return s.Token != token.Illegal && s.Token != token.END
}

func (s *Scanner) scanname() {
	if s.cur >= len(s.src) {
		s.Token = token.END
		s.Value = ""
		return
	}

	chars := make([]rune, 0)
	for ; s.ch != 0 && !unicode.IsSpace(s.ch); s.next() {
		chars = append(chars, s.ch)
	}
	s.Token = token.NAME
	s.Value = string(chars)
}
