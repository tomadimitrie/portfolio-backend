package types

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	Priority uint8
	Title    string
	Value    pq.StringArray `gorm:"type:text[]"`
}

type SkillDTO struct {
	Priority uint8          `json:"priority"`
	Title    string         `json:"title"`
	Value    pq.StringArray `json:"value" gorm:"type:text[]"`
}
