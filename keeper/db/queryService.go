package db

import (
	"database/sql"
	"fmt"
)

type QueryService struct {
	db *sql.DB
}

func NewQueryService(db *sql.DB) *QueryService {
	return &QueryService{db: db}
}

func (q *QueryService) GetAllLogEntries() ([]string, error) {

	rws, err := q.db.Query(`SELECT hostname, application_name FROM log_entry`)
	if err != nil {
		return nil, err
	}
	defer rws.Close()

	var logEntryHost string
	var logEntryApp string
	logEntries := []string{}
	for rws.Next() {
		err := rws.Scan(&logEntryHost, &logEntryApp)
		if err != nil {
			return nil, err
		}
		logEntries = append(logEntries, logEntryHost, logEntryApp)
	}

	fmt.Println(logEntries)

	return logEntries, nil
}

func (q *QueryService) GetApplicationEntries() {
}

// Writes on archives not finished compressing
func (q *QueryService) GetPendingWrites() {
}

func (q *QueryService) WriteLogEntry() {
}
