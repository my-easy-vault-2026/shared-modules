package entities

import (
	"shared-modules/common"
	"shared-modules/utils"
)

type AgentUserPageForm struct {
	AgentID uint64 `json:"-"`
	UserID  uint64 `json:"userID,string"`
	utils.Page
}

type AgentTxReportForm struct {
	AgentID       uint64 `-`
	CreatedAtFrom int64  `json:"createdAtFrom,string" validate:"required"`
	CreatedAtTo   int64  `json:"createdAtTo,string" validate:"required"`
	utils.Page
}

type AgentTxReportExportForm struct {
	AgentID uint64 `-`
	ID      uint64 `json:"id,string" validate:"required"`
}

type AgentReportTaskForm struct {
	CreatedAtFrom string `json:"createdAtFrom"`
	CreatedAtTo   string `json:"createdAtTo"`
}

type AgentUpdatePasswordForm struct {
	AgentID     uint64 `-`
	NewPassword string `json:"newPassword" validate:"required"`
	OldPassword string `json:"oldPassword" validate:"required"`
}

type AgentCardPageForm struct {
	AgentID      uint64            `json:"-"`
	Email        string            `json:"email"`
	CardID       uint64            `json:"cardId,string"`
	UserID       uint64            `json:"userId,string"`
	DiscountCode string            `json:"promotionCode"`
	Status       common.CardStatus `json:"status"`
	utils.Page
}

type AgentDiscountCardUserGetForm struct {
	AgentID uint64 `json:"-"`
	UserID  uint64 `json:"userId,string"`
}
