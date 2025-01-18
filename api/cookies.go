package api

import (
	"os"

	"github.com/gorilla/sessions"
)

func MakeStore() *sessions.CookieStore {
	return sessions.NewCookieStore(
		[]byte(os.Getenv("COOKIE_AUTH_KEY")),
		[]byte(os.Getenv("COOKIE_ENCRYPTION_KEY")),
	)
}
