package scanner

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/nkanaev/numb/pkg/token"
)

type idxerror struct {
	msg string
	idx int
}

func (e *idxerror) Error() string {
	return e.msg
}

func (e *idxerror) Pos() (int, int) {
	return e.idx, e.idx
}

type Scanner struct {
	src []rune
	cur int
	ch  rune

	formatMode bool

	Token                token.Token
	Value                string
	TokenStart, TokenEnd int

	Error error
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

func (s *Scanner) illegalToken(msg string, pos int) {
	s.Token = token.Illegal
	s.Error = &idxerror{msg, pos}
}

func (s *Scanner) scan() {
	s.TokenStart = s.cur
	s.TokenEnd = s.cur

	if s.cur >= len(s.src) {
		s.Token = token.END
		return
	}

	if s.formatMode {
		s.scanFormat()
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
			s.illegalToken(fmt.Sprintf("expected %c", ch), s.TokenStart+1)
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
		s.next()
		runes := make([]rune, 0)
		for s.ch != 0 && s.ch != '}' {
			runes = append(runes, s.ch)
			s.next()
		}
		if s.ch == 0 {
			s.illegalToken("dangling date (missing `}`)", s.TokenStart)
			return
		}
		s.next()
		s.Token = token.DATE
		s.Value = string(runes)
	default:
		s.scanIdent()
	}
	s.TokenEnd = s.cur - 1
}

func (s *Scanner) scanIdent() {
	if unicode.In(s.ch, unicode.Sc) || s.ch == '%' {
		// currency symbol or `%`
		s.Token = token.IDENT
		s.Value = string(s.ch)
		s.next()
		return
	} else if unicode.In(s.ch, unicode.Letter) {
		chars := make([]rune, 0)
		for unicode.In(s.ch, unicode.Letter, unicode.Number) || s.ch == '_' {
			chars = append(chars, s.ch)
			s.next()
		}
		value := string(chars)
		if op, ok := token.StringToOperator[value]; ok {
			s.Token = op
			s.Value = value
			if op == token.IN {
				s.formatMode = true
			}
		} else {
			s.Token = token.IDENT
			s.Value = value
		}
	} else {
		s.illegalToken("unexpected character", s.TokenStart)
	}
}

func (s *Scanner) digits(base int) string {
	// TODO: prohibit , as a separator. ambiguity: gcd(1,2) == gcd(12) || gcd(1, 2)?
	// TODO: space-delimited numbers `100 500`?
	separators := ",_"
	digits := make([]rune, 0)
	accept := "0123456789abcdef"[:base]
	if base == 16 {
		accept += "ABCDEF"
	}
loop:
	for {
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
	if s.Error != nil {
		return false
	}
	s.skipWhitespace()
	s.Value = ""
	s.scan()
	return s.Token != token.Illegal && s.Token != token.END
}

func (s *Scanner) scanFormat() {
	chars := make([]rune, 0)
	if unicode.IsLetter(s.ch) {
		// must cover tz database names & numeric formats
		for unicode.In(s.ch, unicode.Letter) || s.ch == '/' || s.ch == '_' {
			chars = append(chars, s.ch)
			s.next()
		}
		s.Token = token.FORMAT
		s.Value = string(chars)
		s.formatMode = false
	} else {
		s.illegalToken("unexpected symbol for format: "+string(s.ch), s.TokenStart)
	}
}
