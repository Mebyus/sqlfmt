package parser

import "github.com/mebyus/sqlfmt/syntax/token"

func (p *Parser) parseCreateStatement() error {
	if p.next.Kind == token.Index {
		return p.parseCreateIndexStatement()
	}

	p.consumeUnknownStatement()
	return nil
}
