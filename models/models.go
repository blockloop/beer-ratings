package models

// Pagination is used for paginating resultsets
type Pagination struct {
	Page    int
	PerPage int
}

// User is a registered user
type User struct {
	*Address
	ID           int    `json:"id"`
	UUID         string `json:"uuid"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

// Rating is a user rating of a beer
type Rating struct {
	ID     int    `json:"id"`
	UUID   string `json:"uuid"`
	UserID int    `json:"user"`
	Rating int8   `json:"rating"`
	BeerID int    `json:"beer"`
}

// Beer represents a specific beer
type Beer struct {
	ID        int     `json:"id"`
	UUID      string  `json:"uuid"`
	Name      string  `json:"name"`
	BreweryID Brewery `json:"brewery"`
	AvgRating int8    `json:"avg_rating,omitempty"`
}

// Brewery is a brewery
type Brewery struct {
	*Address
	ID   int    `json:"id"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

// Address is a location on the map
type Address struct {
	Line1      string `json:"address_line_1,omitempty"`
	Line2      string `json:"address_line_2,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	Country    string `json:"country,omitempty"`
}
