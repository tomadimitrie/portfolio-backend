package types

import (
	"gorm.io/gorm"
)

type Constant struct {
	gorm.Model
	Key   string
	Value string
}

type ConstantDTO struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
}
