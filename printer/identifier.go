package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeTableName(name ast.TableName) {
	switch n := name.(type) {
	case ast.Identifier:
		p.writeIdentifier(n)
	case ast.QualifiedIdentifier:
		p.writeQualifiedIdentifier(n)
	default:
		panic("unreachable: unexpected table name type")
	}
}

func (p *Printer) writeQualifiedIdentifier(ident ast.QualifiedIdentifier) {
	p.writeIdentifier(ident.SchemaName)
	p.writeToken(ident.Dot)
	p.writeIdentifier(ident.RawTableName)
}

func (p *Printer) writeIdentifier(ident ast.Identifier) {
	p.writeToken(ident.Token)
}
