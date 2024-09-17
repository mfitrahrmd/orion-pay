package repository

import (
	"errors"

	"github.com/mfitrahrmd/orion-pay/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransferRepository struct {
	Db *gorm.DB
}

func NewTransferRepository(db *gorm.DB) *TransferRepository {
	return &TransferRepository{
		Db: db,
	}
}

func (tr *TransferRepository) CreateTransfer(transfer *model.Transfer) error {
	tx := tr.Db.Begin()
	transfer.SentEntry = model.Entry{
		Amount: -transfer.Amount,
	}
	transfer.ReceivedEntry = model.Entry{
		Amount: transfer.Amount,
	}
	if err := tx.Create(&transfer).Error; err != nil {
		tx.Rollback()

		return err
	}
	var wallets []model.Wallet
	if err := tx.Clauses(clause.Locking{Strength: "NO KEY UPDATE"}).Order("id ASC").Find(&wallets, "id IN ?", []uint{transfer.WalletID, transfer.RecipientWalletID}).Error; err != nil {
		tx.Rollback()

		return err
	}
	var sender model.Wallet
	var recipient model.Wallet
	if transfer.WalletID < transfer.RecipientWalletID {
		sender = wallets[0]
		recipient = wallets[1]
	} else {
		sender = wallets[1]
		recipient = wallets[0]

	}
	if sender.Balance < transfer.Amount {
		tx.Rollback()

		return errors.New("insufficient balance")
	}
	sender.Balance -= transfer.Amount
	if err := tx.Save(&sender).Error; err != nil {
		tx.Rollback()

		return err
	}
	recipient.Balance += transfer.Amount
	if err := tx.Save(&recipient).Error; err != nil {
		tx.Rollback()

		return err
	}
	if err := tx.Save(&transfer).Error; err != nil {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}
