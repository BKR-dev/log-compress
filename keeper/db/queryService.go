package db

import "database/sql"

type QueryService struct {
	db *sql.DB
}

// func NewQueryService(db *sql.DB) *QueryService {
// 	return &QueryService{db: db}
// }

func (*QueryService) getApplicationEntries() {
}

func (*QueryService) getPendingWrites() {}
