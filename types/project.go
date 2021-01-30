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
	Priority     uint8
	Title        string
	Description  string
	Link         string
	Repo         string
	Technologies pq.StringArray `gorm:"type:text[]"`
}
