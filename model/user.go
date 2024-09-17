package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `gorm:"not null;index:,unique"`
	Email    *string `gorm:"null;index:,unique"`
	FullName *string `gorm:"null"`

	Wallet *Wallet `gorm:"foreignKey:ID;constraint:onUpdate:CASCADE,OnDelete:SET NULL"`
}
