package entities

import (
	"shared-modules/common"
	"shared-modules/utils"

	"github.com/shopspring/decimal"
)

type MerchantUpdateUserLimitsForm struct {
	ID             uint64           `json:"id,string"`
	Email          string           `json:"email"`
	MonthlyTotal   *decimal.Decimal `json:"monthlyTotal,string"`   // 每月總交易量限額
	DailyTotal     *decimal.Decimal `json:"dailyTotal,string"`     // 每日總交易量限額
	PerTransaction *decimal.Decimal `json:"perTransaction,string"` // 每筆交易金額限額
}

type MerchantGetUserLimitsForm struct {
	UserID uint64 `json:"userId,string"`
	Email  string `json:"email"`
}

type MerchantPageUserLimitsForm struct {
	utils.Page
}

type MerchantCreateUserForm struct {
	Email            string                      `json:"email" validate:"required"`
	CountryCode      int                         `json:"countryCode" validate:"required"`
	PhoneNumber      string                      `json:"phoneNumber" validate:"required"`
	Gender           string                      `json:"gender" validate:"required"`
	FirstName        string                      `json:"firstName" validate:"required"`
	LastName         string                      `json:"lastName" validate:"required"`
	Dob              string                      `json:"dob" validate:"required"` // 出生日期 YYYY-MM-DD
	IdDocumentType   common.IdentityDocumentType `json:"idDocumentType" validate:"required"`
	IdDocumentNumber string                      `json:"idDocumentNumber" validate:"required"`
	AddressLine1     string                      `json:"addressLine1" validate:"required"` // 街道
	AddressLine2     string                      `json:"addressLine2" validate:"required"` // 楼层单元号
	NationCode       string                      `json:"nationCode" validate:"required"`
	City             string                      `json:"city" validate:"required"` // 城市
	PostalCode       string                      `json:"postalCode" validate:"omitempty,digit"`
}

type MerchantUpdateUserProfileForm struct {
	Email            string                      `json:"email" validate:"required"`
	FirstName        string                      `json:"firstName"`
	LastName         string                      `json:"lastName"`
	Dob              string                      `json:"dob"` // 出生日期 YYYY-MM-DD
	IdDocumentType   common.IdentityDocumentType `json:"idDocumentType"`
	IdDocumentNumber string                      `json:"idDocumentNumber"`
	Gender           string                      `json:"gender"`
	CountryCode      int                         `json:"countryCode"`
	PhoneNumber      string                      `json:"phoneNumber"`
	AddressLine1     string                      `json:"addressLine1"` // 街道
	AddressLine2     string                      `json:"addressLine2"` // 楼层单元号
	NationCode       string                      `json:"nationCode"`
	City             string                      `json:"city"` // 城市
	PostalCode       string                      `json:"postalCode" validate:"omitempty,digit"`
}

type MerchantGetUserProfileForm struct {
	Email string `json:"email" validate:"required"`
}

type MerchantPageUserForm struct {
	utils.Page
}

type MerchantUserApplyKycForm struct {
	Email       string `json:"email" validate:"required"`
	RedirectURL string `json:"redirectUrl" validate:"required"`
}
