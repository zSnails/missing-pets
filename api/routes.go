package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zSnails/missing-pet-tracker/api/auth"
	"github.com/zSnails/missing-pet-tracker/api/pets"
	"github.com/zSnails/missing-pet-tracker/storage"
)

var log = logrus.WithField("service", "api")

func Register(r *mux.Router, queries *storage.Queries, db *sql.DB) {
	cookies := MakeStore()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Errorf("route not found %s\n", r.URL.Path)
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	r.Use(logger)

	r.Handle("/api/auth/login", auth.Login(queries, cookies)).Methods("POST")
	r.Handle("/api/auth/login", auth.Logout(queries, cookies)).Methods("DELETE")
	r.Handle("/api/auth/register", auth.Register(queries, cookies)).Methods("POST")
	r.Handle("/api/users/{id}/pets", pets.ListUserPets(queries)).Methods("GET")
	r.Handle("/api/pets", pets.ListAllPetsFilter(queries)).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(authenticated(cookies))
	api.Handle("/me/pets", pets.RegisterUserPet(queries, db)).Methods("POST")
	api.Handle("/me/pets", pets.ListMyPets(queries)).Methods("GET")
	api.Handle("/me/pets/{petId}", pets.GetPet(queries)).Methods("GET")
	api.Handle("/me/pets/{petId}", pets.RemoveUserPet(queries)).Methods("DELETE")
}
