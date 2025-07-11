package repl

import (
	"fmt"
	"interpreter/internal/lexer"
	"interpreter/internal/parser"
	"io"
	"log"
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
	_ = outputString(r.out, "Started to parse an input...\n")

	l := lexer.NewLexer(r.in)
	p := parser.NewParser(l)
	program := p.Parse()

	if len(p.Errors()) > 0 {
		if err := r.printParserErrors(p.Errors()); err != nil {
			log.Fatalf("failed to print errors: %s", err)
		}
	}

	for i, stmt := range program.Statements {
		_ = outputString(r.out, "\n[%d]: %s", i, stmt.String())
	}

	_ = outputString(r.out, "\nParsing completed!")
}

func (r *REPL) printParserErrors(errors []error) error {
	var err error

	if err = outputString(r.errors, "⛔ Found %d syntax error(s):\n", len(errors)); err != nil {
		return err
	}

	for i, e := range errors {
		if err = outputString(r.errors, "\t%d. %s\n\n", i+1, e.Error()); err != nil {
			return err
		}
	}

	return nil
}

func outputString(out io.Writer, s string, args ...any) error {
	if _, err := out.Write([]byte(fmt.Sprintf(s, args...))); err != nil {
		return fmt.Errorf("failed to put a string in output: %w", err)
	}
	return nil
}
