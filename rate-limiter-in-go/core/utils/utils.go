package utils

import (
	"encoding/json"
	"net/http"
	"rate-limiter-in-go/types"
	"strings"
)

func WriteResponse(w http.ResponseWriter, response types.Response) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func GetUserIP(remoteAddr string) string {
	ip := strings.Split(remoteAddr, ":")[0]
	return ip
}
