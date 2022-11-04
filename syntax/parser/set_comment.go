package parser

import (
	"fmt"

	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/ast/statement"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) parseSetCommentStatement() error {
	p.kind = statement.Comment

	commentKeyword := p.tok
	p.advance() // consume COMMENT
	onKeyword, err := p.consume(token.On)
	if err != nil {
		return err
	}

	var stmt ast.Statement
	if p.tok.Kind == token.Column {
		colStmt, err := p.parseSetColumnCommentStatement()
		if err != nil {
			return err
		}
		colStmt.Keywords.Comment = commentKeyword
		colStmt.Keywords.On = onKeyword
		stmt = colStmt
	} else if p.tok.Kind == token.Table {
		tableStmt, err := p.parseSetTableCommentStatement()
		if err != nil {
			return err
		}
		tableStmt.Keywords.Comment = commentKeyword
		tableStmt.Keywords.On = onKeyword
		stmt = tableStmt
	} else {
		return fmt.Errorf("unexpected token [ %v ]", p.tok)
	}
	p.stmts = append(p.stmts, stmt)
	return nil
}

func (p *Parser) parseSetTableCommentStatement() (stmt ast.SetTableCommentStatement, err error) {
	p.kind = statement.TableComment

	tableKeyword := p.tok
	p.advance() // consume "TABLE"
	if p.tok.Kind != token.Identifier && p.tok.Kind != token.QuotedIdentifier {
		return ast.SetTableCommentStatement{},
			fmt.Errorf("expected identifier, got [ %v ]", p.tok)
	}
	firstPartOfName := p.tok
	p.advance()

	var tableName ast.TableName
	if p.tok.Kind == token.Dot {
		dot := p.tok
		p.advance()
		if p.tok.Kind != token.Identifier && p.tok.Kind != token.QuotedIdentifier {
			return ast.SetTableCommentStatement{},
				fmt.Errorf("expected identifier, got [ %v ]", p.tok)
		}
		tableName = ast.QualifiedIdentifier{
			Dot: dot,
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
	isKeyword, err := p.consume(token.Is)
	if err != nil {
		return ast.SetTableCommentStatement{}, err
	}

	if p.tok.Kind != token.String {
		return ast.SetTableCommentStatement{},
			fmt.Errorf("expected string, got [ %v ]", p.tok)
	}
	comment := p.tok
	p.advance()

	semicolon, err := p.consume(token.Semicolon)
	if err != nil {
		return ast.SetTableCommentStatement{}, err
	}
	return ast.SetTableCommentStatement{
		Keywords: ast.SetTableCommentKeywords{
			Table: tableKeyword,
			Is:    isKeyword,
		},
		Semicolon:     semicolon,
		TableName:     tableName,
		CommentString: comment,
	}, nil
}

func (p *Parser) parseSetColumnCommentStatement() (stmt ast.SetColumnCommentStatement, err error) {
	p.kind = statement.ColumnComment

	columnKeyword := p.tok
	p.advance() // consume "COLUMN"
	if p.tok.Kind != token.Identifier && p.tok.Kind != token.QuotedIdentifier {
		return ast.SetColumnCommentStatement{},
			fmt.Errorf("expected identifier, got [ %v ]", p.tok)
	}
	firstPartOfName := p.tok
	p.advance()
	var dot token.Token
	firstDot, err := p.consume(token.Dot)
	if err != nil {
		return ast.SetColumnCommentStatement{},
			err
	}
	if p.tok.Kind != token.Identifier && p.tok.Kind != token.QuotedIdentifier {
		return ast.SetColumnCommentStatement{},
			fmt.Errorf("expected identifier, got [ %v ]", p.tok)
	}
	secondPartOfName := p.tok
	p.advance()

	var tableName ast.TableName
	var columnName token.Token
	if p.tok.Kind == token.Dot {
		secondDot := p.tok
		p.advance()
		if p.tok.Kind != token.Identifier && p.tok.Kind != token.QuotedIdentifier {
			return ast.SetColumnCommentStatement{},
				fmt.Errorf("expected identifier, got [ %v ]", p.tok)
		}
		tableName = ast.QualifiedIdentifier{
			Dot: firstDot,
			SchemaName: ast.Identifier{
				Token: firstPartOfName,
			},
			RawTableName: ast.Identifier{
				Token: secondPartOfName,
			},
		}
		columnName = p.tok
		p.advance()
		dot = secondDot
	} else {
		tableName = ast.Identifier{
			Token: firstPartOfName,
		}
		columnName = secondPartOfName
		dot = firstDot
	}

	isKeyword, err := p.consume(token.Is)
	if err != nil {
		return ast.SetColumnCommentStatement{}, err
	}

	if p.tok.Kind != token.String {
		return ast.SetColumnCommentStatement{},
			fmt.Errorf("expected string, got [ %v ]", p.tok)
	}
	comment := p.tok
	p.advance()

	semicolon, err := p.consume(token.Semicolon)
	if err != nil {
		return ast.SetColumnCommentStatement{}, err
	}
	return ast.SetColumnCommentStatement{
		Keywords: ast.SetColumnCommentKeywords{
			Column: columnKeyword,
			Is:     isKeyword,
		},
		Dot:           dot,
		Semicolon:     semicolon,
		TableName:     tableName,
		ColumnName:    columnName,
		CommentString: comment,
	}, nil
}
