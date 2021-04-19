package scanner

type token int

const (
	Illegal token = iota

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

func (t token) String() string {
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

func (t token) Precedence() int {
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
	return 0
}

var tokenString = map[string]token{
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
