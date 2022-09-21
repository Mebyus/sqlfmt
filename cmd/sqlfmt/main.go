package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mebyus/sqlfmt/printer"
	"github.com/mebyus/sqlfmt/syntax/parser"
)

func fatal(v any) {
	fmt.Println("fatal:", v)
	os.Exit(1)
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fatal("filename was not specified")
	}

	filename := flag.Arg(0)
	stmts, err := parser.ParseFile(filename)
	if err != nil {
		fatal(err)
	}
	err = printer.Print(stmts)
	if err != nil {
		fatal(err)
	}
}
