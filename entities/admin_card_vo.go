package entities

import (
	"shared-modules/common"
	"shared-modules/utils"
	"time"

	"github.com/shopspring/decimal"
)

type AdminCardVO struct {
	ID               uint64                   `json:"id,string"` //可能為空
	UserID           uint64                   `json:"userId,string"`
	Type             common.AssetType         `json:"type"`
	UserFirstName    string                   `json:"userFirstName,omitempty"`
	UserLastName     string                   `json:"userLastName,omitempty"`
	ProductName      string                   `json:"productName,omitempty"`
	PreferredName    string                   `json:"preferredName,omitempty"`
	SecondaryName    string                   `json:"secondaryName,omitempty"`
	Alias            string                   `json:"alias,omitempty"`
	CategoryID       uint64                   `json:"categoryId,string"`
	IssueID          string                   `json:"issueId,omitempty"`
	PANNumber        string                   `json:"panNumber,omitempty"` //可能為空
	Amount           decimal.Decimal          `json:"amount,omitempty"`    //可能為空
	ProviderAmount   *decimal.Decimal         `json:"providerAmount,omitempty"`
	FrozenAmount     decimal.Decimal          `json:"frozenAmount,omitempty"` //可能為空
	Supported        []*SupportedVO           `json:"supported,omitempty"`    //可能為空
	Currency         string                   `json:"currency"`
	CurrencyType     string                   `json:"currencyType"`
	Issuer           string                   `json:"issuer"`
	Format           string                   `json:"format,omitempty"`
	SpendLimit       decimal.Decimal          `json:"spendLimit,omitempty"`
	Design           string                   `json:"design,omitempty"`
	MerchantID       uint64                   `json:"merchantId,omitempty,string"`
	MerchantName     string                   `json:"merchantName,omitempty"`
	BalanceType      string                   `json:"balanceType,omitempty"`
	ForwardType      string                   `json:"forwardType,omitempty"`
	Status           string                   `json:"status"`
	BlockReason      common.CardBlockReason   `json:"blockReason,omitempty"`
	FreezeReason     common.CardFreezeReason  `json:"freezeReason,omitempty"`
	FreezeStatus     string                   `json:"freezeStatus"`
	ReapDeleteStatus common.ReapDeleteStatus  `json:"reapDeleteStatus,omitempty"`
	CreatedAt        time.Time                `json:"createdAt,string"`
	UpdatedAt        time.Time                `json:"updatedAt,string"`
	DeletedAt        *time.Time               `json:"deletedAt,string"`
	LastActiveOn     *int64                   `json:"lastActiveOn,string"`
	Billing          *Billing                 `json:"billing,omitempty"`
	AssetProductId   string                   `json:"assetProductId,omitempty"`
	Vendor           string                   `json:"vendor,omitempty"`
	VendorNum        common.CardProductVendor `json:"vendorNum"`
	TypeNum          common.AssetType         `json:"typeNum"`
	StatusNum        common.CardStatus        `json:"statusNum"`
	FreezeStatusNum  common.CardFreezeStatus  `json:"freezeStatusNum"`
	CardKind         common.CardKind          `json:"cardKind"`
	CardBin          string                   `json:"cardBin"`
	CardDesign       string                   `json:"cardDesign"`
	CurrentBalance   *decimal.Decimal         `json:"currentBalance,string.omitempty"`
}

type Billing struct {
	Name       string          `json:"name"`
	Address    string          `json:"address"`
	City       string          `json:"city"`
	PostalCode string          `json:"postalCode"`
	Country    int             `json:"country"`
	Balance    decimal.Decimal `json:"balance"`
}

type AdminPageCardVO struct {
	utils.PageData[[]*AdminCardVO]
}

type AgentUserTransactionVO struct {
	UserID       uint64                   `json:"user_id,string"`
	Amount       decimal.Decimal          `json:"amount,string"`
	Side         string                   `json:"side"` //交易方 1:from 2:to
	FromCurrency string                   `json:"fromCurrency"`
	FromAmount   string                   `json:"fromAmount"`
	ToCurrency   string                   `json:"toCurrency"`
	ToAmount     string                   `json:"toAmount"`
	Status       common.TransactionStatus `json:"status"` // common.TransactionStatus
	CreatedAt    int64                    `json:"createdAt"`
	CardType     string                   `json:"cardType"` // 卡片類型
}

type AgentUserTransactionReportVO struct {
	UserID       uint64                   `json:"user_id,string"`
	Amount       decimal.Decimal          `json:"amount,string"`
	Side         string                   `json:"side"` //交易方 1:from 2:to
	FromCurrency string                   `json:"fromCurrency"`
	FromAmount   string                   `json:"fromAmount"`
	ToCurrency   string                   `json:"toCurrency"`
	ToAmount     string                   `json:"toAmount"`
	Status       common.TransactionStatus `json:"status"` // common.TransactionStatus
	CreatedAt    string                   `json:"createdAt"`
}

type AgentUserPageTxVO struct {
	utils.PageData[[]*AgentUserTransactionVO]
}

type AgentReportTransactionVO struct {
	UserID   string          `json:"userID,string"`
	OrderNo  string          `json:"orderNo"`
	Amount   decimal.Decimal `json:"amount,string"`
	CardType string          `json:"cardType"`
	//Commission decimal.Decimal `json:"commission,string"`
}

type AdminApplyCardVO struct {
	CardID uint64 `json:"cardID,string"`
	CardNo string `json:"cardNo"`
}

type AdminApplyAndDeleteCardVO struct {
	Success bool   `json:"success"`
	CardID  uint64 `json:"cardID,string"`
	CardNo  string `json:"cardNo"`
	Stage   int    `json:"stage"`
	ErrMsg  string `json:"errMsg"`
}

type AdminCardExcelVO struct {
	ID           uint64                   `json:"id,string"`
	UserID       uint64                   `json:"userId,string"`
	PANNumber    string                   `json:"panNumber"`
	Amount       decimal.Decimal          `json:"amount"`
	FrozenAmount decimal.Decimal          `json:"frozenAmount"`
	Format       string                   `json:"format"`
	Status       string                   `json:"status"`
	BlockReason  common.CardBlockReason   `json:"blockReason"`
	FreezeReason common.CardFreezeReason  `json:"freezeReason"`
	FreezeStatus string                   `json:"freezeStatus"`
	CreatedAt    time.Time                `json:"createdAt,string"`
	UpdatedAt    time.Time                `json:"updatedAt,string"`
	DeletedAt    *time.Time               `json:"deletedAt,string"`
	LastActiveOn *int64                   `json:"lastActiveOn,string"`
	VendorNum    common.CardProductVendor `json:"vendorNum"`
	CardKind     string                   `json:"cardKind"`
	CardBin      string                   `json:"cardBin"`
	CardProvider string                   `json:"cardProvider"`
	CardType     string                   `json:"cardType"`
}
