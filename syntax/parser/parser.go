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
	p.next = p.scanAndSkipComments()
}

func (p *Parser) scanAndSkipComments() (tok token.Token) {
	for {
		tok = p.scanner.Scan()
		if !tok.Kind.IsComment() {
			return tok
		}
		p.comms = append(p.comms, ast.Comment{
			Content: tok,
		})
	}
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
