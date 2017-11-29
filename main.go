package main

import (
	"context"
	"database/sql"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/apex/log"
	"github.com/blockloop/boar"
	"github.com/blockloop/boar-example/handlers"
	"github.com/blockloop/boar-example/middlewares"
	"github.com/blockloop/boar-example/store"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ll := log.Log
	db, err := connectDB()
	if err != nil {
		ll.WithError(err).Fatal("failed to connect to db")
	}

	factory := handlers.NewFactory(
		store.NewRatings(db),
		store.NewUsers(db, ll),
		store.NewBeers(db),
		store.NewBreweries(db, ll),
		ll,
	)

	r := boar.NewRouter()

	r.Use(middlewares.ErrorLogger(ll))
	r.Use(boar.PanicMiddleware)
	r.Use(middlewares.JSONOnly)
	r.Use(middlewares.HTTPLogger(ll))

	/// ROUTES

	// Beers
	r.Get("/beers/:id", factory.GetBeerByID)
	r.Post("/brewery/:brewery_id/beers", factory.CreateBeer)
	// r.Get("/brewery/:brewery_id/beers/:id", factory.GetBeer)

	// Users
	r.Post("/users", factory.CreateUser)
	r.Get("/users/:id", factory.GetUserByID)

	r.MethodFunc(http.MethodGet, "/ping", func(c boar.Context) error {
		return c.WriteJSON(http.StatusOK, boar.JSON{
			"pong": time.Now(),
		})
	})

	/// LISTEN
	addr := ":3000"
	ll.WithField("addr", addr).Info("listening")
	err = http.ListenAndServe(addr, r.RealRouter())
	ll.WithError(err).Info("application exit")
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	schema, err := ioutil.ReadFile("./schema.sql")
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err = db.ExecContext(ctx, string(schema))
	cancel()
	if err != nil {
		return nil, err
	}
	return db, nil
}
