package unauthorized

import (
	"html/template"
	"net/http"

	"github.com/jacksonopp/skuman/internal/helpers"
	"github.com/jacksonopp/skuman/internal/html"
)

func (s UnauthorizedPagesServer) handleLogin() {
	s.r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		files := []string{
			"web/layout/layout.html",
			"web/templates/login.html",
			"web/partials/login-form/login-form.html",
		}
		t, err := template.ParseFiles(files...)

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
