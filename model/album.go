package model

import (
	"github.com/jinzhu/gorm"
)

type Album struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author" gorm:"column:author"`
}

func (Album) TableName() string {
	return "album"
}
