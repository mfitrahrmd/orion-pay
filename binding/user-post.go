package binding

import "github.com/mfitrahrmd/orion-pay/model"

var defaultBalance int = 100000

type UserPost struct {
	Username string `binding:"required"`
	Email    *string
	FullName *string
}

func (up *UserPost) ToUser() model.User {
	return model.User{
		Username: up.Username,
		Email:    up.Email,
		FullName: up.FullName,
		Wallet: &model.Wallet{
			Balance: defaultBalance,
		},
	}
}
