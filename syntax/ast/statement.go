package ast

import "github.com/mebyus/sqlfmt/syntax/token"

type Statement any

type CreateTableStatement struct {
	IsTemporary bool

	Name        TableName
	Columns     []ColumnSpecifier
	Constraints []ConstraintSpecifier

	// Token with a tablespace name if any
	Tablespace *token.Token
}

type ColumnCommentStatement struct {
	TableName TableName
	Name      token.Token
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

type DefaultClause struct {
	Expression Expression
}

type TypeSpecifier struct {
	Spec []token.Token
}

type Expression struct {
}

type LineComment struct {
	Content token.Token
}
