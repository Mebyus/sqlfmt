package statement

type Kind int

const (
	Unknown Kind = iota
	Create
	CreateTable
	Comment
	ColumnComment
	TableComment
	LineComment
)
