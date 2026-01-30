package utils

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"shared-modules/common"
	"shared-modules/logger"

	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

type ExchangeRate struct {
	BaseCurrency  common.Currency
	QuoteCurrency common.Currency
	Rate          decimal.Decimal
	ActualRate    decimal.Decimal
	Timestamp     time.Time
	Purpose       common.RatePurpose
}

func ListExchangeRate(ctx context.Context, baseCurrency common.Currency, quoteCurrencies []common.Currency) ([]*ExchangeRate, error) {

	// 20241202 更新成全都用usd買賣價換算匯率
	return ListUsdBaseExchangeRate(ctx, baseCurrency, quoteCurrencies)

	base := strings.ToUpper(baseCurrency.String())
	redisKey := GetRateRedisKey(base)
	rates := make([]*ExchangeRate, len(quoteCurrencies))

	for i := range quoteCurrencies {
		quote := strings.ToUpper(quoteCurrencies[i].String())
		rate, err := RDB.HGet(ctx, redisKey, quote).Result()
		if err != nil {
			if err == redis.Nil {
				logger.Errorf("币种 %s 对 %s 的汇率不存在\n", base, quote)
			} else {
				logger.Warn("从 Redis 获取数据失败: ", err)
			}
		} else {
			// 將 JSON 字符串反序列化為 ExchangeRate 結構體
			var exchangeRate ExchangeRate
			err = json.Unmarshal([]byte(rate), &exchangeRate)
			if err != nil {
				logger.Error("GetExchangeRate JSON Unmarshal error: ", err)
				return nil, err
			}
			duration := time.Since(exchangeRate.Timestamp)
			if duration > 30*time.Second {
				continue
			}
			rates[i] = &exchangeRate
		}
	}
	return rates, nil
}

func ListUsdBaseExchangeRate(ctx context.Context, baseCurrency common.Currency, quoteCurrencies []common.Currency) ([]*ExchangeRate, error) {
	base := strings.ToUpper(baseCurrency.String())
	redisKey := GetUsdBaseRateRedisKey(base)
	var rates []*ExchangeRate

	for i := range quoteCurrencies {
		quote := strings.ToUpper(quoteCurrencies[i].String())
		rate, err := RDB.HGet(ctx, redisKey, quote).Result()
		if err != nil {
			if err == redis.Nil {
				logger.Errorf("币种 %s 对 %s 的汇率不存在\n", base, quote)
			} else {
				logger.Warn("从 Redis 获取数据失败: ", err)
			}
		} else {
			// 將 JSON 字符串反序列化為 ExchangeRate 結構體
			var exchangeRate ExchangeRate
			err = json.Unmarshal([]byte(rate), &exchangeRate)
			if err != nil {
				logger.Error("GetExchangeRate JSON Unmarshal error: ", err)
				return nil, err
			}
			duration := time.Since(exchangeRate.Timestamp)
			if duration > 30*time.Second {
				continue
			}
			rates = append(rates, &exchangeRate)
		}
	}
	return rates, nil
}

func GetUsdBaseExchangeRate(ctx context.Context, baseCurrency common.Currency, quoteCurrency common.Currency) (*ExchangeRate, error) {
	base := strings.ToUpper(baseCurrency.String())
	redisKey := GetUsdBaseRateRedisKey(base)

	// 從 Redis 中取得對應 BaseCurrency 的匯率資料
	rateData, err := RDB.HGet(ctx, redisKey, strings.ToUpper(quoteCurrency.String())).Result()
	if err != nil {
		if err == redis.Nil {
			logger.Warnf("No exchange rate found for %s to %s", baseCurrency, quoteCurrency)
			return nil, err
		}
		logger.Warn("GetExchangeRate Redis HGet error: ", err)
		return nil, err
	}

	// 將 JSON 字符串反序列化為 ExchangeRate 結構體
	var exchangeRate ExchangeRate
	err = json.Unmarshal([]byte(rateData), &exchangeRate)
	if err != nil {
		logger.Error("GetExchangeRate JSON Unmarshal error: ", err)
		return nil, err
	}

	duration := time.Since(exchangeRate.Timestamp)
	if duration > 30*time.Second {
		return nil, NewBusinessError(ctx, common.CODE_QUOTE_NO_SUCH_RATE)
	}

	return &exchangeRate, nil
}

func GetExchangeRate(ctx context.Context, baseCurrency common.Currency, quoteCurrency common.Currency) (*ExchangeRate, error) {

	results, err := ListUsdBaseExchangeRate(ctx, baseCurrency, []common.Currency{quoteCurrency})
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil
	}
	return results[0], nil

	base := strings.ToUpper(baseCurrency.String())
	redisKey := GetRateRedisKey(base)

	// 從 Redis 中取得對應 BaseCurrency 的匯率資料
	rateData, err := RDB.HGet(ctx, redisKey, strings.ToUpper(quoteCurrency.String())).Result()
	if err != nil {
		if err == redis.Nil {
			logger.Warnf("No exchange rate found for %s to %s", baseCurrency, quoteCurrency)
			return nil, err
		}
		logger.Warn("GetExchangeRate Redis HGet error: ", err)
		return nil, err
	}

	// 將 JSON 字符串反序列化為 ExchangeRate 結構體
	var exchangeRate ExchangeRate
	err = json.Unmarshal([]byte(rateData), &exchangeRate)
	if err != nil {
		logger.Error("GetExchangeRate JSON Unmarshal error: ", err)
		return nil, err
	}

	duration := time.Since(exchangeRate.Timestamp)
	if duration > 30*time.Second {
		return nil, NewBusinessError(ctx, common.CODE_QUOTE_NO_SUCH_RATE)
	}

	return &exchangeRate, nil
}

func ToBittopCurrency(currency string) string {
	switch currency {
	case "MATIC":
		return "POL"
	case "matic":
		return "pol"
	}
	return currency
}

func FromBittopCurrency(currency string) string {
	switch currency {
	case "POL":
		return "MATIC"
	case "pol":
		return "matic"
	}
	return currency
}
