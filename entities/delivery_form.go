package entities

import "shared-modules/common"

type GetDeliveryInfoForm struct {
	CardId uint64 `json:"cardId,string"`
}

type EditDeliveryInfoForm struct {
	CardId        uint64                `json:"cardId,string" validate:"required"`
	Status        common.DeliveryStatus `json:"status" validate:"required"`
	RecipientName string                `json:"recipientName" validate:"required"`
	Address       string                `json:"address" validate:"required"`
	City          string                `json:"city" validate:"required"`
	PhoneNumber   string                `json:"phoneNumber" validate:"required"`
	AreaCode      string                `json:"areaCode" validate:"required"`
	ZipCode       string                `json:"zipCode"`
	Country       string                `json:"country" validate:"required"`
	TrackingLink  string                `json:"trackingLink"`
	TrackingCode  string                `json:"trackingCode"`
}
