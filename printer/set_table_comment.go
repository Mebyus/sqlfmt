package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeSetTableCommentStatement(stmt ast.SetTableCommentStatement) {
	p.writeToken(stmt.Keywords.Comment)
	p.writeToken(stmt.Keywords.On)
	p.writeToken(stmt.Keywords.Table)
	p.writeObjectName(stmt.TableName)
	p.writeToken(stmt.Keywords.Is)
	p.writeToken(stmt.CommentString)
	p.writeToken(stmt.Semicolon)
	p.nl()
}
