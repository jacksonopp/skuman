package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacksonopp/htmx-app/internal/logger"
	"github.com/jacksonopp/htmx-app/internal/middleware"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.LoggerMiddleware)

	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("running"))
	})

	logger.Infoln("listening on port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		logger.Errorln("failed to start server")
	}
}
