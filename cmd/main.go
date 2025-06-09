package main

import (
	"log"
	"os"

	"interpreter/internal/lexer"
	"interpreter/internal/parser"
)

func main() {
	fileReader, err := os.Open("cmd/someFile.txt")
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

	for i, stmt := range program.Statements {
		log.Printf("[%d]: %s", i, stmt.String())
	}
	log.Println("parsing completed!")
}
