package scanner

import "unicode"

type Scanner struct {
	src []rune
	cur int

	Token Token
	Value string
}

func New(line string) *Scanner {
	return &Scanner{
		src: []rune(line),
		Token: Illegal,
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
		start := s.cur
		for ; unicode.IsDigit(s.char()); s.nextChar() {
		}
		s.Token = NUM
		s.Value = string(s.src[start:s.cur])
	case ch == '(':
		s.Token = LPAREN
		s.nextChar()
	case ch == ')':
		s.Token = RPAREN
		s.nextChar()
	case isOp(string([]rune{ch})):
		val := string([]rune{ch})
		s.Token = tokenString[val]
		s.Value = val

		s.nextChar()
		val2 := string([]rune{ch, s.char()})
		if isOp(val2) {
			s.Token = tokenString[val2]
			s.Value = val2
			s.nextChar()
		}
	default:
		s.Token = Illegal
	}
}

func (s *Scanner) Scan() bool {
	s.next()
	return s.Token != Illegal
}
