package token

import "fmt"

type Token struct {
	// Not empty only for tokens which can have
	// arbitrary literal
	//
	// Examples: identifiers, numbers, illegal tokens
	Lit string

	Pos Pos

	// Number of tokens emitted by a stream before this one
	Index int

	Kind Kind
}

func (tok *Token) String() string {
	if tok.Kind.HasStaticLiteral() {
		return fmt.Sprintf("%-12s%s", tok.Pos.String(), Literal[tok.Kind])
	}
	return fmt.Sprintf("%-12s%-12s%s", tok.Pos.String(), Literal[tok.Kind], tok.Lit)
}
