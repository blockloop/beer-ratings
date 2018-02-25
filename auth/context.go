package auth

import (
	"context"
	"net/http"
	"strconv"
)

// FromRequest retrieves a user from an HTTP request
func FromRequest(r *http.Request) User {
	user := FromCtx(r.Context())
	if !user.Nobody() {
		return user
	}

	id, err := strconv.ParseInt(r.Header.Get("x-user-id"), 0, 64)
	if err != nil {
		return nobody
	}
	return User{ID: id}
}

// FromCtx retrieves a user from a context
func FromCtx(ctx context.Context) User {
	if u, ok := ctx.Value("user").(*User); ok {
		return *u
	}
	return nobody
}
