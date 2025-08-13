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
		Short: "Execute a source code file",
		Long: `Execute and interpret the specified source code file.

The command requires a valid source file path provided via --filename flag.
Execution results will be displayed in the console output.`,
		Example: `ipret run -f script.irt
ipret run --filename=script.irt --bench`,
		Args: cobra.NoArgs,
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

	runCmd.Flags().StringP("filename", "f", "", "Path to source file to execute (required)")
	_ = runCmd.MarkFlagRequired("filename")
	runCmd.Flags().Bool("bench", false, "Enable execution time benchmarking\n(shows total duration after completion)")
	runCmd.Flags().Bool("benchmem", false, "Enable detailed benchmarking\n(shows both execution time and memory allocation statistics)")
	runCmd.MarkFlagsMutuallyExclusive("bench", "benchmem")

	return runCmd
}
