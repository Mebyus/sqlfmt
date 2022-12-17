package operator

import "github.com/mebyus/sqlfmt/syntax/token"

type Operator interface {
	Precedence() int
}

type Unary struct {
	// token.Kind is Plus, Minus or Not
	Token token.Token
}

func NewUnary(tok token.Token) *Unary {
	return &Unary{Token: tok}
}

func (u *Unary) Precedence() int {
	switch u.Token.Kind {
	case token.Plus, token.Minus:
		return 4
	case token.Not:
		return 12
	default:
		panic("unreachable: unexpected unary operator")
	}
}

type Binary struct {
	// token.Kind is Plus, Minus, Caret, And, Or, Is, ...
	Token token.Token
}

func NewBinary(tok token.Token) *Binary {
	return &Binary{Token: tok}
}

func (b *Binary) Precedence() int {
	switch b.Token.Kind {
	case token.DoubleColon:
		return 2
	case token.Caret:
		return 5
	case token.Asterisk, token.Slash, token.Percent:
		return 6
	case token.Plus, token.Minus:
		return 7
	case token.Between, token.In, token.Like, token.Ilike:
		return 9
	case token.Less, token.Greater, token.Equal, token.LessEqual, token.GreaterEqual, token.NotEqual, token.NotEqualAlt:
		return 10
	case token.Is:
		return 11
	case token.And:
		return 13
	case token.Or:
		return 14
	default:
		panic("unreachable: unexpected binary operator")
	}
}

type IsNot struct {
	Is  token.Token
	Not token.Token
}

func (i *IsNot) Precedence() int {
	return 11
}

type NotIn struct {
	Not token.Token
	In  token.Token
}

func (n *NotIn) Precedence() int {
	return 9
}

type SimilarTo struct {
	Similar token.Token
	To      token.Token
}

func (s *SimilarTo) Precedence() int {
	return 9
}

type NotSimilarTo struct {
	Not     token.Token
	Similar token.Token
	To      token.Token
}

func (s *NotSimilarTo) Precedence() int {
	return 9
}

type NotLike struct {
	Not  token.Token
	Like token.Token
}

func (n *NotLike) Precedence() int {
	return 9
}

type NotIlike struct {
	Not  token.Token
	Like token.Token
}

func (n *NotIlike) Precedence() int {
	return 9
}
