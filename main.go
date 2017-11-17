package main

import (
	"net/http"
	"time"

	"github.com/apex/log"
	"github.com/blockloop/boar"
	"github.com/blockloop/boar-example/handlers"
	"github.com/blockloop/boar-example/store"
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log := log.WithField("app", "beer-ratings")

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.WithError(err).Fatal("could not connect to datastore")
	}
	db = db.Unsafe()

	factory := handlers.NewFactory(
		store.NewRatings(db),
		store.NewUsers(db),
		store.NewBeers(db),
		store.NewBreweries(db),
		log,
	)

	r := boar.NewRouter()
	r.Get("/beers/:id", factory.GetBeerByID)

	r.MethodFunc(http.MethodGet, "/ping", func(c boar.Context) error {
		return c.WriteJSON(http.StatusOK, boar.JSON{
			"pong": time.Now(),
		})
	})

	addr := ":3000"
	log.WithField("addr", addr).Info("listening")
	log.WithError(r.ListenAndServe(addr)).Info("application exit")
}
