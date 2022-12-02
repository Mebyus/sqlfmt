package wsemitter

import "strings"

type Indenter struct {
	s   string
	buf []byte
}

func ConfigureIndenter(useTabs bool, spaces int) Indenter {
	var s string
	if useTabs {
		s = "\t"
	} else {
		s = strings.Repeat(" ", spaces)
	}
	return Indenter{
		s: s,
	}
}

func (i *Indenter) Inc() {
	i.buf = append(i.buf, []byte(i.s)...)
}

func (i *Indenter) Dec() {
	i.buf = i.buf[:len(i.buf)-len(i.s)]
}

func (i Indenter) Str() string {
	return string(i.buf)
}

