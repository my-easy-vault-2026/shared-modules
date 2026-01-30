package entities

import "github.com/shopspring/decimal"

type CardToCardPreviewForm struct {
	FromCardID uint64           `json:"fromCardId,string" validate:"required"`
	ToCardID   uint64           `json:"toCardId,string"`
	ToUserID   uint64           `json:"toUserId,string"`
	ToEmail    string           `json:"toEmail" validate:"omitempty,email"`
	FromAmount *decimal.Decimal `json:"fromAmount"`
	ToAmount   *decimal.Decimal `json:"toAmount"`
	Remark     *string          `json:"remark,omitempty"`
}

type CardToCardConfirmForm struct {
	ID            string             `json:"id"`
	RawID         string             `json:"rawId"`
	Type          string             `json:"type"`
	Response      VerifyResponseType `json:"response"`
	DeviceID      string
	CardToCardKey string `json:"cardToCardKey" validate:"required"`
	PinCode       string `json:"pinCode"`
}
