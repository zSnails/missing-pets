package auth

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/zSnails/missing-pet-tracker/api/auth/cookies"
	"github.com/zSnails/missing-pet-tracker/response"
	"github.com/zSnails/missing-pet-tracker/storage"
	"golang.org/x/crypto/bcrypt"
)

var log = logrus.WithField("service", "api:auth")

func makeSession(w http.ResponseWriter, r *http.Request, user *storage.CreateUserRow) error {
	sess, err := cookies.Store.Get(r, "Session")
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

func Login(q *storage.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.PostForm.Get("email")
		password := r.PostForm.Get("password")

		log.Debugln("Searching for user in database")
		user, err := q.FindUserByEmail(r.Context(), email)
		if err != nil {
			log.Debugf("Error: %s\n", err.Error())
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
		err = makeSession(w, r, &storage.CreateUserRow{
			ID:      user.ID,
			Name:    user.Name,
			Phone:   user.Phone,
			Email:   user.Email,
			Address: user.Address,
		})
		if err != nil {
			log.Errorf("Could not make session: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		err = json.NewEncoder(w).Encode(response.Response[storage.PetOwner]{
			Code: http.StatusOK,
			Data: user,
		})
		if err != nil {
			log.Errorln(err)
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
	})
}
