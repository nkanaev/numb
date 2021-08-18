package token

import "strings"

type Token int

const (
	Illegal Token = iota
	END

	operator_beg
	OR  // or
	XOR // xor
	AND // and
	REM // mod

	AS  // as
	TO  // to

	chars_beg
	SHL // <<
	SHR // >>
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	EXP // ^
	operator_end

	LPAREN // (
	RPAREN // )
	COMMA  // ,
	ASSIGN // =
	COLON  // :
	chars_end

	NUM_DEC  // 10, 10.1
	NUM_HEX  // 0xdeadbeef
	NUM_OCT  // 0o127
	NUM_BIN  // 0b101
	NUM_SCI  // 1e2, 1.2e-7
	WORD
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
	COLON:  ":",
	ASSIGN: "=",

	NUM_DEC: "NUM_DEC",
	NUM_HEX: "NUM_HEX",
	NUM_OCT: "NUM_OCT",
	NUM_BIN: "NUM_BIN",
	NUM_SCI: "NUM_SCI",
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
	case ASSIGN, COLON:
		return 1
	case AS, TO:
		return 2
	case ADD, SUB, OR, XOR:
		return 3
	case MUL, QUO, REM, SHL, SHR, AND:
		return 4
	case EXP:
		return 5
	}
	return LowestPrec
}

var StringToOperator = map[string]Token{}

var SpecialChars string

func init() {
	for x := operator_beg + 1; x < operator_end; x++ {
		StringToOperator[strings.ToLower(x.String())] = x
	}

	for x := chars_beg + 1; x < chars_end; x++ {
		SpecialChars += x.String()
	}
}
