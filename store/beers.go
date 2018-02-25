package store

import (
	"context"
	"database/sql"

	"github.com/blockloop/beer_ratings/models"
)

func NewBeers(db *sql.DB) Beers {
	return &beers{
		db: db,
	}
}

type beers struct {
	db *sql.DB
}

func (b *beers) Create(ctx context.Context, beer models.Beer) (*models.Beer, error) {
	panic("not implemented")
}

func (b *beers) Search(context.Context, string) ([]*models.Beer, *models.Pagination, error) {
	panic("not implemented")
}

func (b *beers) ForBrewery(ctx context.Context, breweryID int64) ([]*models.Beer, *models.Pagination, error) {
	panic("not implemented")
}

func (b *beers) Get(ctx context.Context, id int64) (*models.Beer, error) {
	panic("not implemented")
}

func (b *beers) Update(context.Context, models.Beer) error {
	panic("not implemented")
}
