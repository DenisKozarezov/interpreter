package main

import (
	"interpreter/internal/repl"
	"log"
	"os"
)

func main() {
	errorsLogger, err := os.OpenFile("cmd/errorsLog.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer func() {
		_ = errorsLogger.Close()
	}()

	fileReader, err := os.OpenFile("cmd/program2.txt", os.O_RDONLY, os.ModeDevice)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer func() {
		_ = fileReader.Close()
	}()

	r := repl.NewREPL(fileReader, log.Writer(), errorsLogger)
	r.StartParser()
}
