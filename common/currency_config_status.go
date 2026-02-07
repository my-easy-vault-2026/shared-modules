package common

type CurrencyConfigStatus int //  狀態:1啟用 2關閉
const (
	CURRENCY_CONFIG_STATUS_ON  CurrencyConfigStatus = 1
	CURRENCY_CONFIG_STATUS_OFF CurrencyConfigStatus = 2
)
