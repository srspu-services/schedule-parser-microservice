package network

import (
	"encoding/json"
	"fmt"

	"github.com/srspu-services/schedule-parser-microservice/internal/models"
	"github.com/srspu-services/schedule-parser-microservice/internal/utils"
	"github.com/srspu-services/schedule-parser-microservice/pkg/httpclient"
)

type Schedule struct {
	client *httpclient.Client
}

func NewSchedule(client *httpclient.Client) *Schedule {
	return &Schedule{client: client}
}

type scheduleJson struct {
	Classes []models.InputSchedule `json:"classes"`
}

func (i *Schedule) GetSchedule(facultyIndex string, groupId int, group models.Group) ([]models.Schedule, error) {
	var input scheduleJson
	var schedule []models.Schedule

	res, err := i.client.NewRequest(fmt.Sprintf("%s/%s/years/%d/groups/%s/schedule",
		URL, facultyIndex, group.Course, group.GroupId))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(res), &input); err != nil {
		return nil, err
	}

	for _, v := range input.Classes {
		isUpWeek := false
		if v.Week == 2 {
			isUpWeek = true
		}
		schedule = append(schedule, models.Schedule{
			GroupId:   groupId,
			UpWeek:    isUpWeek,
			WeekDay:   v.Day,
			StartTime: utils.GetStartTime(v.Class),
			EndTime:   utils.GetEndTime(v.Class),
			Name:      v.Discipline,
			ClassRoom: v.Auditorium,
			ClassType: v.Type,
		})
	}

	return schedule, nil
}
