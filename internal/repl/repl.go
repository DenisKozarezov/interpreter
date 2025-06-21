package repl

import (
	"bufio"
	"fmt"
	"interpreter/internal/lexer"
	"interpreter/internal/parser"
	"io"
	"log"
	"strings"
)

type REPL struct {
	in     io.Reader
	out    io.Writer
	errors io.Writer

	currentLine int64
}

func NewREPL(in io.Reader, out io.Writer, errors io.Writer) *REPL {
	return &REPL{in: in, out: out, errors: errors}
}

func (r *REPL) StartParser() {
	scanner := bufio.NewScanner(r.in)

	_ = outputString(r.out, "Started to parse an input...")

	for {
		if scanned := scanner.Scan(); !scanned {
			break
		}

		r.currentLine++
		line := scanner.Text()

		l := lexer.NewLexer(strings.NewReader(line))
		p := parser.NewParser(l)
		program := p.Parse()

		if len(p.Errors()) > 0 {
			if err := r.printParserErrors(&line, p.Errors()); err != nil {
				log.Fatalf("failed to print errors: %s", err)
			}
			continue
		}

		for i, stmt := range program.Statements {
			_ = outputString(r.out, "\n[%d]: %s", i, stmt.String())
		}
	}
	_ = outputString(r.out, "Parsing completed!")
}

func (r *REPL) printParserErrors(line *string, errors []error) error {
	var err error

	if err = outputString(r.errors, "⛔ Syntax error in input:\n>>> %s <<<\n", *line); err != nil {
		return err
	}

	if err = outputString(r.errors, "Found %d syntax error(s):\n", len(errors)); err != nil {
		return err
	}

	for i, e := range errors {
		if err = outputString(r.errors, "\t%d. [line %d]%s\n\n", i+1, r.currentLine, e.Error()); err != nil {
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
