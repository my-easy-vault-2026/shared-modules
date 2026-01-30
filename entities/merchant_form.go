package entities

import (
	"shared-modules/common"
	"shared-modules/utils"
)

type MerchantRequestForm struct {
	Data string `json:"data"`
	Sign string `json:"sign"`
}

type MerchantRegisterForm struct {
	UserID            uint64 `json:"userId,string" validate:"required"`
	Name              string `json:"name" validate:"required,max=20"`
	CardName          string `json:"cardName"`
	DefaultCategoryID uint64 `json:"defaultCategoryId,string" validate:"required"`
	APIKey            string `json:"apiKey"`
	BalanceType       string `json:"balanceType" validate:"required,oneof='merchant_centralized' 'merchant_decentralized'"`
	Design            string `json:"design" validate:"required"`
	CardBin           string `json:"cardBin" validate:"required"`
}

type MerchantCreateRegisterForm struct {
	Email             string                      `json:"email" validate:"required,email"`
	Name              string                      `json:"name" validate:"required,max=20"`
	Domain            string                      `json:"domain"`
	CardName          string                      `json:"cardName"`
	DefaultCategoryID uint64                      `json:"defaultCategoryId,string" validate:"required"`
	APIKey            string                      `json:"apiKey"`
	BalanceType       string                      `json:"balanceType" validate:"required,oneof='merchant_centralized' 'merchant_decentralized'" default:"merchant_decentralized"`
	FirstName         string                      `json:"firstName" validate:"required"`
	LastName          string                      `json:"lastName" validate:"required"`
	Dob               string                      `json:"dob" validate:"required"`            // 出生日期 YYYY-MM-DD
	IdDocumentType    common.IdentityDocumentType `json:"idDocumentType" validate:"required"` // 1:護照 2:健保 3:身份證 4:稅籍 5:社福 6:駕照
	IdDocumentNumber  string                      `json:"idDocumentNumber" validate:"required"`
	CountryCode       int                         `json:"countryCode" validate:"required"`
	PhoneNumber       string                      `json:"phoneNumber" validate:"required"`
	AddressLine1      string                      `json:"addressLine1" validate:"required"` // 街道
	AddressLine2      string                      `json:"addressLine2" validate:"required"` // 楼层单元号
	NationCode        string                      `json:"nationCode" validate:"required"`
	City              string                      `json:"city" validate:"required"` // 城市
	Design            string                      `json:"design" validate:"required"`
	CardBin           string                      `json:"cardBin" validate"required"`
}
type MerchantSetPublicKeyForm struct {
	PublicKey  string `json:"publicKey" validate:"required"`
	MerchantID uint64 `json:"merchantId,string" validate:"required"`
}

type SetWebhookForm struct {
	URL string `json:"url" validate:"required"`
}

type PageMerchantForm struct {
	Email       string `json:"email" validate:"omitempty,email"`
	Name        string `json:"name"`
	Domain      string `json:"domain" validate:"omitempty"`
	CountryCode int    `json:"countryCode"`
	PhoneNumber string `json:"phoneNumber" validate:"omitempty,numeric"`
	utils.Page
}

type CheckBalanceForm struct {
}
