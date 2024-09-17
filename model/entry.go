package model

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	Amount int `gorm:"not null"`

	TransferID uint
}
