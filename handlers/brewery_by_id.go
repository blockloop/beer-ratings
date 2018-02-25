package handlers

import (
	"net/http"

	"github.com/blockloop/beer_ratings/store"
	"github.com/blockloop/tea"
)

func FindBreweryHandler(db store.Breweries) http.HandlerFunc {
	return tea.Handler(func(w http.ResponseWriter, r *http.Request) (int, interface{}) {
		id, err := tea.URLInt64(r, "id", "")
		if err != nil {
			return tea.StatusError(404)
		}

		return FindBrewery(id, db, w, r)
	})
}

func FindBrewery(id int64, db store.Breweries, w http.ResponseWriter, r *http.Request) (int, interface{}) {
	brewery, err := db.Get(r.Context(), id)
	if err != nil {
		tea.Logger(r).WithError(err).Error("failed to get brewery from DB")
		return tea.StatusError(500)
	}
	if brewery == nil {
		return tea.StatusError(404)
	}

	return 200, brewery
}
