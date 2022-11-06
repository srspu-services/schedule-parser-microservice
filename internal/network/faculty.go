package network

import (
	"encoding/json"

	"github.com/srspu-services/schedule-parser-microservice/internal/models"
	"github.com/srspu-services/schedule-parser-microservice/pkg/httpclient"
	"golang.org/x/exp/maps"
)

type Faculty struct {
	client *httpclient.Client
}

func NewFaculty(client *httpclient.Client) *Faculty {
	return &Faculty{client: client}
}

func (i *Faculty) GetFaculties() ([]models.Faculty, error) {
	var faculties []models.Faculty
	var data map[string]interface{}

	res, err := i.client.NewRequest(URL)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		return nil, err
	}

	keys := maps.Keys(data)
	for _, v := range keys {
		faculties = append(faculties, models.Faculty{Index: v})
	}

	return faculties, nil
}
