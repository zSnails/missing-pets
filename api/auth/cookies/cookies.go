package cookies

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var Store *sessions.CookieStore

func init() {
	if err := godotenv.Load(); err != nil {
        logrus.WithField("service", "api:auth:cookies").Errorln(err)
	}

	Store = sessions.NewCookieStore(
		[]byte(os.Getenv("COOKIE_AUTH_KEY")),
		[]byte(os.Getenv("COOKIE_ENCRYPTION_KEY")),
	)
}
