package parser

import (
	"github.com/nkanaev/numb/pkg/ast"
	"github.com/nkanaev/numb/pkg/scanner"
	"github.com/nkanaev/numb/pkg/token"
	"github.com/nkanaev/numb/pkg/unit"
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

func (p *parser) parsePrimaryExpr() ast.Node {
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
		if p.s.Token == token.LPAREN {
			args := make([]ast.Node, 0)
			p.expect(token.LPAREN)
			for p.s.Token != token.RPAREN {
				args = append(args, p.parseExpr())
				if p.s.Token == token.COMMA {
					p.expect(token.COMMA)
					continue
				}
			}
			p.expect(token.RPAREN)
			return &ast.FunCall{Name: name, Args: args}
		}
		return &ast.Var{Name: name}
	case token.END:
		panic("unexpected end")
	case token.Illegal:
		panic("illegal char")
	}
	panic("unexpected token: " + p.s.Token.String())
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
		tok := p.s.Token

		if tok == token.VAR {
			u := unit.Get(p.s.Value)
			if u == nil {
				panic("unknown unit: " + p.s.Value)
			}
			p.s.Scan()
			lhs = &ast.Unit{Expr: lhs, Unit: u}
			continue
		}

		prec := tok.Precedence()
		if prec < prec1 {
			break
		}
		p.s.Scan()
		rhs := p.parseBinaryExpr(prec + 1)
		lhs = &ast.BinOP{Lhs: lhs, Rhs: rhs, Op: tok}
	}
	return lhs
}

func (p *parser) parseExpr() ast.Node {
	return p.parseBinaryExpr(token.LowestPrec + 1)
}

func (p *parser) parseRoot() ast.Node {
	lhs := p.parseExpr()
	for {
		tok := p.s.Token
		if tok == token.AS {
			p.expect(token.AS)
			if p.s.Token != token.VAR {
				panic("expected format")
			}
			f, ok := value.StringToNumeral[p.s.Value]
			if !ok {
				panic("unknown format: " + p.s.Value)
			}
			lhs = &ast.Format{Expr: lhs, Fmt: f}
			p.expect(token.VAR)
			continue
		}
		if tok == token.TO {
			p.expect(token.TO)
			if p.s.Token != token.VAR {
				panic("expected unit")
			}
			u := unit.Get(p.s.Value)
			if u == nil {
				panic("unknown unit: " + p.s.Value)
			}
			lhs = &ast.Convert{Expr: lhs, Unit: u}
			p.expect(token.VAR)
			continue
		}
		break
	}
	return lhs
}

func Parse(line string) ast.Node {
	s := scanner.New(line)
	p := &parser{s: s}
	s.Scan()
	root := p.parseRoot()
	if s.Token != token.END {
		panic("trailing stuff")
	}
	return root
}
