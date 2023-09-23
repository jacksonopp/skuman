package authorized

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/jacksonopp/skuman/db/db"
)

type AuthorizedPagesServer struct {
	ctx context.Context
	r   *mux.Router
	q   *db.Queries
}

func NewAuthorizedPagesServer(ctx context.Context, r *mux.Router, q *db.Queries) *AuthorizedPagesServer {
	return &AuthorizedPagesServer{ctx, r, q}
}

func (s AuthorizedPagesServer) Run() {
	s.handleValidatePage()
}
