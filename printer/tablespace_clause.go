package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeTablespaceClause(clause ast.TablespaceClause) {
	p.writeToken(clause.Keywords.Tablespace)
	p.writeToken(clause.Name.Token)
}
