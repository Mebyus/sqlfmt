package token

const keywords = endKeyword - beginKeyword - 1

var Keyword = make(map[string]Kind, keywords)

func init() {
	for keyword := Kind(beginKeyword + 1); keyword < endKeyword; keyword++ {
		literal := Literal[keyword]
		Keyword[literal] = keyword
	}
}
