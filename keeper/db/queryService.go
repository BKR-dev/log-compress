package db

import (
	"local/model"

	"gorm.io/gorm"
)

type QueryService struct {
	db *gorm.DB
}

func NewQueryService(db *gorm.DB) *QueryService {
	return &QueryService{db: db}
}

// func (q *QueryService) GetAllLogEntries() ([]string, error) {

// 	rws, err := q.db.Query(`SELECT hostname, application_name FROM log_entry`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rws.Close()

// 	var logEntryHost string
// 	var logEntryApp string
// 	logEntries := []string{}
// 	for rws.Next() {
// 		err := rws.Scan(&logEntryHost, &logEntryApp)
// 		if err != nil {
// 			return nil, err
// 		}
// 		logEntries = append(logEntries, logEntryHost, logEntryApp)
// 	}

// 	fmt.Println(logEntries)

// 	return logEntries, nil
// }

func (q *QueryService) GetAllLogEntriesWithGorm() []model.LogEntry {
	var result []model.LogEntry
	q.db.Raw(`SELECT * FROM log_entries`).Scan(&result)
	return result
}

func (q *QueryService) GetApplicationEntries() {
}

func (q *QueryService) GetPendingWrites() {
}

func (q *QueryService) WriteLogEntry() {
}
