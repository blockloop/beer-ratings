package handlers

import (
	"errors"
	"net/http"

	"github.com/apex/log"
	"github.com/blockloop/boar"
	"github.com/blockloop/boar-example/auth"
	"github.com/blockloop/boar-example/models"
	"github.com/blockloop/boar-example/store"
	"github.com/pborman/uuid"
)

var errBreweryNotFound = errors.New("brewery not found")

type createBeer struct {
	beers     store.Beers
	breweries store.Breweries
	log       log.Interface
	auth      *auth.User

	URLParams struct {
		BreweryID int `url:"brewery_id"`
	}

	Body struct {
		Name    string `json:"name" valid:"required"`
		OwnerID int    `json:"owner_id" valid:"required"`
	}
}

func (h *createBeer) Handle(c boar.Context) error {
	brewery, err := h.breweries.Get(c.Context(), h.URLParams.BreweryID)
	if err != nil {
		return err
	}
	if brewery == nil {
		return boar.NewHTTPError(http.StatusNotFound, errBreweryNotFound)
	}

	beer := models.Beer{
		BreweryID: brewery.ID,
		Name:      h.Body.Name,
		UUID:      uuid.New(),
	}

	newBeer, err := h.beers.Create(c.Context(), beer)
	if err != nil {
		return err
	}

	return c.WriteJSON(http.StatusCreated, newBeer)
}
