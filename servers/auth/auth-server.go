package auth

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/jacksonopp/skuman/db/db"
)

type AuthServer struct {
	ctx context.Context
	r   *mux.Router
	q   *db.Queries
}

func NewAuthServer(ctx context.Context, r *mux.Router, q *db.Queries) *AuthServer {
	return &AuthServer{ctx, r, q}
}

func (s AuthServer) Run() {
	s.handleHealthcheck()
	s.handleCreateAccount()
}
