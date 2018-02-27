package store

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/blockloop/beer_ratings/models"
	"github.com/blockloop/scan"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

var (
	selectBeersCols = scan.Columns(new(models.Beer))
	selectBeers     = sq.Select(selectBeersCols...).From("beers")

	insertBeersCols = scan.Columns(new(models.Beer), "id")
	insertBeers     = sq.Insert("beers").Columns(insertBeersCols...)
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
	if beer.UUID == "" {
		beer.UUID = uuid.New()
	}

	now := time.Now()
	beer.Created = now
	beer.Modified = now

	vals := scan.Values(insertBeersCols, &beer)
	res, err := insertBeers.RunWith(b.db).
		Values(vals...).
		ExecContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	beer.ID, err = res.LastInsertId()
	return &beer, errors.Wrap(err, "failed to lookup new beer ID")
}

func (b *beers) Search(context.Context, string) ([]*models.Beer, *models.Pagination, error) {
	panic("not implemented")
}

func (b *beers) ForBrewery(ctx context.Context, breweryID int64) ([]*models.Beer, *models.Pagination, error) {
	panic("not implemented")
}

func (b *beers) Get(ctx context.Context, id int64) (*models.Beer, error) {
	rows, err := selectBeers.Where(sq.Eq{"id": id}).
		Limit(1).
		RunWith(b.db).
		QueryContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	var beer models.Beer
	if err := scan.Row(&beer, rows); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to scan beer")
	}

	return &beer, nil
}

func (b *beers) Update(context.Context, models.Beer) error {
	panic("not implemented")
}
