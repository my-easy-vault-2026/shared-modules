package entities

import (
	"shared-modules/common"
	"shared-modules/utils"

	"github.com/shopspring/decimal"
)

type CalculateDiscountVO struct {
	Amount       string            `json:"amount"`
	FromDiscount string            `json:"fromDiscount,omitempty"`
	ToDiscount   string            `json:"toDiscount,omitempty"`
	Bonus        string            `json:"bonus,omitempty"`
	Rate         []*ExchangeRateVO `json:"rate"`
	Design       string            `json:"design,omitempty"`
}

type UserInviteSummaryVO struct {
	InviteCount      int64           `json:"inviteCount,string"`
	InviteEarnAmount decimal.Decimal `json:"inviteEarnAmount,string"`
}

type UserInviteRecordVO struct {
	Income    decimal.Decimal `json:"income,string"`
	Email     string          `json:"email"`
	CreatedAt int64           `json:"createdAt,string"`
}
type PageUserInviteRecordsVO struct {
	utils.PageData[[]*UserInviteRecordVO]
}

type AdminDiscountCodeVO struct {
	ID            uint64                      `json:"id,string"`
	Code          string                      `json:"code"`
	Discount      decimal.Decimal             `json:"discount,string"`
	DiscountType  common.DiscountType         `json:"discountType"`
	Target        []string                    `json:"target"`
	StartedAt     int64                       `json:"startedAt"`
	EndedAt       int64                       `json:"endedAt"`
	Status        common.ApplyPromotionStatus `json:"status"`
	Type          common.PromotionCodeType    `json:"type"`
	QuantityLimit int64                       `json:"quantityLimit"` // 限制使用次數
	QuantityUsed  int64                       `json:"quantityUsed"`  // 已使用次數
	CreatedAt     int64                       `json:"createdAt"`     // 創建時間
	UpdatedAt     int64                       `json:"updatedAt"`     // 更新時間
	CreateUser    uint64                      `json:"createUser,string"`
	UpdateUser    uint64                      `json:"updateUser,string"`
}

type AdminDiscountCodePage struct {
	utils.PageData[[]*AdminDiscountCodeVO]
}

type AdminDiscountCodeInfoVO struct {
	ID            uint64                      `json:"id,string"`
	Code          string                      `json:"code"`
	Discount      decimal.Decimal             `json:"discount,string"`
	DiscountType  common.DiscountType         `json:"discountType"`
	Target        []string                    `json:"target"`
	StartedAt     int64                       `json:"startedAt"`
	EndedAt       int64                       `json:"endedAt"`
	Status        common.ApplyPromotionStatus `json:"status"`
	Type          common.PromotionCodeType    `json:"type"`
	QuantityLimit int64                       `json:"quantityLimit"` // 限制使用次數
	QuantityUsed  int64                       `json:"quantityUsed"`  // 已使用次數
	CreatedAt     int64                       `json:"createdAt"`     // 創建時間
	UpdatedAt     int64                       `json:"updatedAt"`     // 更新時間
	CreateUser    uint64                      `json:"createUser,string"`
	UpdateUser    uint64                      `json:"updateUser,string"`
}
