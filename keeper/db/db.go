package db

import (
	"database/sql"
	"fmt"
	"log"

	// sql "github.com/mattn/go-sqlite3"
	// sql "modernc.org/sqlite"
	_ "github.com/glebarez/go-sqlite"
)

func connectToDatabase() (*sql.DB, error) {
	// connect

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func PrintSQLVersion(db *sql.DB) {
	// get SQLite version
	sqlVersion := db.QueryRow("select sqlite_version()")
	var dbVersion string
	_ = sqlVersion.Scan(&dbVersion)
	fmt.Println("SQLite Version: ", dbVersion)
}

func GetDB(dbFilename string) (*sql.DB, error) {
	db, err := connectToDatabase()
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func rotateDbFiles(db *sql.DB) {
	// but how can i dumb the db file!?
}
