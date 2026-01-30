package entities

import "shared-modules/utils"

type MerchantVO struct {
	ID                uint64 `json:"id"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	Domain            string `json:"domain,omitempty"`
	DefaultCategoryID uint64 `json:"defaultCategoryId"`
	DefaultApplyID    uint64 `json:"defaultApplyId"`
	CardID            uint64 `json:"cardId"`
	Forward           string `json:"forward"`
	BalanceType       string `json:"balanceType"`
	CallbackURL       string `json:"callbackUrl,omitempty"` //可能為空
	Status            string `json:"status"`
	CountryCode       int    `json:"countryCode"`
	PhoneNumber       string `json:"phoneNumber"`
	CreatedAt         int64  `json:"createdAt"`
	UpdatedAt         int64  `json:"updatedAt"`
}
type PageMerchantVO struct {
	utils.PageData[[]*MerchantVO]
}
