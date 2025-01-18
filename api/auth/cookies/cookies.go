package cookies

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var Store *sessions.CookieStore

func init() {
	godotenv.Load()

	Store = sessions.NewCookieStore(
		[]byte(os.Getenv("COOKIE_AUTH_KEY")),
		[]byte(os.Getenv("COOKIE_ENCRYPTION_KEY")),
	)
}
