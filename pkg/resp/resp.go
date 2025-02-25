package resp

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}
