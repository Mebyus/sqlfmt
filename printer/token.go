package printer

import "github.com/mebyus/sqlfmt/syntax/token"

func (p *Printer) writeToken(tok token.Token) {
	p.ws(tok.Kind)

	if tok.Kind.HasStaticLiteral() {
		if tok.Kind.IsKeyword() {
			p.writeKeyword(tok.Kind, tok.Index)
			return
		}
		lit := token.Literal[tok.Kind]
		p.writeWithComments(tok.Kind, lit, tok.Index)
		return
	}
	p.writeWithComments(tok.Kind, tok.Lit, tok.Index)
}

func (p *Printer) writeJoinedTokens(tokens []token.Token) {
	for _, tok := range tokens {
		p.writeToken(tok)
	}
}

func (p *Printer) writeKeyword(kind token.Kind, index int) {
	lit := p.keyword[kind]
	p.writeWithComments(kind, lit, index)
}

func (p *Printer) writeWithComments(kind token.Kind, s string, index int) {
	if p.index >= index {
		p.index++
		p.write(s)
		return
	}

	for p.index < index {
		p.writeComment(p.comms[p.next])
		p.next++
		p.index++
	}
	p.ws(kind)
	p.write(s)
	p.index++
}
