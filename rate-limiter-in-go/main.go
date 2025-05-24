package main

import (
	"log"
	"net/http"
	"rate-limiter-in-go/constants"
	"rate-limiter-in-go/core/middleware"
	"rate-limiter-in-go/core/rate_limiter"
	"rate-limiter-in-go/core/utils"
	"rate-limiter-in-go/types"
)

func main() {
	buckets := rate_limiter.TokenBucketInit()

	// Token Bucket rate limited `/v1/resource` request handler
	tokenBucketAlgo := rate_limiter.TokenBucketAlgo(&buckets)
	rateLimitedResourceRequestHandler := middleware.RateLimiterWrap(handleResourceRequest, tokenBucketAlgo)

	// GET `/v1/resource`
	http.HandleFunc("/v1/resource", rateLimitedResourceRequestHandler)

	if err := http.ListenAndServe(":"+constants.PORT, nil); err != nil {
		log.Fatalf("Error in starting the server: %v", err)
	}
}

func handleResourceRequest(w http.ResponseWriter, r *http.Request) {
	var response types.Response

	switch r.Method {
	case http.MethodGet:
		response = types.Response{Success: true, Data: "You will get the resource."}
	default:
		response = types.Response{Success: false, Error: "Only `Get` method supported."}
	}

	utils.WriteResponse(w, response)
}
