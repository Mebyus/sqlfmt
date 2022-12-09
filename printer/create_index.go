package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeCreateIndexStatement(stmt ast.CreateIndexStatement) {
	p.writeToken(stmt.Keywords.Create)
	p.writeToken(stmt.Keywords.Index)
	if stmt.Name != nil {
		p.writeIdentifier(*stmt.Name)
	}
	p.writeToken(stmt.Keywords.On)
	p.writeObjectName(stmt.TableName)
	if stmt.Using != nil {
		p.writeToken(stmt.Using.Keywords.Using)
		p.writeToken(stmt.Using.MethodName)
	}
	p.writeToken(stmt.LeftParentheses)
	for _, column := range stmt.Columns {
		p.writeIndexColumn(column)
	}
	p.writeToken(stmt.RightParentheses)
	if stmt.Where != nil {
		p.writeWhereClause(*stmt.Where)
	}
	p.writeToken(stmt.Semicolon)
	p.nl()
}

func (p *Printer) writeWhereClause(clause ast.WhereClause) {
	p.writeToken(clause.Keywords.Where)
	p.writeExpression(clause.Expression)
}

func (p *Printer) writeIndexColumn(column ast.IndexColumn) {
	if column.Comma != nil {
		p.writeToken(*column.Comma)
	}
	p.writeIdentifier(column.Name)
}
