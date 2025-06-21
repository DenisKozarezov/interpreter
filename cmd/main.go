package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"interpreter/internal/lexer"
	"interpreter/internal/parser"
)

func main() {
	logger, err := os.OpenFile("cmd/errorsLog.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer func() {
		_ = logger.Close()
	}()

	fileReader, err := os.OpenFile("cmd/invalidProgram.txt", os.O_RDONLY, os.ModeDevice)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer func() {
		_ = fileReader.Close()
	}()

	log.Println("parsing the file...")

	l := lexer.NewLexer(fileReader)
	p := parser.NewParser(l)
	program := p.Parse()

	if len(p.Errors()) > 0 {
		if err = printParserErrors(logger, p.Errors()); err != nil {
			log.Fatalf("failed to print errors: %s", err)
		}
		return
	}

	for i, stmt := range program.Statements {
		log.Printf("[%d]: %s", i, stmt.String())
	}
	log.Println("Parsing completed!")
}

func printParserErrors(out io.Writer, errors []error) error {
	var err error
	if err = outputString(out, "========================== PARSE ERRORS ==========================\n"); err != nil {
		return err
	}

	if err = outputString(out, fmt.Sprintf("%d parser errors found:\n", len(errors))); err != nil {
		return err
	}

	for i := range errors {
		if err = outputString(out, fmt.Sprintf("\t%d. %s\n", i+1, errors[i])); err != nil {
			return err
		}
	}
	return nil
}

func outputString(out io.Writer, s string) error {
	if _, err := out.Write([]byte(s)); err != nil {
		return fmt.Errorf("failed to put a string in output: %w", err)
	}
	return nil
}
