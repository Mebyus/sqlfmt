package wsemitter

import "github.com/mebyus/sqlfmt/syntax/token"

type WhitespaceKind int

const (
	None WhitespaceKind = iota
	Newline
	Space
	Indentation
)

type Emitter struct {
	next WhitespaceKind
	ind  Indenter
	pre  Predictor
}

func ConfigureEmitter(useTabs bool, spaces int) Emitter {
	return Emitter{
		next: None,
		ind:  ConfigureIndenter(useTabs, spaces),
		pre:  DefaultPredictor,
	}
}

func (ws *Emitter) Emit(tok token.Kind) string {
	next := ws.next
	ws.next = ws.pre.Predict(tok)
	return ws.emit(ws.pre.Override(next, tok))
}

func (ws *Emitter) emit(kind WhitespaceKind) string {
	switch kind {
	case None:
		return ""
	case Newline:
		return "\n"
	case Space:
		return " "
	case Indentation:
		return ws.ind.Str()
	default:
		panic("unreachable: unexpected whitespace kind")
	}
}

func (ws *Emitter) None() {
	ws.next = None
}

func (ws *Emitter) Indent() {
	ws.next = Indentation
}

func (ws *Emitter) Space() {
	ws.next = Space
}

func (ws *Emitter) Inc() {
	ws.ind.Inc()
}

func (ws *Emitter) Dec() {
	ws.ind.Dec()
}
