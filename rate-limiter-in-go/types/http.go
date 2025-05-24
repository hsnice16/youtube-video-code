package types

import "net/http"

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type Handler func(w http.ResponseWriter, r *http.Request)
