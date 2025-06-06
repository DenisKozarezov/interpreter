package main

import (
	"bufio"
	"fmt"
	"os"

	"interpreter/internal/lexer"
	"interpreter/internal/lexer/tokens"
)

const PROMPT = ">> "

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.NewLexer(line)

		for token := lexer.NextToken(); token.Type != tokens.EOF; token = lexer.NextToken() {
			fmt.Printf("%+v\n", token)
		}
	}
}
