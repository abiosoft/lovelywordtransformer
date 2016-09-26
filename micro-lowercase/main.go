package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/abiosoft/wordtransformer/serviceutil"
)

var (
	portEnvVar = "PORT"
	port       = "80"
)

func init() {
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
}

func main() {
	http.HandleFunc("/", handle)
	log.Println("listening on", port)
	http.ListenAndServe(":"+port, http.HandlerFunc(handle))
}

func handle(w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("word")
	if word == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		serviceutil.LogRequest(w, r, http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, strings.ToLower(word))
	serviceutil.LogRequest(w, r, http.StatusOK)
}
