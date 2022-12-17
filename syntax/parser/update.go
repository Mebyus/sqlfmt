package parser

import (
	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/ast/statement"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) parseUpdateStatement() error {
	p.kind = statement.Update

	updateKeyword := p.tok
	p.advance() // consume UPDATE

	var onlyKeyword *token.Token
	if p.tok.Kind == token.Only {
		tok := p.tok
		p.advance() // consume ONLY

		onlyKeyword = &tok
	}

	tableName, err := p.parseObjectName()
	if err != nil {
		return err
	}

	setKeyword, err := p.consume(token.Set)
	if err != nil {
		return err
	}

	var elements []ast.UpdateListElement
	var elem ast.UpdateListElement
	for {
		assignment, err := p.parseSingleColumnAssignment()
		if err != nil {
			return err
		}
		elem.Assignment = assignment
		elements = append(elements, elem)

		if !p.isComma() {
			break
		}
		tok := p.tok
		elem.Comma = &tok
		p.advance() // consume ","
	}

	var whereClause *ast.WhereClause
	if p.tok.Kind == token.Where {
		whereClause, err = p.parseWhereClause()
		if err != nil {
			return err
		}
	}

	semicolon, err := p.consume(token.Semicolon)
	if err != nil {
		return err
	}

	stmt := ast.UpdateStatement{
		Keywords: ast.UpdateStatementKeywords{
			Update: updateKeyword,
			Only:   onlyKeyword,
			Set:    setKeyword,
		},
		TableName: tableName,
		Elements:  elements,
		Where:     whereClause,
		Semicolon: semicolon,
	}
	p.stmts = append(p.stmts, stmt)
	return nil
}

func (p *Parser) parseSingleColumnAssignment() (ast.SingleColumnAssignment, error) {
	name, err := p.consumeIdentifier()
	if err != nil {
		return ast.SingleColumnAssignment{}, err
	}

	equal, err := p.consume(token.Equal)
	if err != nil {
		return ast.SingleColumnAssignment{}, err
	}

	expression, err := p.parseDefaultableExpression()
	if err != nil {
		return ast.SingleColumnAssignment{}, err
	}

	return ast.SingleColumnAssignment{
		Name:  name,
		Equal: equal,
		Value: expression,
	}, nil
}

func (p *Parser) parseDefaultableExpression() (ast.DefaultableExpression, error) {
	if p.tok.Kind == token.Default {
		defaultKeyword := p.tok
		p.advance() // consume DEFAULT

		return ast.DefaultableExpression{
			Default: &defaultKeyword,
		}, nil
	}

	expression, err := p.parseExpression()
	if err != nil {
		return ast.DefaultableExpression{}, err
	}

	return ast.DefaultableExpression{
		Expression: &expression,
	}, nil
}
