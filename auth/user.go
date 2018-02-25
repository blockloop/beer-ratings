package auth

var nobody = User{}

// User is user authentication information for requests
type User struct {
	ID int64 `json:"id"`
}

// Nobody tells if this User is not a real user
func (user User) Nobody() bool {
	if user.ID == 0 {
		return true
	}
	return false
}
