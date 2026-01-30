package entities

import "github.com/shopspring/decimal"

type TopDownPreviewForm struct {
	FromCardID uint64           `json:"fromCardId,string" validate:"required"`
	ToCardID   uint64           `json:"toCardId,string" validate:"required"`
	FromAmount *decimal.Decimal `json:"fromAmount,string"`
	ToAmount   *decimal.Decimal `json:"toAmount,string"`
}

type TopDownConfirmForm struct {
	TopDownKey string `json:"topDownKey" validate:"required"`
}
