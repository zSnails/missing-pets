package main

import (
	"database/sql"
	"regexp"

	"github.com/mattn/go-sqlite3"
)

var reg = regexp.MustCompile(`[^a-zA-Z0-9áéíóúÁÉÍÓÚ]`)

func removeSpecialCharacters(input string) string {
	return reg.ReplaceAllString(input, "")
}

func main() {
	sql.Register("sqlite_custom", &sqlite3.SQLiteDriver{
		Extensions: []string{},
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			return conn.RegisterFunc("remove_special_characters", removeSpecialCharacters, true)
		},
	})
	db, err := sql.Open("sqlite_custom", "data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
