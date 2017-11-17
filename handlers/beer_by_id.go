package handlers

import (
	"net/http"

	"github.com/apex/log"
	"github.com/blockloop/boar"
	"github.com/blockloop/boar-example/auth"
	"github.com/blockloop/boar-example/store"
)

type beerByID struct {
	db   store.Beers
	log  log.Interface
	auth *auth.User

	URLParams struct {
		BeerID uint64 `url:"id"`
	}
}

func (h *beerByID) Handle(c boar.Context) error {
	beer, err := h.db.Get(c.Context(), h.URLParams.BeerID)
	if err != nil {
		return err
	}
	if beer == nil {
		return boar.ErrEntityNotFound
	}

	return c.WriteJSON(http.StatusOK, beer)
}
