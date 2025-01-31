package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role      string
	Company   string
	Location  string
	Link      string
	Salary    int64
	Remote    bool
	IsActived bool
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
	Remote   bool   `json:"remote"`
}

func (or *CreateOpeningRequest) ValidateCreateOpeningRequest() error {
	if or.Role == "" || or.Company == "" || or.Location == "" || or.Link == "" || or.Salary <= 0 {
		return fmt.Errorf("invalid request")
	}
	return nil
}
