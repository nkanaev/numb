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

	ASSIGN // =

	NUM
	VAR

	KEYWORD
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
	if t == VAR {
		return "VAR"
	}
	if t == ASSIGN {
		return "ASSIGN"
	}
	if t == KEYWORD {
		return "KEYWORD"
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
	case MUL, QUO, REM, SHL, SHR, AND, EXP:
		return 2
	}
	return LowestPrec
}

var TokenString = map[string]Token{
	"or":  OR,
	"xor": XOR,
	"and": AND,
	"<<":  SHL,
	">>":  SHR,

	"+":   ADD,
	"-":   SUB,
	"*":   MUL,
	"/":   QUO,
	"mod": REM,

	"pow": EXP,
}

var Keywords = map[string]int {
	"as": 1,
	"to": 1,
}

func IsKeyword(x string) bool {
	_, ok := Keywords[x]
	return ok
}
