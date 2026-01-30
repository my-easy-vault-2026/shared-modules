package entities

import (
	"shared-modules/common"
	"shared-modules/utils"
	"time"

	"github.com/shopspring/decimal"
)

type EtherfiApplyCardCacheVO struct {
	UserID           uint64                             `json:"userId"`
	OrderID          uint64                             `json:"orderId"`
	CardID           uint64                             `json:"cardId"`
	OrderNo          string                             `json:"orderNo"`
	CategoryID       uint64                             `json:"categoryId"`
	ApplyCardPreview EtherfiCardPreview                 `json:"applyCardPreview"`
	EtherfiUserID    string                             `json:"etherfiUserId,omitempty"`
	RetryCount       int                                `json:"retryCount,omitempty"`
	IssueStatus      common.EtherfiApplyCardCacheStatus `json:"issueStatus,omitempty"`   // "pending" | "failed" | "succeeded"
	NextAttemptAt    int64                              `json:"nextAttemptAt,omitempty"` // unix seconds
	FirstFailedAt    int64                              `json:"firstFailedAt,omitempty"`
	EtherfiId        uint64                             `json:"etherfiId,omitempty"`
}

type EtherfiCardPreview struct {
	UserID                uint64            `json:"userID"`
	DepositAmount         decimal.Decimal   `json:"depositAmount"`
	DepositAmountShortage *decimal.Decimal  `json:"depositAmountShortage,omitempty"`
	FromCardID            uint64            `json:"fromCardID"`
	FromCategoryID        uint64            `json:"fromCategoryID"`
	FromCurrency          common.Currency   `json:"fromCurrency"`
	ToCategoryID          uint64            `json:"toCategoryID"`
	ToCurrency            common.Currency   `json:"toCurrency"`
	ReceiveAmount         decimal.Decimal   `json:"receiveAmount"`
	FromDiscount          *decimal.Decimal  `json:"fromDiscount,omitempty"` // 幣種為來源幣種
	ToDiscount            *decimal.Decimal  `json:"toDiscount,omitempty"`   // 幣種為到帳幣種
	Bonus                 *decimal.Decimal  `json:"bonus,omitempty"`        // 幣種為目的幣種
	PromotionCode         string            `json:"promotionCode,omitempty"`
	Fee                   *decimal.Decimal  `json:"fee"`
	CardFee               *decimal.Decimal  `json:"cardFee"`
	Rate                  []*ExchangeRate   `json:"rate"`
	DisplayRate           *decimal.Decimal  `json:"displayRate"`
	AddressLine1          string            `json:"addressLine1"`
	AddressLine2          string            `json:"addressLine2"`
	AddressLine3          string            `json:"addressLine3"`
	AddressLine4          string            `json:"addressLine4"`
	AddressLine5          string            `json:"addressLine5"`
	State                 string            `json:"state"`
	NationCode            common.NationCode `json:"NationCode"`
	PostalCode            string            `json:"postalCode"`
	City                  string            `json:"city"`
	ExpiredAt             time.Time         `json:"expiredAt"`
	UserTempAddressID     *uint64           `json:"userTempAddressID"` //用戶臨時地址
}

type EtherfiTransactionVO struct {
}

type PaymentOTPInfo struct {
	Code       string  // 924420
	Merchant   string  // PChome
	Currency   string  // TWD
	AmountText string  // 91.00
	Amount     float64 // 91.00
}

type CardIssuePair struct {
	CardID  uint64 `gorm:"column:card_id"`
	IssueID string `gorm:"column:issue_id"`
}

type EtherfiInviteRetryCacheVO struct {
	InviteURL string `json:"inviteUrl"`
	Mail      string `json:"mail"`
}

type EtherfiAcceptRetryCacheVO struct {
	Mail          string                           `json:"mail"`
	AcceptRequest utils.EtherfiAcceptInviteRequest `json:"acceptRequest"`
	UserID        uint64                           `json:"userId"`
}
