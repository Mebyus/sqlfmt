package parser

import (
	"fmt"
	"os"

	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) Parse() (file ast.SQLFile, err error) {
	err = p.parse()
	if err != nil {
		return
	}
	file.Statements = p.stmts
	file.Comments = p.comms
	return
}

func (p *Parser) parse() (err error) {
	for !p.isEOF() {
		err = p.parseStatement()
		if err != nil {
			fmt.Fprintln(os.Stderr, p.pos, p.kind, err)
			p.consumeFlawedStatement()
		}
	}
	return
}

func (p *Parser) consumeFlawedStatement() {
	errorIndex := len(p.stored)
	errorToken := p.tok

	for !p.isEOF() && !p.isSemi() {
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
		Tokens: p.collect(),
	})
}

func (p *Parser) consumeUnknownStatement() {
	for !p.isEOF() && !p.isSemi() {
		p.advance()
	}
	if !p.isEOF() {
		p.advance()
	}
	p.stmts = append(p.stmts, ast.UnknownStatement{
		Tokens: p.collect(),
	})
}

func (p *Parser) collect() []token.Token {
	tokens := p.stored
	p.stored = nil
	return tokens
}

func (p *Parser) start() {
	p.stored = nil
	p.pos = p.tok.Pos
}

func (p *Parser) parseStatement() (err error) {
	p.start()

	switch p.tok.Kind {
	case token.Create:
		return p.parseCreateStatement()
	case token.Comment:
		return p.parseSetCommentStatement()
	default:
		p.consumeUnknownStatement()
		return nil
	}
}

func (p *Parser) consume(kind token.Kind) (token.Token, error) {
	if p.tok.Kind == kind {
		tok := p.tok
		p.advance()
		return tok, nil
	}
	return token.Token{}, fmt.Errorf("expected [ %v ], got [ %v ]", kind, p.tok.String())
}
