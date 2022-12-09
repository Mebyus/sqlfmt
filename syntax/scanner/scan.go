package scanner

import (
	"fmt"
	"strings"

	"github.com/mebyus/sqlfmt/syntax/token"
)

func (s *Scanner) Scan() token.Token {
	tok := s.scan()
	tok.Index = s.index
	s.index++
	fmt.Println(tok.String())
	return tok
}

func (s *Scanner) scan() token.Token {
	if s.c == eof {
		return s.createToken(token.EOF)
	}

	s.skipWhitespace()
	if s.c == eof {
		return s.createToken(token.EOF)
	}

	if isLetterOrUnderscore(s.c) {
		return s.scanName()
	}

	if isDecimalDigit(s.c) {
		return s.scanNumber()
	}

	if s.c == '\'' {
		return s.scanString()
	}

	if s.c == '"' {
		return s.scanQuoted()
	}

	if s.c == '-' && s.next == '-' {
		return s.scanLineComment()
	}

	if s.c == '/' && s.next == '*' {
		return s.scanMultiLineComment()
	}

	return s.scanOther()
}

func (s *Scanner) createToken(kind token.Kind) token.Token {
	return token.Token{
		Kind: kind,
		Pos:  s.pos,
	}
}

func (s *Scanner) scanName() (tok token.Token) {
	tok.Pos = s.pos

	s.storeWord()
	lit := s.collect()
	keyword, ok := token.Keyword[strings.ToUpper(lit)]
	if ok {
		tok.Kind = keyword
		return
	}

	tok.Kind = token.Identifier
	tok.Lit = lit
	return
}

func (s *Scanner) scanString() (tok token.Token) {
	tok.Pos = s.pos
	s.store() // consume '

	for s.c != eof && s.c != '\'' {
		s.store()
	}

	if s.c == eof {
		tok.Kind = token.Illegal
	} else {
		tok.Kind = token.String
		s.store() // consume '
	}
	tok.Lit = s.collect()
	return
}

func (s *Scanner) scanQuoted() (tok token.Token) {
	tok.Pos = s.pos
	s.store() // consume "

	for s.c != eof && s.c != '"' {
		s.store()
	}

	if s.c == eof {
		tok.Kind = token.Illegal
	} else {
		tok.Kind = token.QuotedIdentifier
		s.store() // consume "
	}
	tok.Lit = s.collect()
	return
}

func (s *Scanner) scanDecimalNumber() (tok token.Token) {
	tok.Pos = s.pos

	scannedOnePeriod := false
	for s.c != eof && isDecimalDigitOrPeriod(s.c) {
		s.store()
		if s.c == '.' {
			if scannedOnePeriod {
				break
			}
			scannedOnePeriod = true
		}
	}

	if isAlphanum(s.c) || s.c == '.' {
		s.store()
		s.storeWord()
		tok.Kind = token.Illegal
		tok.Lit = s.collect()
		return
	}

	if s.prev == '.' {
		tok.Kind = token.Illegal
		tok.Lit = s.collect()
		return
	}

	if scannedOnePeriod {
		tok.Kind = token.DecimalFloat
	} else {
		tok.Kind = token.DecimalInteger
	}

	tok.Lit = s.collect()
	return
}

func (s *Scanner) scanNumber() (tok token.Token) {
	if s.c != '0' {
		return s.scanDecimalNumber()
	}

	if s.next == eof {
		tok = token.Token{
			Kind: token.DecimalInteger,
			Pos:  s.pos,
			Lit:  stringFromByte('0'),
		}
		s.advance()
		return
	}

	if s.next == '.' {
		return s.scanDecimalNumber()
	}

	if isAlphanum(s.next) {
		return s.scanIllegalWord()
	}

	tok = token.Token{
		Kind: token.DecimalInteger,
		Pos:  s.pos,
		Lit:  stringFromByte('0'),
	}
	s.advance()
	return
}

func (s *Scanner) scanLineComment() (tok token.Token) {
	tok.Pos = s.pos

	for s.c != eof && s.c != '\n' {
		s.store()
	}

	tok.Kind = token.LineComment
	tok.Lit = s.collect()
	return
}

func (s *Scanner) scanMultiLineComment() (tok token.Token) {
	tok.Pos = s.pos

	for s.c != eof && !(s.c == '*' && s.next == '/') {
		s.store()
	}

	if s.c == eof {
		tok.Kind = token.Illegal
	} else {
		tok.Kind = token.MultiLineComment
		s.store() // consume *
		s.store() // consume /
	}
	tok.Lit = s.collect()
	return
}

func (s *Scanner) scanOneByteToken(kind token.Kind) token.Token {
	tok := s.createToken(kind)
	s.advance()
	return tok
}

func (s *Scanner) scanTwoByteToken(kind token.Kind) token.Token {
	tok := s.createToken(kind)
	s.advance()
	s.advance()
	return tok
}

func (s *Scanner) scanIllegalWord() (tok token.Token) {
	tok = s.createToken(token.Illegal)
	s.storeWord()
	tok.Lit = s.collect()
	return
}

func (s *Scanner) scanIllegalByteToken() token.Token {
	tok := token.Token{
		Kind: token.Illegal,
		Pos:  s.pos,
		Lit:  stringFromByte(byte(s.c)),
	}
	s.advance()
	return tok
}

func (s *Scanner) scanOther() token.Token {
	switch s.c {
	case '(':
		return s.scanOneByteToken(token.LeftParentheses)
	case ')':
		return s.scanOneByteToken(token.RightParentheses)
	case '.':
		return s.scanOneByteToken(token.Dot)
	case ';':
		return s.scanOneByteToken(token.Semicolon)
	case '*':
		return s.scanOneByteToken(token.Asterisk)
	case '+':
		return s.scanOneByteToken(token.Plus)
	case ',':
		return s.scanOneByteToken(token.Comma)
	case '-':
		return s.scanOneByteToken(token.Minus)
	case '/':
		return s.scanOneByteToken(token.Slash)
	case '%':
		return s.scanOneByteToken(token.Percent)
	case '=':
		return s.scanOneByteToken(token.Equal)
	case '!':
		if s.next == '=' {
			return s.scanTwoByteToken(token.NotEqual)
		}
		return s.scanIllegalByteToken()
	case '<':
		if s.next == '=' {
			return s.scanTwoByteToken(token.LessEqual)
		} else if s.next == '>' {
			return s.scanTwoByteToken(token.NotEqualAlt)
		}
		return s.scanOneByteToken(token.Less)
	case '>':
		if s.next == '=' {
			return s.scanTwoByteToken(token.GreaterEqual)
		}
		return s.scanOneByteToken(token.Greater)
	case ':':
		if s.next == ':' {
			return s.scanTwoByteToken(token.DoubleColon)
		}
		return s.scanIllegalByteToken()
	default:
		return s.scanIllegalByteToken()
	}
}
