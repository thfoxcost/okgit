/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	internal "okgit/internal"

	"github.com/spf13/cobra"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Ping() error {
	endpoint, _, err := internal.Endpoint()
	if err != nil {
		return fmt.Errorf("load server endpoint: %w", err)
	}
	resp, err := http.Get(endpoint + "/api/status")
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned: %s", resp.Status)
	}

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	fmt.Printf("\033[32m%s\033[0m\n", result.Status)
	return nil
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Check the connection to your OKGit server",
	Long: `Check whether your configured OKGit server is reachable.
This command sends a request to the server and reports
whether the connection was successful.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Ping()
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
