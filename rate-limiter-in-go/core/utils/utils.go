package utils

import (
	"encoding/json"
	"net/http"
	"rate-limiter-in-go/types"
)

func WriteResponse(w http.ResponseWriter, response types.Response) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
