package unauthorized

import (
	"net/http"

	"github.com/jacksonopp/htmx-app/internal/html"
	"github.com/jacksonopp/htmx-app/internal/logger"
)

func (s UnauthorizedPagesServer) handleLogin() {
	s.r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		t, err := html.GetTemplate("login")
		if err != nil {
			logger.Errorln("error getting login template", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)

	})
}
