package common

import (
	"time"

	"github.com/shopspring/decimal"
)

type ExchangeRate struct {
	BaseCurrency  Currency        `json:"baseCurrency"`
	QuoteCurrency Currency        `json:"quoteCurrency"`
	Rate          decimal.Decimal `json:"rate"`
	Purpose       RatePurpose     `json:"purpose,omitempty"`
	Timestamp     time.Time       `json:"timestamp"`
}

type RatePurpose int // 匯率用途
const (
	RATE_PURPOSE_EXCHANGE RatePurpose = 1
	RATE_PURPOSE_TRANSFER RatePurpose = 2
	RATE_PURPOSE_FEE      RatePurpose = 3
)

func (r RatePurpose) String() string {
	switch r {
	case RATE_PURPOSE_EXCHANGE:
		return "exchange"
	case RATE_PURPOSE_TRANSFER:
		return "transfer"
	}
	return ""
}
