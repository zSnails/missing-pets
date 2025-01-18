package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zSnails/missing-pet-tracker/api/auth"
	"github.com/zSnails/missing-pet-tracker/api/auth/cookies"
	"github.com/zSnails/missing-pet-tracker/api/pets"
	"github.com/zSnails/missing-pet-tracker/storage"
)

var log = logrus.WithField("service", "api")

func Register(r *mux.Router, q *storage.Queries, db *sql.DB) {
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Errorf("route not found %s\n", r.URL.Path)
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	r.Handle("/api/auth/login", auth.Login(q)).Methods("POST")
	r.Handle("/api/auth/login", auth.Logout(q)).Methods("DELETE")
	r.Handle("/api/auth/register", auth.Register(q)).Methods("POST")

	{
		api := r.PathPrefix("/api").Subrouter()
		api.Use(authenticated)
		api.Handle("/me/pets", pets.RegisterUserPet(q, db)).Methods("POST")
		api.Handle("/me/pets", pets.ListMyPets(q)).Methods("GET")
		api.Handle("/me/pets/{petId}", pets.GetPet(q)).Methods("GET")
		api.Handle("/users/{id}/pets", pets.ListUserPets(q)).Methods("GET")
        api.Handle("/me/pets/{petId}", pets.RemoveUserPet(q)).Methods("DELETE")
	}

}

func authenticated(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, err := cookies.Store.Get(r, "Session")
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if sess.IsNew {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		usrData, ok := sess.Values["user-data"].(storage.CreateUserRow)
		if !ok {
			log.Errorln("User is not authenticated")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user-data", usrData)))
	})
}
