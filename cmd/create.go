/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"okgit/internal/client/config"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the configuration file",
	Long:  `Create the configuration directory and initialize the config.json file with default settings.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return config.Create()
	},
}

func init() {
	configCmd.AddCommand(createCmd)
}
