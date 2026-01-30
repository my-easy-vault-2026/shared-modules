package entities

import (
	"shared-modules/common"
	"shared-modules/utils"

	"github.com/shopspring/decimal"
)

type CalculateDiscountForm struct {
	Code       string `json:"code" validate:"required"`
	CategoryID uint64 `json:"categoryId,string" validate:"required"`
	Currency   string `json:"currency" validate:"required"`
}

type VerifyInviteCodeForm struct {
	Code string `json:"code" validate:"required"`
}
type GetUserInviteRecordForm struct {
	utils.Page
}

type AdminDiscountCodePageForm struct {
	Code string `json:"code"`
	utils.Page
}

type AdminDiscountCodeInfoForm struct {
	Id uint64 `json:"id,string" validate:"required"`
}

type AdminDiscountCodeCreateForm struct {
	Code          string               `json:"code" validate:"required"`
	Discount      decimal.Decimal      `json:"discount,string"`
	DiscountType  common.DiscountType  `json:"discountType" validate:"required,oneof=1 2"`
	DiscountGroup common.DiscountGroup `json:"discountGroup"`
	Target        []string             `json:"target"`
	StartedAt     int64                `json:"startedAt"`
	EndedAt       int64                `json:"endedAt"`
	QuantityLimit int64                `json:"quantityLimit"` // 限制使用次數
}

type AdminDiscountCodeUpdateForm struct {
	ID            uint64               `json:"id,string" validate:"required"`
	Target        []string             `json:"target"`
	DiscountGroup common.DiscountGroup `json:"discountGroup"`
	StartedAt     int64                `json:"startedAt"`
	EndedAt       int64                `json:"endedAt"`
	QuantityLimit int64                `json:"quantityLimit"` // 限制使用次數
}
