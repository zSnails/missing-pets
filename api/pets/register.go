package pets

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
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
	size := r.FormValue("size")
	if size == "" {
		return empty, errors.New("size field is empty")
	}
	color := r.FormValue("color")
	if color == "" {
		return empty, errors.New("color field is empty")
	}

	return storage.CreateMissingPetParams{
		Name:     name,
		Type:     _type,
		LastSeen: lastSeen,
		Size:     size,
		Color:    color,
	}, nil
}

const MB32 = 32 << 20

func getFormFiles(r *http.Request, fieldName string) ([]*multipart.FileHeader, error) {
	if r.MultipartForm == nil {
		if err := r.ParseMultipartForm(MB32); err != nil {
			return nil, err
		}
	}
	return r.MultipartForm.File[fieldName], nil
}

func RegisterUserPet(h *storage.Queries, db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tx, err := db.BeginTx(r.Context(), nil)
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer tx.Rollback()

		q := h.WithTx(tx)

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

		if r.MultipartForm != nil {
			log.Infoln("Pet has images")
			headers, err := getFormFiles(r, "images")
			if err != nil {
				log.Errorln(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			log.Debugf("Got a total of %d images\n", len(headers))

			for _, header := range headers {
				func() {
					file, err := header.Open()
					if err != nil {
						log.Errorln(err)
						return
					}
					defer file.Close()
					log.Debugf("Registering photo %s\n", header.Filename)

					read, err := io.ReadAll(file)
					if err != nil {
						log.Errorln(err)
						return
					}
					sum := sha256.Sum256(read)
					log.Debugf("Computed hash %x\n", sum)
					if _, err = q.UploadPhoto(r.Context(), storage.UploadPhotoParams{
						PetID:     pet.ID,
						ImageData: read,
						ApiHash:   fmt.Sprintf("%x", sum),
					}); err != nil {
						log.Errorf("could not register %s due to %s\n", header.Filename, err.Error())
					}
				}()
			}
		} else {
			log.Warnln("Pet has no images")
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
		if err = tx.Commit(); err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	})
}
