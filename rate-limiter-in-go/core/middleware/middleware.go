package middleware

import (
	"net/http"
	"rate-limiter-in-go/core/utils"
	"rate-limiter-in-go/types"
)

func RateLimiterWrap(handler types.Handler, fn func(string) error) types.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(r.RemoteAddr)

		if err == nil {
			handler(w, r)
		} else {
			response := types.Response{Success: false, Error: err.Error()}
			utils.WriteResponse(w, response)
		}
	}
}
