package database

import (
	"database/sql"

	"github.com/srspu-services/schedule-parser-microservice/internal/models"
)

type Schedule struct {
	db *sql.DB
}

func NewSchedule(db *sql.DB) *Schedule {
	return &Schedule{db: db}
}

func (i *Schedule) CreateSchedule(groupId int, schedule models.Schedule) error {
	_, err := i.db.Exec(`INSERT INTO schedule (
		group_id,
		up_week,
		week_day,
		start_time,
		end_time,
		subject,
		classroom,
		class_type
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		groupId,
		schedule.UpWeek,
		schedule.WeekDay,
		schedule.StartTime,
		schedule.EndTime,
		schedule.Name,
		schedule.ClassRoom,
		schedule.ClassType,
	)
	return err
}
