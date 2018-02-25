package handlers

import (
	"bytes"
	"net/http"

	"github.com/apex/log"
	"github.com/blockloop/beer_ratings/store"
	"github.com/blockloop/tea"
	"golang.org/x/crypto/bcrypt"
)

var (
	passwordMismatch = "password and confirm do not match"
)

type createUser struct {
	db  store.Users
	log log.Interface
}

type createUserRequest struct {
	Email           string `json:"email" valid:"email"`
	Password        string `json:"password" valid:"required"`
	PasswordConfirm string `json:"password_confirm" valid:"required"`
}

func CreateUserHandler(users store.Users) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) (int, interface{}) {
		var req createUserRequest
		if err := tea.Body(r, &req); err != nil {
			return tea.Error(400, err.Error())
		}

		password := []byte(req.Password)
		passwordConfirm := []byte(req.PasswordConfirm)

		if !bytes.Equal(password, passwordConfirm) {
			return tea.Error(400, passwordMismatch)
		}

		hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			tea.Logger(r).WithError(err).Error("failed to bcrypt password")
			return tea.StatusError(500)
		}

		user, err := users.Create(r.Context(), req.Email, string(hash))
		if err == nil {
			return http.StatusCreated, user
		}
		if err == store.ErrUsernameTaken || err == store.ErrEmailTaken {
			return tea.Error(400, err.Error())
		}
		tea.Logger(r).WithError(err).Error("failed to create user")
		return tea.StatusError(500)
	}
	return tea.Handler(fn)
}
