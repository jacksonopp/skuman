package csv

import (
	"net/http"

	"github.com/jacksonopp/skuman/internal/html"
	"github.com/jacksonopp/skuman/internal/logger"
	"github.com/jacksonopp/skuman/internal/sku"
)

func (s CsvServer) handleHealthcheck() {
	s.r.Path("/healthcheck").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy"))
	})
}

func (s CsvServer) handleUpload() {
	s.r.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
		items, err := sku.ParseFile(r.FormFile("file"))
		if err != nil {
			logger.Errorln("error parsing file", err)
			http.Error(w, "error parsing file", http.StatusInternalServerError)
			return
		}

		t, err := html.GetComponent("items")
		if err != nil {
			logger.Errorln("error getting items component", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, items)
		if err != nil {
			logger.Errorln("error executing template", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}).Methods("POST")
}
