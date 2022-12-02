package wsemitter

import "github.com/mebyus/sqlfmt/syntax/token"

type Predictor interface {
	Predict(token.Kind) WhitespaceKind
	Override(WhitespaceKind, token.Kind) WhitespaceKind
}

type Simple struct {
	prev token.Kind
}

var DefaultPredictor = &Simple{}

func (s *Simple) Predict(tok token.Kind) WhitespaceKind {
	if tok == token.LeftParentheses || tok == token.Dot {
		return None
	}
	return Space
}

func (s *Simple) Override(kind WhitespaceKind, tok token.Kind) WhitespaceKind {
	// prev := s.prev
	s.prev = tok
	if kind == Newline || kind == Indentation {
		return kind
	}
	if tok == token.Comma || tok == token.Semicolon || tok == token.Dot || tok == token.RightParentheses {
		return None
	}
	return kind
}
