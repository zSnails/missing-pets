package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/zSnails/missing-pet-tracker/response"
	"github.com/zSnails/missing-pet-tracker/storage"
	"golang.org/x/crypto/bcrypt"
)

func validateRegisterParams(r *http.Request) (storage.CreateUserParams, error) {
	var empty storage.CreateUserParams
	username := r.FormValue("username")
	if username == "" {
		return empty, errors.New("username field is empty")
	}
	email := r.FormValue("email")
	if email == "" {
		return empty, errors.New("email field is empty")
	}
	phone := r.FormValue("phone")
	if phone == "" {
		return empty, errors.New("phone field is empty")
	}
	address := r.FormValue("address")
	if address == "" {
		return empty, errors.New("address field is empt")
	}

	return storage.CreateUserParams{
		Name:    username,
		Phone:   phone,
		Email:   email,
		Address: address,
		Hash:    nil,
	}, nil
}

func Register(q *storage.Queries, cookies *sessions.CookieStore) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userData, err := validateRegisterParams(r)
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		password := r.FormValue("password")

		log.Debugln("Generating hash from password")
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Errorf("Could not generate the hash from the password: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Debugln("Hash generated successfully")
		userData.Hash = hash
		log.Debugln("Saving user to the database")
		user, err := q.CreateUser(r.Context(), userData)
		if err != nil {
			log.Errorf("Could not store user: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Debugf("Making session")
		err = makeSession(w, r, cookies, &user)
		if err != nil {
			log.Errorf("Could not make session: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(response.Response[storage.CreateUserRow]{
			Code: http.StatusOK,
			Data: user,
		})
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
