package store

import (
	"context"

	"github.com/blockloop/boar-example/models"
)

// Users is a datastore that stores User information
type Users interface {
	LookupByEmail(ctx context.Context, email string) (*models.User, error)
	LookupByID(ctx context.Context, id int) (*models.User, error)
	LookupByEmailAndPassword(ctx context.Context, email, passwordHash string) (*models.User, error)
	Create(ctx context.Context, email, passwordHash string) (*models.User, error)
}

// Ratings is a datastore that stores beer ratings
type Ratings interface {
	Upsert(context.Context, models.User, models.Rating) (*models.Rating, error)
	Find(context.Context, models.User, models.Beer) (*models.Rating, error)
	Delete(context.Context, models.User, models.Rating) (*models.Rating, error)
}

// Beers is a datastore that holds beers
type Beers interface {
	Create(context.Context, models.Beer) error
	Search(context.Context, string) ([]*models.Beer, *models.Pagination, error)
	ForBrewery(ctx context.Context, breweryID int) ([]*models.Beer, *models.Pagination, error)
	Get(ctx context.Context, id int) (*models.Beer, error)
	Update(context.Context, models.Beer) error
}

// Breweries is a datastore that holds beers
type Breweries interface {
	Create(context.Context, models.Brewery) error
	Search(context.Context, string) ([]*models.Brewery, *models.Pagination, error)
	ForBeer(ctx context.Context, beerID int) (*models.Brewery, error)
	Get(ctx context.Context, id int) (*models.Brewery, error)
	Update(context.Context, models.Brewery) error
}
