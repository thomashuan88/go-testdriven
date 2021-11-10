package models

import (
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Author    string
	Name      string
	PageCount int
}
