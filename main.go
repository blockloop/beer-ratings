package main

import (
	"context"
	"database/sql"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/blockloop/boar"
	"github.com/blockloop/boar-example/handlers"
	"github.com/blockloop/boar-example/store"

	"github.com/apex/log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log := log.WithField("app", "beer-ratings")

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.WithError(err).Fatal("could not connect to datastore")
	}

	schema, err := ioutil.ReadFile("./schema.sql")
	if err != nil {
		log.WithError(err).Fatal("could not read schema file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err = db.ExecContext(ctx, string(schema))
	cancel()
	if err != nil {
		log.WithError(err).Fatal("failed to build schema")
	}

	factory := handlers.NewFactory(
		store.NewRatings(db),
		store.NewUsers(db),
		store.NewBeers(db),
		store.NewBreweries(db),
		log,
	)

	r := boar.NewRouter()
	r.Get("/beers/:id", factory.GetBeerByID)
	r.Post("/users", factory.CreateUser)
	r.Get("/users/:id", factory.GetUserByID)

	r.MethodFunc(http.MethodGet, "/ping", func(c boar.Context) error {
		return c.WriteJSON(http.StatusOK, boar.JSON{
			"pong": time.Now(),
		})
	})

	addr := ":3000"
	log.WithField("addr", addr).Info("listening")
	err = http.ListenAndServe(addr, r.RealRouter())
	log.WithError(err).Info("application exit")
}
