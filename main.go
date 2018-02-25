package main

import (
	"context"
	"database/sql"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/apex/log"
	"github.com/blockloop/beer_ratings/handlers"
	"github.com/blockloop/beer_ratings/store"
	"github.com/blockloop/tea"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ll := log.Log
	db, err := connectDB(os.Getenv("DB"))
	if err != nil {
		ll.WithError(err).Fatal("failed to connect to db")
	}

	// always render JSON
	tea.Responder = render.JSON

	mux := chi.NewMux()
	mux.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		tea.LoggerMiddleware,
		middleware.Recoverer,
	)

	mux.Mount("/debug", middleware.Profiler())

	stores := store.NewStores(db)
	handlers.Register(mux, stores)

	// make error handlers write JSON
	mux.NotFound(tea.NotFound)
	mux.MethodNotAllowed(tea.MethodNotAllowed)

	addr := ":3000"
	ll.WithField("addr", addr).Info("listening")

	err = http.ListenAndServe(addr, mux)
	ll.WithError(err).Info("application exit")
}

func connectDB(dsn string) (*sql.DB, error) {
	if len(dsn) == 0 {
		dsn = "file::memory:?mode=memory&cache=shared"
	}

	db, err := sql.Open("sqlite3", dsn)
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
