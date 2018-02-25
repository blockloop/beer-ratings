package handlers

import (
	"net/http"

	"github.com/apex/log"
	"github.com/blockloop/beer_ratings/auth"
	"github.com/blockloop/beer_ratings/models"
	"github.com/blockloop/beer_ratings/store"
	"github.com/blockloop/tea"
	"github.com/pborman/uuid"
)

type createBeer struct {
	beers     store.Beers
	breweries store.Breweries
	log       log.Interface
	auth      *auth.User

	URLParams struct {
		BreweryID int64 `url:"brewery_id"`
	}
}

type createBeerRequest struct {
	Name    string `json:"name" valid:"required"`
	OwnerID int64  `json:"owner_id" valid:"required"`
}

func CreateBeerHandler(beers store.Beers, breweries store.Breweries) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) (int, interface{}) {
		var req createBeerRequest
		if err := tea.Body(r, req); err != nil {
			return tea.Errorf(400, "bad request: %+v", err)
		}

		breweryID, err := tea.URLInt64(r, "brewery_id", "required")
		if err != nil {
			return tea.Errorf(400, "bad request: brewery_id: %+v", err)
		}

		brewery, err := breweries.Get(r.Context(), breweryID)
		if err != nil {
			tea.Logger(r).WithError(err).Error("failed to get brewery")
			return tea.StatusError(500)
		}
		if brewery == nil {
			return tea.Error(404, "brewery does not exist")
		}

		beer := models.Beer{
			BreweryID: brewery.ID,
			Name:      req.Name,
			UUID:      uuid.New(),
		}

		newBeer, err := beers.Create(r.Context(), beer)
		if err != nil {
			tea.Logger(r).WithError(err).Error("failed to create beer")
			return tea.StatusError(500)
		}

		return http.StatusCreated, newBeer
	}
	return tea.Handler(fn)
}
