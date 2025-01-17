package auth

import (
	"net/http"

	"github.com/zSnails/missing-pet-tracker/storage"
	"golang.org/x/crypto/bcrypt"
)

func Register(q *storage.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		phone := r.FormValue("phone")
		address := r.FormValue("address")

		log.Debugln("Generating hash from password")
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Debugf("Could not generate the hash from the password: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Debugln("Hash generated successfully")

		log.Debugln("Saving user to the database")
		user, err := q.CreateUser(r.Context(), storage.CreateUserParams{
			Name:    username,
			Phone:   phone,
			Email:   email,
			Address: address,
			Hash:    hash,
		})
		if err != nil {
			log.Debugf("Could not store user: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Debugf("Making session")
		err = makeSession(w, r, &user)
		if err != nil {
			log.Debugf("Could not make session: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
