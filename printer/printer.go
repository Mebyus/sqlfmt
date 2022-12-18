package printer

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/mebyus/sqlfmt/printer/wsemitter"
	"github.com/mebyus/sqlfmt/syntax/ast"
	"github.com/mebyus/sqlfmt/syntax/token"
)

type Printer struct {
	buf []byte
	wse wsemitter.Emitter

	// number of tokens already written
	index int

	// index of next comment to be written
	next int

	stmts []ast.Statement
	comms []ast.Comment

	options Options
	keyword []string
	writer  io.Writer
}

func Print(file ast.SQLFile, options Options) error {
	p := &Printer{
		writer:  options.Writer,
		wse:     wsemitter.ConfigureEmitter(options.UseTabs, options.Spaces),
		options: options,
	}
	if options.LowerKeywords {
		p.keyword = token.LowerKeyword[:]
	} else {
		p.keyword = token.Literal[:]
	}
	return p.Print(file)
}

func (p *Printer) Print(file ast.SQLFile) error {
	p.stmts = file.Statements
	p.comms = file.Comments

	start := time.Now()
	p.print()
	fmt.Fprintln(os.Stderr, "writing took:", time.Since(start))

	_, err := p.writer.Write(p.buf)
	return err
}

func (p *Printer) print() {
	for _, stmt := range p.stmts {
		p.writeStatement(stmt)
		p.nl()
	}
	for i := p.next; i < len(p.comms); i++ {
		p.writeComment(p.comms[i])
	}
}

func (p *Printer) write(s string) {
	p.buf = append(p.buf, []byte(s)...)
}

func (p *Printer) ws(tok token.Kind) {
	p.write(p.wse.Emit(tok))
}

func (p *Printer) nl() {
	p.write("\n")
	p.wse.Indent()
}
