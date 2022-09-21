package parser

import (
	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/ast/statement"
)

func (p *Parser) parseLineComment() error {
	p.kind = statement.LineComment

	p.stmts = append(p.stmts, ast.LineComment{
		Content: p.tok,
	})
	p.advance()
	return nil
}
