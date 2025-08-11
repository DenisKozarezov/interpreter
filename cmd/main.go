package main

import (
	"interpreter/cli"
	"os"
)

func init() {
	cli.Init()
}

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
