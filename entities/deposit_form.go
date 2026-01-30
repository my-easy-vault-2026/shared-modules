package entities

import (
	"github.com/shopspring/decimal"
)

type ListFiatForm struct {
}

type GetChangellyOffersForm struct {
	FromAmount   decimal.Decimal `json:"fromAmount" validate:"required"`
	FromCurrency string          `json:"fromCurrency" validate:"required"`
	ToCurrency   string          `json:"toCurrency" validate:"required"`
	Mainnet      string          `json:"mainnet,omitempty"`
	Protocol     string          `json:"protocol,omitempty"`
}

type CreateChangellyOrderForm struct {
	FromCurrency  string `json:"fromCurrency" validate:"required"`
	ToCurrency    string `json:"toCurrency" validate:"required"`
	FromAmount    string `json:"fromAmount" validate:"required"`
	ProviderCode  string `json:"providerCode" validate:"required"`
	PaymentMethod string `json:"paymentMethod" validate:"required"`
	ToCardID      uint64 `json:"toCardID,string" validate:"required"`
	CategoryID    uint64 `json:"categoryId,string,omitempty"`
	Mainnet       string `json:"mainnet,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
}

type PreviewBinancePayOrderForm struct {
	ToCardID uint64          `json:"toCardID,string" validate:"required"`
	Currency string          `json:"currency" validate:"required"`
	Amount   decimal.Decimal `json:"amount" validate:"required"`
}

type CreateBinancePayOrderForm struct {
	ToCardID  uint64          `json:"toCardID,string" validate:"required"`
	Currency  string          `json:"currency" validate:"required"`
	Amount    decimal.Decimal `json:"amount" validate:"required"`
	ReturnUrl string          `json:"returnUrl" validate:"required"`
}

type QueryBinancePayOrderForm struct {
	OrderIds []uint64 `json:"orderIDs,string"`
}
