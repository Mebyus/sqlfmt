package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeUpdateStatement(stmt ast.UpdateStatement) {
	p.writeToken(stmt.Keywords.Update)
	if stmt.Keywords.Only != nil {
		p.writeToken(*stmt.Keywords.Only)
	}
	p.writeObjectName(stmt.TableName)
	p.writeToken(stmt.Keywords.Set)
	p.nl()
	p.wse.Inc()
	for _, elem := range stmt.Elements {
		p.writeUpdateListElement(elem)
		p.nl()
	}
	p.wse.Dec()
	p.nl()
	if stmt.Where != nil {
		p.writeWhereClause(*stmt.Where)
	}
	p.writeToken(stmt.Semicolon)
	p.nl()
}

func (p *Printer) writeUpdateListElement(element ast.UpdateListElement) {
	if element.Comma != nil {
		p.writeToken(*element.Comma)
	}

	switch a := element.Assignment.(type) {
	case ast.SingleColumnAssignment:
		p.writeSingleColumnAssignment(a)
	default:
		panic("unreachable: unexpected update assignment type")
	}
}

func (p *Printer) writeSingleColumnAssignment(assignment ast.SingleColumnAssignment) {
	p.writeIdentifier(assignment.Name)
	p.writeToken(assignment.Equal)
	p.writeDefaultableExpression(assignment.Value)
}

func (p *Printer) writeDefaultableExpression(exp ast.DefaultableExpression) {
	switch {
	case exp.Default != nil:
		p.writeToken(*exp.Default)
	case exp.Expression != nil:
		p.writeExpression(*exp.Expression)
	default:
		panic("unreachable: empty defaultable expression")
	}
}
