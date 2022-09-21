package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeTableCommentStatement(stmt ast.TableCommentStatement) {
	p.write("COMMENT ON TABLE ")
	p.writeTableName(stmt.TableName)
	p.write(" IS ")
	p.writeToken(stmt.Comment)
	p.write(";")
	p.nextLine()
}
