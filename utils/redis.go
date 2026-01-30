package utils

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// 取用 份額
var promotionQuantityIncreaseLuaScript = `
local reserve = tonumber(redis.call('GET', KEYS[1]) or '0')
local limit = tonumber(ARGV[1])
if reserve < limit then
    redis.call('INCR', KEYS[1])
    return 1
else
    return 0
end
`

// 成功或失敗 返回份額
var promotionQuantityRollbackLuaScript = `
local used = tonumber(redis.call('GET', KEYS[1]) or '0')
if used > 0 then
    redis.call('DECR', KEYS[1])
end
redis.call('DEL', KEYS[2])
redis.call('SREM', KEYS[3], ARGV[1])
return 1
`

// 扣除 份額
var promotionQuantityDecrementLuaScript = `
local used = tonumber(redis.call('GET', KEYS[1]) or '0')
if used > 0 then
    redis.call('DECR', KEYS[1])
    return 1
else
    return 0
end
`

var RDB *redis.Client

func InitRedis(addr string, pwd string, db int, poolSize int, tlsVersion string) error {

	tlsConfig := (*tls.Config)(nil)
	switch tlsVersion {
	case "SSLv3":
		tlsConfig = &tls.Config{MinVersion: tls.VersionSSL30}
	case "TLS 1.0":
		tlsConfig = &tls.Config{MinVersion: tls.VersionTLS10}
	case "TLS 1.1":
		tlsConfig = &tls.Config{MinVersion: tls.VersionTLS11}
	case "TLS 1.2":
		tlsConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	case "TLS 1.3":
		tlsConfig = &tls.Config{MinVersion: tls.VersionTLS13}
	}

	RDB = redis.NewClient(&redis.Options{
		Addr:      addr,
		Password:  pwd,
		DB:        db,
		PoolSize:  poolSize,
		TLSConfig: tlsConfig,
	})

	return nil
}

func PushMessage(channel string, message string) error {
	ret := RDB.Publish(context.Background(), channel, message)

	if ret.Err() != nil {
		return ret.Err()
	}
	return nil
}

// 數量監控
func IncrementQuantityUsedWithDBFallback(ctx context.Context, usageKey string, quantityLimit, quantityUsed int64, endedAt time.Time) (bool, error) {

	// 檢查 Redis key 是否存在，否則 fallback
	ttl, err := RDB.TTL(ctx, usageKey).Result()
	if err != nil {
		return false, err
	}
	if ttl <= 0 {
		err = RDB.Set(ctx, usageKey, quantityUsed, 0).Err()
		if err != nil {
			return false, err
		}
		// 根據 EndedAt 設定 TTL
		if !endedAt.IsZero() {
			now := time.Now()
			if endedAt.After(now) {
				_ = RDB.Expire(ctx, usageKey, endedAt.Sub(now))
			}
		}
	}

	// 使用 Lua Script 進行原子性檢查 + 遞增
	script := redis.NewScript(promotionQuantityIncreaseLuaScript)
	result, err := script.Run(ctx, RDB, []string{usageKey}, quantityLimit).Result()
	if err != nil {
		return false, err
	}
	success, ok := result.(int64)
	return ok && success == 1, nil
}

// 用戶預先佔用
func MarkUserReserved(ctx context.Context, userReservKey, reserveCheckKey string, discountID, userID uint64) error {
	// 新增到候選清單中，如果用戶沒有執行 confirm 則需從清單中判斷是否需要回滾
	RDB.SAdd(ctx, reserveCheckKey, fmt.Sprintf("%d:%d", discountID, userID))
	return RDB.Set(ctx, userReservKey, 1, Config.Card.ExpireSeconds*time.Second).Err()
}

// 交易成功 用戶完成佔用
func ConfirmUserReserved(ctx context.Context, userReservKey, reserveCheckKey string, promoID, userID uint64) error {
	entry := fmt.Sprintf("%d:%d", promoID, userID)
	pipe := RDB.TxPipeline()
	pipe.Del(ctx, userReservKey)
	pipe.SRem(ctx, reserveCheckKey, entry)
	_, err := pipe.Exec(ctx)
	return err
}

// 交易失敗，返回取用量
func RollbackQuantityUsedAndUserReserved(ctx context.Context, usageKey, userReservKey, reserveCheckKey string, promoID, userID uint64) (bool, error) {

	// 檢查是否有預佔，沒有就不用回滾
	exists, err := RDB.Exists(ctx, userReservKey).Result()
	if err != nil {
		return false, err
	}
	if exists == 0 {
		return false, nil
	}

	// 執行 Lua：DECR used（如果 > 0）+ DEL reserveKey
	entry := fmt.Sprintf("%d:%d", promoID, userID)
	script := redis.NewScript(promotionQuantityRollbackLuaScript)
	_, err = script.Run(ctx, RDB, []string{usageKey, userReservKey, reserveCheckKey}, entry).Result()

	if err != nil {
		return false, err
	}

	return true, nil
}

func DecrementQuantityUsageSafely(ctx context.Context, usageKey string) (bool, error) {
	script := redis.NewScript(promotionQuantityDecrementLuaScript)

	result, err := script.Run(ctx, RDB, []string{usageKey}).Result()
	if err != nil {
		return false, err
	}

	success, ok := result.(int64)
	return ok && success == 1, nil
}
