package pets

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zSnails/missing-pet-tracker/storage"
)

func RemoveUserPet(q *storage.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		petId, _ := strconv.ParseInt(vars["petId"], 10, 64)
		err := q.RemoveUserPet(r.Context(), petId)
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
            return
		}
	})
}
