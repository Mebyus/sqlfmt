package parser

import (
	"fmt"

	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/ast/operator"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Parser) parseExpression() (ast.Expression, error) {
	operand, err := p.parsePrimaryOperand()
	if err != nil {
		return nil, err
	}
	if !p.tok.Kind.IsBinaryOperator() {
		return operand, nil
	}

	operands := []ast.PrimaryOperand{operand}
	var operators []*operator.Binary
	for p.tok.Kind.IsBinaryOperator() {
		op := operator.NewBinary(p.tok)
		p.advance()

		operand, err = p.parsePrimaryOperand()
		if err != nil {
			return nil, err
		}

		operands = append(operands, operand)
		operators = append(operators, op)
	}

	// handle common cases with 1, 2 or 3 operators by hand
	switch len(operators) {
	case 0:
		panic("unreachable: slice must contain at least one operator")
	case 1:
		return composeBinaryExpressionWithOneOperator(operators, operands), nil
	case 2:
		return composeBinaryExpressionWithTwoOperators(operators, operands), nil
	case 3:
		return composeBinaryExpressionWithThreeOperators(operators, operands), nil
	default:
		return nil, fmt.Errorf("4 or more operators in binary expression not implemented")
	}
}

func composeBinaryExpressionWithOneOperator(ops []*operator.Binary, operands []ast.PrimaryOperand) ast.BinaryExpression {
	return ast.BinaryExpression{
		Operator:  ops[0],
		LeftSide:  operands[0],
		RightSide: operands[1],
	}
}

func composeBinaryExpressionWithTwoOperators(ops []*operator.Binary, operands []ast.PrimaryOperand) ast.BinaryExpression {
	if ops[0].Precedence() <= ops[1].Precedence() { // a + b + c = ((a + b) + c)
		return ast.BinaryExpression{
			Operator: ops[1],
			LeftSide: ast.BinaryExpression{
				Operator:  ops[0],
				LeftSide:  operands[0],
				RightSide: operands[1],
			},
			RightSide: operands[2],
		}
	}
	return ast.BinaryExpression{ // a + b * c = (a + (b * c))
		Operator: ops[0],
		LeftSide: operands[0],
		RightSide: ast.BinaryExpression{
			Operator:  ops[1],
			LeftSide:  operands[1],
			RightSide: operands[2],
		},
	}
}

func composeBinaryExpressionWithThreeOperators(ops []*operator.Binary, operands []ast.PrimaryOperand) ast.BinaryExpression {
	switch {

	case ops[0].Precedence() <= ops[1].Precedence() && ops[1].Precedence() <= ops[2].Precedence():
		// a + b + c + d = (((a + b) + c) + d)
		return ast.BinaryExpression{
			Operator: ops[2],
			LeftSide: ast.BinaryExpression{
				Operator: ops[1],
				LeftSide: ast.BinaryExpression{
					Operator:  ops[0],
					LeftSide:  operands[0],
					RightSide: operands[1],
				},
				RightSide: operands[2],
			},
			RightSide: operands[3],
		}

	case ops[0].Precedence() <= ops[2].Precedence() && ops[2].Precedence() < ops[1].Precedence():
		// a * b + c * d = ((a * b) + (c * d))
		return ast.BinaryExpression{
			Operator: ops[1],
			LeftSide: ast.BinaryExpression{
				Operator: ops[1],
				LeftSide: ast.BinaryExpression{
					Operator:  ops[0],
					LeftSide:  operands[0],
					RightSide: operands[1],
				},
				RightSide: ast.BinaryExpression{
					Operator:  ops[2],
					LeftSide:  operands[2],
					RightSide: operands[3],
				},
			},
		}

	case ops[1].Precedence() < ops[0].Precedence() && ops[0].Precedence() <= ops[2].Precedence():
		// a + b * c + d =  ((a + (b * c)) + d)
		return ast.BinaryExpression{
			Operator: ops[2],
			LeftSide: ast.BinaryExpression{
				Operator: ops[0],
				LeftSide: operands[0],
				RightSide: ast.BinaryExpression{
					Operator:  ops[1],
					LeftSide:  operands[1],
					RightSide: operands[2],
				},
			},
			RightSide: operands[3],
		}

	case ops[1].Precedence() <= ops[2].Precedence() && ops[2].Precedence() < ops[0].Precedence():
		// a + b * c * d = (a + ((b * c) * d))
		return ast.BinaryExpression{
			Operator: ops[1],
			LeftSide: operands[0],
			RightSide: ast.BinaryExpression{
				Operator: ops[2],
				LeftSide: ast.BinaryExpression{
					Operator:  ops[1],
					LeftSide:  operands[1],
					RightSide: operands[2],
				},
				RightSide: operands[3],
			},
		}

	case ops[2].Precedence() < ops[0].Precedence() && ops[0].Precedence() <= ops[1].Precedence():
		// a + b + c * d = ((a + b) + (c * d))
		return ast.BinaryExpression{
			Operator: ops[1],
			LeftSide: ast.BinaryExpression{
				Operator: ops[1],
				LeftSide: ast.BinaryExpression{
					Operator:  ops[0],
					LeftSide:  operands[0],
					RightSide: operands[1],
				},
				RightSide: ast.BinaryExpression{
					Operator:  ops[2],
					LeftSide:  operands[2],
					RightSide: operands[3],
				},
			},
		}

	case ops[2].Precedence() < ops[1].Precedence() && ops[1].Precedence() < ops[0].Precedence():
		// a < b + c * d = (a < (b + (c * d)))
		return ast.BinaryExpression{
			Operator: ops[0],
			LeftSide: operands[0],
			RightSide: ast.BinaryExpression{
				Operator: ops[1],
				LeftSide: operands[1],
				RightSide: ast.BinaryExpression{
					Operator:  ops[2],
					LeftSide:  operands[2],
					RightSide: operands[3],
				},
			},
		}

	default:
		panic("unreachable: switch must cover all cases")
	}
}

func (p *Parser) parsePrimaryOperand() (ast.PrimaryOperand, error) {
	if p.tok.Kind.IsUnaryOperator() {
		unary, err := p.parseUnaryExpression()
		if err != nil {
			return nil, err
		}
		return unary, nil
	}
	operand, err := p.tryParseOperand()
	if err != nil {
		return nil, err
	}
	if operand != nil {
		return operand, nil
	}
	return nil, fmt.Errorf("expected primary operand, got [ %s ]", p.tok.String())
}

func (p *Parser) parseUnaryExpression() (ast.UnaryExpression, error) {
	topExp := ast.UnaryExpression{
		Operator: operator.NewUnary(p.tok),
	}
	p.advance()

	tipExp := &topExp
	for p.tok.Kind.IsUnaryOperator() {
		nextExp := &ast.UnaryExpression{
			Operator: operator.NewUnary(p.tok),
		}
		p.advance()
		tipExp.UnaryOperand = nextExp
		tipExp = nextExp
	}
	operand, err := p.parseOperand()
	if err != nil {
		return ast.UnaryExpression{}, err
	}
	tipExp.UnaryOperand = operand
	return topExp, nil
}

func (p *Parser) parseOperand() (ast.Operand, error) {
	operand, err := p.tryParseOperand()
	if err != nil {
		return nil, err
	}
	if operand == nil {
		return nil, fmt.Errorf("expected operand, got [ %s ]", p.tok.String())
	}
	return operand, nil
}

func (p *Parser) tryParseOperand() (ast.Operand, error) {
	if p.isLit() {
		lit := ast.Literal{Token: p.tok}
		p.advance()
		return lit, nil
	}

	if p.isIdent() {
		ident := ast.Identifier{Token: p.tok}
		p.advance()
		if p.isLeftPar() {
			return nil, fmt.Errorf("call expression not implemented")
		}
		return ident, nil
	}

	if p.isLeftPar() {
		leftParentheses := p.tok
		p.advance()
		exp, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		rightParentheses, err := p.consume(token.RightParentheses)
		if err != nil {
			return nil, err
		}
		return ast.ParenthesizedExpression{
			LeftParentheses:  leftParentheses,
			RightParentheses: rightParentheses,
			Expression:       exp,
		}, nil
	}

	return nil, nil
}
