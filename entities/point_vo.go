package entities

import (
	"shared-modules/utils"

	"github.com/shopspring/decimal"
)

type RewardProductVO struct {
	ID              uint64          `json:"id,string"`
	Type            string          `json:"type"`
	Name            string          `json:"name"`
	Code            string          `json:"code"`
	Summary         string          `json:"summary"`
	Description     interface{}     `json:"description"`
	CategoryCode    string          `json:"categoryCode"`
	VendorCode      string          `json:"vendorCode"`
	VendorName      string          `json:"vendorName"`
	Point           decimal.Decimal `json:"point"`
	PointCurrency   string          `json:"pointCurrency"`
	PointCardLevel  int8            `json:"pointCardLevel"`
	Price           decimal.Decimal `json:"price,omitempty"`
	PriceCurrency   string          `json:"priceCurrency,omitempty"`
	Quantity        *int            `json:"quantity,omitempty"`
	RedemptionTimes int             `json:"redemptionTimes"`
	RedemptionLimit *int            `json:"redemptionLimit,omitempty"`
	CoverURL        string          `json:"coverUrl"`
	ImageURL        string          `json:"imageUrl"`
	IconURL         string          `json:"iconUrl"`
	VendorIconURL   string          `json:"vendorIconUrl"`
	Restriction     string          `json:"restriction"`
	Area            string          `json:"area,omitempty"`
	DisplayOrder    int             `json:"displayOrder"`
	Status          string          `json:"status"`
	StartedAt       int64           `json:"startedAt,string,omitempty"`
	EndedAt         int64           `json:"endedAt,string,omitempty"`
	CreatedAt       int64           `json:"createdAt,string"`
	UpdatedAt       int64           `json:"updatedAt,string"`
}

type ListRewardProductVO struct {
	Records []*RewardProductVO `json:"records"`
}

type PageRewardProductVO struct {
	utils.PageData[[]*RewardProductVO]
}

type RedeemProductVO struct {
	Message string `json:"message"`
}

type RewardOrderVO struct {
	ID             uint64          `json:"id,string"`
	UserID         uint64          `json:"userId,string"`
	ProductID      uint64          `json:"productId,string"`
	Type           string          `json:"type"`
	Name           string          `json:"name"`
	Code           string          `json:"code"`
	Description    interface{}     `json:"description"`
	CategoryCode   string          `json:"categoryCode"`
	VendorCode     string          `json:"vendorCode"`
	VendorName     string          `json:"vendorName"`
	CoverURL       string          `json:"coverUrl"`
	ImageURL       string          `json:"imageUrl"`
	IconURL        string          `json:"iconUrl"`
	OrderNO        string          `json:"orderNo"`
	CardCategoryID uint64          `json:"cardCategoryId,string"`
	CardCurrency   string          `json:"cardCurrency"`
	CardID         uint64          `json:"cardId,string"`
	Point          decimal.Decimal `json:"point"`
	ShippingID     string          `json:"shippingId,omitempty"`
	RecipientName  string          `json:"recipientName,omitempty"`
	CountryCode    int             `json:"countryCode,omitempty"`
	PhoneNumber    string          `json:"phoneNumber,omitempty"`
	AddressLine1   string          `json:"addressLine1,omitempty"`
	AddressLine2   string          `json:"addressLine2,omitempty"`
	NationCode     string          `json:"nationCode,omitempty"`
	PostalCode     string          `json:"postalCode,omitempty"`
	City           string          `json:"city,omitempty"`
	CourierID      string          `json:"courierId,omitempty"`
	CourierName    string          `json:"courierName,omitempty"`
	TrackingNumber string          `json:"trackingNumber,omitempty"`
	TrackingURL    string          `json:"trackingUrl,omitempty"`
	ShippingStatus string          `json:"shippingStatus,omitempty"`
	Remark         string          `json:"remark"`
	Status         string          `json:"status"`
	ShippedAt      int64           `json:"shippedAt,string,omitempty"`
	MailedAt       int64           `json:"mailedAt,string,omitempty"`
	CreatedAt      int64           `json:"createdAt,string"`
	UpdatedAt      int64           `json:"updatedAt,string"`
}

type PageRewardOrderVO struct {
	utils.PageData[[]*RewardOrderVO]
}
