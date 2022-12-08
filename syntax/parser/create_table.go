package parser

import (
	"fmt"

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

	var specs []ast.TablePropertySpecifier
	var spec ast.TablePropertySpecifier
	for {
		property, err := p.parseTableProperty()
		if err != nil {
			return err
		}
		spec.Property = property
		specs = append(specs, spec)

		if !p.isComma() {
			break
		}
		tok := p.tok
		spec.Comma = &tok
		p.advance() // consume ","
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
		Properties:       specs,
	}

	p.stmts = append(p.stmts, stmt)
	return nil
}

func (p *Parser) parseTableProperty() (ast.TableProperty, error) {
	if p.isIdent() {
		spec, err := p.parseColumnSpecifier()
		if err != nil {
			return nil, err
		}
		return spec, nil
	}

	spec, err := p.parseConstraintSpecifier()
	if err != nil {
		return nil, err
	}
	return spec, nil
}

func (p *Parser) parseColumnSpecifier() (ast.ColumnSpecifier, error) {
	name := p.tok
	p.advance()

	var spec ast.TypeSpecifier
	for p.isIdent() || p.tok.Kind == token.With {
		tok := p.tok
		p.advance()

		if tok.Kind == token.With {
			// handle "with time zone" types
			tok.Kind = token.Identifier
			tok.Lit = "with"
		}
		spec.Spec = append(spec.Spec, tok)
	}

	var constraints ast.ColumnConstraints

loop:
	for {
		switch p.tok.Kind {
		case token.Null:
			tok := p.tok
			constraints.Null = &tok
			p.advance() // consume NULL
		case token.Not:
			notNull, err := p.parseNotNullClause()
			if err != nil {
				return ast.ColumnSpecifier{}, err
			}
			constraints.NotNull = notNull
		case token.Primary:
			primaryKey, err := p.parsePrimaryKeyClause()
			if err != nil {
				return ast.ColumnSpecifier{}, err
			}
			constraints.PrimaryKey = primaryKey
		default:
			break loop
		}
	}

	return ast.ColumnSpecifier{
		Name:        ast.Identifier{Token: name},
		Type:        spec,
		Constraints: constraints,
	}, nil
}

func (p *Parser) parseNotNullClause() (*ast.NotNullClause, error) {
	notKeyword := p.tok
	p.advance() // consume NOT

	nullKeyword, err := p.consume(token.Null)
	if err != nil {
		return nil, err
	}

	return &ast.NotNullClause{
		Not:  notKeyword,
		Null: nullKeyword,
	}, nil
}

func (p *Parser) parsePrimaryKeyClause() (*ast.PrimaryKeyClause, error) {
	primaryKeyword := p.tok
	p.advance() // consume PRIMARY

	keyKeyword, err := p.consume(token.Key)
	if err != nil {
		return nil, err
	}

	return &ast.PrimaryKeyClause{
		Primary: primaryKeyword,
		Key:     keyKeyword,
	}, nil
}

func (p *Parser) parseConstraintSpecifier() (ast.ConstraintSpecifier, error) {
	var nameClause *ast.ConstraintNameClause
	var err error

	if p.tok.Kind == token.Constraint {
		nameClause, err = p.parseConstraintNameClause()
		if err != nil {
			return ast.ConstraintSpecifier{}, err
		}
	}

	constraint, err := p.parseTableConstraint()
	if err != nil {
		return ast.ConstraintSpecifier{}, err
	}

	return ast.ConstraintSpecifier{
		Name:       nameClause,
		Constraint: constraint,
	}, nil
}

func (p *Parser) parseTableConstraint() (ast.TableConstraint, error) {
	var constraint ast.TableConstraint
	var err error

	switch p.tok.Kind {
	case token.Unique:
		constraint, err = p.parseUniqueConstraint()
	case token.Foreign:
		constraint, err = p.parseForeignKeyConstraint()
	default:
		return nil, fmt.Errorf("unexpected token [ %v ]", p.tok)
	}

	if err != nil {
		return nil, err
	}
	return constraint, nil
}

func (p *Parser) parseUniqueConstraint() (ast.UniqueConstraint, error) {
	uniqueKeyword := p.tok
	p.advance() // consume UNIQUE

	columns, err := p.parseIdentifierList()
	if err != nil {
		return ast.UniqueConstraint{}, err
	}

	return ast.UniqueConstraint{
		Keywords: ast.UniqueConstraintKeywords{
			Unique: uniqueKeyword,
		},
		Columns: columns,
	}, nil
}

func (p *Parser) parseForeignKeyConstraint() (ast.ForeignKeyConstraint, error) {
	foreignKeyword := p.tok
	p.advance() // consume FOREIGN

	keyKeyword, err := p.consume(token.Key)
	if err != nil {
		return ast.ForeignKeyConstraint{}, err
	}

	columns, err := p.parseIdentifierList()
	if err != nil {
		return ast.ForeignKeyConstraint{}, err
	}

	referencesKeyword, err := p.consume(token.References)
	if err != nil {
		return ast.ForeignKeyConstraint{}, err
	}

	refTableName, err := p.parseObjectName()
	if err != nil {
		return ast.ForeignKeyConstraint{}, err
	}

	var refColumns *ast.IdentifierList
	if p.tok.Kind == token.LeftParentheses {
		list, err := p.parseIdentifierList()
		if err != nil {
			return ast.ForeignKeyConstraint{}, err
		}

		refColumns = &list
	}

	return ast.ForeignKeyConstraint{
		Keywords: ast.ForeignKeyConstraintKeywords{
			Foreign:    foreignKeyword,
			Key:        keyKeyword,
			References: referencesKeyword,
		},
		Columns:      columns,
		RefTableName: refTableName,
		RefColumns:   refColumns,
	}, nil
}

func (p *Parser) parseConstraintNameClause() (*ast.ConstraintNameClause, error) {
	constraintKeyword := p.tok
	p.advance() // consume CONSTRAINT

	name, err := p.consumeIdentifier()
	if err != nil {
		return nil, err
	}

	return &ast.ConstraintNameClause{
		Keywords: ast.ConstraintNameClauseKeywords{
			Constraint: constraintKeyword,
		},
		Name: name,
	}, nil
}
