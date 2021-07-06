package token

type Token int

const (
	Illegal Token = iota
	END

	operator_beg
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
	EXP // ^
	operator_end

	LPAREN // (
	RPAREN // )
	COMMA  // ,

	ASSIGN // =

	NUM
	WORD

	AS
	TO
)

var tokenToString = map[Token]string{
	Illegal: "ILLEGAL",
	END:     "END",

	OR:  "OR",
	XOR: "XOR",
	AND: "AND",
	SHL: "<<",
	SHR: ">>",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "mod",
	EXP: "^",

	LPAREN: "(",
	RPAREN: ")",
	COMMA:  ",",

	ASSIGN: "=",

	NUM: "NUM",
	WORD: "WORD",

	AS: "as",
	TO: "to",
}

func (t Token) String() string {
	return tokenToString[t]
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
	case EXP:
		return 3
	}
	return LowestPrec
}

var StringToOperator = map[string]Token{
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

	"^": EXP,
}

func init() {
	for x := operator_beg + 1; x < operator_end; x++ {
		StringToOperator[x.String()] = x
	}
}
