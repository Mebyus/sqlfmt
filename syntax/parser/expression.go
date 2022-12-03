package parser

import (
	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) parseExpression() (ast.Expression, error) {
	var tokens []token.Token
	for !p.isEOF() && !p.isSemi() {
		tokens = append(tokens, p.tok)
		p.advance()
	}
	return ast.Expression{
		Tokens: tokens,
	}, nil
}
