package model

import (
	"gorm.io/gorm"
)

// used for DTOs and DAOs definition
type LogEntry struct {
	gorm.Model
	Hostname         string
	ApplicationName  string
	StartTime        string
	EndTime          string
	CalendarWeek     int16
	FileSize         int32
	FileLastModified string
}

type Archive struct {
	gorm.Model
	Hostname        string
	ApplicationName string
	CalendarWeek    int16
	FileSize        int32
	FinishTime      string
	Completed       bool
}

// used for API and WEB
type LogString struct {
	ID               string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
	Hostname         string
	ApplicationName  string
	StartTime        string
	EndTime          string
	CalendarWeek     string
	FileSize         string
	FileLastModified string
}

type ArchiveString struct {
	ID              string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
	Hostname        string
	ApplicationName string
	CalendarWeek    string
	FileSize        string
	FinishTime      string
	Completed       string
}
