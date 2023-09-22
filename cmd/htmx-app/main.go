package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacksonopp/htmx-app/internal/middleware"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.LoggerMiddleware)

	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("web/templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	})

	http.ListenAndServe(":3000", r)
}
