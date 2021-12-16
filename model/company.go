package model

import (
	"gorm.io/gorm"
	"time"
)

type Company struct {
	ID        uint
	Name      string
	Phone     string
	Email     string
	Address   string
	TaxNumber string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
