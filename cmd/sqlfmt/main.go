package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mebyus/sqlfmt/syntax/scanner"
	"github.com/mebyus/sqlfmt/syntax/token"
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
	s, err := scanner.FromFile(filename)
	if err != nil {
		fatal(err)
	}
	for {
		tok := s.Scan()
		fmt.Println(tok.String())
		if tok.Kind == token.EOF {
			break
		}
	}
}
