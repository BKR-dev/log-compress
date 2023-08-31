package main

import (
	"fmt"
	"keeper/db"
	"keeper/server"
	"keeper/util"
)

func main() {
	dB, err := db.GetDB()
	if err != nil {
		fmt.Println(err)
	}
	db.PrintSQLVersion(dB)

	server.ServerStart()

	util.GetCalendarWeek()
}
