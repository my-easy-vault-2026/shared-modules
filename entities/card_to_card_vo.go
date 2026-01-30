package entities

import "github.com/shopspring/decimal"

type CardToCardPreviewVO struct {
	ToAmount       string            `json:"toAmount"`
	ToUserID       uint64            `json:"toUserId,string"`
	ToCardID       uint64            `json:"toCardID,string"`
	ToCategory     string            `json:"toCategory,omitempty"`
	ToCategoryID   uint64            `json:"toCategoryId,string"`
	ToCurrency     string            `json:"toCurrency"`
	ToEmail        string            `json:"toEmail,omitempty"`
	FromAmount     string            `json:"fromAmount"`
	FromCardID     uint64            `json:"fromCardId,string"`
	FromCategory   string            `json:"fromCategory,omitempty"`
	FromCategoryID uint64            `json:"fromCategoryId,string"`
	FromCurrency   string            `json:"fromCurrency"`
	ExchangeRate   *decimal.Decimal  `json:"exchangeRate,omitempty"`
	ExchangeFee    string            `json:"exchangeFee,omitempty"`
	Fee            string            `json:"fee,omitempty"`
	Rate           []*ExchangeRateVO `json:"rate"`
	DisplayRate    ExchangeRateVO    `json:"displayRate"`
	Key            string            `json:"key"`
	ExpiredAt      int64             `json:"expiredAt,string"`
}

type CardToCardConfirmVO struct{}
