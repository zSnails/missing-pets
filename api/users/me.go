package users

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/zSnails/missing-pet-tracker/response"
	"github.com/zSnails/missing-pet-tracker/storage"
)

var log = logrus.WithField("service", "api:users")

func Myself(q *storage.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userData := r.Context().Value("user-data").(storage.CreateUserRow)
		if err := json.NewEncoder(w).Encode(response.Response[storage.CreateUserRow]{
			Code: http.StatusOK,
			Data: userData,
		}); err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
