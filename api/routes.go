package api

import (
	"github.com/gorilla/mux"
	"github.com/zSnails/missing-pet-tracker/api/pets"
	"github.com/zSnails/missing-pet-tracker/storage"
)

func Register(r mux.Router, q *storage.Queries) {
	r.Handle("/api/users/{id}/pets", pets.RegisterPet(q)).Methods("POST")
	// TODO: other crud operations
}
