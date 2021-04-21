package token

type Token int

const (
	Illegal Token = iota

	OR  // or
	XOR // xor
	AND // and
	SHL // <<
	SHR // >>

	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // mod
	EXP // pow

	LPAREN // (
	RPAREN // )

	NUM
	VAR
)

func (t Token) String() string {
	if t == Illegal {
		return "ILLEGAL"
	}
	if t == NUM {
		return "NUM"
	}

	for str, tok := range TokenString {
		if tok == t {
			return str
		}
	}	

	if t == LPAREN {
		return "("
	}
	if t == RPAREN {
		return ")"
	}
	return "???"
}

const (
	LowestPrec = 0
)

func (t Token) Precedence() int {
	switch t {
	case ADD, SUB, OR, XOR:
		return 1
	case MUL, QUO, REM, SHL, SHR, AND:
		return 2
	}
	return LowestPrec
}

var TokenString = map[string]Token{
	"or": OR,
	"xor": XOR,
	"and": AND,
	"<<": SHL,
	">>": SHR,

	"+": ADD,
	"-": SUB,
	"*": MUL,
	"/": QUO,
	"mod": REM,

	"pow": EXP,
}
