package entities

type TopUpPreviewVO struct {
	FromAmount      string            `json:"fromAmount"`
	FromCardID      uint64            `json:"fromCardID,string"`
	FromCategory    string            `json:"fromCategory"`
	FromCategoryID  uint64            `json:"fromCategoryID,string"`
	FromCurrency    string            `json:"fromCurrency"`
	ToAmount        string            `json:"toAmount"`
	ToCardID        uint64            `json:"toCardID,string"`
	ToCategory      string            `json:"toCategory"`
	ToCategoryID    uint64            `json:"toCategoryID,string"`
	ToCurrency      string            `json:"toCurrency"`
	Fee             string            `json:"fee"` // 充值手續費
	ExchangeFee     string            `json:"exchangeFee"`
	Rate            []*ExchangeRateVO `json:"rate"`
	DisplayRate     ExchangeRateVO    `json:"displayRate"`
	Key             string            `json:"key"`
	ExpiredAt       int64             `json:"expiredAt,string"`
	PointAmount     string            `json:"pointAmount"`
	PointCardID     uint64            `json:"pointCardID"`
	PointCategory   string            `json:"pointCategory"`
	PointCategoryID uint64            `json:"pointCategoryID"`
	PointCurrency   string            `json:"pointCurrency"`
}

type TopUpDappPreviewVO struct {
	ContractAddress string            `json:"contractAddress"` // 鏈上幣種合約地址
	Decimals        int               `json:"decimals"`        // 鏈上交易精度
	WalletAddress   string            `json:"walletAddress"`
	FromAmount      string            `json:"fromAmount"`
	FromCardID      uint64            `json:"fromCardID,string"`
	FromCategory    string            `json:"fromCategory"`
	FromCategoryID  uint64            `json:"fromCategoryID,string"`
	FromCurrency    string            `json:"fromCurrency"`
	ToAmount        string            `json:"toAmount"`
	ToCardID        uint64            `json:"toCardID,string"`
	ToCategory      string            `json:"toCategory"`
	ToCategoryID    uint64            `json:"toCategoryID,string"`
	ToCurrency      string            `json:"toCurrency"`
	Fee             string            `json:"fee"` // 充值手續費
	ExchangeFee     string            `json:"exchangeFee,omitempty"`
	Rate            []*ExchangeRateVO `json:"rate"`
	DisplayRate     ExchangeRateVO    `json:"displayRate"`
	Key             string            `json:"key"`
	ExpiredAt       int64             `json:"expiredAt,string"`
	PointAmount     string            `json:"pointAmount"`
	PointCardID     uint64            `json:"pointCardID"`
	PointCategory   string            `json:"pointCategory"`
	PointCategoryID uint64            `json:"pointCategoryID"`
	PointCurrency   string            `json:"pointCurrency"`
}

type TopUpConfirmVO struct{}
