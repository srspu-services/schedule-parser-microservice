package models

type Group struct {
	GroupId string `json:"group_id"`
	Name    string `json:"name"`
	Course  int    `json:"course"`
}
