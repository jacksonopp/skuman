package auth

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/mail"

	"github.com/gorilla/mux"
	"github.com/jacksonopp/skuman/db/db"
	"github.com/jacksonopp/skuman/internal/helpers"
	"github.com/jacksonopp/skuman/internal/html"
	"github.com/jacksonopp/skuman/internal/password"
)

func (s AuthServer) handleHealthcheck() {
	s.r.Methods("GET").Path("/healthcheck").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy"))
	})
}

type signupError struct {
	Ok      bool
	Message string
	Email   string
}

func (s AuthServer) handleCreateAccount() {
	s.r.Methods("POST").Path("/account").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.Form["email"][0]
		pw := r.Form["password"][0]
		confirm := r.Form["confirm"][0]

		// validate email
		_, err := mail.ParseAddress(email)
		if err != nil {
			resp := signupError{
				Ok:      false,
				Message: "Your email isn't formatted correctly",
			}
			t, err := html.GetPartial("signup-form")
			if err != nil {
				helpers.InternalServerError(w, r, err)
			}
			t.Execute(w, resp)
			return
		}

		// validate password
		if pw != confirm {
			resp := signupError{
				Ok:      false,
				Message: "Your passwords do not match",
				Email:   email,
			}
			t, err := html.GetPartial("signup-form")
			if err != nil {
				helpers.InternalServerError(w, r, err)
			}
			t.Execute(w, resp)
			return
		}

		hash, err := password.HashPassword(pw)
		if err != nil {
			helpers.InternalServerError(w, r, err)
		}

		abc := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
		rand.Shuffle(len(abc), func(i, j int) {
			abc[i], abc[j] = abc[j], abc[i]
		})
		vc := fmt.Sprintf("%s-%s", string(abc[:4]), string(abc[4:8]))

		_, err = s.q.CreateUser(s.ctx, db.CreateUserParams{
			Email:            email,
			PasswordHash:     hash,
			VerificationCode: helpers.NullString(vc),
		})
		if err != nil {
			resp := signupError{
				Ok:      false,
				Message: "something went wrong creating your user account",
			}
			t, err := html.GetPartial("signup-form")
			if err != nil {
				helpers.InternalServerError(w, r, err)
				return
			}
			t.Execute(w, resp)
			return
		}

		w.Header().Add("HX-Redirect", "/login")
		w.WriteHeader(http.StatusNoContent)
	})
}

func (s AuthServer) handleAccountValidation() {
	s.r.Methods("PATCH", "GET").Path("/validate/{id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_ = vars["id"]
		r.ParseForm()
		vc := r.FormValue("validationCode")

		w.Write([]byte(vc))
	})
}
