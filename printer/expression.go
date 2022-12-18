package printer

import (
	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/ast/operator"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Printer) writeExpression(expression ast.Expression) {
	switch exp := expression.(type) {
	case ast.UnaryExpression:
		p.writeUnaryExpression(exp)
	case ast.BinaryExpression:
		p.writeBinaryExpression(exp)
	case ast.Operand:
		p.writeOperand(exp)
	default:
		panic("unreachable: unexpected expression type")
	}
}

func (p *Printer) writeLiteral(lit ast.Literal) {
	p.writeToken(lit.Token)
}

func (p *Printer) writeParenthesizedExpression(exp ast.ParenthesizedExpression) {
	p.writeToken(exp.LeftParentheses)
	p.writeExpression(exp.Expression)
	p.writeToken(exp.LeftParentheses)
}

func (p *Printer) writeOperand(operand ast.Operand) {
	switch o := operand.(type) {
	case ast.Literal:
		p.writeLiteral(o)
	case ast.Identifier:
		p.writeIdentifier(o)
	case ast.ParenthesizedExpression:
		p.writeParenthesizedExpression(o)
	case ast.FunctionCallExpression:
		p.writeFunctionCallExpression(o)
	default:
		panic("unreachable: unexpected operand type")
	}
}

func (p *Printer) writeFunctionCallExpression(exp ast.FunctionCallExpression) {
	p.writeObjectName(exp.Name)
	p.writeToken(exp.LeftParentheses)
	for _, arg := range exp.Arguments {
		p.writeFunctionCallArgument(arg)
	}
	p.writeToken(exp.LeftParentheses)
}

func (p *Printer) writeFunctionCallArgument(arg ast.FunctionCallArgument) {
	if arg.Comma != nil {
		p.writeToken(*arg.Comma)
	}
	p.writeExpression(arg.Expression)
}

func (p *Printer) writeUnaryExpression(exp ast.UnaryExpression) {
	p.writeOperator(exp.Operator)
	op := exp.Operator.(*operator.Unary)
	if op.Token.Kind != token.Not {
		p.wse.None()
	}
	p.writeUnaryOperand(exp.UnaryOperand)
}

func (p *Printer) writeUnaryOperand(uop ast.UnaryOperand) {
	switch o := uop.(type) {
	case ast.UnaryExpression:
		p.writeUnaryExpression(o)
	case *ast.UnaryExpression:
		p.writeUnaryExpression(*o)
	case ast.Operand:
		p.writeOperand(o)
	default:
		panic("unreachable: unexpected unary operand type")
	}
}

func (p *Printer) writeBinaryExpression(exp ast.BinaryExpression) {
	p.writeExpression(exp.LeftSide)
	p.writeOperator(exp.Operator)
	p.writeExpression(exp.RightSide)
}

func (p *Printer) writeOperator(op operator.Operator) {
	switch o := op.(type) {
	case *operator.Unary:
		p.writeToken(o.Token)
	case *operator.Binary:
		p.writeToken(o.Token)
	default:
		panic("unreachable: unexpected operator type")
	}
}
