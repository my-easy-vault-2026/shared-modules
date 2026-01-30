package entities

// MailingAddressVo represents a mailing address with attributes typically necessary for postal documentation.
type MailingAddressVo struct {
	CardID        uint64 `json:"cardId,string,omitempty"`
	RecipientName string `json:"name,required"`
	Address       string `json:"address,required"`
	City          string `json:"city,required"`
	PhoneNumber   string `json:"phoneNumber,required"`
	AreaCode      string `json:"areaCode,required"`
	ZipCode       string `json:"zipCode,required"`
	Country       string `json:"country,required"`
}
