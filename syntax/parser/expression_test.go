package parser

import (
	"reflect"
	"testing"

	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/ast/operator"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func lit(kind token.Kind, lit string) ast.Literal {
	return ast.Literal{
		Token: token.Token{
			Kind: kind,
			Lit:  lit,
		},
	}
}

func idn(lit string) ast.Identifier {
	return ast.Identifier{
		Token: token.Token{
			Kind: token.Identifier,
			Lit:  lit,
		},
	}
}

func qid(lit string) ast.Identifier {
	return ast.Identifier{
		Token: token.Token{
			Kind: token.QuotedIdentifier,
			Lit:  lit,
		},
	}
}

func par(exp ast.Expression) ast.ParenthesizedExpression {
	return ast.ParenthesizedExpression{
		LeftParentheses: token.Token{
			Kind: token.LeftParentheses,
		},
		RightParentheses: token.Token{
			Kind: token.RightParentheses,
		},
		Expression: exp,
	}
}

func uex(kind token.Kind, op ast.Operand) ast.UnaryExpression {
	return ast.UnaryExpression{
		Operator:     operator.NewUnary(token.Token{Kind: kind}),
		UnaryOperand: op,
	}
}

func bin(kind token.Kind, left ast.Expression, right ast.Expression) ast.BinaryExpression {
	return ast.BinaryExpression{
		Operator:  operator.NewBinary(token.Token{Kind: kind}),
		LeftSide:  left,
		RightSide: right,
	}
}

func TestParseExpression(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    ast.Expression
		wantErr bool
	}{
		{
			name:    "1 empty string",
			str:     "",
			wantErr: true,
		},
		{
			name: "2 decimal integer literal",
			str:  "42",
			want: lit(token.DecimalInteger, "42"),
		},
		{
			name: "3 decimal float literal",
			str:  "42.042",
			want: lit(token.DecimalFloat, "42.042"),
		},
		{
			name: "4 identifier",
			str:  "abc",
			want: idn("abc"),
		},
		{
			name: "5 quoted identifier",
			str:  `"abc"`,
			want: qid(`"abc"`),
		},
		{
			name: "6 integer in parentheses",
			str:  "(3)",
			want: par(lit(token.DecimalInteger, "3")),
		},
		{
			name: "7 unary expression on integer",
			str:  "+49",
			want: uex(token.Plus, lit(token.DecimalInteger, "49")),
		},
		{
			name: "8 unary expression on identifier",
			str:  "not is_empty",
			want: uex(token.Not, idn("is_empty")),
		},
		{
			name: "9 binary expression on integers",
			str:  "49 - 90",
			want: bin(token.Minus, lit(token.DecimalInteger, "49"), lit(token.DecimalInteger, "90")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseExpression(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
