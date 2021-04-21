package scanner

type Token int

const (
	Illegal Token = iota

	OR  // |
	XOR // ^
	AND // &
	SHL // <<
	SHR // >>

	ADD // +
	SUB // -
	MUL // *
	QUO // /

	LPAREN // (
	RPAREN // )

	NUM
)

func (t Token) String() string {
	if t == Illegal {
		return "ILLEGAL"
	}
	if t == NUM {
		return "NUM"
	}

	for str, tok := range tokenString {
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
	case MUL, QUO, SHL, SHR, AND:
		return 5
	}
	return LowestPrec
}

var tokenString = map[string]Token{
	"|": OR,
	"^": XOR,

	"&":  AND,
	"<<": SHL,
	">>": SHR,

	"+": ADD,
	"-": SUB,
	"*": MUL,
	"/": QUO,
}

func isOp(t string) bool {
	_, ok := tokenString[t]
	return ok
}
