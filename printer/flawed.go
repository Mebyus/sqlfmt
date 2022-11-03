package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeFlawedStatement(stmt ast.FlawedStatement) {
	p.writeJoinedTokens(stmt.Tokens)
	p.nl()
}
