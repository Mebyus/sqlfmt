package printer

import "github.com/mebyus/sqlfmt/syntax/token"

func (p *Printer) writeToken(tok token.Token) {
	if tok.Kind.HasStaticLiteral() {
		if tok.Kind.IsKeyword() {
			p.writeKeyword(tok.Kind, tok.Index)
			return
		}
		lit := token.Literal[tok.Kind]
		p.writeWithComments(lit, tok.Index)
		return
	}
	p.writeWithComments(tok.Lit, tok.Index)
}

func (p *Printer) writeJoinedTokens(tokens []token.Token) {
	if len(tokens) == 0 {
		return
	}
	for i := 0; i < len(tokens)-1; i++ {
		tok := tokens[i]
		p.writeToken(tok)
		p.space()
	}
	last := tokens[len(tokens)-1]
	p.writeToken(last)
}

func (p *Printer) writeKeyword(kind token.Kind, index int) {
	lit := p.keyword[kind]
	p.writeWithComments(lit, index)
}

func (p *Printer) writeWithComments(s string, index int) {
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
	p.space()
	p.write(s)
	p.index++
}
