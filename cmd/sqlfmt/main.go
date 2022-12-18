package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

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
	flag.BoolVar(&config.UseStdout, "stdout", false, "use stdout instead of output file")
	flag.BoolVar(&config.KeepGoing, "keep-going", false, "keep parsing even after error limit is reached")
	flag.IntVar(&config.Spaces, "spaces", 4, "number of spaces to use for indentation")
	flag.IntVar(&config.Width, "width", 120, "guideline (not a hard limit) for output width in characters")
	flag.IntVar(&config.MaxErrors, "max-errors", 10, "abort parsing after encountering certain number of errors")
	flag.StringVar(&config.OutputFile, "o", "", "output file")

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

	output, err := configureOutput(config.OutputFile, config.UseStdout)
	if err != nil {
		fatal(err)
	}
	defer output.Close()

	filename := flag.Arg(0)
	b, err := os.ReadFile(filename)
	if err != nil {
		fatal(err)
	}

	start := time.Now()
	file, err := parser.New(parser.Options{
		Input:     b,
		MaxErrors: config.MaxErrors,
		KeepGoing: config.KeepGoing,
	}).Parse()
	if err != nil {
		fatal(err)
	}
	fmt.Fprintln(os.Stderr, "parsing took:", time.Since(start))
	if file.NumberOfErrors > 0 {
		fmt.Fprintln(os.Stderr, "number of errors:", file.NumberOfErrors)
	}

	err = printer.Print(file, printer.Options{
		Writer:        output,
		LowerKeywords: config.LowerKeywords,
		UseTabs:       config.UseTabs,
		Spaces:        config.Spaces,
		Width:         config.Width,
	})
	if err != nil {
		fatal(err)
	}
}

func configureOutput(outputFile string, useStdout bool) (output io.WriteCloser, err error) {
	if useStdout || outputFile == "" {
		return os.Stdout, nil
	}
	dir := filepath.Dir(outputFile)
	if dir != "" {
		err = os.MkdirAll(dir, 0o755)
		if err != nil {
			return nil, err
		}
	}
	file, err := os.Create(outputFile)
	if err != nil {
		return nil, err
	}
	return file, nil
}
