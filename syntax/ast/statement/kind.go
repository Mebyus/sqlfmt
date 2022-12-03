package statement

type Kind int

const (
	Unknown Kind = iota
	Create
	CreateTable
	CreateIndex
	Comment
	ColumnComment
	TableComment
	LineComment
)
