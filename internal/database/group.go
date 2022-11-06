package database

import (
	"database/sql"

	"github.com/srspu-services/schedule-parser-microservice/internal/models"
)

type Group struct {
	db *sql.DB
}

func NewGroup(db *sql.DB) *Group {
	return &Group{db: db}
}

func (i *Group) CreateGroup(group models.Group) (int, error) {
	var id int
	tx, err := i.db.Begin()
	if err != nil {
		return id, err
	}
	stmt, err := tx.Prepare(`INSERT INTO groups(
		group_name,
		group_id,
		course
	) VALUES($1, $2, $3)
	RETURNING id`)

	if err != nil {
		return id, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(
		group.Name,
		group.GroupId,
		group.Course,
	).Scan(&id)

	if err != nil {
		return id, err
	}

	return id, nil
}
