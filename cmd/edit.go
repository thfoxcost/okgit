/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"okgit/internal/client/config"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the configuration file",
	Long:  `Open the okgit configuration file in your preferred editor.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return config.Edit()
	},
}

func init() {
	configCmd.AddCommand(editCmd)
}
