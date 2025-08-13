package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version   = "dev"
	buildDate = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "ipret",
	Short: "Lightweight scripting language interpreter",
	Long: `IPRET (Interpretive Runtime) - fast and minimalistic interpreter for 
custom scripting language with REPL support.

Complete documentation available at https://github.com/DenisKozarezov/interpreter`,
	Example: `ipret --version
ipret run --filename ./someFile.txt --bench
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion, _ := cmd.Flags().GetBool("version"); showVersion {

		}
	},
}

func Execute() error {
	defer func() {
		if recovered := recover(); recovered != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Fatal error occured while running ipret: %s\n", recovered)
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		_, logErr := fmt.Fprintf(os.Stderr, "An error occured while running ipret: %s\n", err)
		return errors.Join(err, logErr)
	}
	return nil
}

func Init() {
	rootCmd.Version = fmt.Sprintf("%s (Build Date: %s)\n", version, buildDate)
	rootCmd.SetVersionTemplate(`{{with .Name}}{{printf "%s " .}}{{end}}{{printf "version: %s" .Version}}`)

	rootCmd.AddCommand(newRunCommand())
}
