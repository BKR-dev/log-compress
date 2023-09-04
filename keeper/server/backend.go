package server

import (
	"local/db"
)

var (
	qS = db.NewQueryService()
)

func returnData(target string) ([]string, error) {
	return nil, nil
}
