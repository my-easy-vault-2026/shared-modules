package common

type CryptoCurrencyStatus int // crypto_currency 狀態:1啟用 2關閉
const (
	CRYPTO_CURRENCY_STATUS_ON  CryptoCurrencyStatus = 1
	CRYPTO_CURRENCY_STATUS_OFF CryptoCurrencyStatus = 2
)
