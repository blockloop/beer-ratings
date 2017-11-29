package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/apex/log"
	"github.com/blockloop/boar"
	"github.com/blockloop/boar-example/store"
	"golang.org/x/crypto/bcrypt"
)

var (
	errPasswordMismatch = errors.New("password and confirm do not match")
)

type createUser struct {
	db  store.Users
	log log.Interface

	Body struct {
		Email           string `json:"email" valid:"email"`
		Password        string `json:"password" valid:"required"`
		PasswordConfirm string `json:"password_confirm" valid:"required"`
	}
}

func (h *createUser) Handle(c boar.Context) error {
	password := []byte(h.Body.Password)
	passwordConfirm := []byte(h.Body.PasswordConfirm)

	if !bytes.Equal(password, passwordConfirm) {
		return boar.NewHTTPError(http.StatusBadRequest, errPasswordMismatch)
	}

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("could not bcrypt password: %v", err)
	}

	user, err := h.db.Create(c.Context(), h.Body.Email, string(hash))
	if err == nil {
		return c.WriteJSON(http.StatusCreated, user)
	}
	if err == store.ErrUsernameTaken || err == store.ErrEmailTaken {
		return boar.NewHTTPError(http.StatusBadRequest, err)
	}
	return err
}
