package model

import ()

type LogEntry struct {
	Hostname string 
	ApplicationNAme string
	startTime string
	endTime string
	calendarWeek int16
	fileSize int32
	fileLastModified string
}
