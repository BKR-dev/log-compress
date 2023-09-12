package server

import (
	"local/model"
	"strconv"
)

// transform all entries into strings
func trfmLstr(logs []model.LogEntry) []model.LogString {
	var logStr model.LogString
	var logStrings []model.LogString
	for _, v := range logs {
		logStr.ID = string(v.ID)
		logStr.CreatedAt = v.CreatedAt.String()
		logStr.UpdatedAt = v.UpdatedAt.String()
		logStr.DeletedAt = v.DeletedAt.Time.String()
		logStr.Hostname = string(v.Hostname)
		logStr.ApplicationName = string(v.ApplicationName)
		logStr.StartTime = string(v.StartTime)
		logStr.EndTime = string(v.EndTime)
		logStr.CalendarWeek = string(v.CalendarWeek)
		logStr.FileSize = string(v.FileSize)
		logStr.FileLastModified = string(v.FileLastModified)
		logStrings = append(logStrings, logStr)
	}
	return logStrings
}

// transform all entries into strings
func trfmAstr(archs []model.Archive) []model.ArchiveString {
	var archString []model.ArchiveString
	var archy model.ArchiveString
	for _, v := range archs {
		archy.ID = string(v.ID)
		archy.CreatedAt = v.CreatedAt.String()
		archy.UpdatedAt = v.UpdatedAt.String()
		archy.DeletedAt = v.DeletedAt.Time.String()
		archy.Hostname = string(v.Hostname)
		archy.ApplicationName = string(v.ApplicationName)
		archy.CalendarWeek = string(v.CalendarWeek)
		archy.FileSize = string(v.FileSize)
		archy.Completed = strconv.FormatBool(v.Completed)
		archString = append(archString, archy)
	}
	return archString
}
