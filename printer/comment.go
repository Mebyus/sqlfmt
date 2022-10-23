package printer

import (
	"strings"

	"github.com/mebyus/sqlfmt/syntax/ast"
)

func (p *Printer) writeLineComment(comment ast.LineComment) {
	p.write("-- ")
	p.write(strings.TrimSpace(comment.Content.Lit))
	p.nextLine()
}