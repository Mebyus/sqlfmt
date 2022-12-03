package parser

import (
	"fmt"

	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) parseObjectName() (ast.ObjectName, error) {
	if !p.isIdent() {
		return nil, fmt.Errorf("expected identifier, got [ %v ]", p.tok)
	}
	firstPartOfName := p.tok
	p.advance()

	if p.tok.Kind != token.Dot {
		return ast.Identifier{
			Token: firstPartOfName,
		}, nil
	}

	dot := p.tok
	p.advance()
	if !p.isIdent() {
		return nil, fmt.Errorf("expected identifier, got [ %v ]", p.tok)
	}
	objectName := ast.QualifiedIdentifier{
		Dot: dot,
		SchemaName: ast.Identifier{
			Token: firstPartOfName,
		},
		RawTableName: ast.Identifier{
			Token: p.tok,
		},
	}
	p.advance()
	return objectName, nil
}
