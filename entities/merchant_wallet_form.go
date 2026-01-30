package entities

type MerchantGetWalletAddressForm struct {
	ID               uint64 `json:"id,string"`
	CardID           uint64 `json:"cardId,string"`
	CategoryID       uint64 `json:"categoryId,string" validate:"required"`
	Mainnet          string `json:"mainnet" validate:"required"`
	Protocol         string `json:"protocol"`
	Currency         string `json:"currency"`
	CreateIfNotExist bool   `json:"createIfNotExist"`
}
