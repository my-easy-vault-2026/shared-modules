package entities

type TopDownPreviewVO struct {
	FromAmount     string            `json:"fromAmount"`
	FromCardID     uint64            `json:"fromCardID,string"`
	FromCategory   string            `json:"fromCategory"`
	FromCategoryID uint64            `json:"fromCategoryID,string"`
	FromCurrency   string            `json:"fromCurrency"`
	ToAmount       string            `json:"toAmount"`
	ToCardID       uint64            `json:"toCardID,string"`
	ToCategory     string            `json:"toCategory"`
	ToCategoryID   uint64            `json:"toCategoryID,string"`
	ToCurrency     string            `json:"toCurrency"`
	Fee            string            `json:"fee"` // 轉出手續費
	ExchangeFee    string            `json:"exchangeFee,omitempty"`
	Rate           []*ExchangeRateVO `json:"rate"`
	DisplayRate    ExchangeRateVO    `json:"displayRate"`
	Key            string            `json:"key"`
	ExpiredAt      int64             `json:"expiredAt,string"`
}

type TopDownConfirmVO struct{}
