package unauthorized

import (
	"context"

	"github.com/gorilla/mux"
)

type UnauthorizedPagesServer struct {
	ctx context.Context
	r   *mux.Router
}

func NewUnauthorizedPagesServer(ctx context.Context, r *mux.Router) *UnauthorizedPagesServer {
	return &UnauthorizedPagesServer{ctx, r}
}

func (s UnauthorizedPagesServer) Run() {
	s.handleLogin()
	s.handleSignup()
}
