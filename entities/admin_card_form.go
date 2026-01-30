package entities

import (
	"shared-modules/common"
	"shared-modules/utils"
	"time"

	"github.com/shopspring/decimal"
)

type AdminPageCardForm struct {
	Role         common.Role       `json:"role"`
	UserID       uint64            `json:"userId,string"`
	CardID       uint64            `json:"cardId,string"`
	Email        string            `json:"email"`
	Vendor       int               `json:"vendor"`
	CardType     common.AssetType  `json:"cardType"`
	CardKind     common.CardKind   `json:"cardKind"`
	Format       common.CardFormat `json:"format"`
	Status       int               `json:"status"`
	FreezeStatus int               `json:"freezeStatus"`
	CategoryID   uint64            `json:"categoryId,string"`
	Category     string            `json:"category"`
	MerchantID   uint64            `json:"merchantId,string"`
	utils.Page
}

type AdminGetCardForm struct {
	ID      uint64 `json:"id,string"`
	IssueID string `json:"issueId"`
}

type AdminDeleteCardForm struct {
	CardID uint64 `json:"cardId,string"`
	Remark string `json:"remark"`
}

type AdminRetryAppleCardViaApplePayForm struct {
	TransactionId string `json:"transactionId"`
	UserId        uint64 `json:"userId,string"`
}

type AdminAgentListUser struct {
	UserId string `json:"userId"`
}

type AdminAgentPageTx struct {
	UserID        uint64 `json:"userID,string"`
	AgentID       uint64 `json:"-"`
	CreatedAtFrom int64  `json:"createdAtFrom,string" validate:"required"`
	CreatedAtTo   int64  `json:"createdAtTo,string" validate:"required"`
	utils.Page
}

type AdminBlockCardForm struct {
	CardID      uint64 `json:"cardId,string" validate:"required"`
	BlockReason string `json:"blockReason" validate:"required"`
}

type AdminUnblockCardForm struct {
	CardID uint64 `json:"cardId,string"`
}

type AdminApplyCard struct {
	UserID     uint64                   `json:"userID,string" validate:"required"`
	Vendor     common.CardProductVendor `json:"vendor" validate:"required"`
	CategoryId uint64                   `json:"categoryId,string" validate:"required"`
	WaiveFee   bool                     `json:"waiveFee"`
}

type AdminApplyCardPreview struct {
	UserID         uint64            `json:"userID"`
	DepositAmount  decimal.Decimal   `json:"depositAmount"`
	ReceiveAmount  decimal.Decimal   `json:"receiveAmount"`
	ToCategoryID   uint64            `json:"toCategoryID"`
	ToCurrency     common.Currency   `json:"toCurrency"`
	Fee            *decimal.Decimal  `json:"fee"`
	CardFee        *decimal.Decimal  `json:"cardFee"`
	NationCode     common.NationCode `json:"NationCode"`
	PostalCode     string            `json:"postalCode"`
	City           string            `json:"city"`
	State          string            `json:"state"`
	ExpiredAt      time.Time         `json:"expiredAt"`
	AddressLine1   string            `json:"addressLine1"`
	AddressLine2   string            `json:"addressLine2"`
	AddressLine3   string            `json:"addressLine3"`
	AddressLine4   string            `json:"addressLine4"`
	AddressLine5   string            `json:"addressLine5"`
	Rate           []*ExchangeRate   `json:"rate"`
	DisplayDecimal int               `json:"displayDecimal"`
}

type ExchangeRate struct {
	BaseCurrency  common.Currency `json:"baseCurrency"`
	QuoteCurrency common.Currency `json:"quoteCurrency"`
	Rate          decimal.Decimal `json:"rate"`
	Timestamp     time.Time       `json:"timestamp"`
}

type AdminCreatePaycryptoPhysicalCardNumberForm struct {
	PANNumbers []string `json:"panNumbers" validate:"required"`
}

type AdminDeleteTransferForm struct {
	UserID       uint64 `json:"userId,string" validate:"required"`
	CardID       uint64 `json:"cardId,string" validate:"required"`
	ToCardID     uint64 `json:"toCardId,string"`
	ToCategoryID uint64 `json:"toCategoryId,string"`
	Force        bool   `json:"force"`
}

type AdminApplyAndDeleteForm struct {
	UserID       uint64                   `json:"userID,string" validate:"required"`
	Vendor       common.CardProductVendor `json:"vendor" validate:"required"`
	CategoryId   uint64                   `json:"categoryId,string" validate:"required"`
	WaiveFee     bool                     `json:"waiveFee"`
	DeleteCardID uint64                   `json:"deleteCardID,string" validate:"required"`
	Force        bool                     `json:"force"`
}

type AdminCardDownloadPageForm struct {
	Role         common.Role       `json:"role"`
	UserID       uint64            `json:"userId,string"`
	CardID       uint64            `json:"cardId,string"`
	Email        string            `json:"email"`
	Vendor       int               `json:"vendor"`
	CardType     common.AssetType  `json:"cardType"`
	CardKind     common.CardKind   `json:"cardKind"`
	Format       common.CardFormat `json:"format"`
	Status       int               `json:"status"`
	FreezeStatus int               `json:"freezeStatus"`
	CategoryID   uint64            `json:"categoryId,string"`
	Category     string            `json:"category"`
	MerchantID   uint64            `json:"merchantId,string"`
	Size         int               `json:"size"`
}
