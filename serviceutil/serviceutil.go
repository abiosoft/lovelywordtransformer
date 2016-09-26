package serviceutil

import (
	"log"
	"net/http"
)

// LogRequest logs an http request.
func LogRequest(w http.ResponseWriter, r *http.Request, status int) {
	log.Println(r.URL.Path, r.RemoteAddr, status)
}
