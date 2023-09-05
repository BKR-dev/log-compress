package model

// used for DTOs and DAOs definition

type LogEntry struct {
	Hostname         string
	ApplicationName  string
	StartTime        string
	EndTime          string
	CalendarWeek     int16
	FileSize         int32
	FileLastModified string
}

type Archive struct {
	Hostname        string
	ApplicationName string
	CalendarWeek    int16
	FileSize        int32
	FinishTime      string
	Completed       bool
}
