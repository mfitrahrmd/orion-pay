package model

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	Balance int `gorm:"not null"`

	Transfers []Transfer `gorm:"foreignKey:WalletID;constraint:onUpdate:CASCADE,OnDelete:SET NULL"`
}
