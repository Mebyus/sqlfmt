package parser

import (
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
		name, err := p.consumeIdentifier()
		if err != nil {
			return err
		}
		column.Name = name
		columns = append(columns, column)

		if !p.isComma() {
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
		whereClause, err = p.parseWhereClause()
		if err != nil {
			return err
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
