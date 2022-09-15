package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeTableName(name ast.TableName) {
	switch n := name.(type) {
	case ast.Identifier:
		p.writeIdentifier(n)
	case ast.QualifiedIdentifier:
		p.writeIdentifier(n.SchemaName)
		p.write(".")
		p.writeIdentifier(n.RawTableName)
	default:
		panic("unexpected table name type")
	}
}

func (p *Printer) writeIdentifier(ident ast.Identifier) {
	p.writeToken(ident.Token)
}
