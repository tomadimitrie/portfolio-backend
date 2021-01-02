package types

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type About struct {
	gorm.Model
	Priority uint8
	Title    string
	Value    string
	Details  pq.StringArray `gorm:"type:text[]"`
}

type AboutDTO struct {
	Priority uint8
	Title    string
	Value    string
	Details  pq.StringArray `gorm:"type:text[]"`
}
