package ast

import "github.com/mebyus/sqlfmt/syntax/token"

type UpdateStatement struct {
	Keywords  UpdateStatementKeywords
	TableName ObjectName
	Elements  []UpdateListElement
	Where     *WhereClause
	Semicolon token.Token
}

type UpdateStatementKeywords struct {
	Update token.Token
	Set    token.Token
	Only   *token.Token
}

type UpdateListElement struct {
	Comma      *token.Token
	Assignment UpdateAssignment
}

// <UpdateAssignment> = <SingleColumnAssignment> |
type UpdateAssignment any

type SingleColumnAssignment struct {
	Name  Identifier
	Equal token.Token
	Value DefaultableExpression
}

type DefaultableExpression struct {
	Default    *token.Token
	Expression *Expression
}
