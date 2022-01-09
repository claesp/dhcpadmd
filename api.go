package main

import (
	"encoding/json"
	"net/http"
)

func apiPing(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(CONFIG)
}
