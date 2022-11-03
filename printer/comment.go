package printer

import (
	"strings"

	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Printer) writeLineComment(comment ast.Comment) {
	p.write("-- ")
	p.write(strings.TrimSpace(comment.Content.Lit[2:]))
	p.nl()
}

func (p *Printer) writeMultiLineComment(comment ast.Comment) {
	p.write(comment.Content.Lit)
}

func (p *Printer) writeComment(comment ast.Comment) {
	kind := comment.Content.Kind
	switch kind {
	case token.LineComment:
		p.writeLineComment(comment)
	case token.MultiLineComment:
		p.writeMultiLineComment(comment)
	default:
		panic("unreachable: comment token expected, got " + kind.String())
	}
}
