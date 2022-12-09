package statement

type Kind int

const (
	Unknown Kind = iota
	Create
	CreateTable
	CreateIndex
	SetComment
	SetColumnComment
	SetTableComment
	Update
)

type Category int

const (
	// Data Definition Language
	DDL Category = 1 + iota

	// Data Manipulation Language
	DML

	// Data Control Language
	DCL

	// Transaction Control Language
	TCL
)

var Name = [...]string{
	Unknown:          "unknown",
	Create:           "create",
	CreateTable:      "create table",
	CreateIndex:      "create index",
	SetComment:       "set comment",
	SetColumnComment: "set column comment",
	SetTableComment:  "set table comment",
	Update:           "update",
}

func (k Kind) String() string {
	return Name[k] + " statement"
}

func (k Kind) Category() Category {
	switch k {
	case Create, CreateTable, CreateIndex, SetComment, SetColumnComment, SetTableComment:
		return DDL
	case Update:
		return DML
	default:
		panic("unreachable: unexpected statement kind")
	}
}
