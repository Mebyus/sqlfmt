package printer

import "io"

type Options struct {
	Writer io.Writer

	LowerKeywords bool

	// use tabs instead of spaces for indentation
	UseTabs bool

	// number of spaces to use for indentation, only takes effect
	// if UseTabs = false
	Spaces int

	// guideline (not a hard limit) for output width in characters
	Width int
}

var DefaultOptions = Options{
	UseTabs: true,
	Spaces:  4,
	Width:   120,
}
