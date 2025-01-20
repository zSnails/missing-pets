package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"github.com/zSnails/missing-pet-tracker/storage"
)

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(logrus.Fields{
			"remote":  r.RemoteAddr,
			"path":    r.URL.Path,
			"referer": r.Referer(),
			"method":  r.Method,
		}).Infoln()
		h.ServeHTTP(w, r)
	})
}

func authenticated(c *sessions.CookieStore) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sess, err := c.Get(r, "Session")
			if err != nil {
				log.Errorln(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if sess.IsNew {
				log.Errorf("Unauthorized access attempt to %s\n", r.URL.Path)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			usrData, ok := sess.Values["user-data"].(storage.CreateUserRow)
			if !ok {
				log.Errorf("Unauthorized access attempt to %s\n", r.URL.Path)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user-data", usrData)))
		})
	}
}
