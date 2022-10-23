package printer

import "github.com/mebyus/sqlfmt/syntax/token"

func (p *Printer) writeToken(tok token.Token) {
	if tok.Kind.HasStaticLiteral() {
		p.write(token.Literal[tok.Kind])
		return
	}
	p.write(tok.Lit)
}
