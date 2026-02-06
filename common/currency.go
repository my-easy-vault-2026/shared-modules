package common

import "strings"

type Currency uint64 // 貨幣
const (
	// crypto
	CURRENCY_USDT  Currency = 1001
	CURRENCY_ETH   Currency = 1002
	CURRENCY_BTC   Currency = 1003
	CURRENCY_USDC  Currency = 1004
	CURRENCY_DAI   Currency = 1005
	CURRENCY_WBTC  Currency = 1006
	CURRENCY_TRX   Currency = 1007
	CURRENCY_ADA   Currency = 1008
	CURRENCY_BCH   Currency = 1009
	CURRENCY_DOGE  Currency = 1010
	CURRENCY_LTC   Currency = 1011
	CURRENCY_XRP   Currency = 1012
	CURRENCY_SOL   Currency = 1013
	CURRENCY_BNB   Currency = 1014
	CURRENCY_ETC   Currency = 1015
	CURRENCY_MATIC Currency = 1016
	CURRENCY_TON   Currency = 1017
	CURRENCY_AVAX  Currency = 1018

	// fiat
	CURRENCY_USD        Currency = 2001
	CURRENCY_EUR        Currency = 2002
	CURRENCY_JPY        Currency = 2003
	CURRENCY_GBP        Currency = 2004
	CURRENCY_CNY        Currency = 2005
	CURRENCY_CAD        Currency = 2006
	CURRENCY_AUD        Currency = 2007
	CURRENCY_CHF        Currency = 2008
	CURRENCY_HKD        Currency = 2009
	CURRENCY_SEK        Currency = 2010
	CURRENCY_NZD        Currency = 2011
	CURRENCY_MXN        Currency = 2012
	CURRENCY_SGD        Currency = 2013
	CURRENCY_NOK        Currency = 2014
	CURRENCY_KRW        Currency = 2015
	CURRENCY_TRY        Currency = 2016
	CURRENCY_INR        Currency = 2017
	CURRENCY_BRL        Currency = 2018
	CURRENCY_ZAR        Currency = 2019
	CURRENCY_RUB        Currency = 2020
	CURRENCY_TWD        Currency = 2021
	CURRENCY_DKK        Currency = 2022
	CURRENCY_PLN        Currency = 2023
	CURRENCY_ILS        Currency = 2024
	CURRENCY_AED        Currency = 2025
	CURRENCY_ARS        Currency = 2026
	CURRENCY_MYR        Currency = 2027
	CURRENCY_THB        Currency = 2028
	CURRENCY_SAR        Currency = 2029
	CURRENCY_CZK        Currency = 2030
	CURRENCY_UAH        Currency = 2031
	CURRENCY_PHP        Currency = 2032
	CURRENCY_MOP        Currency = 2033
	CURRENCY_EGP        Currency = 2034
	CURRENCY_MAD        Currency = 2035
	CURRENCY_KES        Currency = 2036
	CURRENCY_RON        Currency = 2037
	CURRENCY_BGN        Currency = 2038
	CURRENCY_VND        Currency = 2039
	CURRENCY_COP        Currency = 2040
	CURRENCY_OTHER_FIAT Currency = 2999
)

func (c Currency) String() string {
	switch c {
	// crypto
	case CURRENCY_USDT:
		return "usdt"
	case CURRENCY_ETH:
		return "eth"
	case CURRENCY_USD:
		return "usd"
	case CURRENCY_BTC:
		return "btc"
	case CURRENCY_USDC:
		return "usdc"
	case CURRENCY_DAI:
		return "dai"
	case CURRENCY_WBTC:
		return "wbtc"
	case CURRENCY_TRX:
		return "trx"
	case CURRENCY_ADA:
		return "ada"
	case CURRENCY_BCH:
		return "bch"
	case CURRENCY_DOGE:
		return "doge"
	case CURRENCY_LTC:
		return "ltc"
	case CURRENCY_XRP:
		return "xrp"
	case CURRENCY_SOL:
		return "sol"
	case CURRENCY_BNB:
		return "bnb"
	case CURRENCY_ETC:
		return "etc"
	case CURRENCY_MATIC:
		return "matic"
	case CURRENCY_TON:
		return "ton"
	case CURRENCY_AVAX:
		return "avax"
	// fiat
	case CURRENCY_EUR:
		return "eur"
	case CURRENCY_JPY:
		return "jpy"
	case CURRENCY_GBP:
		return "gbp"
	case CURRENCY_CNY:
		return "cny"
	case CURRENCY_CAD:
		return "cad"
	case CURRENCY_AUD:
		return "aud"
	case CURRENCY_CHF:
		return "chf"
	case CURRENCY_HKD:
		return "hkd"
	case CURRENCY_SEK:
		return "sek"
	case CURRENCY_NZD:
		return "nzd"
	case CURRENCY_MXN:
		return "mxn"
	case CURRENCY_SGD:
		return "sgd"
	case CURRENCY_NOK:
		return "nok"
	case CURRENCY_KRW:
		return "krw"
	case CURRENCY_TRY:
		return "try"
	case CURRENCY_INR:
		return "inr"
	case CURRENCY_BRL:
		return "brl"
	case CURRENCY_ZAR:
		return "zar"
	case CURRENCY_RUB:
		return "rub"
	case CURRENCY_TWD:
		return "twd"
	case CURRENCY_DKK:
		return "dkk"
	case CURRENCY_PLN:
		return "pln"
	case CURRENCY_ILS:
		return "ils"
	case CURRENCY_AED:
		return "aed"
	case CURRENCY_ARS:
		return "ars"
	case CURRENCY_MYR:
		return "myr"
	case CURRENCY_THB:
		return "thb"
	case CURRENCY_SAR:
		return "sar"
	case CURRENCY_CZK:
		return "czk"
	case CURRENCY_UAH:
		return "uah"
	case CURRENCY_PHP:
		return "php"
	case CURRENCY_MOP:
		return "mop"
	case CURRENCY_EGP:
		return "egp"
	case CURRENCY_MAD:
		return "mad"
	case CURRENCY_KES:
		return "kes"
	case CURRENCY_RON:
		return "ron"
	case CURRENCY_BGN:
		return "bgn"
	case CURRENCY_VND:
		return "vnd"
	case CURRENCY_COP:
		return "cop"
	case CURRENCY_OTHER_FIAT:
		return "other_fiat"
	}
	return ""
}

func (c Currency) IsValid() bool {
	switch c {
	case
		// crypto
		CURRENCY_USDT,
		CURRENCY_ETH,
		CURRENCY_BTC,
		CURRENCY_USDC,
		CURRENCY_DAI,
		CURRENCY_WBTC,
		CURRENCY_TRX,
		CURRENCY_ADA,
		CURRENCY_BCH,
		CURRENCY_DOGE,
		CURRENCY_LTC,
		CURRENCY_XRP,
		CURRENCY_SOL,
		CURRENCY_BNB,
		CURRENCY_ETC,
		CURRENCY_MATIC,
		// fiat
		CURRENCY_USD,
		CURRENCY_EUR,
		CURRENCY_JPY,
		CURRENCY_GBP,
		CURRENCY_CNY,
		CURRENCY_CAD,
		CURRENCY_AUD,
		CURRENCY_CHF,
		CURRENCY_HKD,
		CURRENCY_SEK,
		CURRENCY_NZD,
		CURRENCY_MXN,
		CURRENCY_SGD,
		CURRENCY_NOK,
		CURRENCY_KRW,
		CURRENCY_TRY,
		CURRENCY_INR,
		CURRENCY_BRL,
		CURRENCY_ZAR,
		CURRENCY_RUB,
		CURRENCY_TWD,
		CURRENCY_DKK,
		CURRENCY_PLN,
		CURRENCY_ILS,
		CURRENCY_AED,
		CURRENCY_ARS,
		CURRENCY_MYR,
		CURRENCY_THB,
		CURRENCY_SAR,
		CURRENCY_CZK,
		CURRENCY_UAH,
		CURRENCY_PHP,
		CURRENCY_MOP,
		CURRENCY_EGP,
		CURRENCY_MAD,
		CURRENCY_KES,
		CURRENCY_RON,
		CURRENCY_BGN,
		CURRENCY_VND,
		CURRENCY_COP,
		CURRENCY_OTHER_FIAT:
		return true
	}
	return false
}

func (Currency) FromString(s string) Currency {
	sLower := strings.ToLower(s)
	switch sLower {
	// crypto
	case "usdt":
		return CURRENCY_USDT
	case "eth":
		return CURRENCY_ETH
	case "btc":
		return CURRENCY_BTC
	case "usdc":
		return CURRENCY_USDC
	case "dai":
		return CURRENCY_DAI
	case "wbtc":
		return CURRENCY_WBTC
	case "trx":
		return CURRENCY_TRX
	case "ada":
		return CURRENCY_ADA
	case "bch":
		return CURRENCY_BCH
	case "doge":
		return CURRENCY_DOGE
	case "ltc":
		return CURRENCY_LTC
	case "xrp":
		return CURRENCY_XRP
	case "sol":
		return CURRENCY_SOL
	case "bnb":
		return CURRENCY_BNB
	case "etc":
		return CURRENCY_ETC
	case "matic":
		return CURRENCY_MATIC
	case "ton":
		return CURRENCY_TON
	case "avax":
		return CURRENCY_AVAX

	// fiat
	case "usd":
		return CURRENCY_USD
	case "eur":
		return CURRENCY_EUR
	case "jpy":
		return CURRENCY_JPY
	case "gbp":
		return CURRENCY_GBP
	case "cny":
		return CURRENCY_CNY
	case "cad":
		return CURRENCY_CAD
	case "aud":
		return CURRENCY_AUD
	case "chf":
		return CURRENCY_CHF
	case "hkd":
		return CURRENCY_HKD
	case "sek":
		return CURRENCY_SEK
	case "nzd":
		return CURRENCY_NZD
	case "mxn":
		return CURRENCY_MXN
	case "sgd":
		return CURRENCY_SGD
	case "nok":
		return CURRENCY_NOK
	case "krw":
		return CURRENCY_KRW
	case "try":
		return CURRENCY_TRY
	case "inr":
		return CURRENCY_INR
	case "brl":
		return CURRENCY_BRL
	case "zar":
		return CURRENCY_ZAR
	case "rub":
		return CURRENCY_RUB
	case "twd":
		return CURRENCY_TWD
	case "dkk":
		return CURRENCY_DKK
	case "pln":
		return CURRENCY_PLN
	case "ils":
		return CURRENCY_ILS
	case "aed":
		return CURRENCY_AED
	case "ars":
		return CURRENCY_ARS
	case "myr":
		return CURRENCY_MYR
	case "thb":
		return CURRENCY_THB
	case "sar":
		return CURRENCY_SAR
	case "czk":
		return CURRENCY_CZK
	case "uah":
		return CURRENCY_UAH
	case "php":
		return CURRENCY_PHP
	case "mop":
		return CURRENCY_MOP
	case "egp":
		return CURRENCY_EGP
	case "mad":
		return CURRENCY_MAD
	case "kes":
		return CURRENCY_KES
	case "ron":
		return CURRENCY_RON
	case "bgn":
		return CURRENCY_BGN
	case "vnd":
		return CURRENCY_VND
	case "cop":
		return CURRENCY_COP
	case "other_fiat":
		return CURRENCY_OTHER_FIAT
	}
	return 0
}

func (c Currency) AssetType() AssetType {
	switch {
	case c >= 1000 && c < 2000:
		return ASSET_TYPE_CRYPTO
	case c >= 2000 && c < 3000:
		return ASSET_TYPE_FIAT
	case c >= 10000 && c < 20000:
		return ASSET_TYPE_POINT
	}
	return 0
}
