package pets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zSnails/missing-pet-tracker/response"
	"github.com/zSnails/missing-pet-tracker/storage"
)

func GetPet(q *storage.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usrData := r.Context().Value("user-data").(storage.CreateUserRow)
		vars := mux.Vars(r)
		petId, _ := strconv.ParseInt(vars["petId"], 10, 64)

		pet, err := q.GetPetByOwnerAndId(r.Context(), storage.GetPetByOwnerAndIdParams{
			ID:      petId,
			OwnerID: usrData.ID,
		})
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(response.Response[storage.GetPetByOwnerAndIdRow]{
			Code: http.StatusOK,
			Data: pet,
		})
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func ListMyPets(q *storage.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usrData := r.Context().Value("user-data").(storage.CreateUserRow)
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

func getQueryIntDefault(r url.Values, name string, dft int64) int64 {
	if !r.Has(name) {
		return dft
	}

	num, err := strconv.ParseInt(r.Get(name), 10, 64)
	if err != nil {
		log.Errorln(err)
		return dft // BUG: this will probably introduce a bug at some point
	}

	return num
}

func ListAllPetsFilter(q *storage.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("name")
		limit := getQueryIntDefault(r.URL.Query(), "count", 10)
		offset := getQueryIntDefault(r.URL.Query(), "page", 0)

		name := fmt.Sprintf("%%%s%%", query)
		pets, err := q.GetAllPetsNameFilter(r.Context(), storage.GetAllPetsNameFilterParams{
			Name:   name,
			Limit:  limit,
			Offset: offset * limit,
		})
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(response.Response[[]storage.GetAllPetsNameFilterRow]{
			Code: http.StatusOK,
			Data: pets,
		})
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
