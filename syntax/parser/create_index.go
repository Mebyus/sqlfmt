package parser

import (
	"fmt"

	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/ast/statement"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) parseCreateIndexStatement() error {
	p.kind = statement.CreateIndex

	createKeyword := p.tok
	indexKeyword := p.next
	p.advance() // consume CREATE
	p.advance() // consume INDEX

	var indexName *ast.Identifier
	if p.isIdent() {
		indexName = &ast.Identifier{
			Token: p.tok,
		}
		p.advance()
	}

	onKeyword, err := p.consume(token.On)
	if err != nil {
		return err
	}

	tableName, err := p.parseObjectName()
	if err != nil {
		return err
	}

	var usingClause *ast.IndexUsingClause
	if p.tok.Kind == token.Using {
		usingKeyword := p.tok
		p.advance() // consume USING

		methodName, err := p.consume(token.Identifier)
		if err != nil {
			return err
		}

		usingClause = &ast.IndexUsingClause{
			Keywords: ast.IndexUsingClauseKeywords{
				Using: usingKeyword,
			},
			MethodName: methodName,
		}
	}

	leftParentheses, err := p.consume(token.LeftParentheses)
	if err != nil {
		return err
	}

	var columns []ast.IndexColumn
	var column ast.IndexColumn
	for {
		if !p.isIdent() {
			return fmt.Errorf("expected identifier, got [ %v ]", p.tok)
		}
		column.Name = ast.Identifier{
			Token: p.tok,
		}
		p.advance()
		columns = append(columns, column)

		if p.tok.Kind != token.Comma {
			break
		}
		tok := p.tok
		column.Comma = &tok
		p.advance() // consume ","
	}

	rightParentheses, err := p.consume(token.RightParentheses)
	if err != nil {
		return err
	}

	var whereClause *ast.WhereClause
	if p.tok.Kind == token.Where {
		whereKeyword := p.tok
		p.advance() // consume WHERE

		expression, err := p.parseExpression()
		if err != nil {
			return err
		}

		whereClause = &ast.WhereClause{
			Keywords: ast.WhereClauseKeywords{
				Where: whereKeyword,
			},
			Expression: expression,
		}
	}

	semicolon, err := p.consume(token.Semicolon)
	if err != nil {
		return err
	}

	stmt := ast.CreateIndexStatement{
		Keywords: ast.CreateIndexKeywords{
			Create: createKeyword,
			Index:  indexKeyword,
			On:     onKeyword,
		},
		LeftParentheses:  leftParentheses,
		RightParentheses: rightParentheses,
		Name:             indexName,
		TableName:        tableName,
		Semicolon:        semicolon,
		Columns:          columns,
		Where:            whereClause,
		Using:            usingClause,
	}

	p.stmts = append(p.stmts, stmt)
	return nil
}
