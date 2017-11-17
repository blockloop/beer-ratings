package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/blockloop/boar-example/models"
	"github.com/jmoiron/sqlx"
)

const schema = `
CREATE TABLE IF NOT EXISTS beers (
	id UNSIGNED BIG INT PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	brewery_id UNSIGNED BIG INT NOT NULL,
	avg_rating TINYINT NOT NULL
)
`

func NewBeers(db *sqlx.DB) Beers {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	db.MustExecContext(ctx, schema)
	return &beers{
		db: db,
	}
}

type beers struct {
	db *sqlx.DB
}

func (b *beers) Create(context.Context, models.Beer) error {
	panic("not implemented")
}

func (b *beers) Search(context.Context, string) ([]*models.Beer, *models.Pagination, error) {
	panic("not implemented")
}

func (b *beers) ForBrewery(ctx context.Context, breweryID uint64) ([]*models.Beer, *models.Pagination, error) {
	panic("not implemented")
}

func (b *beers) Get(ctx context.Context, id uint64) (*models.Beer, error) {
	var beer models.Beer
	query := `SELECT * FROM beers WHERE id = $1`
	err := b.db.GetContext(ctx, &beer, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		err = fmt.Errorf("cannot find beer id %d: %v", id, err)
	}
	return &beer, err
}

func (b *beers) Update(context.Context, models.Beer) error {
	panic("not implemented")
}
