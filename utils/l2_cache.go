package utils

import (
	"shared-modules/common"
	"shared-modules/logger"

	"context"
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type L2CacheConfig struct {
	Level         []common.L2CacheLevel
	ExpireSeconds time.Duration
}

type User struct {
	ID uint64
}

// 僅支援select單表查詢
func L2CQuery[T interface{}](ctx context.Context, db *gorm.DB, qf func(tx *gorm.DB) *gorm.DB, rf func(tx *gorm.DB) (T, error), config *L2CacheConfig) (T, error) {

	if config == nil {
		return *new(T), NewBusinessError(ctx, common.CODE_MISSING_PARAMETER)
	}

	sql := db.ToSQL(qf)
	sql = strings.ReplaceAll(sql, "`", "")
	split := strings.Split(sql, " ")
	table := "unknown"
	for i, s := range split {
		if strings.ToLower(s) == "from" {
			if len(split) > i+1 {
				table = split[i+1]
				break
			}
		}
	}

	switch ret := RDB.Get(ctx, GetL2CacheKey(table, sql)); true {
	case ret.Err() != nil && !errors.Is(ret.Err(), redis.Nil):
		return *new(T), ret.Err()
	case ret.Err() == nil && ret.Val() != "":
		t := *new(T)
		err := json.Unmarshal([]byte(ret.Val()), &t)
		if err == nil {
			return t, nil
		}
		logger.Warnf("unmarshal failed [%s][%s], %#v", GetL2CacheKey(table, sql), ret.Val(), err)
	case ret.Err() == nil && ret.Val() == "":
		return *new(T), nil
	case errors.Is(ret.Err(), redis.Nil):
	default:
		logger.Warnf("redis get failed [%s], %#v", GetL2CacheKey(table, sql), ret.Err())
	}

	ret := qf(db)

	reqID, _ := GetMDCValue("reqId")
	result, err := rf(ret)
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return *new(T), NewBusinessError(ctx, common.CODE_INVALID_PARAMETER)
	}
	resultCopy := deepCopy(result)
	go func(res T) {
		value := ""
		SetMDCValue("reqId", reqID)
		if err == nil && !reflect.ValueOf(res).IsZero() {
			resJSON, err := json.Marshal(res)
			if err != nil {
				logger.Warnf("marshal failed: [%s], ", GetL2CacheKey(table, sql), err)
				return
			}
			value = string(resJSON)
		}

		if err := RDB.SetEx(ctx, GetL2CacheKey(table, sql), value, config.ExpireSeconds*time.Second).Err(); err != nil {
			logger.Warnf("redis cache failed: [%s], ", GetL2CacheKey(table, sql), err)
			return
		}

		if err := RDB.SAdd(ctx, GetL2CacheKeysKey(table), GetL2CacheKey(table, sql)).Err(); err != nil {
			logger.Warnf("redis sadd failed: [%s], %v", GetL2CacheKeysKey(table), err)
			return
		}

		switch ret := RDB.TTL(ctx, GetL2CacheKeysKey(table)); true {
		case ret.Err() != nil && !errors.Is(ret.Err(), redis.Nil):
			logger.Warnf("redis ttl failed: [%s], %v", GetL2CacheKeysKey(table), err)
			return
		case errors.Is(ret.Err(), redis.Nil):
			logger.Warnf("redis expire not found: [%#v]", GetL2CacheKeysKey(table))
			if err := RDB.SAdd(ctx, GetL2CacheKeysKey(table), GetL2CacheKey(table, sql)).Err(); err != nil {
				logger.Warnf("redis sadd failed: [%s], %v", GetL2CacheKeysKey(table), err)
				return
			}
			if err := RDB.Expire(ctx, GetL2CacheKeysKey(table), config.ExpireSeconds*time.Second).Err(); err != nil {
				logger.Warnf("redis expire failed: [%s], %v", GetL2CacheKeysKey(table), err)
				return
			}
		case ret.Val() <= config.ExpireSeconds*time.Second:
			if err := RDB.Expire(ctx, GetL2CacheKeysKey(table), config.ExpireSeconds*time.Second).Err(); err != nil {
				logger.Warnf("redis expire failed: [%s], %v", GetL2CacheKeysKey(table), err)
				return
			}
		default:
			logger.Errorf("redis expire unknown condition: [%s][%#v]", GetL2CacheKeysKey(table), ret)
		}

		if err := RDB.SAdd(ctx, GetL2CacheTablesKey(), table).Err(); err != nil {
			logger.Warnf("redis sadd failed: [%s], %v", GetL2CacheKeysKey(table), err)
			return
		}

	}(resultCopy)
	return result, nil
}

func deepCopy[T any](src T) T {
	b, _ := json.Marshal(src) // 先序列化
	var dst T
	_ = json.Unmarshal(b, &dst) // 再反序列化
	return dst                  // dst 與 src 完全獨立
}
