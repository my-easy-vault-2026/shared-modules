package common

// http header
const (
	HEADER_X_TOKEN      = "X-Token" // 令牌
	HEADER_X_UID        = "X-Uid"   // 用戶id
	HEADER_X_DEVICE_ID  = "DeviceId"
	HEADER_X_EXPIRED_TS = "X-Expired" // 令牌過期時間
	HEADER_X_EXTEND     = "X-Extend"  // 延伸資料
	HEADER_X_CONVERT    = "X-Convert" // 轉換資料

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

	// user
	CODE_USER_NO_SUCH_USER                       = 10000010 // 無此用戶
	CODE_USER_ALREADY_HAS_PIN_CODE               = 10000020 // 用戶已有pin code
	CODE_USER_PIN_CODE_CONFIRMATION_FAILED       = 10000030 // 用戶Pin code確認失敗
	CODE_USER_PIN_CODE_INVALID_FORMAT            = 10000040 // pin code格式錯誤
	CODE_USER_NO_IDENTITY                        = 10000050 // kyc信息未上传
	CODE_USER_KYC_DUPLICATED                     = 10000060 // 重复kyc
	CODE_USER_NO_FACE_DETECTION                  = 10000070 // 没有扫脸
	CODE_USER_UNABLE_DEACTIVATE_BY_REMAIN_ASSETS = 10000080 // 尚有資產餘額，無法註銷帳戶
	CODE_USER_EXCEEDS_MONTHLY_LIMIT              = 10000090 // 超过当月限额
	CODE_USER_INVALID_PHONE_NUMBER               = 10000100 // 電話號碼不正確
	CODE_USER_EXCEEDS_DAILY_LIMIT                = 10000110 // 超过当日限额
	CODE_USER_EXCEEDS_DAILY_LIMIT_COUNT          = 10000120 // 超過當日次數限制
	CODE_USER_EXCEEDS_MONTHLY_LIMIT_COUNT        = 10000130 // 超過當月次數限制
	CODE_USER_INCORRECT_OTP                      = 10000140 // OTP不正確
	CODE_USER_OTP_RETRY_LIMIT_EXCEEDED           = 10000150 // 重试超过5次，请10分钟后重试
	CODE_USER_EXCEEDS_PER_TRANSACTION            = 10000160 // 單次交易金額超過限額
	CODE_USER_LIMIT_SETTING_CANNOT_EXCEED_LIMIT  = 10000170 // 設定額度不可超過上限，幫我翻譯成英文
	CODE_USER_NO_SUCH_LIMIT                      = 10000180 // 無此限額
	CODE_USER_STATUS_BLOCKED                     = 10000190 // 使用者帳號被封鎖
	CODE_USER_EXIST                              = 10000200 // user exists
	CODE_USER_INVALID_DATE_OF_BIRTH              = 10000210 // 生日填寫不正確
	CODE_USER_NAME_TOO_LONG                      = 10000220 // 使用者名稱過長
	CODE_USER_NAME_FIELD_TOO_LONG                = 10000230 // 使用者 firstName or lastName 過長
	CODE_USER_KYC_INVALID_SESSION_ID             = 10000241 // 无效 session id
	CODE_USER_KYC_INVALID_COMPANY_ID             = 10000242 // 无效 company id
	CODE_USER_KYC_INVALID_LIVENESS_IMAGE         = 10000243 // (非pre_creation)活体检测图片无效
	CODE_USER_KYC_LIVENESS_NOT_MATCH_ID          = 10000244 // (非pre_creation)活体检测图片和证件照对比未通过
	CODE_USER_KYC_LIVENESS_COMPARE_EXCEPTION     = 10000245 // (非pre_creation)活体检测图片和证件照对比异常(未检测人脸或未知异)
	CODE_USER_KYC_LIVENESS_FAILED                = 10000246 // (非pre_creation)活体检测未通过
	CODE_USER_KYC_UPLOAD_IMAGE_AWS_FAILED        = 10000247 // 上传图片到 aws 失败
	CODE_USER_KYC_INTERNAL_ERROR                 = 10000248 // 内部错误
	CODE_USER_KYC_TOO_MANY_FAILURES              = 10000249 // 失败次数超过限制，暂时锁定该用户认证
	CODE_USER_KYC_INVALID_COMPANY_USER_ID        = 10000250 // 无效 company user id
	CODE_USER_KYC_UNRECOGNIZED_IMAGE_TYPE        = 10000251 // 不能识别的图片类型
	CODE_USER_KYC_USER_NOT_FOUND                 = 10000252 // (非pre_creation)用户不存在或状态不正确
	CODE_USER_KYC_IMAGE_SAVE_FAILED              = 10000253 // 图片保存失败
	CODE_USER_KYC_GET_AWS_IMAGE_URL_FAILED       = 10000254 // 获取aws图片url失败
	CODE_USER_KYC_OCR_FAILED                     = 10000255 // (非pre_creation)OCR 失败
	CODE_USER_KYC_TIMEOUT                        = 10000256 // (非pre_creation)认证超时失败
	CODE_USER_KYC_INVALID_TYPE                   = 10000257 // (非pre_creation)无效类型
	CODE_USER_KYC_MISSING_IMAGE                  = 10000258 // (非pre_creation)缺少图片
	CODE_USER_KYC_USER_IN_BLACKLIST              = 10000259 // (非pre_creation)用户已经在黑名单内
	CODE_USER_KYC_GET_GLOBAL_LOCK_FAILED         = 10000260 // 获取全局锁失败
	CODE_USER_KYC_SIGNATURE_VERIFICATION_FAILED  = 10000261 // 签名验证失败
	CODE_USER_KYC_FILE_INTEGRITY_VERIFICATION    = 10000262 // (非pre_creation)文件完整性验证失败
	CODE_USER_KYC_DUPLICATE_REGISTRATION         = 10000263 // (非pre_creation)用户重复注册
	CODE_USER_KYC_THIRD_PARTY_ERROR              = 10000264 // (非pre_creation)第三方错误
	CODE_USER_KYC_ALREADY_PASSED                 = 10000270 // 已通过KYC

	// auth
	CODE_AUTH_INCORRECT_EMAIL_OR_OTP    = 20000010 // email或OTP不正確
	CODE_AUTH_OTP_RETRY_LIMIT_EXCEEDED  = 20000020 // 重试超过5次，请10分钟后重试
	CODE_AUTH_TOKEN_EXPIRED             = 20000030 // token過期
	CODE_AUTH_ADMIN_USER_NOT_FOUND      = 20000040 // 管理員不存在
	CODE_BIOMETRIC_VERIFY_NOT_EXIST     = 20000050 // 發起的驗證不存在
	CODE_BIOMETRIC_VERIFY_EXIST         = 20000060 // 發起的驗證請求已存在
	CODE_VERIFY_FAIL                    = 20000070 // 驗證失敗
	CODE_BIOMETRIC_IS_NOT_VERIFY_DEVICE = 20000080 // 该账号已经设置passkey，当前操作需要到支持passkey的设备进行；如设备丢失请走passkey更换流程
	CODE_CANNOT_DELETE_CURRENT_DEVICE   = 20000090 // 不可删除当前设备
	CODE_CANNOT_DELETE_MAIN_DEVICE      = 20000100 // 不可刪除主設備
	CODE_AUTH_REAP_SIGN_EXPIRED         = 20000110 // reap簽名過期
	CODE_AUTH_REAP_SIGN_NOT_MATCH       = 20000120 // reap簽名不正確
	CODE_AUTH_AGENT_USER_NOT_FOUND      = 20000130 // 代理不存在
	CODE_AUTH_UNLOCK_FAILED             = 20000140 // 解鎖失敗
	CODE_AUTH_UNLOCK_LIMIT_EXCEEDED     = 20000141 // 解鎖失敗超過上限，強制登出

	// card
	CODE_CARD_INCORRECT_PIN_CODE                      = 30000010 // pin code不正確
	CODE_CARD_NO_SUCH_ASSET_CATEGORY                  = 30000020 // 無此類別
	CODE_CARD_NO_SUCH_ASSET                           = 30000030 // 無此資產
	CODE_CARD_ASSET_ALREADY_EXIST                     = 30000040 // 資產已存在
	CODE_CARD_ASSET_NOT_EXIST                         = 30000041 // 資產不存在
	CODE_CARD_ACTIVATION_DEPOSIT_NOT_ENOUGH           = 30000050 // 未達激活押金最小額
	CODE_CARD_APPLY_EXPIRED                           = 30000060 // 申請已過期
	CODE_CARD_INSUFFICIENT_FUNDS                      = 30000070 // 資產餘額不足
	CODE_CARD_ALREADY_DEPOSITED                       = 30000080 // 開卡已經存款
	CODE_CARD_CARD_ALREADY_CREATED                    = 30000081 // 卡片已經開卡
	CODE_CARD_CARD_ALREADY_ACTIVATED                  = 30000082 // 卡片已經激活
	CODE_CARD_CARD_ALREADY_HAS_PAN                    = 30000083 // 卡片已經取得隱私資訊
	CODE_CARD_DEPOSIT_FROM_CARD_NOT_SUPPORTED         = 30000090 // 不支援用卡片存款
	CODE_CARD_KYC_LEVEL_INSUFFICIENT                  = 30000100 // KYC等級不足
	CODE_CARD_NO_CARD_ASSET                           = 30000110 // 卡片無資產
	CODE_CARD_NO_SUCH_CARD                            = 30000120 // 沒這張卡片
	CODE_CARD_WALLET_ASSET_MISMATCH                   = 30000130 // 錢包與資產不匹配
	CODE_CARD_CARD_ALREADY_EXIST                      = 30000140 // 已有此張卡片
	CODE_CARD_NO_SUCH_APPLY_PROMOTION                 = 30000150 // 無此開卡促銷
	CODE_CARD_STATUS_EXCEPTION                        = 30000160 // 卡片状态异常
	CODE_CARD_NO_SUCH_CURRENCY                        = 30000170 // 無此幣種
	CODE_CARD_NO_PASSED_KYC                           = 30000180 // 無已通過的KYC
	CODE_CARD_NO_PROFILE_DETAIL                       = 30000190 // 用户信息不存在
	CODE_CARD_PROFILE_DETAIL_INSUFFICIENT             = 30000200 // 用户信息不完整
	CODE_CARD_ADDRESS_DETAIL_INSUFFICIENT             = 30000210 // 地址信息不完整
	CODE_CARD_REAP_CREATE_FAILED                      = 30000220 // reap開卡失敗
	CODE_CARD_INVALID_BALANCE_TYPE                    = 30000230 // 無效的餘額模式
	CODE_CARD_CATEGORY_NOT_ALLOWED_BY_MERCHANT        = 30000240 // 商戶不可使用此類型
	CODE_CARD_CATEGORY_NOT_ALLOWED                    = 30000250 // 不可使用此類型
	CODE_CARD_NOT_OPEN_FOR_APPLICATION                = 30000251 // 卡片未開放申請
	CODE_CARD_NO_SUCH_MERCHANT                        = 30000260 // 無此商戶
	CODE_CARD_ASSET_REMAINS                           = 30000270 // 請將資金提出再執行刪卡操作
	CODE_CARD_INVALID_ASSET_TYPE                      = 30000280 // 無效的資產類型
	CODE_CARD_MERCHANT_DEFAULT_APPLY_CATEGORY_NOT_SET = 30000290 // 商戶為設定預設開卡種類
	CODE_CARD_NOT_SHIPPABLE                           = 30000330 // 不可運送
	CODE_CARD_REAP_SHIP_FAILED                        = 30000340 // reap運送失敗
	CODE_CARD_REAP_GET_SHIPPING_ORDER_FAILED          = 30000350 // reap取得運送單號失敗
	CODE_CARD_REAP_ACTIVATE_CARD_FAILED               = 30000360 // reap激活卡片失敗
	CODE_CARD_SHIP_FAILED                             = 30000370 // 運送失敗
	CODE_CARD_INVALID_NATION_CODE                     = 30000380 // 無效的國碼
	CODE_CARD_PAN_OR_CODE_MISMATCH                    = 30000390 // 卡號或激活碼不符
	CODE_CARD_DELETE_FUNDING_PROCESS_NOT_COMPLETE     = 30000400 // 仍有資金流程未完成
	CODE_CARD_REAP_UPDATE_PIN_FAILED                  = 30000410 // reap更新PIN碼失敗
	CODE_CARD_STATUS_BLOCKED                          = 30000420 // 您的卡片目前处于blocked状态，请联系客服解决
	CODE_CARD_INVALID_APPLICATION_FUND                = 30000430 // 無效的資金來源
	CODE_CARD_KYC_LEVEL_REQUIRE_1                     = 30000440 // KYC等級需要1級
	CODE_CARD_KYC_LEVEL_REQUIRE_2                     = 30000450 // KYC等級需要2級
	CODE_CARD_KYC_LEVEL_REQUIRE_3                     = 30000460 // KYC等級需要3級
	CODE_CARD_USER_PROFILE_NOT_PROVIDE                = 30000470 // 未填寫用戶資料
	CODE_CARD_EMPTY_PROFILE_USER_REQUIRE_1            = 30000480 // 空白資料用戶KYC需要等級1
	CODE_CARD_EMPTY_PROFILE_USER_REQUIRE_2            = 30000490 // 空白資料用戶KYC需要等級2
	CODE_CARD_EMPTY_PROFILE_USER_REQUIRE_3            = 30000500 // 空白資料用戶KYC需要等級3
	CODE_CARD_NO_SUCH_PROVIDER                        = 30000510 // invalid card provider
	CODE_CARD_CREATE_NEED_PROMOTION                   = 30000520 // 開卡需要折扣碼
	CODE_CARD_UNBLOCK_NOT_ALLOWED                     = 30000530 // 不允許解除封鎖
	CODE_CARD_STATUS_NOT_IN_DELETE                    = 30000540 // 卡片尚未被刪除
	CODE_CARD_CARD_ALREADY_SHIPPED                    = 30000550 // 卡片已經寄出
	CODE_CARD_CARD_DELIVERY_ADDRESS_ALREADY_EXIST     = 30000560 // 卡片地址已存在
	CODE_CARD_NEGATIVE_BALANCE_BEFORE_UNBLOCK         = 30000570 // 卡片須先充值至正餘額才可解鎖
	CODE_CARD_PAN_EXIST                               = 30000580 // 卡號已存在
	CODE_CARD_PAN_IN_USE                              = 30000590 // 卡號已使用

	// account
	CODE_ACCOUNT_NO_SUCH_ASSET            = 40000010 // 無此資產
	CODE_ACCOUNT_INVALID_TARGET           = 40000020 // 無效的對象
	CODE_ACCOUNT_INVALID_TRANSACTION_TYPE = 40000030 // 無效的交易類型
	CODE_ACCOUNT_NO_SUCH_MERCHANT         = 40000040 // 無此商戶
	CODE_ACCOUNT_NO_SUCH_USER             = 40000050 // 無此用戶

	// wallet
	CODE_WALLET_NO_SUCH_WALLET                 = 50000010 // 無此錢包
	CODE_WALLET_NO_SUCH_WALLET_ADDRESS         = 50000020 // 無此錢包地址
	CODE_WALLET_NO_SUCH_MAINNET                = 50000030 // 無此主網路
	CODE_WALLET_NO_SUCH_PROTOCOL               = 50000040 // 無此協議
	CODE_WALLET_NO_SUCH_CURRENCY               = 50000050 // 無此幣種
	CODE_WALLET_NO_SUCH_CATEGORY               = 50000060 // 無此資產種類
	CODE_WALLET_INVALID_CATEGORY               = 50000070 // 不可用的資產種類
	CODE_WALLET_ALREADY_EXIST                  = 50000080 // 錢包已存在
	CODE_WALLET_NO_SUCH_FLAG                   = 50000090 // 無此合約
	CODE_WALLET_CURRENCY_MAINNET_NOT_SUPPORTED = 50000100 // 不支援的幣種與主網路
	CODE_WALLET_ADDRESS_POOL_EMPTY             = 50000110 // 地址池無資料
	CODE_WALLET_MAINNET_PROTOCOL_NOT_SUPPORTED = 50000120 // 不支援的主網路與協議
	CODE_WALLET_NOT_RECEIVABLE                 = 50000130 // 不可接收

	// quote
	CODE_QUOTE_NO_SUCH_CURRENCY = 60000010 // 無此幣種
	CODE_QUOTE_NO_SUCH_PURPOSE  = 60000020 // 無此理由
	CODE_QUOTE_NO_SUCH_RATE     = 60000030 // 無此匯率

	// system
	CODE_SYSTEM_NO_SUCH_PARAMETER = 70000010 // 無此參數

	// exchange
	CODE_EXCHANGE_NO_SUCH_CURRENCY     = 80000010 // 無此幣種
	CODE_EXCHANGE_NO_SUCH_CARD         = 80000020 // 無此卡片
	CODE_EXCHANGE_INVALID_TARGET       = 80000030 // 無效的對象
	CODE_EXCHANGE_ALREADY_PROCESSED    = 80000040 // 兌換已經執行
	CODE_EXCHANGE_INSUFFICIENT_FUND    = 80000050 // 資金不足
	CODE_EXCHANGE_NO_SUCH_CATEGORY     = 80000060 // 無此資產
	CODE_EXCHANGE_WALLET_ALREADY_EXIST = 80000070 // 錢包已存在
	CODE_EXCHANGE_FROM_CARD_FROZEN     = 80000080 // 來源卡片已凍結
	CODE_EXCHANGE_TO_CARD_FROZEN       = 80000081 // 目標卡片已凍結
	CODE_EXCHANGE_INVALID_ROLE         = 80000090 // 無效的角色
	CODE_EXCHANGE_EXCHANGE_EXPIRED     = 80000100 // 兌換已過期
	CODE_EXCHANGE_NO_SUCH_MERCHANT     = 80000110 // 無此商戶
	CODE_EXCHANGE_INVALID_CATEGORY     = 80000120 // 無效的資產
	CODE_EXCHANGE_NO_SUCH_USER         = 80000130 // 無此用戶
	CODE_EXCHANGE_AMOUNT_TOO_LOW       = 80000140 // 金額過低
	CODE_EXCHANGE_AMOUNT_TOO_HIGH      = 80000150 // 金額過高

	// reap
	CODE_REAP_NO_SUCH_CARD                       = 90000010 // 無此卡片
	CODE_REAP_INSUFFICIENT_FUNDS                 = 90000020 // 資產餘額不足
	CODE_REAP_INVALID_CURRENCY                   = 90000030 // 不可用此幣種
	CODE_REAP_NO_SUCH_TRANSACTION                = 90000040 // 無此交易
	CODE_REAP_TRANSACTION_ALREADY_PROCESSED      = 90000050 // 交易已經處理
	CODE_REAP_TRANSACTION_ALREADY_REVERSED       = 90000051 // 交易已經回饋
	CODE_REAP_TRANSACTION_ALREADY_DECLINED       = 90000052 // 交易已經取消
	CODE_REAP_AMOUNT_MISMATCH                    = 90000070 // 金額不符
	CODE_REAP_DATA_LOSS                          = 90000080 // 資料丟失
	CODE_REAP_ANOMALY                            = 90000090 // 異常交易
	CODE_REAP_ASSET_NOT_FROZEN                   = 90000100 // 資產未凍結
	CODE_REAP_THREEDS_NOTIFICATION_NOT_FOUND     = 90000110 // 3ds通知未找到
	CODE_REAP_INCORRECT_PIN_CODE                 = 90000120 // pin code不正確
	CODE_REAP_NO_SUCH_MERCHANT                   = 90000130 // 無此商戶
	CODE_REAP_SINGLE_TRANSACTION_EXCEEDS_LIMIT   = 90000140 // 單筆交易超出限額
	CODE_REAP_DAILY_TOTAL_AMOUNT_EXCEEDS_LIMIT   = 90000150 // 當日交易總額超出限額
	CODE_REAP_DAILY_TRANSACTION_COUNT_EXCEEDED   = 90000160 // 當日交易次數已達上限
	CODE_REAP_MONTHLY_TRANSACTION_COUNT_EXCEEDED = 90000170 // 當月交易次數已達上限
	CODE_REAP_MONTHLY_TOTAL_AMOUNT_EXCEEDS_LIMIT = 90000180 // 當月交易總額超出限額
	CODE_REAP_INSUFFICIENT_BALANCE               = 90000190 // 餘額不足
	CODE_REAP_INVALID_PIN_FORMAT                 = 90000200 // 無效的PIN碼格式，請重設PIN碼
	CODE_REAP_MORE_THAN_TWO_IDENTICAL_DIGITS     = 90000201 // 超過2個相同數字，請重設PIN碼
	CODE_REAP_MORE_THAN_TWO_SEQUENTAIL_DIGITS    = 90000202 // 超過2個連續數字，請重設PIN碼
	CODE_REAP_INVALID_REASON_CODE                = 90000203 // 無效的錯誤原因

	// reap auth transaction fail
	CODE_REAP_DECLINE_CODE_DEFAULT                          = 90000300 // 未預期錯誤
	CODE_REAP_DECLINE_CODE_MERCHANT_NOT_ACCEPT_CARD         = 90000301 // 商户未設置接受此類型的卡。請嘗試其他支付方式。
	CODE_REAP_DECLINE_CODE_DECLINED_BY_CLIENT               = 90000302 // 交易因“不予承兑”被拒绝。
	CODE_REAP_DECLINE_CODE_CARD_STATUS_LOST                 = 90000303 // 此卡已被掛失。
	CODE_REAP_DECLINE_CODE_CARD_DESTROYED                   = 90000304 // 此卡已失效。
	CODE_REAP_DECLINE_CODE_INSUFFICIENT_FUNDS               = 90000305 // 餘額不足，無法完成交易。請充值或更换支付方式。
	CODE_REAP_DECLINE_CODE_CARD_EXPIRED                     = 90000306 // 此卡已過期。請使用有效卡片。
	CODE_REAP_DECLINE_CODE_INCORRECT_PIN                    = 90000307 // 輸入的 PIN 碼錯誤。請重試。
	CODE_REAP_DECLINE_CODE_TRANSACTION_NOT_PERMITTED        = 90000308 // 處理交易時出錯。請聯繫 Reap 客服。
	CODE_REAP_DECLINE_CODE_POS_TRANSACTION_NOT_ALLOWED      = 90000309 // 此銷售點不允許該交易。
	CODE_REAP_DECLINE_CODE_ATM_LIMIT_EXCEEDED               = 90000310 // 您已超過提款機取款限額。請嘗試較小金額或等待限額重置。
	CODE_REAP_DECLINE_CODE_WITHDRAWAL_AMOUNT_LIMIT_EXCEEDED = 90000311 // 處理交易時出錯。請聯繫 Reap 客服。
	CODE_REAP_DECLINE_CODE_WITHDRAWAL_NOT_PERMITTED         = 90000312 // 此國家不允許使用此卡提款。
	CODE_REAP_DECLINE_CODE_WITHDRAWAL_FREQUENCY_EXCEEDED    = 90000313 // 您已超過允許的提款次數。請稍後再試。
	CODE_REAP_DECLINE_CODE_INVALID_CARD_STATUS              = 90000314 // 卡片狀態異常。請聯繫 Reap。
	CODE_REAP_DECLINE_CODE_EXCEEDING_PIN_ATTEMPTS           = 90000315 // 由於超過錯誤提款密碼次數，卡片已被鎖定。
	CODE_REAP_DECLINE_CODE_ISSUER_NOT_PARTICIPATE           = 90000316 // Reap 確實接受此商戶類別的交易。
	CODE_REAP_DECLINE_CODE_TRANSACTION_TIMEOUT              = 90000317 // 交易處理時間過長，請再試一次。
	CODE_REAP_DECLINE_CODE_PIN_VERIFICATION_FAILED          = 90000318 // 系統無法驗證您的密碼，可能因技術問題或輸入錯誤，交易無法完成。
	CODE_REAP_DECLINE_CODE_ISSUER_INOPERATIVE               = 90000319 // 發卡方或交換系統無法運行。
	CODE_REAP_DECLINE_CODE_TRANSACTION_RECORD_DISCREPANCY   = 90000320 // 支付系統的交易記錄存在不一致。
	CODE_REAP_DECLINE_CODE_UNEXPECTED_ERROR                 = 90000321 // 在支付完全處理前發生了意外錯誤。
	CODE_REAP_DECLINE_CODE_PROCESSING_ERROR                 = 90000322 // 處理交易時出現問題。請聯繫 Reap 獲取支援。
	CODE_REAP_DECLINE_CODE_INVALID_CVV                      = 90000323 // 輸入的 CVV 不正確。請檢查後重試。
	CODE_REAP_DECLINE_CODE_VERIFICATION_FAILED              = 90000324 // 提供的資料驗證失敗。請確認後再試。
	CODE_REAP_DECLINE_CODE_CARD_FROZEN                      = 90000325 // 卡片已凍結。請聯繫您的帳戶管理員。
	CODE_REAP_DECLINE_CODE_EXCEEDING_CVV_ATTEMPTS           = 90000326 // 由於多次輸入錯誤的 CVV，此卡已被鎖定。
	CODE_REAP_DECLINE_CODE_EXCEEDING_EXPIRY_DATE_ATTEMPTS   = 90000327 // 卡片因多次輸入錯誤的到期日而被鎖定。
	CODE_REAP_DECLINE_CODE_CARD_NOT_ACTIVATED               = 90000328 // 卡片尚未啟用。請透過 Reap 行動應用程式或 Reap 控制台啟用。
	CODE_REAP_DECLINE_CODE_WITHDRAWAL_NOT_ALLOWED           = 90000329 // 此卡不允許提款。
	CODE_REAP_DECLINE_CODE_DAILY_SPENDING_LIMIT_EXCEEDED    = 90000330 // 此卡的每日消費上限已超出。請聯繫管理員以申請提高限額。
	CODE_REAP_DECLINE_CODE_SPENDING_FREQUENCY_EXCEEDED      = 90000331 // 此卡的消費頻率已超出允許範圍。請聯繫管理員。
	CODE_REAP_DECLINE_CODE_CARD_BLOCKED                     = 90000332 // 卡片已被鎖定。
	CODE_REAP_DECLINE_CODE_CATEGORY_INSUFFICIENT_FUNDS      = 90000333 // 此卡餘額不足，無法完成此商戶類別的交易。請充值。
	CODE_REAP_DECLINE_CODE_BUDGET_INSUFFICIENT_FUNDS        = 90000334 // 與此卡關聯的預算資金不足。請聯繫管理員以增加預算。
	CODE_REAP_DECLINE_CODE_BUDGET_FROZEN                    = 90000335 // 與此卡關聯的預算已被冻结。請聯繫 Reap。
	CODE_REAP_DECLINE_CODE_FUND_ISSUE                       = 90000336 // 系统检测到资金异常，请联系运营团队协助处理。

	// order
	CODE_ORDER_NO_SUCH_CATEGORY       = 100000010 // 無此類別
	CODE_ORDER_USER_HAS_NO_SUCH_CARD  = 100000020 // 用戶沒有這張卡
	CODE_ORDER_USER_HAS_NO_SUCH_ORDER = 100000030 // 用戶沒有這订单

	// promotion
	CODE_PROMOTIOM_NO_SUCH_CODE         = 110000010 // 無此促銷代碼
	CODE_PROMOTIOM_CODE_NOT_APPLICABLE  = 110000020 // 促銷代碼不適用
	CODE_PROMOTIOM_NO_SUCH_CURRENCY     = 110000030 // 無此幣種
	CODE_PROMOTION_DISCOUNT_REACH_LIMIT = 110000040 // 折扣碼已達上限，請稍候再試
	CODE_PROMOTION_EXIST                = 110000050 // had promotion code

	CODE_COUNTRY_CODE_SHOULD_BE_NUMBER = 120000000

	// top up
	CODE_TOP_UP_NO_SUCH_CARD         = 130000010 // 無此卡片
	CODE_TOP_UP_INVALID_TARGET       = 130000020 // 無效的充值對象
	CODE_TOP_UP_INCORRECT_PIN_CODE   = 130000030 // pin code不正確
	CODE_TOP_UP_TOP_UP_EXPIRED       = 130000060 // 充值已過期
	CODE_TOP_UP_INSUFFICIENT_FUND    = 130000070 // 資金不足
	CODE_TOP_UP_ALREADY_PROCESSED    = 130000080 // 充值已處理
	CODE_TOP_UP_FROM_CARD_FROZEN     = 130000090 // 來源卡片已凍結
	CODE_TOP_UP_TO_CARD_FROZEN       = 130000091 // 目標卡片已凍結
	CODE_TOP_UP_FROM_CARD_BLOCKED    = 130000092 // 來源卡片已封鎖
	CODE_TOP_UP_TO_CARD_BLOCKED      = 130000093 // 目標卡片已封鎖
	CODE_TOP_UP_NO_SUCH_USER         = 130000100 // 無此用戶
	CODE_TOP_UP_AMOUNT_TOO_LOW       = 130000110 // 金額過低
	CODE_TOP_UP_AMOUNT_TOO_HIGH      = 130000120 // 金額過高
	CODE_TOP_UP_NO_SUCH_CATEGORY     = 130000130 // 無此資產類別
	CODE_TOP_UP_INVALID_CATEGORY     = 130000140 // 無效的資產類別
	CODE_TOP_UP_WHALE_AMOUNT_TOO_LOW = 130000150 // 金額過低

	// transfer
	CODE_TRANSFER_NO_SUCH_CARD                   = 140000010 // 無此卡片
	CODE_TRANSFER_INVALID_TARGET                 = 140000020 // 無效的充值對象
	CODE_TRANSFER_INCORRECT_PIN_CODE             = 140000030 // pin code不正確
	CODE_TRANSFER_TRANSFER_EXPIRED               = 140000060 // 轉帳已過期
	CODE_TRANSFER_INSUFFICIENT_FUND              = 140000070 // 資金不足
	CODE_TRANSFER_ALREADY_PROCESSED              = 140000080 // 轉帳已處理
	CODE_TRANSFER_NO_SUCH_USER                   = 140000090 // 無此用戶
	CODE_TRANSFER_NO_SUCH_CATEGORY               = 140000100 // 無此資產種類
	CODE_TRANSFER_NO_SUCH_CURRENCY               = 140000110 // 無此幣種
	CODE_TRANSFER_WALLET_ALREADY_EXIST           = 140000120 // 錢包已存在
	CODE_TRANSFER_INVALID_CATEGORY               = 140000130 // 無效的資產種類
	CODE_TRANSFER_NO_SUCH_ADDRESS                = 140000140 // 無此地址
	CODE_TRANSFER_FROM_CARD_FROZEN               = 140000150 // 來源卡片已凍結
	CODE_TRANSFER_TO_CARD_FROZEN                 = 140000151 // 目標卡片已凍結
	CODE_TRANSFER_CHANGE_TO_WITHDRAW             = 140000160 // 更改為提領
	CODE_TRANSFER_SELF_TRANSFER                  = 140000170 // 轉帳給自己
	CODE_TRANSFER_AMOUNT_TOO_LOW                 = 140000180 // 金額過低
	CODE_TRANSFER_AMOUNT_TOO_HIGH                = 140000190 // 金額過高
	CODE_TRANSFER_NO_SUCH_MAINNET                = 140000200 // 無此主網路
	CODE_TRANSFER_NO_SUCH_PROTOCOL               = 140000210 // 無此協議
	CODE_TRANSFER_MAINNET_PROTOCOL_NOT_SUPPORTED = 140000220 // 不支援的主網路與協議
	CODE_TRANSFER_TRANSFER_TO_MERCHANT_CARD      = 140000230 // 轉帳至商戶卡片
	CODE_TRANSFER_FROM_CARD_BLOCKED              = 140000240 // 來源卡片已封鎖
	CODE_TRANSFER_TO_CARD_BLOCKED                = 140000250 // 目標卡片已封鎖
	CODE_TRANSFER_TO_USER_BLOCKED                = 140000260 // 目標使用者已封鎖
	CODE_TRANSFER_FROM_USER_BLOCKED              = 140000270 // 來源使用者已封鎖

	// top down
	CODE_TOP_DOWN_NO_SUCH_CARD         = 150000010 // 無此卡片
	CODE_TOP_DOWN_INVALID_TARGET       = 150000020 // 無效的充值對象
	CODE_TOP_DOWN_INCORRECT_PIN_CODE   = 150000030 // pin code不正確
	CODE_TOP_DOWN_TOP_DOWN_EXPIRED     = 150000060 // 轉出已過期
	CODE_TOP_DOWN_INSUFFICIENT_FUND    = 150000070 // 資金不足
	CODE_TOP_DOWN_ALREADY_PROCESSED    = 150000080 // 轉出已處理
	CODE_TOP_DOWN_FROM_CARD_FROZEN     = 150000090 // 來源卡片已凍結
	CODE_TOP_DOWN_TO_CARD_FROZEN       = 150000091 // 目標卡片已凍結
	CODE_TOP_DOWN_FROM_CARD_BLOCKED    = 150000092 // 來源卡片已鎖定
	CODE_TOP_DOWN_TO_CARD_BLOCKED      = 150000093 // 目標卡片已鎖定
	CODE_TOP_DOWN_NO_SUCH_USER         = 150000100 // 無此用戶
	CODE_TOP_DOWN_AMOUNT_TOO_LOW       = 150000110 // 金額過低
	CODE_TOP_DOWN_AMOUNT_TOO_HIGH      = 150000120 // 金額過高
	CODE_TOP_DOWN_WHALE_AMOUNT_TOO_LOW = 150000130 // whale金額過低

	// withdraw
	CODE_WITHDRAW_NO_SUCH_CARD                   = 160000010 // 無此卡片
	CODE_WITHDRAW_INSUFFICIENT_FUND              = 160000020 // 資金不足
	CODE_WITHDRAW_NO_SUCH_MAINNET                = 160000030 // 無此主網路
	CODE_WITHDRAW_NO_SUCH_PROTOCOL               = 160000040 // 無此協議
	CODE_WITHDRAW_MAINNET_PROTOCOL_NOT_SUPPORTED = 160000050 // 不支援的網路與協議
	CODE_WITHDRAW_ALREADY_PROCESSED              = 160000060 // 提領已處理
	CODE_WITHDRAW_NO_SIGN_IDENTITY               = 160000070 // 沒有提幣簽名身份
	CODE_WITHDRAW_INVALID_SIGNATURE              = 160000080 // 無效的簽名
	CODE_WITHDRAW_FROM_CARD_FROZEN               = 160000090 // 來源卡片已凍結
	CODE_WITHDRAW_NO_SUCH_CATEGORY               = 160000100 // 無此資產種類
	CODE_WITHDRAW_NO_SUCH_USER                   = 160000110 // 無此用戶
	CODE_WITHDRAW_WITHDRAW_NOT_SUPPORTED         = 160000120 // 不支援提幣
	CODE_WITHDRAW_INCORRECT_PIN_CODE             = 160000130 // 用戶Pin code錯誤
	CODE_WITHDRAW_WITHDRAW_EXPIRED               = 160000140 // 提幣過期
	CODE_WITHDRAW_INVALID_CHANNEL                = 160000150 // 無效的管道
	CODE_WITHDRAW_AMOUNT_LESS_THAN_FEE           = 160000160 // 金額小於手續費
	CODE_WITHDRAW_AMOUNT_TOO_LOW                 = 160000170 // 金額過低
	CODE_WITHDRAW_AMOUNT_TOO_HIGH                = 160000180 // 金額過高
	CODE_WITHDRAW_FROM_USER_BLOCKED              = 160000190 // 來源使用者已封鎖
	CODE_WITHDRAW_FROM_CARD_BLOCKED              = 160000200 // 來源卡片已凍結

	// admin
	CODE_ADMIN_USER_ALREADY_EXIST  = 170000010 // 用戶已存在
	CODE_ADMIN_INVALID_ADMIN_LEVEL = 170000020 // 無效的管理員等級

	// coinsdo
	CODE_COINSDO_NO_SUCH_ORDER          = 180000010 // 不存在的訂單
	CODE_COINSDO_NO_SUCH_MAINNET        = 180000020 // 沒有這個主網路
	CODE_COINSDO_NO_CASE_SENSITIVE_DATA = 180000030 // 沒有區分大小寫資訊
	CODE_COINSDO_WRONG_ADDRESS_FORMAT   = 180000040 // 錯誤的地址格式
	CODE_COINSDO_ADDRESS_POOL_EMPTY     = 180000050 // 地址池為空

	// card to card
	CODE_C2C_NO_SUCH_CARD              = 190000010 // 無此卡片
	CODE_C2C_INVALID_TARGET            = 190000020 // 無效的充值對象
	CODE_C2C_INCORRECT_PIN_CODE        = 190000030 // pin code不正確
	CODE_C2C_CARD_TO_CARD_EXPIRED      = 190000060 // 轉帳已過期
	CODE_C2C_INSUFFICIENT_FUND         = 190000070 // 資金不足
	CODE_C2C_ALREADY_PROCESSED         = 190000080 // 轉帳已處理
	CODE_C2C_NO_SUCH_USER              = 190000090 // 無此用戶
	CODE_C2C_NO_SUCH_CATEGORY          = 190000100 // 無此資產種類
	CODE_C2C_NO_SUCH_CURRENCY          = 190000110 // 無此幣種
	CODE_C2C_WALLET_ALREADY_EXIST      = 190000120 // 錢包已存在
	CODE_C2C_INVALID_CATEGORY          = 190000130 // 無效的資產種類
	CODE_C2C_NO_SUCH_ADDRESS           = 190000140 // 無此地址
	CODE_C2C_FROM_CARD_FROZEN          = 190000150 // 來源卡片已凍結
	CODE_C2C_TO_CARD_FROZEN            = 190000151 // 目標卡片已凍結
	CODE_C2C_RECEIVER_HAS_NO_SUCH_CARD = 190000160 // 收方無此卡片
	CODE_C2C_PIN_CODE_INCORRECT        = 190000170 // Pin code錯誤
	CODE_C2C_AMOUNT_TOO_LOW            = 190000180 // 金額過低
	CODE_C2C_AMOUNT_TOO_HIGH           = 190000190 // 金額過高
	CODE_C2C_TRANSFER_TO_MERCHANT_CARD = 190000200 // 轉帳至商戶卡片
	CODE_C2C_FROM_USER_BLOCKED         = 190000240 // 使用者被鎖
	CODE_C2C_TO_USER_BLOCKED           = 190000210 // 收方被鎖住
	CODE_C2C_FROM_CARD_BLOCKED         = 190000220 // 卡被鎖住
	CODE_C2C_TO_CARD_BLOCKED           = 190000230 // 收方卡被鎖住

	// merchant
	CODE_MERCHANT_NO_SUCH_API_KEY                = 200000010 // 無此api密鑰
	CODE_MERCHANT_NO_SUCH_USER                   = 200000020 // 無此用戶
	CODE_MERCHANT_INVALID_USER_ROLE              = 200000030 // 無效的用戶角色
	CODE_MERCHANT_MERCHANT_ALREADY_REGISTERED    = 200000040 // 商戶已註冊
	CODE_MERCHANT_ADMIN_NOT_FROM_MERCHANT        = 200000050 // 管理員非屬於商戶
	CODE_MERCHANT_NO_SUCH_CATEGORY               = 200000060 // 無此資產種類
	CODE_MERCHANT_NO_SUCH_CURRENCY               = 200000070 // 無此幣種
	CODE_MERCHANT_INVALID_CATEGORY               = 200000080 // 無效的資產種類
	CODE_MERCHANT_NO_SUCH_CARD                   = 200000090 // 無此卡片
	CODE_MERCHANT_THREEDS_NOTIFICATION_NOT_FOUND = 200000100 // 3ds通知未找到
	CODE_MERCHANT_ASSET_NOT_EMPTY                = 200000110 // 資產不為空
	CODE_MERCHANT_INSUFFICIENT_ADMIN_AUTHORITY   = 200000120 // 管理員權限不足
	CODE_MERCHANT_NO_SUCH_MERCHANT               = 200000130 // 無此商戶
	CODE_MERCHANT_INVALID_DEFAULT_CATEGORY       = 200000140 // 無效的預設資產類型
	CODE_MERCHANT_USER_EXIST                     = 200000150 // 用戶已存在
	CODE_MERCHANT_INVALID_PHONE_NUMBER           = 200000160 // 無效的手機號碼
	CODE_MERCHANT_INVALID_CARD_CATEGORY          = 200000170 // 無效的卡片類型
	CODE_MERCHANT_INVALID_DATE_OF_BIRTH          = 200000180 // 無效的生日
	CODE_MERCHANT_EMAIL_ALREADY_TAKEN            = 200000190 // email已被使用
	CODE_MERCHANT_NO_SUCH_ORDER                  = 200000200 // 無此訂單
	CODE_MERCHANT_INSUFFICIENT_FUNDS             = 200000210 // 資產餘額不足
	CODE_MERCHANT_MERCHANT_NOT_TRUSTED           = 200000220 // 商戶未被信任

	// chippay express
	CODE_CHIPPAY_EXPRESS_NO_SUCH_CARD                         = 210000010 // 無此卡片
	CODE_CHIPPAY_EXPRESS_NO_IDENTITY                          = 210000020 // 無通過的KYC
	CODE_CHIPPAY_EXPRESS_IDENTITY_DOCUMENT_TYPE_NOT_SUPPORTED = 210000030 // 不支援的證件類型
	CODE_CHIPPAY_EXPRESS_NO_SUCH_ASSET_CATEGORY               = 210000040 // 無此類別
	CODE_CHIPPAY_EXPRESS_AMOUNT_TOO_LOW                       = 210000050 // 金額過低
	CODE_CHIPPAY_EXPRESS_AMOUNT_TOO_HIGH                      = 210000060 // 金額過高
	CODE_CHIPPAY_EXPRESS_INVALID_NAME_FORMAT                  = 210000070 // 無效的名稱格式
	CODE_CHIPPAY_EXPRESS_INCOMPLETE_ORDER                     = 210000080 // 存在未完成的訂單
	CODE_CHIPPAY_EXPRESS_KYC_LEVEL_INSUFFICIENT               = 210000090 // KYC等級不足
	CODE_CHIPPAY_EXPRESS_USER_FROZEN                          = 210000100 // 用戶被鎖定
	CODE_CHIPPAY_EXPRESS_INVALID_TARGET                       = 210000110 // 無效的對象

	// websocket
	CODE_WEBSOCKET_NO_SUCH_USER                    = 210000010 // 無此用戶
	CODE_WEBSOCKET_CREATE_WEBSOCKET_CHANNEL_FAILED = 210000020 // 創建websocket通道失敗

	// merchant adjust balance
	CODE_MERCHANT_ADJUST_BALANCE_NO_SUCH_CARD                      = 230000010 // 無此卡片
	CODE_MERCHANT_ADJUST_BALANCE_MERCHANT_ASSET_INSUFFICIENT_FUNDS = 230000020 // 商戶資產餘額不足
	CODE_MERCHANT_ADJUST_BALANCE_USER_ASSET_INSUFFICIENT_FUNDS     = 230000030 // 用戶資產餘額不足
	CODE_MERCHANT_ADJUST_BALANCE_CARD_FROZEN                       = 230000040 // 卡片凍結
	CODE_MERCHANT_ADJUST_BALANCE_BALANCE_INCONSISTENT              = 230000050 // 餘額不一致
	CODE_MERCHANT_ADJUST_BALANCE_BALANCE_DEDUCTION_NOT_SUPPORTED   = 230000060 // 不支援餘額扣除

	// open
	CODE_OPEN_NO_SUCH_LOG = 240000010 // 無此日誌

	// point
	CODE_POINT_NO_SUCH_REWARD_PRODUCT      = 250000010 // 無此兌換商品
	CODE_POINT_INVALID_REWARD_PRODUCT      = 250000020 // 無效的兌換商品
	CODE_POINT_REWARD_PRODUCT_OUT_OF_STOCK = 250000030 // 兌換商品缺貨
	CODE_POINT_INSUFFICIENT_BALANCE        = 250000040 // 餘額不足
	CODE_POINT_REWARD_PRODUCT_EXPIRED      = 250000050 // 商品過期

	// financial
	CODE_FINANCIAL_NO_SUCH_CATEGORY = 260000010 // 無此資產類型
	CODE_FINANCIAL_INVALID_CATEGORY = 260000020 // 無效的資產類型
	CODE_FINANCIAL_INVALID_CURRENCY = 260000030 // 無效的幣種
	CODE_FINANCIAL_NO_SUCH_CARD     = 260000040 // 卡片不存在
	CODE_FINANCIAL_INVALID_TARGET   = 260000050 // 無效的對象
	CODE_FINANCIAL_PRODUCT_DISABLED = 260000060 // 理財產品未開啟

	// third party deposit
	CODE_DEPOSITED_CURRENTLY_UNAVAILABLE_EXCHANGE_PAIR = 270000010 // 目前该货币对无法交易，请稍候再试

	// whale
	CODE_WHALE_CREATE_CARD_FAILED    = 280000010 // 開卡失敗
	CODE_WHALE_NO_SUCH_TRANSACTION   = 280000040 // 無此交易
	CODE_WHALE_CARD_INACTIVE         = 280000050 // 卡片狀態未啟用
	CODE_WHALE_CARD_LIMIT_US         = 280000060 // 卡片帳單限制區域 US
	CODE_WHALE_CARD_LIMIT_US_ADDRESS = 280000070 // 卡片帳單限制區域 US 地址
	CODE_WHALE_CARD_UNFROZEN_FAILED  = 280000080 // 卡片解凍失敗

	// notification
	CODE_NOTIFICATION_ENGLISH_NOT_SET      = 290000010 // 公告格式錯誤
	CODE_NOTIFYICATION_DATE_FORMAT         = 290000020 // 日期格式錯誤
	CODE_NOTIFICATION_EXIST                = 290000030 // notification exists
	CODE_NOTIFICATION_CHINESE_NOT_SET      = 290000040 // 公告格式錯誤
	CODE_NOTIFICATION_CONTEXT_LENGTH_LIMIT = 290000050 // 內文長度限制

	// paycrypto
	CODE_PAYCRYPTO_REAPPLY_KYC                     = 300000010 // 請重新申請KYC，補足缺漏資料
	CODE_PAYCRYPTO_KYC_FAILED                      = 300000020 // 開卡KYC失敗
	CODE_PAYCRYPTO_KYC_PENDING                     = 300000030 // 請等待KYC審核
	CODE_PAYCRYPTO_EXCEED_APPLY_LIMIT              = 300000040 // 超過申請上限
	CODE_PAYCRYPTO_KYC_NOT_FINISHED                = 300000050 // KYC尚未完成
	CODE_PAYCRYPTO_DOCUMENT_TYPE_NOT_SUPPORTED     = 300000060 // 不支援的證件型式
	CODE_PAYCRYPTO_ACTIVATE_FAILED                 = 300000070 // 啟用卡片失敗
	CODE_PAYCRYPTO_PHYSICAL_INVENTORY_NO_AVAILABLE = 300000080 // 實體卡庫存不足
	CODE_PAYCRYPTO_PIN_CODE_LENGTH_INVALID         = 300000090 // 密码长度必须为4位
	CODE_PAYCRYPTO_PIN_CODE_ALL_SAME_DIGITS        = 300000100 // 密码不能为全部相同的数字
	CODE_PAYCRYPTO_PIN_CODE_CONTAINS_PHONE_NUMBER  = 300000110 // 密码不能包含手机号
	CODE_PAYCRYPTO_PIN_CODE_SAME_AS_CARD_LAST_FOUR = 300000120 // 密码不能与卡号后四位相同
	CODE_PAYCRYPTO_NOT_ACTIVATED                   = 300000130 // 卡片尚未激活

	// binance pay
	CODE_BINANCE_PAY_INCORRECT_ORDER_STATUS = 310000010 // order狀態異常
	CODE_BINANCE_PAY_CREATE_ORDER_FAILED    = 310000020 // can not create order
	CODE_BINANCE_PAY_DEPOSIT_LIMIT          = 310000030 // deposit amount limit
	CODE_BINANCE_PAY_ORDER_NUMBER_LIMIT     = 310000040 // order size limit

	// file access
	CODE_FILE_ACCESS_TYPE_INCORRECT = 320000010 // file access type incorrect

	// delivery
	CODE_DELIVERY_INVALID_ORDER_STATUS = 330000010 // invalid order status
	CODE_DELIVERY_ORDER_NOT_FOUND      = 330000020 // delivery order not found

	// etherfi
	CODE_ETHERFI_SPENDING_LIMIT_NOT_EQUAL = 340000010 // spending limit not equal
	CODE_ETHERFI_NO_SUCH_MERCHANT         = 340000020 // no such merchant

	// finance report
	CODE_FINANCE_REPORT_TABLE_INCORRECT = 350000010 // finance report table incorrect
	CODE_FINANCE_REPORT_RANGE_TOO_LARGE = 350000020 // finance report range too large

	// physical card order
	CODE_PHYSICAL_CARD_ORDER_CARD_NOT_FOUND = 360000010 // card not found
	CODE_PHYSICAL_CARD_ORDER_MAX_SIZE       = 360000020 // max item size

)

// system account
const (
	SYSTEM_ACCOUNT_PLATFORM uint64 = 1
	SYSTEM_ACCOUNT_FEE      uint64 = 2
	SYSTEM_ACCOUNT_AUDIT    uint64 = 3
	SYSTEM_ACCOUNT_INTEREST uint64 = 4
)

type MerchantAdjustBalanceOrderStatus int // 充值狀態
const (
	MERCHANT_ADJUST_BALANCE_ORDER_STATUS_PENDING MerchantAdjustBalanceOrderStatus = 1
	MERCHANT_ADJUST_BALANCE_ORDER_STATUS_SUCCESS MerchantAdjustBalanceOrderStatus = 2
	MERCHANT_ADJUST_BALANCE_ORDER_STATUS_FAILED  MerchantAdjustBalanceOrderStatus = 3
)

type AdminLevel int // 管理等級
const (
	ADMIN_LEVEL_ANALYTICS        AdminLevel = 10
	ADMIN_LEVEL_CUSTOMER_SERVICE AdminLevel = 20
	ADMIN_LEVEL_FINANCE          AdminLevel = 30
	ADMIN_LEVEL_OPERATION        AdminLevel = 40
	ADMIN_LEVEL_ADMIN            AdminLevel = 90
)

type AdminUserStatus int // 管理員狀態
const (
	ADMIN_USER_STATUS_ENABLED  AdminUserStatus = 1
	ADMIN_USER_STATUS_DISABLED AdminUserStatus = 2
)

type ApplicationStatus int // 開卡狀態
const (
	APPLICATION_STATUS_PENDING                 ApplicationStatus = 1
	APPLICATION_STATUS_CREATED                 ApplicationStatus = 2
	APPLICATION_STATUS_DEPOSITED               ApplicationStatus = 3
	APPLICATION_STATUS_FAILED                  ApplicationStatus = 4
	APPLICATION_STATUS_WHALE_DEPOSITED         ApplicationStatus = 5
	APPLICATION_STATUS_WHALE_CREATED           ApplicationStatus = 6
	APPLICATION_STATUS_WHALE_DEPOSIT_FAILED    ApplicationStatus = 7
	APPLICATION_STATUS_PAYCRYPTO_DEPOSITED     ApplicationStatus = 8
	APPLICATION_STATUS_PAYCRYPTO_CREATED       ApplicationStatus = 9
	APPLICATION_STATUS_PAYCRYPTO_NOT_ACTIVATED ApplicationStatus = 10
	APPLICATION_STATUS_PAYCRYPTO_ACTIVATED     ApplicationStatus = 11
)

type ApplyPromotionRuleType int // 開卡促銷規則種類
const (
	APPLY_PROMOTION_RULE_TYPE_ALL     ApplyPromotionRuleType = 1 // 所有卡片
	APPLY_PROMOTION_RULE_TYPE_PARTIAL ApplyPromotionRuleType = 2 // 部分卡片
)

type BiometricRegisterType int // 生物辨識註冊類型
const (
	AFTER_REGISTER  BiometricRegisterType = 1 // 注册后开启
	IN_REGISTER     BiometricRegisterType = 2 // 注册阶段开启
	CHANGE_REGISTER BiometricRegisterType = 3 // 更改验证设备
)

type APIAuthorityStatus int // api權限狀態
const (
	API_AUTHORITY_STATUS_ENABLED  ApplicationStatus = 1
	API_AUTHORITY_STATUS_DISABLED ApplicationStatus = 2
)

type ApplyPromotionStatus int // 開卡促銷狀態
const (
	APPLY_PROMOTION_STATUS_ENABLED  ApplyPromotionStatus = 1
	APPLY_PROMOTION_STATUS_DISABLED ApplyPromotionStatus = 2
)

type AssetType int // 資產類型
const (
	ASSET_TYPE_CRYPTO       AssetType = 1
	ASSET_TYPE_FIAT         AssetType = 2
	ASSET_TYPE_CARD_PRODUCT AssetType = 3
	ASSET_TYPE_POINT        AssetType = 4
	ASSET_TYPE_AUTO_YIELD   AssetType = 5
)

type CardKind int // 卡片類型
const (
	CARD_KIND_REAP_VIRTUAL         CardKind = 1
	CARD_KIND_REAP_PHYSICAL        CardKind = 2
	CARD_KIND_WHALE                CardKind = 3
	CARD_KIND_WHALE_AI             CardKind = 4
	CARD_KIND_PAYCRYPTO_VIRTUAL    CardKind = 5
	CARD_KIND_REAP_MONETA_VIRTUAL  CardKind = 6
	CARD_KIND_REAP_MONETA_PHYSICAL CardKind = 7
	CARD_KIND_ETHERFI_VIRTUAL      CardKind = 8
	CARD_KIND_ETHERFI_PHYSICAL     CardKind = 9
	CARD_KIND_PAYCRYPTO_PHYSICAL   CardKind = 10
)

type CardType int // 卡片類型
const (
	CARD_TYPE_REAP      CardType = 1
	CARD_TYPE_WHALE     CardType = 2
	CARD_TYPE_WHALE_AI  CardType = 3
	CARD_TYPE_PAYCRYPTO CardType = 4
	CARD_TYPE_ETHERFI   CardType = 5
)

type ATMToggle int // atm開關
const (
	ATM_TOGGLE_ENABLED  ATMToggle = 1
	ATM_TOGGLE_DISABLED ATMToggle = 2
)

type Auto3DSStatus int // 自動3DS開關
const (
	AUTO_3DS_STATUS_ENABLED  Auto3DSStatus = 1
	AUTO_3DS_STATUS_DISABLED Auto3DSStatus = 2
)

type AutoTopUpStatus int // 自動充值開關
const (
	AUTO_TOP_UP_STATUS_ENABLED  AutoTopUpStatus = 1
	AUTO_TOP_UP_STATUS_DISABLED AutoTopUpStatus = 2
)

type BalanceType int // 餘額類型
const (
	BALANCE_TYPE_CENTRALIZED            BalanceType = 1
	BALANCE_TYPE_DECENTRALIZED          BalanceType = 2
	BALANCE_TYPE_MERCHANT_CENTRALIZED   BalanceType = 3
	BALANCE_TYPE_MERCHANT_DECENTRALIZED BalanceType = 4
	BALANCE_TYPE_MERCHANT_STANDARD      BalanceType = 5
)

type BillType int // 帳變類型
const (
	BILL_TYPE_APPLY                   BillType = 100
	BILL_TYPE_TRANSFER                BillType = 200
	BILL_TYPE_EXCHANGE                BillType = 300
	BILL_TYPE_WITHDRAW                BillType = 400
	BILL_TYPE_PAY_REAP                BillType = 500
	BILL_TYPE_DEPOSIT                 BillType = 600
	BILL_TYPE_TOP_UP                  BillType = 700
	BILL_TYPE_TOP_DOWN                BillType = 800
	BILL_TYPE_REAP_REFUND             BillType = 900
	BILL_TYPE_REAP_VERIFY             BillType = 1000
	BILL_TYPE_CARD_TO_CARD            BillType = 1100
	BILL_TYPE_MANUAL                  BillType = 1200
	BILL_TYPE_MERCHANT_AUTO_EXCHANGE  BillType = 1400
	BILL_TYPE_CHIPPAY_EXPRESS         BillType = 1500
	BILL_TYPE_MERCHANT_ADJUST_BALANCE BillType = 1600
	BILL_TYPE_POINT_ACCURAL           BillType = 1700
	BILL_TYPE_POINT_REDEEM            BillType = 1800
	BILL_TYPE_INTEREST                BillType = 1900
	BILL_TYPE_FINANCIAL_TRANSFER      BillType = 2000
	BILL_TYPE_WHALE                   BillType = 2100
	BILL_TYPE_DISTRIBUTE              BillType = 2300
	BILL_TYPE_PAYCRYPTO               BillType = 2400
	BILL_TYPE_BINANCE_PAY             BillType = 2500
	BILL_TYPE_SYSTEM_CHARGE           BillType = 2600
	BILL_TYPE_ETHERFI                 BillType = 3300
)
const (
	BILL_TYPE_APPLY_FEE_DEDUCT                      BillType = 101 // 101 用戶扣兌換手續費
	BILL_TYPE_APPLY_FEE_ADD                         BillType = 102 // 102 手續費帳戶入帳兌換手續費
	BILL_TYPE_APPLY_CARD_FEE_DEDUCT                 BillType = 103 // 103 用戶扣開卡卡片手續費
	BILL_TYPE_APPLY_CARD_FEE_ADD                    BillType = 104 // 104 手續費帳戶入帳開卡卡片手續費
	BILL_TYPE_APPLY_AUTO_REACTIVATE_FEE_DEDUCT      BillType = 105 // 105 自動重新激活時用戶扣兌換手續費
	BILL_TYPE_APPLY_AUTO_REACTIVATE_FEE_ADD         BillType = 106 // 106 自動重新激活時手續費帳戶入帳兌換手續費
	BILL_TYPE_APPLY_AUTO_REACTIVATE_CARD_FEE_DEDUCT BillType = 107 // 107 自動重新激活時用戶扣開卡卡片手續費
	BILL_TYPE_APPLY_AUTO_REACTIVATE_CARD_FEE_ADD    BillType = 108 // 108 自動重新激活時手續費帳戶入帳開卡卡片手續費
	BILL_TYPE_APPLY_MERCHANT_FEE_DEDUCT             BillType = 109 // 109 商戶用戶扣兌換手續費
	BILL_TYPE_APPLY_MERCHANT_FEE_ADD                BillType = 110 // 110 手續費帳戶入帳商戶兌換手續費
	BILL_TYPE_APPLY_MERCHANT_CARD_FEE_DEDUCT        BillType = 111 // 111 商戶用戶扣開卡卡片手續費
	BILL_TYPE_APPLY_MERCHANT_CARD_FEE_ADD           BillType = 112 // 112 手續費帳戶入帳商戶開卡卡片手續費
	BILL_TYPE_APPLY_WHALE_FEE_DEDUCT                BillType = 113 // 113 whale開卡用戶扣兌換手續費
	BILL_TYPE_APPLY_WHALE_FEE_ADD                   BillType = 114 // 114 whale手續費帳戶入帳兌換手續費
	BILL_TYPE_APPLY_WHALE_CARD_FEE_DEDUCT           BillType = 115 // 115 whale用戶扣開卡卡片手續費
	BILL_TYPE_APPLY_WHALE_CARD_FEE_ADD              BillType = 116 // 116 whale手續費帳戶入帳開卡卡片手續費
	BILL_TYPE_APPLY_WHALE_RECEIVE_ADD               BillType = 117 // 117 whale用戶開卡加錢
	BILL_TYPE_APPLY_WHALE_RECEIVE_DEDUCT            BillType = 118 // 118 whale平台開卡扣錢
	BILL_TYPE_APPLY_PAYCRYPTO_FEE_DEDUCT            BillType = 121 // 121 paycrypto開卡用戶扣兌換手續費
	BILL_TYPE_APPLY_PAYCRYPTO_FEE_ADD               BillType = 122 // 122 paycrypto手續費帳戶入帳兌換手續費
	BILL_TYPE_APPLY_PAYCRYPTO_CARD_FEE_DEDUCT       BillType = 123 // 123 paycrypto用戶扣開卡卡片手續費
	BILL_TYPE_APPLY_PAYCRYPTO_CARD_FEE_ADD          BillType = 124 // 124 paycrypto手續費帳戶入帳開卡卡片手續費
	BILL_TYPE_APPLY_PAYCRYPTO_RECEIVE_ADD           BillType = 125 // 125 paycrypto用戶開卡加錢
	BILL_TYPE_APPLY_PAYCRYPTO_RECEIVE_DEDUCT        BillType = 126 // 126 paycrypto平台開卡扣錢
	BILL_TYPE_APPLY_ETHERFI_FEE_DEDUCT              BillType = 127 // 127 etherfi開卡用戶扣兌換手續費
	BILL_TYPE_APPLY_ETHERFI_FEE_ADD                 BillType = 128 // 128 etherfi手續費帳戶入帳兌換手續費
	BILL_TYPE_APPLY_ETHERFI_CARD_FEE_DEDUCT         BillType = 129 // 129 etherfi用戶扣開卡卡片手續費
	BILL_TYPE_APPLY_ETHERFI_CARD_FEE_ADD            BillType = 130 // 130 etherfi手續費帳戶入帳開卡卡片手續費
	BILL_TYPE_APPLY_ETHERFI_RECEIVE_ADD             BillType = 131 // 131 etherfi用戶開卡加錢
	BILL_TYPE_APPLY_ETHERFI_RECEIVE_DEDUCT          BillType = 132 // 132 etherfi平台開卡扣錢

	BILL_TYPE_TRANSFER_SEND_DEDUCT                       BillType = 201 // 201 發方用戶扣款
	BILL_TYPE_TRANSFER_SEND_FEE_DEDUCT                   BillType = 202 // 202 發方用戶扣手續費
	BILL_TYPE_TRANSFER_SEND_ADD                          BillType = 203 // 203 平台帳戶入帳
	BILL_TYPE_TRANSFER_SEND_FEE_ADD                      BillType = 204 // 204 手續費帳戶入帳手續費
	BILL_TYPE_TRANSFER_RECEIVE_ADD                       BillType = 205 // 205 收方用戶入款
	BILL_TYPE_TRANSFER_RECEIVE_DEDUCT                    BillType = 206 // 206 平台扣款
	BILL_TYPE_TRANSFER_SEND_EXCHANGE_FEE_DEDUCT          BillType = 207 // 207 發方用戶扣匯差費
	BILL_TYPE_TRANSFER_SEND_EXCHANGE_FEE_ADD             BillType = 208 // 208 手續費帳戶入帳匯差費
	BILL_TYPE_TRANSFER_MERCHANT_SEND_DEDUCT              BillType = 209 // 209 商戶轉帳子用戶 發方用戶扣款
	BILL_TYPE_TRANSFER_MERCHANT_SEND_FEE_DEDUCT          BillType = 210 // 210 商戶轉帳子用戶 發方用戶扣手續費
	BILL_TYPE_TRANSFER_MERCHANT_SEND_ADD                 BillType = 211 // 211 商戶轉帳子用戶 平台帳戶入帳
	BILL_TYPE_TRANSFER_MERCHANT_SEND_FEE_ADD             BillType = 212 // 212 商戶轉帳子用戶 手續費帳戶入帳手續費
	BILL_TYPE_TRANSFER_MERCHANT_RECEIVE_ADD              BillType = 213 // 213 商戶轉帳子用戶 收方用戶入款
	BILL_TYPE_TRANSFER_MERCHANT_RECEIVE_DEDUCT           BillType = 214 // 214 商戶轉帳子用戶 平台扣款
	BILL_TYPE_TRANSFER_MERCHANT_SEND_EXCHANGE_FEE_DEDUCT BillType = 215 // 215 商戶轉帳子用戶 發方用戶扣匯差費
	BILL_TYPE_TRANSFER_MERCHANT_SEND_EXCHANGE_FEE_ADD    BillType = 216 // 216 商戶轉帳子用戶 手續費帳戶入帳匯差費

	BILL_TYPE_EXCHANGE_SEND_DEDUCT              BillType = 301 // 301 用戶扣款
	BILL_TYPE_EXCHANGE_SEND_FEE_DEDUCT          BillType = 302 // 302 用戶扣手續費
	BILL_TYPE_EXCHANGE_SEND_ADD                 BillType = 303 // 303 平台帳戶入帳
	BILL_TYPE_EXCHANGE_SEND_FEE_ADD             BillType = 304 // 304 手續費帳戶入帳手續費
	BILL_TYPE_EXCHANGE_RECEIVE_ADD              BillType = 305 // 305 用戶入款
	BILL_TYPE_EXCHANGE_RECEIVE_DEDUCT           BillType = 306 // 306 平台扣款
	BILL_TYPE_EXCHANGE_MERCHANT_SEND_DEDUCT     BillType = 307 // 307 商戶兌換 商戶扣款
	BILL_TYPE_EXCHANGE_MERCHANT_SEND_FEE_DEDUCT BillType = 308 // 308 商戶兌換 商戶扣手續費
	BILL_TYPE_EXCHANGE_MERCHANT_SEND_ADD        BillType = 309 // 309 商戶兌換 平台帳戶入帳
	BILL_TYPE_EXCHANGE_MERCHANT_SEND_FEE_ADD    BillType = 310 // 310 商戶兌換 手續費帳戶入帳手續費
	BILL_TYPE_EXCHANGE_MERCHANT_RECEIVE_ADD     BillType = 311 // 311 商戶兌換 商戶入款
	BILL_TYPE_EXCHANGE_MERCHANT_RECEIVE_DEDUCT  BillType = 312 // 312 商戶兌換 平台扣款

	BILL_TYPE_WITHDRAW_SEND_DEDUCT                     BillType = 401 // 401 發方用戶扣款
	BILL_TYPE_WITHDRAW_SEND_FEE_DEDUCT                 BillType = 402 // 402 發方用戶扣款
	BILL_TYPE_WITHDRAW_SEND_ADD                        BillType = 403 // 403 審核帳戶入帳
	BILL_TYPE_WITHDRAW_SEND_FEE_ADD                    BillType = 404 // 404 手續費帳戶入帳手續費
	BILL_TYPE_WITHDRAW_REJECT_DEDUCT                   BillType = 405 // 405 因駁回將審核帳戶扣回
	BILL_TYPE_WITHDRAW_REJECT_FEE_DEDUCT               BillType = 406 // 406 因駁回將手續費帳戶扣回
	BILL_TYPE_WITHDRAW_REJECT_ADD                      BillType = 407 // 407 因駁回將發方用戶入帳
	BILL_TYPE_WITHDRAW_REJECT_FEE_ADD                  BillType = 408 // 408 因駁回將手續費帳戶扣回手續費
	BILL_TYPE_WITHDRAW_CANCEL_DEDUCT                   BillType = 409 // 409 因取消將審核帳戶扣回
	BILL_TYPE_WITHDRAW_CANCEL_ADD                      BillType = 410 // 410 因取消將發方用戶入帳
	BILL_TYPE_WITHDRAW_CANCEL_FEE_DEDUCT               BillType = 411 // 411 因取消將手續費帳戶扣回手續費
	BILL_TYPE_WITHDRAW_CANCEL_FEE_ADD                  BillType = 412 // 412 因取消將發方用戶入帳手續費
	BILL_TYPE_WITHDRAW_SUCCESS_DEDUCT                  BillType = 413 // 413 因成功將審核帳戶扣回
	BILL_TYPE_WITHDRAW_SUCCESS_ADD                     BillType = 414 // 414 因成功將平台帳戶入帳
	BILL_TYPE_WITHDRAW_TRANSFER_SEND_DEDUCT            BillType = 415 // 415 轉帳轉提幣 發方用戶扣款
	BILL_TYPE_WITHDRAW_TRANSFER_SEND_FEE_DEDUCT        BillType = 416 // 416 轉帳轉提幣 發方用戶扣款
	BILL_TYPE_WITHDRAW_TRANSFER_SEND_ADD               BillType = 417 // 417 轉帳轉提幣 審核帳戶入帳
	BILL_TYPE_WITHDRAW_TRANSFER_SEND_FEE_ADD           BillType = 418 // 418 轉帳轉提幣 手續費帳戶入帳手續費
	BILL_TYPE_WITHDRAW_TRANSFER_FAILED_SEND_DEDUCT     BillType = 419 // 419 轉帳轉提幣 失敗時審核帳戶扣款
	BILL_TYPE_WITHDRAW_TRANSFER_FAILED_SEND_FEE_DEDUCT BillType = 420 // 420 轉帳轉提幣 失敗時手續費帳戶扣手續費
	BILL_TYPE_WITHDRAW_TRANSFER_FAILED_SEND_ADD        BillType = 421 // 421 轉帳轉提幣 失敗時發方用戶入帳
	BILL_TYPE_WITHDRAW_TRANSFER_FAILED_SEND_FEE_ADD    BillType = 422 // 422 轉帳轉提幣 失敗時發方用戶入帳手續費

	BILL_TYPE_PAY_REAP_AUTHORIZATION_FREEZE                             BillType = 501 // 501 用戶凍結
	BILL_TYPE_PAY_REAP_TRANSACTION_DEDUCT_FREEZED                       BillType = 502 // 502 用戶凍結扣除
	BILL_TYPE_PAY_REAP_TRANSACTION_ADD                                  BillType = 503 // 503 平台入帳
	BILL_TYPE_PAY_REAP_TRANSACTION_REVERSAL_UNFREEZE                    BillType = 504 // 504 用戶回沖解凍
	BILL_TYPE_PAY_REAP_TRANSACTION_REVERSAL_PARTIAL_UNFREEZE            BillType = 541 // 541 用戶部分回沖解凍
	BILL_TYPE_PAY_REAP_DECLINE_UNFREEZE                                 BillType = 505 // 505 用戶取消解凍
	BILL_TYPE_PAY_REAP_TRANSACTION_DIRECT_DEDUCT                        BillType = 506 // 506 用戶直接扣除
	BILL_TYPE_PAY_REAP_TRANSACTION_DIRECT_ADD                           BillType = 507 // 507 平台直接入帳
	BILL_TYPE_PAY_REAP_TRANSACTION_OFFLINE_DEDUCT                       BillType = 508 // 508 用戶離線直接扣除
	BILL_TYPE_PAY_REAP_TRANSACTION_OFFLINE_ADD                          BillType = 509 // 509 平台離線直接入帳
	BILL_TYPE_PAY_REAP_TRANSACTION_REFUND_DEDUCT                        BillType = 510 // 510 平台退款扣除
	BILL_TYPE_PAY_REAP_TRANSACTION_REFUND_ADD                           BillType = 511 // 511 商戶用戶退款入帳
	BILL_TYPE_PAY_REAP_TRANSACTION_DRIECT_AFTER_DECLINE_ADD             BillType = 512 // 512 拒絕後清算平台入帳
	BILL_TYPE_PAY_REAP_TRANSACTION_DRIECT_AFTER_DECLINE_DEDUCT          BillType = 513 // 513 拒絕後清算用戶扣款
	BILL_TYPE_PAY_REAP_AUTHORIZATION_FEE_DEDUCT                         BillType = 514 // 514 用戶授權費扣除
	BILL_TYPE_PAY_REAP_AUTHORIZATION_FEE_ADD                            BillType = 515 // 515 平台授權費入帳
	BILL_TYPE_PAY_REAP_MERCHANT_AUTHORIZATION_FREEZE                    BillType = 521 // 521 商戶用戶凍結
	BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_DEDUCT_FREEZED              BillType = 522 // 522 商戶用戶凍結扣除
	BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_ADD                         BillType = 523 // 523 商戶交易時平台入帳
	BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_REVERSAL_UNFREEZE           BillType = 524 // 524 商戶用戶回沖解凍
	BILL_TYPE_PAY_REAP_MERCHANT_DECLINE_UNFREEZE                        BillType = 525 // 525 商戶用戶取消解凍
	BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_DIRECT_DEDUCT               BillType = 526 // 526 商戶用戶直接扣除
	BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_DIRECT_ADD                  BillType = 527 // 527 商戶交易時平台直接入帳
	BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_OFFLINE_DEDUCT              BillType = 528 // 528 商戶用戶離線直接扣除
	BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_OFFLINE_ADD                 BillType = 529 // 529 商戶交易時平台離線直接入帳
	BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_REFUND_DEDUCT               BillType = 520 // 520 商戶交易時平台退款扣除
	BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_REFUND_ADD                  BillType = 531 // 531 商戶用戶退款入帳
	BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_DRIECT_AFTER_DECLINE_ADD    BillType = 532 // 532 商戶交易時拒絕後清算平台入帳
	BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_DRIECT_AFTER_DECLINE_DEDUCT BillType = 533 // 533 拒絕後清算商戶用戶扣款
	BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_REVERSAL_PARTIAL_UNFREEZE   BillType = 542 // 542 商戶用戶部分回沖解凍

	BILL_TYPE_DEPOSIT_RECEIVE_ADD    BillType = 601 // 601 用戶入帳
	BILL_TYPE_DEPOSIT_RECEIVE_DEDUCT BillType = 602 // 602 平台扣款
	BILL_TYPE_DEPOSIT_FEE_DEDUCT     BillType = 603 // 603 平台扣手續費
	BILL_TYPE_DEPOSIT_FEE_ADD        BillType = 604 // 604 收續費帳戶入帳手續費

	BILL_TYPE_TOP_UP_SEND_DEDUCT                           BillType = 701 // 701 用戶扣款
	BILL_TYPE_TOP_UP_SEND_FEE_DEDUCT                       BillType = 702 // 702 用戶扣手續費
	BILL_TYPE_TOP_UP_SEND_ADD                              BillType = 703 // 703 平台帳戶入帳
	BILL_TYPE_TOP_UP_SEND_FEE_ADD                          BillType = 704 // 704 手續費帳戶入帳手續費
	BILL_TYPE_TOP_UP_RECEIVE_ADD                           BillType = 705 // 705 用戶入款
	BILL_TYPE_TOP_UP_RECEIVE_DEDUCT                        BillType = 706 // 706 平台扣款
	BILL_TYPE_TOP_UP_SEND_EXCHANGE_FEE_DEDUCT              BillType = 707 // 707 用戶扣 匯差費
	BILL_TYPE_TOP_UP_SEND_EXCHANGE_FEE_ADD                 BillType = 708 // 708 手續費帳戶入帳匯差費
	BILL_TYPE_TOP_UP_AUTO_SEND_DEDUCT                      BillType = 711 // 711 自動充值用戶扣款
	BILL_TYPE_TOP_UP_AUTO_SEND_FEE_DEDUCT                  BillType = 712 // 712 自動充值用戶扣手續費
	BILL_TYPE_TOP_UP_AUTO_SEND_ADD                         BillType = 713 // 713 自動充值平台帳戶入帳
	BILL_TYPE_TOP_UP_AUTO_SEND_FEE_ADD                     BillType = 714 // 714 自動充值手續費帳戶入帳手續費
	BILL_TYPE_TOP_UP_AUTO_RECEIVE_ADD                      BillType = 715 // 715 自動充值用戶入款
	BILL_TYPE_TOP_UP_AUTO_RECEIVE_DEDUCT                   BillType = 716 // 716 自動充值平台扣款
	BILL_TYPE_TOP_UP_AUTO_SEND_EXCHANGE_FEE_DEDUCT         BillType = 717 // 717 自動充值用戶扣 匯差費
	BILL_TYPE_TOP_UP_AUTO_SEND_EXCHANGE_FEE_ADD            BillType = 718 // 718 自動充值手續費帳戶入帳匯差費
	BILL_TYPE_TOP_UP_WHALE_SEND_FREEZE                     BillType = 721 // 721 whale用戶凍結
	BILL_TYPE_TOP_UP_WHALE_SEND_FEE_FREEZE                 BillType = 722 // 722 whale用戶手續費凍結
	BILL_TYPE_TOP_UP_WHALE_SEND_FROZEN_DEDUCT              BillType = 723 // 723 whale用戶扣凍結款
	BILL_TYPE_TOP_UP_WHALE_SEND_FEE_FROZEN_DEDUCT          BillType = 724 // 724 whale用戶扣凍結手續費
	BILL_TYPE_TOP_UP_WHALE_SEND_ADD                        BillType = 725 // 725 whale平台帳戶入帳
	BILL_TYPE_TOP_UP_WHALE_SEND_FEE_ADD                    BillType = 726 // 726 whale手續費帳戶入帳手續費
	BILL_TYPE_TOP_UP_WHALE_RECEIVE_ADD                     BillType = 727 // 727 whale用戶入款
	BILL_TYPE_TOP_UP_WHALE_RECEIVE_DEDUCT                  BillType = 728 // 728 whale平台扣款
	BILL_TYPE_TOP_UP_WHALE_SEND_EXCHANGE_FEE_FREEZE        BillType = 729 // 729 whale用戶凍結匯差費
	BILL_TYPE_TOP_UP_WHALE_SEND_EXCHANGE_FEE_FROZEN_DEDUCT BillType = 730 // 730 whale用戶扣凍結匯差費
	BILL_TYPE_TOP_UP_WHALE_SEND_EXCHANGE_FEE_ADD           BillType = 731 // 731 whale手續費帳戶入帳匯差費
	BILL_TYPE_TOP_UP_WHALE_SEND_UNFREEZE                   BillType = 732 // 732 whale用戶解凍
	BILL_TYPE_TOP_UP_WHALE_SEND_FEE_UNFREEZE               BillType = 733 // 733 whale用戶手續費解凍
	BILL_TYPE_TOP_UP_WHALE_SEND_EXCHANGE_FEE_UNFREEZE      BillType = 734 // 734 whale用戶解凍匯差費
	BILL_TYPE_TOP_UP_PAYCRYPTO_SEND_DEDUCT                 BillType = 741 // 741 paycrypto用戶扣款
	BILL_TYPE_TOP_UP_PAYCRYPTO_SEND_FEE_DEDUCT             BillType = 742 // 742 paycrypto用戶扣手續費
	BILL_TYPE_TOP_UP_PAYCRYPTO_SEND_ADD                    BillType = 743 // 743 paycrypto平台帳戶入帳
	BILL_TYPE_TOP_UP_PAYCRYPTO_SEND_FEE_ADD                BillType = 744 // 744 paycrypto手續費帳戶入帳手續費
	BILL_TYPE_TOP_UP_PAYCRYPTO_RECEIVE_ADD                 BillType = 745 // 745 paycrypto用戶入款
	BILL_TYPE_TOP_UP_PAYCRYPTO_RECEIVE_DEDUCT              BillType = 746 // 746 paycrypto平台扣款
	BILL_TYPE_TOP_UP_PAYCRYPTO_SEND_EXCHANGE_FEE_DEDUCT    BillType = 747 // 747 paycrypto用戶扣匯差費
	BILL_TYPE_TOP_UP_PAYCRYPTO_SEND_EXCHANGE_FEE_ADD       BillType = 748 // 748 paycrypto手續費帳戶入帳匯差費

	BILL_TYPE_TOP_DOWN_SEND_DEDUCT                           BillType = 801 // 801 用戶扣款
	BILL_TYPE_TOP_DOWN_SEND_FEE_DEDUCT                       BillType = 802 // 802 用戶扣手續費
	BILL_TYPE_TOP_DOWN_SEND_ADD                              BillType = 803 // 803 平台帳戶入帳
	BILL_TYPE_TOP_DOWN_SEND_FEE_ADD                          BillType = 804 // 804 手續費帳戶入帳手續費
	BILL_TYPE_TOP_DOWN_RECEIVE_ADD                           BillType = 805 // 805 用戶入款
	BILL_TYPE_TOP_DOWN_RECEIVE_DEDUCT                        BillType = 806 // 806 平台扣款
	BILL_TYPE_TOP_DOWN_SEND_EXCHANGE_FEE_ADD                 BillType = 807 // 807 手續費帳戶入帳匯差
	BILL_TYPE_TOP_DOWN_SEND_EXCHANGE_FEE_DEDUCT              BillType = 808 // 802 用戶扣匯差
	BILL_TYPE_TOP_DOWN_WHALE_SEND_FREEZE                     BillType = 821 // 821 whale用戶凍結款
	BILL_TYPE_TOP_DOWN_WHALE_SEND_FROZEN_DEDUCT              BillType = 822 // 822 whale用戶扣凍結款
	BILL_TYPE_TOP_DOWN_WHALE_SEND_FEE_FREEZE                 BillType = 823 // 823 whale用戶凍結手續費
	BILL_TYPE_TOP_DOWN_WHALE_SEND_FEE_FROZEN_DEDUCT          BillType = 824 // 824 whale用戶扣凍結手續費
	BILL_TYPE_TOP_DOWN_WHALE_SEND_ADD                        BillType = 825 // 825 whale平台帳戶入帳
	BILL_TYPE_TOP_DOWN_WHALE_SEND_FEE_ADD                    BillType = 826 // 826 whale手續費帳戶入帳手續費
	BILL_TYPE_TOP_DOWN_WHALE_RECEIVE_ADD                     BillType = 827 // 827 whale用戶入款
	BILL_TYPE_TOP_DOWN_WHALE_RECEIVE_DEDUCT                  BillType = 828 // 828 whale平台扣款
	BILL_TYPE_TOP_DOWN_WHALE_SEND_EXCHANGE_FEE_ADD           BillType = 829 // 829 whale手續費帳戶入帳匯差
	BILL_TYPE_TOP_DOWN_WHALE_SEND_EXCHANGE_FEE_FREEZE        BillType = 830 // 830 whale用戶凍結匯差
	BILL_TYPE_TOP_DOWN_WHALE_SEND_EXCHANGE_FEE_FROZEN_DEDUCT BillType = 831 // 831 whale用戶扣凍結匯差
	BILL_TYPE_TOP_DOWN_WHALE_SEND_UNFREEZE                   BillType = 832 // 832 whale用戶結凍款
	BILL_TYPE_TOP_DOWN_WHALE_SEND_FEE_UNFREEZE               BillType = 833 // 833 whale用戶解凍手續費
	BILL_TYPE_TOP_DOWN_WHALE_SEND_EXCHANGE_FEE_UNFREEZE      BillType = 834 // 834 whale用戶解凍匯差
	BILL_TYPE_TOP_DOWN_DELETE_CARD_SEND_DEDUCT               BillType = 835 // 835 刪卡用戶扣款
	BILL_TYPE_TOP_DOWN_DELETE_CARD_SEND_FEE_DEDUCT           BillType = 836 // 836 刪卡用戶扣手續費
	BILL_TYPE_TOP_DOWN_DELETE_CARD_SEND_ADD                  BillType = 837 // 837 刪卡平台帳戶入帳
	BILL_TYPE_TOP_DOWN_DELETE_CARD_SEND_FEE_ADD              BillType = 838 // 838 刪卡手續費帳戶入帳手續費
	BILL_TYPE_TOP_DOWN_DELETE_CARD_RECEIVE_ADD               BillType = 839 // 839 刪卡用戶入款
	BILL_TYPE_TOP_DOWN_DELETE_CARD_RECEIVE_DEDUCT            BillType = 840 // 840 刪卡平台扣款
	BILL_TYPE_TOP_DOWN_DELETE_CARD_SEND_EXCHANGE_FEE_ADD     BillType = 841 // 841 刪卡手續費帳戶入帳匯差
	BILL_TYPE_TOP_DOWN_DELETE_CARD_SEND_EXCHANGE_FEE_DEDUCT  BillType = 842 // 842 刪卡用戶扣匯差
	BILL_TYPE_TOP_DOWN_ADMIN_DELETE_SEND_DEDUCT              BillType = 843 // 843 管理後台刪卡發方用戶扣款
	BILL_TYPE_TOP_DOWN_ADMIN_DELETE_RECEIVE_ADD              BillType = 844 // 844 管理後台刪卡收方用戶入帳
	BILL_TYPE_TOP_DOWN_ADMIN_DELETE_RECEIVE_DEDUCT           BillType = 845 // 845 管理後台刪卡平台帳戶扣款
	BILL_TYPE_TOP_DOWN_ADMIN_DELETE_SEND_ADD                 BillType = 846 // 846 管理後台刪卡平台帳戶入帳

	BILL_TYPE_REFUND_REAP_TRANSACTION_REFUND_DEDUCT BillType = 901 // 901 平台退款扣除
	BILL_TYPE_REFUND_REAP_TRANSACTION_REFUND_ADD    BillType = 902 // 902 用戶退款入帳

	BILL_TYPE_CARD_TO_CARD_SEND_DEDUCT                                          BillType = 1101 // 1101 發方用戶扣款
	BILL_TYPE_CARD_TO_CARD_SEND_FEE_DEDUCT                                      BillType = 1102 // 1102 發方用戶扣手續費
	BILL_TYPE_CARD_TO_CARD_SEND_ADD                                             BillType = 1103 // 1103 平台帳戶入帳
	BILL_TYPE_CARD_TO_CARD_SEND_FEE_ADD                                         BillType = 1104 // 1104 手續費帳戶入帳手續費
	BILL_TYPE_CARD_TO_CARD_RECEIVE_ADD                                          BillType = 1105 // 1105 收方用戶入款
	BILL_TYPE_CARD_TO_CARD_RECEIVE_DEDUCT                                       BillType = 1106 // 1106 平台扣款
	BILL_TYPE_CARD_TO_CARD_SEND_EXCHANGE_FEE_DEDUCT                             BillType = 1107 // 1107 發方用戶扣匯差費
	BILL_TYPE_CARD_TO_CARD_SEND_EXCHANGE_FEE_ADD                                BillType = 1108 // 1108 手續費帳戶入帳匯差費
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FREEZE                                    BillType = 1121 // 1121 whale發方用戶凍結款
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FROZEN_DEDUCT                             BillType = 1122 // 1122 whale發方用戶扣凍結款
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FEE_FREEZE                                BillType = 1123 // 1123 whale發方用戶凍結手續費
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FEE_FROZEN_DEDUCT                         BillType = 1124 // 1124 whale發方用戶扣凍結手續費
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_ADD                                       BillType = 1125 // 1125 whale平台帳戶入帳
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FEE_ADD                                   BillType = 1126 // 1126 whale手續費帳戶入帳手續費
	BILL_TYPE_CARD_TO_CARD_WHALE_RECEIVE_ADD                                    BillType = 1127 // 1127 whale收方用戶入款
	BILL_TYPE_CARD_TO_CARD_WHALE_RECEIVE_DEDUCT                                 BillType = 1128 // 1128 whale平台扣款
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_EXCHANGE_FEE_FREEZE                       BillType = 1129 // 1129 whale發方用戶凍結匯差費
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_EXCHANGE_FEE_FROZEN_DEDUCT                BillType = 1130 // 1130 whale發方用戶扣凍結匯差費
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_EXCHANGE_FEE_ADD                          BillType = 1131 // 1131 whale手續費帳戶入帳匯差費
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_UNFREEZE                                  BillType = 1132 // 1132 whale發方用戶解凍款
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FEE_UNFREEZE                              BillType = 1133 // 1133 whale發方用戶解凍手續費
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_EXCHANGE_FEE_UNFREEZE                     BillType = 1134 // 1134 whale發方用戶解凍匯差費
	BILL_TYPE_CARD_TO_CARD_WHALE_SELF_SEND_DEDUCT                               BillType = 1141 // 1141 whale自轉發方用戶扣款
	BILL_TYPE_CARD_TO_CARD_WHALE_SELF_SEND_FEE_DEDUCT                           BillType = 1142 // 1142 whale自轉發方用戶扣手續費
	BILL_TYPE_CARD_TO_CARD_WHALE_SELF_SEND_ADD                                  BillType = 1143 // 1143 whale自轉平台帳戶入帳
	BILL_TYPE_CARD_TO_CARD_WHALE_SELF_SEND_FEE_ADD                              BillType = 1144 // 1144 whale自轉手續費帳戶入帳手續費
	BILL_TYPE_CARD_TO_CARD_WHALE_SELF_RECEIVE_ADD                               BillType = 1145 // 1145 whale自轉收方用戶入款
	BILL_TYPE_CARD_TO_CARD_WHALE_SELF_RECEIVE_DEDUCT                            BillType = 1146 // 1146 whale自轉平台扣款
	BILL_TYPE_CARD_TO_CARD_WHALE_SELF_SEND_EXCHANGE_FEE_DEDUCT                  BillType = 1147 // 1147 whale自轉發方用戶扣匯差費
	BILL_TYPE_CARD_TO_CARD_WHALE_SELF_SEND_EXCHANGE_FEE_ADD                     BillType = 1148 // 1148 whale自轉手續費帳戶入帳匯差費
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FAILED_HALFWAY_FROZEN_DEDUCT              BillType = 1151 // 1151 whale半途失敗發方用戶扣凍結款
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FAILED_HALFWAY_ADD                        BillType = 1152 // 1152 whale半途失敗平台帳戶入帳
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FAILED_HALFWAY_FEE_FROZEN_DEDUCT          BillType = 1153 // 1153 whale半途失敗發方用戶扣凍結手續費
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FAILED_HALFWAY_FEE_ADD                    BillType = 1154 // 1154 whale半途失敗手續費帳戶入帳手續費
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FAILED_HALFWAY_EXCHANGE_FEE_FROZEN_DEDUCT BillType = 1155 // 1155 whale半途失敗發方用戶扣凍結匯差費
	BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FAILED_HALFWAY_EXCHANGE_FEE_ADD           BillType = 1156 // 1156 whale半途失敗手續費帳戶入帳匯差費
	BILL_TYPE_CARD_TO_CARD_ADMIN_DELETE_SEND_DEDUCT                             BillType = 1161 // 1161 管理後台刪卡發方用戶扣款
	BILL_TYPE_CARD_TO_CARD_ADMIN_DELETE_SEND_ADD                                BillType = 1162 // 1162 管理後台刪卡收方用戶入帳
	BILL_TYPE_CARD_TO_CARD_ETHERFI_SEND_ADD                                     BillType = 1163 // 1163 平台帳戶入帳
	BILL_TYPE_CARD_TO_CARD_ETHERFI_SEND_FEE_ADD                                 BillType = 1164 // 1164 手續費帳戶入帳手續費
	BILL_TYPE_CARD_TO_CARD_ETHERFI_RECEIVE_DEDUCT                               BillType = 1165 // 1165 平台扣款
	BILL_TYPE_CARD_TO_CARD_ETHERFI_SEND_DEDUCT                                  BillType = 1166 // 1166 etherfi 發方用戶扣款
	BILL_TYPE_CARD_TO_CARD_ETHERFI_SEND_FEE_DEDUCT                              BillType = 1167 // 1167 etherfi 發方用戶扣手續費
	BILL_TYPE_CARD_TO_CARD_ETHERFI_RECEIVE_ADD                                  BillType = 1168 // 1158 etherfi 收方用戶入款
	BILL_TYPE_CARD_TO_CARD_ETHERFI_SEND_EXCHANGE_FEE_ADD                        BillType = 1169 // 1169 手續費帳戶入帳匯差費
	BILL_TYPE_CARD_TO_CARD_ETHERFI_SEND_EXCHANGE_FEE_DEDUCT                     BillType = 1170 // 1170 發方用戶扣匯差費
	BILL_TYPE_CARD_TO_CARD_PAYCRYPTO_SEND_DEDUCT                                BillType = 1171 // 1171 paycrypto發方用戶扣款
	BILL_TYPE_CARD_TO_CARD_PAYCRYPTO_SEND_FEE_DEDUCT                            BillType = 1172 // 1172 paycrypto發方用戶扣手續費
	BILL_TYPE_CARD_TO_CARD_PAYCRYPTO_SEND_ADD                                   BillType = 1173 // 1173 paycrypto平台帳戶入帳
	BILL_TYPE_CARD_TO_CARD_PAYCRYPTO_SEND_FEE_ADD                               BillType = 1174 // 1174 paycrypto手續費帳戶入帳手續費
	BILL_TYPE_CARD_TO_CARD_PAYCRYPTO_RECEIVE_ADD                                BillType = 1175 // 1175 paycrypto收方用戶入款
	BILL_TYPE_CARD_TO_CARD_PAYCRYPTO_RECEIVE_DEDUCT                             BillType = 1176 // 1176 paycrypto平台扣款
	BILL_TYPE_CARD_TO_CARD_PAYCRYPTO_SEND_EXCHANGE_FEE_DEDUCT                   BillType = 1177 // 1177 paycrypto發方用戶扣匯差費
	BILL_TYPE_CARD_TO_CARD_PAYCRYPTO_SEND_EXCHANGE_FEE_ADD                      BillType = 1178 // 1178 paycrypto手續費帳戶入帳匯差費

	BILL_TYPE_MANUAL_ADD_ADD                     BillType = 1201 // 1201 加錢帳變 用戶加錢
	BILL_TYPE_MANUAL_ADD_DEDUCT                  BillType = 1202 // 1202 加錢帳變 對方扣錢
	BILL_TYPE_MANUAL_DEDUCT_DEDUCT               BillType = 1203 // 1203 扣錢帳變 用戶扣錢
	BILL_TYPE_MANUAL_DEDUCT_ADD                  BillType = 1204 // 1204 扣錢帳變 對方加錢
	BILL_TYPE_MANUAL_FREEZE_FREEZE               BillType = 1205 // 1205 凍結帳變 用戶凍結
	BILL_TYPE_MANUAL_UNFREEZE_UNFREEZE           BillType = 1206 // 1206 解凍帳變 對方解凍
	BILL_TYPE_MANUAL_ADD_FREEZE_ADD_FREEZE       BillType = 1207 // 1207 加凍結帳變 用戶加凍結
	BILL_TYPE_MANUAL_ADD_FREEZE_DEDUCT           BillType = 1208 // 1208 加凍結帳變 對方扣錢
	BILL_TYPE_MANUAL_DEDUCT_FREEZE_DEDUCT_FREEZE BillType = 1209 // 1209 扣凍結帳變 用戶扣凍結
	BILL_TYPE_MANUAL_DEDUCT_FREEZE_ADD           BillType = 1210 // 1210 扣凍結帳變 對方加錢

	BILL_TYPE_MERCHANT_AUTO_EXCHANGE_TIMED_SEND_DEDUCT     BillType = 1407 // 1407 商戶自動兌換 定時商戶扣款
	BILL_TYPE_MERCHANT_AUTO_EXCHANGE_TIMED_SEND_FEE_DEDUCT BillType = 1408 // 1408 商戶自動兌換 定時商戶扣手續費
	BILL_TYPE_MERCHANT_AUTO_EXCHANGE_TIMED_SEND_ADD        BillType = 1409 // 1409 商戶自動兌換 定時平台帳戶入帳
	BILL_TYPE_MERCHANT_AUTO_EXCHANGE_TIMED_SEND_FEE_ADD    BillType = 1410 // 1410 商戶自動兌換 定時手續費帳戶入帳手續費
	BILL_TYPE_MERCHANT_AUTO_EXCHANGE_TIMED_RECEIVE_ADD     BillType = 1411 // 1411 商戶自動兌換 定時商戶入款
	BILL_TYPE_MERCHANT_AUTO_EXCHANGE_TIMED_RECEIVE_DEDUCT  BillType = 1412 // 1412 商戶自動兌換 定時平台扣款

	BILL_TYPE_CHIPPAY_EXPRESS_RECEIVE_ADD              BillType = 1501 // 1501 用戶入帳
	BILL_TYPE_CHIPPAY_EXPRESS_RECEIVE_DEDUCT           BillType = 1502 // 1502 平台扣款
	BILL_TYPE_CHIPPAY_EXPRESS_FEE_DEDUCT               BillType = 1503 // 1503 平台扣手續費
	BILL_TYPE_CHIPPAY_EXPRESS_FEE_ADD                  BillType = 1504 // 1504 收續費帳戶入帳手續費
	BILL_TYPE_CHIPPAY_EXPRESS_PAYCRYPTO_RECEIVE_ADD    BillType = 1505 // 1505 paycrypto用戶入帳
	BILL_TYPE_CHIPPAY_EXPRESS_PAYCRYPTO_RECEIVE_DEDUCT BillType = 1506 // 1506 paycrypto平台扣款
	BILL_TYPE_CHIPPAY_EXPRESS_PAYCRYPTO_FEE_DEDUCT     BillType = 1507 // 1507 paycrypto平台扣手續費
	BILL_TYPE_CHIPPAY_EXPRESS_PAYCRYPTO_FEE_ADD        BillType = 1508 // 1508 paycrypto收續費帳戶入帳手續費

	BILL_TYPE_MERCHANT_AUTO_EXCHANGE_ON_DEPOSIT_SEND_DEDUCT     BillType = 1421 // 1421 商戶自動兌換 儲值時商戶扣款
	BILL_TYPE_MERCHANT_AUTO_EXCHANGE_ON_DEPOSIT_SEND_FEE_DEDUCT BillType = 1422 // 1422 商戶自動兌換 儲值時商戶扣手續費
	BILL_TYPE_MERCHANT_AUTO_EXCHANGE_ON_DEPOSIT_SEND_ADD        BillType = 1423 // 1423 商戶自動兌換 儲值時平台帳戶入帳
	BILL_TYPE_MERCHANT_AUTO_EXCHANGE_ON_DEPOSIT_SEND_FEE_ADD    BillType = 1424 // 1424 商戶自動兌換 儲值時手續費帳戶入帳手續費
	BILL_TYPE_MERCHANT_AUTO_EXCHANGE_ON_DEPOSIT_RECEIVE_ADD     BillType = 1425 // 1425 商戶自動兌換 儲值時商戶入款
	BILL_TYPE_MERCHANT_AUTO_EXCHANGE_ON_DEPOSIT_RECEIVE_DEDUCT  BillType = 1426 // 1426 商戶自動兌換 儲值時平台扣款

	BILL_TYPE_MERCHANT_ADJUST_BALANCE_POSITIVE_SEND_DEDUCT        BillType = 1601 // 1601 加錢時商戶扣款
	BILL_TYPE_MERCHANT_ADJUST_BALANCE_POSITIVE_SEND_FEE_DEDUCT    BillType = 1602 // 1602 加錢時商戶扣手續費
	BILL_TYPE_MERCHANT_ADJUST_BALANCE_POSITIVE_SEND_ADD           BillType = 1603 // 1603 加錢時平台帳戶入帳
	BILL_TYPE_MERCHANT_ADJUST_BALANCE_POSITIVE_SEND_FEE_ADD       BillType = 1604 // 1604 加錢時手續費帳戶入帳手續費
	BILL_TYPE_MERCHANT_ADJUST_BALANCE_POSITIVE_RECEIVE_ADD        BillType = 1605 // 1605 加錢時商戶子用戶入帳
	BILL_TYPE_MERCHANT_ADJUST_BALANCE_POSITIVE_RECEIVE_DEDUCT     BillType = 1606 // 1606 加錢時平台扣商戶子用戶入帳金額
	BILL_TYPE_MERCHANT_ADJUST_BALANCE_NEGATIVE_RECEIVE_ADD        BillType = 1611 // 1611 扣錢時商戶入帳
	BILL_TYPE_MERCHANT_ADJUST_BALANCE_NEGATIVE_RECEIVE_FEE_DEDUCT BillType = 1612 // 1612 扣錢時商戶扣手續費
	BILL_TYPE_MERCHANT_ADJUST_BALANCE_NEGATIVE_RECEIVE_DEDUCT     BillType = 1613 // 1613 扣錢時平台帳戶扣款
	BILL_TYPE_MERCHANT_ADJUST_BALANCE_NEGATIVE_RECEIVE_FEE_ADD    BillType = 1614 // 1614 扣錢時手續費帳戶入帳手續費
	BILL_TYPE_MERCHANT_ADJUST_BALANCE_NEGATIVE_SEND_DEDUCT        BillType = 1615 // 1615 扣錢時商戶子用戶扣款
	BILL_TYPE_MERCHANT_ADJUST_BALANCE_NEGATIVE_SEND_ADD           BillType = 1616 // 1616 扣錢時平台加商戶子用戶扣款金額

	BILL_TYPE_POINT_ACCURAL_POINT_ADD                    BillType = 1701 // 1701 用戶加點數
	BILL_TYPE_POINT_ACCURAL_POINT_DEDUCT                 BillType = 1702 // 1702 平台帳戶扣點數
	BILL_TYPE_POINT_ACCURAL_WHALE_POINT_ADD              BillType = 1703 // 1703 whale用戶加點數
	BILL_TYPE_POINT_ACCURAL_WHALE_POINT_DEDUCT           BillType = 1704 // 1704 whale平台帳戶扣點數
	BILL_TYPE_POINT_ACCURAL_CHIPPAY_EXPRESS_POINT_ADD    BillType = 1705 // 1705 cp用戶加點數
	BILL_TYPE_POINT_ACCURAL_CHIPPAY_EXPRESS_POINT_DEDUCT BillType = 1706 // 1706 cp平台帳戶扣點數
	BILL_TYPE_POINT_ACCURAL_PAYCRYPTO_POINT_ADD          BillType = 1707 // 1707 paycrypto用戶加點數
	BILL_TYPE_POINT_ACCURAL_PAYCRYPTO_POINT_DEDUCT       BillType = 1708 // 1708 paycrypto平台帳戶扣點數

	BILL_TYPE_POINT_REDEEM_POINT_DEDUCT BillType = 1801 // 1801 用戶扣款
	BILL_TYPE_POINT_REDEEM_POINT_ADD    BillType = 1802 // 1802 平台帳戶入帳

	BILL_TYPE_INTEREST_ADD    BillType = 1901 // 1901 用戶加利息
	BILL_TYPE_INTEREST_DEDUCT BillType = 1902 // 1902 利息帳戶扣利息

	BILL_TYPE_FINANCIAL_TRANSFER_SEND_DEDUCT BillType = 2001 // 2001 發方卡片扣款
	BILL_TYPE_FINANCIAL_TRANSFER_RECEIVE_ADD BillType = 2002 // 2002 收方卡片入款

	BILL_TYPE_WHALE_PAY_CONSUMPTION_PENDING_FREEZE           BillType = 2101 // 2101 用戶消費凍結
	BILL_TYPE_WHALE_PAY_CONSUMPTION_CLOSED_DEDUCT_FROZEN     BillType = 2102 // 2102 用戶凍結扣除
	BILL_TYPE_WHALE_PAY_CONSUMPTION_CLOSED_ADD               BillType = 2103 // 2103 平台入帳
	BILL_TYPE_WHALE_PAY_CONSUMPTION_CLOSED_UNFREEZE          BillType = 2104 // 2104 用戶解凍
	BILL_TYPE_WHALE_PAY_REVERSAL_DEDUCT                      BillType = 2105 // 2105 平台回沖扣除
	BILL_TYPE_WHALE_PAY_REVERSAL_ADD                         BillType = 2106 // 2106 用戶回沖入帳
	BILL_TYPE_WHALE_PAY_REVERSAL_FEE_DEDUCT                  BillType = 2107 // 2107 用戶回沖手續費扣除
	BILL_TYPE_WHALE_PAY_REVERSAL_FEE_ADD                     BillType = 2108 // 2108 平台回沖手續費入帳
	BILL_TYPE_WHALE_PAY_CREDIT_DEDUCT                        BillType = 2109 // 2109 平台退款扣除
	BILL_TYPE_WHALE_PAY_CREDIT_ADD                           BillType = 2110 // 2110 用戶退款入帳
	BILL_TYPE_WHALE_PAY_CREDIT_FEE_DEDUCT                    BillType = 2111 // 2111 平台退款手續費扣除
	BILL_TYPE_WHALE_PAY_CREDIT_FEE_ADD                       BillType = 2112 // 2112 平台退款手續費入帳
	BILL_TYPE_WHALE_PAY_CONSUMPTION_PENDING_FEE_FREEZE       BillType = 2113 // 2113 用戶消費凍結手續費
	BILL_TYPE_WHALE_PAY_CONSUMPTION_CLOSED_FEE_DEDUCT_FROZEN BillType = 2114 // 2114 用戶消費凍結手續費扣除
	BILL_TYPE_WHALE_PAY_CONSUMPTION_CLOSED_FEE_ADD           BillType = 2115 // 2115 用戶消費手續費入帳
	BILL_TYPE_WHALE_PAY_CONSUMPTION_DIRECT_CLOSED_DEDUCT     BillType = 2116 // 2116 用戶直接結清扣除
	BILL_TYPE_WHALE_PAY_CONSUMPTION_DIRECT_CLOSED_ADD        BillType = 2117 // 2117 平台直接結清入帳
	BILL_TYPE_WHALE_PAY_CONSUMPTION_DIRECT_CLOSED_FEE_DEDUCT BillType = 2118 // 2118 用戶直接結清手續費扣除
	BILL_TYPE_WHALE_PAY_CONSUMPTION_DIRECT_CLOSED_FEE_ADD    BillType = 2119 // 2119 用戶直接結清手續費入帳
	BILL_TYPE_WHALE_PAY_CONSUMPTION_FAIL_UNFREEZE            BillType = 2120 // 2120 用戶交易失敗解凍
	BILL_TYPE_WHALE_PAY_CONSUMPTION_FAIL_FEE_UNFREEZE        BillType = 2121 // 2121 用戶交易失敗解凍手續費

	BILL_TYPE_DISTRIBUTE_DISCOUNT_DEDUCT BillType = 2301 // 2301 開卡佣金扣除
	BILL_TYPE_DISTRIBUTE_DISCOUNT_ADD    BillType = 2302 // 2302 開卡佣金入帳

	BILL_TYPE_PAYCRYPTO_PAY_FAIL_DEDUCT     BillType = 2401 // 2401 用戶消費失敗扣除
	BILL_TYPE_PAYCRYPTO_PAY_FAIL_ADD        BillType = 2402 // 2402 平台入帳用戶消費失敗
	BILL_TYPE_PAYCRYPTO_PAY_REVERSAL_DEDUCT BillType = 2403 // 2403 平台回沖扣除
	BILL_TYPE_PAYCRYPTO_PAY_REVERSAL_ADD    BillType = 2404 // 2404 用戶回沖入帳
	BILL_TYPE_PAYCRYPTO_PAY_REFUND_DEDUCT   BillType = 2405 // 2405 平台退款扣除
	BILL_TYPE_PAYCRYPTO_PAY_REFUND_ADD      BillType = 2406 // 2406 用戶退款入帳
	BILL_TYPE_PAYCRYPTO_PAY_PAY_DEDUCT      BillType = 2407 // 2407 用戶消費扣除
	BILL_TYPE_PAYCRYPTO_PAY_PAY_ADD         BillType = 2408 // 2408 平台入帳用戶消費
	BILL_TYPE_PAYCRYPTO_PAY_CLOSE_DEDUCT    BillType = 2409 // 2409 用戶結清調整扣除
	BILL_TYPE_PAYCRYPTO_PAY_CLOSE_ADD       BillType = 2410 // 2410 平台入帳用戶結清調整

	BILL_TYPE_BINANCE_PAY_DEPOSIT_RECEIVE_DEDUCT BillType = 2501 // 2501 平台扣幣
	BILL_TYPE_BINANCE_PAY_DEPOSIT_FEE_DEDUCT     BillType = 2502 // 2502 平台扣手續費
	BILL_TYPE_BINANCE_PAY_DEPOSIT_FEE_ADD        BillType = 2503 // 2503 手續費帳戶加手續費
	BILL_TYPE_BINANCE_PAY_DEPOSIT_RECEIVE_ADD    BillType = 2504 // 2504 用戶加幣

	BILL_TYPE_SYSTEM_CHARGE_DECLINE_FINE_DEDUCT BillType = 2601 // 2601 拒刷罰款用戶扣錢
	BILL_TYPE_SYSTEM_CHARGE_DECLINE_FINE_ADD    BillType = 2602 // 2602 拒刷罰款平台收錢

	BILL_TYPE_ADMIN_USER_DELETE_ADD          BillType = 2701 // 2701 平台入帳
	BILL_TYPE_ADMIN_USER_DELETE_DEDUT        BillType = 2702 // 2702 用戶扣除
	BILL_TYPE_ADMIN_USER_DELETE_FREEZE_ADD   BillType = 2703 // 2703 平台凍結款入帳
	BILL_TYPE_ADMIN_USER_DELETE_FREEZE_DEDUT BillType = 2704 // 2704 用戶凍結款扣除

	BILL_TYPE_ADMIN_CARD_DELETE_ADD   BillType = 2801 // 2801 平台入帳
	BILL_TYPE_ADMIN_CARD_DELETE_DEDUT BillType = 2802 // 2802 用戶扣除

	BILL_TYPE_ETHERFI_RECEIVE_ADD      BillType = 3301
	BILL_TYPE_ETHERFI_RECEIVE_DEDUCT   BillType = 3302
	BILL_TYPE_ETHERFI_RECEIVE_FREEZE   BillType = 3303
	BILL_TYPE_ETHERFI_RECEIVE_UNFREEZE BillType = 3304
)

type BusinessErrorArg string // 业务错误参数
const (
	BUSINESS_ERROR_ARG_ATTEMPT BusinessErrorArg = "ATTEMPT"
)

type CallbackCriteria int // 回調成功判斷
const (
	CALLBACK_CRITERIA_STATUS_CODE_200 CallbackCriteria = 1
	CALLBACK_CRITERIA_CODE_OK         CallbackCriteria = 2
	CALLBACK_CRITERIA_MSG_OK          CallbackCriteria = 3
)

type CallbackScene int // 回調場景
const (
	CALLBACK_SCENE_REQ_AUTH              CallbackScene = 10
	CALLBACK_SCENE_TX_AUTH               CallbackScene = 11
	CALLBACK_SCENE_TX_ADV                CallbackScene = 12
	CALLBACK_SCENE_TX_REV                CallbackScene = 13
	CALLBACK_SCENE_TX_REF                CallbackScene = 14
	CALLBACK_SCENE_TX_CLR                CallbackScene = 15
	CALLBACK_SCENE_VERIFY                CallbackScene = 16
	CALLBACK_SCENE_VERIFIED              CallbackScene = 17
	CALLBACK_SCENE_3DS                   CallbackScene = 20
	CALLBACK_SCENE_WALLET_OTP            CallbackScene = 30
	CALLBACK_SCENE_CARD_STATUS           CallbackScene = 40
	CALLBACK_SCENE_CARD_SHIPPING_STATUS  CallbackScene = 41
	CALLBACK_SCENE_CARD_FREEZE_STATUS    CallbackScene = 42
	CALLBACK_SCENE_BALANCE_CHANGE_EXPORT CallbackScene = 50
	CALLBACK_SCENE_KYC_RESULT            CallbackScene = 60
)

type CallbackStatus int // 回調狀態
const (
	CALLBACK_STATUS_PENDING       CallbackStatus = 1
	CALLBACK_STATUS_SUCCESS       CallbackStatus = 2
	CALLBACK_STATUS_PENDING_RETRY CallbackStatus = 3
	CALLBACK_STATUS_FAILED        CallbackStatus = 4
)

type CallbackType int // 回調類型
const (
	CALLBACK_TYPE_MERCHANT_PAY         CallbackType = 1
	CALLBACK_TYPE_MERCHANT_3DS         CallbackType = 2
	CALLBACK_TYPE_MERCHANT_WALLET_OTP  CallbackType = 3
	CALLBACK_TYPE_MERCHANT_CARD_STATUS CallbackType = 4
	CALLBACK_TYPE_MERCHANT_EXPORT      CallbackType = 5
	CALLBACK_TYPE_MERCHANT_KYC_RESULT  CallbackType = 6
)

type CardStatus int // 卡片狀態
const (
	CARD_STATUS_NOT_CREATED   CardStatus = 1
	CARD_STATUS_CREATED       CardStatus = 2
	CARD_STATUS_NOT_ACTIVATED CardStatus = 3
	CARD_STATUS_ACTIVATED     CardStatus = 4
	CARD_STATUS_DEACTIVATED   CardStatus = 5
	CARD_STATUS_DELETED       CardStatus = 6
	CARD_STATUS_BLOCKED       CardStatus = 7
)

type DisplayStatus string // 前端卡片狀態
const (
	DISPLAY_STATUS_ACTIVATED               DisplayStatus = "activated"
	DISPLAY_STATUS_NOT_ACTIVATED           DisplayStatus = "not_activated"
	DISPLAY_STATUS_BLOCKED                 DisplayStatus = "blocked"
	DISPLAY_STATUS_FROZEN                  DisplayStatus = "frozen"
	DISPLAY_STATUS_CREATED                 DisplayStatus = "created"
	DISPLAY_STATUS_NOT_CREATED             DisplayStatus = "not_created"
	DISPLAY_STATUS_BLOCKED_DUE_TO_DECLINES DisplayStatus = "blocked_due_to_declines"
)

type CardEarningStatus int // 卡片計息狀態
const (
	CARD_EARNING_STATUS_ENABLED  CardEarningStatus = 1
	CARD_EARNING_STATUS_DISABLED CardEarningStatus = 2
)

type CardFreezeStatus int // 卡片狀態
const (
	CARD_FREEZE_STATUS_UNFROZEN CardFreezeStatus = 1
	CARD_FREEZE_STATUS_FROZEN   CardFreezeStatus = 2
)

type CardRiskyStatus int // 卡片風險狀態
const (
	CARD_RISKY_STATUS_UNRISKY CardRiskyStatus = 1
	CARD_RISKY_STATUS_RISKY   CardRiskyStatus = 2
)

type CardFormat int // 卡片格式
const (
	CARD_FORMAT_PHYSICAL CardFormat = 1
	CARD_FORMAT_VIRTUAL  CardFormat = 2
	CARD_FORMAT_GENERAL  CardFormat = 3
)

type CardOrganization int // 信用卡組織
const (
	CARD_ORGANIZATION_VISA       CardOrganization = 1
	CARD_ORGANIZATION_MASTERCARD CardOrganization = 2
	CARD_ORGANIZATION_AMEX       CardOrganization = 3
	CARD_ORGANIZATION_DISCOVER   CardOrganization = 4
	CARD_ORGANIZATION_JCB        CardOrganization = 5
	CARD_ORGANIZATION_DINERS     CardOrganization = 6
	CARD_ORGANIZATION_UNIONPAY   CardOrganization = 7
)

type CardProductVendor int // 卡片廠商
const (
	CARD_PRODUCT_VENDOR_REAP      CardProductVendor = 1
	CARD_PRODUCT_VENDOR_WHALE     CardProductVendor = 2
	CARD_PRODUCT_VENDOR_PAYCRYPTO CardProductVendor = 3
	CARD_PRODUCT_VENDOR_ETHERFI   CardProductVendor = 4
)

type CardToCardStatus int // 轉卡狀態
const (
	CARD_TO_CARD_STATUS_PENDING        CardToCardStatus = 1
	CARD_TO_CARD_STATUS_SUCCESS        CardToCardStatus = 2
	CARD_TO_CARD_STATUS_FAILED         CardToCardStatus = 3
	CARD_TO_CARD_STATUS_PROCESSING     CardToCardStatus = 4
	CARD_TO_CARD_STATUS_FAILED_HALFWAY CardToCardStatus = 5
)

type CaseSensitive int // 區分大小寫
const (
	CASE_SENSITIVE_YES CaseSensitive = 1
	CASE_SENSITIVE_NO  CaseSensitive = 2
)

type CoinsdoResponseCode int // coinsdo回傳碼
const (
	COINSDO_RESPONSE_CODE_TOKEN_NOT_SUPPORTED                            CoinsdoResponseCode = 4006 // 幣種不支援
	COINSDO_RESPONSE_CODE_BUSINESS_ID_ALREADY_EXIST                      CoinsdoResponseCode = 4008 // 業務id已經存在
	COINSDO_RESPONSE_CODE_FAILED_TO_ADD_NEW_DISPATCH_RECORD              CoinsdoResponseCode = 4009 // 新增派發記錄失敗
	COINSDO_RESPONSE_CODE_FAILED_TO_VERIFY_DEVICE_FOR_APPROVAL           CoinsdoResponseCode = 4011 // 驗證設備以進行批准失敗
	COINSDO_RESPONSE_CODE_FIRST_APPROVER_DEVICE_NOT_EXISTS               CoinsdoResponseCode = 4012 // 數據批准者的（第一批准者）設備不存在
	COINSDO_RESPONSE_CODE_APPROVAL_DEVICE_NOT_BOUND                      CoinsdoResponseCode = 4013 // 批准設備未綁定至批准帳戶
	COINSDO_RESPONSE_CODE_FIRST_APPROVER_SIGNATURE_FAILED                CoinsdoResponseCode = 4014 // 數據批准者（第一批准者）的簽名驗證失敗
	COINSDO_RESPONSE_CODE_AUTHORIZATION_CODE_VERIFICATION_FAILED         CoinsdoResponseCode = 4015 // 授權碼驗證失敗
	COINSDO_RESPONSE_CODE_AUTHORIZATION_CODE_EXPIRED                     CoinsdoResponseCode = 4016 // 授權碼已過期或不存在
	COINSDO_RESPONSE_CODE_DEVICE_VERIFICATION_FAILED                     CoinsdoResponseCode = 4017 // 設備驗證失敗
	COINSDO_RESPONSE_CODE_ADDING_NEW_APPROVAL_DEVICE_FAILED              CoinsdoResponseCode = 4018 // 新增批准設備失敗
	COINSDO_RESPONSE_CODE_SIGNATURE_VERIFICATION_FAILED                  CoinsdoResponseCode = 4019 // 通過SHA256與RSA的簽名驗證失敗
	COINSDO_RESPONSE_CODE_CORPORATE_SUB_ACCOUNT_NOT_FOUND                CoinsdoResponseCode = 4020 // 找不到企業子帳戶
	COINSDO_RESPONSE_CODE_NO_DATA_SUBMISSION_RIGHTS                      CoinsdoResponseCode = 4021 // 沒有數據提交權限
	COINSDO_RESPONSE_CODE_NO_FIRST_APPROVER_RIGHTS                       CoinsdoResponseCode = 4022 // 沒有數據批准者（第一批准者）權限
	COINSDO_RESPONSE_CODE_TOKEN_LARGEST_DECIMAL_EXCEEDED                 CoinsdoResponseCode = 4023 // 超過代幣的最大十進制
	COINSDO_RESPONSE_CODE_INCOMPLETE_PARAMETERS_FIRST_APPROVER           CoinsdoResponseCode = 4024 // 執行批准者（第一批准者）執行時找到不完整參數
	COINSDO_RESPONSE_CODE_DISPATCH_ADDRESS_WRONG_FORMAT                  CoinsdoResponseCode = 4025 // 派發地址格式錯誤
	COINSDO_RESPONSE_CODE_INCOMPLETE_PARAMETERS_SECOND_APPROVER          CoinsdoResponseCode = 4026 // 執行批准者（第二批准者）執行時找到不完整參數
	COINSDO_RESPONSE_CODE_SECOND_APPROVER_DEVICE_NOT_EXISTS              CoinsdoResponseCode = 4027 // 執行批准者（第二批准者）的設備不存在
	COINSDO_RESPONSE_CODE_VERIFY_SECOND_APPROVER_SIGNATURE_FAILED        CoinsdoResponseCode = 4028 // 驗證執行批准者（第二批准者）的簽名失敗
	COINSDO_RESPONSE_CODE_NO_RIGHTS_TO_EXECUTE_APPROVAL                  CoinsdoResponseCode = 4029 // 沒有執行批准（第二批准）的權限
	COINSDO_RESPONSE_CODE_DISPATCH_AMOUNT_EXCEEDED_FIRST_APPROVER_LIMIT  CoinsdoResponseCode = 4030 // 派發金額超過數據批准者（第一批准者）角色的限制
	COINSDO_RESPONSE_CODE_DISPATCH_AMOUNT_EXCEEDED_SECOND_APPROVER_LIMIT CoinsdoResponseCode = 4031 // 派發金額超過執行批准者（第二批准者）角色的限制
	COINSDO_RESPONSE_CODE_FAILED_TO_VERIFY_WALLET                        CoinsdoResponseCode = 4032 // 驗證錢包失敗
	COINSDO_RESPONSE_CODE_WALLET_NOT_EXISTS                              CoinsdoResponseCode = 4033 // 錢包不存在
	COINSDO_RESPONSE_CODE_FAILED_TO_VERIFY_DISPATCH_ADDRESS              CoinsdoResponseCode = 4038 // 驗證派發地址失敗
	COINSDO_RESPONSE_CODE_WITHDRAWAL_ACCOUNT_DEACTIVATED                 CoinsdoResponseCode = 4039 // 提現帳戶已被停用
	COINSDO_RESPONSE_CODE_FIRST_APPROVER_DEVICE_DEACTIVATED              CoinsdoResponseCode = 4040 // 數據批准者（第一批准者）的綁定設備已被停用
	COINSDO_RESPONSE_CODE_SECOND_APPROVER_DEVICE_DEACTIVATED             CoinsdoResponseCode = 4041 // 執行批准者（第二批准者）的綁定設備已被停用
	COINSDO_RESPONSE_CODE_FAILED_TO_VERIFY_WHITELISTED_ADDRESS           CoinsdoResponseCode = 4046 // 驗證白名單派發地址失敗
	COINSDO_RESPONSE_CODE_DISPATCH_ADDRESS_NOT_FOUND                     CoinsdoResponseCode = 4047 // 派發地址不在白名單中
	COINSDO_RESPONSE_CODE_DISPATCH_AMOUNT_BELOW_MINIMUM                  CoinsdoResponseCode = 4048 // 派發金額低於最低要求金額
	COINSDO_RESPONSE_CODE_DATA_SUBMITTER_NOT_API_ACCOUNT                 CoinsdoResponseCode = 4049 // 數據提交者帳戶不屬於API帳戶
	COINSDO_RESPONSE_CODE_APPROVAL_DEVICE_NO_REQUIRED_PERMISSION         CoinsdoResponseCode = 4056 // 與數據批准帳戶關聯的批准設備沒有所需的權限
	COINSDO_RESPONSE_CODE_EXECUTION_DEVICE_NO_REQUIRED_PERMISSION        CoinsdoResponseCode = 4057 // 與執行批准帳戶關聯的批准設備沒有所需的權限
	COINSDO_RESPONSE_CODE_FAILED_TO_VERIFY_BUSINESS_ID                   CoinsdoResponseCode = 4065 // 驗證業務ID失敗
	COINSDO_RESPONSE_CODE_FAILED_TO_VERIFY_SUPPORTED_TOKENS              CoinsdoResponseCode = 4066 // 驗證支持的代幣失敗
	COINSDO_RESPONSE_CODE_COINS_DO_ID_MUST_BE_WHOLE_NUMBER               CoinsdoResponseCode = 4067 // CoinsDo ID 必須是整數
	COINSDO_RESPONSE_CODE_SUB_ACCOUNT_NOT_EXISTS                         CoinsdoResponseCode = 4069 // 子帳戶不存在
	COINSDO_RESPONSE_CODE_SELECT_A_COIN                                  CoinsdoResponseCode = 4070 // 請選擇一種幣
)

type ChippayResponseCode int // coinsdo回傳碼
const (
	CHIPPAY_RESPONSE_CODE_OK               ChippayResponseCode = 200
	CHIPPAY_RESPONSE_CODE_NO_SUCH_USER     ChippayResponseCode = 6024
	CHIPPAY_RESPONSE_CODE_INCOMPLETE_ORDER ChippayResponseCode = 6044
	CHIPPAY_RESPONSE_CODE_AMOUNT_TOO_LOW   ChippayResponseCode = 6062
	CHIPPAY_RESPONSE_CODE_USER_FROZEN      ChippayResponseCode = 6081
	CHIPPAY_RESPONSE_CODE_AMOUNT_TOO_HIGH  ChippayResponseCode = 6091
	CHIPPAY_RESPONSE_CODE_INVALID_NAME     ChippayResponseCode = 6093
)

type CoinsdoWithdrawStatus string // coinsdo提幣狀態
const (
	COINSDO_WITHDRAW_STATUS_APPROVED  CoinsdoWithdrawStatus = "1"
	COINSDO_WITHDRAW_STATUS_REJECTED  CoinsdoWithdrawStatus = "2"
	COINSDO_WITHDRAW_STATUS_CANCELLED CoinsdoWithdrawStatus = "5"
)

type CoinfaceMain int // coinface 主要驗證帳號
const (
	COINFACE_MAIN_YES CoinfaceMain = 1
	COINFACE_MAIN_NO  CoinfaceMain = 2
)

type CoinTokenType int // 主幣或代幣
const (
	COIN_TOKEN_TYPE_BASE  CoinTokenType = 1
	COIN_TOKEN_TYPE_TOKEN CoinTokenType = 2
)

type ContentChannel int // 文案渠道
const (
	CONTENT_CHANNEL_APP ContentChannel = 1
	CONTENT_CHANNEL_WEB ContentChannel = 2
)

type ContentScene int // 文案場景
const (
	CONTENT_SCENE_GENERAL             ContentScene = 1
	CONTENT_SCENE_ERROR_CODE          ContentScene = 2
	CONTENT_SCENE_CARD_FORMAT         ContentScene = 100
	CONTENT_SCENE_REWARD_PRODUCT      ContentScene = 200
	CONTENT_SCENE_CRYPTO_CARD         ContentScene = 300
	CONTENT_SCENE_PAY_METHOD          ContentScene = 400
	CONTENT_SCENE_WHALE_CARD_TYPE_BIN ContentScene = 500
	CONTENT_SCENE_PAYCRYPTO_CARD      ContentScene = 600
	CONTENT_SCENE_ETHERFI_CARD        ContentScene = 700
)

type ContentStatus int // 結構化文案狀態
const (
	CONTENT_STATUS_ACTIVATED   ContentStatus = 1
	CONTENT_STATUS_DEACTIVATED ContentStatus = 2
)

type CPIdentityDocumentType int // 身分證類型
const (
	CP_IDENTITY_DOCUMENT_TYPE_ID_CARD  CPIdentityDocumentType = 1
	CP_IDENTITY_DOCUMENT_TYPE_PASSPORT CPIdentityDocumentType = 2
)

type CPExpressDirection int // chippay快捷買單方向
const (
	CP_EXPRESS_DIRECTION_USER_BUY  CPExpressDirection = 1
	CP_EXPRESS_DIRECTION_USER_SELL CPExpressDirection = 2
)

type CPExpressStatus int // chippay快捷買單狀態
const (
	CP_EXPRESS_STATUS_PENDING          CPExpressStatus = 1
	CP_EXPRESS_STATUS_SUCCESS          CPExpressStatus = 2
	CP_EXPRESS_STATUS_FAILED           CPExpressStatus = 3
	CP_EXPRESS_STATUS_CHIPPAY_FAILED   CPExpressStatus = 4
	CP_EXPRESS_STATUS_CHIPPAY_CANCEL   CPExpressStatus = 5
	CP_EXPRESS_STATUS_PAYCRYPTO_FAILED CPExpressStatus = 6
)

type CPExpressWebhookStatus int // chippay回調狀態
const (
	CP_EXPRESS_WEBHOOK_STATUS_FAILED       CPExpressWebhookStatus = 0
	CP_EXPRESS_WEBHOOK_STATUS_SUCCESS      CPExpressWebhookStatus = 1
	CP_EXPRESS_WEBHOOK_STATUS_BATCH_FAILED CPExpressWebhookStatus = 2
)

type CPPaymentMethod int // chippay支付途徑
const (
	CPPAYMENT_METHOD_WECHAT_PAY CPPaymentMethod = 1
	CPPAYMENT_METHOD_ALIPAY     CPPaymentMethod = 2
	CPPAYMENT_METHOD_BANK       CPPaymentMethod = 3
)

type DeliveryStatus int // 寄送狀態
const (
	DELIVERY_STATUS_CARD_NOT_ORDERED   DeliveryStatus = 1
	DELIVERY_STATUS_CARD_IN_PRODUCTION DeliveryStatus = 2
	DELIVERY_STATUS_CARD_SENT_TO_USER  DeliveryStatus = 3
	DELIVERY_STATUS_CARD_ACTIVATED     DeliveryStatus = 4
	DELIVERY_STATUS_NOT_PHYSICAL_CARD  DeliveryStatus = 5
)

type ReapCode string // 開卡錯誤編碼
const (
	REAP_CODE_ERROR_INTERNAL_SERVER_ERROR ReapCode = "999999"
	// create card
	REAP_CODE_INVALID_CARD_EXPIRY_DATE                    ReapCode = "0302008"
	REAP_CODE_INVALID_CARD_EXPIRY_DATE2                   ReapCode = "0302005"
	REAP_CODE_INVALID_PHONE_NUMBER                        ReapCode = "0302009"
	REAP_CODE_SPEND_LIMIT_IS_REQUIRED                     ReapCode = "0302012"
	REAP_CODE_CRYPTO_CURRENCY_IS_REQUIRED                 ReapCode = "0321002"
	REAP_CODE_INSUFFICIENT_BALANCE                        ReapCode = "0302007"
	REAP_CODE_COUNTRY_NOT_ISSUABLE                        ReapCode = "0315001"
	REAP_CODE_SECONDARY_CARD_NAME_NOT_SUPPORTED_BY_TYPE   ReapCode = "0302019"
	REAP_CODE_SECONDARY_CARD_NAME_NOT_SUPPORTED_BY_DESIGN ReapCode = "0309007"
	REAP_CODE_INSUFFICIENT_CREDIT                         ReapCode = "0401001"
	REAP_CODE_WITHDRAWAL_WALLET_NOT_ENABLED               ReapCode = "0321001"
	REAP_CODE_ROOT_BUDGET_NOT_FOUND                       ReapCode = "0302006"
	REAP_CODE_CARD_DESIGN_NOT_FOUND                       ReapCode = "0309002"
	REAP_CODE_CANNOT_CREATE_UNDER_THIS_BUSINESS           ReapCode = "0302020"
	REAP_CODE_SYSTEM_MAINTANCE                            ReapCode = "0302010"
	REAP_CODE_INPUT_INVALID                               ReapCode = "input_invalid"

	// ship card
	REAP_CODE_CANNOT_SHIP_VIRTUAL_CARD                   ReapCode = "0302014"
	REAP_CODE_CARD_ALREADY_IN_PRODUCTION_OR_SHIPPED      ReapCode = "0302015"
	REAP_CODE_CARD_RESTRICTED                            ReapCode = "0302024"
	REAP_CODE_ZONE_REQUIRED_FOR_SHIPMENTS                ReapCode = "0315003"
	REAP_CODE_SHIPPING_UNSUPPORTED                       ReapCode = "0315002"
	REAP_CODE_PHONE_DIAL_CODE_MISMATCH                   ReapCode = "0315005"
	REAP_CODE_INVALID_SHIPPING_PHONE_NUMBER              ReapCode = "0315004"
	REAP_CODE_CARD_SHIPPING_BLOCKED_DURING_BIN_MIGRATION ReapCode = "0315006"
	REAP_CODE_CARD_DESIGN_OUT_OF_STOCK                   ReapCode = "0309001"
	REAP_CODE_ERROR_UPDATING_DELIVERY_ADDRESS            ReapCode = "0302004"
	// REAP_CODE_SHIPPING_RESTRICTED                        ReapCode = "0315001"

	// activate
	REAP_CODE_CARD_ALREADY_ACTIVATED ReapCode = "0304002"

	// unblock card
	REAP_CODE_WITHDRAWAL_PENDING                   ReapCode = "0302023"
	REAP_CODE_CARD_NOT_FROZEN_BY_SUSPECTED_FRAUD   ReapCode = "0302018"
	REAP_CODE_CANNOT_DELETE_COMPANY_MAIN_CARD      ReapCode = "0801003"
	REAP_CODE_STATUS_AND_REASON_REQUIRED           ReapCode = "0801005"
	REAP_CODE_CARD_NOT_FOUND_OR_DELETED_0302026    ReapCode = "0302026"
	REAP_CODE_CARD_STATUS_DESTROYED_NOT_REVERSABLE ReapCode = "0302010"
	REAP_CODE_CARD_NOT_FOUND_0801001               ReapCode = "0801001"
	REAP_CODE_CARD_EXPIRED_CANNOT_UNBLOCK          ReapCode = "0302031"

	// block card
	REAP_CODE_CARD_STATUS_DESTROYED_NOT_REVERSIBLE ReapCode = "0302010"
	REAP_CODE_CARD_ALREADY_BLOCKED                 ReapCode = "0302030"
	REAP_CODE_CARD_EXPIRED_CANNOT_BLOCK            ReapCode = "0302031"
	REAP_CODE_CARD_NOT_FOUND_OR_DELETED            ReapCode = "0302026"
)

type Currency uint64 // 貨幣
const (
	// crypto
	CURRENCY_USDT  Currency = 101
	CURRENCY_ETH   Currency = 102
	CURRENCY_BTC   Currency = 103
	CURRENCY_USDC  Currency = 104
	CURRENCY_DAI   Currency = 105
	CURRENCY_WBTC  Currency = 106
	CURRENCY_TRX   Currency = 107
	CURRENCY_ADA   Currency = 108
	CURRENCY_BCH   Currency = 109
	CURRENCY_DOGE  Currency = 110
	CURRENCY_LTC   Currency = 111
	CURRENCY_XRP   Currency = 112
	CURRENCY_SOL   Currency = 113
	CURRENCY_BNB   Currency = 114
	CURRENCY_ETC   Currency = 115
	CURRENCY_MATIC Currency = 116
	CURRENCY_TON   Currency = 117
	CURRENCY_AVAX  Currency = 118

	// fiat
	CURRENCY_USD        Currency = 201
	CURRENCY_EUR        Currency = 202
	CURRENCY_JPY        Currency = 203
	CURRENCY_GBP        Currency = 204
	CURRENCY_CNY        Currency = 205
	CURRENCY_CAD        Currency = 206
	CURRENCY_AUD        Currency = 207
	CURRENCY_CHF        Currency = 208
	CURRENCY_HKD        Currency = 209
	CURRENCY_SEK        Currency = 210
	CURRENCY_NZD        Currency = 211
	CURRENCY_MXN        Currency = 212
	CURRENCY_SGD        Currency = 213
	CURRENCY_NOK        Currency = 214
	CURRENCY_KRW        Currency = 215
	CURRENCY_TRY        Currency = 216
	CURRENCY_INR        Currency = 217
	CURRENCY_BRL        Currency = 218
	CURRENCY_ZAR        Currency = 219
	CURRENCY_RUB        Currency = 220
	CURRENCY_TWD        Currency = 221
	CURRENCY_DKK        Currency = 222
	CURRENCY_PLN        Currency = 223
	CURRENCY_ILS        Currency = 224
	CURRENCY_AED        Currency = 225
	CURRENCY_ARS        Currency = 226
	CURRENCY_MYR        Currency = 227
	CURRENCY_THB        Currency = 228
	CURRENCY_SAR        Currency = 229
	CURRENCY_CZK        Currency = 230
	CURRENCY_UAH        Currency = 231
	CURRENCY_PHP        Currency = 232
	CURRENCY_MOP        Currency = 233
	CURRENCY_EGP        Currency = 234
	CURRENCY_MAD        Currency = 235
	CURRENCY_KES        Currency = 236
	CURRENCY_RON        Currency = 237
	CURRENCY_BGN        Currency = 238
	CURRENCY_VND        Currency = 239
	CURRENCY_COP        Currency = 240
	CURRENCY_OTHER_FIAT Currency = 299

	//point
	CURRENCY_EPOINT Currency = 1001
)

type CurrencyConfigFeeType int // 幣種設定手續費類型
const (
	CURRENCY_CONFIG_FEE_TYPE_AMOUNT     CurrencyConfigFeeType = 1
	CURRENCY_CONFIG_FEE_TYPE_PERCENTAGE CurrencyConfigFeeType = 2
)

type CurrencyConfigStatus int // 幣種設定狀態
const (
	CURRENCY_CONFIG_STATUS_ACTIVE   CurrencyConfigStatus = 1
	CURRENCY_CONFIG_STATUS_INACTIVE CurrencyConfigStatus = 2
	CURRENCY_CONFIG_STATUS_UNSET    CurrencyConfigStatus = 3
)

type CurrencyType int // 貨幣類型
const (
	CURRENCY_TYPE_CRYPTO CurrencyType = 1
	CURRENCY_TYPE_FIAT   CurrencyType = 2
	CURRENCY_TYPE_POINT  CurrencyType = 4
)

type DatabaseRole string // 資料庫角色
const (
	DATABASE_ROLE_MAIN     DatabaseRole = "main"
	DATABASE_ROLE_READ     DatabaseRole = "read"
	DATABASE_ROLE_SNAPSHOT DatabaseRole = "snapshot"
)

type DeclinedCode string // 信用卡拒絕碼 https://www.tidalcommerce.com/learn/card-decline-codes https://support.reap.global/en/articles/701120
const (
	DECLINED_CODE_OK                               DeclinedCode = "00"  // OK
	DECLINED_CODE_MERCHANT_NOT_ACCEPT_CARD         DeclinedCode = "03"  // The merchant is not set up to accept this type of card. Please try using a different payment method.
	DECLINED_CODE_DECLINED_BY_CLIENT               DeclinedCode = "05"  // Transaction declined by client or The system has timed out.
	DECLINED_CODE_INVALID_CARD_NUMBER              DeclinedCode = "14"  // The card number is invalid, and your terminal is having trouble finding the relevant account.
	DECLINED_CODE_CARD_STATUS_LOST                 DeclinedCode = "41"  // The card's status is lost.
	DECLINED_CODE_CARD_DESTROYED                   DeclinedCode = "46"  // The card has been destroyed.
	DECLINED_CODE_INSUFFICIENT_FUNDS               DeclinedCode = "51"  // Insufficient funds.
	DECLINED_CODE_CARD_EXPIRED                     DeclinedCode = "54"  // The card has expired. Please use a valid card.
	DECLINED_CODE_INCORRECT_PIN                    DeclinedCode = "55"  // The entered PIN is incorrect. Please try again.
	DECLINED_CODE_INSUFFICIENT_FUNDS_REAL_TIME     DeclinedCode = "5E"  // Insufficient funds.
	DECLINED_CODE_TRANSACTION_NOT_PERMITTED        DeclinedCode = "57"  // Transaction is not permitted to the cardholder
	DECLINED_CODE_POS_TRANSACTION_NOT_ALLOWED      DeclinedCode = "58"  // The transaction is not allowed at this point of sale.
	DECLINED_CODE_ATM_LIMIT_EXCEEDED               DeclinedCode = "59"  // Something went wrong when processing the transaction. Please contact Reap for support.
	DECLINED_CODE_WITHDRAWAL_AMOUNT_LIMIT_EXCEEDED DeclinedCode = "61"  // The ATM withdrawal amount limit has been exceeded.
	DECLINED_CODE_WITHDRAWAL_NOT_PERMITTED         DeclinedCode = "62"  // ATM withdrawal is not permitted in the merchant's country.
	DECLINED_CODE_WITHDRAWAL_FREQUENCY_EXCEEDED    DeclinedCode = "65"  // The ATM withdrawal frequency has been exceeded.
	DECLINED_CODE_INVALID_CARD_STATUS              DeclinedCode = "70"  // Invalid card status. Please contact the issuer.
	DECLINED_CODE_EXCEEDING_PIN_ATTEMPTS           DeclinedCode = "75"  // The card has been blocked due to exceeding the maximum number of incorrect ATM PIN attempts.
	DECLINED_CODE_ISSUER_NOT_PARTICIPATE           DeclinedCode = "77"  // The issuer does not participate in that merchant category.
	DECLINED_CODE_TRANSACTION_TIMEOUT              DeclinedCode = "82"  // Transaction took too long to process, please try again.
	DECLINED_CODE_PIN_VERIFICATION_FAILED          DeclinedCode = "86"  // The system was unable to verify your PIN due to a technical issue or incorrect input.
	DECLINED_CODE_ISSUER_INOPERATIVE               DeclinedCode = "91"  // Issuer or switch is inoperative.
	DECLINED_CODE_TRANSACTION_RECORD_DISCREPANCY   DeclinedCode = "N9"  // There's a discrepancy between the transaction records of the payment systems involved.
	DECLINED_CODE_UNEXPECTED_ERROR                 DeclinedCode = "N0"  // An unexpected error happened before the payment could be fully processed.
	DECLINED_CODE_PROCESSING_ERROR                 DeclinedCode = "N3"  // Something went wrong when processing the transaction. Please contact Reap for support.
	DECLINED_CODE_INVALID_CVV                      DeclinedCode = "N7"  // The CVV entered is incorrect. Please check and try again.
	DECLINED_CODE_VERIFICATION_FAILED              DeclinedCode = "6P"  // Verification of the provided data failed. Please verify and try again.
	DECLINED_CODE_CARD_FROZEN                      DeclinedCode = "R1"  // Cardholder has frozen the card.
	DECLINED_CODE_EXCEEDING_CVV_ATTEMPTS           DeclinedCode = "R2"  // The card has been blocked due to exceeding the maximum number of incorrect CVV transaction attempts.
	DECLINED_CODE_EXCEEDING_EXPIRY_DATE_ATTEMPTS   DeclinedCode = "R3"  // The card has been blocked due to exceeding the maximum number of incorrect expiry date transaction attempts.
	DECLINED_CODE_CARD_NOT_ACTIVATED               DeclinedCode = "R4"  // The card is not activated yet. Please activate it from the Reap mobile app or the Reap Dashboard.
	DECLINED_CODE_WITHDRAWAL_NOT_ALLOWED           DeclinedCode = "R5"  // Withdrawals are not allowed for this card.
	DECLINED_CODE_DAILY_SPENDING_LIMIT_EXCEEDED    DeclinedCode = "R6"  // The daily ceiling allowed to spend for this card has been exceeded. Please contact your Admin.
	DECLINED_CODE_SPENDING_FREQUENCY_EXCEEDED      DeclinedCode = "R7"  // The spending frequency allowed for this card has been exceeded. Please contact your Admin.
	DECLINED_CODE_CARD_BLOCKED                     DeclinedCode = "R8"  // The card is blocked.
	DECLINED_CODE_CATEGORY_INSUFFICIENT_FUNDS      DeclinedCode = "R9"  // Not enough balance to complete the transaction for this merchant category. Please add funds.
	DECLINED_CODE_BUDGET_INSUFFICIENT_FUNDS        DeclinedCode = "R10" // Insufficient funds in the budget associated with this card. Please contact your Admin.
	DECLINED_CODE_BUDGET_FROZEN                    DeclinedCode = "R11" // The budget associated with this card has been frozen. Please contact Reap.
)

type DepositOrderStatus int // 入帳單狀態
const (
	DEPOSIT_ORDER_STATUS_PENDING   DepositOrderStatus = 1
	DEPOSIT_ORDER_STATUS_CONFIRMED DepositOrderStatus = 2
	DEPOSIT_ORDER_STATUS_DEPOSITED DepositOrderStatus = 3
	DEPOSIT_ORDER_STATUS_FAILED    DepositOrderStatus = 4
)

type DiscountType int // 優惠類型
const (
	DISCOUNT_TYPE_PERCENTAGE DiscountType = 1
	DISCOUNT_TYPE_AMOUNT     DiscountType = 2
	DISCOUNT_TYPE_BONUS      DiscountType = 3
)

type DiscountGroup int // 優惠群組
const (
	DISCOUNT_GROUP_GENERAL DiscountGroup = 1
	DISCOUNT_GROUP_MONETA  DiscountGroup = 2
)

type EncryptType int // 加密類型
const (
	ENCRYPT_TYPE_BASE64 EncryptType = 1
	ENCRYPT_TYPE_AES    EncryptType = 2
	ENCRYPT_TYPE_RSA    EncryptType = 3
)

type ExchangeTriggerMode int // 兑幣觸發模式
const (
	EXCHANGE_TRIGGER_MODE_ON_REQUEST          ExchangeTriggerMode = 1
	EXCHANGE_TRIGGER_MODE_AUTO                ExchangeTriggerMode = 2
	EXCHANGE_TRIGGER_MODE_MERCHANT_ON_REQUEST ExchangeTriggerMode = 3
	EXCHANGE_TRIGGER_MODE_MERCHANT_AUTO       ExchangeTriggerMode = 4
	EXCHANGE_TRIGGER_MODE_MERCHANT_ON_DEPOSIT ExchangeTriggerMode = 5
)

type ExternalRequestLogStage int // 外部請求日誌階段
const (
	EXTERNAL_REQUEST_LOG_STAGE_PRELOG   ExternalRequestLogStage = 1
	EXTERNAL_REQUEST_LOG_STAGE_UPDATE_1 ExternalRequestLogStage = 2
	EXTERNAL_REQUEST_LOG_STAGE_UPDATE_2 ExternalRequestLogStage = 3
	EXTERNAL_REQUEST_LOG_STAGE_UPDATE_3 ExternalRequestLogStage = 4
	EXTERNAL_REQUEST_LOG_STAGE_POSTLOG  ExternalRequestLogStage = 99
)

type CategoryUsage int // 一般用途
const (
	CATEGORY_USAGE_USER_APPLY            CategoryUsage = 0b00000000000000000000000000000001 // 1
	CATEGORY_USAGE_USER_DISPLAY          CategoryUsage = 0b00000000000000000000000000000010 // 2
	CATEGORY_USAGE_MERCHANT_APPLY        CategoryUsage = 0b00000000000000000000000000000100 // 4
	CATEGORY_USAGE_MERCHANT_DISPLAY      CategoryUsage = 0b00000000000000000000000000001000 // 8
	CATEGORY_USAGE_MERCHANT_USER_APPLY   CategoryUsage = 0b00000000000000000000000000010000 // 16
	CATEGORY_USAGE_MERCHANT_USER_DISPLAY CategoryUsage = 0b00000000000000000000000000100000 // 32
	CATEGORY_USAGE_ADMIN_DISPLAY         CategoryUsage = 0b00000000000000000000000001000000 // 64
	CATEGORY_USAGE_ADDRESS_POOL          CategoryUsage = 0b00000000000000000000000010000000 // 128
	CATEGORY_USAGE_QUOTE                 CategoryUsage = 0b00000000000000000000000100000000 // 256
	CATEGORY_USAGE_USER_REWARD           CategoryUsage = 0b00000000000000000000001000000000 // 512
	CATEGORY_USAGE_USER_AUTO_YIELD       CategoryUsage = 0b00000000000000000000010000000000 // 1024
	CATEGORY_USAGE_FIAT_SUPPORT          CategoryUsage = 0b00000000000000000000100000000000 // 2048
	CATEGORY_USAGE_USER_DELETE           CategoryUsage = 0b00000000000000000001000000000000 // 4096
	CATEGORY_USAGE_USER_APPLE_PAY_APPLY  CategoryUsage = 0b00000000000000000010000000000000 // 8192
	CATEGORY_USAGE_ADMIN_DELETE          CategoryUsage = 0b00000000000000000100000000000000 // 16384
	CATEGORY_USAGE_ADMIN_APPLY           CategoryUsage = 0b00000000000000001000000000000000 // 32768
	CATEGORY_USAGE_POINT                 CategoryUsage = 0b00000000000000010000000000000000 // 65536
	CATEGORY_USAGE_ALL                   CategoryUsage = 0b00111111111111111111111111111111 // 1073741823
	CATEGORY_USAGE_NONE                  CategoryUsage = 0b01000000000000000000000000000000 // 1073741824
)

type CategoryFundInUsage int // 進帳入金用途
const (
	CATEGORY_FUND_IN_USAGE_EXCHANGE_IN           CategoryFundInUsage = 0b00000000000000000000000000000001 // 1
	CATEGORY_FUND_IN_USAGE_TRANSFER_IN           CategoryFundInUsage = 0b00000000000000000000000000000010 // 2
	CATEGORY_FUND_IN_USAGE_TOP_UP_IN             CategoryFundInUsage = 0b00000000000000000000000000000100 // 4
	CATEGORY_FUND_IN_USAGE_TOP_DOWN_IN           CategoryFundInUsage = 0b00000000000000000000000000001000 // 8
	CATEGORY_FUND_IN_USAGE_CARD_TO_CARD_IN       CategoryFundInUsage = 0b00000000000000000000000000010000 // 16
	CATEGORY_FUND_IN_USAGE_SELF_CARD_TO_CARD_IN  CategoryFundInUsage = 0b00000000000000000000000000100000 // 32
	CATEGORY_FUND_IN_USAGE_CROSS_CARD_TO_CARD_IN CategoryFundInUsage = 0b00000000000000000000000001000000 // 64
	CATEGORY_FUND_IN_USAGE_APPLE                 CategoryFundInUsage = 0b00000000000000000000000010000000 // 128
	CATEGORY_FUND_IN_USAGE_CHANGELLY             CategoryFundInUsage = 0b00000000000000000000000100000000 // 256
	CATEGORY_FUND_IN_USAGE_CHIPPAY               CategoryFundInUsage = 0b00000000000000000000001000000000 // 512
	CATEGORY_FUND_IN_USAGE_ALL                   CategoryFundInUsage = 0b00111111111111111111111111111111 // 1073741823
	CATEGORY_FUND_IN_USAGE_NONE                  CategoryFundInUsage = 0b01000000000000000000000000000000 // 1073741824
)

type CategoryFundOutUsage int // 出帳出金用途
const (
	CATEGORY_FUND_OUT_USAGE_EXCHANGE_OUT           CategoryFundOutUsage = 0b00000000000000000000000000000001 // 1
	CATEGORY_FUND_OUT_USAGE_TRANSFER_OUT           CategoryFundOutUsage = 0b00000000000000000000000000000010 // 2
	CATEGORY_FUND_OUT_USAGE_TOP_UP_OUT             CategoryFundOutUsage = 0b00000000000000000000000000000100 // 4
	CATEGORY_FUND_OUT_USAGE_TOP_DOWN_OUT           CategoryFundOutUsage = 0b00000000000000000000000000001000 // 8
	CATEGORY_FUND_OUT_USAGE_CARD_TO_CARD_OUT       CategoryFundOutUsage = 0b00000000000000000000000000010000 // 16
	CATEGORY_FUND_OUT_USAGE_SELF_CARD_TO_CARD_OUT  CategoryFundOutUsage = 0b00000000000000000000000000100000 // 32
	CATEGORY_FUND_OUT_USAGE_CROSS_CARD_TO_CARD_OUT CategoryFundOutUsage = 0b00000000000000000000000001000000 // 64
	CATEGORY_FUND_OUT_USAGE_WITHDRAW               CategoryFundOutUsage = 0b00000000000000000000000010000000 // 128
	CATEGORY_FUND_OUT_USAGE_APPLY                  CategoryFundOutUsage = 0b00000000000000000000000100000000 // 256
	CATEGORY_FUND_OUT_USAGE_ALL                    CategoryFundOutUsage = 0b00111111111111111111111111111111 // 1073741823
	CATEGORY_FUND_OUT_USAGE_NONE                   CategoryFundOutUsage = 0b01000000000000000000000000000000 // 1073741824
)

type CategoryFrontendUsage int // 前端顯示
const (
	CATEGORY_FRONT_USAGE_ENABLE_APPLE_PAY       CategoryFrontendUsage = 0b00000000000000000000000000000001 // 1
	CATEGORY_FRONT_USAGE_ENABLE_DELETE          CategoryFrontendUsage = 0b00000000000000000000000000000010 // 2
	CATEGORY_FRONT_USAGE_ENABLE_SPENDING_LIMIT  CategoryFrontendUsage = 0b00000000000000000000000000000100 // 4
	CATEGORY_FRONT_USAGE_ENABLE_FINANCE         CategoryFrontendUsage = 0b00000000000000000000000000001000 // 8
	CATEGORY_FRONT_USAGE_ENABLE_QR_CODE         CategoryFrontendUsage = 0b00000000000000000000000000010000 // 16
	CATEGORY_FRONT_USAGE_ENABLE_CONVERT         CategoryFrontendUsage = 0b00000000000000000000000000100000 // 32
	CATEGORY_FRONT_USAGE_ENABLE_CHANGELLY       CategoryFrontendUsage = 0b00000000000000000000000001000000 // 64
	CATEGORY_FRONT_USAGE_ENABLE_CHIPPAY         CategoryFrontendUsage = 0b00000000000000000000000010000000 // 128
	CATEGORY_FRONT_USAGE_ENABLE_APPLE_DEPOSIT   CategoryFrontendUsage = 0b00000000000000000000000100000000 // 256
	CATEGORY_FRONT_USAGE_ENABLE_SEND            CategoryFrontendUsage = 0b00000000000000000000001000000000 // 512
	CATEGORY_FRONT_USAGE_ENABLE_TO_CONVERT      CategoryFrontendUsage = 0b00000000000000000000010000000000 // 1024
	CATEGORY_FRONT_USAGE_ENABLE_BINANCE_PAY     CategoryFrontendUsage = 0b00000000000000000000100000000000 // 2048
	CATEGORY_FRONT_USAGE_ENABLE_GOOGLE_PAY      CategoryFrontendUsage = 0b00000000000000000001000000000000 // 4096
	CATEGORY_FRONT_USAGE_ENABLE_PAYPAL          CategoryFrontendUsage = 0b00000000000000000010000000000000 // 8192
	CATEGORY_FRONT_USAGE_ENABLE_DEPOSIT         CategoryFrontendUsage = 0b00000000000000000100000000000000 // 16384
	CATEGORY_FRONT_USAGE_ENABLE_ACTIVATE        CategoryFrontendUsage = 0b00000000000000001000000000000000 // 32768
	CATEGORY_FRONT_USAGE_ENABLE_FROM_CONVERSION CategoryFrontendUsage = 0b00000000000000010000000000000000 // 65536 当为 true 时，表示该卡可以作为"资金来源"，执行资金转换（转出操作）。
	CATEGORY_FRONT_USAGE_ENABLE_TO_CONVERSION   CategoryFrontendUsage = 0b00000000000000100000000000000000 // 131072 当为 true 时，表示该卡可以作为"资金目标"，接收转换资金（转入操作）。
	CATEGORY_FRONT_USAGE_ENABLE_APPLY_SELECT    CategoryFrontendUsage = 0b00000000000001000000000000000000 // 262144
	CATEGORY_FRONT_USAGE_ENABLE_ALI_PAY         CategoryFrontendUsage = 0b00000000000010000000000000000000 // 524288
	CATEGORy_FRONT_USAGE_ENABLE_WECHAT_PAY      CategoryFrontendUsage = 0b00000000000100000000000000000000 // 1048576
	CATEGORY_FRONT_USAGE_ALL                    CategoryFrontendUsage = 0b00111111111111111111111111111111 // 1073741823
	CATEGORY_FRONT_USAGE_NONE                   CategoryFrontendUsage = 0b01000000000000000000000000000000 // 1073741824
)

type CountryGroupType int // 支持國家類型
const (
	COUNTRY_GROUP_TYPE_SUPPORTED   CountryGroupType = 1 // 支持的國家
	COUNTRY_GROUP_TYPE_UNSUPPORTED CountryGroupType = 2 // 不支持的國家
)

type GroupName string // 群組名稱
const (
	GROUP_NAME_REGULAR_USER           GroupName = "regular_user"
	GROUP_NAME_REGULAR_ADMIN          GroupName = "regular_admin"
	GROUP_NAME_REGULAR_MERCHANT       GroupName = "regular_merchant"
	GROUP_NAME_REGULAR_MERCHANT_ADMIN GroupName = "regular_merchant_admin"
	GROUP_NAME_REGULAR_MERCHANT_USER  GroupName = "regular_merchant_user"
)

type EPointLevel int // e積分等級
const (
	EPOINT_LEVEL_BRONZE EPointLevel = 1
	EPOINT_LEVEL_SILVER EPointLevel = 2
	EPOINT_LEVEL_GOLD   EPointLevel = 3
)

type ExchangeStatus int // 兌換狀態
const (
	EXCHANGE_STATUS_PENDING ExchangeStatus = 1
	EXCHANGE_STATUS_SUCCESS ExchangeStatus = 2
	EXCHANGE_STATUS_FAILED  ExchangeStatus = 3
)

type ExportStatus string // 导出状态
const (
	EXPORT_STATUS_PENDING    ExportStatus = "pending"
	EXPORT_STATUS_PROCESSING ExportStatus = "processing"
	EXPORT_STATUS_COMPLETED  ExportStatus = "completed"
	EXPORT_STATUS_FAILED     ExportStatus = "failed"
)

type ExternalProvider int // 外部提供商
const (
	EXTERNAL_PROVIDER_COINSDO    ExternalProvider = 1
	EXTERNAL_PROVIDER_COINFACE   ExternalProvider = 2
	EXTERNAL_PROVIDER_REAP       ExternalProvider = 3
	EXTERNAL_PROVIDER_CHIPPAY    ExternalProvider = 4
	EXTERNAL_PROVIDER_CHANGELLY  ExternalProvider = 5
	EXTERNAL_PROVIDER_WHALE      ExternalProvider = 6
	EXTERNAL_PROVIDER_PAYCRYPTO  ExternalProvider = 7
	EXTERNAL_PROVIDER_BINANCEPAY ExternalProvider = 8
	EXTERNAL_PROVIDER_MAILGUN    ExternalProvider = 9
)

type FeeType string // 手續費種類
const (
	FEE_TYPE_EXCHANGE      FeeType = "exchange"
	FEE_TYPE_DEPOSIT       FeeType = "deposit"
	FEE_TYPE_TRANSFER      FeeType = "transfer"
	FEE_TYPE_WITHDRAW      FeeType = "withdraw"
	FEE_TYPE_CARD          FeeType = "card"
	FEE_TYPE_PHYSICAL_CARD FeeType = "physical_card"
	FEE_TYPE_DELIVERY      FeeType = "delivery"
	FEE_TYPE_TOP_UP        FeeType = "top_up"
	FEE_TYPE_TOP_DOWN      FeeType = "top_down"
	FEE_TYPE_CARD_TO_CARD  FeeType = "card_to_card"
	FEE_TYPE_ATM           FeeType = "atm"
	FEE_TYPE_FX            FeeType = "fx"
)

type FileType int // 檔案類型
const (
	FILE_TYPE_TEXT  FileType = 1
	FILE_TYPE_IMAGE FileType = 2
	FILE_TYPE_AUDIO FileType = 3
	FILE_TYPE_VIDEO FileType = 4
	FILE_TYPE_FILE  FileType = 5
	FILE_TYPE_ZIP   FileType = 6
	FILE_TYPE_EXE   FileType = 7
	FILE_TYPE_FONT  FileType = 8
)

type FilePurpose int // 檔案用途
const (
	FILE_PURPOSE_KYC2_FACE_PHOTO1    FilePurpose = 1
	FILE_PURPOSE_KYC3_FACE_PHOTO1    FilePurpose = 2
	FILE_PURPOSE_KYC3_FACE_PHOTO2    FilePurpose = 3
	FILE_PURPOSE_KYC3_DOCUMENT_PHOTO FilePurpose = 4
	FILE_PURPOSE_AGENT_REPORT        FilePurpose = 5
	FILE_PURPOSE_MERCHANT_BILL       FilePurpose = 6
)

type FileUploadStatus int // 檔案上傳狀態
const (
	FILE_UPLOAD_STATUS_UPLOADING FileUploadStatus = 1
	FILE_UPLOAD_STATUS_UPLOADED  FileUploadStatus = 2
	FILE_UPLOAD_STATUS_FAILED    FileUploadStatus = 3
)

type FileStatus int // 檔案狀態
const (
	FILE_STATUS_ENABLED  FileStatus = 1
	FILE_STATUS_DISABLED FileStatus = 2
)

type FinancialCode string // 理財代碼
const (
	FINANCIAL_CODE_AUTO_YIELD FinancialCode = "auto_yield"
)

type FinancialProductSupportType int // 理財產品支持類型
const (
	FINANCIAL_PRODUCT_SUPPORT_TYPE_ALL      FinancialProductSupportType = 1
	FINANCIAL_PRODUCT_SUPPORT_TYPE_CURRENCY FinancialProductSupportType = 2
)

type FinancialProductStatus int // 理財產品狀態
const (
	FINANCIAL_PRODUCT_STATUS_ACTIVE   FinancialProductStatus = 1
	FINANCIAL_PRODUCT_STATUS_INACTIVE FinancialProductStatus = 2
)

type FinancialTransferDirection int // 理財轉帳方向
const (
	FINANCIAL_TRANSFER_DIRECTION_IN  FinancialTransferDirection = 1 // 存入理財帳戶
	FINANCIAL_TRANSFER_DIRECTION_OUT FinancialTransferDirection = 2 // 理財帳戶領出
)

type FinancialTransferStatus int // 理財轉帳狀態
const (
	FINANCIAL_TRANSFER_STATUS_PENDING    FinancialTransferStatus = 1
	FINANCIAL_TRANSFER_STATUS_PROCESSING FinancialTransferStatus = 2
	FINANCIAL_TRANSFER_STATUS_SUCCESS    FinancialTransferStatus = 3
	FINANCIAL_TRANSFER_STATUS_FAILED     FinancialTransferStatus = 4
)

type FinancialTransferTriggerMethod int // 理財轉帳方式
const (
	FINANCIAL_TRANSFER_TRIGGER_METHOD_MANUAL  FinancialTransferTriggerMethod = 1
	FINANCIAL_TRANSFER_TRIGGER_METHOD_INSTANT FinancialTransferTriggerMethod = 2
	FINANCIAL_TRANSFER_TRIGGER_METHOD_AUTO    FinancialTransferTriggerMethod = 3
)

type FinancialReportTable string // 理財報表
const (
	FINANCIAL_REPORT_TABLE_APPLY_ORDER                   FinancialReportTable = "apply_order"
	FINANCIAL_REPORT_TABLE_CARD_TO_CARD_ORDER            FinancialReportTable = "card_to_card_order"
	FINANCIAL_REPORT_TABLE_CHIPPAY_EXPRESS_ORDER         FinancialReportTable = "chippay_express_order"
	FINANCIAL_REPORT_TABLE_DEPOSIT_ORDER                 FinancialReportTable = "deposit_order"
	FINANCIAL_REPORT_TABLE_EXCHANGE_ORDER                FinancialReportTable = "exchange_order"
	FINANCIAL_REPORT_TABLE_FINANCIAL_TRANSFER_ORDER      FinancialReportTable = "financial_transfer_order"
	FINANCIAL_REPORT_TABLE_INTEREST_ORDER                FinancialReportTable = "interest_order"
	FINANCIAL_REPORT_TABLE_MERCHANT_ADJUST_BALANCE_ORDER FinancialReportTable = "merchant_adjust_balance_order"
	FINANCIAL_REPORT_TABLE_POINT_ACCURAL_ORDER           FinancialReportTable = "point_accural_order"
	FINANCIAL_REPORT_TABLE_SYSTEM_CHARGE_ORDER           FinancialReportTable = "system_charge_order"
	FINANCIAL_REPORT_TABLE_TERMINATE_ORDER               FinancialReportTable = "terminate_order"
	FINANCIAL_REPORT_TABLE_TOP_DOWN_ORDER                FinancialReportTable = "top_down_order"
	FINANCIAL_REPORT_TABLE_TOP_UP_ORDER                  FinancialReportTable = "top_up_order"
	FINANCIAL_REPORT_TABLE_TRANSFER_ORDER                FinancialReportTable = "transfer_order"
	FINANCIAL_REPORT_TABLE_WITHDRAW_ORDER                FinancialReportTable = "withdraw_order"
	FINANCIAL_REPORT_TABLE_REAP_TRANSACTION              FinancialReportTable = "reap_transaction"
	FINANCIAL_REPORT_TABLE_WHALE_TRANSACTION             FinancialReportTable = "whale_transaction"
	FINANCIAL_REPORT_TABLE_PAYCRYPTO_TRANSACTION         FinancialReportTable = "paycrypto_transaction"
	FINANCIAL_REPORT_TABLE_ETHERFI_TRANSACTION           FinancialReportTable = "etherfi_transaction"
	FINANCIAL_REPORT_TABLE_BILL                          FinancialReportTable = "bill"
)

type ForwardType int // 轉發類型
const (
	FORWARD_TYPE_REAP     ForwardType = 0b1     // 1
	FORWARD_TYPE_EMAIL    ForwardType = 0b10    // 2
	FORWARD_TYPE_SMS      ForwardType = 0b100   // 4
	FORWARD_TYPE_APP      ForwardType = 0b1000  // 8
	FORWARD_TYPE_MERCHANT ForwardType = 0b10000 // 16
)

type Gender int // 性別
const (
	GENDER_MALE   Gender = 1
	GENDER_FEMALE Gender = 2
)

type HTTPMethod int // http方法
const (
	HTTP_METHOD_POST   = 1
	HTTP_METHOD_GEt    = 2
	HTTP_METHOD_PUT    = 3
	HTTP_METHOD_DELETE = 4
)

type IdentityDocumentType int // reap證件類型
const (
	IDENTITY_DOCUMENT_TYPE_PASSPORT       IdentityDocumentType = 1
	IDENTITY_DOCUMENT_TYPE_HEALTH         IdentityDocumentType = 2
	IDENTITY_DOCUMENT_TYPE_NATIONAL_ID    IdentityDocumentType = 3
	IDENTITY_DOCUMENT_TYPE_TAX_ID_NUMBER  IdentityDocumentType = 4
	IDENTITY_DOCUMENT_TYPE_SOCIAL_SERVICE IdentityDocumentType = 5
	IDENTITY_DOCUMENT_TYPE_DRIVER         IdentityDocumentType = 6
)

type IdentityStatus int // kyc
const (
	IDENTITY_STATUS_PENDING IdentityStatus = 1
	IDENTITY_STATUS_SUCCESS IdentityStatus = 2
	IDENTITY_STATUS_FAILED  IdentityStatus = 3
)

type InterestOrderStatus int // 利息訂單狀態
const (
	INTEREST_ORDER_STATUS_PENDING InterestOrderStatus = 1
	INTEREST_ORDER_STATUS_SUCCESS InterestOrderStatus = 2
	INTEREST_ORDER_STATUS_FAILED  InterestOrderStatus = 3
)

type KYCLevel string // kyc level 因為在golang中0常被當成不存在，因此enum改用string表達
const (
	// 无KYC
	KYC_LEVEL_0 KYCLevel = "0"
	// 上传证照信息 II期
	KYC_LEVEL_1 KYCLevel = "1"
	// 活体识别/人证合一通过
	KYC_LEVEL_2 KYCLevel = "2"
	// 上传地址/收入证明
	KYC_LEVEL_3 KYCLevel = "3"
)

type KYCStatus int // 检测状态:1.认证中 2.通过 3.拒绝 4.已創建 5.失敗
const (
	KYC_STATUS_PENDING KYCStatus = 1
	KYC_STATUS_PASS    KYCStatus = 2
	KYC_STATUS_REJECT  KYCStatus = 3
	KYC_STATUS_CREATED KYCStatus = 4
	KYC_STATUS_FAILED  KYCStatus = 5
)

type KYCRequestStatus string // 申請狀態
const (
	KYC_REQUEST_STATUS_NONE            KYCRequestStatus = "none"
	KYC_REQUEST_STATUS_LEVEL_2_PENDING KYCRequestStatus = "level_2_pending"
	KYC_REQUEST_STATUS_LEVEL_3_PENDING KYCRequestStatus = "level_3_pending"
)

type L2CacheLevel int // 二級緩存等級
const (
	L2_CACHE_LEVEL_MEMORY L2CacheLevel = 1
	L2_CACHE_LEVEL_REDIS  L2CacheLevel = 2
)

type Language string // 語言
const (
	// 中文
	LANGUAGE_CHINESE                Language = "zh"
	LANGUAGE_SIMPLIFIED_CHINESE     Language = "zh_hans"
	LANGUAGE_TRADITIONAL_CHINESE_TW Language = "zh_hant"
	LANGUAGE_TRADITIONAL_CHINESE_HK Language = "zh_hk"
	LANGUAGE_SIMPLIFIED_CHINESE_SG  Language = "zh_sg"
	LANGUAGE_ENGLISH                Language = "en"
	LANGUAGE_ENGLISH_US             Language = "en_us"
	LANGUAGE_ENGLISH_GB             Language = "en_gb"
	LANGUAGE_ENGLISH_AU             Language = "en_au"
	LANGUAGE_ENGLISH_CA             Language = "en_ca"
	LANGUAGE_GERMAN                 Language = "de"
	LANGUAGE_GERMAN_DE              Language = "de_de"
	LANGUAGE_GERMAN_AT              Language = "de_at"
	LANGUAGE_GERMAN_CH              Language = "de_ch"
	LANGUAGE_FRENCH                 Language = "fr"
	LANGUAGE_FRENCH_FR              Language = "fr_fr"
	LANGUAGE_FRENCH_CA              Language = "fr_ca"
	LANGUAGE_SPANISH                Language = "es"
	LANGUAGE_SPANISH_ES             Language = "es_es"
	LANGUAGE_SPANISH_MX             Language = "es_mx"
	LANGUAGE_ITALIAN                Language = "it"
	LANGUAGE_PORTUGUESE             Language = "pt"
	LANGUAGE_PORTUGUESE_BR          Language = "pt_br"
	LANGUAGE_DUTCH                  Language = "nl"
	LANGUAGE_RUSSIAN                Language = "ru"
	LANGUAGE_POLISH                 Language = "pl"
	LANGUAGE_TURKISH                Language = "tr"
	LANGUAGE_JAPANESE               Language = "ja"
	LANGUAGE_KOREAN                 Language = "ko"
	LANGUAGE_THAI                   Language = "th"
	LANGUAGE_VIETNAMESE             Language = "vi"
	LANGUAGE_HINDI                  Language = "hi"
	LANGUAGE_ARABIC                 Language = "ar"
	LANGUAGE_AFRIKAANS              Language = "af"
	LANGUAGE_BULGARIAN              Language = "bg"
	LANGUAGE_CATALAN                Language = "ca"
	LANGUAGE_CZECH                  Language = "cs"
	LANGUAGE_DANISH                 Language = "da"
	LANGUAGE_GREEK                  Language = "el"
	LANGUAGE_FINNISH                Language = "fi"
	LANGUAGE_HEBREW                 Language = "he"
	LANGUAGE_HUNGARIAN              Language = "hu"
	LANGUAGE_INDONESIAN             Language = "id"
	LANGUAGE_NORWEGIAN              Language = "no"
	LANGUAGE_ROMANIAN               Language = "ro"
	LANGUAGE_SLOVAK                 Language = "sk"
	LANGUAGE_SWEDISH                Language = "sv"
	LANGUAGE_UKRAINIAN              Language = "uk"
)

type LoginType int // 登入類型
const (
	LOGIN_TYPE_USER_NO_PIN_CODE LoginType = 1
	LOGIN_TYPE_USER             LoginType = 2
	LOGIN_TYPE_ADMIN            LoginType = 3
)

type Mainnet int // 主網路
const (
	MAINNET_BTC   Mainnet = 1
	MAINNET_ETH   Mainnet = 2
	MAINNET_TRX   Mainnet = 3
	MAINNET_ADA   Mainnet = 4
	MAINNET_BCH   Mainnet = 5
	MAINNET_DOGE  Mainnet = 6
	MAINNET_LTC   Mainnet = 7
	MAINNET_XRP   Mainnet = 8
	MAINNET_SOL   Mainnet = 9
	MAINNET_BSC   Mainnet = 10
	MAINNET_ETC   Mainnet = 11
	MAINNET_MATIC Mainnet = 12
	MAINNET_TON   Mainnet = 13
	MAINNET_AVAXC Mainnet = 14
)

type MerchantBalanceDeductToggle int // 商戶扣除餘額開關
const (
	MERCHANT_BALANCE_DEDUCT_TOGGLE_ON  MerchantBalanceDeductToggle = 1
	MERCHANT_BALANCE_DEDUCT_TOGGLE_OFF MerchantBalanceDeductToggle = 2
)

type MerchantTier int // 商戶等級
const (
	MERCHANT_TIER_NORMAL MerchantTier = 1
	MERCHANT_TIER_TRUST  MerchantTier = 2
)

type MerchantReturnCode string

const (
	MERCHANT_RETURN_CODE_OK                               MerchantReturnCode = "00" // OK
	MERCHANT_RETURN_CODE_DECLINED_BY_CLIENT               MerchantReturnCode = "05" // Transaction declined by client or The system has timed out.
	MERCHANT_RETURN_CODE_INVALID_CARD_NUMBER              MerchantReturnCode = "14" // The card number is invalid, and your terminal is having trouble finding the relevant account.
	MERCHANT_RETURN_CODE_CARD_STATUS_LOST                 MerchantReturnCode = "41" // The card's status is lost.
	MERCHANT_RETURN_CODE_CARD_DESTROYED                   MerchantReturnCode = "46" // The card has been destroyed.
	MERCHANT_RETURN_CODE_INSUFFICIENT_FUNDS               MerchantReturnCode = "51" // Insufficient funds.
	MERCHANT_RETURN_CODE_INSUFFICIENT_FUNDS_REAL_TIME     MerchantReturnCode = "5E" // Insufficient funds.
	MERCHANT_RETURN_CODE_TRANSACTION_NOT_PERMITTED        MerchantReturnCode = "57" // Transaction is not permitted to the cardholder
	MERCHANT_RETURN_CODE_WITHDRAWAL_AMOUNT_LIMIT_EXCEEDED MerchantReturnCode = "61" // The ATM withdrawal amount limit has been exceeded.
	MERCHANT_RETURN_CODE_WITHDRAWAL_NOT_PERMITTED         MerchantReturnCode = "62" // ATM withdrawal is not permitted in the merchant's country.
	MERCHANT_RETURN_CODE_WITHDRAWAL_FREQUENCY_EXCEEDED    MerchantReturnCode = "65" // The ATM withdrawal frequency has been exceeded.
	MERCHANT_RETURN_CODE_INVALID_CARD_STATUS              MerchantReturnCode = "70" // Invalid card status. Please contact the issuer.
	MERCHANT_RETURN_CODE_EXCEEDING_PIN_ATTEMPTS           MerchantReturnCode = "75" // The card has been blocked due to exceeding the maximum number of incorrect ATM PIN attempts.
	MERCHANT_RETURN_CODE_ISSUER_NOT_PARTICIPATE           MerchantReturnCode = "77" // The issuer does not participate in that merchant category.
	MERCHANT_RETURN_CODE_CARD_FROZEN                      MerchantReturnCode = "R1" // Cardholder has frozen the card.
	MERCHANT_RETURN_CODE_EXCEEDING_CVV_ATTEMPTS           MerchantReturnCode = "R2" // The card has been blocked due to exceeding the maximum number of incorrect CVV transaction attempts.
	MERCHANT_RETURN_CODE_EXCEEDING_EXPIRY_DATE_ATTEMPTS   MerchantReturnCode = "R3" // The card has been blocked due to exceeding the maximum number of incorrect expiry date transaction attempts.
)

type MerchantEventName string // 事件類型
const (
	MERCHANT_EVENT_NAME_REQUEST                 MerchantEventName = "request"                // Triggered when authorization is initialized
	MERCHANT_EVENT_NAME_AUTHORIZATION           MerchantEventName = "authorization"          // Triggered when a transaction is authorized.
	MERCHANT_EVENT_NAME_AUTHORIZATION_ADVICE    MerchantEventName = "authorization.advice"   // Triggered when a transaction has been declined
	MERCHANT_EVENT_NAME_AUTHORIZATION_REVERSAL  MerchantEventName = "authorization.reversal" // Triggered when an authorized transaction amount has been reversed.
	MERCHANT_EVENT_NAME_AUTHORIZATION_CLEARING  MerchantEventName = "authorization.clearing" // Triggered when an authorized transaction amount has been cleared.
	MERCHANT_EVENT_NAME_REFUND                  MerchantEventName = "refund"                 // Triggered when an authorized transaction amount has been refunded.
	MERCHANT_EVENT_NAME_RECEIVED                MerchantEventName = "received"
	MERCHANT_EVENT_NAME_WALLET_OPT_NOTIFICATION MerchantEventName = "wallet_tokenization_otp"
	MERCHANT_EVENT_NAME_CARD_STATUS             MerchantEventName = "card_status"
	MERCHANT_EVENT_NAME_SHIPPING_STATUS         MerchantEventName = "shipping_status"
	MERCHANT_EVENT_NAME_FREEZE_STATUS           MerchantEventName = "freeze_status"
	MERCHANT_EVENT_NAME_BALANCE_CHANGE_REPORT   MerchantEventName = "balance_change"
	MERCHANT_EVENT_NAME_KYC_RESUTL              MerchantEventName = "kyc_result"
)

type MerchantEventType string // 事件類型
const (
	MERCHANT_EVENT_TYPE_TRANSACTION   MerchantEventType = "transaction"
	MERCHANT_EVENT_TYPE_AUTHORIZATION MerchantEventType = "authorization"
	MERCHANT_EVENT_TYPE_NOTIFICATION  MerchantEventType = "notification"
	MERCHANT_EVENT_TYPE_CARD          MerchantEventType = "card"
	MERCHANT_EVENT_TYPE_ACCOUNT       MerchantEventType = "account"
	MERCHANT_EVENT_TYPE_FILES         MerchantEventType = "files"
	MERCHANT_EVENT_TYPE_SHIPPING      MerchantEventType = "shipping"
	MERCHANT_EVENT_TYPE_KYC           MerchantEventType = "kyc"
)

type MerchantStatus int // 商戶狀態
const (
	MERCHANT_STATUS_ACTIVE   MerchantStatus = 1
	MERCHANT_STATUS_DISABLED MerchantStatus = 2
	MERCHANT_STATUS_FROZEN   MerchantStatus = 3
	MERCHANT_STATUS_DELETED  MerchantStatus = 4
)

type ManualStatus int // 手動訂單狀態
const (
	MANUAL_STATUS_PENDING       ManualStatus = 1
	MANUAL_STATUS_PENDING_AUDIT ManualStatus = 2
	MANUAL_STATUS_AUDITED       ManualStatus = 3
	MANUAL_STATUS_SUCCESS       ManualStatus = 4
	MANUAL_STATUS_AUDIT_FAILED  ManualStatus = 5
	MANUAL_STATUS_FAILED        ManualStatus = 6
)

type MessagePurpose int // 訊息用途
const (
	MESSAGE_PURPOSE_USER_LOGIN_OR_REGISTER MessagePurpose = 1
	MESSAGE_PURPOSE_VERIFIED               MessagePurpose = 2
	MESSAGE_PURPOSE_FORGET_PIN_CODE        MessagePurpose = 3
	MESSAGE_PURPOSE_REMOVE_PASSKEY         MessagePurpose = 4
	MESSAGE_PURPOSE_SET_CARD_PIN_CODE      MessagePurpose = 5
)

type NationCode string // 國碼
const (
	NATION_CODE_USA NationCode = "USA"
	NATION_CODE_JPN NationCode = "JPN"
	NATION_CODE_AUT NationCode = "AUT" // EU Austria
	NATION_CODE_BEL NationCode = "BEL" // EU Belgium
	NATION_CODE_BGR NationCode = "BGR" // EU Bulgaria
	NATION_CODE_CYP NationCode = "CYP" // EU Cyprus
	NATION_CODE_CZE NationCode = "CZE" // EU Czech Republic
	NATION_CODE_DEU NationCode = "DEU" // EU Germany
	NATION_CODE_DNK NationCode = "DNK" // EU Denmark
	NATION_CODE_ESP NationCode = "ESP" // EU Spain
	NATION_CODE_EST NationCode = "EST" // EU Estonia
	NATION_CODE_FIN NationCode = "FIN" // EU Finland
	NATION_CODE_FRA NationCode = "FRA" // EU France
	NATION_CODE_GRC NationCode = "GRC" // EU Greece
	NATION_CODE_HUN NationCode = "HUN" // EU Hungary
	NATION_CODE_IRL NationCode = "IRL" // EU Ireland
	NATION_CODE_ITA NationCode = "ITA" // EU Italy
	NATION_CODE_LAT NationCode = "LAT" // EU Latvia
	NATION_CODE_LIT NationCode = "LIT" // EU Lithuania
	NATION_CODE_LUX NationCode = "LUX" // EU Luxembourg
	NATION_CODE_MLT NationCode = "MLT" // EU Malta
	NATION_CODE_NLD NationCode = "NLD" // EU Netherlands
	NATION_CODE_POL NationCode = "POL" // EU Poland
	NATION_CODE_PRT NationCode = "PRT" // EU Portugal
	NATION_CODE_ROU NationCode = "ROU" // EU Romania
	NATION_CODE_SVK NationCode = "SVK" // EU Slovakia
	NATION_CODE_SVN NationCode = "SVN" // EU Slovenia
	NATION_CODE_SWE NationCode = "SWE" // EU Sweden
	NATION_CODE_GBR NationCode = "GBR"
	NATION_CODE_CHN NationCode = "CHN"
	NATION_CODE_CAN NationCode = "CAN"
	NATION_CODE_AUS NationCode = "AUS"
	NATION_CODE_CHE NationCode = "CHE"
	NATION_CODE_HKG NationCode = "HKG"
	NATION_CODE_NZL NationCode = "NZL"
	NATION_CODE_MEX NationCode = "MEX"
	NATION_CODE_SGP NationCode = "SGP"
	NATION_CODE_NOR NationCode = "NOR"
	NATION_CODE_KOR NationCode = "KOR"
	NATION_CODE_TUR NationCode = "TUR"
	NATION_CODE_IND NationCode = "IND"
	NATION_CODE_BRA NationCode = "BRA"
	NATION_CODE_ZAF NationCode = "ZAF"
	NATION_CODE_RUS NationCode = "RUS"
	NATION_CODE_TWN NationCode = "TWN"
	NATION_CODE_ISR NationCode = "ISR"
	NATION_CODE_ARE NationCode = "ARE"
	NATION_CODE_ARG NationCode = "ARG"
	NATION_CODE_MYS NationCode = "MYS"
	NATION_CODE_THA NationCode = "THA"
	NATION_CODE_SAU NationCode = "SAU"
	NATION_CODE_IDN NationCode = "IDN" // 印尼 (Indonesia)
	NATION_CODE_PAK NationCode = "PAK" // 巴基斯坦 (Pakistan)
	NATION_CODE_NGA NationCode = "NGA" // 奈及利亞 (Nigeria)
	NATION_CODE_EGY NationCode = "EGY" // 埃及 (Egypt)
	NATION_CODE_VNM NationCode = "VNM" // 越南 (Vietnam)
	NATION_CODE_PHL NationCode = "PHL" // 菲律賓 (Philippines)
	NATION_CODE_COL NationCode = "COL" // 哥倫比亞 (Colombia)
	NATION_CODE_CHL NationCode = "CHL" // 智利 (Chile)
	NATION_CODE_PER NationCode = "PER" // 秘魯 (Peru)
	NATION_CODE_DZA NationCode = "DZA" // 阿爾及利亞 (Algeria)
	NATION_CODE_MAR NationCode = "MAR" // 摩洛哥 (Morocco)
	NATION_CODE_KEN NationCode = "KEN" // 肯亞 (Kenya)
	NATION_CODE_UKR NationCode = "UKR" // 烏克蘭 (Ukraine)
	NATION_CODE_AFG NationCode = "AFG" // 阿富汗 (Afghanistan)
	NATION_CODE_ETH NationCode = "ETH" // 埃塞俄比亞 (Ethiopia)
	NATION_CODE_COD NationCode = "COD" // 剛果民主共和國 (Democratic Republic of the Congo)
	NATION_CODE_TZA NationCode = "TZA" // 坦尚尼亞 (Tanzania)
	NATION_CODE_UZB NationCode = "UZB" // 烏茲別克 (Uzbekistan)
	NATION_CODE_ALA NationCode = "ALA" // Åland Islands
	NATION_CODE_ALB NationCode = "ALB" // Albania
	NATION_CODE_ASM NationCode = "ASM" // American Samoa
	NATION_CODE_AND NationCode = "AND" // Andorra
	NATION_CODE_AGO NationCode = "AGO" // Angola
	NATION_CODE_AIA NationCode = "AIA" // Anguilla
	NATION_CODE_ATA NationCode = "ATA" // Antarctica
	NATION_CODE_ATG NationCode = "ATG" // Antigua and Barbuda
	NATION_CODE_ARM NationCode = "ARM" // Armenia
	NATION_CODE_ABW NationCode = "ABW" // Aruba
	NATION_CODE_AZE NationCode = "AZE" // Azerbaijan
	NATION_CODE_BHS NationCode = "BHS" // Bahamas
	NATION_CODE_BHR NationCode = "BHR" // Bahrain
	NATION_CODE_BGD NationCode = "BGD" // Bangladesh
	NATION_CODE_BRB NationCode = "BRB" // Barbados
	NATION_CODE_BLR NationCode = "BLR" // Belarus
	NATION_CODE_BLZ NationCode = "BLZ" // Belize
	NATION_CODE_BEN NationCode = "BEN" // Benin
	NATION_CODE_BMU NationCode = "BMU" // Bermuda
	NATION_CODE_BTN NationCode = "BTN" // Bhutan
	NATION_CODE_BOL NationCode = "BOL" // Bolivia (Plurinational State of)
	NATION_CODE_BES NationCode = "BES" // Bonaire, Sint Eustatius and Saba
	NATION_CODE_BIH NationCode = "BIH" // Bosnia and Herzegovina
	NATION_CODE_BWA NationCode = "BWA" // Botswana
	NATION_CODE_BVT NationCode = "BVT" // Bouvet Island
	NATION_CODE_IOT NationCode = "IOT" // British Indian Ocean Territory
	NATION_CODE_BRN NationCode = "BRN" // Brunei Darussalam
	NATION_CODE_BFA NationCode = "BFA" // Burkina Faso
	NATION_CODE_BDI NationCode = "BDI" // Burundi
	NATION_CODE_CPV NationCode = "CPV" // Cabo Verde
	NATION_CODE_KHM NationCode = "KHM" // Cambodia
	NATION_CODE_CMR NationCode = "CMR" // Cameroon
	NATION_CODE_CYM NationCode = "CYM" // Cayman Islands
	NATION_CODE_CAF NationCode = "CAF" // Central African Republic
	NATION_CODE_TCD NationCode = "TCD" // Chad
	NATION_CODE_CXR NationCode = "CXR" // Christmas Island
	NATION_CODE_CCK NationCode = "CCK" // Cocos (Keeling) Islands
	NATION_CODE_COM NationCode = "COM" // Comoros
	NATION_CODE_COG NationCode = "COG" // Congo
	NATION_CODE_COK NationCode = "COK" // Cook Islands
	NATION_CODE_CRI NationCode = "CRI" // Costa Rica
	NATION_CODE_CIV NationCode = "CIV" // Côte d'Ivoire
	NATION_CODE_HRV NationCode = "HRV" // Croatia
	NATION_CODE_CUB NationCode = "CUB" // Cuba
	NATION_CODE_CUW NationCode = "CUW" // Curaçao
	NATION_CODE_DJI NationCode = "DJI" // Djibouti
	NATION_CODE_DMA NationCode = "DMA" // Dominica
	NATION_CODE_DOM NationCode = "DOM" // Dominican Republic
	NATION_CODE_ECU NationCode = "ECU" // Ecuador
	NATION_CODE_SLV NationCode = "SLV" // El Salvador
	NATION_CODE_GNQ NationCode = "GNQ" // Equatorial Guinea
	NATION_CODE_ERI NationCode = "ERI" // Eritrea
	NATION_CODE_SWZ NationCode = "SWZ" // Eswatini
	NATION_CODE_FLK NationCode = "FLK" // Falkland Islands (Malvinas)
	NATION_CODE_FRO NationCode = "FRO" // Faroe Islands
	NATION_CODE_FJI NationCode = "FJI" // Fiji
	NATION_CODE_GUF NationCode = "GUF" // French Guiana
	NATION_CODE_PYF NationCode = "PYF" // French Polynesia
	NATION_CODE_ATF NationCode = "ATF" // French Southern Territories
	NATION_CODE_GAB NationCode = "GAB" // Gabon
	NATION_CODE_GMB NationCode = "GMB" // Gambia
	NATION_CODE_GEO NationCode = "GEO" // Georgia
	NATION_CODE_GHA NationCode = "GHA" // Ghana
	NATION_CODE_GIB NationCode = "GIB" // Gibraltar
	NATION_CODE_GRL NationCode = "GRL" // Greenland
	NATION_CODE_GRD NationCode = "GRD" // Grenada
	NATION_CODE_GLP NationCode = "GLP" // Guadeloupe
	NATION_CODE_GUM NationCode = "GUM" // Guam
	NATION_CODE_GTM NationCode = "GTM" // Guatemala
	NATION_CODE_GGY NationCode = "GGY" // Guernsey
	NATION_CODE_GIN NationCode = "GIN" // Guinea
	NATION_CODE_GNB NationCode = "GNB" // Guinea-Bissau
	NATION_CODE_GUY NationCode = "GUY" // Guyana
	NATION_CODE_HTI NationCode = "HTI" // Haiti
	NATION_CODE_HMD NationCode = "HMD" // Heard Island and McDonald Islands
	NATION_CODE_VAT NationCode = "VAT" // Holy See
	NATION_CODE_HND NationCode = "HND" // Honduras
	NATION_CODE_ISL NationCode = "ISL" // Iceland
	NATION_CODE_IRN NationCode = "IRN" // Iran (Islamic Republic of)
	NATION_CODE_IRQ NationCode = "IRQ" // Iraq
	NATION_CODE_IMN NationCode = "IMN" // Isle of Man
	NATION_CODE_JAM NationCode = "JAM" // Jamaica
	NATION_CODE_JEY NationCode = "JEY" // Jersey
	NATION_CODE_JOR NationCode = "JOR" // Jordan
	NATION_CODE_KAZ NationCode = "KAZ" // Kazakhstan
	NATION_CODE_KIR NationCode = "KIR" // Kiribati
	NATION_CODE_PRK NationCode = "PRK" // Korea (Democratic People's Republic of)
	NATION_CODE_KWT NationCode = "KWT" // Kuwait
	NATION_CODE_KGZ NationCode = "KGZ" // Kyrgyzstan
	NATION_CODE_LAO NationCode = "LAO" // Lao People's Democratic Republic
	NATION_CODE_LVA NationCode = "LVA" // Latvia
	NATION_CODE_LBN NationCode = "LBN" // Lebanon
	NATION_CODE_LSO NationCode = "LSO" // Lesotho
	NATION_CODE_LBR NationCode = "LBR" // Liberia
	NATION_CODE_LBY NationCode = "LBY" // Libya
	NATION_CODE_LIE NationCode = "LIE" // Liechtenstein
	NATION_CODE_LTU NationCode = "LTU" // Lithuania
	NATION_CODE_MAC NationCode = "MAC" // Macao
	NATION_CODE_MDG NationCode = "MDG" // Madagascar
	NATION_CODE_MWI NationCode = "MWI" // Malawi
	NATION_CODE_MDV NationCode = "MDV" // Maldives
	NATION_CODE_MLI NationCode = "MLI" // Mali
	NATION_CODE_MHL NationCode = "MHL" // Marshall Islands
	NATION_CODE_MTQ NationCode = "MTQ" // Martinique
	NATION_CODE_MRT NationCode = "MRT" // Mauritania
	NATION_CODE_MUS NationCode = "MUS" // Mauritius
	NATION_CODE_MYT NationCode = "MYT" // Mayotte
	NATION_CODE_FSM NationCode = "FSM" // Micronesia (Federated States of)
	NATION_CODE_MDA NationCode = "MDA" // Moldova, Republic of
	NATION_CODE_MCO NationCode = "MCO" // Monaco
	NATION_CODE_MNG NationCode = "MNG" // Mongolia
	NATION_CODE_MNE NationCode = "MNE" // Montenegro
	NATION_CODE_MSR NationCode = "MSR" // Montserrat
	NATION_CODE_MOZ NationCode = "MOZ" // Mozambique
	NATION_CODE_MMR NationCode = "MMR" // Myanmar
	NATION_CODE_NAM NationCode = "NAM" // Namibia
	NATION_CODE_NRU NationCode = "NRU" // Nauru
	NATION_CODE_NPL NationCode = "NPL" // Nepal
	NATION_CODE_NCL NationCode = "NCL" // New Caledonia
	NATION_CODE_NIC NationCode = "NIC" // Nicaragua
	NATION_CODE_NER NationCode = "NER" // Niger
	NATION_CODE_NIU NationCode = "NIU" // Niue
	NATION_CODE_NFK NationCode = "NFK" // Norfolk Island
	NATION_CODE_MKD NationCode = "MKD" // North Macedonia
	NATION_CODE_MNP NationCode = "MNP" // Northern Mariana Islands
	NATION_CODE_OMN NationCode = "OMN" // Oman
	NATION_CODE_PLW NationCode = "PLW" // Palau
	NATION_CODE_PSE NationCode = "PSE" // Palestine, State of
	NATION_CODE_PAN NationCode = "PAN" // Panama
	NATION_CODE_PNG NationCode = "PNG" // Papua New Guinea
	NATION_CODE_PRY NationCode = "PRY" // Paraguay
	NATION_CODE_PCN NationCode = "PCN" // Pitcairn
	NATION_CODE_PRI NationCode = "PRI" // Puerto Rico
	NATION_CODE_QAT NationCode = "QAT" // Qatar
	NATION_CODE_REU NationCode = "REU" // Réunion
	NATION_CODE_RWA NationCode = "RWA" // Rwanda
	NATION_CODE_BLM NationCode = "BLM" // Saint Barthélemy
	NATION_CODE_SHN NationCode = "SHN" // Saint Helena, Ascension and Tristan da Cunha
	NATION_CODE_KNA NationCode = "KNA" // Saint Kitts and Nevis
	NATION_CODE_LCA NationCode = "LCA" // Saint Lucia
	NATION_CODE_MAF NationCode = "MAF" // Saint Martin (French part)
	NATION_CODE_SPM NationCode = "SPM" // Saint Pierre and Miquelon
	NATION_CODE_VCT NationCode = "VCT" // Saint Vincent and the Grenadines
	NATION_CODE_WSM NationCode = "WSM" // Samoa
	NATION_CODE_SMR NationCode = "SMR" // San Marino
	NATION_CODE_STP NationCode = "STP" // Sao Tome and Principe
	NATION_CODE_SEN NationCode = "SEN" // Senegal
	NATION_CODE_SRB NationCode = "SRB" // Serbia
	NATION_CODE_SYC NationCode = "SYC" // Seychelles
	NATION_CODE_SLE NationCode = "SLE" // Sierra Leone
	NATION_CODE_SXM NationCode = "SXM" // Sint Maarten (Dutch part)
	NATION_CODE_SLB NationCode = "SLB" // Solomon Islands
	NATION_CODE_SOM NationCode = "SOM" // Somalia
	NATION_CODE_SGS NationCode = "SGS" // South Georgia and the South Sandwich Islands
	NATION_CODE_SSD NationCode = "SSD" // South Sudan
	NATION_CODE_LKA NationCode = "LKA" // Sri Lanka
	NATION_CODE_SDN NationCode = "SDN" // Sudan
	NATION_CODE_SUR NationCode = "SUR" // Suriname
	NATION_CODE_SJM NationCode = "SJM" // Svalbard and Jan Mayen
	NATION_CODE_SYR NationCode = "SYR" // Syrian Arab Republic
	NATION_CODE_TJK NationCode = "TJK" // Tajikistan
	NATION_CODE_TLS NationCode = "TLS" // Timor-Leste
	NATION_CODE_TGO NationCode = "TGO" // Togo
	NATION_CODE_TKL NationCode = "TKL" // Tokelau
	NATION_CODE_TON NationCode = "TON" // Tonga
	NATION_CODE_TTO NationCode = "TTO" // Trinidad and Tobago
	NATION_CODE_TUN NationCode = "TUN" // Tunisia
	NATION_CODE_TKM NationCode = "TKM" // Turkmenistan
	NATION_CODE_TCA NationCode = "TCA" // Turks and Caicos Islands
	NATION_CODE_TUV NationCode = "TUV" // Tuvalu
	NATION_CODE_UGA NationCode = "UGA" // Uganda
	NATION_CODE_UMI NationCode = "UMI" // United States Minor Outlying Islands
	NATION_CODE_URY NationCode = "URY" // Uruguay
	NATION_CODE_VUT NationCode = "VUT" // Vanuatu
	NATION_CODE_VEN NationCode = "VEN" // Venezuela (Bolivarian Republic of)
	NATION_CODE_VGB NationCode = "VGB" // Virgin Islands (British)
	NATION_CODE_VIR NationCode = "VIR" // Virgin Islands (U.S.)
	NATION_CODE_WLF NationCode = "WLF" // Wallis and Futuna
	NATION_CODE_ESH NationCode = "ESH" // Western Sahara
	NATION_CODE_YEM NationCode = "YEM" // Yemen
	NATION_CODE_ZMB NationCode = "ZMB" // Zambia
	NATION_CODE_ZWE NationCode = "ZWE" // Zimbabwe
)

type NotifyType int // 推播方式
const (
	NOTIFY_TYPE_EMAIL   NotifyType = 1
	NOTIFY_TYPE_SMS     NotifyType = 2
	NOTIFY_TYPE_APP     NotifyType = 3
	NOTIFY_TYPE_MESSAGE NotifyType = 4
)

type NotifyStatus int // 推播狀態
const (
	NOTIFY_STATUS_SENDING NotifyStatus = 1
	NOTIFY_STATUS_SENDED  NotifyStatus = 2
	NOTIFY_STATUS_SUCCESS NotifyStatus = 3
	NOTIFY_STATUS_FAILED  NotifyStatus = 4
)

type AnnouncementStatus uint8 // 1: draft, 2: wait for publish, 3: published, 4: unpublished
const (
	ANNOUNCEMENT_STATUS_DRAFT       AnnouncementStatus = 1
	ANNOUNCEMENT_STATUS_SCHEDULED   AnnouncementStatus = 2
	ANNOUNCEMENT_STATUS_PUBLISHED   AnnouncementStatus = 3
	ANNOUNCEMENT_STATUS_UNPUBLISHED AnnouncementStatus = 4
)

type OrderDirection int // 排序方向
const (
	ORDER_DIRECTION_ASC  OrderDirection = 1
	ORDER_DIRECTION_DESC OrderDirection = 2
)

type Protocol int // 協議
const (
	PROTOCOL_OMNI        Protocol = 1
	PROTOCOL_ERC20       Protocol = 2
	PROTOCOL_TRC20       Protocol = 3
	PROTOCOL_DEP20       Protocol = 4
	PROTOCOL_SPL         Protocol = 5
	PROTOCOL_BEP20       Protocol = 6
	PROTOCOL_MATIC_ERC20 Protocol = 7
)

type IconType int // 參數鍵
const (
	ICON_TYPE_WALLET IconType = 1
)

type IconKey string // 參數鍵
const (
	ICON_KEY_APPLE_PAY  IconKey = "wallet_apple_pay"
	ICON_KEY_GOOGLE_PAY IconKey = "wallet_google_pay"
	ICON_KEY_PAYPAL     IconKey = "wallet_paypal"
	ICON_KEY_ALI        IconKey = "wallet_alipay"
	ICON_KEY_WECHAT     IconKey = "wallet_wechatpay"
)

type ParameterCategory int // 參數類別
const (
	PARAMETER_CATEGORY_ACCOUNT                 ParameterCategory = 1
	PARAMETER_CATEGORY_ADDRESS_POOL            ParameterCategory = 2
	PARAMETER_CATEGORY_ADMIN                   ParameterCategory = 3
	PARAMETER_CATEGORY_AUTH                    ParameterCategory = 4
	PARAMETER_CATEGORY_CARD                    ParameterCategory = 5
	PARAMETER_CATEGORY_CARD_TO_CARD            ParameterCategory = 6
	PARAMETER_CATEGORY_COINS_DO                ParameterCategory = 7
	PARAMETER_CATEGORY_DEPOSIT                 ParameterCategory = 8
	PARAMETER_CATEGORY_EXCHANGE                ParameterCategory = 9
	PARAMETER_CATEGORY_NOTIFY                  ParameterCategory = 10
	PARAMETER_CATEGORY_ORDER                   ParameterCategory = 11
	PARAMETER_CATEGORY_PROMOTION               ParameterCategory = 12
	PARAMETER_CATEGORY_QUOTE                   ParameterCategory = 13
	PARAMETER_CATEGORY_REAP                    ParameterCategory = 14
	PARAMETER_CATEGORY_SYSTEM                  ParameterCategory = 15
	PARAMETER_CATEGORY_TOP_DOWN                ParameterCategory = 16
	PARAMETER_CATEGORY_TOP_UP                  ParameterCategory = 17
	PARAMETER_CATEGORY_TRANSFER                ParameterCategory = 18
	PARAMETER_CATEGORY_USER                    ParameterCategory = 19
	PARAMETER_CATEGORY_WALLET                  ParameterCategory = 20
	PARAMETER_CATEGORY_WITHDRAW                ParameterCategory = 21
	PARAMETER_CATEGORY_MERCHANT                ParameterCategory = 22
	PARAMETER_CATEGORY_USER_LIMIT              ParameterCategory = 23
	PARAMETER_CATEGORY_MERCHANT_ADJUST_BALANCE ParameterCategory = 24
	PARAMETER_CATEGORY_POINT                   ParameterCategory = 25
	PARAMETER_CATEGORY_CHIPPAY_EXPRESS         ParameterCategory = 26
	PARAMETER_CATEGORY_FINANCIAL               ParameterCategory = 27
	PARAMETER_CATEGORY_WHALE                   ParameterCategory = 28
	PARAMETER_CATEGORY_PAYCEYPTO               ParameterCategory = 29
	PARAMETER_CATEGORY_ETHERFI                 ParameterCategory = 30
)

// parameter key PARAMETER_KEY_{category}_{key}
type ParameterKey string // 參數鍵
const (
	PARAMETER_KEY_EXCHANGE_EXCHANGE_FEE                                        ParameterKey = "exchange_fee"
	PARAMETER_KEY_EXCHANGE_MERCHANT_EXCHANGE_FEE                               ParameterKey = "merchant_exchange_fee"
	PARAMETER_KEY_EXCHANGE_MERCHANT_AUTO_EXCHANGE_FEE                          ParameterKey = "merchant_auto_exchange_fee"
	PARAMETER_KEY_DEPOSIT_DEPOSIT_FEE                                          ParameterKey = "deposit_fee"
	PARAMETER_KEY_WITHDRAW_WITHDRAW_FEE                                        ParameterKey = "withdraw_fee"
	PARAMETER_KEY_CARD_APPLY_EXCHANGE_FEE                                      ParameterKey = "apply_exchange_fee"
	PARAMETER_KEY_TOP_UP_TOP_UP_FEE                                            ParameterKey = "top_up_fee"
	PARAMETER_KEY_TOP_UP_WHALE_TOP_UP_FEE                                      ParameterKey = "whale_top_up_fee"
	PARAMETER_KEY_TOP_UP_PAYCRYPTO_TOP_UP_FEE                                  ParameterKey = "paycrypto_top_up_fee"
	PARAMETER_KEY_TOP_UP_ETHERFI_TOP_UP_FEE                                    ParameterKey = "etherfi_top_up_fee"
	PARAMETER_KEY_TOP_UP_TOP_UP_EXCHANGE_FEE                                   ParameterKey = "top_up_exchange_fee"
	PARAMETER_KEY_TOP_UP_ETHERFI_TOP_UP_EXCHANGE_FEE                           ParameterKey = "etherfi_top_up_exchange_fee"
	PARAMETER_KEY_TOP_DOWN_TOP_DOWN_FEE                                        ParameterKey = "top_down_fee"
	PARAMETER_KEY_TOP_DOWN_TOP_DOWN_EXCHANGE_FEE                               ParameterKey = "top_down_exchange_fee"
	PARAMETER_KEY_TRANSFER_TRANSFER_FEE                                        ParameterKey = "transfer_fee"
	PARAMETER_KEY_TRANSFER_TRANSFER_EXCHANGE_FEE                               ParameterKey = "transfer_exchange_fee"
	PARAMETER_KEY_TRANSFER_MERCHANT_TRANSFER_FEE                               ParameterKey = "merchant_transfer_fee"
	PARAMETER_KEY_TRANSFER_MERCHANT_TRANSFER_EXCHANGE_FEE                      ParameterKey = "merchant_transfer_exchange_fee"
	PARAMETER_KEY_PAY_AUTO_TOP_UP_PRIORITY                                     ParameterKey = "pay_auto_top_up_priority"
	PARAMETER_KEY_CARD_TO_CARD_CARD_TO_CARD_FEE                                ParameterKey = "card_to_card_fee"
	PARAMETER_KEY_CARD_TO_CARD_CARD_TO_CARD_EXCHANGE_FEE                       ParameterKey = "card_to_card_exchange_fee"
	PARAMETER_KEY_CARD_TO_CARD_SELF_CARD_TO_CARD_FEE                           ParameterKey = "self_card_to_card_fee"
	PARAMETER_KEY_CARD_TO_CARD_SELF_CARD_TO_CARD_EXCHANGE_FEE                  ParameterKey = "self_card_to_card_exchange_fee"
	PARAMETER_KEY_CARD_TO_CARD_WHALE_CARD_TO_CARD_FEE                          ParameterKey = "whale_card_to_card_fee"
	PARAMETER_KEY_CARD_TO_CARD_WHALE_CARD_TO_CARD_EXCHANGE_FEE                 ParameterKey = "whale_card_to_card_exchange_fee"
	PARAMETER_KEY_CARD_TO_CARD_WHALE_REAP_CARD_TO_CARD_FEE                     ParameterKey = "whale_reap_card_to_card_fee"
	PARAMETER_KEY_CARD_TO_CARD_WHALE_REAP_CARD_TO_CARD_EXCHANGE_FEE            ParameterKey = "whale_reap_card_to_card_exchange_fee"
	PARAMETER_KEY_CARD_TO_CARD_WHALE_SELF_CARD_TO_CARD_FEE                     ParameterKey = "whale_self_card_to_card_fee"
	PARAMETER_KEY_CARD_TO_CARD_WHALE_SELF_CARD_TO_CARD_EXCHANGE_FEE            ParameterKey = "whale_self_card_to_card_exchange_fee"
	PARAMETER_KEY_CARD_TO_CARD_PAYCRYPTO_REAP_CARD_TO_CARD_FEE                 ParameterKey = "paycrypto_reap_card_to_card_fee"
	PARAMETER_KEY_CARD_TO_CARD_PAYCRYPTO_REAP_CARD_TO_CARD_EXCHANGE_FEE        ParameterKey = "paycrypto_reap_card_to_card_exchange_fee"
	PARAMETER_KEY_WHALE_CARD_LIMIT_COUNTRY_US                                  ParameterKey = "whale_card_limit_country_us"
	PARAMETER_KEY_MERCHANT_ADJUST_BALANCE_MERCHANT_ADJUST_BALANCE_FEE          ParameterKey = "merchant_adjust_balance_fee"
	PARAMETER_KEY_MERCHANT_ADJUST_BALANCE_MERCHANT_ADJUST_BALANCE_EXCHANGE_FEE ParameterKey = "merchant_adjust_balance_exchange_fee"
	PARAMETER_KEY_REAP_DECLINE_RATE                                            ParameterKey = "reap_decline_rate"
	PARAMETER_KEY_REAP_DECLINE_INTERVAL                                        ParameterKey = "reap_decline_interval"
	PARAMETER_KEY_REAP_DECLINE_FINE                                            ParameterKey = "reap_decline_fine"
	PARAMETER_KEY_REAP_MAX_CONSECUTIVE_FAILURES                                ParameterKey = "reap_max_consecutive_failures"
	PARAMETER_KEY_REAP_NO_FEE_THRESHOLD                                        ParameterKey = "reap_no_fee_threshold"
	PARAMETER_KEY_REAP_AUTHORIZATION_FEE                                       ParameterKey = "reap_authorization_fee"
	PARAMETER_KEY_REAP_AUTHORIZATION_FEE_CHARGING_MODE                         ParameterKey = "reap_authorization_fee_charging_mode"
	PARAMETER_KEY_REAP_AUTHORIZATION_FEE_COUNTRY                               ParameterKey = "reap_authorization_fee_country"
	PARAMETER_KEY_REAP_ATM_RESTRICTED_COUNTRIES                                ParameterKey = "reap_atm_restricted_countries"
	PARAMETER_KEY_OPT_LIMIT_TOP                                                ParameterKey = "opt_limit_top"
	PARAMETER_KEY_OPT_LIMIT_INTERVAL                                           ParameterKey = "opt_limit_interval"
	PARAMETER_KEY_OPT_LIMIT_LOCK_TIME                                          ParameterKey = "opt_limit_lock_time"
	PARAMETER_KEY_MERCHANT_DEAULT_ACTIVATION_DEPOSIT                           ParameterKey = "merchant_default_activation_deposit"
	PARAMETER_KEY_USER_LIMIT_MONTHLY_TOTAL                                     ParameterKey = "user_limit_monthly_total"
	PARAMETER_KEY_USER_LIMIT_MONTHLY_COUNT                                     ParameterKey = "user_limit_monthly_count"
	PARAMETER_KEY_USER_LIMIT_DAILY_TOTAL                                       ParameterKey = "user_limit_daily_total"
	PARAMETER_KEY_USER_LIMIT_DAILY_COUNT                                       ParameterKey = "user_limit_daily_count"
	PARAMETER_KEY_USER_LIMIT_PER_TRANSACTION                                   ParameterKey = "user_limit_per_transaction"
	PARAMETER_KEY_USER_LIMIT_PER_TRANSACTION_MAX                               ParameterKey = "user_limit_per_transaction_max"
	PARAMETER_KEY_MAIL_TEMPLATE_STYLE                                          ParameterKey = "mail_template_style"
	PARAMETER_KEY_MAIL_TEMPLATE_FOOTER                                         ParameterKey = "mail_template_footer"
	PARAMETER_KEY_IOS_VERSION_LIMIT                                            ParameterKey = "ios_version_limit"
	PARAMETER_KEY_ANDROID_VERSION_LIMIT                                        ParameterKey = "android_version_limit"
	PARAMETER_KEY_TOP_UP_POINT_TYPE                                            ParameterKey = "top_up_point_type"
	PARAMETER_KEY_CHECK_IN_POINT_TYPE                                          ParameterKey = "check_in_point_type"
	PARAMETER_KEY_REFERRAL_REGISTER_POINT_TYPE                                 ParameterKey = "referral_register_point_type"
	PARAMETER_KEY_REFERRAL_APPLY_POINT_TYPE                                    ParameterKey = "referral_apply_point_type"
	PARAMETER_KEY_EPOINT_LEVEL_REQUIREMENT                                     ParameterKey = "epoint_level_requirement"
	PARAMETER_KEY_CHIPPAY_EXPRESS_CRYPTO_FEE                                   ParameterKey = "chippay_express_crypto_fee"
	PARAMETER_KEY_CHIPPAY_EXPRESS_CARD_PRODUCT_FEE                             ParameterKey = "chippay_express_card_product_fee"
	PARAMETER_KEY_REAP_BALANCE_ALERT_THRESHOLD                                 ParameterKey = "reap_balance_alert_threshold"
	PARAMETER_KEY_WHALE_BALANCE_ALERT_THRESHOLD                                ParameterKey = "whale_balance_alert_threshold"
	PARAMETER_KEY_PAYCRYPTO_BALANCE_ALERT_THRESHOLD                            ParameterKey = "paycrypto_balance_alert_threshold"
	PARAMETER_KEY_AUTO_YIELD_THRESHOLD                                         ParameterKey = "auto_yield_threshold"
	PARAMETER_KEY_AUTO_YIELD_CRYPTO_BUTTON_TOGGLE                              ParameterKey = "auto_yield_crypto_button_toggle"
	PARAMETER_KEY_CHANGELLY_USDT_DEFAULT_MAINNET                               ParameterKey = "changelly_usdt_default_mainnet"
	PARAMETER_KEY_USER_INVITE_CODE_RATIO                                       ParameterKey = "user_invite_code_ratio"       // commission
	PARAMETER_KEY_USER_INVITE_CODE_APPLY_RATIO                                 ParameterKey = "user_invite_code_apply_ratio" // promotion discount
	PARAMETER_KEY_EXCHANGE_SOURCE                                              ParameterKey = "exchange_rate_source"
	PARAMETER_KEY_BINANCE_PAY_DEPOSIT_FEE                                      ParameterKey = "binance_pay_deposit_fee"
	PARAMETER_KEY_BINANCE_PAY_CARD_DEPOSIT_FEE                                 ParameterKey = "binance_pay_card_deposit_fee"
	PARAMETER_KEY_BINANCE_PAY_CARD_DEPOSIT_FEE_REAP                            ParameterKey = "binance_pay_card_deposit_fee_reap"
	PARAMETER_KEY_BINANCE_PAY_CARD_DEPOSIT_FEE_PAYCRYPTO                       ParameterKey = "binance_pay_card_deposit_fee_paycrypto"
	PARAMETER_KEY_BINANCE_PAY_CARD_DEPOSIT_FEE_BINANCEPAY                      ParameterKey = "binance_pay_card_deposit_fee_binancepay"
	PARAMETER_KEY_BINANCE_PAY_DEPOSIT_LIMIT                                    ParameterKey = "binance_pay_deposit_min_amount"
	PARAMETER_KEY_TELEGRAM_BOT_SYSTEM_NOTICE_CHAT_ID                           ParameterKey = "telegram_bot_system_notice_chat_id"
	PARAMETER_KEY_MONETA_PROMOTION_CODE                                        ParameterKey = "moneta_promotion_code"
	PARAMETER_KEY_ISSUE_CARD_LIMITED_COUNTRY_LIST                              ParameterKey = "issue_card_limited_country_list"
	PARAMETER_KEY_CARD_BLOCK_DELETE_PERIOD                                     ParameterKey = "card_block_delete_period"
	PARAMETER_KEY_CARD_TRANSACTION_OVER_LIMIT_SETTING                          ParameterKey = "card_transaction_over_limit_setting"
	PARAMETER_KEY_DEPOSIT_OVER_LIMIT_SETTING                                   ParameterKey = "deposit_over_limit_setting"
	PARAMETER_KEY_MERCHANT_BALANCE_SETTING                                     ParameterKey = "merchant_balance_setting"
	PARAMETER_KEY_BINANCE_PAY_DEPOSIT_LIMIT_USDT                               ParameterKey = "binance_pay_deposit_min_amount_usdt"
	PARAMETER_KEY_BINANCE_PAY_DEPOSIT_LIMIT_ETH                                ParameterKey = "binance_pay_deposit_min_amount_eth"
	PARAMETER_KEY_BINANCE_PAY_DEPOSIT_LIMIT_BTC                                ParameterKey = "binance_pay_deposit_min_amount_btc"
	PARAMETER_KEY_BINANCE_PAY_DEPOSIT_LIMIT_USDC                               ParameterKey = "binance_pay_deposit_min_amount_usdc"
	PARAMETER_KEY_BINANCE_PAY_DEPOSIT_LIMIT_DOGE                               ParameterKey = "binance_pay_deposit_min_amount_doge"
	PARAMETER_KEY_BINANCE_PAY_DEPOSIT_LIMIT_TRX                                ParameterKey = "binance_pay_deposit_min_amount_trx"
	PARAMETER_KEY_BINANCE_PAY_DEPOSIT_LIMIT_ADA                                ParameterKey = "binance_pay_deposit_min_amount_ada"
	PARAMETER_KEY_BINANCE_PAY_DEPOSIT_LIMIT_SOL                                ParameterKey = "binance_pay_deposit_min_amount_sol"
	PARAMETER_KEY_BINANCE_PAY_DEPOSIT_LIMIT_BNB                                ParameterKey = "binance_pay_deposit_min_amount_bnb"
	PARAMETER_KEY_BINANCE_PAY_DEPOSIT_LIMIT_XRP                                ParameterKey = "binance_pay_deposit_min_amount_xrp"
	PARAMETER_KEY_ETHERFI_RAIN_PUBLICKEY                                       ParameterKey = "etherfi_rain_publickey"
	PARAMETER_KEY_ETHERFI_ADMIN_MAIL                                           ParameterKey = "etherfi_admin_mail"
	PARAMETER_KEY_ETHERFI_CREATE_CARD_AMOUNT                                   ParameterKey = "etherfi_creat_card_amount"
	PARAMETER_KEY_CARD_APPLY_USER_WHITE_LIST                                   ParameterKey = "card_apply_user_white_list"
	PARAMETER_KEY_ETHERFI_BALANCE_ALERT_THRESHOLD                              ParameterKey = "etherfi_balance_alert_threshold"
	PARAMETER_KEY_CARD_AMOUNT_LOW_LIMIT_SETTING                                ParameterKey = "card_amount_low_limit_setting"
	PARAMETER_KEY_PHYSICAL_CARD_CARD_USER_FILTER                               ParameterKey = "physical_card_card_user_filter"
	PARAMETER_KEY_ETHERFI_API_BATCH                                            ParameterKey = "etherfi_api_batch"
)

type ParameterStatus int // 參數狀態
const (
	PARAMETER_STATUS_ACTIVATED   ParameterStatus = 1
	PARAMETER_STATUS_DEACTIVATED ParameterStatus = 2
)

type ParameterValueType int // 參數數值型別
const (
	PARAMETER_VALUE_TYPE_STRING          ParameterValueType = 10
	PARAMETER_VALUE_TYPE_NUMBER          ParameterValueType = 20
	PARAMETER_VALUE_TYPE_PERCENTAGE      ParameterValueType = 21
	PARAMETER_VALUE_TYPE_AMOUNT          ParameterValueType = 22
	PARAMETER_VALUE_TYPE_BOOLEAN         ParameterValueType = 30
	PARAMETER_VALUE_TYPE_TIME            ParameterValueType = 40
	PARAMETER_VALUE_TYPE_UNIX_TIME       ParameterValueType = 41
	PARAMETER_VALUE_TYPE_UNIX_TIME_MILLI ParameterValueType = 42
	PARAMETER_VALUE_TYPE_RFC3339         ParameterValueType = 43
	PARAMETER_VALUE_TYPE_INTERVAL        ParameterValueType = 50
	PARAMETER_VALUE_TYPE_ARRAY           ParameterValueType = 60
	PARAMETER_VALUE_TYPE_EXPRESSION      ParameterValueType = 70
	PARAMETER_VALUE_TYPE_REGEX           ParameterValueType = 71
	PARAMETER_VALUE_TYPE_ARITHMETIC      ParameterValueType = 72
	PARAMETER_VALUE_TYPE_TEXT            ParameterValueType = 80
	PARAMETER_VALUE_TYPE_JSON            ParameterValueType = 81
	PARAMETER_VALUE_TYPE_YAML            ParameterValueType = 82
	PARAMETER_VALUE_TYPE_XML             ParameterValueType = 83
)

type Passageway string // 通道
const (
	PASSAGEWAY_SENDGRID Passageway = "sendgrid"
	PASSAGEWAY_MAILGUN  Passageway = "mailgun"
)

type PassagewayStatus int // 通道狀態
const (
	PASSAGEWAY_STATUS_ACTIVE   PassagewayStatus = 1
	PASSAGEWAY_STATUS_INACTIVE PassagewayStatus = 2
)

type PaycryptoChannel int // paycrypto交易管道
const (
	PAYCRYPTO_CHANNEL_ATM PaycryptoChannel = 1
)

type PaycryptoPhysicalCardNumberStatus int // paycrypto 實體卡卡號狀態
const (
	PAYCRYPTO_PHYSICAL_CARD_NUMBER_STATUS_UNUSED PaycryptoPhysicalCardNumberStatus = 1
	PAYCRYPTO_PHYSICAL_CARD_NUMBER_STATUS_LOCK   PaycryptoPhysicalCardNumberStatus = 2
	PAYCRYPTO_PHYSICAL_CARD_NUMBER_STATUS_USED   PaycryptoPhysicalCardNumberStatus = 3
)

type PaycryptoDepositOrderPurpose int // paycrypto充帳目的
const (
	PAYCRYPTO_DEPOSIT_ORDER_PURPOSE_TOP_UP PaycryptoDepositOrderPurpose = 1
	PAYCRYPTO_DEPOSIT_ORDER_PURPOSE_CP_BUY PaycryptoDepositOrderPurpose = 2
)

type PaycryptoWebhookAction string // paycrypto回調事件
const (
	PAYCRYPTO_WEBHOOK_ACTION_KYC_STATUS                  PaycryptoWebhookAction = "kyc-status"
	PAYCRYPTO_WEBHOOK_ACTION_CARD_APPLICATION_READY      PaycryptoWebhookAction = "card-application-ready"
	PAYCRYPTO_WEBHOOK_ACTION_CARD_STATUS                 PaycryptoWebhookAction = "card-status"
	PAYCRYPTO_WEBHOOK_ACTION_TX_STATUS                   PaycryptoWebhookAction = "tx-status"
	PAYCRYPTO_WEBHOOK_ACTION_CREDITCARD_BALANCE_UPDATED  PaycryptoWebhookAction = "creditcard-balance-updated"
	PAYCRYPTO_WEBHOOK_ACTION_DEPOSITCARD_BALANCE_UPDATED PaycryptoWebhookAction = "depositcard-balance-updated"
	PAYCRYPTO_WEBHOOK_ACTION_CARD_3DS_OTP                PaycryptoWebhookAction = "card-3ds-otp"
	PAYCRYPTO_WEBHOOK_ACTION_CARD_LOCK_STATUS            PaycryptoWebhookAction = "card-lock-status"
)

// 状态码: 0.处理中 1.成功 2.失败
type PaycryptoWebhookCardStatus int // paycrypto卡片狀態
const (
	PAYCRYPTO_WEBHOOK_CARD_STATUS_PROCESSING PaycryptoWebhookCardStatus = 0
	PAYCRYPTO_WEBHOOK_CARD_STATUS_SUCCESS    PaycryptoWebhookCardStatus = 1
	PAYCRYPTO_WEBHOOK_CARD_STATUS_FAILED     PaycryptoWebhookCardStatus = 2
)

// 1.冻结 2.解冻 3.挂失 4.重置密码 5.补卡 6.重发PIN
type PaycryptoWebhookCardRequestType int // 卡片工單類型
const (
	PAYCRYPTO_WEBHOOK_CARD_REQUEST_TYPE_FROZEN    PaycryptoWebhookCardRequestType = 1
	PAYCRYPTO_WEBHOOK_CARD_REQUEST_TYPE_UNFROZEN  PaycryptoWebhookCardRequestType = 2
	PAYCRYPTO_WEBHOOK_CARD_REQUEST_TYPE_LOST      PaycryptoWebhookCardRequestType = 3
	PAYCRYPTO_WEBHOOK_CARD_REQUEST_TYPE_RESET_PIN PaycryptoWebhookCardRequestType = 4
	PAYCRYPTO_WEBHOOK_CARD_REQUEST_TYPE_REPAIR    PaycryptoWebhookCardRequestType = 5
)

type PaycryptoWebhookKYCStatus int // paycrypto KYC回調狀態
const (
	PAYCRYPTO_WEBHOOK_KYC_STATUS_PASS PaycryptoWebhookKYCStatus = 1
	PAYCRYPTO_WEBHOOK_KYC_STATUS_FAIL PaycryptoWebhookKYCStatus = 2
)

type PaycryptoWebhookTXStatus int // paycrypto交易狀態
const (
	PAYCRYPTO_WEBHOOK_TX_STATUS_SUCCESS   PaycryptoWebhookTXStatus = 1
	PAYCRYPTO_WEBHOOK_TX_STATUS_FAILED    PaycryptoWebhookTXStatus = 2
	PAYCRYPTO_WEBHOOK_TX_STATUS_CANCELLED PaycryptoWebhookTXStatus = 5
)

type PaycryptoPayStatus int // paycrypto支付狀態
const (
	PAYCRYPTO_PAY_STATUS_NONE     PaycryptoPayStatus = 0
	PAYCRYPTO_PAY_STATUS_PENDING  PaycryptoPayStatus = 1
	PAYCRYPTO_PAY_STATUS_CLOSED   PaycryptoPayStatus = 2
	PAYCRYPTO_PAY_STATUS_FAILED   PaycryptoPayStatus = 3
	PAYCRYPTO_PAY_STATUS_REFUNDED PaycryptoPayStatus = 4
	PAYCRYPTO_PAY_STATUS_REVERSAL PaycryptoPayStatus = 5
)

// 交易类型，1.消费 2.充值 3.取款 4.转账(转入) 5.转账(转出) 6.其他 7.结算调整 8.退款 9.消费失败 10.verification（绑卡验证交易），11.void（撤销,付款后撤销，退回资金）
type PaycryptoTxRecordType int // paycrypto交易紀錄類型

const (
	PAYCRYPTO_TX_RECORD_TYPE_PAY          PaycryptoTxRecordType = 1  // 消费
	PAYCRYPTO_TX_RECORD_TYPE_DEPOSIT      PaycryptoTxRecordType = 2  // 充值
	PAYCRYPTO_TX_RECORD_TYPE_WITHDRAWAL   PaycryptoTxRecordType = 3  // 取款
	PAYCRYPTO_TX_RECORD_TYPE_TRANSFER_IN  PaycryptoTxRecordType = 4  // 转账(转入)
	PAYCRYPTO_TX_RECORD_TYPE_TRANSFER_OUT PaycryptoTxRecordType = 5  // 转账(转出)
	PAYCRYPTO_TX_RECORD_TYPE_OTHER        PaycryptoTxRecordType = 6  // 其他
	PAYCRYPTO_TX_RECORD_TYPE_CLOSE        PaycryptoTxRecordType = 7  // 结算调整
	PAYCRYPTO_TX_RECORD_TYPE_REFUND       PaycryptoTxRecordType = 8  // 退款
	PAYCRYPTO_TX_RECORD_TYPE_FAILED       PaycryptoTxRecordType = 9  // 消费失败
	PAYCRYPTO_TX_RECORD_TYPE_VERIFICATION PaycryptoTxRecordType = 10 // verification（绑卡验证交易）
	PAYCRYPTO_TX_RECORD_TYPE_VOID         PaycryptoTxRecordType = 11 // void（撤销,付款后撤销，退回资金）
)

// 状态码: 0 冻结， 1 激活成功， 2未激活， 3. 激活待审核， 4. 激活审核失败, 5. 申请失败(卡片正在制作中，请过会再申请) 6.已注销
type PaycryptoCardActivateStatus int // 卡片激活狀態
const (
	PAYCRYPTO_CARD_ACTIVATE_STATUS_FROZEN        PaycryptoCardActivateStatus = 0
	PAYCRYPTO_CARD_ACTIVATE_STATUS_ACTIVE        PaycryptoCardActivateStatus = 1
	PAYCRYPTO_CARD_ACTIVATE_STATUS_UNACTIVE      PaycryptoCardActivateStatus = 2
	PAYCRYPTO_CARD_ACTIVATE_STATUS_WAIT_REVIEW   PaycryptoCardActivateStatus = 3
	PAYCRYPTO_CARD_ACTIVATE_STATUS_REVIEW_FAILED PaycryptoCardActivateStatus = 4
	PAYCRYPTO_CARD_ACTIVATE_STATUS_APPLY_FAILED  PaycryptoCardActivateStatus = 5
	PAYCRYPTO_CARD_ACTIVATE_STATUS_CANCEL        PaycryptoCardActivateStatus = 6
)

// 状态码: 0.处理中 1.成功 2.失败
type PaycryptoCardRequestStatus int // 卡片工單狀態
const (
	PAYCRYPTO_CARD_REQUEST_STATUS_PROCESSING PaycryptoCardRequestStatus = 0
	PAYCRYPTO_CARD_REQUEST_STATUS_SUCCESS    PaycryptoCardRequestStatus = 1
	PAYCRYPTO_CARD_REQUEST_STATUS_FAILED     PaycryptoCardRequestStatus = 2
)

// 1.冻结 2.解冻 3.挂失 4.重置密码 5.补卡 6.重发PIN
type PaycryptoCardRequestType int // 卡片工單類型
const (
	PAYCRYPTO_CARD_REQUEST_TYPE_FROZEN    PaycryptoCardRequestType = 1
	PAYCRYPTO_CARD_REQUEST_TYPE_UNFROZEN  PaycryptoCardRequestType = 2
	PAYCRYPTO_CARD_REQUEST_TYPE_LOST      PaycryptoCardRequestType = 3
	PAYCRYPTO_CARD_REQUEST_TYPE_RESET_PIN PaycryptoCardRequestType = 4
	PAYCRYPTO_CARD_REQUEST_TYPE_REPAIR    PaycryptoCardRequestType = 5
	PAYCRYPTO_CARD_REQUEST_TYPE_REPIN     PaycryptoCardRequestType = 6
)

// 状态：2. 开卡申请成功(未激活)， 5. 申请失败(卡片正在制作中，请过会再申请)
type PaycryptoCardApplicationStatus int // 卡片申请狀態
const (
	PAYCRYPTO_CARD_APPLICATION_STATUS_SUCCESS PaycryptoCardApplicationStatus = 2
	PAYCRYPTO_CARD_APPLICATION_STATUS_FAILED  PaycryptoCardApplicationStatus = 5
)

type PaycryptoDepositOrderStatus int // 充值狀態
const (
	PAYCRYPTO_DEPOSIT_ORDER_STATUS_PENDING    PaycryptoDepositOrderStatus = 1
	PAYCRYPTO_DEPOSIT_ORDER_STATUS_PROCESSING PaycryptoDepositOrderStatus = 2
	PAYCRYPTO_DEPOSIT_ORDER_STATUS_SUCCESS    PaycryptoDepositOrderStatus = 3
	PAYCRYPTO_DEPOSIT_ORDER_STATUS_FAILED     PaycryptoDepositOrderStatus = 4
)

// 状态码: 0 已提交， 1 认证通过(开卡成功)， 2 认证未通过， 3 认证中， 4 提交信息处理中 6 已退款
type PaycryptoKYCStatus int // KYC狀態
const (
	PAYCRYPTO_KYC_STATUS_SUBMITTED PaycryptoKYCStatus = 0
	PAYCRYPTO_KYC_STATUS_PASSED    PaycryptoKYCStatus = 1
	PAYCRYPTO_KYC_STATUS_FAILED    PaycryptoKYCStatus = 2
	PAYCRYPTO_KYC_STATUS_PROCESS   PaycryptoKYCStatus = 3
	PAYCRYPTO_KYC_STATUS_REFUNDED  PaycryptoKYCStatus = 6
)

type PaycryptoKYCSubmitStatus int // KYC 提交狀態
const (
	PAYCRYPTO_KYC_SUBMIT_STATUS_NOT_REQUESTED PaycryptoKYCSubmitStatus = 1
	PAYCRYPTO_KYC_SUBMIT_STATUS_REQUESTED     PaycryptoKYCSubmitStatus = 2
	PAYCRYPTO_KYC_SUBMIT_STATUS_PENDING       PaycryptoKYCSubmitStatus = 3
	PAYCRYPTO_KYC_SUBMIT_STATUS_SUCCESS       PaycryptoKYCSubmitStatus = 4
	PAYCRYPTO_KYC_SUBMIT_STATUS_FAILED        PaycryptoKYCSubmitStatus = 5
)

// 交易状态。0、3、4:待处理中，1:充值成功，2充值失败，5:充值失败
type PaycryptoTXStatus int // 交易狀態
const (
	PAYCRYPTO_TX_STATUS_PENDING0 PaycryptoTXStatus = 0
	PAYCRYPTO_TX_STATUS_SUCCESS  PaycryptoTXStatus = 1
	PAYCRYPTO_TX_STATUS_FAILED2  PaycryptoTXStatus = 2
	PAYCRYPTO_TX_STATUS_PENDING3 PaycryptoTXStatus = 3
	PAYCRYPTO_TX_STATUS_PENDING4 PaycryptoTXStatus = 4
	PAYCRYPTO_TX_STATUS_FAILED5  PaycryptoTXStatus = 5
)

type PointSource int // 點數來源
const (
	POINT_SOURCE_TOP_UP                          PointSource = 1
	POINT_SOURCE_CHECK_IN                        PointSource = 2
	POINT_SOURCE_REFERRAL_REGISTER               PointSource = 3
	POINT_SOURCE_REFERRED_REGISTER               PointSource = 4
	POINT_SOURCE_REFERRAL_APPLY                  PointSource = 5
	POINT_SOURCE_REFERRED_APPLY                  PointSource = 6
	POINT_SOURCE_GIFTED                          PointSource = 7
	POINT_SOURCE_GAME_REWARD                     PointSource = 8
	POINT_SOURCE_CHIPPAY_EXPRESS_TO_CARD_PRODUCT PointSource = 9
)

type PointStatus int // 點數狀態
const (
	POINT_STATUS_USABLE   PointStatus = 1
	POINT_STATUS_USED     PointStatus = 2
	POINT_STATUS_EXPIRED  PointStatus = 3
	POINT_STATUS_PENDING  PointStatus = 4
	POINT_STATUS_FAILED   PointStatus = 5
	POINT_STATUS_REVERSAL PointStatus = 6
)

type PreviewPurpose int // 預覽目的
const (
	PREVIEW_PURPOSE_APPLY        PreviewPurpose = 1
	PREVIEW_PURPOSE_TRANSFER     PreviewPurpose = 2
	PREVIEW_PURPOSE_EXCHANGE     PreviewPurpose = 3
	PREVIEW_PURPOSE_WITHDRAW     PreviewPurpose = 4
	PREVIEW_PURPOSE_TOP_UP       PreviewPurpose = 7
	PREVIEW_PURPOSE_TOP_DOWN     PreviewPurpose = 8
	PREVIEW_PURPOSE_CARD_TO_CARD PreviewPurpose = 11
	PREVIEW_PURPOSE_UNBLOCK      PreviewPurpose = 12
)

type PromotionType int // 促銷類型
const (
	PROMOTIOM_TYPE_APPLY      PromotionType = 1
	PROMOTIOM_TYPE_EXCHANGE   PromotionType = 2
	PROMOTIOM_TYPE_TRANSFER   PromotionType = 3
	PROMOTIOM_TYPE_WITHDRAWAL PromotionType = 4
)

type RatePurpose int // 匯率用途
const (
	RATE_PURPOSE_APPLY               RatePurpose = 1
	RATE_PURPOSE_EXCHANGE            RatePurpose = 2
	RATE_PURPOSE_TRANSFER            RatePurpose = 3
	RATE_PURPOSE_CARD_TO_CARD        RatePurpose = 4
	RATE_PURPOSE_TOP_UP              RatePurpose = 5
	RATE_PURPOSE_TOP_DOWN            RatePurpose = 6
	RATE_PURPOSE_FEE                 RatePurpose = 7
	RATE_PURPOSE_LIMIT               RatePurpose = 8
	RATE_PURPOSE_MERCHANT_EXCHANGE   RatePurpose = 9
	RATE_PURPOSE_MERCHANT_TRANSFER   RatePurpose = 10
	RATE_PURPOSE_DISCOUNT            RatePurpose = 11
	RATE_PURPOSE_DISPLAY             RatePurpose = 12
	RATE_PURPOSE_CHIPPAY_EXPRESS_BUY RatePurpose = 13
	RATE_PURPOSE_TOP_UP_POINT        RatePurpose = 14
)

type ReapAffectBalance int // reap影響金額
const (
	REAP_AFFECT_BALANCE_YES ReapAffectBalance = 1
	REAP_AFFECT_BALANCE_NO  ReapAffectBalance = 2
)

type ReapCardType string // reap卡片種類
const (
	REAP_CARD_TYPE_VIRTUAL  ReapCardType = "Virtual"
	REAP_CARD_TYPE_PHYSICAL ReapCardType = "Physical"
)

type ReapConsumerType string // reap客戶種類
const (
	REAP_CONSUMER_TYPE_BUSINESS ReapConsumerType = "Business"
	REAP_CONSUMER_TYPE_CONSUMER ReapConsumerType = "Consumer"
)

type ReapChannel int // reap交易管道
const (
	REAP_CHANNEL_ATM         ReapChannel = 1
	REAP_CHANNEL_POS         ReapChannel = 2
	REAP_CHANNEL_ECOMMERCE   ReapChannel = 3
	REAP_CHANNEL_VISA_DIRECT ReapChannel = 4
)

type ReapDataLoss int // reap數據損失
const (
	REAP_DATA_LOSS_OK          ReapDataLoss = 1
	REAP_DATA_LOSS_DATA_LOSS   ReapDataLoss = 2
	REAP_DATA_LOSS_NOT_FREEZED ReapDataLoss = 3
	REAP_DATA_LOSS_RECOVERED   ReapDataLoss = 4
)

type ReapDeleteStatus int // reap刪除狀態
const (
	REAP_DELETE_STATUS_PENDING ReapDeleteStatus = 1
	REAP_DELETE_STATUS_CLOSE   ReapDeleteStatus = 2
)

type ReapEventName string // 事件類型
const (
	REAP_EVENT_NAME_REQUEST                ReapEventName = "request"                // Triggered when authorization is initialized
	REAP_EVENT_NAME_AUTHORIZATION          ReapEventName = "authorization"          // Triggered when a transaction is authorized.
	REAP_EVENT_NAME_AUTHORIZATION_ADVICE   ReapEventName = "authorization.advice"   // Triggered when a transaction has been declined
	REAP_EVENT_NAME_AUTHORIZATION_REVERSAL ReapEventName = "authorization.reversal" // Triggered when an authorized transaction amount has been reversed.
	REAP_EVENT_NAME_AUTHORIZATION_CLEARING ReapEventName = "authorization.clearing" // Triggered when an authorized transaction amount has been cleared.
	REAP_EVENT_NAME_REFUND                 ReapEventName = "refund"                 // Triggered when an authorized transaction amount has been refunded.

	REAP_EVENT_NAME_BIOMETRIC_AUTH_NOTIFICATION_RECEIVED ReapEventName = "biometric_auth_notification_received"
	REAP_EVENT_NAME_WALLET_OPT_NOTIFICATION              ReapEventName = "wallet_tokenization_otp"

	REAP_EVENT_NAME_CARD_STATUS ReapEventName = "card_status"

	REAP_EVENT_NAME_MONTHLY_FILE ReapEventName = "monthly_file" // Triggered when a monthly file is generated for a cardholder.
	REAP_EVENT_NAME_DAILY_FILE   ReapEventName = "daily_file"   // Triggered when a daily file is generated for a cardholder.

	REAP_EVENT_NAME_SHIPPING_ORDER_CONFIRMATION ReapEventName = "shipping_order_confirmation"
)

type ReapEventType string // 事件類型
const (
	REAP_EVENT_TYPE_TRANSATION    ReapEventType = "transaction"
	REAP_EVENT_TYPE_AUTHORIZATION ReapEventType = "authorization"
	REAP_EVENT_TYPE_CARD          ReapEventType = "card"
	REAP_EVENT_TYPE_NOTIFICATION  ReapEventType = "notification"
	REAP_EVENT_TYPE_ACCOUNT       ReapEventType = "account"
	REAP_EVENT_TYPE_FILES         ReapEventType = "files"
	REAP_EVENT_TYPE_SHIPPING      ReapEventType = "shipping"
)

type ReapAuthorizationFeeChargingMode int // reap 授權費支付方式
const (
	REAP_AUTHORIZATION_FEE_CHARGING_MODE_ALL             ReapAuthorizationFeeChargingMode = 1
	REAP_AUTHORIZATION_FEE_CHARGING_MODE_COUNTRY_INCLUDE ReapAuthorizationFeeChargingMode = 2
	REAP_AUTHORIZATION_FEE_CHARGING_MODE_COUNTRY_EXCLUDE ReapAuthorizationFeeChargingMode = 3
	REAP_AUTHORIZATION_FEE_CHARGING_MODE_NONE            ReapAuthorizationFeeChargingMode = 4
)

type UserDeleteReason int

const (
	USER_STATUS_EUSD_ADMIN_DELETE            UserDeleteReason = 1
	USER_STATUS_EUSD_CUSTOMER_REQUIRE_DELETE UserDeleteReason = 2
)

type UserBlockReason int

const (
	USER_STATUS_EUSD_USER_BLOCK             UserBlockReason = 1
	USER_STATUS_EUSD_CUSTOMER_REQUIRE_BLOCK UserBlockReason = 2
	USER_STATUS_EUSD_REAP_KYC               UserBlockReason = 3
)

type CardBlockReason string // reap卡片狀態
const (
	// Fraud Block Reasons
	CARD_STATUS_EUSD_USER_BLOCK             CardBlockReason = "EUSD_USER_BLOCK"
	CARD_STATUS_EUSD_USER_ALL_BLOCK         CardBlockReason = "EUSD_USER_ALL_BLOCK"
	CARD_STATUS_EUSD_CUSTOMER_REQUIRE_BLOCK CardBlockReason = "EUSD_CUSTOMER_REQUIRE_BLOCK"
	CARD_STATUS_EUSD_REAP_KYC               CardBlockReason = "EUSD_REAP_KYC"

	CARD_STATUS_SUSPECTED_FRAUD_CVV    CardBlockReason = "SUSPECTED_FRAUD_CVV"
	CARD_STATUS_SUSPECTED_FRAUD_EXPIRY CardBlockReason = "SUSPECTED_FRAUD_EXPIRY"
	CARD_STATUS_SUSPECTED_FRAUD_PIN    CardBlockReason = "SUSPECTED_FRAUD_PIN"

	// Fund Limitations
	CARD_STATUS_FUND_INSUFFICIENT_BALANCE            CardBlockReason = "INSUFFICIENT_BALANCE"
	CARD_STATUS_FUND_DAILY_WITHDRAW_EXCEEDED         CardBlockReason = "DAILY_WITHDRAW_EXCEEDED"
	CARD_STATUS_FUND_DAILY_TX_EXCEEDED               CardBlockReason = "DAILY_TX_EXCEEDED"
	CARD_STATUS_FUND_TX_AMOUNT_EXCEEDED              CardBlockReason = "TX_AMOUNT_EXCEEDED"
	CARD_STATUS_FUND_EXCEED_MAX_CONSECUTIVE_FAILURES CardBlockReason = "EXCEED_MAX_CONSECUTIVE_FAILURES"

	// Security Issues
	CARD_STATUS_SECURITY_RISKY_TX         CardBlockReason = "RISKY_TX"
	CARD_STATUS_SECURITY_INCORRECT_PIN    CardBlockReason = "INCORRECT_PIN"
	CARD_STATUS_SECURITY_UNUSUAL_LOCATION CardBlockReason = "UNUSUAL_LOCATION"
	CARD_STATUS_SECURITY_FOREIGN_TX       CardBlockReason = "FOREIGN_TX"
	CARD_STATUS_SECURITY_CARD_LOST        CardBlockReason = "CARD_LOST"
	CARD_STATUS_SECURITY_CARD_STOLEN      CardBlockReason = "CARD_STOLEN"
	CARD_STATUS_SECURITY_RISKY_PAYEE      CardBlockReason = "RISKY_PAYEE"

	// Technical Issues
	CARD_STATUS_TECH_PAYEE_AUTH_FAIL CardBlockReason = "PAYEE_AUTH_FAIL"
	CARD_STATUS_TECH_DEVICE_STRICT   CardBlockReason = "DEVICE_STRICT"
	CARD_STATUS_TECH_PAYEE_STRICT    CardBlockReason = "PAYEE_STRICT"

	// Legal or Compliance
	CARD_STATUS_LEGAL_SUSPICIOUS_TX       CardBlockReason = "SUSPICIOUS_TX"
	CARD_STATUS_LEGAL_FOREIGN_COMPLIANCE  CardBlockReason = "FOREIGN_COMPLIANCE"
	CARD_STATUS_LEGAL_RESTRICTED_CATEGORY CardBlockReason = "RESTRICTED_CATEGORY"
	CARD_STATUS_LEGAL_FRAUD_RISK          CardBlockReason = "FRAUD_RISK"
	CARD_STATUS_LEGAL_AGE_LIMIT           CardBlockReason = "AGE_LIMIT"

	// Account Issues
	CARD_STATUS_ACCOUNT_IDENTITY_INCORRECT     CardBlockReason = "IDENTITY_INCORRECT"     // holder identity incorrect
	CARD_STATUS_ACCOUNT_RESTRICTED_BY_PROVIDER CardBlockReason = "RESTRICTED_BY_PROVIDER" // caused by provider
	CARD_STATUS_ACCOUNT_RESTRICTED_BY_HOLDER   CardBlockReason = "RESTRICTED_BY_HOLDER"   // caused by holder
	CARD_STATUS_ACCOUNT_MAIN_ACCOUNT           CardBlockReason = "MAIN_ACCOUNT"
	CARD_STATUS_ACCOUNT_AUTH_FAIL              CardBlockReason = "AUTH_FAIL"
	CARD_STATUS_ACCOUNT_ANNUAL_FEE_UNPAID      CardBlockReason = "ANNUAL_FEE_UNPAID"
	CARD_STATUS_ACCOUNT_PAST_DUE               CardBlockReason = "PAST_DUE"
	CARD_STATUS_ACCOUNT_NOT_IN_USE             CardBlockReason = "NOT_IN_USE"
	CARD_STATUS_ACCOUNT_INACTIVE               CardBlockReason = "INACTIVE"
	CARD_STATUS_ACCOUNT_HOLDER_DECEASED        CardBlockReason = "HOLDER_DECEASED" // holder is deceased
	CARD_STATUS_ACCOUNT_HOLDER_BANKRUPT        CardBlockReason = "HOLDER_BANKRUPT" // holder financial issues
	CARD_STATUS_ACCOUNT_INVALID_KYC            CardBlockReason = "INVALID_KYC"
)

type CardFreezeReason int // reap凍結原因
var (
	CARD_FREEZE_REASON_REAP_MANUAL           CardFreezeReason = 1
	CARD_FREEZE_REASON_REAP_NEGATIVE_BALANCE CardFreezeReason = 2
	CARD_FREEZE_REASON_WHALE_MANUAL          CardFreezeReason = 3
	CARD_FREEZE_REASON_WHALE_BY_SYSTEM       CardFreezeReason = 4
	CARD_FREEZE_REASON_EUSD_BLOCK            CardFreezeReason = 5
	CARD_FREEZE_REASON_PAYCRYPTO_MANUAL      CardFreezeReason = 6
	CARD_FREEZE_REASON_PAYCRYPTO_DELETED     CardFreezeReason = 7
	CARD_FREEZE_REASON_ETHERFI_MANUAL        CardFreezeReason = 8
)

type ReapResponseCode string // Reap回傳錯誤碼
const (
	REAP_RESPONSE_CODE_3DS_NOT_FOUND                               ReapResponseCode = "0313006"
	REAP_RESPONSE_CODE_INVALID_CARD_EXPIRY_DATE                    ReapResponseCode = "0302008"
	REAP_RESPONSE_CODE_INVALID_CARD_EXPIRY_DATE2                   ReapResponseCode = "0302005"
	REAP_RESPONSE_CODE_INVALID_PHONE_NUMBER                        ReapResponseCode = "0302009"
	REAP_RESPONSE_CODE_SPEND_LIMIT_IS_REQUIRED                     ReapResponseCode = "0302012"
	REAP_RESPONSE_CODE_CRYPTO_CURRENCY_IS_REQUIRED                 ReapResponseCode = "0321002"
	REAP_RESPONSE_CODE_INSUFFICIENT_BALANCE                        ReapResponseCode = "0302007"
	REAP_RESPONSE_CODE_COUNTRY_NOT_ISSUABLE                        ReapResponseCode = "0315001"
	REAP_RESPONSE_CODE_SECONDARY_CARD_NAME_NOT_SUPPORTED_BY_TYPE   ReapResponseCode = "0302019"
	REAP_RESPONSE_CODE_SECONDARY_CARD_NAME_NOT_SUPPORTED_BY_DESIGN ReapResponseCode = "0309007"
	REAP_RESPONSE_CODE_INSUFFICIENT_CREDIT                         ReapResponseCode = "0401001"
	REAP_RESPONSE_CODE_WITHDRAWAL_WALLET_NOT_ENABLED               ReapResponseCode = "0321001"
	REAP_RESPONSE_CODE_ROOT_BUDGET_NOT_FOUND                       ReapResponseCode = "0302006"
	REAP_RESPONSE_CODE_CARD_DESIGN_NOT_FOUND                       ReapResponseCode = "0309002"
	REAP_RESPONSE_CODE_CANNOT_CREATE_UNDER_THIS_BUSINESS           ReapResponseCode = "0302020"
	REAP_RESPONSE_CODE_SYSTEM_MAINTANCE                            ReapResponseCode = "0302010"
)

type ReapShippingStatus string // reap運送狀態
const (
	REAP_SHIPPING_STATUS_IMCOMPLETE          ReapShippingStatus = "incomplete"
	REAP_SHIPPING_STATUS_PENDING_APPROVAL    ReapShippingStatus = "pending_approval"
	REAP_SHIPPING_STATUS_PENDING_FULFILMENT  ReapShippingStatus = "pending_fulfilment"
	REAP_SHIPPING_STATUS_FULFILLED           ReapShippingStatus = "fulfilled"
	REAP_SHIPPING_STATUS_CANCELLED           ReapShippingStatus = "canceled"
	REAP_SHIPPING_STATUS_ORDER_IN_PRODUCTION ReapShippingStatus = "order_in_productio"
)

type ReapTransactionStatus int // reap交易狀態
const (
	REAP_TRANSACTION_STATUS_REQUESTING         ReapTransactionStatus = 1
	REAP_TRANSACTION_STATUS_AUTHORIZING        ReapTransactionStatus = 2
	REAP_TRANSACTION_STATUS_PENDING            ReapTransactionStatus = 3
	REAP_TRANSACTION_STATUS_REVERSAL           ReapTransactionStatus = 4
	REAP_TRANSACTION_STATUS_CLEARED            ReapTransactionStatus = 5
	REAP_TRANSACTION_STATUS_DECLINED           ReapTransactionStatus = 6
	REAP_TRANSACTION_STATUS_VOID               ReapTransactionStatus = 7
	REAP_TRANSACTION_STATUS_REFUNDED           ReapTransactionStatus = 8
	REAP_TRANSACTION_STATUS_UNRELATED_PENDING  ReapTransactionStatus = 9
	REAP_TRANSACTION_STATUS_UNRELATED_REFUNDED ReapTransactionStatus = 10
	REAP_TRANSACTION_STATUS_UNRELATED_FAILED   ReapTransactionStatus = 11
	REAP_TRANSACTION_STATUS_VERIFY_PENDING     ReapTransactionStatus = 12
	REAP_TRANSACTION_STATUS_VERIFY_SUCCESS     ReapTransactionStatus = 13
	REAP_TRANSACTION_STATUS_VERIFY_FAILED      ReapTransactionStatus = 14
)

type ReapTransactionType int // reap交易類型
const (
	REAP_TRANSACTION_TYPE_PAY    ReapTransactionType = 1
	REAP_TRANSACTION_TYPE_REFUND ReapTransactionType = 2
	REAP_TRANSACTION_TYPE_VERIFY ReapTransactionType = 3
)

type ReapWallet int // reap錢包
const (
	REAP_WALLET_APPLE_PAY ReapWallet = 1
)

type ReapIdDocType string // reap證件類型
const (
	REAP_ID_DOC_TYPE_PASSPORT       ReapIdDocType = "Passport"
	REAP_ID_DOC_TYPE_NATIONAL_ID    ReapIdDocType = "NationalID"
	REAP_ID_DOC_TYPE_DRIVER         ReapIdDocType = "DriversLicense"
	REAP_ID_DOC_TYPE_HEALTH         ReapIdDocType = "Health"
	REAP_ID_DOC_TYPE_TAX_ID_NUMBER  ReapIdDocType = "TaxIDNumber"
	REAP_ID_DOC_TYPE_SOCIAL_SERVICE ReapIdDocType = "SocialService"
)

type RewardOrderStatus int // 兌換訂單狀態
const (
	REWARD_ORDER_STATUS_PENDING    RewardOrderStatus = 1
	REWARD_ORDER_STATUS_PROCESSING RewardOrderStatus = 2
	REWARD_ORDER_STATUS_SUCCESS    RewardOrderStatus = 3
	REWARD_ORDER_STATUS_FAILED     RewardOrderStatus = 4
)

type RewardProductRestriction int // 兌換商品限制
const (
	REWARD_PRODUCT_RESTRICTION_UNRESTRICTED         RewardProductRestriction = 1
	REWARD_PRODUCT_RESTRICTION_PARTIALLY_ALLOWED    RewardProductRestriction = 2
	REWARD_PRODUCT_RESTRICTION_PARTIALLY_RESTRICTED RewardProductRestriction = 3
)

type RewardProductStatus int // 兌換商品狀態
const (
	REWARD_PRODUCT_STATUS_ACTIVE   RewardProductStatus = 1
	REWARD_PRODUCT_STATUS_INACTIVE RewardProductStatus = 2
)

type RewardProductCategoryStatus int // 兌換商品狀態
const (
	REWARD_PRODUCT_CATEGORY_STATUS_ACTIVE   RewardProductCategoryStatus = 1
	REWARD_PRODUCT_CATEGORY_STATUS_INACTIVE RewardProductCategoryStatus = 2
)

type RewardProductType int // 兌換商品類型
const (
	REWARD_PRODUCT_TYPE_VOUCHER  RewardProductType = 1
	REWARD_PRODUCT_TYPE_PHYSICAL RewardProductType = 2
	REWARD_PRODUCT_TYPE_GAME     RewardProductType = 3
	REWARD_PRODUCT_TYPE_WAIVER   RewardProductType = 4
)

type PincodeRequired int // 要求設置密碼
const (
	PINCODE_REQUIRED_YES       PincodeRequired = 1
	PINCODE_REQUIRED_NO        PincodeRequired = 2
	PINCODE_REQUIRED_DONT_CARE PincodeRequired = 3
)

type Role int // 角色
const (
	ROLE_GUEST          Role = 1
	ROLE_USER           Role = 2
	ROLE_ADMIN          Role = 3
	ROLE_MERCHANT       Role = 4
	ROLE_MERCHANT_ADMIN Role = 5
	ROLE_MERCHANT_USER  Role = 6
	ROLE_SYSTEM         Role = 9
)

type SingleFlightGroup int // singleflight組別
const (
	SINGLE_FLIGHT_GROUP_OPEN_LOG SingleFlightGroup = 1
)

type SnapshotMissingStatus int // 快照缺失狀態
const (
	SNAPSHOT_MISSING_STATUS_NOT_MISSING SnapshotMissingStatus = 1
	SNAPSHOT_MISSING_STATUS_MISSING     SnapshotMissingStatus = 2
)

type SpecialValue int // 特殊值
const (
	SPECIAL_VALUE_NULL      SpecialValue = 1
	SPECIAL_VALUE_UNLIMITED SpecialValue = 2
	SPECIAL_VALUE_POS_INF   SpecialValue = 3
	SPECIAL_VALUE_NEG_INF   SpecialValue = 4
	SPECIAL_VALUE_AUTO      SpecialValue = 5
	SPECIAL_VALUE_DEFAULT   SpecialValue = 6
	SPECIAL_VALUE_DYNAMIC   SpecialValue = 7
)

type TransactionType int // 交易動作
const (
	TRANSACTION_TYPE_ASSET_ADD           TransactionType = 1
	TRANSACTION_TYPE_ASSET_DEDUCT        TransactionType = 2
	TRANSACTION_TYPE_ASSET_FREEZE        TransactionType = 3
	TRANSACTION_TYPE_ASSET_UNFREEZE      TransactionType = 4
	TRANSACTION_TYPE_FROZEN_ASSET_ADD    TransactionType = 5
	TRANSACTION_TYPE_FROZEN_ASSET_DEDUCT TransactionType = 6
)

type TransferChannel int // 轉帳途徑
const (
	TRANSFER_CHANNEL_USER_ID  TransferChannel = 1
	TRANSFER_CHANNEL_EMAIL    TransferChannel = 2
	TRANSFER_CHANNEL_CARD_ID  TransferChannel = 3
	TRANSFER_CHANNEL_ADDRESS  TransferChannel = 4 // 一般轉帳專用，轉卡無此途徑
	TRANSFER_CHANNEL_WITHDRAW TransferChannel = 5 // 轉帳時輸入外部地址，自動轉提幣
)

type TransferStatus int // 轉帳狀態
const (
	TRANSFER_STATUS_PENDING TransferStatus = 1
	TRANSFER_STATUS_SUCCESS TransferStatus = 2
	TRANSFER_STATUS_FAILED  TransferStatus = 3
)

type WalletAddressStatus int // 錢包地址狀態
const (
	WALLET_ADDRESS_STATUS_ENABLED  WalletAddressStatus = 1
	WALLET_ADDRESS_STATUS_DISABLED WalletAddressStatus = 2
)

type WalletStatus int // 錢包狀態
const (
	WALLET_STATUS_ENABLED  WalletStatus = 1
	WALLET_STATUS_DISABLED WalletStatus = 2
)

type LockPurpose int // 鎖用途
const (
	LOCK_PURPOSE_APPLY_CONFIRM                  LockPurpose = 1
	LOCK_PURPOSE_REAP_TRANSACTION               LockPurpose = 2
	LOCK_PURPOSE_EXCHANGE_CONFIRM               LockPurpose = 3
	LOCK_PURPOSE_TRANSFER_CONFIRM               LockPurpose = 4
	LOCK_PURPOSE_TOP_UP_CONFIRM                 LockPurpose = 5
	LOCK_PURPOSE_TOP_DOWN_CONFIRM               LockPurpose = 6
	LOCK_PURPOSE_THREEDS_DAO                    LockPurpose = 7
	LOCK_PURPOSE_CARD_TO_CARD_CONFIRM           LockPurpose = 8
	LOCK_PURPOSE_WITHDRAW_CONFIRM               LockPurpose = 9
	LOCK_PURPOSE_MERCHANT_APPLY                 LockPurpose = 10
	LOCK_PURPOSE_CALLBACK                       LockPurpose = 11
	LOCK_PURPOSE_MANUAL                         LockPurpose = 12
	LOCK_PURPOSE_MERCHANT_THREEDS_DAO           LockPurpose = 13
	LOCK_PURPOSE_CREATE_POINT                   LockPurpose = 14
	LOCK_PURPOSE_CREATE_WALLET                  LockPurpose = 15
	LOCK_PURPOSE_TASK                           LockPurpose = 16
	LOCK_PURPOSE_CONSUME_POINT                  LockPurpose = 17
	LOCK_PURPOSE_CREATE_AUTO_YIELD              LockPurpose = 18
	LOCK_PURPOSE_DISTRIBUTE_AUTO_YIELD          LockPurpose = 19
	LOCK_PURPOSE_GET_TEMP_ADDRESS               LockPurpose = 20
	LOCK_PURPOSE_WHALE_WEBHOOK_BALANCE_EDIT     LockPurpose = 21
	LOCK_PURPOSE_WHALE_WEBHOOK_TRANSFER         LockPurpose = 22
	LOCK_PURPOSE_WHALE_WEBHOOK_CARD_TRANSACTION LockPurpose = 23
	LOCK_PURPOSE_WHALE_THREEDS_DAO              LockPurpose = 24
	LOCK_PURPOSE_CARD_STATUS_FROZEN             LockPurpose = 25
	LOCK_PURPOSE_WHALE_STATUS_DELETED           LockPurpose = 26
	LOCK_PURPOSE_WHALE_CHECK_ADJUST             LockPurpose = 27
	LOCK_PURPOSE_PAYCRYPTO_KYC_SUBMIT           LockPurpose = 28
	LOCK_PURPOSE_PAYCRYPTO_ACTIVATE             LockPurpose = 29
	LOCK_PURPOSE_PAYCRYPTO_WEBHOOK              LockPurpose = 30
	LOCK_PURPOSE_ADMIN_APPLY_CONFIRM            LockPurpose = 31
	LOCK_PURPOSE_ADMIN_DELETE_USER              LockPurpose = 32
	LOCK_PURPOSE_ADMIN_DELETE_CARD              LockPurpose = 33
	LOCK_PURPOSE_ETHEFI_SYNC_TRANSACTION        LockPurpose = 34
	LOCK_PURPOSE_ETHEFI_SYNC_CARD_TRANSACTION   LockPurpose = 35
	LOCK_PURPOSE_ETHEFI_SYNC_CARD_EVENT         LockPurpose = 36
	LOCK_PURPOSE_PIN_UNLOCK                     LockPurpose = 37
	LOCK_PURPOSE_ETHEFI_SYNC_SPENDING_LIMIT     LockPurpose = 38
	LOCK_PURPOSE_PHYSICAL_CARD_APPLY            LockPurpose = 41
	LOCK_PURPOSE_FINANCE_REPORT                 LockPurpose = 39
	LOCK_PURPOSE_ETHERFI_APPLY_CONFIRM          LockPurpose = 42
	LOCK_PURPOSE_ETHERFI_REPORT_DOWNLOAD        LockPurpose = 44
)

type ShippingStatus int // 貨運狀態
const (
	SHIPPING_STATUS_NOT_REQUESTED       ShippingStatus = 1
	SHIPPING_STATUS_IN_PROGRESS         ShippingStatus = 2
	SHIPPING_STATUS_IMCOMPLETE          ShippingStatus = 3
	SHIPPING_STATUS_PENDING_APPROVAL    ShippingStatus = 4
	SHIPPING_STATUS_PENDING_FULFILMENT  ShippingStatus = 5
	SHIPPING_STATUS_FULFILLED           ShippingStatus = 6
	SHIPPING_STATUS_CANCELLED           ShippingStatus = 7
	SHIPPING_STATUS_ORDER_IN_PRODUCTION ShippingStatus = 8
	SHIPPING_STATUS_REJECTED            ShippingStatus = 9
)

type SystemChargePurpose int // 系統扣款用途
const (
	SYSTEM_CHARGE_PURPOSE_DECLINE_FINE SystemChargePurpose = 1
)

type SystemChargeStatus int // 系統充值狀態
const (
	SYSTEM_CHARGE_STATUS_PENDING SystemChargeStatus = 1
	SYSTEM_CHARGE_STATUS_SUCCESS SystemChargeStatus = 2
	SYSTEM_CHARGE_STATUS_FAILED  SystemChargeStatus = 3
)

type TemplateVariable string // 模板參數
const (
	TEMPLATE_VARIABLE_RENEWAL_FEE                   TemplateVariable = "renewal_fee"
	TEMPLATE_VARIABLE_REAP_DECLINE_FINE             TemplateVariable = "reap_decline_fine"
	TEMPLATE_VARIABLE_TOP_UP_EXCHANGE_FEE           TemplateVariable = "top_up_exchange_fee"
	TEMPLATE_VARIABLE_TOP_DOWN_EXCHANGE_FEE         TemplateVariable = "top_down_exchange_fee"
	TEMPLATE_VARIABLE_REAP_ATM_RESTRICTED_COUNTRIES TemplateVariable = "reap_atm_restricted_countries"
	TEMPLATE_VARIABLE_ATM_DAILY_COUNT               TemplateVariable = "atm_daily_count"
	TEMPLATE_VARIABLE_ATM_MONTHLY_COUNT             TemplateVariable = "atm_monthly_count"
	TEMPLATE_VARIABLE_ATM_DAILY_TOTAL               TemplateVariable = "atm_daily_total"
	TEMPLATE_VARIABLE_ATM_MONTHLY_TOTAL             TemplateVariable = "atm_monthly_total"
	TEMPLATE_VARIABLE_ATM_YEARLY_TOTAL              TemplateVariable = "atm_yearly_total"
)

type ThreedsResponseType int // 3ds回應方式
const (
	THREEDS_RESPONSE_TYPE_IN_APP ThreedsResponseType = 1
	THREEDS_RESPONSE_TYPE_OTP    ThreedsResponseType = 2
)

type ThreedsStatus int // 3ds狀態
const (
	THREEDS_STATUS_PENDING        ThreedsStatus = 1
	THREEDS_STATUS_AUTHORIZING    ThreedsStatus = 2
	THREEDS_STATUS_AUTHORIZED     ThreedsStatus = 3
	THREEDS_STATUS_DECLINING      ThreedsStatus = 4
	THREEDS_STATUS_DECLINED       ThreedsStatus = 5
	THREEDS_STATUS_TIMEOUT        ThreedsStatus = 6
	THREEDS_STATUS_AUTO_SCHEDULED ThreedsStatus = 7
)

type TopUpOrderStatus int // 充值狀態
const (
	TOP_UP_ORDER_STATUS_PENDING    TopUpOrderStatus = 1
	TOP_UP_ORDER_STATUS_SUCCESS    TopUpOrderStatus = 2
	TOP_UP_ORDER_STATUS_FAILED     TopUpOrderStatus = 3
	TOP_UP_ORDER_STATUS_PROCESSING TopUpOrderStatus = 4
)

type TopDownOrderStatus int // 提現狀態
const (
	TOP_DOWN_ORDER_STATUS_PENDING    TopDownOrderStatus = 1
	TOP_DOWN_ORDER_STATUS_SUCCESS    TopDownOrderStatus = 2
	TOP_DOWN_ORDER_STATUS_FAILED     TopDownOrderStatus = 3
	TOP_DOWN_ORDER_STATUS_PROCESSING TopDownOrderStatus = 4
)

type TransactionSide int // 交易方
const (
	TRANSACTION_SIDE_FROM TransactionSide = 1
	TRANSACTION_SIDE_TO   TransactionSide = 2
)

type TransactionStatus int // 交易紀錄狀態
const (
	TRANSACTION_STATUS_APPLY_PENDING   TransactionStatus = 101
	TRANSACTION_STATUS_APPLY_CREATED   TransactionStatus = 102
	TRANSACTION_STATUS_APPLY_DEPOSITED TransactionStatus = 103
	TRANSACTION_STATUS_APPLY_FAILED    TransactionStatus = 104

	TRANSACTION_STATUS_TRANSFER_PENDING    TransactionStatus = 201
	TRANSACTION_STATUS_TRANSFER_PROCESSING TransactionStatus = 202
	TRANSACTION_STATUS_TRANSFER_SUCCESS    TransactionStatus = 203
	TRANSACTION_STATUS_TRANSFER_FAILED     TransactionStatus = 204

	TRANSACTION_STATUS_EXCHANGE_PENDING    TransactionStatus = 301
	TRANSACTION_STATUS_EXCHANGE_PROCESSING TransactionStatus = 302
	TRANSACTION_STATUS_EXCHANGE_SUCCESS    TransactionStatus = 303
	TRANSACTION_STATUS_EXCHANGE_FAILED     TransactionStatus = 304

	TRANSACTION_STATUS_WITHDRAW_PENDING           TransactionStatus = 401
	TRANSACTION_STATUS_WITHDRAW_FAILED            TransactionStatus = 402
	TRANSACTION_STATUS_WITHDRAW_SUCCESS           TransactionStatus = 403
	TRANSACTION_STATUS_WITHDRAW_FREEZED           TransactionStatus = 404
	TRANSACTION_STATUS_WITHDRAW_COINSDO_PENDING   TransactionStatus = 405
	TRANSACTION_STATUS_WITHDRAW_COINSDO_PASSED    TransactionStatus = 406
	TRANSACTION_STATUS_WITHDRAW_COINSDO_REJECTED  TransactionStatus = 407
	TRANSACTION_STATUS_WITHDRAW_COINSDO_CANCELLED TransactionStatus = 408
	TRANSACTION_STATUS_WITHDRAW_COINSDO_FAILED    TransactionStatus = 409

	TRANSACTION_STATUS_PAY_REQUESTING  TransactionStatus = 501
	TRANSACTION_STATUS_PAY_AUTHORIZING TransactionStatus = 502
	TRANSACTION_STATUS_PAY_PENDING     TransactionStatus = 503
	TRANSACTION_STATUS_PAY_REVERSAL    TransactionStatus = 504
	TRANSACTION_STATUS_PAY_CLEARED     TransactionStatus = 505
	TRANSACTION_STATUS_PAY_DECLINED    TransactionStatus = 506
	TRANSACTION_STATUS_PAY_VOID        TransactionStatus = 507
	TRANSACTION_STATUS_PAY_REFUNDED    TransactionStatus = 508

	TRANSACTION_STATUS_DEPOSIT_PENDING   TransactionStatus = 601
	TRANSACTION_STATUS_DEPOSIT_CONFIRMED TransactionStatus = 602
	TRANSACTION_STATUS_DEPOSIT_DEPOSITED TransactionStatus = 603
	TRANSACTION_STATUS_DEPOSIT_FAILED    TransactionStatus = 604

	TRANSACTION_STATUS_TOP_UP_PENDING    TransactionStatus = 701
	TRANSACTION_STATUS_TOP_UP_SUCCESS    TransactionStatus = 702
	TRANSACTION_STATUS_TOP_UP_FAILED     TransactionStatus = 703
	TRANSACTION_STATUS_TOP_UP_PROCESSING TransactionStatus = 704

	TRANSACTION_STATUS_TOP_DOWN_PENDING    TransactionStatus = 801
	TRANSACTION_STATUS_TOP_DOWN_SUCCESS    TransactionStatus = 802
	TRANSACTION_STATUS_TOP_DOWN_FAILED     TransactionStatus = 803
	TRANSACTION_STATUS_TOP_DOWN_PROCESSING TransactionStatus = 804

	TRANSACTION_STATUS_REFUND_PENDING TransactionStatus = 909
	TRANSACTION_STATUS_REFUND_SUCCESS TransactionStatus = 910
	TRANSACTION_STATUS_REFUND_FAILED  TransactionStatus = 911

	TRANSACTION_STATUS_VERIFY_PENDING TransactionStatus = 912
	TRANSACTION_STATUS_VERIFY_SUCCESS TransactionStatus = 913
	TRANSACTION_STATUS_VERIFY_FAILED  TransactionStatus = 914

	TRANSACTION_STATUS_CARD_TO_CARD_PENDING        TransactionStatus = 1101
	TRANSACTION_STATUS_CARD_TO_CARD_SUCCESS        TransactionStatus = 1102
	TRANSACTION_STATUS_CARD_TO_CARD_FAILED         TransactionStatus = 1103
	TRANSACTION_STATUS_CARD_TO_CARD_PROCESSING     TransactionStatus = 1104
	TRANSACTION_STATUS_CARD_TO_CARD_FAILED_HALFWAY TransactionStatus = 1105

	TRANSACTION_STATUS_MANUAL_PENDING       TransactionStatus = 1201
	TRANSACTION_STATUS_MANUAL_PENDING_AUDIT TransactionStatus = 1202
	TRANSACTION_STATUS_MANUAL_AUDITED       TransactionStatus = 1203
	TRANSACTION_STATUS_MANUAL_SUCCESS       TransactionStatus = 1204
	TRANSACTION_STATUS_MANUAL_AUDIT_FAILED  TransactionStatus = 1205
	TRANSACTION_STATUS_MANUAL_FAILED        TransactionStatus = 1206

	TRANSACTION_STATUS_MERCHANT_AUTO_EXCHANGE_PENDING    TransactionStatus = 1401
	TRANSACTION_STATUS_MERCHANT_AUTO_EXCHANGE_PROCESSING TransactionStatus = 1402
	TRANSACTION_STATUS_MERCHANT_AUTO_EXCHANGE_SUCCESS    TransactionStatus = 1403
	TRANSACTION_STATUS_MERCHANT_AUTO_EXCHANGE_FAILED     TransactionStatus = 1404

	TRANSACTION_STATUS_CHIPPAY_EXPRESS_PENDING    TransactionStatus = 1501
	TRANSACTION_STATUS_CHIPPAY_EXPRESS_PROCESSING TransactionStatus = 1502
	TRANSACTION_STATUS_CHIPPAY_EXPRESS_SUCCESS    TransactionStatus = 1503
	TRANSACTION_STATUS_CHIPPAY_EXPRESS_FAILED     TransactionStatus = 1504

	TRANSACTION_STATUS_MERCHANT_ADJUST_BALANCE_PENDING TransactionStatus = 1601
	TRANSACTION_STATUS_MERCHANT_ADJUST_BALANCE_SUCCESS TransactionStatus = 1602
	TRANSACTION_STATUS_MERCHANT_ADJUST_BALANCE_FAILED  TransactionStatus = 1603

	TRANSACTION_STATUS_POINT_ACCURAL_PENDING TransactionStatus = 1701
	TRANSACTION_STATUS_POINT_ACCURAL_SUCCESS TransactionStatus = 1702
	TRANSACTION_STATUS_POINT_ACCURAL_FAILED  TransactionStatus = 1703

	TRANSACTION_STATUS_POINT_REDEEM_PENDING    TransactionStatus = 1801
	TRANSACTION_STATUS_POINT_REDEEM_PROCESSING TransactionStatus = 1802
	TRANSACTION_STATUS_POINT_REDEEM_SUCCESS    TransactionStatus = 1803
	TRANSACTION_STATUS_POINT_REDEEM_FAILED     TransactionStatus = 1804

	TRANSACTION_STATUS_INTEREST_PENDING TransactionStatus = 1901
	TRANSACTION_STATUS_INTEREST_SUCCESS TransactionStatus = 1902
	TRANSACTION_STATUS_INTEREST_FAILED  TransactionStatus = 1903

	TRANSACTION_STATUS_FINANCIAL_TRANSFER_PENDING TransactionStatus = 2001
	TRANSACTION_STATUS_FINANCIAL_TRANSFER_SUCCESS TransactionStatus = 2002
	TRANSACTION_STATUS_FINANCIAL_TRANSFER_FAILED  TransactionStatus = 2003

	TRANSACTION_STATUS_WHALE_PAY_PENDING  TransactionStatus = 2101
	TRANSACTION_STATUS_WHALE_PAY_CLOSED   TransactionStatus = 2102
	TRANSACTION_STATUS_WHALE_PAY_FAILED   TransactionStatus = 2103
	TRANSACTION_STATUS_WHALE_PAY_CREDIT   TransactionStatus = 2104
	TRANSACTION_STATUS_WHALE_PAY_REVERSAL TransactionStatus = 2105

	TRANSACTION_STATUS_DISTRIBUTE_DISCOUNT_PENDING    TransactionStatus = 2301
	TRANSACTION_STATUS_DISTRIBUTE_DISCOUNT_PROCESSING TransactionStatus = 2302
	TRANSACTION_STATUS_DISTRIBUTE_DISCOUNT_SUCCESS    TransactionStatus = 2303
	TRANSACTION_STATUS_DISTRIBUTE_DISCOUNT_FAILED     TransactionStatus = 2304

	TRANSACTION_STATUS_PAYCRYPTO_PAY_PENDING  TransactionStatus = 2401
	TRANSACTION_STATUS_PAYCRYPTO_PAY_CLOSED   TransactionStatus = 2402
	TRANSACTION_STATUS_PAYCRYPTO_PAY_FAILED   TransactionStatus = 2403
	TRANSACTION_STATUS_PAYCRYPTO_PAY_REFUNDED TransactionStatus = 2404
	TRANSACTION_STATUS_PAYCRYPTO_PAY_REVERSAL TransactionStatus = 2405

	TRANSACTION_STATUS_PAYCRYPTO_REFUND_CLOSED TransactionStatus = 2502

	TRANSACTION_STATUS_BINANCE_PAY_PENDING    TransactionStatus = 2601
	TRANSACTION_STATUS_BINANCE_PAY_SUCCESS    TransactionStatus = 2602
	TRANSACTION_STATUS_BINANCE_PAY_FAILED     TransactionStatus = 2603
	TRANSACTION_STATUS_BINANCE_PAY_PROCESSING TransactionStatus = 2604
	TRANSACTION_STATUS_BINANCE_PAY_CLOSED     TransactionStatus = 2605

	TRANSACTION_STATUS_SYSTEM_CHARGE_PENDING TransactionStatus = 2801
	TRANSACTION_STATUS_SYSTEM_CHARGE_SUCCESS TransactionStatus = 2802
	TRANSACTION_STATUS_SYSTEM_CHARGE_FAILED  TransactionStatus = 2803

	TRANSACTION_STATUS_ETHERFI_PENDING   TransactionStatus = 3305
	TRANSACTION_STATUS_ETHERFI_CLEARED   TransactionStatus = 3306
	TRANSACTION_STATUS_ETHERFI_CANCELLED TransactionStatus = 3307
	TRANSACTION_STATUS_ETHERFI_DECLINED  TransactionStatus = 3308
	TRANSACTION_STATUS_ETHERFI_REFUND    TransactionStatus = 3309
	TRANSACTION_STATUS_ETHERFI_SUCCESS   TransactionStatus = 3310
	TRANSACTION_STATUS_ETHERFI_FAILED    TransactionStatus = 3311

	TRANSACTION_STATUS_ADMIN_DELETE_CARD_PENDING TransactionStatus = 3201
	TRANSACTION_STATUS_ADMIN_DELETE_CARD_SUCCESS TransactionStatus = 3202
)

type TransactionRecordType int // 交易紀錄類型
const (
	TRANSACTION_RECORD_TYPE_APPLY                   TransactionRecordType = 1
	TRANSACTION_RECORD_TYPE_TRANSFER                TransactionRecordType = 2
	TRANSACTION_RECORD_TYPE_EXCHANGE                TransactionRecordType = 3
	TRANSACTION_RECORD_TYPE_WITHDRAW                TransactionRecordType = 4
	TRANSACTION_RECORD_TYPE_PAY                     TransactionRecordType = 5
	TRANSACTION_RECORD_TYPE_DEPOSIT                 TransactionRecordType = 6
	TRANSACTION_RECORD_TYPE_TOP_UP                  TransactionRecordType = 7
	TRANSACTION_RECORD_TYPE_TOP_DOWN                TransactionRecordType = 8
	TRANSACTION_RECORD_TYPE_REFUND                  TransactionRecordType = 9
	TRANSACTION_RECORD_TYPE_VERIFY                  TransactionRecordType = 10
	TRANSACTION_RECORD_TYPE_CARD_TO_CARD            TransactionRecordType = 11
	TRANSACTION_RECORD_TYPE_MANUAL                  TransactionRecordType = 12 // 手動帳變的時候才用
	TRANSACTION_RECORD_TYPE_DELETE_CARD             TransactionRecordType = 13 //刪卡帳戶清零
	TRANSACTION_RECORD_TYPE_MERCHANT_AUTO_EXCHANGE  TransactionRecordType = 14
	TRANSACTION_RECORD_TYPE_CHIPPAY_EXPRESS         TransactionRecordType = 15
	TRANSACTION_RECORD_TYPE_MERCHANT_ADJUST_BALANCE TransactionRecordType = 16
	TRANSACTION_RECORD_TYPE_POINT_ACCURAL           TransactionRecordType = 17
	TRANSACTION_RECORD_TYPE_POINT_REDEEM            TransactionRecordType = 18
	TRANSACTION_RECORD_TYPE_INTEREST                TransactionRecordType = 19
	TRANSACTION_RECORD_TYPE_FINANCIAL_TRANSFER      TransactionRecordType = 20
	TRANSACTION_RECORD_TYPE_WHALE_PAY               TransactionRecordType = 21
	TRANSACTION_RECORD_TYPE_WHALE_REFUND            TransactionRecordType = 22
	TRANSACTION_RECORD_TYPE_DISTRIBUTE_APPLY        TransactionRecordType = 23
	TRANSACTION_RECORD_TYPE_PAYCRYPTO_PAY           TransactionRecordType = 24
	TRANSACTION_RECORD_TYPE_PAYCRYPTO_REFUND        TransactionRecordType = 25
	TRANSACTION_RECORD_TYPE_PAYCRYPTO_VERIFY        TransactionRecordType = 26
	TRANSACTION_RECORD_TYPE_BINANCE_PAY             TransactionRecordType = 27
	TRANSACTION_RECORD_TYPE_SYSTEM_CHARGE           TransactionRecordType = 28
	TRANSACTION_RECORD_TYPE_ADMIN_USER_DELETE       TransactionRecordType = 29
	TRANSACTION_RECORD_TYPE_ADMIN_CARD_DELETE       TransactionRecordType = 32
	TRANSACTION_RECORD_TYPE_ETHERFI                 TransactionRecordType = 33
	TRANSACTION_RECORD_TYPE_ETHERFI_REFUND          TransactionRecordType = 34
)

type TranslateMsg string // 翻譯訊息
const (
	// order
	TRANSLATE_MSG_ORDER_TITLE_APPLY                            TranslateMsg = "text_order_title_apply"
	TRANSLATE_MSG_ORDER_TITLE_TRANSFER_FROM                    TranslateMsg = "text_order_title_transfer_from"
	TRANSLATE_MSG_ORDER_TITLE_TRANSFER_TO                      TranslateMsg = "text_order_title_transfer_to"
	TRANSLATE_MSG_ORDER_TITLE_EXCHANGE                         TranslateMsg = "text_order_title_exchange"
	TRANSLATE_MSG_ORDER_TITLE_DEPOSIT                          TranslateMsg = "text_order_title_deposit"
	TRANSLATE_MSG_ORDER_TITLE_TOP_UP_TO                        TranslateMsg = "text_order_title_top_up_to"
	TRANSLATE_MSG_ORDER_TITLE_TOP_UP_FROM                      TranslateMsg = "text_order_title_top_up_from"
	TRANSLATE_MSG_ORDER_TITLE_TOP_DOWN_FROM                    TranslateMsg = "text_order_title_top_down_from"
	TRANSLATE_MSG_ORDER_TITLE_TOP_DOWN_TO                      TranslateMsg = "text_order_title_top_down_to"
	TRANSLATE_MSG_ORDER_TITLE_VERIFY                           TranslateMsg = "text_order_title_verify"
	TRANSLATE_MSG_ORDER_TITLE_CARD_TO_CARD_FROM                TranslateMsg = "text_order_title_card_to_card_from"
	TRANSLATE_MSG_ORDER_TITLE_CARD_TO_CARD_TO                  TranslateMsg = "text_order_title_card_to_card_to"
	TRANSLATE_MSG_ORDER_TITLE_WITHDRAW                         TranslateMsg = "text_order_title_withdraw"
	TRANSLATE_MSG_ORDER_TITLE_REFUND                           TranslateMsg = "text_order_title_refund"
	TRANSLATE_MSG_ORDER_TITLE_PAY                              TranslateMsg = "text_order_title_pay"
	TRANSLATE_MSG_ORDER_TITLE_CP_EXPRESS                       TranslateMsg = "text_order_title_cp_express"
	TRANSLATE_MSG_ORDER_TITLE_POINT_ACCURAL                    TranslateMsg = "text_order_title_point_accural"
	TRANSLATE_MSG_ORDER_TITLE_POINT_REDEEM                     TranslateMsg = "text_order_title_point_redeem"
	TRANSLATE_MSG_ORDER_TITLE_INTEREST                         TranslateMsg = "text_order_title_interest"
	TRANSLATE_MSG_ORDER_TITLE_INTEREST_AUTO_YIELD              TranslateMsg = "text_order_title_interest_auto_yield"
	TRANSLATE_MSG_ORDER_TITLE_FINANCIAL_TRANSFER_FROM          TranslateMsg = "text_order_title_financial_transfer_from"
	TRANSLATE_MSG_ORDER_TITLE_FINANCIAL_TRANSFER_TO            TranslateMsg = "text_order_title_financial_transfer_to"
	TRANSLATE_MSG_ORDER_TITLE_FINANCIAL_TRANSFER_TO_AUTO_YIELD TranslateMsg = "text_order_title_financial_transfer_to_auto_yield"
	TRANSLATE_MSG_ORDER_TITLE_DISTRIBUTE_APPLY                 TranslateMsg = "text_order_title_distribute_apply"
	TRANSLATE_MSG_ORDER_TITLE_SYSTEM_CHARGE_DECLINE_FINE       TranslateMsg = "text_order_title_system_charge_decline_fine"
	TRANSLATE_MSG_ORDER_TITLE_MERCHANT_ADJUST_BALANCE          TranslateMsg = "text_order_title_merchant_adjust_balance"
	// card
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0302008       TranslateMsg = "text_card_reap_apply_code_0302008"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0302005       TranslateMsg = "text_card_reap_apply_code_0302005"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0302009       TranslateMsg = "text_card_reap_apply_code_0302009"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0302012       TranslateMsg = "text_card_reap_apply_code_0302012"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0321002       TranslateMsg = "text_card_reap_apply_code_0321002"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0302007       TranslateMsg = "text_card_reap_apply_code_0302007"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0315001       TranslateMsg = "text_card_reap_apply_code_0315001"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0302019       TranslateMsg = "text_card_reap_apply_code_0302019"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0309007       TranslateMsg = "text_card_reap_apply_code_0309007"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0401001       TranslateMsg = "text_card_reap_apply_code_0401001"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0321001       TranslateMsg = "text_card_reap_apply_code_0321001"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0302006       TranslateMsg = "text_card_reap_apply_code_0302006"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0309002       TranslateMsg = "text_card_reap_apply_code_0309002"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0302020       TranslateMsg = "text_card_reap_apply_code_0302020"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0302010       TranslateMsg = "text_card_reap_apply_code_0302010"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_INPUT_INVALID TranslateMsg = "text_card_reap_apply_code_input_invalid"
	// ship
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0302014 TranslateMsg = "text_card_reap_ship_code_0302014"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0302015 TranslateMsg = "text_card_reap_ship_code_0302015"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0302024 TranslateMsg = "text_card_reap_ship_code_0302024"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0315003 TranslateMsg = "text_card_reap_ship_code_0315003"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0315002 TranslateMsg = "text_card_reap_ship_code_0315002"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0315005 TranslateMsg = "text_card_reap_ship_code_0315005"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0315004 TranslateMsg = "text_card_reap_ship_code_0315004"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0315006 TranslateMsg = "text_card_reap_ship_code_0315006"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0309001 TranslateMsg = "text_card_reap_ship_code_0309001"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_0302004 TranslateMsg = "text_card_reap_ship_code_0302004"
	TRANSLATE_MSG_CARD_REAP_APPLY_CODE_999999  TranslateMsg = "text_card_reap_ship_code_999999"
	// cp
	TRANSLATE_MSG_CP_FROZEN TranslateMsg = "text_cp_frozen"
	// country
	TRANSLATE_MSG_SYSTEM_COUNTRY_USA TranslateMsg = "text_system_country_usa"
	TRANSLATE_MSG_SYSTEM_COUNTRY_JPN TranslateMsg = "text_system_country_jpn"
	TRANSLATE_MSG_SYSTEM_COUNTRY_AUT TranslateMsg = "text_system_country_aut"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BEL TranslateMsg = "text_system_country_bel"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BGR TranslateMsg = "text_system_country_bgr"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CYP TranslateMsg = "text_system_country_cyp"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CZE TranslateMsg = "text_system_country_cze"
	TRANSLATE_MSG_SYSTEM_COUNTRY_DEU TranslateMsg = "text_system_country_deu"
	TRANSLATE_MSG_SYSTEM_COUNTRY_DNK TranslateMsg = "text_system_country_dnk"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ESP TranslateMsg = "text_system_country_esp"
	TRANSLATE_MSG_SYSTEM_COUNTRY_EST TranslateMsg = "text_system_country_est"
	TRANSLATE_MSG_SYSTEM_COUNTRY_FIN TranslateMsg = "text_system_country_fin"
	TRANSLATE_MSG_SYSTEM_COUNTRY_FRA TranslateMsg = "text_system_country_fra"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GRC TranslateMsg = "text_system_country_grc"
	TRANSLATE_MSG_SYSTEM_COUNTRY_HUN TranslateMsg = "text_system_country_hun"
	TRANSLATE_MSG_SYSTEM_COUNTRY_IRL TranslateMsg = "text_system_country_irl"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ITA TranslateMsg = "text_system_country_ita"
	TRANSLATE_MSG_SYSTEM_COUNTRY_LAT TranslateMsg = "text_system_country_lat"
	TRANSLATE_MSG_SYSTEM_COUNTRY_LIT TranslateMsg = "text_system_country_lit"
	TRANSLATE_MSG_SYSTEM_COUNTRY_LUX TranslateMsg = "text_system_country_lux"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MLT TranslateMsg = "text_system_country_mlt"
	TRANSLATE_MSG_SYSTEM_COUNTRY_NLD TranslateMsg = "text_system_country_nld"
	TRANSLATE_MSG_SYSTEM_COUNTRY_POL TranslateMsg = "text_system_country_pol"
	TRANSLATE_MSG_SYSTEM_COUNTRY_PRT TranslateMsg = "text_system_country_prt"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ROU TranslateMsg = "text_system_country_rou"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SVK TranslateMsg = "text_system_country_svk"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SVN TranslateMsg = "text_system_country_svn"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SWE TranslateMsg = "text_system_country_swe"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GBR TranslateMsg = "text_system_country_gbr"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CHN TranslateMsg = "text_system_country_chn"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CAN TranslateMsg = "text_system_country_can"
	TRANSLATE_MSG_SYSTEM_COUNTRY_AUS TranslateMsg = "text_system_country_aus"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CHE TranslateMsg = "text_system_country_che"
	TRANSLATE_MSG_SYSTEM_COUNTRY_HKG TranslateMsg = "text_system_country_hkg"
	TRANSLATE_MSG_SYSTEM_COUNTRY_NZL TranslateMsg = "text_system_country_nzl"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MEX TranslateMsg = "text_system_country_mex"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SGP TranslateMsg = "text_system_country_sgp"
	TRANSLATE_MSG_SYSTEM_COUNTRY_NOR TranslateMsg = "text_system_country_nor"
	TRANSLATE_MSG_SYSTEM_COUNTRY_KOR TranslateMsg = "text_system_country_kor"
	TRANSLATE_MSG_SYSTEM_COUNTRY_TUR TranslateMsg = "text_system_country_tur"
	TRANSLATE_MSG_SYSTEM_COUNTRY_IND TranslateMsg = "text_system_country_ind"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BRA TranslateMsg = "text_system_country_bra"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ZAF TranslateMsg = "text_system_country_zaf"
	TRANSLATE_MSG_SYSTEM_COUNTRY_RUS TranslateMsg = "text_system_country_rus"
	TRANSLATE_MSG_SYSTEM_COUNTRY_TWN TranslateMsg = "text_system_country_twn"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ISR TranslateMsg = "text_system_country_isr"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ARE TranslateMsg = "text_system_country_are"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ARG TranslateMsg = "text_system_country_arg"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MYS TranslateMsg = "text_system_country_mys"
	TRANSLATE_MSG_SYSTEM_COUNTRY_THA TranslateMsg = "text_system_country_tha"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SAU TranslateMsg = "text_system_country_sau"
	TRANSLATE_MSG_SYSTEM_COUNTRY_IDN TranslateMsg = "text_system_country_idn"
	TRANSLATE_MSG_SYSTEM_COUNTRY_PAK TranslateMsg = "text_system_country_pak"
	TRANSLATE_MSG_SYSTEM_COUNTRY_NGA TranslateMsg = "text_system_country_nga"
	TRANSLATE_MSG_SYSTEM_COUNTRY_EGY TranslateMsg = "text_system_country_egy"
	TRANSLATE_MSG_SYSTEM_COUNTRY_VNM TranslateMsg = "text_system_country_vnm"
	TRANSLATE_MSG_SYSTEM_COUNTRY_PHL TranslateMsg = "text_system_country_phl"
	TRANSLATE_MSG_SYSTEM_COUNTRY_COL TranslateMsg = "text_system_country_col"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CHL TranslateMsg = "text_system_country_chl"
	TRANSLATE_MSG_SYSTEM_COUNTRY_PER TranslateMsg = "text_system_country_per"
	TRANSLATE_MSG_SYSTEM_COUNTRY_DZA TranslateMsg = "text_system_country_dza"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MAR TranslateMsg = "text_system_country_mar"
	TRANSLATE_MSG_SYSTEM_COUNTRY_KEN TranslateMsg = "text_system_country_ken"
	TRANSLATE_MSG_SYSTEM_COUNTRY_UKR TranslateMsg = "text_system_country_ukr"
	TRANSLATE_MSG_SYSTEM_COUNTRY_AFG TranslateMsg = "text_system_country_afg"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ETH TranslateMsg = "text_system_country_eth"
	TRANSLATE_MSG_SYSTEM_COUNTRY_COD TranslateMsg = "text_system_country_cod"
	TRANSLATE_MSG_SYSTEM_COUNTRY_TZA TranslateMsg = "text_system_country_tza"
	TRANSLATE_MSG_SYSTEM_COUNTRY_UZB TranslateMsg = "text_system_country_uzb"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ALA TranslateMsg = "text_system_country_ala"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ALB TranslateMsg = "text_system_country_alb"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ASM TranslateMsg = "text_system_country_asm"
	TRANSLATE_MSG_SYSTEM_COUNTRY_AND TranslateMsg = "text_system_country_and"
	TRANSLATE_MSG_SYSTEM_COUNTRY_AGO TranslateMsg = "text_system_country_ago"
	TRANSLATE_MSG_SYSTEM_COUNTRY_AIA TranslateMsg = "text_system_country_aia"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ATA TranslateMsg = "text_system_country_ata"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ATG TranslateMsg = "text_system_country_atg"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ARM TranslateMsg = "text_system_country_arm"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ABW TranslateMsg = "text_system_country_abw"
	TRANSLATE_MSG_SYSTEM_COUNTRY_AZE TranslateMsg = "text_system_country_aze"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BHS TranslateMsg = "text_system_country_bhs"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BHR TranslateMsg = "text_system_country_bhr"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BGD TranslateMsg = "text_system_country_bgd"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BRB TranslateMsg = "text_system_country_brb"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BLR TranslateMsg = "text_system_country_blr"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BLZ TranslateMsg = "text_system_country_blz"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BEN TranslateMsg = "text_system_country_ben"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BMU TranslateMsg = "text_system_country_bmu"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BTN TranslateMsg = "text_system_country_btn"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BOL TranslateMsg = "text_system_country_bol"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BES TranslateMsg = "text_system_country_bes"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BIH TranslateMsg = "text_system_country_bih"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BWA TranslateMsg = "text_system_country_bwa"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BVT TranslateMsg = "text_system_country_bvt"
	TRANSLATE_MSG_SYSTEM_COUNTRY_IOT TranslateMsg = "text_system_country_iot"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BRN TranslateMsg = "text_system_country_brn"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BFA TranslateMsg = "text_system_country_bfa"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BDI TranslateMsg = "text_system_country_bdi"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CPV TranslateMsg = "text_system_country_cpv"
	TRANSLATE_MSG_SYSTEM_COUNTRY_KHM TranslateMsg = "text_system_country_khm"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CMR TranslateMsg = "text_system_country_cmr"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CYM TranslateMsg = "text_system_country_cym"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CAF TranslateMsg = "text_system_country_caf"
	TRANSLATE_MSG_SYSTEM_COUNTRY_TCD TranslateMsg = "text_system_country_tcd"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CXR TranslateMsg = "text_system_country_cxr"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CCK TranslateMsg = "text_system_country_cck"
	TRANSLATE_MSG_SYSTEM_COUNTRY_COM TranslateMsg = "text_system_country_com"
	TRANSLATE_MSG_SYSTEM_COUNTRY_COG TranslateMsg = "text_system_country_cog"
	TRANSLATE_MSG_SYSTEM_COUNTRY_COK TranslateMsg = "text_system_country_cok"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CRI TranslateMsg = "text_system_country_cri"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CIV TranslateMsg = "text_system_country_civ"
	TRANSLATE_MSG_SYSTEM_COUNTRY_HRV TranslateMsg = "text_system_country_hrv"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CUB TranslateMsg = "text_system_country_cub"
	TRANSLATE_MSG_SYSTEM_COUNTRY_CUW TranslateMsg = "text_system_country_cuw"
	TRANSLATE_MSG_SYSTEM_COUNTRY_DJI TranslateMsg = "text_system_country_dji"
	TRANSLATE_MSG_SYSTEM_COUNTRY_DMA TranslateMsg = "text_system_country_dma"
	TRANSLATE_MSG_SYSTEM_COUNTRY_DOM TranslateMsg = "text_system_country_dom"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ECU TranslateMsg = "text_system_country_ecu"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SLV TranslateMsg = "text_system_country_slv"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GNQ TranslateMsg = "text_system_country_gnq"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ERI TranslateMsg = "text_system_country_eri"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SWZ TranslateMsg = "text_system_country_swz"
	TRANSLATE_MSG_SYSTEM_COUNTRY_FLK TranslateMsg = "text_system_country_flk"
	TRANSLATE_MSG_SYSTEM_COUNTRY_FRO TranslateMsg = "text_system_country_fro"
	TRANSLATE_MSG_SYSTEM_COUNTRY_FJI TranslateMsg = "text_system_country_fji"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GUF TranslateMsg = "text_system_country_guf"
	TRANSLATE_MSG_SYSTEM_COUNTRY_PYF TranslateMsg = "text_system_country_pyf"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ATF TranslateMsg = "text_system_country_atf"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GAB TranslateMsg = "text_system_country_gab"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GMB TranslateMsg = "text_system_country_gmb"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GEO TranslateMsg = "text_system_country_geo"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GHA TranslateMsg = "text_system_country_gha"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GIB TranslateMsg = "text_system_country_gib"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GRL TranslateMsg = "text_system_country_grl"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GRD TranslateMsg = "text_system_country_grd"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GLP TranslateMsg = "text_system_country_glp"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GUM TranslateMsg = "text_system_country_gum"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GTM TranslateMsg = "text_system_country_gtm"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GGY TranslateMsg = "text_system_country_ggy"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GIN TranslateMsg = "text_system_country_gin"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GNB TranslateMsg = "text_system_country_gnb"
	TRANSLATE_MSG_SYSTEM_COUNTRY_GUY TranslateMsg = "text_system_country_guy"
	TRANSLATE_MSG_SYSTEM_COUNTRY_HTI TranslateMsg = "text_system_country_hti"
	TRANSLATE_MSG_SYSTEM_COUNTRY_HMD TranslateMsg = "text_system_country_hmd"
	TRANSLATE_MSG_SYSTEM_COUNTRY_VAT TranslateMsg = "text_system_country_vat"
	TRANSLATE_MSG_SYSTEM_COUNTRY_HND TranslateMsg = "text_system_country_hnd"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ISL TranslateMsg = "text_system_country_isl"
	TRANSLATE_MSG_SYSTEM_COUNTRY_IRN TranslateMsg = "text_system_country_irn"
	TRANSLATE_MSG_SYSTEM_COUNTRY_IRQ TranslateMsg = "text_system_country_irq"
	TRANSLATE_MSG_SYSTEM_COUNTRY_IMN TranslateMsg = "text_system_country_imn"
	TRANSLATE_MSG_SYSTEM_COUNTRY_JAM TranslateMsg = "text_system_country_jam"
	TRANSLATE_MSG_SYSTEM_COUNTRY_JEY TranslateMsg = "text_system_country_jey"
	TRANSLATE_MSG_SYSTEM_COUNTRY_JOR TranslateMsg = "text_system_country_jor"
	TRANSLATE_MSG_SYSTEM_COUNTRY_KAZ TranslateMsg = "text_system_country_kaz"
	TRANSLATE_MSG_SYSTEM_COUNTRY_KIR TranslateMsg = "text_system_country_kir"
	TRANSLATE_MSG_SYSTEM_COUNTRY_PRK TranslateMsg = "text_system_country_prk"
	TRANSLATE_MSG_SYSTEM_COUNTRY_KWT TranslateMsg = "text_system_country_kwt"
	TRANSLATE_MSG_SYSTEM_COUNTRY_KGZ TranslateMsg = "text_system_country_kgz"
	TRANSLATE_MSG_SYSTEM_COUNTRY_LAO TranslateMsg = "text_system_country_lao"
	TRANSLATE_MSG_SYSTEM_COUNTRY_LVA TranslateMsg = "text_system_country_lva"
	TRANSLATE_MSG_SYSTEM_COUNTRY_LBN TranslateMsg = "text_system_country_lbn"
	TRANSLATE_MSG_SYSTEM_COUNTRY_LSO TranslateMsg = "text_system_country_lso"
	TRANSLATE_MSG_SYSTEM_COUNTRY_LBR TranslateMsg = "text_system_country_lbr"
	TRANSLATE_MSG_SYSTEM_COUNTRY_LBY TranslateMsg = "text_system_country_lby"
	TRANSLATE_MSG_SYSTEM_COUNTRY_LIE TranslateMsg = "text_system_country_lie"
	TRANSLATE_MSG_SYSTEM_COUNTRY_LTU TranslateMsg = "text_system_country_ltu"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MAC TranslateMsg = "text_system_country_mac"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MDG TranslateMsg = "text_system_country_mdg"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MWI TranslateMsg = "text_system_country_mwi"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MDV TranslateMsg = "text_system_country_mdv"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MLI TranslateMsg = "text_system_country_mli"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MHL TranslateMsg = "text_system_country_mhl"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MTQ TranslateMsg = "text_system_country_mtq"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MRT TranslateMsg = "text_system_country_mrt"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MUS TranslateMsg = "text_system_country_mus"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MYT TranslateMsg = "text_system_country_myt"
	TRANSLATE_MSG_SYSTEM_COUNTRY_FSM TranslateMsg = "text_system_country_fsm"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MDA TranslateMsg = "text_system_country_mda"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MCO TranslateMsg = "text_system_country_mco"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MNG TranslateMsg = "text_system_country_mng"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MNE TranslateMsg = "text_system_country_mne"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MSR TranslateMsg = "text_system_country_msr"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MOZ TranslateMsg = "text_system_country_moz"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MMR TranslateMsg = "text_system_country_mmr"
	TRANSLATE_MSG_SYSTEM_COUNTRY_NAM TranslateMsg = "text_system_country_nam"
	TRANSLATE_MSG_SYSTEM_COUNTRY_NRU TranslateMsg = "text_system_country_nru"
	TRANSLATE_MSG_SYSTEM_COUNTRY_NPL TranslateMsg = "text_system_country_npl"
	TRANSLATE_MSG_SYSTEM_COUNTRY_NCL TranslateMsg = "text_system_country_ncl"
	TRANSLATE_MSG_SYSTEM_COUNTRY_NIC TranslateMsg = "text_system_country_nic"
	TRANSLATE_MSG_SYSTEM_COUNTRY_NER TranslateMsg = "text_system_country_ner"
	TRANSLATE_MSG_SYSTEM_COUNTRY_NIU TranslateMsg = "text_system_country_niu"
	TRANSLATE_MSG_SYSTEM_COUNTRY_NFK TranslateMsg = "text_system_country_nfk"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MKD TranslateMsg = "text_system_country_mkd"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MNP TranslateMsg = "text_system_country_mnp"
	TRANSLATE_MSG_SYSTEM_COUNTRY_OMN TranslateMsg = "text_system_country_omn"
	TRANSLATE_MSG_SYSTEM_COUNTRY_PLW TranslateMsg = "text_system_country_plw"
	TRANSLATE_MSG_SYSTEM_COUNTRY_PSE TranslateMsg = "text_system_country_pse"
	TRANSLATE_MSG_SYSTEM_COUNTRY_PAN TranslateMsg = "text_system_country_pan"
	TRANSLATE_MSG_SYSTEM_COUNTRY_PNG TranslateMsg = "text_system_country_png"
	TRANSLATE_MSG_SYSTEM_COUNTRY_PRY TranslateMsg = "text_system_country_pry"
	TRANSLATE_MSG_SYSTEM_COUNTRY_PCN TranslateMsg = "text_system_country_pcn"
	TRANSLATE_MSG_SYSTEM_COUNTRY_PRI TranslateMsg = "text_system_country_pri"
	TRANSLATE_MSG_SYSTEM_COUNTRY_QAT TranslateMsg = "text_system_country_qat"
	TRANSLATE_MSG_SYSTEM_COUNTRY_REU TranslateMsg = "text_system_country_reu"
	TRANSLATE_MSG_SYSTEM_COUNTRY_RWA TranslateMsg = "text_system_country_rwa"
	TRANSLATE_MSG_SYSTEM_COUNTRY_BLM TranslateMsg = "text_system_country_blm"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SHN TranslateMsg = "text_system_country_shn"
	TRANSLATE_MSG_SYSTEM_COUNTRY_KNA TranslateMsg = "text_system_country_kna"
	TRANSLATE_MSG_SYSTEM_COUNTRY_LCA TranslateMsg = "text_system_country_lca"
	TRANSLATE_MSG_SYSTEM_COUNTRY_MAF TranslateMsg = "text_system_country_maf"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SPM TranslateMsg = "text_system_country_spm"
	TRANSLATE_MSG_SYSTEM_COUNTRY_VCT TranslateMsg = "text_system_country_vct"
	TRANSLATE_MSG_SYSTEM_COUNTRY_WSM TranslateMsg = "text_system_country_wsm"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SMR TranslateMsg = "text_system_country_smr"
	TRANSLATE_MSG_SYSTEM_COUNTRY_STP TranslateMsg = "text_system_country_stp"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SEN TranslateMsg = "text_system_country_sen"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SRB TranslateMsg = "text_system_country_srb"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SYC TranslateMsg = "text_system_country_syc"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SLE TranslateMsg = "text_system_country_sle"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SXM TranslateMsg = "text_system_country_sxm"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SLB TranslateMsg = "text_system_country_slb"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SOM TranslateMsg = "text_system_country_som"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SGS TranslateMsg = "text_system_country_sgs"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SSD TranslateMsg = "text_system_country_ssd"
	TRANSLATE_MSG_SYSTEM_COUNTRY_LKA TranslateMsg = "text_system_country_lka"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SDN TranslateMsg = "text_system_country_sdn"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SUR TranslateMsg = "text_system_country_sur"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SJM TranslateMsg = "text_system_country_sjm"
	TRANSLATE_MSG_SYSTEM_COUNTRY_SYR TranslateMsg = "text_system_country_syr"
	TRANSLATE_MSG_SYSTEM_COUNTRY_TJK TranslateMsg = "text_system_country_tjk"
	TRANSLATE_MSG_SYSTEM_COUNTRY_TLS TranslateMsg = "text_system_country_tls"
	TRANSLATE_MSG_SYSTEM_COUNTRY_TGO TranslateMsg = "text_system_country_tgo"
	TRANSLATE_MSG_SYSTEM_COUNTRY_TKL TranslateMsg = "text_system_country_tkl"
	TRANSLATE_MSG_SYSTEM_COUNTRY_TON TranslateMsg = "text_system_country_ton"
	TRANSLATE_MSG_SYSTEM_COUNTRY_TTO TranslateMsg = "text_system_country_tto"
	TRANSLATE_MSG_SYSTEM_COUNTRY_TUN TranslateMsg = "text_system_country_tun"
	TRANSLATE_MSG_SYSTEM_COUNTRY_TKM TranslateMsg = "text_system_country_tkm"
	TRANSLATE_MSG_SYSTEM_COUNTRY_TCA TranslateMsg = "text_system_country_tca"
	TRANSLATE_MSG_SYSTEM_COUNTRY_TUV TranslateMsg = "text_system_country_tuv"
	TRANSLATE_MSG_SYSTEM_COUNTRY_UGA TranslateMsg = "text_system_country_uga"
	TRANSLATE_MSG_SYSTEM_COUNTRY_UMI TranslateMsg = "text_system_country_umi"
	TRANSLATE_MSG_SYSTEM_COUNTRY_URY TranslateMsg = "text_system_country_ury"
	TRANSLATE_MSG_SYSTEM_COUNTRY_VUT TranslateMsg = "text_system_country_vut"
	TRANSLATE_MSG_SYSTEM_COUNTRY_VEN TranslateMsg = "text_system_country_ven"
	TRANSLATE_MSG_SYSTEM_COUNTRY_VGB TranslateMsg = "text_system_country_vgb"
	TRANSLATE_MSG_SYSTEM_COUNTRY_VIR TranslateMsg = "text_system_country_vir"
	TRANSLATE_MSG_SYSTEM_COUNTRY_WLF TranslateMsg = "text_system_country_wlf"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ESH TranslateMsg = "text_system_country_esh"
	TRANSLATE_MSG_SYSTEM_COUNTRY_YEM TranslateMsg = "text_system_country_yem"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ZMB TranslateMsg = "text_system_country_zmb"
	TRANSLATE_MSG_SYSTEM_COUNTRY_ZWE TranslateMsg = "text_system_country_zwe"
)

type UnitType int // 單位類型
const (
	UNIT_TYPE_DOLLAR      UnitType = 100
	UNIT_TYPE_NANOSECOND  UnitType = 200
	UNIT_TYPE_MICROSECOND UnitType = 201
	UNIT_TYPE_MILLISECOND UnitType = 202
	UNIT_TYPE_SECOND      UnitType = 203
	UNIT_TYPE_MINUTE      UnitType = 204
	UNIT_TYPE_HOUR        UnitType = 205
	UNIT_TYPE_DAY         UnitType = 206
	UNIT_TYPE_WEEK        UnitType = 207
	UNIT_TYPE_MONTH       UnitType = 208
	UNIT_TYPE_YEAR        UnitType = 209
)

type WhaleAdjustType int // whale調整餘額類型
const (
	WHALE_ADJUST_TYPE_BALANCE_TYPE_CREDIT WhaleAdjustType = 1
	WHALE_ADJUST_TYPE_BALANCE_TYPE_DEDUCT WhaleAdjustType = 2
)

type WhaleAdjustBalancePurpose int // whale調整餘額目的
const (
	WHALE_ADJUST_TYPE_BALANCE_PURPOSE_APPLY         WhaleAdjustBalancePurpose = 1
	WHALE_ADJUST_TYPE_BALANCE_PURPOSE_TOP_UP        WhaleAdjustBalancePurpose = 2
	WHALE_ADJUST_TYPE_BALANCE_PURPOSE_TOP_DOWN      WhaleAdjustBalancePurpose = 3
	WHALE_ADJUST_TYPE_BALANCE_PURPOSE_TRANSFER_FROM WhaleAdjustBalancePurpose = 4
	WHALE_ADJUST_TYPE_BALANCE_PURPOSE_TRANSFER_TO   WhaleAdjustBalancePurpose = 5
	WHALE_ADJUST_TYPE_BALANCE_PURPOSE_UNKNOWN       WhaleAdjustBalancePurpose = 127
)

type WhaleAdjustBalanceOrderStatus int // whale調整餘額狀態
const (
	WHALE_ADJUST_BALANCE_ORDER_STATUS_REQUESTED  WhaleAdjustBalanceOrderStatus = 1
	WHALE_ADJUST_BALANCE_ORDER_STATUS_PROCESSING WhaleAdjustBalanceOrderStatus = 2
	WHALE_ADJUST_BALANCE_ORDER_STATUS_SUCCESS    WhaleAdjustBalanceOrderStatus = 3
	WHALE_ADJUST_BALANCE_ORDER_STATUS_FAILED     WhaleAdjustBalanceOrderStatus = 4
)

type WhaleBusinessType string // whale業務類型
const (
	WHALE_BUSINESS_TYPE_CREATE_CARD                 WhaleBusinessType = "CreateCard"
	WHALE_BUSINESS_TYPE_CARD_FROZEN                 WhaleBusinessType = "CardFrozen"
	WHALE_BUSINESS_TYPE_CARD_RISKY                  WhaleBusinessType = "CardRisky"
	WHALE_BUSINESS_TYPE_CARD_FROZEN_BY_SYSTEM       WhaleBusinessType = "CardFrozenBySystem"
	WHALE_BUSINESS_TYPE_CARD_DELETED                WhaleBusinessType = "CardDeleted"
	WHALE_BUSINESS_TYPE_CARD_ACTIVATED              WhaleBusinessType = "CardActivated"
	WHALE_BUSINESS_TYPE_CARD_TRANSACTION            WhaleBusinessType = "CardTransaction"
	WHALE_BUSINESS_TYPE_BALANCE_DEPOSIT_TRANSACTION WhaleBusinessType = "BalanceDepositTransaction"
	WHALE_BUSINESS_TYPE_BALANCE_EDIT                WhaleBusinessType = "BalanceEdit"
	WHALE_BUSINESS_TYPE_CARD_3DS_OTP                WhaleBusinessType = "Card3dsOtp"
)

type WhaleCardType int // whale卡片類型
const (
	WHALE_CARD_TYPE_VISA             WhaleCardType = 1
	WHALE_CARD_TYPE_VISA_PREMIUM     WhaleCardType = 2
	WHALE_CARD_TYPE_VISA_PREMIUM_USD WhaleCardType = 3
	WHALE_CARD_TYPE_VISA_PREMIUM_HKD WhaleCardType = 4
	WHALE_CARD_TYPE_MASTER           WhaleCardType = 5
)

type WhaleCardInfoStatus string // whale卡片資料狀態
const (
	WHALE_CARD_INFO_STATUS_ACTIVATED        WhaleCardInfoStatus = "activated"
	WHALE_CARD_INFO_STATUS_INACTIVE         WhaleCardInfoStatus = "inactive"
	WHALE_CARD_INFO_STATUS_RISKY            WhaleCardInfoStatus = "risky"
	WHALE_CARD_INFO_STATUS_FROZEN           WhaleCardInfoStatus = "Frozen"
	WHALE_CARD_INFO_STATUS_ACTIVE           WhaleCardInfoStatus = "Active"
	WHALE_CARD_INFO_STATUS_FROZEN_BY_SYSTEM WhaleCardInfoStatus = "frozen_by_system"
	WHALE_CARD_INFO_STATUS_DELETED          WhaleCardInfoStatus = "Deleted"
)

type WhaleCardStatus string // whale卡片狀態
const (
	WHALE_CARD_STATUS_ACTIVATED        WhaleCardStatus = "activated"
	WHALE_CARD_STATUS_INACTIVE         WhaleCardStatus = "inactive"
	WHALE_CARD_STATUS_RISKY            WhaleCardStatus = "risky"
	WHALE_CARD_STATUS_FROZEN           WhaleCardStatus = "Frozen"
	WHALE_CARD_STATUS_ACTIVE           WhaleCardStatus = "Active"
	WHALE_CARD_STATUS_FROZEN_BY_SYSTEM WhaleCardStatus = "frozen_by_system"
	WHALE_CARD_STATUS_DELETED          WhaleCardStatus = "Deleted"
)

type WhaleCardTransactionType string // whale卡片交易類型
const (
	WHALE_CARD_TRANSACTION_TYPE_CONSUMPTION WhaleCardTransactionType = "Consumption"
	WHALE_CARD_TRANSACTION_TYPE_CREDIT      WhaleCardTransactionType = "Credit"
	WHALE_CARD_TRANSACTION_TYPE_REVERSAL    WhaleCardTransactionType = "Reversal"
	WHALE_CARD_TRANSACTION_TYPE_FROZEN      WhaleCardTransactionType = "Frozen"
	WHALE_CARD_TRANSACTION_TYPE_UNFROZEN    WhaleCardTransactionType = "UnFrozen"
)

type WhaleCardTransactionStatus string // whale卡片交易狀態
const (
	WHALE_CARD_TRANSACTION_STATUS_PENDING WhaleCardTransactionStatus = "Pending"
	WHALE_CARD_TRANSACTION_STATUS_CLOSED  WhaleCardTransactionStatus = "Closed"
	WHALE_CARD_TRANSACTION_STATUS_FAIL    WhaleCardTransactionStatus = "Fail"
)

type WhaleCardTransferOrderStatus int // whale轉卡訂單狀態
const (
	WHALE_CARD_TRANSFER_ORDER_STATUS_PROCESSING WhaleCardTransferOrderStatus = 1
	WHALE_CARD_TRANSFER_ORDER_STATUS_SUCCESS    WhaleCardTransferOrderStatus = 2
	WHALE_CARD_TRANSFER_ORDER_STATUS_FAILED     WhaleCardTransferOrderStatus = 3
)

type WhaleTransferStatus int // whale轉卡狀態
const (
	WHALE_TRANSFER_STATUS_FAILED     WhaleTransferStatus = -1
	WHALE_TRANSFER_STATUS_PROCESSING WhaleTransferStatus = 1
	WHALE_TRANSFER_STATUS_SUCCESS    WhaleTransferStatus = 2
)

type WhaleEditBalanceType string // whale調整餘額類型
const (
	WHALE_EDIT_BALANCE_TYPE_CREDIT WhaleEditBalanceType = "credit"
	WHALE_EDIT_BALANCE_TYPE_DEDUCT WhaleEditBalanceType = "deduct"
)

type WhaleEditBalanceStatus int // whale更新餘額狀態
const (
	WHALE_EDIT_BALANCE_STATUS_UNKNOWN    WhaleEditBalanceStatus = -1
	WHALE_EDIT_BALANCE_STATUS_PROCESSING WhaleEditBalanceStatus = 0
	WHALE_EDIT_BALANCE_STATUS_SUCCESS    WhaleEditBalanceStatus = 1
	WHALE_EDIT_BALANCE_STATUS_FAILED     WhaleEditBalanceStatus = 2
)

type WhalePayStatus int // whale支付狀態
const (
	WHALE_PAY_STATUS_NONE     WhalePayStatus = 0
	WHALE_PAY_STATUS_PENDING  WhalePayStatus = 1
	WHALE_PAY_STATUS_CLOSED   WhalePayStatus = 2
	WHALE_PAY_STATUS_FAILED   WhalePayStatus = 3
	WHALE_PAY_STATUS_CREDIT   WhalePayStatus = 4
	WHALE_PAY_STATUS_REVERSAL WhalePayStatus = 5
)

type WhaleTransactionType int // whale交易類型
const (
	WHALE_TRANSACTION_TYPE_CONSUMPTION WhaleTransactionType = 1
	WHALE_TRANSACTION_TYPE_CREDIT      WhaleTransactionType = 2
	WHALE_TRANSACTION_TYPE_REVERSAL    WhaleTransactionType = 3
)

type WithdrawAutoLevel int // 提幣自動等級
const (
	WITHDRAW_AUTO_LEVEL_NONE         WithdrawAutoLevel = 0
	WITHDRAW_AUTO_LEVEL_AUTO_REVIEW  WithdrawAutoLevel = 1
	WITHDRAW_AUTO_LEVEL_AUTO_APPROVE WithdrawAutoLevel = 2
)

type WithdrawScene int // 提領場景
const (
	WITHDRAW_SCENE_NO_REVIEW      WithdrawScene = 1
	WITHDRAW_SCENE_ADMIN_REVIEWED WithdrawScene = 2
	WITHDRAW_SCENE_ADMIN_APPROVED WithdrawScene = 3
	WITHDRAW_SCENE_AUTO_REVIEWED  WithdrawScene = 4
	WITHDRAW_SCENE_AUTO_APPROVED  WithdrawScene = 5
)

type WithdrawStatus int // 提領狀態
const (
	WITHDRAW_STATUS_PENDING           WithdrawStatus = 1
	WITHDRAW_STATUS_FAILED            WithdrawStatus = 2
	WITHDRAW_STATUS_SUCCESS           WithdrawStatus = 3
	WITHDRAW_STATUS_FREEZED           WithdrawStatus = 4
	WITHDRAW_STATUS_COINSDO_PENDING   WithdrawStatus = 5
	WITHDRAW_STATUS_COINSDO_PASSED    WithdrawStatus = 6
	WITHDRAW_STATUS_COINSDO_REJECTED  WithdrawStatus = 7
	WITHDRAW_STATUS_COINSDO_CANCELLED WithdrawStatus = 8
	WITHDRAW_STATUS_COINSDO_FAILED    WithdrawStatus = 9
)

type WorkerPool int // 工作池
const (
	WORKER_POOL_DEFAULT          WorkerPool = 1
	WORKER_POOL_OPEN_REQUEST_LOG WorkerPool = 2
)

type AddressPoolStatus int // 參數狀態
const (
	ADDRESS_POOL_AVAILABLE   AddressPoolStatus = 1
	ADDRESS_POOL_UNAVAILABLE AddressPoolStatus = 2
)

type CryptoCurrencyStatus int // crypto_currency 狀態:1啟用 2關閉
const (
	CRYPTO_CURRENCY_STATUS_ON  CryptoCurrencyStatus = 1
	CRYPTO_CURRENCY_STATUS_OFF CryptoCurrencyStatus = 2
)

type PayChannel string

const (
	APP_STORE PayChannel = "APP_STORE"
	EUSD      PayChannel = "EUSD"
)

type UserLimitsType int // 用戶限流類型
const (
	CARD_LIMIT         UserLimitsType = 1
	TRANSFER_LIMIT     UserLimitsType = 2
	CP_EXPRESS_LIMIT   UserLimitsType = 3
	ATM_WITHDRAW_LIMIT UserLimitsType = 4
)

type PasskeyVerifyStatus int // passkey verify 狀態
const (
	PASSKEY_VERIFY_STATUS_PENDING PasskeyVerifyStatus = 1
	PASSKEY_VERIFY_STATUS_SUCCESS PasskeyVerifyStatus = 2
	PASSKEY_VERIFY_STATUS_FAILED  PasskeyVerifyStatus = 3
)

type AppsCategory int // 系統 app 類別
const (
	APP_ALL                             = 0
	APP_CATEGORY_AI_TOOL   AppsCategory = 1
	APP_CATEGORY_EXCHANGES AppsCategory = 2
	APP_CAREGORY_GAME      AppsCategory = 3
)

type MsgOPCode string // 訊息行為碼
const (
	MSG_OPCODE_READ                             MsgOPCode = "READ"
	MSG_OPCODE_FORWARD_3DS                      MsgOPCode = "3DS"
	MSG_OPCODE_FORWARD_3DS_BALANCED             MsgOPCode = "3DS_BALANCED"
	MSG_OPCODE_OTP                              MsgOPCode = "OTP"
	MSG_OPCODE_LOGIN_NOTIFY                     MsgOPCode = "LOGIN_NOTIFY"
	MSG_OPCODE_CARD_TRANSACTION                 MsgOPCode = "CARD_TRANSACTION"
	MSG_OPCODE_RECHARGE_SUCCESSFUL              MsgOPCode = "RECHARGE_SUCCESSFUL"
	MSG_OPCODE_CARD_FROZEN                      MsgOPCode = "CARD_FROZEN"
	MSG_OPCODE_CARD_UNFROZEN                    MsgOPCode = "CARD_UNFROZEN"
	MSG_OPCODE_CARD_DELETE                      MsgOPCode = "CARD_DELETE"
	MSG_OPCODE_KYC_PASSED                       MsgOPCode = "KYC_PASSED"
	MSG_OPCODE_STATEMENT                        MsgOPCode = "STATEMENT"
	MSG_OPCODE_WITHDRAW_REQUEST                 MsgOPCode = "WITHDRAW_REQUEST"
	MSG_OPCODE_WITHDRAW_SUCCESS                 MsgOPCode = "WITHDRAW_SUCCESS"
	MSG_OPCODE_REQUEST_3DS                      MsgOPCode = "REQUEST_3DS"
	MSG_OPCODE_REAP_WALLET_BIND_OTP             MsgOPCode = "REAP_WALLET_BIND_OTP"
	MSG_OPCODE_REAP_TRANSACTION_FAIL            MsgOPCode = "REAP_TRANSACTION_FAIL"
	MSG_OPCODE_EXTRENAL_REQUEST_LOG             MsgOPCode = "EXTRENAL_REQUEST_LOG"
	MSG_OPCODE_EXTRENAL_REQUEST_LOG_PARTITIONED MsgOPCode = "EXTRENAL_REQUEST_LOG_PARTITIONED"
	MSG_OPCODE_DAPP_APPLY_CARD_SUCCESS          MsgOPCode = "DAPP_APPLY_CARD_SUCESS"
	MSG_OPCODE_ADMIN_NOTICE                     MsgOPCode = "ADMIN_NOTICE"
	MSG_OPCODE_REQUEST_WHALE_3DS                MsgOPCode = "REQUEST_WHALE_3DS"
	MSG_OPCODE_WHALE_CARD_RISKY                 MsgOPCode = "WHALE_CARD_RISKY"
	MSG_OPCODE_USER_BLOCKED                     MsgOPCode = "USER_BLOCKED"
	MSG_OPCODE_KYC_NOTIFY                       MsgOPCode = "KYC_NOTIFY"
	MSG_OPCODE_USER_UNBLOCKED                   MsgOPCode = "USER_UNBLOCKED"
	MSG_OPCODE_KYC_NOTIFY_FINAL                 MsgOPCode = "KYC_NOTIFY_FINAL"
	MSG_OPCODE_CARD_BLOCKED                     MsgOPCode = "CARD_BLOCKED"
	MSG_OPCODE_CARD_UNBLOCKED                   MsgOPCode = "CARD_UNBLOCKED"
	MSG_OPCODE_MERCHANT_EXPORT_BALANCE_CHANGE   MsgOPCode = "MERCHANT_EXPORT_BALANCE_CHANGE"
	MSG_OPCODE_USER_DELETE                      MsgOPCode = "USER_DELETE"
	MSG_OPCODE_CUSTOMISED                       MsgOPCode = "CUSTOMISED"
	MSG_OPCODE_CARD_BLOCKED_CONSECUTIVE_FAILURE MsgOPCode = "CARD_BLOCKED_CONSECUTIVE_FAILURE"
	MSG_OPCODE_CARD_REINSTATE_AFTER_BLOCKED     MsgOPCode = "CARD_REINSTATE_AFTER_BLOCKED"
	MSG_OPCODE_DELIVERY_CARD_IN_PROGRESS        MsgOPCode = "DELIVERY_CARD_IN_PROGRESS"
	MSG_OPCODE_DELIVERY_CARD_SENT_TO_USER       MsgOPCode = "DELIVERY_CARD_SENT_TO_USER"
	MSG_OPCODE_CARD_TRANSACTION_OTP             MsgOPCode = "CARD_TRANSACTION_OTP"
	MSG_OPCODE_APPLY_CARD_SUCCESS               MsgOPCode = "APPLY_CARD_SUCCESS"
	MSG_OPCODE_APPLY_CARD_FAIL                  MsgOPCode = "APPLY_CARD_FAIL"
	MSG_OPCODE_CRYPTO_CARD_BLOCKED              MsgOPCode = "CRYPTO_CARD_BLOCKED"
	MSG_OPCODE_CRYPTO_CARD_UNBLOCKED            MsgOPCode = "CRYPTO_CARD_UNBLOCKED"
	MSG_OPCODE_CARD_AMOUNT_LOW_WARNING          MsgOPCode = "CARD_AMOUNT_LOW_WARNING"
)

type UserMessageReadStatus int // 參數狀態
const (
	USER_MESSAGE_UNREAD UserMessageReadStatus = 0
	USER_MESSAGE_READ   UserMessageReadStatus = 1
)

type UserMessageType int

const (
	USER_MESSAGE              UserMessageType = 1
	USER_MESSAGE_REQUEST      UserMessageType = 2
	USER_MESSAGE_ADMIN_NOTICE UserMessageType = 3
)

type MailChannel string // 邮件通道
const (
	MAIL_CHANNEL_MAILGUN  MailChannel = "MAILGUN"
	MAIL_CHANNEL_SENDGRID MailChannel = "SENDGRID"
)

type AgentReportType int // 代理報表類型
const (
	AGENT_REPORT_CODE_MONTHLY_TRANSACTION AgentReportType = 1
)

// thirdparty 交易狀態
type ThirdPartyDepositStatus int

const (
	TRANSACTION_STATUS_CREATED  ThirdPartyDepositStatus = 1 // 已創建
	TRANSACTION_STATUS_PENDING  ThirdPartyDepositStatus = 2 // 處理中
	TRANSACTION_STATUS_HOLD     ThirdPartyDepositStatus = 3 // 待審核
	TRANSACTION_STATUS_REFUNDED ThirdPartyDepositStatus = 4 // 已退款
	TRANSACTION_STATUS_EXPIRED  ThirdPartyDepositStatus = 5 // 已過期
	TRANSACTION_STATUS_FAILED   ThirdPartyDepositStatus = 6 // 失敗
	TRANSACTION_STATUS_COMPLETE ThirdPartyDepositStatus = 7 // 已完成
)

type ThirdPartyPlatform string

const (
	THIRD_PARTY_PLATFORM_BINANCE_PAY ThirdPartyPlatform = "BINANCEPAY"
)

const (
	FIAT_TO_CRYPTO_DEPOSIT_CHANGELLY = "changelly"
)

type TransactionAddressPoolStatus int // 交易地址池狀態
const (
	TRANSACTION_ADDRESS_AVAILABLE TransactionAddressPoolStatus = 1
	TRANSACTION_ADDRESS_INUSE     TransactionAddressPoolStatus = 2
)

type UserTempAddressUsage int // 用戶臨時領用地址用途
const (
	ADDRESS_USAGE_ORIGINAL        UserTempAddressUsage = 1
	ADDRESS_USAGE_DAPP_CARD_APPLY UserTempAddressUsage = 2
	ADDRESS_USAGE_DAPP_DEPOSIT    UserTempAddressUsage = 3
	ADDRESS_USAGE_CHANGELLY       UserTempAddressUsage = 4
	ADDRESS_USAGE_BINANCEPAY      UserTempAddressUsage = 5
)

type UserTempAddressRealseStatus int //用戶地址資源釋放狀態
const (
	TEMP_ADDRESS_RELEASED     UserTempAddressRealseStatus = 1
	TEMP_ADDRESS_NOT_RELEASED UserTempAddressRealseStatus = 2
)

type UnblockMode int // 解鎖模式
const (
	UNBLOCK_MODE_DIRECT       UnblockMode = 1
	UNBLOCK_MODE_TOP_UP       UnblockMode = 2
	UNBLOCK_MODE_CARD_TO_CARD UnblockMode = 3
)

type UserBlockStatus int // 表示使用者帳戶的鎖定狀態
const (
	USER_BLOCK_STATUS_NORMAL UserBlockStatus = 1
	USER_BLOCK_STATUS_BLOCK  UserBlockStatus = 2
)

type UserTerminateStatus int

const (
	USER_TERMINATED_STATUS_ACTIVE     UserTerminateStatus = 1
	USER_TERMINATED_STATUS_TERMINATED UserTerminateStatus = 2
)

type ExchangeRateSource string

const (
	EXCHANGE_RATE_SOURCE_BITOP   ExchangeRateSource = "BITOP"
	EXCHANGE_RATE_SOURCE_BINANCE ExchangeRateSource = "BINANCE"
)

type FileAssessSource string

const (
	FILE_ACCESS_SOURCE_S3 = "s3"
)

type CoinfaceKycResponseCode int // KYC 返回代碼
const (
	COINFACE_KYC_RESPONSE_SUCCESS                       CoinfaceKycResponseCode = 0    //成功
	COINFACE_KYC_RESPONSE_INVALID_SESSION_ID            CoinfaceKycResponseCode = 1001 // 无效 session id
	COINFACE_KYC_RESPONSE_INVALID_COMPANY_ID            CoinfaceKycResponseCode = 1002 // 无效 company id
	COINFACE_KYC_RESPONSE_INVALID_LIVENESS_IMAGE        CoinfaceKycResponseCode = 1003 // 活体检测图片无效
	COINFACE_KYC_RESPONSE_LIVENESS_NOT_MATCH_ID         CoinfaceKycResponseCode = 1004 // 活体检测图片和证件照对比未通过
	COINFACE_KYC_RESPONSE_LIVENESS_COMPARE_EXCEPTION    CoinfaceKycResponseCode = 1005 // 活体检测图片和证件照对比异常(未检测人脸或未知异
	COINFACE_KYC_RESPONSE_LIVENESS_FAILED               CoinfaceKycResponseCode = 1006 // 活体检测未通过
	COINFACE_KYC_RESPONSE_UPLOAD_IMAGE_AWS_FAILED       CoinfaceKycResponseCode = 1007 // 上传图片到 aws 失败
	COINFACE_KYC_RESPONSE_INTERNAL_ERROR                CoinfaceKycResponseCode = 1008 // 内部错误
	COINFACE_KYC_RESPONSE_TOO_MANY_FAILURES             CoinfaceKycResponseCode = 1009 // 失败次数超过限制，暂时锁定该用户认证
	COINFACE_KYC_RESPONSE_INVALID_COMPANY_USER_ID       CoinfaceKycResponseCode = 1010 // 无效 company user id
	COINFACE_KYC_RESPONSE_UNRECOGNIZED_IMAGE_TYPE       CoinfaceKycResponseCode = 1011 // 不能识别的图片类型
	COINFACE_KYC_RESPONSE_USER_NOT_FOUND                CoinfaceKycResponseCode = 1012 // 用户不存在或状态不正确
	COINFACE_KYC_RESPONSE_IMAGE_SAVE_FAILED             CoinfaceKycResponseCode = 1013 // 图片保存失败
	COINFACE_KYC_RESPONSE_GET_AWS_IMAGE_URL_FAILED      CoinfaceKycResponseCode = 1014 // 获取aws图片url失败
	COINFACE_KYC_RESPONSE_OCR_FAILED                    CoinfaceKycResponseCode = 1015 // OCR 失败
	COINFACE_KYC_RESPONSE_TIMEOUT                       CoinfaceKycResponseCode = 1016 // 认证超时失败
	COINFACE_KYC_RESPONSE_INVALID_TYPE                  CoinfaceKycResponseCode = 1017 // 无效类型
	COINFACE_KYC_RESPONSE_MISSING_IMAGE                 CoinfaceKycResponseCode = 1018 // 缺少图片
	COINFACE_KYC_RESPONSE_USER_IN_BLACKLIST             CoinfaceKycResponseCode = 1019 // 用户已经在黑名单内
	COINFACE_KYC_RESPONSE_GET_GLOBAL_LOCK_FAILED        CoinfaceKycResponseCode = 1020 // 获取全局锁失败
	COINFACE_KYC_RESPONSE_SIGNATURE_VERIFICATION_FAILED CoinfaceKycResponseCode = 1021 // 签名验证失败
	COINFACE_KYC_RESPONSE_FILE_INTEGRITY_VERIFICATION   CoinfaceKycResponseCode = 1022 // 文件完整性验证失败
	COINFACE_KYC_RESPONSE_DUPLICATE_REGISTRATION        CoinfaceKycResponseCode = 1023 // 用户重复注册
	COINFACE_KYC_RESPONSE_THIRD_PARTY_ERROR             CoinfaceKycResponseCode = 1024 // 第三方错误
)

type ReapReportRange int // REAP BILL REPORT RAGENGE TYPE
const (
	REAP_REPORT_RANGE_DAILY   ReapReportRange = 1
	REAP_REPORT_RANGE_MONTHLY ReapReportRange = 2
)

type PromotionCodeType int // 推廣碼類型
const (
	PROMOOTION_CODE_TYPE_INVITE  PromotionCodeType = 1 // 邀請碼
	PROMOTION_CODE_TYPE_DISCOUNT PromotionCodeType = 2 // 折扣碼
)

type TerminateOrderStatus int

const (
	TERMINATE_ORDER_STATUS_PENDING TerminateOrderStatus = 1
	TERMINATE_ORDER_STATUS_SUCCESS TerminateOrderStatus = 2
	TERMINATE_ORDER_STATUS_FAILED  TerminateOrderStatus = 3
)

type TopDownOrderPurpose int

const (
	TOP_DOWN_ORDER_PURPOSE_TOP_DOWN  TopDownOrderPurpose = 1
	TOP_DOWN_ORDER_PURPOSE_TERMINATE TopDownOrderPurpose = 2
)

type EtherfiRole string

const (
	ETHERFI_ROLE_OWNER    EtherfiRole = "owner"
	ETHERFI_ROLE_ADMIN    EtherfiRole = "admin"
	ETHERFI_ROLE_EMPLOYEE EtherfiRole = "employee"
)

type EtherfiCardType string

const (
	ETHERFI_CARD_VIRTUAL  EtherfiCardType = "Virtual"
	ETHERFI_CARD_PHYSICAL EtherfiCardType = "Physical"
)

type EtherfiMailType int

const (
	ETHERFI_MAIL_TYPE_UNKNOWN     EtherfiMailType = 1
	ETHERFI_MAIL_TYPE_OTP_CODE    EtherfiMailType = 2
	ETHERFI_MAIL_TYPE_INVITE      EtherfiMailType = 3
	ETHERFI_MAIL_TYPE_BINDING     EtherfiMailType = 4
	ETHERFI_MAIL_TYPE_TRANSACTION EtherfiMailType = 5
)

type EtherfiTransactionHistoryEvent string

const (
	ETHERFI_ADDING_FUNDS      EtherfiTransactionHistoryEvent = "ADDING_FUNDS"
	ETHERFI_WITHDRAWALS       EtherfiTransactionHistoryEvent = "WITHDRAWALS"
	ETHERFI_CONVERSIONS       EtherfiTransactionHistoryEvent = "CONVERSIONS"
	ETHERFI_REPAYS            EtherfiTransactionHistoryEvent = "REPAYS"
	ETHERFI_REMOVED           EtherfiTransactionHistoryEvent = "REMOVED"
	ETHERFI_REFERRALS         EtherfiTransactionHistoryEvent = "REFERRALS"
	ETHERFI_CARD_TRANSACTIONS EtherfiTransactionHistoryEvent = "CARD_TRANSACTIONS"
)

type EtherfiTransactionStatus string

const (
	ETHERFI_PENDING   EtherfiTransactionStatus = "PENDING"
	ETHERFI_CLEARED   EtherfiTransactionStatus = "CLEARED"
	ETHERFI_DECLINED  EtherfiTransactionStatus = "DECLINED"
	ETHERFI_CANCELLED EtherfiTransactionStatus = "CANCELLED"
	ETHERFI_REFUND    EtherfiTransactionStatus = "REFUND"
	ETHERFI_ADDAUTH   EtherfiTransactionStatus = "ADDAUTH"
)

type EtherfiTransactionLifecycleEventAction string

const (
	ETHERFI_REQUESTED EtherfiTransactionLifecycleEventAction = "requested"
	ETHERFI_CREATED   EtherfiTransactionLifecycleEventAction = "created"
	ETHERFI_UPDATED   EtherfiTransactionLifecycleEventAction = "updated"
	ETHERFI_COMPLETED EtherfiTransactionLifecycleEventAction = "completed"
)

type EtherfiTransactionEventType string

const (
	RAIN_TRANSACTION EtherfiTransactionEventType = "rain_transaction"
)

type EtherfiApplyCardCacheStatus int

const (
	ETHERFI_APPLY_CARD_CACHE_STATUS_PENDING EtherfiApplyCardCacheStatus = 1
	ETHERFI_APPLY_CARD_CACHE_STATUS_SUCCESS EtherfiApplyCardCacheStatus = 2
	ETHERFI_APPLY_CARD_CACHE_STATUS_FAILED  EtherfiApplyCardCacheStatus = 3
)

type EtherfiTransactionLifecycleEventStatus int

const (
	ETHERFI_LIFECYCLE_EVENT_STATUS_PENDING   EtherfiTransactionLifecycleEventStatus = 1
	ETHERFI_LIFECYCLE_EVENT_STATUS_COMPLETED EtherfiTransactionLifecycleEventStatus = 2
)

type PhysicalCardRequestItemStatus int

const (
	PHYSICAL_CARD_REQUEST_ITEM_STATUS_PENDING PhysicalCardRequestItemStatus = 1
	PHYSICAL_CARD_REQUEST_ITEM_STATUS_SUCCESS PhysicalCardRequestItemStatus = 2
	PHYSICAL_CARD_REQUEST_ITEM_STATUS_FAIIED  PhysicalCardRequestItemStatus = 3
)

type PhysicalCardRequestOrderMerchant int

const (
	PHYSICAL_CARD_REQUEST_ORDER_MERCHANT_REAP   PhysicalCardRequestOrderMerchant = 1
	PHYSICAL_CARD_REQUEST_ORDER_MERCHANT_MONETA PhysicalCardRequestOrderMerchant = 2
	PHYSICAL_CARD_REQUEST_ORDER_MERCHANT_JINBEL PhysicalCardRequestOrderMerchant = 3
)

const PHYSICAL_CARD_APPLY_ORDER_POSITION = uint64(9999)

const PHYSICAL_CARD_AAPLY_ORDER_ITEM_POSITION = uint64(279999)

const PHYSICAL_CARD_APPLY_ORDER_ITEM_ADDRESS_LINE_1 = "#%d Prosperity Place"

const (
	PHYSICAL_CARD_APPLY_REAP_BULK_SHIPPING_ADDRESS_LINE1      = "TechPlus Consulting, INFINITY8 MyTOWN, L3-023"
	PHYSICAL_CARD_APPLY_REAP_BULK_SHIPPING_ADDRESS_LINE2      = "MyTOWN Shopping Centre, 6, Jln Cochrane, SEK 90"
	PHYSICAL_CARD_APPLY_REAP_BULK_SHIPPING_ADDRESS_CITY       = "Kuala Lumpur"
	PHYSICAL_CARD_APPLY_REAP_BULK_SHIPPING_ADDRESS_POSTALCODE = "55100"
	PHYSICAL_CARD_APPLY_REAP_BULK_SHIPPING_ADDRESS_COUNTRY    = "MY"
	PHYSICAL_CARD_APPLY_REAP_SHIPPING_ADDRESS_LINE1           = ""
	PHYSICAL_CARD_APPLY_REAP_SHIPPING_ADDRESS_LINE2           = "6 SHING YIP STREET, KWUN TONG"
	PHYSICAL_CARD_APPLY_REAP_SHIPPING_ADDRESS_ZONE            = "HongKong"
	PHYSICAL_CARD_APPLY_REAP_SHIPPING_ADDRESS_CITY            = "HongKong"
	PHYSICAL_CARD_APPLY_REAP_SHIPPING_ADDRESS_POSTAL_CODE     = "999077"
	PHYSICAL_CARD_APPLY_REAP_SHIPPING_ADDRESS_COUNTRY         = "HK"

	PHYSICAL_CARD_APPLY_REAP_RECIPIENT_INFORMATION_RECIPIENT_FIRST_NAME   = "WYNN"
	PHYSICAL_CARD_APPLY_REAP_RECIPIENT_INFORMATION_RECIPIENT_LAST_NAME    = "ZHAO"
	PHYSICAL_CARD_APPLY_REAP_RECIPIENT_INFORMATION_RECIPIENT_PHONE_NUMBER = "126681531"
	PHYSICAL_CARD_APPLY_REAP_RECIPIENT_INFORMATION_RECIPIENT_DIAL_CODE    = "60"
	PHYSICAL_CARD_APPLY_REAP_RECIPIENT_INFORMATION_RECIPIENT_EMAIL        = "zwynn90@gmail.com"
)

type EtherfiConfigEnable int

const (
	ETHERFI_CONFIG_ENABLE_ENABLE  EtherfiConfigEnable = 1
	ETHERFI_CONFIG_ENABLE_DISABLE EtherfiConfigEnable = 2
	ETHERFI_CONFIG_ENABLE_LOCK    EtherfiConfigEnable = 3
)

type EtherfiReportSource int

const (
	ETHERFI_REPORT_SOURCE_EUSD    EtherfiReportSource = 1
	ETHERFI_REPORT_SOURCE_ETHERFI EtherfiReportSource = 2
)
