package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)

type user struct {
	Name    string
	Address string
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/page.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}
	t.Execute(w, nil)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	expire := time.Now().AddDate(0, 0, 1)
	userObject := user{
		r.FormValue("name"),
		r.FormValue("address"),
	}
	cookie := http.Cookie{
		Name:    "token",
		Value:   userObject.Name + ":" + userObject.Address,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", 302)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleGet).Methods(http.MethodGet)
	router.HandleFunc("/", handlePost).Methods(http.MethodPost)

	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Panic(err)
	}
}
