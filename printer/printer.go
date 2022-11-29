package printer

import (
	"io"
	"os"

	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/token"
)

type Printer struct {
	buf         []byte
	indentation Indentation

	// number of tokens already written
	index int

	// index of next comment to be written
	next int

	comms []ast.Comment

	options Options
	keyword []string
	writer  io.Writer
}

func Print(file ast.SQLFile, options Options) error {
	p := &Printer{
		writer:      os.Stdout,
		indentation: InitIdentation(options.UseTabs, options.Spaces),
		options:     options,
	}
	if options.LowerKeywords {
		p.keyword = token.LowerKeyword[:]
	} else {
		p.keyword = token.Literal[:]
	}
	return p.Print(file)
}

func (p *Printer) Print(file ast.SQLFile) error {
	p.comms = file.Comments

	for _, stmt := range file.Statements {
		p.writeStatement(stmt)
		p.nl()
	}
	for i := p.next; i < len(p.comms); i++ {
		p.writeComment(p.comms[i])
	}

	_, err := p.writer.Write(p.buf)
	return err
}

func (p *Printer) write(s string) {
	p.buf = append(p.buf, []byte(s)...)
}

func (p *Printer) space() {
	p.write(" ")
}

func (p *Printer) nl() {
	p.write("\n")
	p.indent()
}

func (p *Printer) indent() {
	p.write(p.indentation.Str())
}
