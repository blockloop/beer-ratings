package handlers

import (
	"net/http"

	"github.com/apex/log"
	"github.com/blockloop/beer_ratings/store"
	"github.com/blockloop/tea"
)

type userByID struct {
	db  store.Users
	log log.Interface

	URLParams struct {
		ID int64 `url:"id" valid:"required"`
	}
}

func UserByIDHandler(db store.Users) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) (int, interface{}) {
		userID, err := tea.URLInt64(r, "id", "required")
		if err != nil {
			return tea.Errorf(400, "bad request: %+v", err)
		}

		user, err := db.LookupByID(r.Context(), userID)
		if err != nil {
			tea.Logger(r).WithError(err).WithField("user.id", userID).
				Error("failed to find user by ID")
			return tea.StatusError(500)
		}

		if user == nil {
			return tea.Error(404, "user not found")
		}

		return 200, user
	}
	return tea.Handler(fn)
}
