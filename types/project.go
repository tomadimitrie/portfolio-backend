package types

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Priority     uint8
	Title        string
	Description  string
	Link         string
	Repo         string
	Technologies pq.StringArray `gorm:"type:text[]"`
}

type ProjectDTO struct {
	Priority     uint8          `json:"priority"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	Link         string         `json:"link"`
	Repo         string         `json:"repo"`
	Technologies pq.StringArray `json:"technologies" gorm:"type:text[]"`
}
