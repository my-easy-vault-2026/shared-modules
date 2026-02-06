package common

// error code 編碼方式: CODE_{模組名稱}_{錯誤類別}
const (
	// general
	CODE_UNKNOWN                       = 0     // 未知
	CODE_OK                            = 1     // 成功
	CODE_SYSTEM_ERROR                  = 10000 // 系統錯誤
	CODE_UNKOWN_ERROR                  = 10001 // 未知錯誤
	CODE_SYSTEM_BUSY                   = 10010 // 系統忙碌
	CODE_TOO_MANY_REQUEST              = 10020 // 請求過多
	CODE_DB_ERROR                      = 10030 // 資料庫錯誤
	CODE_REDIS_ERROR                   = 10040 // 緩存錯誤
	CODE_INTERNAL_API_ERROR            = 10050 // 內部呼叫錯誤
	CODE_EXTERNAL_API_ERROR            = 10051 // 外部呼叫錯誤
	CODE_INTERNAL_API_TIMEOUT          = 10052 // 內部呼叫超時
	CODE_EXTERNAL_API_TIMEOUT          = 10053 // 外部呼叫超時
	CODE_NOT_LOGIN                     = 10060 // 未登入
	CODE_BLOCKED                       = 10061 // 用戶已封鎖
	CODE_NO_PIN_CODE                   = 10062 // 用戶未設定pin code
	CODE_ALREADY_HAS_PIN_CODE          = 10063 // 用戶已設定pin code
	CODE_INVALID_SIGNATURE             = 10070 // 簽名不符
	CODE_DATA_NOT_EXIST                = 10080 // 資料不存在
	CODE_DATA_EXISTED                  = 10081 // 資料已存在
	CODE_REQUEST_BODY_INVALID_FORMAT   = 10090 // 請求內容格式錯誤
	CODE_REQUEST_HEADER_INVALID_FORMAT = 10100 // 請求標頭格式錯誤
	CODE_TRANSACTION_FAILED            = 10110 // 帳變失敗
	CODE_NO_SUCH_TRANSACTION_TYPE      = 10111 // 無此交易類別
	CODE_ADD_ASSETS_FAILED             = 10112 // 資產增加失敗
	CODE_DEDUCT_ASSETS_FAILED          = 10113 // 資產減少失敗
	CODE_FREEZE_ASSETS_FAILED          = 10114 // 資產凍結失敗
	CODE_UNFREEZE_ASSETS_FAILED        = 10115 // 資產解凍失敗
	CODE_ADD_FREEZED_ASSETS_FAILED     = 10116 // 凍結資產增加失敗
	CODE_DEDUCT_FREEZED_ASSETS_FAILED  = 10117 // 凍結資產減少失敗
	CODE_NO_SUCH_CURRENCY              = 10130 // 沒有這個幣種
	CODE_MALFORMED_DATA                = 10140 // 資料錯誤
	CODE_INVALID_PARAMETER             = 10150 // 不可用的參數
	CODE_UNIMPLEMENTED                 = 10160 // 未實現
	CODE_NO_PERMISSION                 = 10170 // 無權限
	CODE_REDIS_SET_FAILED              = 10180 // 緩存寫入失敗
	CODE_REDIS_GET_FAILED              = 10181 // 緩存讀取失敗
	CODE_REDIS_UPDATE_FAILED           = 10182 // 緩存更新失敗
	CODE_REDIS_DELETE_FAILED           = 10183 // 緩存刪除失敗
	CODE_REDIS_EXPIRE_FAILED           = 10184 // 緩存過期失敗
	CODE_DB_INSERT_FAILED              = 10190 // 資料庫寫入失敗
	CODE_DB_SELECT_FAILED              = 10191 // 資料庫讀取失敗
	CODE_DB_UPDATE_FAILED              = 10192 // 資料庫更新失敗
	CODE_DB_DELETE_FAILED              = 10193 // 資料庫刪除失敗
	CODE_MISSING_PARAMETER             = 10200 // 缺少參數
	CODE_ENCRYPT_FAILED                = 10210 // 加密失敗
	CODE_DECRYPT_FAILED                = 10211 // 解密失敗
	CODE_SIGN_FAILED                   = 10212 // 簽名失敗
	CODE_SIGNATURE_VERIFICATION_FAILED = 10213 // 驗證簽名失敗
	CODE_ILLEGAL_OPERATION             = 10220 // 不合法的操作
	CODE_APP_VERSION_OUTDATED          = 10230 // APP版本過低，請更新至最新版本APP
	CODE_RETRY_LATER                   = 10240 // 請稍後重試
	CODE_PAGE_SIZE_EXCEEDED            = 10250 // 頁面大小超過限制
	CODE_INSUFFICIENT_MEMORY           = 10260 // 記憶體不足
	CODE_INSUFFICIENT_CPU              = 10261 // CPU不足

	CODE_NO_SUCH_USER              = 100010 // 無此用戶
	CODE_EMAIL_OR_PIN_CODE_INVALID = 100020 // 電郵或pin code不正確
	CODE_INVALED_PIN_CODE          = 100021 // pin code不正確
	CODE_ASSET_TRANSACTION_FAILED  = 100030 // 資產交易失敗
	CODE_INVALID_TARGET            = 100040 // 無效的對象
	CODE_NO_SUCH_WALLET            = 100050 // 無此錢包
	CODE_INSUFFICIENT_FUNDS        = 100060 // 資產不足
	CODE_PREVIEW_EXPIRED           = 100070 // 預覽已過期
	CODE_SELF_TRANSFER             = 100080 // 轉帳給自己
)
