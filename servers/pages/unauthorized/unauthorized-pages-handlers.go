package unauthorized

import (
	"net/http"

	"github.com/jacksonopp/skuman/internal/helpers"
	"github.com/jacksonopp/skuman/internal/html"
)

func (s UnauthorizedPagesServer) handleLogin() {
	s.r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		t, err := html.GetTemplate("login")
		if err != nil {
			helpers.InternalServerError(w, r, err)
			return
		}
		t.Execute(w, nil)
	})
}

func (s UnauthorizedPagesServer) handleSignup() {
	s.r.Methods("GET").Path("/sign-up").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := html.GetTemplate("sign-up")
		if err != nil {
			helpers.InternalServerError(w, r, err)
			return
		}
		t.Execute(w, nil)
	})
}
