package pets

import (
	"net/http"

	"github.com/zSnails/missing-pet-tracker/storage"
)

func RegisterPet(q *storage.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
