package ast

import (
	"github.com/mebyus/sqlfmt/syntax/ast/operator"
	"github.com/mebyus/sqlfmt/syntax/token"
)

// <Expression> = <PrimaryOperand> | <BinaryExpression>
type Expression any

// <PrimaryOperand> = <Operand> | <UnaryExpression>
type PrimaryOperand any

// <Operand> = <Literal> | <Identifier> | <ParenthesizedExpression> | <FunctionCallExpression>
type Operand any

type Literal struct {
	// token.Kind is String, DecimalInteger, DecimalFloat, True, False or Null
	Token token.Token
}

// <ParenthesizedExpression> = "(" <Expression> ")"
type ParenthesizedExpression struct {
	LeftParentheses  token.Token
	RightParentheses token.Token
	Expression       Expression
}

// <FunctionCallExpression> = <ObjectName> "(" { <Expression> ["," <Expression>] } ")"
type FunctionCallExpression struct {
	Name             ObjectName
	LeftParentheses  token.Token
	RightParentheses token.Token
	Arguments        []FunctionCallArgument
}

type FunctionCallArgument struct {
	Comma      *token.Token
	Expression Expression
}

// <UnaryExpression> = <UnaryOperator> <UnaryOperand>
type UnaryExpression struct {
	Operator     operator.Operator
	UnaryOperand UnaryOperand
}

// <UnaryOperand> = <Operand> | <UnaryExpression>
type UnaryOperand any

// <BinaryExpression> = <Expression> <BinaryOperator> <Expression>
type BinaryExpression struct {
	Operator  operator.Operator
	LeftSide  Expression
	RightSide Expression
}
