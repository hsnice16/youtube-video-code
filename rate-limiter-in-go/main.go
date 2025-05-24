package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rate-limiter-in-go/types"
)

const PORT = "8080"

func main() {
	// GET `/v1/resource`
	http.HandleFunc("/v1/resource", handleResourceRequest)

	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatalf("Error in starting the server: %v", err)
	}
}

func handleResourceRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		fmt.Printf("Request: %#v", r)
		response := types.Response{Success: true, Data: "You will get the resource."}
	default:
		response := types.Response{Success: false, Error: "Only `Get` method supported."}
		encoder := json.NewEncoder(w)
		encoder.Encode(response)
	}
}
