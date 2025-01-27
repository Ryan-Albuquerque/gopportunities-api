package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)


func Init() error {
	return fmt.Errorf("erro aqui")
}