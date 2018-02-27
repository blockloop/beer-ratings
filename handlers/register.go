package handlers

import (
	"net/http"

	"github.com/blockloop/beer_ratings/store"
	"github.com/go-chi/render"
)

type muxer interface {
	Get(string, http.HandlerFunc)
	Post(string, http.HandlerFunc)
	// Method(method string, route string, h http.HandlerFunc)
}

// Register registers the routes with the handlers
func Register(mux muxer, db *store.Stores) {
	mux.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, render.M{"pong": ""})
	})

	// Beers
	mux.Get(`/beers/{id}`, FindBeerHandler(db.Beers))
	mux.Post("/brewery", CreateBreweryHandler(db.Breweries))
	mux.Get("/brewery/{id}", FindBreweryHandler(db.Breweries))
	mux.Post("/brewery/{brewery_id}/beers", CreateBeerHandler(db.Beers, db.Breweries))

	// Users
	mux.Post("/users", CreateUserHandler(db.Users))
	mux.Get("/users/{id}", UserByIDHandler(db.Users))

	// r.MethodFunc(http.MethodGet, "/ping", func(c boar.Context) error {
	// 	return c.WriteJSON(http.StatusOK, boar.JSON{
	// 		"pong": time.Now(),
	// 	})
	// })
}
