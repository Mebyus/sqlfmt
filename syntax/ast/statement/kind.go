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
)

var Name = [...]string{
	Unknown:          "unknown",
	Create:           "create",
	CreateTable:      "create table",
	CreateIndex:      "create index",
	SetComment:       "set comment",
	SetColumnComment: "set column comment",
	SetTableComment:  "set table comment",
}

func (kind Kind) String() string {
	return Name[kind] + " statement"
}
