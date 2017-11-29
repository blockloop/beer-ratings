package store

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/blockloop/boar-example/models"
	"github.com/blockloop/scan"
	"github.com/pborman/uuid"
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
	res, err := sq.Insert("brewery_id, name, uuid, avg_rating").
		Into("beers").
		Values(beer.BreweryID, beer.Name, uuid.New(), beer.AvgRating).
		RunWith(b.db).
		ExecContext(ctx)

	if err != nil {
		return nil, fmt.Errorf("could not execute insert query: %v", err)
	}

	bigid, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("could not determine last inserted id: %v", err)
	}

	beer.ID = int(bigid)
	return &beer, nil
}

func (b *beers) Search(context.Context, string) ([]*models.Beer, *models.Pagination, error) {
	panic("not implemented")
}

func (b *beers) ForBrewery(ctx context.Context, breweryID int) ([]*models.Beer, *models.Pagination, error) {
	panic("not implemented")
}

func (b *beers) Get(ctx context.Context, id int) (*models.Beer, error) {
	rows, err := sq.Select("id, name, brewery_id").
		From("beers").
		Where("id = ?", id).
		RunWith(b.db).
		QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}

	var beer models.Beer
	if err := scan.Row(&beer, rows); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("scan row failed for selecting beer id %d: %v", id, err)
	}

	return &beer, nil
}

func (b *beers) Update(context.Context, models.Beer) error {
	panic("not implemented")
}
