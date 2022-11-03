package parser

func (p *Parser) parseCreateStatement() error {
	// p.kind = statement.Create

	// for !p.isEOF() && p.tok.Kind != token.Semicolon {
	// 	p.advance()
	// }
	// p.advance()
	p.consumeUnknownStatement()
	return nil
}
