package handlers

import (
	"net/http"

	"github.com/apex/log"
	"github.com/blockloop/boar"
	"github.com/blockloop/boar-example/store"
)

type userByID struct {
	db  store.Users
	log log.Interface

	URLParams struct {
		ID int `url:"id" valid:"required"`
	}
}

func (u *userByID) Handle(c boar.Context) error {
	user, err := u.db.LookupByID(c.Context(), u.URLParams.ID)
	if err != nil {
		u.log.WithError(err).WithField("user.id", u.URLParams.ID).
			Error("failed to find user by ID")
		return err
	}

	if user == nil {
		return boar.ErrEntityNotFound
	}

	return c.WriteJSON(http.StatusOK, user)
}
