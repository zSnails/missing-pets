package pets

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zSnails/missing-pet-tracker/response"
	"github.com/zSnails/missing-pet-tracker/storage"
)

func ListMyPets(q *storage.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usrData := r.Context().Value("user-data").(storage.PetOwner)
		listPets(w, r, q, usrData.ID)
	})
}

func ListUserPets(q *storage.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId, _ := strconv.ParseInt(vars["id"], 10, 64)
		_, err := q.FindUserById(r.Context(), userId)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		listPets(w, r, q, userId)
	})
}

func listPets(w http.ResponseWriter, r *http.Request, q *storage.Queries, userID int64) {
	pets, err := q.GetUserPets(r.Context(), userID)
	if err != nil {
		log.Errorln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(response.Response[[]storage.GetUserPetsRow]{
		Code: http.StatusOK,
		Data: pets,
	})
	if err != nil {
		log.Errorln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
