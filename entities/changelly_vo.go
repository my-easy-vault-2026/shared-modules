package entities

type ChangellyNotifyVO struct {
	Status string `json:"status,omitempty"`
}

type ChangellyOfferVO struct {
	ProviderCode     string `json:"providerCode"`
	ProviderIconUrl  string `json:"providerIconUrl"`
	OfferID          string `json:"offerId,omitempty"`
	AmountExpectedTo string `json:"amountExpectedTo"`
	Method           string `json:"method"`
	MethodName       string `json:"methodName"`
	Rate             string `json:"rate"`
	InvertedRate     string `json:"invertedRate"`
	Fee              string `json:"fee"`
}
