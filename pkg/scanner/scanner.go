package scanner

import "unicode"

type parser struct {
	src []rune
	cur int

	tok token
	val string
}

func New(line string) *parser {
	return &parser{
		src: []rune(line),
		tok: Illegal,
	}
}

func (p *parser) char() rune {
	if p.cur >= len(p.src) {
		return rune(0)
	}
	return p.src[p.cur]
}

func (p *parser) nextChar() {
	p.cur += 1
}

func (p *parser) next() {
	for ; unicode.IsSpace(p.char()); p.nextChar() {
	}
	ch := p.char()
	switch {
	case unicode.IsDigit(ch):
		start := p.cur
		for ; unicode.IsDigit(p.char()); p.nextChar() {
		}
		p.tok = NUM
		p.val = string(p.src[start:p.cur])
	case isOp(string([]rune{ch})):
		val := string([]rune{ch})
		p.tok = tokenString[val]
		p.val = val

		p.nextChar()
		val2 := string([]rune{ch, p.char()})
		if isOp(val2) {
			p.tok = tokenString[val2]
			p.val = val2
			p.nextChar()
		}
	default:
		p.tok = Illegal
	}
}
