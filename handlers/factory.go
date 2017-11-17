package handlers

import (
	"github.com/apex/log"
	"github.com/blockloop/boar"
	"github.com/blockloop/boar-example/store"
)

// Factory is a builder for handlers
type Factory struct {
	users     store.Users
	ratings   store.Ratings
	beers     store.Beers
	breweries store.Breweries
	log       log.Interface
}

// NewFactory creates a new factory
func NewFactory(ratings store.Ratings, users store.Users, beers store.Beers, breweries store.Breweries, log log.Interface) *Factory {
	return &Factory{
		users:     users,
		ratings:   ratings,
		beers:     beers,
		breweries: breweries,
		log:       log,
	}
}

// GetBeerByID finds a beer by ID
func (f *Factory) GetBeerByID(c boar.Context) (boar.Handler, error) {
	return &beerByID{
		db:  f.beers,
		log: f.log,
	}, nil
}
