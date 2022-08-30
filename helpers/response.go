package helpers

import (
	"encoding/json"
	"net/http"
)

// ----------> This function sets header 
func Response(w http.ResponseWriter, code int, message interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}
