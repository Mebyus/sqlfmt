package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeStatement(stmt ast.Statement) {
	switch s := stmt.(type) {
	case ast.CreateTableStatement:
		p.writeCreateTableStatement(s)
	case ast.LineComment:
		p.writeLineComment(s)
	case ast.ColumnCommentStatement:
		p.writeColumnCommentStatement(s)
	}
}
