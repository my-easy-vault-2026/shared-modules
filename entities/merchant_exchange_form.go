package entities

import "github.com/shopspring/decimal"

type MerchantExchangePreviewForm struct {
	FromCardID uint64           `json:"fromCardId,string" validate:"required"`
	ToCardID   uint64           `json:"toCardId,string"`
	ToCategory string           `json:"toCategory"`
	FromAmount *decimal.Decimal `json:"fromAmount"`
	ToAmount   *decimal.Decimal `json:"toAmount"`
}

type MerchantExchangeConfirmForm struct {
	ExchangeKey string `json:"exchangeKey" validate:"required"`
}
