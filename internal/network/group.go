package network

import (
	"encoding/json"
	"fmt"

	"github.com/srspu-services/schedule-parser-microservice/internal/models"
	"github.com/srspu-services/schedule-parser-microservice/pkg/httpclient"
)

type Group struct {
	client *httpclient.Client
}

func NewGroup(client *httpclient.Client) *Group {
	return &Group{client: client}
}

func (i *Group) GetGroup(faculty models.Faculty, course int) ([]models.Group, error) {
	var resJson [][]string
	var groups []models.Group

	res, err := i.client.NewRequest(fmt.Sprintf("%s/%s/years/%d/groups/", URL, faculty.Index, course))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(res), &resJson); err != nil {
		return nil, err
	}

	for _, v := range resJson {
		groups = append(groups, models.Group{
			GroupId: v[0],
			Name:    v[1],
			Course:  course,
		})
	}

	return groups, nil
}
