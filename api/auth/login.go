package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"github.com/zSnails/missing-pet-tracker/response"
	"github.com/zSnails/missing-pet-tracker/storage"
	"golang.org/x/crypto/bcrypt"
)

var log = logrus.WithField("service", "api:auth")

func makeSession(w http.ResponseWriter, r *http.Request, cookies sessions.Store, user *storage.CreateUserRow) error {
	sess, err := cookies.Get(r, "Session")
	if err != nil {
		return err
	}

	sess.Options.HttpOnly = true
	sess.Options.Path = "/"
	sess.Options.MaxAge = 3600
	sess.Options.SameSite = http.SameSiteLaxMode

	sess.Values["user-data"] = user

	err = sess.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

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

func Login(q *storage.Queries, cookies sessions.Store) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		log.Debugln("Searching for user in database")
		user, err := q.FindUserByEmail(r.Context(), email)
		if err != nil {
			log.Errorf("Error: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		log.Debugf("User found under the name %s\n", user.Name)

		log.Debugln("Comparing stored hash with given password")
		err = bcrypt.CompareHashAndPassword(user.Hash, []byte(password))
		if err != nil {
			log.Errorf("Could not compare: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		log.Debugln("Making session")
		sessUser := storage.CreateUserRow{
			ID:      user.ID,
			Name:    user.Name,
			Phone:   user.Phone,
			Email:   user.Email,
			Address: user.Address,
		}
		err = makeSession(w, r, cookies, &sessUser)
		if err != nil {
			log.Errorf("Could not make session: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		err = json.NewEncoder(w).Encode(response.Response[storage.CreateUserRow]{
			Code: http.StatusOK,
			Data: sessUser,
		})
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
	})
}
