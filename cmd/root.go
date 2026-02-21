package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quran",
	Short: "Minimal HTTP server service Quran data",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("still on progress..")
		return errors.New("in development")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
