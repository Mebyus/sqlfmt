package scanner

import "github.com/mebyus/sqlfmt/syntax/token"

type Stream interface {
	Scan() token.Token
}

// Eraser is a Stream decorator. It erases all position information
// from tokens
type Eraser struct {
	stream Stream
}

func NewEraser(stream Stream) *Eraser {
	return &Eraser{
		stream: stream,
	}
}

func (e *Eraser) Scan() token.Token {
	tok := e.stream.Scan()
	tok.Erase()
	return tok
}
