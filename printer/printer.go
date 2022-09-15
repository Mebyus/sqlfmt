package printer

import (
	"io"
	"os"

	"github.com/mebyus/sqlfmt/syntax/ast"
)

type Printer struct {
	buf         []byte
	indentation Indentation

	writer io.Writer
}

func Print(stmts []ast.Statement) error {
	p := &Printer{
		writer: os.Stdout,
		indentation: Indentation{
			s: "    ",
		},
	}
	return p.Print(stmts)
}

func (p *Printer) Print(stmts []ast.Statement) error {
	for _, stmt := range stmts {
		p.writeStatement(stmt)
	}
	_, err := p.writer.Write(p.buf)
	return err
}

func (p *Printer) write(s string) {
	p.buf = append(p.buf, []byte(s)...)
}

func (p *Printer) nextLine() {
	p.write("\n")
	p.indent()
}

func (p *Printer) indent() {
	p.write(p.indentation.Str())
}
