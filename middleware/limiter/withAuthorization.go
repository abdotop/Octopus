package limiter

import (
	"errors"
	"net/http"

	"github.com/abdotop/octopus"
	ululeLimiter "github.com/ulule/limiter/v3"
)

type StatusResponses map[octopus.StatusCode]map[string]interface{}

// WithAuthorization creates a middleware for authorization and rate limiting.
// It checks for the Authorization header and applies rate limiting using the provided limiter.
// Custom error responses can be defined using the statusResponses parameter.
func WithAuthorization(ul *ululeLimiter.Limiter, statusResponses StatusResponses) octopus.HandlerFunc {
    return func(c *octopus.Ctx) {
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            sendErrorResponse(c, octopus.StatusUnauthorized, statusResponses)
            return
        }
        if code, err := limitByToken(ul, authHeader, c); err != nil {
            sendErrorResponse(c, code, statusResponses)
            return
        }
        c.Next()
    }
}

func limitByToken(limiter *ululeLimiter.Limiter, key string, c *octopus.Ctx) (octopus.StatusCode, error) {
	r, ok := c.Values.Get("request")
	if !ok {
		return octopus.StatusInternalServerError, errors.New("request not found in context")
	}
	request, ok := r.(*http.Request)
	if !ok {
		return octopus.StatusInternalServerError, errors.New("invalid request type")
	}
	ctx, err := limiter.Get(request.Context(), key)
	if err != nil {
		return octopus.StatusInternalServerError, err
	}
	if ctx.Reached {
		return octopus.StatusTooManyRequests, errors.New("rate limit exceeded")
	}
	return 0, nil
}

func sendErrorResponse(c *octopus.Ctx, statusCode octopus.StatusCode, statusResponses StatusResponses) {
	response, ok := statusResponses[statusCode]
	if !ok {
		response = map[string]interface{}{
			"error":   statusCode,
			"message": "An error occurred",
		}
	}
	c.Status(statusCode).JSON(response)
}
