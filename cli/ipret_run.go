package cli

import (
	"errors"
	"fmt"
	"interpreter/internal/repl"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func newRunCommand() *cobra.Command {
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Starts to interpret source code.",
		Long:  "Starts to interpret source code specified in --filename (-f) flag.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			filename, _ := cmd.Flags().GetString("filename")
			if strings.TrimSpace(filename) == "" {
				return errors.New("--filename argument is empty")
			}

			fileReader, err := os.OpenFile(filename, os.O_RDONLY, os.ModeDevice)
			if err != nil {
				return fmt.Errorf("failed to open file: %w", err)
			}
			defer func() { _ = fileReader.Close() }()

			r := repl.NewREPL(fileReader, os.Stdout, os.Stderr)

			isBench, _ := cmd.Flags().GetBool("bench")

			start := time.Now()
			r.StartParser()
			if end := time.Since(start).Seconds(); isBench {
				cmd.Printf("Benchmark result: %f seconds\n", end)
			}
			return nil
		},
	}

	runCmd.Flags().StringP("filename", "f", "", "Path to the source file.")
	_ = runCmd.MarkFlagRequired("filename")
	runCmd.Flags().Bool("bench", true, "Turns on a benchmark and shows total time after execution.")
	runCmd.Flags().Bool("benchmem", true, "Turns on a benchmark and shows both total time and memory consumption after execution.")
	runCmd.MarkFlagsMutuallyExclusive("bench", "benchmem")

	return runCmd
}
