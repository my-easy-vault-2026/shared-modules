package entities

import (
	"shared-modules/common"
	"shared-modules/utils"

	"github.com/shopspring/decimal"
)

type MerchantUserLimitsVO struct {
	ID                  uint64                `json:"id,string"`           // 主鍵
	UserID              uint64                `json:"userID,string"`       // 用戶ID，與用戶表關聯
	Email               string                `json:"email"`               // 信箱
	LimitType           common.UserLimitsType `json:"limitType"`           // 限額類型：卡片限額或轉帳限額
	MonthlyTotal        decimal.Decimal       `json:"monthlyTotal"`        // 每月總交易量限額
	CurrentMonthlyTotal decimal.Decimal       `json:"currentMonthlyTotal"` // 當前每月總交易量
	DailyTotal          decimal.Decimal       `json:"dailyTotal"`          // 每日總交易量限額
	CurrentDailyTotal   decimal.Decimal       `json:"currentDailyTotal"`   // 當前每日總交易量
	DailyCount          int                   `json:"dailyCount"`          // 每日消費次數限額
	CurrentDailyCount   int                   `json:"currentDailyCount"`   // 當前每日消費次數
	MonthlyCount        int                   `json:"monthlyCount"`        // 每月消費次數限額
	CurrentMonthlyCount int                   `json:"currentMonthlyCount"` // 當前每月消費次數
	PerTransaction      decimal.Decimal       `json:"perTransaction"`      // 每筆交易金額限額
	MonthlyTotalMax     decimal.Decimal       `json:"monthlyTotalMax"`     //每月交易限額上限
	DailyTotalMax       decimal.Decimal       `json:"dailyTotalMax"`       //每日交易限額上限
	PerTransactionMax   decimal.Decimal       `json:"perTransactionMax"`   //單次交易金額上限
}

type MerchantPageUserLimitsDataVO struct {
	ID                  uint64                `json:"id,string"`           // 主鍵
	UserID              uint64                `json:"userID,string"`       // 用戶ID，與用戶表關聯
	Email               string                `json:"email"`               // 信箱
	LimitType           common.UserLimitsType `json:"limitType"`           // 限額類型：卡片限額或轉帳限額
	MonthlyTotal        decimal.Decimal       `json:"monthlyTotal"`        // 每月總交易量限額
	CurrentMonthlyTotal decimal.Decimal       `json:"currentMonthlyTotal"` // 當前每月總交易量
	DailyTotal          decimal.Decimal       `json:"dailyTotal"`          // 每日總交易量限額
	CurrentDailyTotal   decimal.Decimal       `json:"currentDailyTotal"`   // 當前每日總交易量
	DailyCount          int                   `json:"dailyCount"`          // 每日消費次數限額
	CurrentDailyCount   int                   `json:"currentDailyCount"`   // 當前每日消費次數
	MonthlyCount        int                   `json:"monthlyCount"`        // 每月消費次數限額
	CurrentMonthlyCount int                   `json:"currentMonthlyCount"` // 當前每月消費次數
	PerTransaction      decimal.Decimal       `json:"perTransaction"`      // 每筆交易金額限額
}

type MerchantCreateUserVO struct {
	ID uint64 `json:"id,string"`
}

type MerchantPageUserLimitsVO struct {
	utils.PageData[[]*MerchantPageUserLimitsDataVO]
}

type MerchantUpdateUserProfileVO struct {
	Success bool `json:"success"`
}

type MerchantUserProfileVO struct {
	Gender           string                      `json:"gender,omitempty"`
	CountryCode      int                         `json:"countryCode,omitempty"`
	PhoneNumber      string                      `json:"phoneNumber,omitempty"`
	FirstName        string                      `json:"firstName,omitempty"`
	LastName         string                      `json:"lastName,omitempty"`
	Dob              string                      `json:"dob,omitempty"`
	IdDocumentType   common.IdentityDocumentType `json:"idDocumentType,omitempty"`
	IdDocumentNumber string                      `json:"idDocumentNumber,omitempty"`
	AddressLine1     string                      `json:"addressLine1,omitempty"`
	AddressLine2     string                      `json:"addressLine2,omitempty"`
	NationCode       string                      `json:"nationCode,omitempty"`
	PostalCode       string                      `json:"postalCode,omitempty"`
	City             string                      `json:"city,omitempty"`
	KYCLevel         string                      `json:"kycLevel,omitempty"`
}

type MerchantPageUserVO struct {
	utils.PageData[[]string]
}
