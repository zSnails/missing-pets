package api

import (
	"os"

	"github.com/gorilla/sessions"
)

// TODO: store cookies in the database and only send their id
func MakeStore() *sessions.CookieStore {
	return sessions.NewCookieStore(
		[]byte(os.Getenv("COOKIE_AUTH_KEY")),
		[]byte(os.Getenv("COOKIE_ENCRYPTION_KEY")),
	)
}
