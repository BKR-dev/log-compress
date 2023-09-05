package model

import ()

// used for DTOs and DAOs definition

type LogEntry struct {
	Hostname string 
	ApplicationName string
	startTime string
	endTime string
	calendarWeek int16
	fileSize int32
	fileLastModified string
}

type Archive struct {
	Hostname string
	ApplicationName string
	calendarWeek int16
	fileSize int32
	finishTime string
	completed bool
}
