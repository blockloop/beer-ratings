package handlers

import (
	"net/http"

	"github.com/blockloop/beer_ratings/request"
	"github.com/blockloop/beer_ratings/store"
	"github.com/blockloop/tea"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Register registers the routes with the handlers
func Register(mux *chi.Mux, db *store.Stores) {
	mux.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, render.M{"pong": ""})
	})

	// Beers
	mux.Get(`/beers/{id}`, FindBeerHandler(db.Beers))
	mux.Get("/brewery/{id}", FindBreweryHandler(db.Breweries))
	mux.With(authMw(db.Users)).
		Post("/brewery", CreateBreweryHandler(db.Breweries))
	mux.With(authMw(db.Users)).
		Post("/brewery/{brewery_id}/beers", CreateBeerHandler(db.Beers, db.Breweries))

	// Users
	mux.Post("/users", CreateUserHandler(db.Users))
	mux.Get("/users/{id}", UserByIDHandler(db.Users))
}

func authMw(users store.Users) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			userUUID := r.Header.Get("x-user-uuid")
			if userUUID == "" {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			ctx := r.Context()
			user, err := users.LookupByUUID(ctx, userUUID)
			if err != nil {
				tea.Logger(r).WithError(err).Error("failed to lookup user")
				w.WriteHeader(500)
				return
			}

			if user == nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			ctx = request.WithUser(ctx, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
