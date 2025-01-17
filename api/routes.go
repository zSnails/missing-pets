package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zSnails/missing-pet-tracker/api/auth"
	"github.com/zSnails/missing-pet-tracker/api/pets"
	"github.com/zSnails/missing-pet-tracker/storage"
)

var log = logrus.WithField("service", "api")

func Register(r *mux.Router, q *storage.Queries) {
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Errorf("route not found %s\n", r.URL.Path)
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	r.Handle("/api/users/{id}/pets", pets.RegisterPet(q)).Methods("POST")
	// TODO: other crud operations

	r.Handle("/api/auth/login", auth.Login(q)).Methods("POST")
	r.Handle("/api/auth/register", auth.Register(q)).Methods("POST")
}
