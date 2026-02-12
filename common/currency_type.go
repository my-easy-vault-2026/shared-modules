package common

type CurrencyType int // 貨幣類型
const (
	CURRENCY_TYPE_CRYPTO CurrencyType = 1
	CURRENCY_TYPE_FIAT   CurrencyType = 2
	CURRENCY_TYPE_POINT  CurrencyType = 3
)

func (ct CurrencyType) String() string {
	switch ct {
	case CURRENCY_TYPE_CRYPTO:
		return "crypto"
	case CURRENCY_TYPE_FIAT:
		return "fiat"
	case CURRENCY_TYPE_POINT:
		return "point"
	}
	return ""
}
