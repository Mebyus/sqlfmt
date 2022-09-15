package printer

import "github.com/mebyus/sqlfmt/syntax/token"

func (p *Printer) writeToken(tok token.Token) {
	p.write(tok.Lit)
}
