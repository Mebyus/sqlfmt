package parser

import (
	"fmt"

	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) parseExpression() (ast.Expression, error) {
	var tokens []token.Token
	for !p.isEOF() && !p.isSemi() && !p.isComma() && p.tok.Kind != token.Where {
		tokens = append(tokens, p.tok)
		p.advance()
	}

	if len(tokens) == 0 {
		return ast.Expression{}, fmt.Errorf("empty expression")
	}

	return ast.Expression{
		Tokens: tokens,
	}, nil
}
