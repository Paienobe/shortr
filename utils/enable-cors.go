package utils

import "net/http"

func EnableCors(w *http.ResponseWriter) {
	// (*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	// (*w).Header().Set("Access-Control-Allow-Credentials", "true")
}
