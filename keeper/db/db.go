package db

import (
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func createDbFile() string {
	dbFile := "test.db"
	file, err := os.Create(dbFile)
	defer file.Close()

	if errors.Is(err, &fs.PathError{}) {
		fmt.Println("could not create ", dbFile)
	}
	// sad that this wont work
	if errors.Is(err, fs.ErrExist) {
		fmt.Println("db file already exists")
	} else if err != nil {
		fmt.Println("you done goofed up m8, ", err)
	}
	return dbFile
}

func connectToDatabase(dbFile string) (*sql.DB, error) {
	// connect

	db, err := sql.Open("sqlite", dbFile)
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

func GetDB() (*sql.DB, error) {
	dbFile := createDbFile()
	db, err := connectToDatabase(dbFile)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func rotateDbFiles(db *sql.DB) {
	// but how can i dumb the db file!?

}
