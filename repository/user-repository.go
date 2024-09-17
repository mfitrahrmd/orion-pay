package repository

import (
	"github.com/mfitrahrmd/orion-pay/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	if err := ur.Db.Debug().Create(user).Error; err != nil {
		return err
	}
	if err := ur.Db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUsers(users *[]model.User) error {
	tx := ur.Db.Begin()
	if err := tx.Preload("Wallet").Find(users).Error; err != nil {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (ur *UserRepository) GetUser(user *model.User) error {
	tx := ur.Db.Begin()
	if err := tx.Preload("Wallet").Find(user).Error; err != nil {
		tx.Rollback()

		return err
	}

	return nil
}
