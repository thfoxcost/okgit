/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	config "okgit/internal/client/config"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove the generated config file",
	Long:  "removes the JSON config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return config.Rm()
	},
}

func init() {
	configCmd.AddCommand(rmCmd)
}
