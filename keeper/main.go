package main

import (
	"fmt"
	"local/db"
	"local/server"
	"local/util"
)

func main() {
	dB, err := db.GetDB()
	if err != nil {
		fmt.Println(err)
	}

	db.PrintSQLVersion(dB)
	// os.Exit(1)
	util.GetCalendarWeek()
	server.ServerStart()
}
