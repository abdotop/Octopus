package limiter

import (
	"net/http"

	"github.com/abdotop/octopus"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"
)

// NewMiddleware creates a new Octopus middleware from a ulule/limiter stdlib middleware
func NewMiddleware(mh *stdlib.Middleware) octopus.HandlerFunc {
	return func(c *octopus.Ctx) {
		handler := mh.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// After the rate limiting middleware has been applied,
			// we transfer control back to Octopus
			c.Values.Set("request", r)
			c.Values.Set("response", w)
			c.Next()
		}))

		// Retrieve the HTTP request from the Octopus context
		r, ok := c.Values.Get("request")
		if !ok {
			c.SendStatus(octopus.StatusInternalServerError)
			return
		}
		request, ok := r.(*http.Request)
		if !ok {
			c.SendStatus(octopus.StatusInternalServerError)
			return
		}

		// Execute the rate limiting middleware
		handler.ServeHTTP(c.Response(), request)
	}
}
