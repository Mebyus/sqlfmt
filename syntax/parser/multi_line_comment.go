package parser

import (
	"github.com/mebyus/sqlfmt/syntax/ast"
)

func (p *Parser) parseMultiLineComment() error {
	p.comms = append(p.comms, ast.Comment{
		Content: p.tok,
	})
	p.advance()
	return nil
}
