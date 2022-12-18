package main

type Config struct {
	LowerKeywords bool

	// use tabs instead of spaces for indentation
	UseTabs bool

	// number of spaces to use for indentation, only takes effect
	// if UseTabs = false
	Spaces int

	// guideline (not a hard limit) for output width in characters
	Width int

	// how many errors parser should tolerate before aborting
	//
	// if 0 abort after the first one
	MaxErrors int

	// ignore all parser errors
	//
	// if true MaxErrors setting is ignored
	KeepGoing bool

	OutputFile string

	UseStdout bool
}
