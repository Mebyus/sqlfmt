package main

type Config struct {
	LowerKeywords bool

	// use tabs instead of spaces for indentation
	UseTabs bool

	// number of spaces to use for indentation, only takes effect
	// if UseTabs = false
	Spaces int
}
