package models

import (
	"database/sql"
	"time"
)

type MstProjects struct {
	Id             int
	Title          string
	Client         string
	Date           time.Time
	DateCompletion time.Time
	MediaUuid      sql.NullString
	Description    string
	Position       string
	CreatedAt      time.Time
	CreatedBy      string
	UpdatedAt      sql.NullTime
	UpdatedBy      sql.NullString
	DeletedAt      sql.NullTime
}
