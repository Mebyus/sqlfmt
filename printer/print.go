package printer

import (
	"fmt"

	"github.com/mebyus/sqlfmt/syntax/ast"
)

func (p *Printer) writeStatement(stmt ast.Statement) {
	switch s := stmt.(type) {
	case ast.CreateTableStatement:
		p.writeCreateTableStatement(s)
	case ast.CreateIndexStatement:
		p.writeCreateIndexStatement(s)
	case ast.SetColumnCommentStatement:
		p.writeSetColumnCommentStatement(s)
	case ast.SetTableCommentStatement:
		p.writeSetTableCommentStatement(s)
	case ast.UnknownStatement:
		p.writeUnknownStatement(s)
	case ast.FlawedStatement:
		p.writeFlawedStatement(s)
	default:
		panic(fmt.Sprintf("unknown statement type %T %v", stmt, stmt))
	}
}
