package token

type Kind int

const (
	EOF Kind = iota

	Dot              // .
	Comma            // ,
	Semicolon        // ;
	Asterisk         // *
	Slash            // /
	Minus            // -
	Plus             // +
	Percent          // %
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

	endKeyword
	noStaticLiteral

	LineComment      // -- it's a line comment
	MultiLineComment // /* it's a multi-line comment */
	String           // 'string literal'
	DecimalInteger   // 54163
	DecimalFloat     // 10.432
	Identifier       // my_table
	QuotedIdentifier // "my_table"
	Illegal
)

func (kind Kind) IsKeyword() bool {
	return beginKeyword < kind && kind < endKeyword
}

func (kind Kind) HasStaticLiteral() bool {
	return kind < noStaticLiteral
}
