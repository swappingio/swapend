package validation

import "fmt"

func ValidateEmail(string email) error {
	err := emailx.Validate(email)
	if err != nil {
		return fmt.Errorf("Email is not valid.")
	}

	email = emailx.Normalize(email)

	return email
}
