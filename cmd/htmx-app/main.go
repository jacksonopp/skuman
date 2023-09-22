package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacksonopp/htmx-app/internal/html"
	"github.com/jacksonopp/htmx-app/internal/logger"
	"github.com/jacksonopp/htmx-app/internal/middleware"
	"github.com/jacksonopp/htmx-app/internal/types"
	"github.com/jacksonopp/htmx-app/servers/csv"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.LoggerMiddleware)

	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	servers := []types.Server{}

	csvRouter := r.PathPrefix("/api/csv").Subrouter()
	csvServer := csv.NewCsvRouter(csvRouter)
	servers = append(servers, csvServer)

	for _, server := range servers {
		server.Run()
	}

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := html.GetTemplate("index")
		if err != nil {
			logger.Errorln("error getting login template", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	})

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		t, err := html.GetTemplate("login")
		if err != nil {
			logger.Errorln("error getting login template", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	})

	http.ListenAndServe(":3000", r)
}
