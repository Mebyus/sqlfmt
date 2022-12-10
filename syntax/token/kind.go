package token

type Kind int

const (
	empty Kind = iota

	EOF

	Dot              // .
	Comma            // ,
	Semicolon        // ;
	Asterisk         // *
	Slash            // /
	Minus            // -
	Plus             // +
	Percent          // %
	Caret            // ^
	Equal            // =
	Greater          // >
	Less             // <
	GreaterEqual     // >=
	LessEqual        // <=
	NotEqual         // !=
	NotEqualAlt      // <>
	DoubleColon      // ::
	LeftParentheses  // (
	RightParentheses // )
	LeftSquare       // [
	RightSquare      // ]

	beginKeyword

	Select
	Insert
	Update
	Delete
	Create
	Truncate
	Alter
	Set
	Default
	Limit
	Null
	Join
	Left
	Right
	On
	Full
	Outer
	Inner
	Into
	Values
	Is
	In
	As
	To
	Union
	All
	With
	Exists
	Any
	Case
	When
	Then
	End
	Having
	Group
	By
	Order
	Index
	Drop
	Table
	Distinct
	From
	Where
	Like
	Ilike
	Similar
	Between
	And
	Or
	Not
	Asc
	Primary
	Key
	Foreign
	Unique
	Add
	Constraint
	Check
	Desc
	Column
	Comment
	Trigger
	Cascade
	Nothing
	Returning
	Nulls
	First
	Last
	Conflict
	Do
	References
	Sequence
	Temporary
	Temp
	Using
	Tablespace
	Only

	endKeyword
	noStaticLiteral

	LineComment      // -- it's a line comment
	MultiLineComment // /* it's a multi-line comment */
	String           // 'string literal'
	DecimalInteger   // 54163
	DecimalFloat     // 10.432
	Identifier       // my_table
	QuotedIdentifier // "my_table"
	Illegal          // 4grevs - would-be-identifier but starts with a digit
)

func (kind Kind) String() string {
	return Literal[kind]
}

func (kind Kind) IsEmpty() bool {
	return kind == empty
}

func (kind Kind) IsKeyword() bool {
	return beginKeyword < kind && kind < endKeyword
}

func (kind Kind) IsComment() bool {
	return kind == LineComment || kind == MultiLineComment
}

func (kind Kind) HasStaticLiteral() bool {
	return kind < noStaticLiteral
}
