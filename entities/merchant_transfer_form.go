package entities

import "github.com/shopspring/decimal"

type MerchantTransferPreviewForm struct {
	FromCardID uint64           `json:"fromCardId,string" validate:"required"`
	ToCardID   uint64           `json:"toCardId,string"`
	ToUserID   uint64           `json:"toUserId,string"`
	ToEmail    string           `json:"toEmail" validate:"omitempty,email"`
	ToAddress  string           `json:"toAddress"`
	Mainnet    string           `json:"mainnet"`
	Protocol   string           `json:"protocol"`
	FromAmount *decimal.Decimal `json:"fromAmount"`
	ToAmount   *decimal.Decimal `json:"toAmount"`
}

type MerchantTransferConfirmForm struct {
	TransferKey string `json:"transferKey" validate:"required"`
}
