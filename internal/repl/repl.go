package repl

import (
	"fmt"
	"io"
	"os"

	"interpreter/internal/evaluator"
	"interpreter/internal/lexer"
	"interpreter/internal/parser"
)

type REPL struct {
	stdIn  io.ReaderAt
	stdOut io.Writer
	stdErr io.Writer
}

func NewREPL(stdIn io.ReaderAt, stdOut io.Writer, stdErr io.Writer) *REPL {
	return &REPL{stdIn: stdIn, stdOut: stdOut, stdErr: stdErr}
}

func (r *REPL) StartParser() {
	log(r.stdOut, "Started to parse an input...")

	l := lexer.NewLexer(r.stdIn)
	p := parser.NewParser(l)
	program := p.Parse()

	if len(p.Errors()) > 0 {
		r.printParserErrors(p.Errors())
		log(r.stdOut, "Parser error found! Stopping the interpreter...")
		os.Exit(1)
	}

	v := evaluator.NewASTVisitor()
	if result := evaluator.EvaluateStatement(program, v); result != nil {
		log(r.stdOut, "%s", result.Inspect())
	}

	log(r.stdOut, "Parsing completed!")
}

func (r *REPL) printParserErrors(errors []error) {
	log(r.stdErr, "⛔ Found %d syntax error(s):", len(errors))

	for i, e := range errors {
		logf(r.stdErr, "\t%d. %s\n", i+1, e.Error())
	}
}

func log(w io.Writer, s string, args ...any) {
	logf(w, s+"\n", args...)
}

func logf(w io.Writer, s string, args ...any) {
	_, _ = fmt.Fprintf(w, s, args...)
}
