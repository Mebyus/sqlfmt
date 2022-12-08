package parser

import (
	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) parseIdentifierList() (ast.IdentifierList, error) {
	leftParentheses, err := p.consume(token.LeftParentheses)
	if err != nil {
		return ast.IdentifierList{}, err
	}

	var elements []ast.IdentifierElement
	var elem ast.IdentifierElement
	for {
		name, err := p.consumeIdentifier()
		if err != nil {
			return ast.IdentifierList{}, err
		}
		elem.Name = name
		elements = append(elements, elem)

		if !p.isComma() {
			break
		}
		tok := p.tok
		elem.Comma = &tok
		p.advance() // consume ","
	}

	rightParentheses, err := p.consume(token.RightParentheses)
	if err != nil {
		return ast.IdentifierList{}, err
	}

	return ast.IdentifierList{
		LeftParentheses:  leftParentheses,
		RightParentheses: rightParentheses,
		Elements:         elements,
	}, nil
}
