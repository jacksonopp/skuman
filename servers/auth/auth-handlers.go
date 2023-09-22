package auth

import "net/http"

func (s AuthServer) handleHealthcheck() {
	s.r.Methods("GET").Path("/healthcheck").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy"))
	})
}

func (s AuthServer) handleCreateAccount() {
	s.r.Methods("POST").Path("/account").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("working on it"))
	})
}
