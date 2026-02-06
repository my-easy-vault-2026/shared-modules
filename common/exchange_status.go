package common

type ExchangeStatus int // 兌換狀態
const (
	EXCHANGE_STATUS_PENDING ExchangeStatus = 1
	EXCHANGE_STATUS_SUCCESS ExchangeStatus = 2
	EXCHANGE_STATUS_FAILED  ExchangeStatus = 3
)
