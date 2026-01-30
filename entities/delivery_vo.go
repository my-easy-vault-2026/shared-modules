package entities

import (
	"shared-modules/common"
	"time"
)

type DeliveryInfoVO struct {
	CardId        uint64                   `json:"cardId,string"`
	UserId        uint64                   `json:"userId,string"`
	Vendor        common.CardProductVendor `json:"vendor"`
	Status        common.DeliveryStatus    `json:"status"`
	RecipientName string                   `json:"recipientName"`
	Address       string                   `json:"address"`
	City          string                   `json:"city"`
	PhoneNumber   string                   `json:"phoneNumber"`
	AreaCode      string                   `json:"areaCode"`
	ZipCode       string                   `json:"zipCode"`
	Country       string                   `json:"country"`
	TrackingLink  string                   `json:"trackingLink"`
	TrackingCode  string                   `json:"trackingCode"`
	ShippedAt     *time.Time               `json:"shippedAt"`
	CreatedAt     *time.Time               `json:"createdAt"`
	UpdatedAt     *time.Time               `json:"updatedAt"`
}
