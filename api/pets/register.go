package pets

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/zSnails/missing-pet-tracker/response"
	"github.com/zSnails/missing-pet-tracker/storage"
)

var log = logrus.WithField("service", "api:pets")

func validatePetData(r *http.Request) (storage.CreateMissingPetParams, error) {
	var empty storage.CreateMissingPetParams
	name := r.FormValue("name")
	if name == "" {
		return empty, errors.New("name field is empty")
	}
	_type := r.FormValue("type")
	if _type == "" {
		return empty, errors.New("type field is empty")
	}
	lastSeen := r.FormValue("last-seen")
	if lastSeen == "" {
		return empty, errors.New("last-seen field is empty")
	}

	return storage.CreateMissingPetParams{
		Name:     name,
		Type:     _type,
		LastSeen: lastSeen,
	}, nil
}

func RegisterUserPet(q *storage.Queries, db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usrData := r.Context().Value("user-data").(storage.CreateUserRow)

		petData, err := validatePetData(r)
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		petData.OwnerID = usrData.ID

		pet, err := q.CreateMissingPet(r.Context(), petData)
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(response.Response[storage.CreateMissingPetRow]{
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
