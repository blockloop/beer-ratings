package store

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/apex/log"
	"github.com/blockloop/boar-example/models"
	"github.com/blockloop/scan"
)

func NewBreweries(db *sql.DB, ll log.Interface) Breweries {
	return &breweries{
		db:  db,
		log: ll,
	}
}

type breweries struct {
	db  *sql.DB
	log log.Interface
}

func (b *breweries) Create(context.Context, models.Brewery) (*models.Brewery, error) {
	panic("not implemented")
}

func (b *breweries) Search(context.Context, string) ([]*models.Brewery, *models.Pagination, error) {
	panic("not implemented")
}

func (b *breweries) ForBeer(ctx context.Context, beerID int) (*models.Brewery, error) {
	panic("not implemented")
}

func (b *breweries) Get(ctx context.Context, id int) (brewery *models.Brewery, err error) {
	rows, err := sq.Select("*").
		From("breweries").
		Where("id = ?", id).
		RunWith(b.db).
		QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}

	err = scan.Row(brewery, rows)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return
}

func (b *breweries) Update(context.Context, models.Brewery) error {
	panic("not implemented")
}
