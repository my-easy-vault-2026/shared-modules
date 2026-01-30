package entities

type ChangellyCallbackPayload struct {
	RedirectUrl     string      `json:"redirectUrl"`
	OrderId         string      `json:"orderId"`
	Status          string      `json:"status"`
	ExternalUserId  string      `json:"externalUserId"`
	ExternalOrderId string      `json:"externalOrderId"`
	ProviderCode    string      `json:"providerCode"`
	CurrencyFrom    string      `json:"currencyFrom"`
	CurrencyTo      string      `json:"currencyTo"`
	AmountFrom      string      `json:"amountFrom"`
	Country         string      `json:"country"`
	State           string      `json:"state"`
	IP              string      `json:"ip"`
	WalletAddress   string      `json:"walletAddress"`
	WalletExtraId   string      `json:"walletExtraId"`
	PaymentMethod   string      `json:"paymentMethod"`
	UserAgent       string      `json:"userAgent"`
	Metadata        interface{} `json:"metadata"`
	CreatedAt       string      `json:"createdAt"`
	UpdatedAt       string      `json:"updatedAt"`
	PayinAmount     string      `json:"payinAmount"`
	PayoutAmount    string      `json:"payoutAmount"`
	PayinCurrency   string      `json:"payinCurrency"`
	PayoutCurrency  string      `json:"payoutCurrency"`
}
