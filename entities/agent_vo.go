package entities

import (
	"shared-modules/common"
	"shared-modules/utils"
	"time"

	"github.com/shopspring/decimal"
)

type AgentUserVO struct {
	UserID    uint64 `json:"userId,string"`
	Email     string `json:"email"`
	UpdatedAt int64  `json:"updatedAt,string"`
}

type PageAgentUserVO struct {
	utils.PageData[[]*AgentUserVO]
}

type AgentTxReportVO struct {
	ID        uint64 `json:"id,string"`
	YearMonth string `json:"yearMonth"`
	FileName  string `json:"fileName"`
	CreatedAt int64  `json:"createdAt,string"`
}

type AgentTxReportPageVO struct {
	utils.PageData[[]*AgentTxReportVO]
}

type AgentCardPageVO struct {
	utils.PageData[[]*AgentCardVO]
}

type AgentCardVO struct {
	ID             uint64           `json:"id,string"` //可能為空
	UserID         uint64           `json:"userId,string"`
	Type           int              `json:"type"`
	Email          string           `json:"email"`
	Amount         decimal.Decimal  `json:"amount,omitempty"` //可能為空
	ProviderAmount *decimal.Decimal `json:"providerAmount,omitempty"`
	FrozenAmount   decimal.Decimal  `json:"frozenAmount,omitempty"` //可能為空
	Currency       string           `json:"currency"`
	CurrencyType   string           `json:"currencyType"`
	DiscountCode   string           `json:"discountCode"`
	Status         int              `json:"status"`
	CreatedAt      time.Time        `json:"createdAt,string"`
	UpdatedAt      time.Time        `json:"updatedAt,string"`
}

type AgentDiscountCardUserGetCardVO struct {
	ID           uint64            `json:"id,string"` //可能為空
	Currency     string            `json:"currency"`
	AssetType    common.AssetType  `json:"assetType"`
	DiscountCode string            `json:"discountCode"`
	Balance      decimal.Decimal   `json:"balance,omitempty"` //可能為空
	Format       common.CardFormat `json:"format"`
	Status       int               `json:"status"`
}

type AgentDiscountCardUserGetVO struct {
	UserID uint64                            `json:"userId,string"`
	Email  string                            `json:"email"`
	Status common.UserBlockStatus            `json:"status"`
	Cards  []*AgentDiscountCardUserGetCardVO `json:"cards"`
}
