package auth

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"net/mail"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jacksonopp/skuman/db/db"
	"github.com/jacksonopp/skuman/internal/helpers"
	"github.com/jacksonopp/skuman/internal/html"
	"github.com/jacksonopp/skuman/internal/logger"
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

func (s AuthServer) handleAccountVerification() {
	s.r.Methods("PATCH", "GET").Path("/validate/{id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if err := r.ParseForm(); err != nil {
			// TODO: create form
			logger.Errorln("error parsing form")
			return
		}

		id := vars["id"]
		vc := r.FormValue("code")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			//  TODO handle error
			return
		}

		_, err = s.q.SetVerification(s.ctx, db.SetVerificationParams{
			VerificationCode: helpers.NullString(vc),
			ID:               int64(idInt),
		})

		if err != nil {
			file := "web/partials/verify-form/verify-error.html"
			t, err := template.ParseFiles(file)
			if err != nil {
				helpers.InternalServerError(w, r, err)
				return
			}

			data := struct {
				Message string
			}{
				"Please check that you have the correct verification code.",
			}

			logger.Warningf("user id [%s] and vericiation code [%s] did not match", id, vc)

			t.Execute(w, data)

			return
		}

		w.Write([]byte(vc + id))
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
