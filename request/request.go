package request

import (
	"context"

	"github.com/blockloop/beer_ratings/models"
)

type key int

const userKey key = iota

// WithUser returns a new context with the given user embedded
func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// User fetches a user from a context
func User(ctx context.Context) (user *models.User) {
	return ctx.Value(userKey).(*models.User)
}
