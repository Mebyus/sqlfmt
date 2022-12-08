package printer

import "github.com/mebyus/sqlfmt/syntax/ast"

func (p *Printer) writeIdentifierList(list ast.IdentifierList) {
	p.writeToken(list.LeftParentheses)
	for _, elem := range list.Elements {
		if elem.Comma != nil {
			p.writeToken(*elem.Comma)
		}
		p.writeIdentifier(elem.Name)
	}
	p.writeToken(list.RightParentheses)
}
