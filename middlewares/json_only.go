package middlewares

import (
	"strings"

	"github.com/apex/log"
	"github.com/blockloop/boar"
)

// JSONOnly denies requests with media types other than JSON
func JSONOnly(next boar.HandlerFunc) boar.HandlerFunc {
	contentType := "content-type"
	applicationJSON := "application/json"

	return func(c boar.Context) error {
		log.WithFields(log.Fields{
			"method":       c.Request().Method,
			"content-type": c.Request().Header.Get(contentType),
		}).Info("JSONOnly")

		if c.Request().Method[0] != 'P' ||
			strings.Contains(c.Request().Header.Get(contentType), applicationJSON) {
			return next(c)
		}

		return boar.ErrUnsupportedMediaType
	}
}
