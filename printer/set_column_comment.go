package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeSetColumnCommentStatement(stmt ast.SetColumnCommentStatement) {
	p.writeToken(stmt.Keywords.Comment)
	p.writeToken(stmt.Keywords.On)
	p.writeToken(stmt.Keywords.Column)
	p.writeObjectName(stmt.TableName)
	p.writeToken(stmt.Dot)
	p.writeToken(stmt.ColumnName)
	p.writeToken(stmt.Keywords.Is)
	p.writeToken(stmt.CommentString)
	p.writeToken(stmt.Semicolon)
	p.nl()
}
