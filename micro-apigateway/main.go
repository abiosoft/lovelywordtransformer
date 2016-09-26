package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/abiosoft/wordtransformer/serviceutil"
)

var (
	portEnvVar = "PORT"
	port       = "80"

	// services
	upperCaseServiceEnv = "UPPER_CASE_URL"
	lowerCaseServiceEnv = "LOWER_CASE_URL"
	titleCaseServiceEnv = "TITLE_CASE_URL"
	reverseServiceEnv   = "REVERSE_URL"

	servicesEnv = []string{
		upperCaseServiceEnv,
		lowerCaseServiceEnv,
		titleCaseServiceEnv,
		reverseServiceEnv,
	}

	services = map[string]service{}
)

type service struct {
	url      string
	endpoint string
}

func (s service) endpointURL(word string) string {
	return s.url + "/" + s.endpoint + "?word=" + url.QueryEscape(word)
}

func init() {
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	for _, env := range servicesEnv {
		services[env] = service{
			url:      os.Getenv(env),
			endpoint: strings.ToLower(strings.Split(env, "_")[0]),
		}
	}
}

func main() {
	http.HandleFunc("/upper", handleUppercase)
	log.Println("listening on", port)
	http.ListenAndServe(":"+port, nil)
}

func handleUppercase(w http.ResponseWriter, r *http.Request) {
	service, ok := services[r.URL.Path]
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	status, err := handleService(w, r, service)
	serviceutil.LogRequest(w, r, status)
	if err != nil {
		log.Println("Error:", err)
	}
}

func handleService(w http.ResponseWriter, r *http.Request, s service) (int, error) {
	resp, err := http.Get(s.endpointURL(r.FormValue("word")))
	if err != nil {
		return http.StatusServiceUnavailable, err
	}
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	return resp.StatusCode, err
}
