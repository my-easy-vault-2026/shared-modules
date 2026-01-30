package entities

// UpdateMailingAddressForm contains parameters for updating an existing mailing address
type UpdateMailingAddressForm struct {
	ID            int    `json:"id,string"`
	CardID        uint64 `json:"cardId,string"`
	RecipientName string `json:"name,required"`
	Address       string `json:"address,required"`
	City          string `json:"city,required"`
	PhoneNumber   string `json:"phoneNumber,required"`
	AreaCode      string `json:"areaCode,required"`
	ZipCode       string `json:"zipCode"`
	Country       string `json:"country,required"`
}

type GetMailingAddressForm struct {
	CardID uint64 `json:"cardId,string"`
}
