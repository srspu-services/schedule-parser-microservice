package network

import "github.com/srspu-services/schedule-parser-microservice/internal/models"

const URL = "https://schedule.npi-tu.ru/api/v1/faculties/"

type FacultyInterfase interface {
	GetFaculties() ([]models.Faculty, error)
}

type GroupInterface interface {
	GetGroup(facul, course int) ([]models.Group, error)
}

type ScheduleInterface interface {
	GetSchedule(groupName string) ([]models.Schedule, error)
}

type NetworkInterfaces struct {
	FacultyInterfase
	GroupInterface
	ScheduleInterface
}

func NewNetwork() *NetworkInterfaces {
	return nil
}
