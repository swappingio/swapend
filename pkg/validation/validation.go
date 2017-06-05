package validation

import (
	"fmt"

	"github.com/goware/emailx"
)

func ValidateEmail(email string) (string, error) {
	err := emailx.Validate(email)
	if err != nil {
		return "", fmt.Errorf("Email is not valid.")
	}

	email = emailx.Normalize(email)

	return email, nil
}

func ValidateUsername(username string) error {
	return nil
}

func ValidatePassword(password string) error {
	return nil
}

func Validate(something string) error {
	return nil
}
