package db

import (
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"local/model"
	"local/util"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// "modernc.org/sqlite"
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

func connectToDatabase(dbFile string) (*gorm.DB, error) {
	// connect to db

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
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

func createTablesWithGorm(db *gorm.DB, LogEntry *model.LogEntry, Archive *model.Archive) {
	db.AutoMigrate(&LogEntry, &Archive)
}

func populateTables(db *sql.DB) {

	var populateArchives strings.Builder
	var populateLogs strings.Builder
	for i := 1; i < 1001; i++ {

		fmt.Fprint(&populateLogs, "INSERT INTO log_entry ( id, hostname, application_name, start_time, end_time, calendar_week, file_size, file_last_modified) VALUES (")
		fmt.Fprint(&populateLogs, "?, ?, ?, ?, ?, ?, ?, ?);")

		fmt.Fprint(&populateArchives, "INSERT INTO archives ( id, hostname, application_name, calendar_week, file_size, finish_time, completed) VALUES (")
		fmt.Fprint(&populateArchives, "?, ?, ?, ?, ?, ?, ?);")

		_, err := db.Exec(populateLogs.String(), i, "hostname"+string(i), "app"+string(i), time.DateTime, time.DateTime, i, i, time.DateTime, 1)
		if err != nil {
			fmt.Println(err)
		}
		_, err = db.Exec(populateArchives.String(), i, i+i, i, i, i, time.DateTime, 1)
		if err != nil {
			fmt.Println(err)
		}
		populateArchives.Reset()
		populateLogs.Reset()
	}
	fmt.Println("added 1000 entries for each table")
}

func populateTablesWithGorm(db *gorm.DB, LogEntry *model.LogEntry, Archive *model.Archive) {
	var logEntires []model.LogEntry
	var archives []model.Archive

	for i := 1; i < 1001; i++ {

		entrie := model.LogEntry{
			Hostname:         "Hostname_" + strconv.Itoa(i),
			ApplicationName:  "AppName_" + strconv.Itoa(i),
			StartTime:        time.DateTime,
			EndTime:          time.DateTime,
			CalendarWeek:     util.RndmCW(),
			FileSize:         int32((i * 1_073_741_824) / 10),
			FileLastModified: time.DateTime,
		}
		logEntires = append(logEntires, entrie)

		archy := model.Archive{
			Hostname:        "Hostname_" + strconv.Itoa(i),
			ApplicationName: "AppName_" + strconv.Itoa(i),
			CalendarWeek:    util.RndmCW(),
			FileSize:        int32((i * 1_073_741_824) / 10),
			FinishTime:      time.DateTime,
			Completed:       util.RndmB(),
		}
		archives = append(archives, archy)
	}
	db.CreateInBatches(archives, 100)
	db.CreateInBatches(logEntires, 100)
}

func PrintSQLVersion(db *gorm.DB) {
	// get SQLite version
	var dbVersion string
	db.Raw("select sqlite_version()").Scan(&dbVersion)
	fmt.Println("SQLite Version: ", dbVersion)
}

func GetDB() (*gorm.DB, error) {
	dbFile := createDbFile()
	db, err := connectToDatabase(dbFile)
	createTablesWithGorm(db, &model.LogEntry{}, &model.Archive{})
	populateTablesWithGorm(db, &model.LogEntry{}, &model.Archive{})
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
