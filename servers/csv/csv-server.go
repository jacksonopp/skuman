package csv

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/jacksonopp/skuman/db/db"
	"github.com/jacksonopp/skuman/internal/logger"
)

type CsvServer struct {
	ctx context.Context
	r   *mux.Router
	q   *db.Queries
}

func NewCsvRouter(ctx context.Context, r *mux.Router, q *db.Queries) *CsvServer {
	return &CsvServer{ctx, r, q}
}

func (s CsvServer) Run() {
	logger.Infoln("running csv server")
	s.handleHealthcheck()
	s.handleUpload()
}
