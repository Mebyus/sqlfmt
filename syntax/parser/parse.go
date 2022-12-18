package parser

import (
	"fmt"
	"os"

	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/scanner"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) Parse() (file ast.SQLFile, err error) {
	err = p.parse()
	if err != nil {
		return
	}
	file.Statements = p.stmts
	file.Comments = p.comms
	file.NumberOfErrors = p.nerr
	return
}

func (p *Parser) parse() error {
	for !p.isEOF() {
		err := p.parseStatement()
		if err != nil {
			if p.nerr <= p.options.MaxErrors {
				fmt.Fprintln(os.Stderr, "error:", p.pos, p.kind, ">>", err)
			}
			p.nerr++
			if !p.options.KeepGoing && p.nerr > p.options.MaxErrors {
				return fmt.Errorf("too many errors")
			}

			p.consumeFlawedStatement()
		}
	}
	return nil
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
	case token.Update:
		return p.parseUpdateStatement()
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

// ParseExpression is a pure function for usage in unit tests
func ParseExpression(str string) (ast.Expression, error) {
	p := FromStream(scanner.NewEraser(scanner.FromString(str)))
	return p.parseExpression()
}
