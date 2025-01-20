package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zSnails/missing-pet-tracker/response"
	"github.com/zSnails/missing-pet-tracker/storage"
)

func Info(q *storage.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userId, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		info, err := q.GetContactInfo(r.Context(), userId)
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(response.Response[storage.GetContactInfoRow]{
			Code: http.StatusOK,
			Data: info,
		})
	})
}
