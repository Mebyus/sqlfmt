package ast

import (
	"github.com/mebyus/sqlfmt/syntax/ast/statement"
	"github.com/mebyus/sqlfmt/syntax/token"
)

type Statement any

type CreateTableStatement struct {
	Keywords         CreateTableKeywords
	LeftParentheses  token.Token
	RightParentheses token.Token
	Semicolon        token.Token

	Name       ObjectName
	Properties []TablePropertySpecifier

	Tablespace *TablespaceClause
}

type CreateTableKeywords struct {
	Create    token.Token
	Table     token.Token
	Temporary *token.Token
	Temp      *token.Token
}

type CreateIndexStatement struct {
	Keywords         CreateIndexKeywords
	Using            *IndexUsingClause
	Tablespace       *TablespaceClause
	Where            *WhereClause
	LeftParentheses  token.Token
	RightParentheses token.Token
	Semicolon        token.Token

	Name      *Identifier
	TableName ObjectName
	Columns   []IndexColumn
}

type IndexColumn struct {
	Name  Identifier
	Comma *token.Token
}

type IndexUsingClause struct {
	Keywords IndexUsingClauseKeywords

	// token.Kind is Identifier
	MethodName token.Token
}

type IndexUsingClauseKeywords struct {
	Using token.Token
}

type CreateIndexKeywords struct {
	Create       token.Token
	Index        token.Token
	On           token.Token
	Unique       *token.Token
	Concurrently *token.Token
}

type TablespaceClause struct {
	Keywords TablespaceClauseKeywords
	Name     Identifier
}

type TablespaceClauseKeywords struct {
	Tablespace token.Token
}

type WhereClause struct {
	Keywords   WhereClauseKeywords
	Expression Expression
}

type WhereClauseKeywords struct {
	Where token.Token
}

// <SetColumnCommentStatement> = "COMMENT" "ON" "COLUMN" <TableName> "." <Identifier>
// "IS" <String> ";"
type SetColumnCommentStatement struct {
	Keywords      SetColumnCommentKeywords
	Dot           token.Token
	Semicolon     token.Token
	TableName     ObjectName
	ColumnName    token.Token
	CommentString token.Token
}

type SetColumnCommentKeywords struct {
	Comment token.Token
	On      token.Token
	Column  token.Token
	Is      token.Token
}

// <SetTableCommentStatement> = "COMMENT" "ON" "TABLE" <TableName> "IS" <String> ";"
type SetTableCommentStatement struct {
	Keywords      SetTableCommentKeywords
	Semicolon     token.Token
	TableName     ObjectName
	CommentString token.Token
}

type SetTableCommentKeywords struct {
	Comment token.Token
	On      token.Token
	Table   token.Token
	Is      token.Token
}

// <ObjectName> = <Identifier> | <QualifiedIdentifier>
type ObjectName any

// <Identifier> = <RawIdentifier> | <QuotedIdentifier>
type Identifier struct {
	// Token.Kind is an Identifier or QuotedIdentifier
	Token token.Token
}

// <QualifiedIdentifier> = <Identifier> "." <Identifier>
type QualifiedIdentifier struct {
	// token.Kind is Dot
	Dot token.Token

	SchemaName   Identifier
	RawTableName Identifier
}

type TablePropertySpecifier struct {
	Comma    *token.Token
	Property TableProperty
}

// <TableProperty> = <ColumnSpecifier> | <ConstraintSpecifier>
type TableProperty any

type ColumnSpecifier struct {
	Name        Identifier
	Type        TypeSpecifier
	Constraints ColumnConstraints
}

type ColumnConstraints struct {
	Null       *token.Token
	NotNull    *NotNullClause
	PrimaryKey *PrimaryKeyClause
}

type NotNullClause struct {
	Not  token.Token
	Null token.Token
}

type PrimaryKeyClause struct {
	Primary token.Token
	Key     token.Token
}

type ConstraintSpecifier struct {
	Name       *ConstraintNameClause
	Constraint TableConstraint
}

type ConstraintNameClause struct {
	Keywords ConstraintNameClauseKeywords
	Name     Identifier
}

type ConstraintNameClauseKeywords struct {
	Constraint token.Token
}

// <TableConstraint> = <ForeignKeyConstraint> |
type TableConstraint any

type ForeignKeyConstraint struct {
	Keywords     ForeignKeyConstraintKeywords
	Columns      IdentifierList
	RefTableName ObjectName
	RefColumns   *IdentifierList
}

type IdentifierList struct {
	LeftParentheses  token.Token
	RightParentheses token.Token

	Elements []IdentifierElement
}

type IdentifierElement struct {
	Name  Identifier
	Comma *token.Token
}

type ForeignKeyConstraintKeywords struct {
	Foreign    token.Token
	Key        token.Token
	References token.Token
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
	Tokens []token.Token
}

// <Comment> = <LineComment> | <MultiLineComment>
type Comment struct {
	// Token.Kind is an LineComment or MultiLineComment
	Content token.Token

	// Comment is inlined if it appears on the line which contains non-comment tokens
	Inlined bool
}
