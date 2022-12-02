package ast

import (
	"github.com/mebyus/sqlfmt/syntax/ast/statement"
	"github.com/mebyus/sqlfmt/syntax/token"
)

type Statement any

type CreateTableStatement struct {
	Keywords         CreateTableKeywords
	Temporary        *token.Token
	LeftParentheses  token.Token
	RightParentheses token.Token
	Semicolon        token.Token

	Name        TableName
	Columns     []ColumnSpecifier
	Constraints []ConstraintSpecifier

	// Token with a tablespace name if any
	Tablespace *token.Token
}

type CreateTableKeywords struct {
	Create token.Token
	Table  token.Token
}

type CreateIndexStatement struct {
	Keywords   CreateIndexKeywords
	Tablespace *TablespaceClause
	Where      *WhereClause
	Semicolon  token.Token

	Name    TableName
	Columns []IndexColumn
}

type IndexColumn struct {
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
	TableName     TableName
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
	TableName     TableName
	CommentString token.Token
}

type SetTableCommentKeywords struct {
	Comment token.Token
	On      token.Token
	Table   token.Token
	Is      token.Token
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
	// token.Kind is Dot
	Dot token.Token

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
	Tokens []token.Token
}

// <Comment> = <LineComment> | <MultiLineComment>
type Comment struct {
	// Token.Kind is an LineComment or MultiLineComment
	Content token.Token

	// Comment is inlined if it appears on the line which contains non-comment tokens
	Inlined bool
}
