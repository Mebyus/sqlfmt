package parser

import (
	"github.com/mebyus/sqlfmt/syntax/ast/statement"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) parseCreateStatement() error {
	p.kind = statement.Create

	for !p.isEOF() && p.tok.Kind != token.Semicolon {
		p.advance()
	}
	p.advance()
	return nil
}
