package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type FiatVO struct {
	ID       uint64 `json:"id,string"`
	Name     string `json:"name"`
	Currency string `json:"currency"`
	IconUrl  string `json:"iconUrl"`
}

type ListFiatVO struct {
	Records []*FiatVO `json:"records"`
}

type ListOfferVO struct {
	Records []*ChangellyOfferVO `json:"records"`
}

type PreviewBinancePayOrderVO struct {
	ToCardID     uint64           `json:"toCardID,string"`
	FromAmount   decimal.Decimal  `json:"fromAmount"`
	FromCurrency string           `json:"fromCurrency"`
	ToAmount     decimal.Decimal  `json:"toAmount"`
	ToCurrency   string           `json:"toCurrency"`
	Fee          *decimal.Decimal `json:"fee"`
	FeeRate      decimal.Decimal  `json:"feeRate"`
	ExchangeRate decimal.Decimal  `json:"exchangeRate"`
	ExpiredAt    time.Time        `json:"expiredAt"`
}

type CreateBinancePayOrderVO struct {
	OrderId     string `json:"orderID"`
	RedirectUrl string `json:"redirectUrl"`
}

type QueryBinancePayOrderVO struct {
	OrderId      string `json:"orderID"`
	Status       string `json:"status"`
	Currency     string `json:"currency"`
	OrderAmount  string `json:"orderAmount"`
	Commission   string `json:"commission"`
	CreateTime   int64  `json:"createTime"`
	TransactTime int64  `json:"transactTime"`
}
