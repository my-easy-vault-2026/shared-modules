package entities

import (
	"shared-modules/common"
	"shared-modules/utils"

	"github.com/shopspring/decimal"
)

type CreateAdminUserForm struct {
	Name          string              `json:"name" validate:"required"`
	Password      string              `json:"password" validate:"required"`
	Level         common.AdminLevel   `json:"level" validate:"required"`
	Role          common.Role         `json:"role" validate:"required"`
	PromotionCode string              `json:"promotionCode" validate:"omitempty,min=4,max=20"`
	DiscountType  common.DiscountType `json:"discountType"`
	Discount      decimal.Decimal     `json:"discount"`
	GroupIds      []uint64            `json:"groupIds"`
}

type CreateMerchantAdminUserForm struct {
	Name       string `json:"name" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Level      string `json:"level" validate:"required"`
	MerchantID uint64 `json:"merchantId,string" validate:"required"`
}

type AdminQueryPageForm struct {
	Name    string                 `json:"name"`
	AdminID uint64                 `json:"adminID,string"`
	Level   common.AdminLevel      `json:"level"`
	Status  common.AdminUserStatus `json:"status"`
	Role    common.Role            `json:"role"`
	utils.Page
}

type AdminUpdatePasswordForm struct {
	AdminID     uint64 `json:"-"`
	NewPassword string `json:"newPassword" validate:"required"`
	OldPassword string `json:"oldPassword" validate:"required"`
}

type AdminQueryAccountPageForm struct {
	Name   string                 `json:"name"`
	ID     uint64                 `json:"id,string"`
	Level  common.AdminLevel      `json:"level"`
	Status common.AdminUserStatus `json:"status"`
	utils.Page
}

type AdminQueryAccountForm struct {
	ID uint64 `json:"id,string" validate:"required"`
}

type AdminCreateAccountForm struct {
	Name     string            `json:"name" validate:"required"`
	Password string            `json:"password" validate:"required"`
	Level    common.AdminLevel `json:"level" validate:"required"`
	GroupIds []uint64          `json:"groupIds,string"`
}

type AdminEditAccountForm struct {
	ID       uint64            `json:"id,string"  validate:"required"`
	Level    common.AdminLevel `json:"level"`
	GroupIds []uint64          `json:"groupIds,string"`
}

type AdminActivateAccountForm struct {
	ID uint64 `json:"id,string"  validate:"required"`
}

type AdminDeactivateAccountForm struct {
	ID uint64 `json:"id,string"  validate:"required"`
}

type AdminUpdateAccountPasswordForm struct {
	ID          uint64 `json:"id,string" validate:"required"`
	NewPassword string `json:"password" validate:"required"`
	//OldPassword string `json:"oldPassword" validate:"required"`
}

type AgentQueryAccountPageForm struct {
	Name          string                 `json:"name"`
	ID            uint64                 `json:"id,string"`
	Status        common.AdminUserStatus `json:"status"`
	InviteCode    string                 `json:"inviteCode"`
	PromotionCode string                 `json:"promotionCode"`
	utils.Page
}

type AgentQueryAccountForm struct {
	ID uint64 `json:"id,string" validate:"required"`
}

type AgentCreateAccountForm struct {
	Name          string              `json:"name" validate:"required"`
	Password      string              `json:"password" validate:"required"`
	Level         common.AdminLevel   `json:"level" validate:"required"`
	PromotionCode string              `json:"promotionCode" validate:"omitempty,min=4,max=20"`
	DiscountType  common.DiscountType `json:"discountType"`
	Discount      decimal.Decimal     `json:"discount"`
}

type AgentEditAccountForm struct {
	ID           uint64              `json:"id,string"`
	DiscountType common.DiscountType `json:"discountType" validate:"required"`
	Discount     decimal.Decimal     `json:"discount" validate:"required"`
}

type AgentActivateAccountForm struct {
	ID uint64 `json:"id,string"  validate:"required"`
}

type AgentDeactivateAccountForm struct {
	ID uint64 `json:"id,string"  validate:"required"`
}

type AgentUpdateAccountPasswordForm struct {
	ID          uint64 `json:"id,string" validate:"required"`
	NewPassword string `json:"password" validate:"required"`
}
