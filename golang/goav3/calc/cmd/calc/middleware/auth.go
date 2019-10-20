package middleware

import (
	"context"
	"net/http"
	"strings"
)

// WithAuthToken is a HTTP server middleware that reads the value of the
// Authorization header and if present writes it in the request context.
func WithAuthToken() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		// A HTTP handler is a function.
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req := r
			// Grab Authorization header and initialize request context with it.
			if bearerToken := r.Header.Get("Authorization"); bearerToken != "" {
				ctx := context.WithValue(r.Context(), AuthTokenKey, strings.Replace(bearerToken, "Bearer ", "", 1))
				req = r.WithContext(ctx)
			}

			// Call initial handler.
			h.ServeHTTP(w, req)
		})
	}
}
