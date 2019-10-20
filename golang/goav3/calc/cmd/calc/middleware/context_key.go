package middleware

type (
	// private type used to define context keys
	ctxKey int
)

const (
	// AuthTokenKey is the request context key used to store the request ID
	// created by the RequestID middleware.
	AuthTokenKey ctxKey = iota + 1
)
