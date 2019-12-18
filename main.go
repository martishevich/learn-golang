package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Host       string      `json:"host"`
	UserAgent  string      `json:"user_agent"`
	RequestUri string      `json:"request_uri"`
	Header     http.Header `json:"headers"`
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	responseJSON, err := json.Marshal(response{
		r.Host,
		r.Header.Get("User-Agent"),
		r.RequestURI,
		r.Header,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(responseJSON)
}
func main() {
	http.HandleFunc("/", handleRoot)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Panic(err)
	}
}
