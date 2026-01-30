package utils

import (
	"shared-modules/common"
	"strconv"
	"strings"

	"time"
)

// redis key prefix
var (
	prefixOTP                          = "otp"
	prefixToken                        = "token"
	prefixWsToken                      = "ws_token"
	prefixUnlockAttempt                = "unlock_attempt"
	prefixPreview                      = "preview"
	prefixLock                         = "lock"
	prefixRate                         = "rate"
	prefixUsdBaseRate                  = "usd_base_rate"
	kycSession                         = "kycSession"
	passkeyChallenge                   = "passkeyChallenge"
	otpLimit                           = "otpLimit"
	deviceToken                        = "deviceToken"
	passkeyVerify                      = "passkeyVerify"
	prefixLoginData                    = "login_data"
	prefixTokenBucket                  = "token_bucket"
	prefixTokenBucketData              = "token_bucket_data"
	prefixThreeds                      = "threeds"
	prefixThreedsMerchant              = "threeds:merchant"
	prefixL2CacheCache                 = "l2cache:cache"
	prefixL2CacheKeys                  = "l2cache:keys"
	prefixL2CacheTables                = "l2cache:tables"
	prefixPubsub                       = "pubsub"
	prefixQueue                        = "queue"
	prefixMsgList                      = "msg_list"
	prefixWebsocketNode                = "websocket:node"
	prefixFirebaseToken                = "firebase_token"
	prefixSnapshotBalance              = "snapshot_balance"
	prefixReapOtpLimit                 = "reap_otp_limit"
	prefixUserTempAddress              = "user_temp_address"
	prefixCurrencyIcon                 = "changelly_currency_icon"
	prefixChangellyOrder               = "changelly_order"
	prefixChangellyProvider            = "changelly_provider"
	prefixCheckCardStatus              = "check_card_status"
	prefixFileDownload                 = "file_download"
	prefixReapBillReport               = "reap_bill_report"
	prefixPromotionDiscountUsage       = "promotion_discount_usage"
	prefixPromotionDiscountUserReserve = "promotion_discount_reserve"
	prefixPromotionReserveCheck        = "promotion_discount_reserve_check"
	prefixMerchantUserCardApply        = "merchant_user_card_apply"
	prefixWithdrawOrderAlert           = "withdraw_order_alert"
	prefixEtherfiApplyCard             = "etherfi_apply_card"
	prefixEtherfiSystemAccountCookie   = "etherfi_system_account_cookie"
	prefixEtherfiUserCookie            = "etherfi_user_cookie"
	prefixEtherfiLoginMail             = "etherfi_login_mail"
	prefixEtherfiApplyCardRetry        = "etherfi_create_card_retry"
	prefixEtherfiSyncTx                = "etherfi_sync_tx"
	prefixEtherfiInviteRetry           = "etherfi_invite_retry"
	prefixEtherfiAcceptRetry           = "etherfi_accept_retry"
	prefixCardAmountLowAlert           = "card_amount_low_alert"
)

// otp:{purpose}:{notify type}:{otp code}
func GetOTPRedisKey(purpose common.MessagePurpose, notifyType common.NotifyType, otpCode string) string {
	return prefixOTP + ":" + purpose.String() + ":" + notifyType.String() + ":" + otpCode
}

// token:{token}
func GetTokenRedisKey(token string) string {
	return prefixToken + ":" + token
}

// unlock_attempt:{token}
func GetUnlockTokenAttemptRedisKey(token string) string {
	key := prefixUnlockAttempt + ":" + token
	return key
}

// ws_token:{role}:{token}
func GetWsTokenRedisKey(role common.Role, token string) string {
	return prefixWsToken + ":" + role.String() + ":" + token
}

// login_data:{user id}
func GetLoginDataRedisKey(userID uint64, token string) string {
	return prefixLoginData + ":" + strconv.FormatUint(userID, 10) + ":" + token
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

// threeds:{user id}
func GetThreedsRedisKey(userID uint64) string {
	return prefixThreeds + strconv.FormatUint(userID, 10)
}

// threeds:merchant:{merchant id}
func GetMerchantThreedsRedisKey(merchantID uint64) string {
	return prefixThreedsMerchant + ":" + strconv.FormatUint(merchantID, 10)
}

func GetRateRedisKey(currency string) string {
	return prefixRate + ":" + currency
}

func GetUsdBaseRateRedisKey(currency string) string {
	return prefixUsdBaseRate + ":" + currency
}

func GetKycSessionKey(session string) string {
	return kycSession + ":" + session
}

func GetPasskeyChallengeKey(userID string) string {
	return passkeyChallenge + ":" + userID
}

// token_bucket:{url}:{ip/user id}
func GetTokenBucketRedisKey(url string, ipUserID string) string {
	return prefixTokenBucket + ":" + url + ":" + ipUserID
}

// token_bucket_data:{url}:{ip/user id}
func GetTokenBucketDataRedisKey(url string, ipUserID string) string {
	return prefixTokenBucketData + ":" + url + ":" + ipUserID
}

func GetOptLimitKey(account string) string {
	return otpLimit + ":" + account
}

func GetDeviceTokenRedisKey(userID uint64, deviceID string) string {
	return deviceToken + ":" + strconv.FormatUint(userID, 10) + ":" + deviceID
}

func GetPasskeyVerifyRedisKey(userID, deviceID string) string {
	return passkeyVerify + ":" + userID + ":" + deviceID
}

func GetPasskeyVerifyPrefixKey(userID string) string {
	return passkeyVerify + ":" + userID + ":" + "*"
}

func GetReapOtpLimitRedisKey(issueID, phoneNum, otpCode string) string {
	return prefixReapOtpLimit + ":" + issueID + ":" + phoneNum + ":" + otpCode
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

// websocket:node:{user id}
func GetFirebaseTokenKey(userID uint64) string {
	key := prefixFirebaseToken + ":" + strconv.FormatUint(userID, 10)
	return key
}

func GeChangellyCurrencyIconKey(currency string) string {
	upperCurrency := strings.ToUpper(currency)
	key := prefixCurrencyIcon + ":" + upperCurrency
	return key
}

func GetThirdPartyOrderKey(walletAddress, usage string) string {
	key := prefixChangellyOrder + ":" + walletAddress + ":" + usage
	return key
}

func GetChangellyProviderKey(providerCode string) string {
	code := strings.ToLower(providerCode)
	key := prefixChangellyProvider + ":" + code
	return key
}

// filedownload:token
func GetFileDownloadKey(token string) string {
	code := strings.ToLower(token)
	return prefixFileDownload + ":" + code
}

// snapshot_balance:{financial code}:{card id}:{timestamp}
func GetSnapshotRedisKey(financialCode common.FinancialCode, cardId uint64, t time.Time) string {
	return prefixSnapshotBalance + ":" + string(financialCode) + ":" + strconv.FormatUint(cardId, 10) + ":" + t.Format("20060102T150405")
}

func GetUserTempAddressKey(address string) string {
	return prefixUserTempAddress + ":" + address
}

func GetCheckCardStatusKey(card uint64) string {
	return prefixCheckCardStatus + ":" + strconv.FormatUint(card, 10)
}

func GetReapBillReportKey(referenceID string) string {
	return prefixReapBillReport + ":" + referenceID
}

func GetDiscountUsageKey(discountID uint64) string {
	return prefixPromotionDiscountUsage + ":" + strconv.FormatUint(discountID, 10)
}

func GetDiscountUserReserveKey(discountID, userID uint64) string {
	return prefixPromotionDiscountUserReserve + ":" + strconv.FormatUint(discountID, 10) + ":" + strconv.FormatUint(userID, 10)
}

func GetDiscountReserveCheckKey() string {
	return prefixPromotionReserveCheck
}

func GetMerchantUserCardApplyKey(userID, cardID uint64) string {
	return prefixMerchantUserCardApply + ":" + strconv.FormatUint(userID, 10) + ":" + strconv.FormatUint(cardID, 10)
}

func GetWithdrawOrderAlertKey() string {
	return prefixWithdrawOrderAlert + ":overdue:set"
}

func GetEtherfiApplyCardKey(systemEmail string) string {
	return prefixEtherfiApplyCard + ":" + systemEmail
}

func GetEtherfiAdminCookieKey(adminMail string) string {
	return prefixEtherfiSystemAccountCookie + ":" + adminMail
}

func GetEtherfiUserCookieKey(userID uint64) string {
	return prefixEtherfiUserCookie + ":" + strconv.FormatUint(userID, 10)
}

func GetEtherfiUserLoginKey(etherfiMail string) string {
	return prefixEtherfiLoginMail + ":" + etherfiMail
}

func GetEtherfiCreateCardRetryKey() string {
	return prefixEtherfiApplyCardRetry
}

func GetEtherfiSyncTxKey(category string) string {
	return prefixEtherfiSyncTx + ":" + category
}

func GetEtherfiInviteRetryKey() string {
	return prefixEtherfiInviteRetry
}

func GetEtherfiAcceptRetryKey() string {
	return prefixEtherfiAcceptRetry
}

func GetCardAmountLowAlertKey(cardID uint64) string {
	return prefixCardAmountLowAlert + ":" + strconv.FormatUint(cardID, 10)
}
