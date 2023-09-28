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
	tx  *db.TxStore
}

func NewAuthServer(ctx context.Context, r *mux.Router, q *db.Queries, tx *db.TxStore) *AuthServer {
	return &AuthServer{ctx, r, q, tx}
}

func (s AuthServer) Run() {
	s.handleHealthcheck()
	s.handleCreateAccount()
	s.handleAccountVerification()
	s.handleLogin()
}
