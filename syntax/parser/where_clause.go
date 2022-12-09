package parser

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Parser) parseWhereClause() (*ast.WhereClause, error) {
	whereKeyword := p.tok
	p.advance() // consume WHERE

	expression, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	return &ast.WhereClause{
		Keywords: ast.WhereClauseKeywords{
			Where: whereKeyword,
		},
		Expression: expression,
	}, nil
}
