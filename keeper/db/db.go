package db

import (
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

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

type LogEntry struct {
	Hostname         string
	ApplicationName  string
	StartTime        string
	EndTime          string
	CalendarWeek     int16
	FileSize         int32
	FileLastModified string
}

type Archive struct {
	Hostname        string
	ApplicationName string
	CalendarWeek    int16
	FileSize        int32
	FinishTime      string
	Completed       bool
}

// creates tables for models
func createTables(db *sql.DB) {
	const createTableLogs string = `
	CREATE TABLE IF NOT EXISTS log_entry(
	id INTEGER NOT NULL PRIMARY KEY,
	hostname STRING NOT NULL,
	application_name STRING NOT NULL,
	start_time DATETIME NOT NULL,
	end_time DATETIME NOT NULL,
	calendar_week INTEGER NOT NULL,
	file_size INTEGER NOT NULL,
	file_last_modified DATETIME NOT NULL
	);`

	const createTableArchive string = `
	CREATE TABLE IF NOT EXISTS archives(
		id INTEGER NOT NULL PRIMARY KEY,
		hostname STRING NOT NULL,
		application_name STRING NOT NULL,
		calendar_week INTEGER NOT NULL,
		file_size INTEGER NOT NULL,
		finish_time DATETIME NOT NULL,
		completed BOOLEAN NOT NULL
	);`

	rslt, err := db.Exec(createTableLogs)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rslt.RowsAffected())

	rslt, err = db.Exec(createTableArchive)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rslt.RowsAffected())
}

func populateTables(db *sql.DB) {

	// var populateLogs string
	// var populateArchives strings.Builder
	var populateLogs strings.Builder
	for i := 1; i < 11; i++ {

		fmt.Fprint(&populateLogs, "INSERT INTO log_entry ( id, hostname, application_name, start_time, end_time, calendar_week, file_size, file_last_modified) VALUES (")
		fmt.Fprint(&populateLogs, "?, ?, ?, ?, ?, ?, ?, ?);")
		// fmt.Fprint(&populateLogs, ")")

		fmt.Println(populateLogs.String())

		// fmt.Fprint(&populateArchives, "INSERT INTO archives ( id, hostname, application_name, calendar_week, file_size, finish_time, completed) VALUES (")
		// fmt.Fprintf(&populateArchives, "%d, hostname_%d, app-%d, %d, %d, %s, %d", i, i+i, i, i, i, time.DateTime, 1)
		// fmt.Fprint(&populateArchives, ");")

		// fmt.Println(populateArchives.String())

		_, err := db.Exec(populateLogs.String(), i, "hostname"+string(i), "app"+string(i), time.DateTime, time.DateTime, i, i, time.DateTime, 1)
		if err != nil {
			fmt.Println(err)
		}
		// _, err = db.Exec(populateArchives.String())
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// populateArchives.Reset()
		populateLogs.Reset()
	}
	fmt.Println("added 1000 entries for each table")
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
	createTables(db)
	populateTables(db)
	rotateDbFiles()
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func rotateDbFiles() {
	// but how can i dumb the db file!?
	cmd := exec.Command("sqlite3", "test.db", fmt.Sprintf(".backup '%s'", "backup.db"))
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("backup cmd failed ", err)
	}
}
