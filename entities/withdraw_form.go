package entities

import "github.com/shopspring/decimal"

type WithdrawForm struct {
	FromCardID uint64          `json:"fromCardId,string" validate:"required"`
	ToAddress  string          `json:"toAddress" validate:"required"`
	Mainnet    string          `json:"mainnet" validate:"required"`
	Protocol   string          `json:"protocol"`
	ToAmount   decimal.Decimal `json:"toAmount" validate:"required"`
}

type PassFirstReviewForm struct {
	DeviceID string `json:"deviceId" validate:"required"`
	OrderNO  string `json:"orderNo" validate:"required"`
	Sign     string `josn:"sign" validate:"required"`
	Remark   string `json:"remark"`
}

type CheckWithdrawOrderForm struct {
	LookbackMinutes       int `json:"lookbackMinutes" validate:"required"`
	AlertThresholdMinutes int `json:"alertThresholdMinutes" validate:"required"`
}
