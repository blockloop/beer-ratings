package middlewares

import (
	"time"

	"github.com/apex/log"
	"github.com/blockloop/boar"
)

// HTTPLogger is an HTTP middleware that logs request/response times and statuses
func HTTPLogger(ll log.Interface) boar.Middleware {
	return func(next boar.HandlerFunc) boar.HandlerFunc {
		return func(c boar.Context) error {
			start := time.Now()
			err := next(c)
			dur := time.Since(start)

			ll.WithFields(log.Fields{
				"path":     c.Request().URL.Path,
				"status":   c.Response().Status(),
				"sent":     c.Response().Len(),
				"duration": dur,
			}).Info("HTTP request received")

			return err
		}
	}
}
