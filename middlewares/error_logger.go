package middlewares

import (
	"github.com/apex/log"
	"github.com/blockloop/boar"
)

// ErrorLogger logs errors occuring in HTTP handlers
func ErrorLogger(ll log.Interface) boar.Middleware {
	return func(next boar.HandlerFunc) boar.HandlerFunc {
		return func(c boar.Context) error {
			err := next(c)

			if err == nil {
				return nil
			}
			if httperr, ok := err.(boar.HTTPError); ok {
				if httperr.Status() < 499 {
					return nil
				}
			}

			ll.WithError(err).Error("an error occurred")

			return err
		}
	}
}
