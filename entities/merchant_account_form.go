package entities

import (
	"shared-modules/common"
	"shared-modules/utils"
	"time"
)

type MerchantExportBalanceChangeForm struct {
	ExportToken string      `json:"exportToken" validate:"required"`
	StartTime   time.Time   `json:"startTime,omitempty"`
	EndTime     time.Time   `json:"endTime,omitempty"`
	MerchantID  uint64      `json:"merchantId,omitempty,string" validate:"required"`
	Role        common.Role `json:"role,omitempty"`
	UserIDs     []uint64    `json:"userIds,omitempty,string"`
	CardIDs     []uint64    `json:"cardIds,omitempty,string"`
	IncludeZero bool        `json:"includeZero"`
}

type MerchantRequestExportBalanceChangeForm struct {
	StartTime   int64    `json:"startTime"`
	EndTime     int64    `json:"endTime"`
	Role        string   `json:"role,omitempty"`
	UserIDs     []string `json:"userIds,omitempty"`
	CardIDs     []string `json:"cardIds,omitempty"`
	Emails      []string `json:"emails,omitempty"`
	IncludeZero bool     `json:"includeZero"`
}

type MerchantPageBillForm struct {
	Email         string          `json:"email,omitempty"`
	AssetIDIn     []uint64        `json:"assetIdIn,omitempty,string"`
	AssetID       uint64          `json:"assetId,omitempty,string"`
	UserID        uint64          `json:"userId,omitempty,string"`
	CategoryID    uint64          `json:"categoryId,omitempty,string"`
	Currency      common.Currency `json:"currency,omitempty"`
	IncludeZero   bool            `json:"includeZero"`
	CreatedAtFrom int64           `json:"createdAtFrom,string" validate:"required"`
	CreatedAtTo   int64           `json:"createdAtTo,string" validate:"required"`
	utils.Page
}
