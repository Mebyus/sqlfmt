package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeExpression(exp ast.Expression) {
	p.writeJoinedTokens(exp.Tokens)
}
