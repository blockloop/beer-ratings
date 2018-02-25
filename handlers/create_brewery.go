package handlers

import (
	"net/http"

	"github.com/blockloop/beer_ratings/auth"
	"github.com/blockloop/beer_ratings/models"
	"github.com/blockloop/beer_ratings/store"
	"github.com/blockloop/tea"
)

type createBreweryRequest struct {
	models.Address
	Name    string `json:"name" validate:"required"`
	OwnerID int64  `json:"owner_id"`
}

// CreateBreweryHandler creates a new http.HandlerFunc which can be used
// to create breweries
func CreateBreweryHandler(db store.Breweries) http.HandlerFunc {
	return tea.Handler(func(w http.ResponseWriter, r *http.Request) (int, interface{}) {
		return CreateBrewery(db, w, r)
	})
}

// CreateBrewery creates a brewery with the provided data store
func CreateBrewery(db store.Breweries, w http.ResponseWriter, r *http.Request) (int, interface{}) {
	user := auth.FromRequest(r)
	if user.Nobody() {
		return tea.StatusError(http.StatusForbidden)
	}

	var req createBreweryRequest
	if err := tea.Body(r, &req); err != nil {
		return tea.Error(http.StatusBadRequest, err.Error())
	}

	brewery, err := db.Create(r.Context(), user.ID, models.Brewery{
		Name:    req.Name,
		Address: req.Address,
		OwnerID: req.OwnerID,
	})
	if err != nil {
		tea.Logger(r).WithError(err).Error("failed to create brewery")
		return tea.StatusError(http.StatusInternalServerError)
	}
	return http.StatusCreated, brewery
}
