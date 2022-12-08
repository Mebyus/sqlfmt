package printer

import (
	"fmt"

	"github.com/mebyus/sqlfmt/syntax/ast"
)

func (p *Printer) writeCreateTableStatement(stmt ast.CreateTableStatement) {
	p.writeToken(stmt.Keywords.Create)
	if stmt.Keywords.Temporary != nil {
		p.writeToken(*stmt.Keywords.Temporary)
	}
	if stmt.Keywords.Temp != nil {
		p.writeToken(*stmt.Keywords.Temp)
	}
	p.writeToken(stmt.Keywords.Table)
	p.writeTableName(stmt.Name)
	p.writeToken(stmt.LeftParentheses)

	p.wse.Inc()
	p.nl()
	for _, property := range stmt.Properties {
		p.writeTablePropertySpecifier(property)
		p.nl()
	}
	p.wse.Dec()
	p.nl()

	p.writeToken(stmt.RightParentheses)
	if stmt.Tablespace != nil {
		p.wse.Inc()
		p.nl()
		p.writeTablespaceClause(*stmt.Tablespace)
		p.wse.Dec()
	}
	p.writeToken(stmt.Semicolon)
	p.nl()
}

func (p *Printer) writeTablePropertySpecifier(spec ast.TablePropertySpecifier) {
	if spec.Comma != nil {
		p.writeToken(*spec.Comma)
	}

	switch property := spec.Property.(type) {
	case ast.ColumnSpecifier:
		p.writeColumnSpecifier(property)
	case ast.ConstraintSpecifier:
		p.writeConstraintSpecifier(property)
	default:
		panic(fmt.Sprintf("unreachable: unknown table property type %T %v", property, property))
	}
}

func (p *Printer) writeColumnSpecifier(spec ast.ColumnSpecifier) {
	p.writeIdentifier(spec.Name)
	p.writeTypeSpecifier(spec.Type)
	p.writeColumnConstraints(spec.Constraints)
}

func (p *Printer) writeTypeSpecifier(spec ast.TypeSpecifier) {
	p.writeJoinedTokens(spec.Spec)
}

func (p *Printer) writeColumnConstraints(cc ast.ColumnConstraints) {
	if cc.Null != nil {
		p.writeToken(*cc.Null)
	}
	if cc.NotNull != nil {
		p.writeToken(cc.NotNull.Not)
		p.writeToken(cc.NotNull.Null)
	}
	if cc.PrimaryKey != nil {
		p.writeToken(cc.PrimaryKey.Primary)
		p.writeToken(cc.PrimaryKey.Key)
	}
}

func (p *Printer) writeConstraintSpecifier(spec ast.ConstraintSpecifier) {
	if spec.Name != nil {
		p.writeToken(spec.Name.Keywords.Constraint)
		p.writeIdentifier(spec.Name.Name)
	}

	p.writeTableConstraint(spec.Constraint)
}

func (p *Printer) writeTableConstraint(constraint ast.TableConstraint) {
	switch c := constraint.(type) {
	case ast.ForeignKeyConstraint:
		p.writeForeignKeyConstraint(c)
	default:
		panic(fmt.Sprintf("unreachable: unknown table constraint type %T %v", c, c))
	}
}

func (p *Printer) writeForeignKeyConstraint(fk ast.ForeignKeyConstraint) {
	p.writeToken(fk.Keywords.Foreign)
	p.writeToken(fk.Keywords.Key)
	p.writeIdentifierList(fk.Columns)
	p.writeToken(fk.Keywords.References)
	p.writeTableName(fk.RefTableName)
	if fk.RefColumns != nil {
		p.writeIdentifierList(*fk.RefColumns)
	}
}
