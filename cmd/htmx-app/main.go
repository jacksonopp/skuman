package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jacksonopp/skuman/db/db"
	"github.com/jacksonopp/skuman/internal/html"
	"github.com/jacksonopp/skuman/internal/logger"
	"github.com/jacksonopp/skuman/internal/middleware"
	"github.com/jacksonopp/skuman/internal/types"
	"github.com/jacksonopp/skuman/servers/auth"
	"github.com/jacksonopp/skuman/servers/csv"
	"github.com/jacksonopp/skuman/servers/pages/unauthorized"

	_ "github.com/lib/pq"
)

// go:embed schema.sql
var ddl string

func openDb() (*db.Queries, context.Context, error) {
	ctx := context.Background()

	sqldb, err := sql.Open("postgres", "port=5438 user=postgres password=postgres dbname=skuman sslmode=disable")
	if err != nil {
		return nil, nil, err
	}

	if _, err := sqldb.ExecContext(ctx, ddl); err != nil {
		return nil, nil, err
	}

	queries := db.New(sqldb)
	return queries, ctx, nil
}

func main() {
	q, ctx, err := openDb()
	if err != nil {
		logger.Errorln("failed to open db: ", err)
		os.Exit(1)
	}

	r := mux.NewRouter()
	r.Use(middleware.LoggerMiddleware)

	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	servers := []types.Server{}

	unauthorizedPagesRouter := r.PathPrefix("").Subrouter()
	unauthorizedPagesServer := unauthorized.NewUnauthorizedPagesServer(ctx, unauthorizedPagesRouter)
	servers = append(servers, unauthorizedPagesServer)

	csvRouter := r.PathPrefix("/api/csv").Subrouter()
	csvServer := csv.NewCsvRouter(ctx, csvRouter, q)
	servers = append(servers, csvServer)

	authRouter := r.PathPrefix("/api/auth").Subrouter()
	authServer := auth.NewAuthServer(ctx, authRouter, q)
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
