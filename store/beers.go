package store

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/blockloop/boar-example/models"
)

func NewBeers(db *sql.DB) Beers {
	return &beers{
		db: db,
	}
}

type beers struct {
	db *sql.DB
}

func (b *beers) Create(context.Context, models.Beer) error {
	panic("not implemented")
}

func (b *beers) Search(context.Context, string) ([]*models.Beer, *models.Pagination, error) {
	panic("not implemented")
}

func (b *beers) ForBrewery(ctx context.Context, breweryID int) ([]*models.Beer, *models.Pagination, error) {
	panic("not implemented")
}

func (b *beers) Get(ctx context.Context, id int) (*models.Beer, error) {
	var beer models.Beer
	row := sq.Select("id, name, brewery_id").
		From("beers").
		Where("id = ?", id).
		RunWith(b.db).
		QueryRow()

	if err := row.Scan(&beer.ID, &beer.Name, &beer.BreweryID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("scan row failed for selecting beer id %d: %v", row, err)
	}

	return &beer, nil
}

func (b *beers) Update(context.Context, models.Beer) error {
	panic("not implemented")
}
