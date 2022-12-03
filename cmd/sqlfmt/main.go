package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mebyus/sqlfmt/printer"
	"github.com/mebyus/sqlfmt/syntax/parser"
)

func fatal(v any) {
	fmt.Fprintln(os.Stderr, "fatal:", v)
	os.Exit(1)
}

func main() {
	var config Config

	flag.BoolVar(&config.LowerKeywords, "lower", false, "format keywords in lowercase")
	flag.BoolVar(&config.UseTabs, "tabs", false, "use tabs instead of spaces for indentation")
	flag.IntVar(&config.Spaces, "spaces", 4, "number of spaces to use for indentation")
	flag.IntVar(&config.Width, "width", 120, "guideline (not a hard limit) for output width in characters")

	flag.Parse()

	if flag.NArg() == 0 {
		fatal("filename was not specified")
	}

	if config.Spaces < 0 {
		fatal("spaces cannot be negative")
	}

	if config.Width <= 0 {
		fatal("width must be positive")
	}

	filename := flag.Arg(0)
	file, err := parser.ParseFile(filename)
	if err != nil {
		fatal(err)
	}
	err = printer.Print(file, printer.Options{
		Writer:        os.Stdout,
		LowerKeywords: config.LowerKeywords,
		UseTabs:       config.UseTabs,
		Spaces:        config.Spaces,
		Width:         config.Width,
	})
	if err != nil {
		fatal(err)
	}
}
