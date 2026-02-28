package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/fatahnuram/quran/internal/server"
	"github.com/spf13/cobra"
)

var (
	// flags
	port int

	// cmd
	rootCmd = &cobra.Command{
		Use:     "quran",
		Short:   "Minimal HTTP server service Quran data",
		Version: VERSION,
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.ServeHttp(port)
		},
	}
)

func init() {
	rootCmd.Flags().IntVarP(&port, "port", "p", 8087, "HTTP server port")
	rootCmd.SetVersionTemplate("{{.Version}}\n")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Info(fmt.Sprintf("got runtime error: %v", err))
		os.Exit(1)
	}
}
