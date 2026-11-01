package cmd

import "github.com/spf13/cobra"

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage okgit configuration",
	Long:  "Create, list, update, and remove okgit configuration.",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
