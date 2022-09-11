package scanner

func isLetterOrUnderscore(b int) bool {
	return ('a' <= b && b <= 'z') || b == '_' || ('A' <= b && b <= 'Z')
}

func isAlphanum(b int) bool {
	return ('a' <= b && b <= 'z') || b == '_' || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9')
}

func isLetter(b int) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z')
}

func isDecimalDigit(b int) bool {
	return '0' <= b && b <= '9'
}

func isDecimalDigitOrPeriod(b int) bool {
	return ('0' <= b && b <= '9') || b == '.'
}

func isWhitespace(b int) bool {
	return b == ' ' || b == '\n' || b == '\t' || b == '\r'
}

func stringFromByte(b byte) string {
	return string([]byte{byte(b)})
}
