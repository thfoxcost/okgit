package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Show information about okgit",
	Long:  "Show a short description and repo GitHub link for okgit.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`okgit is a CLI-first code hosting platform for people who live in the terminal.
Create repositories, browse commits, manage issues, and review pull requests
without ever opening a browser. Everything is built around the command line:
fast, scriptable, and keyboard-friendly.

To learn more, visit:
https://github.com/thfoxcost/okgit`)
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
