package entities

import (
	"shared-modules/utils"
)

type PageRewardProductForm struct {
	utils.Page
}

type PageRewardOrderForm struct {
	CardID uint64 `json:"cardId,string"`
	utils.Page
}

type ListRewardOrderForm struct {
	OrderNOIn []string
}

type GetRewardProductForm struct {
	ID   uint64 `json:"id"`
	Code string `json:"code"`
}

type ListRewardProductForm struct {
	IDIn []uint64 `json:"idIn"`
}

type RedeemRewardProductForm struct {
	ID uint64 `json:"id,string" validate:"required"`
}
