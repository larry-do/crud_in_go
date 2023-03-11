package model

import (
	"github.com/jinzhu/gorm"
)

type Artist struct {
	gorm.Model
	Name string `json:"name"`
}

func (Artist) TableName() string {
	return "artist"
}
