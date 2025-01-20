package images

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zSnails/missing-pet-tracker/storage"
)

var log = logrus.WithField("service", "api:images")

func Serve(q *storage.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hash := vars["hash"]
		image, err := q.RetrieveImage(r.Context(), hash)
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "image/png")
		wrote, err := w.Write(image.ImageData)
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Debugf("wrote %d bytes\n", wrote)
	})
}
