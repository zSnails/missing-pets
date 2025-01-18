package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/zSnails/missing-pet-tracker/storage"
)

func Logout(q *storage.Queries, cookies *sessions.CookieStore) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := cookies.Get(r, "Session")
		if err != nil {
			log.Errorln(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sess, err := cookies.New(r, "Session")
		if err != nil {
			log.Errorln(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sess.Options.Path = "/"

		err = sess.Save(r, w)
		if err != nil {
			log.Errorln(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
