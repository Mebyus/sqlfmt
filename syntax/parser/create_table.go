package parser

import (
	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/ast/statement"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) parseCreateTableStatement() error {
	p.kind = statement.CreateTable

	createKeyword := p.tok
	tableKeyword := p.next
	p.advance() // consume CREATE
	p.advance() // consume TABLE

	var temporaryKeyword, tempKeyword *token.Token
	if p.tok.Kind == token.Temporary {
		tok := p.tok
		temporaryKeyword = &tok
		p.advance()
	} else if p.tok.Kind == token.Temp {
		tok := p.tok
		tempKeyword = &tok
		p.advance()
	}

	tableName, err := p.parseObjectName()
	if err != nil {
		return err
	}

	leftParentheses, err := p.consume(token.LeftParentheses)
	if err != nil {
		return err
	}

	rightParentheses, err := p.consume(token.RightParentheses)
	if err != nil {
		return err
	}

	var tablespaceClause *ast.TablespaceClause
	if p.tok.Kind == token.Tablespace {
		tablespaceClause, err = p.parseTablespaceClause()
		if err != nil {
			return err
		}
	}

	semicolon, err := p.consume(token.Semicolon)
	if err != nil {
		return err
	}

	stmt := ast.CreateTableStatement{
		Keywords: ast.CreateTableKeywords{
			Create:    createKeyword,
			Table:     tableKeyword,
			Temporary: temporaryKeyword,
			Temp:      tempKeyword,
		},
		LeftParentheses:  leftParentheses,
		RightParentheses: rightParentheses,
		Semicolon:        semicolon,
		Name:             tableName,
		Tablespace:       tablespaceClause,
	}

	p.stmts = append(p.stmts, stmt)
	return nil
}
