package csv

import (
	"github.com/gorilla/mux"
	"github.com/jacksonopp/htmx-app/internal/logger"
)

type CsvServer struct {
	r *mux.Router
}

func NewCsvRouter(r *mux.Router) *CsvServer {
	return &CsvServer{r}
}

func (s CsvServer) Run() {
	logger.Infoln("running csv server")
	s.handleHealthcheck()
	s.handleUpload()
}
