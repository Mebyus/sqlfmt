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
	Some
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
	True
	False

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

func (k Kind) String() string {
	return Literal[k]
}

func (k Kind) IsEmpty() bool {
	return k == empty
}

func (k Kind) IsKeyword() bool {
	return beginKeyword < k && k < endKeyword
}

func (k Kind) IsComment() bool {
	return k == LineComment || k == MultiLineComment
}

func (k Kind) HasStaticLiteral() bool {
	return k < noStaticLiteral
}

func (k Kind) IsIdentifier() bool {
	return k == Identifier || k == QuotedIdentifier
}

func (k Kind) IsLiteral() bool {
	switch k {
	case String, DecimalInteger, DecimalFloat, True, False, Null:
		return true
	default:
		return false
	}
}

func (k Kind) IsUnaryOperator() bool {
	switch k {
	case Plus, Minus, Not:
		return true
	default:
		return false
	}
}

func (k Kind) IsBinaryOperator() bool {
	switch k {
	case
		DoubleColon, Caret, Asterisk, Slash, Percent, Plus, Minus, Between, In, Like, Ilike, Less, Greater,
		Equal, LessEqual, GreaterEqual, NotEqual, NotEqualAlt, Is, And, Or:

		return true
	default:
		return false
	}
}
