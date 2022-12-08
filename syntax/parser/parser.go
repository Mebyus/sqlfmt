package parser

import (
	"fmt"
	"io"

	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/ast/statement"
	"github.com/mebyus/sqlfmt/syntax/scanner"
	"github.com/mebyus/sqlfmt/syntax/token"
)

type Parser struct {
	scanner *scanner.Scanner
	stmts   []ast.Statement
	comms   []ast.Comment

	// token at current Parser position
	tok token.Token

	// next token
	next token.Token

	// kind of statement which is being parsed at current position
	kind statement.Kind

	// position of first token in current statement
	pos token.Pos

	stored []token.Token
}

func (p *Parser) advance() {
	p.stored = append(p.stored, p.tok)
	p.tok = p.next
	p.next = p.scanAndSkipComments()
}

func (p *Parser) scanAndSkipComments() (tok token.Token) {
	for {
		tok = p.scanner.Scan()
		if !tok.Kind.IsComment() {
			return tok
		}

		inlined := !p.tok.Kind.IsEmpty() && tok.Pos.Line == p.tok.Pos.Line

		p.comms = append(p.comms, ast.Comment{
			Inlined: inlined,
			Content: tok,
		})
	}
}

func (p *Parser) isEOF() bool {
	return p.tok.Kind == token.EOF
}

func (p *Parser) isSemi() bool {
	return p.tok.Kind == token.Semicolon
}

func (p *Parser) isComma() bool {
	return p.tok.Kind == token.Comma
}

func (p *Parser) isIdent() bool {
	return p.tok.Kind == token.Identifier || p.tok.Kind == token.QuotedIdentifier
}

func (p *Parser) consumeIdentifier() (ast.Identifier, error) {
	if !p.isIdent() {
		return ast.Identifier{}, fmt.Errorf("expected identifier, got [ %v ]", p.tok)
	}
	ident := ast.Identifier{
		Token: p.tok,
	}
	p.advance()
	return ident, nil
}

func FromReader(r io.Reader) (p *Parser, err error) {
	s, err := scanner.FromReader(r)
	if err != nil {
		return
	}
	p = FromScanner(s)
	return
}

func FromScanner(s *scanner.Scanner) (p *Parser) {
	p = &Parser{
		scanner: s,
	}

	// init Parser buffer
	p.advance()
	p.advance()
	return
}

func FromBytes(b []byte) *Parser {
	return FromScanner(scanner.FromBytes(b))
}

func FromFile(filename string) (p *Parser, err error) {
	s, err := scanner.FromFile(filename)
	if err != nil {
		return
	}
	p = FromScanner(s)
	return
}

func ParseBytes(b []byte) (file ast.SQLFile, err error) {
	p := FromBytes(b)
	file, err = p.Parse()
	return
}

func ParseFile(filename string) (file ast.SQLFile, err error) {
	p, err := FromFile(filename)
	if err != nil {
		return
	}
	file, err = p.Parse()
	return
}

func Parse(r io.Reader) (file ast.SQLFile, err error) {
	p, err := FromReader(r)
	if err != nil {
		return
	}
	file, err = p.Parse()
	return
}
