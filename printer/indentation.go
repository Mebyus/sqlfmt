package printer

type Indentation struct {
	s   string
	buf []byte
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
