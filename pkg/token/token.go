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
	COLON  // :

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
	COLON: ":",
	ASSIGN: "=",

	NUM:  "NUM",
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
	case ASSIGN:
		return 1
	case AS:
		return 2
	case TO:
		return 3
	case ADD, SUB, OR, XOR:
		return 4
	case MUL, QUO, REM, SHL, SHR, AND:
		return 5
	case EXP:
		return 6
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
