package printer

import "strings"

type Indentation struct {
	s   string
	buf []byte
}

func InitIdentation(useTabs bool, spaces int) Indentation {
	var s string
	if useTabs {
		s = "\t"
	} else {
		s = strings.Repeat(" ", spaces)
	}
	return Indentation{
		s: s,
	}
}

func (i *Indentation) Inc() {
	i.buf = append(i.buf, []byte(i.s)...)
}

func (i *Indentation) Dec() {
	i.buf = i.buf[:len(i.buf)-len(i.s)]
}

func (i Indentation) Str() string {
	return string(i.buf)
}
