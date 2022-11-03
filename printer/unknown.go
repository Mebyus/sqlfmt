package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeUnknownStatement(stmt ast.UnknownStatement) {
	p.writeJoinedTokens(stmt.Tokens)
	p.nl()
}
