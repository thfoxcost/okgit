package auth

import (
	"encoding/base64"
	"fmt"
	"os"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/term"
)

func Signup() (string, string, string, error) {
	var username string
	var email string

	fmt.Print("Username: ")
	_, err := fmt.Scan(&username)
	if err != nil {
		return "", "", "", err
	}

	fmt.Print("Email: ")
	_, err = fmt.Scan(&email)
	if err != nil {
		return "", "", "", err
	}

	fmt.Print("Password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()

	if err != nil {
		return "", "", "", err
	}

	err = validation.Validate(username,
		validation.Required,
		validation.Length(3, 20),
	)
	if err != nil {
		return "", "", "", fmt.Errorf("invalid username: %w", err)
	}

	err = validation.Validate(email,
		validation.Required,
		is.Email,
	)
	if err != nil {
		return "", "", "", fmt.Errorf("invalid email: %w", err)
	}

	err = validation.Validate(string(password),
		validation.Required,
		validation.Length(8, 50),
	)
	if err != nil {
		return "", "", "", fmt.Errorf("invalid password: %w", err)
	}

	username64 := base64.StdEncoding.EncodeToString([]byte(username))
	email64 := base64.StdEncoding.EncodeToString([]byte(email))
	password64 := base64.StdEncoding.EncodeToString(password)

	return username64, email64, password64, nil
}
