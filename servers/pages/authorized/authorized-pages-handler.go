package authorized

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacksonopp/skuman/internal/helpers"
	"github.com/jacksonopp/skuman/internal/html"
)

func (s AuthorizedPagesServer) handleValidatePage() {
	s.r.HandleFunc("/validate/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		vc := r.URL.Query().Get("code")

		t, err := html.GetTemplate("validate")
		if err != nil {
			helpers.InternalServerError(w, r, err)
			return
		}

		data := struct {
			ID   string
			Code string
		}{
			id, vc,
		}

		t.Execute(w, data)
	})
}
