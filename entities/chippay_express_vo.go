package entities

type ChippayExpressRequestVO struct {
	OrderNO string `json:"orderNo"`
	URL     string `json:"url"`
}

type ChippayExpressWebhookDataVO struct {
	OTCOrderNUM     string `json:"otcOrderNum"`
	CompanyOrderNUM string `json:"companyOrderNum"`
}

type ChippayExpressWebhookVO struct {
	Code    int                          `json:"code"`
	Data    *ChippayExpressWebhookDataVO `json:"data,omitempty"`
	Message string                       `json:"msg"`
}

type ChippayExpressGetProcessingOrderVO struct {
	Exists       bool   `json:"exists"`
	Frozen       bool   `json:"frozen"`
	CardID       uint64 `json:"card,string,omitempty"`
	OrderNO      string `json:"orderNo,omitempty"`
	RedirectLink string `json:"redirectLink,omitempty"`
	Message      string `json:"message"`
}
