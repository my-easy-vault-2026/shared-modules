package entities

import (
	"shared-modules/common"

	"github.com/shopspring/decimal"
)

type TopUpPreviewForm struct {
	FromCardID        uint64                      `json:"fromCardId,string" validate:"required"`
	ToCardID          uint64                      `json:"toCardId,string" validate:"required"`
	FromAmount        *decimal.Decimal            `json:"fromAmount"`
	ToAmount          *decimal.Decimal            `json:"toAmount"`
	Usage             common.UserTempAddressUsage `json:"-"`
	Address           string                      `json:"-"`
	UserTempAddressID uint64                      `json:"-"`
}

type TopUpConfirmForm struct {
	TopUpKey string                      `json:"topUpKey" validate:"required"`
	Usage    common.UserTempAddressUsage `json:"-"`
	Address  string                      `json:"-"`
	// PinCode  string `json:"pinCode" validate:"required"`
}

type DappCardDepositApplyForm struct {
	Currency   string           `json:"currency" validate:"required"`
	Mainnet    string           `json:"mainnet" validate:"required"`
	ToCardID   uint64           `json:"toCardId,string" validate:"required"`
	FromAmount *decimal.Decimal `json:"fromAmount"`
	ToAmount   *decimal.Decimal `json:"toAmount"`
}

type DappCreateDepositOrderForm struct {
	TXHash      string `json:"txHash" validate:"required"`
	TXIndex     string `json:"txIndex" validate:"required"`
	FromAddress string `json:"fromAddress" validate:"required"`
	ToAddress   string `json:"toAddress" validate:"required"`
}
