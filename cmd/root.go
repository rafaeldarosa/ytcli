// cmd/root.go
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	outputFormat string
	maxResults   int
)

var rootCmd = &cobra.Command{
	Use:   "ytcli",
	Short: "YouTube CLI - Terminal interface for YouTube",
	Long: `A feature-rich CLI application for interacting with YouTube.
Manage subscriptions, discover trending videos, and more.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "pretty", "Output format (pretty/simple/json)")
	rootCmd.PersistentFlags().IntVarP(&maxResults, "limit", "l", 5, "Maximum number of results to show")
}
