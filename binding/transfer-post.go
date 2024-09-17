package binding

import "github.com/mfitrahrmd/orion-pay/model"

type TransferPost struct {
	Amount      int  `binding:"required"`
	SenderID    uint `binding:"required"`
	RecipientID uint `binding:"required"`
}

func (tp *TransferPost) ToTransfer() *model.Transfer {
	return &model.Transfer{
		Amount:            tp.Amount,
		RecipientWalletID: tp.RecipientID,
		WalletID:          tp.SenderID,
	}
}
