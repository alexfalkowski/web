package security

import "github.com/alexfalkowski/go-service/v2/net/http"

const contentSecurityPolicy = "default-src 'self'; " +
	"script-src 'self' https://cdn.jsdelivr.net; " +
	"style-src 'self' https://cdn.jsdelivr.net; " +
	"img-src 'self'; " +
	"connect-src 'self'; " +
	"font-src 'self'; " +
	"object-src 'none'; " +
	"base-uri 'self'; " +
	"frame-ancestors 'none'"

// NewHandlers returns the HTTP middleware handlers used by the site.
func NewHandlers() []http.ChainedHandler {
	return []http.ChainedHandler{Headers{}}
}

// Headers adds browser security headers to HTTP responses.
type Headers struct{}

// ServeHTTP implements negroni.Handler.
func (Headers) ServeHTTP(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	header := res.Header()
	header.Set("Content-Security-Policy", contentSecurityPolicy)
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("Referrer-Policy", "strict-origin-when-cross-origin")
	header.Set("X-Frame-Options", "DENY")
	header.Set("Permissions-Policy", "camera=(), geolocation=(), microphone=()")

	next(res, req)
}
