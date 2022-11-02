package parser

import (
	"fmt"

	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) Parse() (stmts []ast.Statement, err error) {
	err = p.parse()
	if err != nil {
		return
	}
	stmts = p.stmts
	return
}

func (p *Parser) parse() (err error) {
	for !p.isEOF() {
		err = p.parseStatement()
		if err != nil {
			fmt.Println(err)
			p.consumeFlawedStatement()
		}
	}
	return
}

func (p *Parser) consumeFlawedStatement() {
	errorIndex := len(p.stored)
	errorToken := p.tok

	for !p.isEOF() && p.tok.Kind != token.Semicolon {
		p.advance()
	}
	if !p.isEOF() {
		p.advance()
	}
	p.stmts = append(p.stmts, ast.FlawedStatement{
		Error: ast.Error{
			Index:     errorIndex,
			Token:     errorToken,
			Statement: p.kind,
		},
		Tokens: p.stored,
	})
}

func (p *Parser) consumeUnknownStatement() {
	for !p.isEOF() && p.tok.Kind != token.Semicolon {
		p.advance()
	}
	if !p.isEOF() {
		p.advance()
	}
	p.stmts = append(p.stmts, ast.UnknownStatement{
		Tokens: p.stored,
	})
}

func (p *Parser) startNewStatement() {
	p.stored = nil
}

func (p *Parser) parseStatement() (err error) {
	p.startNewStatement()

	switch p.tok.Kind {
	case token.Create:
		return p.parseCreateStatement()
	case token.Comment:
		return p.parseCommentStatement()
	case token.LineComment:
		return p.parseLineComment()
	case token.MultiLineComment:
		return p.parseMultiLineComment()
	default:
		p.consumeUnknownStatement()
		return nil
	}
}

func (p *Parser) consume(kind token.Kind) error {
	if p.tok.Kind == kind {
		p.advance()
		return nil
	}
	return fmt.Errorf("expected [ %v ], got [ %v ]", kind, p.tok.String())
}
