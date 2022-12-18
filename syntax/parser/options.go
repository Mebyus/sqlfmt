package parser

type Options struct {
	Input []byte

	// how many errors parser should tolerate before aborting
	//
	// if 0 abort after the first one
	MaxErrors int

	// ignore all parser errors
	//
	// if true MaxErrors setting is ignored
	KeepGoing bool
}
