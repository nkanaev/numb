package parser

import (
	"errors"

	"github.com/nkanaev/numb/pkg/ast"
	"github.com/nkanaev/numb/pkg/ratutils"
	"github.com/nkanaev/numb/pkg/scanner"
	"github.com/nkanaev/numb/pkg/timeutil"
	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/value"
)

type parser struct {
	s *scanner.Scanner
}

type syntaxerror struct {
	Err              string
	PosStart, PosEnd int
}

func (e *syntaxerror) Pos() (int, int) {
	return e.PosStart, e.PosEnd
}

func (e *syntaxerror) Error() string {
	return e.Err
}

func (p *parser) expect(t token.Token) {
	if p.s.Token != t {
		msg := "expected " + t.String() + ", got " + p.s.Token.String()
		panic(&syntaxerror{msg, p.s.TokenStart, p.s.TokenEnd})
	}
	p.s.Scan()
}

func (p *parser) parsePrimaryExpr() ast.Node {
	switch p.s.Token {
	case token.LPAREN:
		p.s.Scan()
		expr := p.parseExpr()
		p.expect(token.RPAREN)
		return &ast.ParenExpr{Expr: expr}
	case token.NUM_DEC:
		val := p.s.Value
		p.s.Scan()
		return &ast.Literal{value.Number{Num: ratutils.Must(val)}}
	case token.NUM_HEX:
		val := p.s.Value
		p.s.Scan()
		return &ast.Literal{value.Number{Num: ratutils.Must(val), Fmt: value.HEX}}
	case token.NUM_OCT:
		val := p.s.Value
		p.s.Scan()
		return &ast.Literal{value.Number{Num: ratutils.Must(val), Fmt: value.OCT}}
	case token.NUM_BIN:
		val := p.s.Value
		p.s.Scan()
		return &ast.Literal{value.Number{Num: ratutils.Must(val), Fmt: value.BIN}}
	case token.NUM_SCI:
		val := p.s.Value
		p.s.Scan()
		return &ast.Literal{value.Number{Num: ratutils.Must(val), Fmt: value.SCI}}
	case token.DATE:
		t, err := timeutil.Parse(p.s.Value)
		if err != nil {
			panic(&syntaxerror{err.Error(), p.s.TokenStart, p.s.TokenEnd})
		}
		p.s.Scan()
		return &ast.Literal{value.NewTime(t)}
	case token.IDENT:
		name := p.s.Value
		p.s.Scan()
		if p.s.Token == token.LPAREN {
			p.s.Scan()
			args := make([]ast.Node, 0)
			var arg ast.Node

		argloop:
			for {
				switch p.s.Token {
				case token.COMMA:
					if arg == nil {
						panic(&syntaxerror{"missing expression before comma", p.s.TokenStart, p.s.TokenStart})
					}
					p.s.Scan()
					args = append(args, arg)
					arg = nil
				case token.RPAREN:
					p.s.Scan()
					if arg != nil {
						args = append(args, arg)
						arg = nil
					}
					break argloop
				case token.END:
					panic(&syntaxerror{"unexpected end of line, expected )", p.s.TokenStart, p.s.TokenStart})
				default:
					arg = p.parseExpr()
				}
			}
			return &ast.FunCall{Name: name, Args: args}
		}
		return &ast.Var{Name: name}
	case token.Illegal:
		panic(p.s.Error)
	default:
		panic(&syntaxerror{"unexpected token: " + p.s.Token.String(), p.s.TokenStart, p.s.TokenEnd})
	}
}

func (p *parser) parseUnaryExpr() ast.Node {
	if p.s.Token == token.ADD || p.s.Token == token.SUB {
		tok := p.s.Token
		p.s.Scan()
		return &ast.Unary{Op: tok, Expr: p.parseUnaryExpr()}
	}
	expr := p.parsePrimaryExpr()
	return expr
}

func (p *parser) parseBinaryExpr(prec1 int) ast.Node {
	lhs := p.parseUnaryExpr()
	for {
		tok := p.s.Token
		implicit := false
		prec := tok.Precedence()

		if tok == token.IDENT {
			tok = token.MUL
			implicit = true
			prec = tok.Precedence()
		}

		if prec < prec1 {
			break
		}

		if !implicit {
			p.s.Scan()
		}

		if tok == token.IN {
			name := p.s.Value
			p.expect(token.FORMAT)
			lhs = &ast.Format{Expr: lhs, Fmt: name}
		} else if tok == token.ASSIGN || tok == token.COLON {
			if _, ok := lhs.(*ast.Var); !ok {
				panic("expected varname, got " + lhs.String())
			}
			name := lhs.String()
			rhs := p.parseBinaryExpr(prec)
			lhs = &ast.Assign{Name: name, Expr: rhs, Unit: tok == token.COLON}
		} else {
			rhs := p.parseBinaryExpr(prec + 1)
			lhs = &ast.BinOP{Lhs: lhs, Rhs: rhs, Op: tok, Implicit: implicit}
		}
	}
	return lhs
}

func (p *parser) parseExpr() ast.Node {
	return p.parseBinaryExpr(token.LowestPrec + 1)
}

func Parse(line string) (node ast.Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			node = nil
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown error")
			}
		}
	}()

	s := scanner.New(line)
	p := &parser{s: s}
	s.Scan()
	node = p.parseExpr()
	if s.Token == token.Illegal {
		panic(s.Error)
	}
	if s.Token != token.END {
		panic(&syntaxerror{"invalid syntax", s.TokenStart, s.TokenStart})
	}
	return
}

func Eval(expr string, env map[string]value.Value) (val value.Value, err error) {
	node, err := Parse(expr)
	if err != nil {
		return nil, err
	}
	return node.Eval(env)
}
