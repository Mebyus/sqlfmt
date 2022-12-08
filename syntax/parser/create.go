package parser

import (
	"github.com/mebyus/sqlfmt/syntax/ast/statement"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) parseCreateStatement() error {
	p.kind = statement.Create

	if p.next.Kind == token.Index {
		return p.parseCreateIndexStatement()
	}
	if p.next.Kind == token.Table {
		return p.parseCreateTableStatement()
	}

	p.consumeUnknownStatement()
	return nil
}
