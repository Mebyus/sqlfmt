package token

import "strings"

const keywords = endKeyword - beginKeyword - 1

var Keyword = make(map[string]Kind, keywords)

var LowerKeyword = [...]string{
	endKeyword: "",
}

func init() {
	for keyword := Kind(beginKeyword + 1); keyword < endKeyword; keyword++ {
		literal := Literal[keyword]
		Keyword[literal] = keyword
		LowerKeyword[keyword] = strings.ToLower(literal)
	}
}
