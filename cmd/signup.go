/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"

	"okgit/internal"
	"okgit/internal/client/auth"
)

type SingupData struct {
	Username string `json:username`
	Email    string `json:email`
	Password string `json:password`
}

func SignupHandler(data SingupData) error {
	baseURL, _, err := internal.Endpoint()
	SignupRoute := baseURL + "/api/auth/signup"

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		http.MethodPost,
		SignupRoute,
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("signup failed: %s", string(body))
	}

	fmt.Println(string(body))
	return nil
}

var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		username64, email64, password64, err := auth.Signup()
		if err != nil {
			fmt.Println("Signup failed:", err)
			return
		}
		data := SingupData{
			Username: username64,
			Email:    email64,
			Password: string(password64),
		}
		errH := SignupHandler(data)
		if errH != nil {
			fmt.Println("Signup failed:", err)
			return
		}
	},
}

func init() {
	authCmd.AddCommand(signupCmd)
}
