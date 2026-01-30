package entities

import (
	"shared-modules/common"
	"shared-modules/utils"

	"github.com/shopspring/decimal"
)

type MerchantCardApplyForm struct {
	Data string `json:"data"`
	Sign string `json:"sign"`
}

type MerchantCardApplyDataForm struct {
	FromCategory     string                      `json:"fromCategory"`
	CategoryID       uint64                      `json:"categoryID,string"`
	Email            string                      `json:"email" validate:"required"`
	CountryCode      int                         `json:"countryCode"`
	PhoneNumber      string                      `json:"phoneNumber"`
	FirstName        string                      `json:"firstName"`
	LastName         string                      `json:"lastName"`
	Dob              string                      `json:"dob"`          // 出生日期 YYYY-MM-DD
	AddressLine1     string                      `json:"addressLine1"` // 街道
	AddressLine2     string                      `json:"addressLine2"` // 楼层单元号
	NationCode       string                      `json:"nationCode"`
	City             string                      `json:"city"` // 城市
	IdDocumentType   common.IdentityDocumentType `json:"idDocumentType"`
	IdDocumentNumber string                      `json:"idDocumentNumber"`
}

type MerchantUpdateForwardTypeForm struct {
	CardID   uint64 `json:"cardId,string" validate:"required"`
	UserID   uint64 `json:"userId,string" validate:"required"`
	REAP     *bool  `json:"reap"`
	Email    *bool  `json:"email"`
	SMS      *bool  `json:"sms"`
	APP      *bool  `json:"app"`
	Merchant *bool  `json:"merchant"`
}

type MerchantAnswer3DSForm struct {
	CardID  uint64 `json:"cardId,string" validate:"required"`
	ID      uint64 `json:"Id,string" validate:"required"`
	Approve *bool  `json:"approve" validate:"required"`
}

type MerchantGetPrivacyInfoForm struct {
	ID uint64 `json:"id,string" validate:"required"`
}

type MerchantCardFreezeForm struct {
	CardId uint64 `json:"cardId,string" validate:"required"`
	Freeze *bool  `json:"freeze" validate:"required"`
}

type MerchantCardBlockForm struct {
	CardId uint64 `json:"cardId,string" validate:"required"`
	Block  *bool  `json:"block" validate:"required"`
	Reason string `json:"reason"`
}

type MerchantPageCardForm struct {
	ID         uint64 `json:"id,string"`
	Email      string `json:"email"`
	CategoryID uint64 `json:"categoryId,string"`
	UserID     uint64 `json:"userId,string"`
	IssueID    string `json:"issueId"`
	Currency   string `json:"currency"`
	utils.Page
}

type MerchantGetUserCardForm struct {
	CardID uint64 `json:"cardId,string"`
}

type UpdateForwardTypeForm struct {
	CardID   uint64 `json:"cardId,string" validate:"required"`
	REAP     *bool  `json:"reap"`
	Email    *bool  `json:"email"`
	SMS      *bool  `json:"sms"`
	APP      *bool  `json:"app"`
	Merchant *bool  `json:"merchant"`
}

type MerchantGetCardForm struct {
	ID uint64 `json:"id,string" validate:"required"`
}

type MerchantDeleteCardForm struct {
	CardID uint64 `json:"cardId,string"`
	Remark string `json:"remark"`
}

type MerchantCheck3dsForm struct {
}

type MerchantAdjustBalanceForm struct {
	CardID uint64          `json:"cardId,string" validate:"required"`
	Amount decimal.Decimal `json:"amount" validate:"required"`
}

type MerchantCardCategoryVO struct {
	ID                uint64           `json:"id,string"`
	Name              string           `json:"name"`
	PreferredName     string           `json:"preferredName,omitempty"`
	SecondaryName     string           `json:"secondaryName,omitempty"`
	Type              string           `json:"type"`
	Currency          string           `json:"currency"`
	CurrencyType      string           `json:"currencyType"`
	Organization      string           `json:"organization,omitempty"`
	AnnualFee         string           `json:"annualFee,omitempty"`
	ActivationDeposit string           `json:"activationDeposit,omitempty"`
	Format            string           `json:"format,omitempty"`
	SpendLimit        *decimal.Decimal `json:"spendLimit,omitempty"`
	CreatedAt         int64            `json:"createdAt,string"`
	UpdatedAt         int64            `json:"updatedAt,string"`
}

type ListMerchantCardCategoryVO struct {
	Records []*MerchantCardCategoryVO `json:"records"`
}

type MerchantActivatePhysicalCardForm struct {
	CardID         uint64 `json:"cardId,string" validate:"required"`
	ActivationCode string `json:"activationCode" validate:"required"`
	PIN            string `json:"pin"`
}

type MerchantSaveDeliveryAddressForm struct {
	CardID        uint64 `json:"cardId,string" validate:"required"`
	RecipientName string `json:"name" validate:"required"`
	Address       string `json:"address" validate:"required"`
	City          string `json:"city" validate:"required"`
	PhoneNumber   string `json:"phoneNumber" validate:"required"`
	AreaCode      string `json:"areaCode" validate:"required"`
	ZipCode       string `json:"zipCode" validate:"omitempty"`
	Country       string `json:"country" validate:"required"`
}

type MerchantGetDeliveryAddressForm struct {
	CardID uint64 `json:"cardId,string" validate:"required"`
}

type MerchantUpdatePINForm struct {
	CardID uint64 `json:"cardId,string" validate:"required"`
	PIN    string `json:"pin,required" validate:"required,digit"`
}
