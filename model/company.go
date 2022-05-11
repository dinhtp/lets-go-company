package model

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name      string
	Phone     string
	Email     string
	Address   string
	TaxNumber string
}
