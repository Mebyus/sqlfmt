package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeUnknownStatement(stmt ast.UnknownStatement) {
	for _, tok := range stmt.Tokens {
		p.writeToken(tok)
		p.write(" ")
	}
	p.nextLine()
}
