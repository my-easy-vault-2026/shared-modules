package entities

type WithdrawVO struct {
	ID             uint64 `json:"id,string"`
	OrderNO        string `json:"orderNo"`
	UserID         uint64 `json:"userId,string"`
	ToAmount       string `json:"toAmount"`
	ToAddress      string `json:"toAddress"`
	FromAmount     string `json:"fromAmount"`
	FromCardID     uint64 `json:"fromCardId,string"`
	FromCategoryID uint64 `json:"fromCategoryId,string"`
	FromCurrency   string `json:"fromCurrency"`
	FromAddress    string `json:"fromAddress"`
	Mainnet        string `json:"mainnet"`
	Protocol       string `json:"protocol"`
	WithdrawFee    string `json:"withdrawFee"`
}

type PassFirstReviewVO struct {
}
