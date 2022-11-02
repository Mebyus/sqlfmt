package parser

import (
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

	stored []token.Token
}

func (p *Parser) advance() {
	p.stored = append(p.stored, p.tok)
	p.tok = p.next
	p.next = p.scanner.Scan()
}

func (p *Parser) isEOF() bool {
	return p.tok.Kind == token.EOF
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

func ParseBytes(b []byte) (stmts []ast.Statement, err error) {
	p := FromBytes(b)
	stmts, err = p.Parse()
	return
}

func ParseFile(filename string) (stmts []ast.Statement, err error) {
	p, err := FromFile(filename)
	if err != nil {
		return
	}
	stmts, err = p.Parse()
	return
}

func Parse(r io.Reader) (stmts []ast.Statement, err error) {
	p, err := FromReader(r)
	if err != nil {
		return
	}
	stmts, err = p.Parse()
	return
}
