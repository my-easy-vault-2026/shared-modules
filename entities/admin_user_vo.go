package entities

import (
	"shared-modules/common"
	"shared-modules/utils"
	"time"

	"github.com/shopspring/decimal"
)

type GetAdminInfoVO struct {
	ID               uint64                     `json:"id,string"`
	Email            string                     `json:"email"`
	FinishedIdentity string                     `json:"finishedIdentity,omitempty"`
	CountryCode      int                        `json:"countryCode,omitempty"` //可能為空
	PhoneNumber      string                     `json:"phoneNumber,omitempty"` //可能為空
	KycLevel         common.KYCLevel            `json:"kycLevel,omitempty"`
	KycStatus        common.KYCStatus           `json:"kycStatus,omitempty"`
	Gender           string                     `json:"gender,omitempty"`
	Channel          string                     `json:"channel,omitempty"`
	EPoint           decimal.Decimal            `json:"ePoint,omitempty"`
	EPointLevel      string                     `json:"ePointLevel,omitempty"`
	EPointCardID     uint64                     `json:"ePointCardId,string,omitempty"`
	Auto3DS          string                     `json:"auto3ds,omitempty"`
	AutoTopUp        string                     `json:"autoTopUp,omitempty"`
	ATMToggle        string                     `json:"atmToggle,omitempty"`
	CreatedAt        int64                      `json:"createdAt,string,omitempty"`
	UpdatedAt        int64                      `json:"updatedAt,string,omitempty"`
	DeletedAt        *time.Time                 `json:"deletedAt,string,omitempty"`
	Identity         IdentityVO                 `json:"identity,omitempty"`
	Device           DeviceInfo                 `json:"device,omitempty"`
	Language         common.Language            `json:"language,omitempty"`
	PromotionCode    string                     `json:"promotionCode,omitempty"`
	PromotionLink    string                     `json:"promotionLink,omitempty"`
	ReferrerCode     string                     `json:"referrerCode,omitempty"`
	BlockStatus      common.UserBlockStatus     `json:"blockStatus,omitempty"`
	BlockReason      common.UserBlockReason     `json:"blockReason,omitempty"`
	TerminateStatus  common.UserTerminateStatus `json:"terminateStatus,omitempty"`
	DeleteReason     common.UserDeleteReason    `json:"deleteReason,omitempty"`
	LastLoginAt      int64                      `json:"lastLoginAt,string,omitempty"`
	Role             common.Role                `json:"role,omitempty"`
	MerchantId       uint64                     `json:"merchantId,string,omitempty"`
	MerchantName     string                     `json:"merchantName,omitempty"`
}

type PageAdminUserVo struct {
	utils.PageData[[]*GetAdminInfoVO]
}

type AdminValidateUserVO []*AdminValidateUserVOUnit

type AdminValidateUserVOUnit struct {
	UserId    uint64 `json:"userId,string,omitempty"`
	Email     string `json:"email,omitempty"`
	ErrorCode int    `json:"errorCode"`
}
