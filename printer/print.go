package printer

import (
	"fmt"

	"github.com/mebyus/sqlfmt/syntax/ast"
)

func (p *Printer) writeStatement(stmt ast.Statement) {
	switch s := stmt.(type) {
	case ast.CreateTableStatement:
		p.writeCreateTableStatement(s)
	case ast.Comment:
		p.writeLineComment(s)
	case ast.ColumnCommentStatement:
		p.writeColumnCommentStatement(s)
	case ast.TableCommentStatement:
		p.writeTableCommentStatement(s)
	case ast.UnknownStatement:
		p.writeUnknownStatement(s)
	case ast.FlawedStatement:
		p.writeFlawedStatement(s)
	default:
		panic(fmt.Sprintf("unknown statement type %T %v", stmt, stmt))
	}
}
