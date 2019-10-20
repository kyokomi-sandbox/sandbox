package middleware

import (
	"context"
	"log"

	goa "goa.design/goa/v3/pkg"
)

// ErrorLogger is an endpoint middleware that logs errors using the given
// logger. All log entries start with the given prefix.
func ErrorLogger(l *log.Logger, prefix string) func(goa.Endpoint) goa.Endpoint {
	return func(e goa.Endpoint) goa.Endpoint {
		// A Goa endpoint is itself a function.
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			// Call the original endpoint function.
			res, err := e(ctx, req)
			// Log any error.
			if err != nil {
				l.Printf("[ERROR] %s: %+v", prefix, err)
			}
			// Return endpoint results.
			return res, err
		}
	}
}
