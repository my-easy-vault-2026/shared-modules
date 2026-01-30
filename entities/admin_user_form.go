package entities

import (
	"shared-modules/common"
	"shared-modules/utils"
)

type AdminUserQueryPageForm struct {
	Role            common.Role                `json:"role" validate:"required"`
	UserId          uint64                     `json:"userId,string"`
	UserIds         []uint64                   `json:"userIds,string"`
	Email           string                     `json:"email"`
	KycLevel        common.KYCLevel            `json:"kycLevel"`
	TerminateStatus common.UserTerminateStatus `json:"terminateStatus"`
	BlockStatus     common.UserBlockStatus     `json:"blockStatus"`
	NationCode      common.NationCode          `json:"nationCode"`
	ReferrerCode    string                     `json:"referrerCode"`
	OrderBy         string
	OrderDirection  common.OrderDirection
	utils.Page
}

type AdminUser struct {
	UserId uint64 `json:"userId,string" validate:"required"`
}

type AdminBlockUser struct {
	UserId      uint64 `json:"userId,string" validate:"required"`
	BlockReason string `json:"blockReason" validate:"required"`
}

type AdminUnblockUser struct {
	UserId uint64 `json:"userId,string" validate:"required"`
	CardID uint64 `json:"cardId,string"`
}

type AdminValidateUserForm struct {
	UserIds []string `json:"userIds"`
	Emails  []string `json:"emails"`
}

type AdminUserDownloadPageForm struct {
	Role            common.Role                `json:"role" validate:"required"`
	UserId          uint64                     `json:"userId,string"`
	Email           string                     `json:"email"`
	KycLevel        common.KYCLevel            `json:"kycLevel"`
	TerminateStatus common.UserTerminateStatus `json:"terminateStatus"`
	BlockStatus     common.UserBlockStatus     `json:"blockStatus"`
	NationCode      common.NationCode          `json:"nationCode"`
	ReferrerCode    string                     `json:"referrerCode"`
	Size            int                        `json:"size"`
}

type AdminDeleteUserAccountForm struct {
	UserId       uint64                  `json:"userId,string" validate:"required"`
	DeleteReason common.UserDeleteReason `json:"deleteReason" validate:"required"`
}

type AdminUserSaveReferenceCodeForm struct {
	UserId        uint64 `json:"userId,string" validate:"required"`
	ReferenceCode string `json:"referrerCode" validate:"required"`
}
