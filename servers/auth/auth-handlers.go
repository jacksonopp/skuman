package auth

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/mail"
	"strings"

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

		vc := createVerificationCode()

		user, err := s.q.CreateUser(s.ctx, db.CreateUserParams{
			Email:            email,
			PasswordHash:     hash,
			VerificationCode: helpers.NullString(vc),
		})
		if err != nil {
			fmt.Println(err)

			dupeEmail := strings.Contains(err.Error(), "duplicate key")

			var resp signupError

			if dupeEmail {
				resp = signupError{
					Ok:      false,
					Message: "A user with that email already exists",
				}
			} else {
				resp = signupError{
					Ok:      false,
					Message: "Something went wrong creating your user account. Please try again",
				}
			}

			t, err := html.GetPartial("signup-form")
			if err != nil {
				helpers.InternalServerError(w, r, err)
				return
			}
			t.Execute(w, resp)
			return
		}

		redirectPath := fmt.Sprintf("/validate/%d", user.ID)

		w.Header().Add("HX-Redirect", redirectPath)
		w.WriteHeader(http.StatusNoContent)
	})
}

func createVerificationCode() string {
	abc := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	vcSlc := []rune{}
	for i := 0; i < 6; i++ {
		idx := rand.Intn(len(abc))
		vcSlc = append(vcSlc, abc[idx])
	}
	return string(vcSlc)
}

func (s AuthServer) handleAccountValidation() {
	s.r.Methods("PATCH", "GET").Path("/validate/{id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if err := r.ParseForm(); err != nil {
			// TODO: create form
			return
		}

		id := vars["id"]
		vc := r.FormValue("validationCode")

		w.Write([]byte(vc + id))
	})
}
