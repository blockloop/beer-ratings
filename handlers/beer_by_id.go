package handlers

import (
	"net/http"

	"github.com/blockloop/beer_ratings/store"
	"github.com/blockloop/tea"
)

func FindBeerHandler(db store.Beers) http.HandlerFunc {
	return tea.Handler(func(w http.ResponseWriter, r *http.Request) (int, interface{}) {
		id, err := tea.URLInt64(r, "id", "")
		if err != nil {
			return tea.StatusError(404)
		}

		return FindBeer(id, db, w, r)
	})
}

func FindBeer(id int64, db store.Beers, w http.ResponseWriter, r *http.Request) (int, interface{}) {
	beer, err := db.Get(r.Context(), id)
	if err != nil {
		tea.Logger(r).WithError(err).Error("failed to get beer from DB")
		return tea.StatusError(500)
	}
	if beer == nil {
		return tea.StatusError(404)
	}

	return 200, beer
}
