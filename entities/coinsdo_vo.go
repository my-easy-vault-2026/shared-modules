package entities

type CoinsdoNotifyVO struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type WithdrawNotifyVO struct {
	BusinessID string `json:"businessId"`
}

type CoinsendDeviceInfoResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    Data   `json:"data"`
}

// Data 包含回應中的主要數據
type Data struct {
	OnlineStatus       int               `json:"onlineStatus"`
	UseStatus          int               `json:"useStatus"`
	AddressList        []CoinsendAddress `json:"addressList"`
	DeviceExpiryTime   string            `json:"deviceExpiryTime"`
	DeviceRegisterTime string            `json:"deviceRegisterTime"`
	DevicePublicKey    string            `json:"devicePublicKey"`
}

// Address 表示地址列表中的每一個項目
type CoinsendAddress struct {
	Currency     string `json:"currency"`
	Flag         string `json:"flag"`
	CoinsDoId    int    `json:"coinsDoId"`
	CurrencyName string `json:"currencyName"`
	Address      string `json:"address"`
	Balance      string `json:"balance"`
	AddrIndex    int    `json:"addrIndex"`
}
