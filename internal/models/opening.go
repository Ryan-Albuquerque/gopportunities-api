package models

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role string
	Company string
	Location string
	Link string
	Salary int64
	Remote bool
	IsActived bool
}