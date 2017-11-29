package store

import (
	"errors"
	"strings"
)

var (
	// ErrEmailTaken is an error indicating that the email address is already in use
	ErrEmailTaken = errors.New("email is already in use")

	// ErrUsernameTaken is an error indicating that the username address is already in use
	ErrUsernameTaken = errors.New("username is already in use")
)

func isUniqueConstraint(err error) (yes bool, causeError error) {
	if err == nil {
		return false, nil
	}
	errstr := err.Error()
	if !strings.Contains(errstr, "UNIQUE constraint failed") {
		return false, nil
	}

	column := ""
	for i := len(errstr) - 1; i > 0; i-- {
		r := errstr[i]
		if r == '.' {
			break
		}
		column = string(r) + column
	}
	switch column {
	case "email":
		return true, ErrEmailTaken
	case "username":
		return true, ErrUsernameTaken
	default:
		return true, err
	}
}
