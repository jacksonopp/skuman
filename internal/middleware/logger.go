package middleware

import (
	"net/http"

	"github.com/jacksonopp/htmx-app/internal/logger"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infoln("Incoming request: ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
