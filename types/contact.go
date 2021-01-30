package types

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Priority uint8
	Title    string
	Value    string
	Href     string
}

type ContactDTO struct {
	Priority uint8  `json:"priority"`
	Title    string `json:"title"`
	Value    string `json:"value"`
	Href     string
}
