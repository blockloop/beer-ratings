package models

import "time"

// Pagination is used for paginating resultsets
type Pagination struct {
	Page    int64
	PerPage int64
}

// User is a registered user
type User struct {
	Address      `db:"-"`
	ID           int64     `json:"id" db:"id"`
	UUID         string    `json:"uuid" db:"uuid"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Created      time.Time `json:"created" db:"created"`
	Modified     time.Time `json:"modified" db:"modified"`
}

// Rating is a user rating of a beer
type Rating struct {
	ID       int64     `json:"id" db:"id"`
	UUID     string    `json:"uuid" db:"uuid"`
	UserID   int64     `json:"user" db:"user_id"`
	Rating   int8      `json:"rating" db:"rating"`
	BeerID   int64     `json:"beer" db:"beer_id"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
}

// Beer represents a specific beer
type Beer struct {
	ID        int64     `json:"id" db:"id"`
	UUID      string    `json:"uuid" db:"uuid"`
	Name      string    `json:"name" db:"name"`
	BreweryID int64     `json:"brewery_id" db:"brewery_id"`
	AvgRating int8      `json:"avg_rating,omitempty" db:"avg_rating"`
	Created   time.Time `json:"created" db:"created"`
	Modified  time.Time `json:"modified" db:"modified"`
}

// Brewery is a brewery
type Brewery struct {
	Address  `db:"-"`
	ID       int64     `json:"id" db:"id"`
	OwnerID  int64     `json:"owner_id" db:"owner_id"`
	UUID     string    `json:"uuid" db:"uuid"`
	Name     string    `json:"name" db:"name"`
	Verified bool      `json:"verified" db:"verified"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
}

// Address is a location on the map
type Address struct {
	Line1      string    `json:"address_line_1,omitempty" db:"address_line_1"`
	Line2      string    `json:"address_line_2,omitempty" db:"address_line_2"`
	City       string    `json:"city,omitempty" db:"city"`
	State      string    `json:"state,omitempty" db:"state"`
	PostalCode string    `json:"postal_code,omitempty" db:"postal_code"`
	Country    string    `json:"country,omitempty" db:"country"`
	Created    time.Time `json:"created" db:"created"`
	Modified   time.Time `json:"modified" db:"modified"`
}
