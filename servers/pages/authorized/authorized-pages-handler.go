package authorized

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacksonopp/skuman/internal/helpers"
)

func (s AuthorizedPagesServer) handleVerifyPage() {
	s.r.HandleFunc("/verify/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		vc := r.URL.Query().Get("code")

		files := []string{
			"web/layout/layout.html",
			"web/templates/verify.html",
			"web/partials/verify-form/verify-form.html",
		}

		t, err := template.ParseFiles(files...)
		if err != nil {
			helpers.InternalServerError(w, r, err)
			return
		}

		data := struct {
			ID           string
			Code         string
			Error        bool
			ErrorMessage string
		}{
			ID:    id,
			Code:  vc,
			Error: false,
		}

		t.Execute(w, data)
	})
}
