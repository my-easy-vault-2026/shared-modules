package entities

import (
	"shared-modules/common"
	"shared-modules/utils"

	"github.com/shopspring/decimal"
)

type MerchantUpdateForwardTypeVO struct {
}

type MerchantPrivacyInfoVO struct {
	CardNumber string `json:"panNumber"`
	ExpiryDate string `json:"expiryDate"`
	CCV        string `json:"ccv"`
}

type MerchantPageCardVO struct {
	utils.PageData[[]*MerchantCardVO]
}

type MerchantCardApplyVO struct {
	UserID uint64 `json:"userId,string"`
	CardID uint64 `json:"cardId,string"`
}

type UpdateForwardTypeVO struct {
	Reap     bool `json:"reap"`
	Email    bool `json:"email"`
	SMS      bool `json:"sms"`
	APP      bool `json:"app"`
	Merchant bool `json:"merchant"`
}

type Merchant3dsVO struct {
	ID                   uint64            `json:"id,string"`
	CardID               uint64            `json:"cardId,string"`
	MerchantID           uint64            `json:"merchantId,string,omitempty"`
	Alias                string            `json:"alias" gorm:"default:null"`
	PANNumberSuffix      string            `json:"panNumberSuffix" gorm:"default:null"`
	Status               string            `json:"status"`
	MerchantName         string            `json:"merchantName"`
	TransactionAmount    string            `json:"transactionAmount"`
	MerchantCountryCode  common.NationCode `json:"merchantCountryCode"`
	TransactionCurrency  string            `json:"transactionCurrency"`
	TransactionTimestamp int64             `json:"transactionTimestamp,string"`
	ExpiredAt            int64             `json:"expiredAt,string,omitempty" gorm:"default:null"`
	CreatedAt            int64             `json:"createdAt,string,omitempty" gorm:"default:null"`
	UpdatedAt            int64             `json:"updatedAt,string,omitempty" gorm:"default:null"`
}

type MerchantCheck3dsVO struct {
	Records []*Merchant3dsVO `json:"records"`
}

type MerchantAdjustBalanceVO struct {
	BalanceBefore decimal.Decimal `json:"balanceBefore"`
	BalanceAfter  decimal.Decimal `json:"balanceAfter"`
	Currency      string          `json:"currency"`
}

type MerchantCardVO struct {
	ID               uint64                 `json:"id,string,omitempty"`
	UserID           uint64                 `json:"userId,string"`
	Email            string                 `json:"email"`
	CategoryID       uint64                 `json:"categoryId,string"`
	Vendor           string                 `json:"vendor"`
	PreferredName    string                 `json:"preferredName"`
	Type             string                 `json:"type"`
	Organization     string                 `json:"organization,omitempty"`
	Amount           decimal.Decimal        `json:"amount,omitempty"`
	PANNumber        string                 `json:"panNumber,omitempty"`
	Issuer           string                 `json:"issuer"`
	Currency         string                 `json:"currency"`
	PointLevel       string                 `json:"pointLevel,omitempty"`       //可能為空
	PointToNextLevel *decimal.Decimal       `json:"pointToNextLevel,omitempty"` //
	NextPointLevel   string                 `json:"nextPointLevel,omitempty"`   //可能為空
	Alias            string                 `json:"alias"`
	Format           string                 `json:"format"`
	Supported        []*SupportedVO         `json:"supported,omitempty"`
	CardSwitch       map[string]interface{} `json:"cardSwitch,omitempty"`
	BalanceType      string                 `json:"balanceType,omitempty"`
	ForwardType      string                 `json:"forwardType,omitempty"`
	AutoYieldStatus  string                 `json:"autoYieldStatus,omitempty"`
	Design           string                 `json:"design,omitempty"`
	DisplayStatus    common.DisplayStatus   `json:"displayStatus,omitempty"`
	DeliveryStatus   string                 `json:"deliveryStatus,omitempty"`
	TrackingLink     string                 `json:"trackingLink,omitempty"`
	Status           string                 `json:"status"`
	FreezeStatus     string                 `json:"freezeStatus"`
	HasCardIdentity  bool                   `json:"hasCardIdentity"`
	CreatedAt        int64                  `json:"createdAt,string"`
	UpdatedAt        int64                  `json:"updatedAt,string"`
}

type MerchantGetDeliveryAddressVO struct {
	RecipientName string `json:"name,required" validate:"required"`
	Address       string `json:"address,required" validate:"required"`
	City          string `json:"city,required" validate:"required"`
	PhoneNumber   string `json:"phoneNumber,required" validate:"required"`
	AreaCode      string `json:"areaCode,required" validate:"required"`
	ZipCode       string `json:"zipCode,required" validate:"required"`
	Country       string `json:"country,required" validate:"required"`
}

type MerchantActivatePhysicalCardVO struct {
}

type MerchantSaveDeliveryAddressVO struct {
}

type MerchantUpdatePINVO struct {
}
