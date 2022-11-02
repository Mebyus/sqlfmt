package ast

import (
	"github.com/mebyus/sqlfmt/syntax/ast/statement"
	"github.com/mebyus/sqlfmt/syntax/token"
)

type Statement any

type CreateTableStatement struct {
	IsTemporary bool

	Name        TableName
	Columns     []ColumnSpecifier
	Constraints []ConstraintSpecifier

	// Token with a tablespace name if any
	Tablespace *token.Token
}

// <ColumnCommentStatement> = "COMMENT" "ON" "COLUMN" <TableName> "." <Identifier>
// "IS" <String> ";"
type ColumnCommentStatement struct {
	TableName TableName
	Name      token.Token
	Comment   token.Token
}

// <TableCommentStatement> = "COMMENT" "ON" "TABLE" <TableName> "IS" <String> ";"
type TableCommentStatement struct {
	TableName TableName
	Comment   token.Token
}

// <TableName> = <Identifier> | <QualifiedIdentifier>
type TableName any

// <Identifier> = <RawIdentifier> | <QuotedIdentifier>
type Identifier struct {
	// Token.Kind is an Identifier or QuotedIdentifier
	Token token.Token
}

// <QualifiedIdentifier> = <Identifier> "." <Identifier>
type QualifiedIdentifier struct {
	SchemaName   Identifier
	RawTableName Identifier
}

type ColumnSpecifier struct {
	IsNotNull    bool
	IsPrimaryKey bool
	Name         token.Token
	Type         TypeSpecifier
	Default      *DefaultClause
}

type ConstraintSpecifier struct {
	Name token.Token
}

type Error struct {
	// Index of token which produced an error inside
	// statement's list of tokens
	Index int

	// Token which produced error upon parsing a statement
	Token token.Token

	// Kind of statement which parser was trying to parse and failed
	Statement statement.Kind
}

type FlawedStatement struct {
	Error Error

	// List of tokens which failed to produce a valid statement
	Tokens []token.Token
}

type UnknownStatement struct {
	Tokens []token.Token
}

type DefaultClause struct {
	Expression Expression
}

type TypeSpecifier struct {
	Spec []token.Token
}

type Expression struct {
}

// <Comment> = <LineComment> | <MultiLineComment>
type Comment struct {
	// Token.Kind is an LineComment or MultiLineComment
	Content token.Token
}
