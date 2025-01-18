package main

import (
	"database/sql"
	"encoding/gob"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/zSnails/missing-pet-tracker/api"
	"github.com/zSnails/missing-pet-tracker/storage"
)

var reg = regexp.MustCompile(`[^a-zA-Z0-9áéíóúÁÉÍÓÚ]`)

func removeSpecialCharacters(input string) string {
	return reg.ReplaceAllString(input, "")
}

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Errorln(err)
	}
	gob.Register(storage.CreateUserRow{})
}

var log = logrus.WithField("service", "entry")

func main() {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	sql.Register("sqlite_custom", &sqlite3.SQLiteDriver{
		Extensions: []string{},
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			return conn.RegisterFunc("remove_special_characters", removeSpecialCharacters, true)
		},
	})
	db, err := sql.Open("sqlite_custom", "data.db")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	api.Register(router, storage.New(db), db)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Error(err)
	}
}
