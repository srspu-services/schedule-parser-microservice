package models

import "time"

type Schedule struct {
	GroupId   int
	UpWeek    bool
	WeekDay   int
	StartTime time.Time
	EndTime   time.Time
	Name      string
	ClassRoom string
	ClassType string
}

type InputSchedule struct {
	Week       int    `json:"week"`
	Day        int    `json:"day"`
	Class      int    `json:"class"`
	Auditorium string `json:"auditorium"`
	Lecturer   string `json:"lecturer"`
	Discipline string `json:"discipline"`
	Type       string `json:"type"`
}
