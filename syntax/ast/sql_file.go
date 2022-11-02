package ast

type SQLFile struct {
	Statements []Statement
	Comments   []Comment
}
