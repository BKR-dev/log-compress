package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

func connectToDatabase(dbFilename string) (*sql.DB, error) {
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
	db, err := connectToDatabase(dbFilename)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
