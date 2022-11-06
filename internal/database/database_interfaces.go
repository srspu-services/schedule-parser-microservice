package database

import (
	"database/sql"

	"github.com/srspu-services/schedule-parser-microservice/internal/models"
)

type GroupInterface interface {
	CreateGroup(group models.Group) (int, error)
}

type ScheduleInterface interface {
	CreateSchedule(groupId int, schedule models.Schedule) error
}

type DatabaseInterfaces struct {
	GroupInterface
	ScheduleInterface
}

func NewDatabaseInterfaces(db *sql.DB) *DatabaseInterfaces {
	return &DatabaseInterfaces{
		GroupInterface:    NewGroup(db),
		ScheduleInterface: NewSchedule(db),
	}
}
