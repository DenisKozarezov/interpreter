package main

import (
	"log"
	"os"

	"interpreter/internal/repl"
)

func main() {
	errorsLogger, err := os.OpenFile("example/errorsLog.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer func() {
		_ = errorsLogger.Close()
	}()

	fileReader, err := os.OpenFile("example/invalidProgram.txt", os.O_RDONLY, os.ModeDevice)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer func() {
		_ = fileReader.Close()
	}()

	r := repl.NewREPL(fileReader, log.Writer(), errorsLogger)
	r.StartParser()
}
