package cli

import (
	"errors"
	"fmt"
	"os"
	"runtime"

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
	rootCmd.Version = version

	// Register custom template functions
	cobra.AddTemplateFunc("BuildDate", func() string { return buildDate })
	cobra.AddTemplateFunc("GOOS", func() string { return runtime.GOOS })
	cobra.AddTemplateFunc("GOARCH", func() string { return runtime.GOARCH })
	cobra.AddTemplateFunc("GoVersion", func() string { return runtime.Version() })

	// Set the version template
	rootCmd.SetVersionTemplate(`Version:    {{.Version}}
Built:      {{BuildDate}}
Platform:   {{GOOS}}/{{GOARCH}}
Go Version: {{GoVersion}}
`)

	rootCmd.AddCommand(newRunCommand())
}
