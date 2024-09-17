package repository

import (
	"log"
	"testing"

	"github.com/mfitrahrmd/orion-pay/model"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createRandomWallet(db *gorm.DB) model.Wallet {
	wallet := model.Wallet{
		Balance: 1000000,
	}
	db.Create(&wallet)

	return wallet
}

func TestCreateTransaction(t *testing.T) {
	db, err := gorm.Open(postgres.Open("postgres://dev:dev@localhost:5432/orion_pay?sslmode=disable"))
	if err != nil {
		log.Fatalln(err)
	}
	repo := NewTransferRepository(db)
	amount := 1000
	errs := make(chan error)
	ress := make(chan *model.Transfer)
	for i := 0; i < 100; i++ {
		go func() {
			transfer := model.Transfer{
				Amount:            amount,
				WalletID:          1,
				RecipientWalletID: 2,
			}
			err := repo.CreateTransfer(&transfer)
			errs <- err
			ress <- &transfer
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			transfer := model.Transfer{
				Amount:            amount,
				WalletID:          2,
				RecipientWalletID: 1,
			}
			err := repo.CreateTransfer(&transfer)
			errs <- err
			ress <- &transfer
		}()
	}
	for i := 0; i < 200; i++ {
		err := <-errs
		res := <-ress
		require.NoError(t, err)
		require.Equal(t, amount, res.Amount)
	}
}
