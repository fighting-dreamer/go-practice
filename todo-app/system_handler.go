package main

import "net/http"

// have the logic for the system handler and also the response.

func SystemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status: "pong"}`))
}
