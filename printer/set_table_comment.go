package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeSetTableCommentStatement(stmt ast.SetTableCommentStatement) {
	p.writeToken(stmt.Keywords.Comment)
	p.space()
	p.writeToken(stmt.Keywords.On)
	p.space()
	p.writeToken(stmt.Keywords.Table)
	p.space()
	p.writeTableName(stmt.TableName)
	p.space()
	p.writeToken(stmt.Keywords.Is)
	p.space()
	p.writeToken(stmt.CommentString)
	p.writeToken(stmt.Semicolon)
	p.nl()
}
