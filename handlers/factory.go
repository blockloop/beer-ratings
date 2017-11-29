package handlers

import (
	"github.com/apex/log"
	"github.com/blockloop/boar"
	"github.com/blockloop/boar-example/auth"
	"github.com/blockloop/boar-example/store"
	"github.com/pborman/uuid"
)

// Factory is a builder for handlers
type Factory struct {
	users     store.Users
	ratings   store.Ratings
	beers     store.Beers
	breweries store.Breweries
	log       log.Interface
}

// NewFactory creates a new factory
func NewFactory(ratings store.Ratings, users store.Users, beers store.Beers, breweries store.Breweries, log log.Interface) *Factory {
	return &Factory{
		users:     users,
		ratings:   ratings,
		beers:     beers,
		breweries: breweries,
		log:       log,
	}
}

// GetBeerByID finds a beer by ID
func (f *Factory) GetBeerByID(c boar.Context) (boar.Handler, error) {
	return &beerByID{
		db:  f.beers,
		log: f.log,
	}, nil
}

func (f *Factory) CreateUser(c boar.Context) (boar.Handler, error) {
	return &createUser{
		db:  f.users,
		log: f.log,
	}, nil
}

func (f *Factory) GetUserByID(c boar.Context) (boar.Handler, error) {
	return &userByID{
		db:  f.users,
		log: f.log,
	}, nil
}

func (f *Factory) CreateBeer(c boar.Context) (boar.Handler, error) {
	auth, err := f.auth(c)
	if err != nil {
		return nil, err
	}
	if auth == nil {
		return nil, boar.ErrUnauthorized
	}

	h := &createBeer{
		beers:     f.beers,
		breweries: f.breweries,
		log:       f.log.WithFields(f.logFields(c, auth)),
		auth:      auth,
	}

	return h, nil
}

func (f *Factory) auth(c boar.Context) (*auth.User, error) {
	u, p, ok := c.Request().BasicAuth()
	if !ok {
		return nil, nil
	}
	user, err := f.users.LookupByEmailAndPassword(c.Context(), u, p)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return &auth.User{ID: user.ID}, nil
}

func (f *Factory) logFields(c boar.Context, user *auth.User) log.Fields {
	fields := log.Fields{}

	rid := c.Request().Header.Get("request-id")
	if len(rid) == 0 {
		rid = uuid.New()
	}
	fields["request.id"] = rid

	if user != nil {
		fields["user.id"] = user.ID
	}
	return fields
}
