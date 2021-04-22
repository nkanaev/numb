package parser

import (
	"fmt"

	"github.com/nkanaev/numb/pkg/ast"
	"github.com/nkanaev/numb/pkg/scanner"
	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/value"
)

type parser struct {
	s *scanner.Scanner
}

func (p *parser) expect(t token.Token) {
	if p.s.Token != t {
		panic("expected " + t.String())
	}
	p.s.Scan()
}

func (p *parser) parseOperand() ast.Node {
	switch p.s.Token {
	case token.LPAREN:
		p.s.Scan()
		expr := p.parseExpr()
		p.expect(token.RPAREN)
		return &ast.ParenExpr{Expr: expr}
	case token.NUM:
		val := p.s.Value
		p.s.Scan()
		return value.Parse(val)
	case token.VAR:
		name := p.s.Value
		p.expect(token.VAR)
		if p.s.Token == token.ASSIGN {
			p.expect(token.ASSIGN)
			expr := p.parseExpr()
			return &ast.Assign{Name: name, Expr: expr}
		}
		return &ast.Var{Name: name}
	}
	fmt.Println(p.s.Token)
	panic("die")
}

func (p *parser) parsePrimaryExpr() ast.Node {
	expr := p.parseOperand()
	for {
		if p.s.Token == token.KEYWORD && p.s.Value == "as" {
			p.expect(token.KEYWORD)
			if p.s.Token != token.VAR {
				panic("expected format")
			}
			f, ok := value.StringToNumeral[p.s.Value]
			if !ok {
				panic("unknown format: " + p.s.Value)
			}
			expr = &ast.Format{Expr: expr, Fmt: f}
			p.expect(token.VAR)
			continue
		}
		break
	}
	return expr
}

func (p *parser) parseUnaryExpr() ast.Node {
	if p.s.Token == token.ADD || p.s.Token == token.SUB {
		tok := p.s.Token
		p.s.Scan()
		return &ast.Unary{Op: tok, Expr: p.parseExpr()}
	}
	return p.parsePrimaryExpr()
}

func (p *parser) parseBinaryExpr(prec1 int) ast.Node {
	lhs := p.parseUnaryExpr()
	for {
		prec := p.s.Token.Precedence()
		if prec < prec1 {
			break
		}
		tok := p.s.Token
		p.s.Scan()
		rhs := p.parseBinaryExpr(prec + 1)
		lhs = &ast.BinOP{Lhs: lhs, Rhs: rhs, Op: tok}
	}
	return lhs
}

func (p *parser) parseExpr() ast.Node {
	return p.parseBinaryExpr(token.LowestPrec + 1)
}

func Parse(line string) ast.Node {
	s := scanner.New(line)
	p := &parser{s: s}
	s.Scan()
	return p.parseExpr()
}
