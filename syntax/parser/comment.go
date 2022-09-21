package parser

import (
	"fmt"

	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/ast/statement"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) parseCommentStatement() error {
	p.kind = statement.Comment

	p.advance() // consume COMMENT
	err := p.consume(token.On)
	if err != nil {
		return err
	}
	var stmt ast.Statement
	if p.tok.Kind == token.Column {
		stmt, err = p.parseColumnCommentStatement()
		if err != nil {
			return err
		}
	} else if p.tok.Kind == token.Table {
		stmt, err = p.parseTableCommentStatement()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unexpected token [ %v ]", p.tok)
	}
	p.stmts = append(p.stmts, stmt)
	return nil
}

func (p *Parser) parseTableCommentStatement() (stmt ast.TableCommentStatement, err error) {
	p.kind = statement.TableComment

	p.advance() // consume "TABLE"
	if p.tok.Kind != token.Identifier && p.tok.Kind != token.QuotedIdentifier {
		return ast.TableCommentStatement{},
			fmt.Errorf("expected identifier, got [ %v ]", p.tok)
	}
	firstPartOfName := p.tok
	p.advance()

	var tableName ast.TableName
	if p.tok.Kind == token.Dot {
		p.advance()
		if p.tok.Kind != token.Identifier && p.tok.Kind != token.QuotedIdentifier {
			return ast.TableCommentStatement{},
				fmt.Errorf("expected identifier, got [ %v ]", p.tok)
		}
		tableName = ast.QualifiedIdentifier{
			SchemaName: ast.Identifier{
				Token: firstPartOfName,
			},
			RawTableName: ast.Identifier{
				Token: p.tok,
			},
		}
		p.advance()
	} else {
		tableName = ast.Identifier{
			Token: firstPartOfName,
		}
	}
	err = p.consume(token.Is)
	if err != nil {
		return ast.TableCommentStatement{}, err
	}

	if p.tok.Kind != token.String {
		return ast.TableCommentStatement{},
			fmt.Errorf("expected string, got [ %v ]", p.tok)
	}
	comment := p.tok
	p.advance()

	err = p.consume(token.Semicolon)
	if err != nil {
		return ast.TableCommentStatement{}, err
	}
	return ast.TableCommentStatement{
		TableName: tableName,
		Comment:   comment,
	}, nil
}

func (p *Parser) parseColumnCommentStatement() (stmt ast.ColumnCommentStatement, err error) {
	p.kind = statement.ColumnComment

	p.advance() // consume "COLUMN"
	if p.tok.Kind != token.Identifier && p.tok.Kind != token.QuotedIdentifier {
		return ast.ColumnCommentStatement{},
			fmt.Errorf("expected identifier, got [ %v ]", p.tok)
	}
	firstPartOfName := p.tok
	p.advance()
	err = p.consume(token.Dot)
	if err != nil {
		return ast.ColumnCommentStatement{},
			err
	}
	if p.tok.Kind != token.Identifier && p.tok.Kind != token.QuotedIdentifier {
		return ast.ColumnCommentStatement{},
			fmt.Errorf("expected identifier, got [ %v ]", p.tok)
	}
	secondPartOfName := p.tok
	p.advance()

	var tableName ast.TableName
	var columnName token.Token
	if p.tok.Kind == token.Dot {
		p.advance()
		if p.tok.Kind != token.Identifier && p.tok.Kind != token.QuotedIdentifier {
			return ast.ColumnCommentStatement{},
				fmt.Errorf("expected identifier, got [ %v ]", p.tok)
		}
		tableName = ast.QualifiedIdentifier{
			SchemaName: ast.Identifier{
				Token: firstPartOfName,
			},
			RawTableName: ast.Identifier{
				Token: secondPartOfName,
			},
		}
		columnName = p.tok
		p.advance()
	} else {
		tableName = ast.Identifier{
			Token: firstPartOfName,
		}
		columnName = secondPartOfName
	}

	err = p.consume(token.Is)
	if err != nil {
		return ast.ColumnCommentStatement{}, err
	}

	if p.tok.Kind != token.String {
		return ast.ColumnCommentStatement{},
			fmt.Errorf("expected string, got [ %v ]", p.tok)
	}
	comment := p.tok
	p.advance()

	err = p.consume(token.Semicolon)
	if err != nil {
		return ast.ColumnCommentStatement{}, err
	}
	return ast.ColumnCommentStatement{
		TableName: tableName,
		Name:      columnName,
		Comment:   comment,
	}, nil
}
