package printer

import (
	"github.com/mebyus/sqlfmt/syntax/ast"
)

func (p *Printer) writeCreateTableStatement(stmt ast.CreateTableStatement) {
	p.writeToken(stmt.Keywords.Create)
	p.space()
	if stmt.Temporary != nil {
		p.writeToken(*stmt.Temporary)
		p.space()
	}
	p.writeToken(stmt.Keywords.Table)
	p.space()
	p.writeTableName(stmt.Name)
	p.space()
	p.writeToken(stmt.LeftParentheses)

	p.indentation.Inc()
	p.nl()
	if len(stmt.Columns) > 0 {
		for i := 0; i < len(stmt.Columns)-1; i++ {
			p.writeColumnSpecifier(stmt.Columns[i])
			p.write(",")
			p.nl()
		}
		last := stmt.Columns[len(stmt.Columns)-1]
		p.writeColumnSpecifier(last)
	}

	if len(stmt.Constraints) > 0 {
		for i := 0; i < len(stmt.Constraints)-1; i++ {
			p.writeConstraintSpecifier(stmt.Constraints[i])
			p.write(",")
			p.nl()
		}
		last := stmt.Constraints[len(stmt.Constraints)-1]
		p.writeConstraintSpecifier(last)
	}
	p.indentation.Dec()
	p.nl()

	p.writeToken(stmt.RightParentheses)
	if stmt.Tablespace != nil {
		p.space()
		p.write("TABLESPACE")
		p.space()
		p.writeToken(*stmt.Tablespace)
	}
	p.writeToken(stmt.Semicolon)
	p.nl()
}

func (p *Printer) writeColumnSpecifier(spec ast.ColumnSpecifier) {
	p.writeToken(spec.Name)
	p.space()
	p.writeJoinedTokens(spec.Type.Spec)
	if spec.IsPrimaryKey {
		p.space()
		p.write("PRIMARY")
		p.space()
		p.write("KEY")
	}
	if spec.IsNotNull {
		p.space()
		p.write("NOT")
		p.space()
		p.write("NULL")
	}
	if spec.Default != nil {
		p.space()
		p.write("DEFAULT")
	}
}

func (p *Printer) writeConstraintSpecifier(spec ast.ConstraintSpecifier) {
	p.write("CONSTRAINT")
	p.space()
	p.writeToken(spec.Name)
}
