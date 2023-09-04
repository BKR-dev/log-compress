package db

import "database/sql"

type QueryService struct {
	qS *sql.DB
}

func NewQueryService() {
	qs := QueryService
	return qs
}

func (*QueryService) getApplicationEntries() {
}


func (*QueryService) getOpenWrites() {}
