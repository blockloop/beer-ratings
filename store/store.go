package store

import (
	"database/sql"
	"errors"
	"strings"
)

var (
	// ErrEmailTaken is an error indicating that the email address is already in use
	ErrEmailTaken = errors.New("email is already in use")

	// ErrUsernameTaken is an error indicating that the username address is already in use
	ErrUsernameTaken = errors.New("username is already in use")
)

// Stores is a group of datastores
type Stores struct {
	Beers     Beers
	Breweries Breweries
	Users     Users
}

// NewStores creates a new group of datastores with the provided
// database connection
func NewStores(db *sql.DB) *Stores {
	return &Stores{
		Beers:     NewBeers(db),
		Breweries: NewBreweries(db),
		Users:     NewUsers(db),
	}
}

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
