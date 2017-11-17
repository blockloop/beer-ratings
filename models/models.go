package models

// Pagination is used for paginating resultsets
type Pagination struct {
	Page    int
	PerPage int
}

// User is a registered user
type User struct {
	*Address
	ID           uint64 `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	Location     string `json:"location" db:"location"`
}

// Rating is a user rating of a beer
type Rating struct {
	ID     uint64 `json:"id" db:"id"`
	UserID uint64 `json:"user" db:"user_id"`
	Rating int8   `json:"rating" db:"rating"`
	BeerID uint64 `json:"beer" db:"beer_id"`
}

// Beer represents a specific beer
type Beer struct {
	ID        uint64  `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	BreweryID Brewery `json:"brewery" db:"brewery_id"`
	AvgRating int8    `json:"avg_rating" db:"avg_rating"`
}

// Brewery is a brewery
type Brewery struct {
	*Address
	ID   uint64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// Address is a location on the map
type Address struct {
	ID         uint64 `json:"id" db:"id"`
	Line1      string `json:"address_line_1" db:"address_line_1"`
	Line2      string `json:"address_line_2" db:"address_line_2"`
	City       string `json:"city" db:"city"`
	State      string `json:"state" db:"state"`
	PostalCode string `json:"postal_code" db:"postal_code"`
	Country    string `json:"country" db:"country"`
}
