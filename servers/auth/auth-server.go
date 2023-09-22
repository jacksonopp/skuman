package auth

import (
	"context"

	"github.com/gorilla/mux"
)

type AuthServer struct {
	ctx context.Context
	r   *mux.Router
}

func NewAuthServer(ctx context.Context, r *mux.Router) *AuthServer {
	return &AuthServer{ctx, r}
}

func (s AuthServer) Run() {
	s.handleHealthcheck()
}
