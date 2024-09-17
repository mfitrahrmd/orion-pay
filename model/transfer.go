package model

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model
	Amount            int  `gorm:"not null"`
	RecipientWalletID uint `gorm:"not null"`

	WalletID uint

	SentEntry     Entry `gorm:"foreignKey:TransferID;constraint:onUpdate:CASCADE,OnDelete:SET NULL"`
	ReceivedEntry Entry `gorm:"foreignKey:TransferID;constraint:onUpdate:CASCADE,OnDelete:SET NULL"`
}
