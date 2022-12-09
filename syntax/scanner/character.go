package scanner

func isLetterOrUnderscore(c int) bool {
	return isLetter(c) || c == '_'
}

func isAlphanum(c int) bool {
	return isLetterOrUnderscore(c) || isDecimalDigit(c)
}

func isLetter(c int) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isDecimalDigit(c int) bool {
	return '0' <= c && c <= '9'
}

func isDecimalDigitOrPeriod(c int) bool {
	return isDecimalDigit(c) || c == '.'
}

func isWhitespace(c int) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == '\r'
}

func stringFromByte(b byte) string {
	return string([]byte{byte(b)})
}
