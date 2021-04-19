package scanner

type Token int

const (
	Illegal Token = iota

	LAND // &&
	LOR  // ||

	OR  // |
	XOR // ^
	AND // &
	SHL // <<
	SHR // >>

	ADD // +
	SUB // -
	MUL // *
	DIV // /
	MOD // %

	EQL // ==
	LSS // <
	GTR // >
	NEQ // !=
	LEQ // <=
	GEQ // >=

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
	return "???"
}

const (
	LowestPrec = 0
)

func (t Token) Precedence() int {
	switch t {
	case LOR:
		return 1
	case LAND:
		return 2
	case EQL, NEQ, LSS, LEQ, GTR, GEQ:
		return 3
	case ADD, SUB, OR, XOR:
		return 4
	case MUL, DIV, MOD, SHL, SHR, AND:
		return 5
	}
	return LowestPrec
}

var tokenString = map[string]Token{
	"&&": LAND,
	"||": LOR,

	"|": OR,
	"^": XOR,

	"&":  AND,
	"<<": SHL,
	">>": SHR,

	"+": ADD,
	"-": SUB,
	"*": MUL,
	"/": DIV,
	"%": MOD,

	"==": EQL,
	"<":  LSS,
	">":  GTR,
	"!=": NEQ,
	"<=": LEQ,
	">=": GEQ,
}

func isOp(t string) bool {
	_, ok := tokenString[t]
	return ok
}
