package printer

import (
	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/token"
)

func (p *Printer) writeCreateTableStatement(stmt ast.CreateTableStatement) {
	p.write(token.Literal[token.Create])
	p.write(" ")
	if stmt.IsTemporary {
		p.write("TEMPORARY ")
	}
	p.write(token.Literal[token.Table])
	p.write(" ")
	p.writeTableName(stmt.Name)
	p.write(" (")

	p.indentation.Inc()
	p.nextLine()
	if len(stmt.Columns) > 0 {
		for i := 0; i < len(stmt.Columns)-1; i++ {
			p.writeColumnSpecifier(stmt.Columns[i])
			p.write(",")
			p.nextLine()
		}
		last := stmt.Columns[len(stmt.Columns)-1]
		p.writeColumnSpecifier(last)
	}

	if len(stmt.Constraints) > 0 {
		for i := 0; i < len(stmt.Constraints)-1; i++ {
			p.writeConstraintSpecifier(stmt.Constraints[i])
			p.write(",")
			p.nextLine()
		}
		last := stmt.Constraints[len(stmt.Constraints)-1]
		p.writeConstraintSpecifier(last)
	}
	p.indentation.Dec()
	p.nextLine()

	p.write(")")
	if stmt.Tablespace != nil {
		p.write(" TABLESPACE ")
		p.writeToken(*stmt.Tablespace)
	}
	p.write(";")
	p.nextLine()
}

func (p *Printer) writeColumnSpecifier(spec ast.ColumnSpecifier) {
	p.write(spec.Name.Lit)
	p.write(" ")
	if len(spec.Type.Spec) > 0 {
		for i := 0; i < len(spec.Type.Spec)-1; i++ {
			p.writeToken(spec.Type.Spec[i])
			p.write(" ")
		}
		last := spec.Type.Spec[len(spec.Type.Spec)-1]
		p.writeToken(last)
	}
	if spec.IsPrimaryKey {
		p.write(" PRIMARY KEY")
	}
	if spec.IsNotNull {
		p.write(" NOT NULL")
	}
	if spec.Default != nil {
		p.write(" DEFAULT")
	}
}

func (p *Printer) writeConstraintSpecifier(spec ast.ConstraintSpecifier) {
	p.write("CONSTRAINT ")
	p.writeToken(spec.Name)
}
