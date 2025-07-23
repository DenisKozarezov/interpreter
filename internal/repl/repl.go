package repl

import (
	"fmt"
	"io"
	"log"
	"os"

	"interpreter/internal/evaluator"
	"interpreter/internal/lexer"
	"interpreter/internal/parser"
)

type REPL struct {
	in     io.ReaderAt
	out    io.Writer
	errors io.Writer
}

func NewREPL(in io.ReaderAt, out io.Writer, errors io.Writer) *REPL {
	return &REPL{in: in, out: out, errors: errors}
}

func (r *REPL) StartParser() {
	r.log(r.out, "Started to parse an input...")

	l := lexer.NewLexer(r.in)
	p := parser.NewParser(l)
	program := p.Parse()

	if len(p.Errors()) > 0 {
		r.printParserErrors(p.Errors())
		r.log(r.out, "Parser errors found! Stopping the interpreter...")
		os.Exit(1)
	}

	if result := evaluator.EvaluateStatement(program); result != nil {
		r.log(r.out, "%s", result.Inspect())
	}

	r.log(r.out, "Parsing completed!")
}

func (r *REPL) printParserErrors(errors []error) {
	r.log(r.errors, "⛔ Found %d syntax error(s):", len(errors))

	for i, e := range errors {
		r.log(r.errors, "\t%d. %s\n", i+1, e.Error())
	}
}

func (r *REPL) log(w io.Writer, s string, args ...any) {
	if _, err := fmt.Fprintf(w, s+"\n", args...); err != nil {
		log.Fatal("failed to put a string in output: %w", err)
	}
}
