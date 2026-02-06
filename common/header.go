package common

const (
	HEADER_X_TOKEN      = "X-Token" // 令牌
	HEADER_X_UID        = "X-Uid"   // 用戶id
	HEADER_X_DEVICE_ID  = "DeviceId"
	HEADER_X_EXPIRED_TS = "X-Expired" // 令牌過期時間

	HEADER_X_REAL_IP             = "X-Real-Ip"              // nginx轉發前ip
	HEADER_X_GROUP_IDS           = "X-GroupIds"             // 用戶群組
	HEADER_X_ROLE                = "X-Role"                 // 用戶角色
	HEADER_X_RATELIMIT_LIMIT     = "X-RateLimit-Limit"      // 限流上限
	HEADER_X_RATELIMIT_REMAINING = "X-RateLimit-Remaining"  // 剩餘次數
	HEADER_X_RATELIMIT_USED      = "X-RateLimit-Reset-Used" // 使用次數
	HEADER_X_RATELIMIT_RESET     = "X-RateLimit-Reset"      // 重置時間
	HEADER_X_APP_VERSION         = "X-App-version"          // app 版本
	HEADER_X_REQUEST_ID          = "X-Request-Id"           // 請求id
	HEADER_ACCEPT_LANGUAGED      = "Accept-Language"        // 語系
)
