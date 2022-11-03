package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeSetColumnCommentStatement(stmt ast.SetColumnCommentStatement) {
	p.writeToken(stmt.Keywords.Comment)
	p.space()
	p.writeToken(stmt.Keywords.On)
	p.space()
	p.writeToken(stmt.Keywords.Column)
	p.space()
	p.writeTableName(stmt.TableName)
	p.writeToken(stmt.Dot)
	p.writeToken(stmt.ColumnName)
	p.space()
	p.writeToken(stmt.Keywords.Is)
	p.space()
	p.writeToken(stmt.CommentString)
	p.writeToken(stmt.Semicolon)
	p.nl()
}
