package utils

import (
	"strconv"

	"github.com/my-easy-vault-2026/shared-modules/common"
)

// redis key prefix
var (
	prefixOTP             = "otp"
	prefixToken           = "token"
	prefixWsToken         = "ws_token"
	prefixPreview         = "preview"
	prefixLock            = "lock"
	prefixRate            = "rate"
	prefixLoginData       = "login_data"
	prefixTokenBucket     = "token_bucket"
	prefixTokenBucketData = "token_bucket_data"
	prefixL2CacheCache    = "l2cache:cache"
	prefixL2CacheKeys     = "l2cache:keys"
	prefixL2CacheTables   = "l2cache:tables"
	prefixPubsub          = "pubsub"
	prefixQueue           = "queue"
	prefixMsgList         = "msg_list"
	prefixWebsocketNode   = "websocket:node"
)

// token:{token}
func GetTokenRedisKey(token string) string {
	return prefixToken + ":" + token
}

// ws_token:{role}:{token}
func GetWsTokenRedisKey(role common.Role, token string) string {
	return prefixWsToken + ":" + role.String() + ":" + token
}

// login_data:{user id}
func GetLoginDataRedisKey(userID uint64) string {
	return prefixLoginData + ":" + strconv.FormatUint(userID, 10)
}

// preview:{purpose}:{key}
func GetPreviewRedisKey(purpose common.PreviewPurpose, key string) string {
	return prefixPreview + ":" + purpose.String() + ":" + key
}

// lock:{lock purpose}:{sub}:{token}
func GetGlobalLockKey(lockPurpose common.LockPurpose, token string, sub ...string) string {
	middle := ""
	for _, s := range sub {
		middle += s + ":"
	}
	return prefixLock + ":" + lockPurpose.String() + ":" + middle + token
}

func GetRateRedisKey(currency string) string {
	return prefixRate + ":" + currency
}

// token_bucket:{url}:{ip/user id}
func GetTokenBucketRedisKey(url string, ipUserID string) string {
	return prefixTokenBucket + ":" + url + ":" + ipUserID
}

// token_bucket_data:{url}:{ip/user id}
func GetTokenBucketDataRedisKey(url string, ipUserID string) string {
	return prefixTokenBucketData + ":" + url + ":" + ipUserID
}

// l2cache:cache:{table}:{sub}:{query}
func GetL2CacheKey(table string, query string, sub ...string) string {
	key := prefixL2CacheCache + ":" + table + ":"
	for _, s := range sub {
		key += s + ":"
	}
	return key + query
}

// l2cache:tables
func GetL2CacheTablesKey() string {
	return prefixL2CacheTables
}

// l2cache:keys:{table}
func GetL2CacheKeysKey(table string) string {
	return prefixL2CacheKeys + ":" + table
}

// msg_list:{queue}:{sub...}
func GetMsgListKey(queueName string, sub ...string) string {
	key := prefixPubsub + ":" + queueName
	for _, s := range sub {
		key += ":" + s
	}
	return key
}

// pubsub:{queue}:{sub...}
func GetPubsubKey(queueName string, sub ...string) string {
	key := prefixPubsub + ":" + queueName
	for _, s := range sub {
		key += ":" + s
	}
	return key
}

// queue:{queue}:{sub...}
func GetQueueKey(queueName string, sub ...string) string {
	key := prefixQueue + ":" + queueName
	for _, s := range sub {
		key += ":" + s
	}
	return key
}

// websocket:node:{user id}
func GetWebsocketNodeKey(userID uint64) string {
	key := prefixWebsocketNode + ":" + strconv.FormatUint(userID, 10)
	return key
}
