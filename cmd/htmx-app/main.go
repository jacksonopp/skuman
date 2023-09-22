package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacksonopp/skuman/internal/html"
	"github.com/jacksonopp/skuman/internal/logger"
	"github.com/jacksonopp/skuman/internal/middleware"
	"github.com/jacksonopp/skuman/internal/types"
	"github.com/jacksonopp/skuman/servers/auth"
	"github.com/jacksonopp/skuman/servers/csv"
	"github.com/jacksonopp/skuman/servers/pages/unauthorized"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	r := mux.NewRouter()
	r.Use(middleware.LoggerMiddleware)

	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	servers := []types.Server{}

	unauthorizedPagesRouter := r.PathPrefix("").Subrouter()
	unauthorizedPagesServer := unauthorized.NewUnauthorizedPagesServer(ctx, unauthorizedPagesRouter)
	servers = append(servers, unauthorizedPagesServer)

	csvRouter := r.PathPrefix("/api/csv").Subrouter()
	csvServer := csv.NewCsvRouter(csvRouter)
	servers = append(servers, csvServer)

	authRouter := r.PathPrefix("/api/auth").Subrouter()
	authServer := auth.NewAuthServer(ctx, authRouter)
	servers = append(servers, authServer)

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

	http.ListenAndServe(":3000", r)
}
