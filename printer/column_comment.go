package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeColumnCommentStatement(stmt ast.ColumnCommentStatement) {
	p.write("COMMENT ON COLUMN ")
	p.writeTableName(stmt.TableName)
	p.write(".")
	p.writeToken(stmt.Name)
	p.write(" IS ")
	p.writeToken(stmt.Comment)
	p.write(";")
	p.nextLine()
}
