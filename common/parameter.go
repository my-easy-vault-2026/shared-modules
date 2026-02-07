package common

type ParameterCategory int // 參數類別
const (
	PARAMETER_CATEGORY_ACCOUNT  ParameterCategory = 1
	PARAMETER_CATEGORY_AUTH     ParameterCategory = 2
	PARAMETER_CATEGORY_USER     ParameterCategory = 3
	PARAMETER_CATEGORY_WALLET   ParameterCategory = 4
	PARAMETER_CATEGORY_EXCHANGE ParameterCategory = 5
	PARAMETER_CATEGORY_TRANSFER ParameterCategory = 6
	PARAMETER_CATEGORY_ORDER    ParameterCategory = 7
	PARAMETER_CATEGORY_QUOTE    ParameterCategory = 8
	PARAMETER_CATEGORY_SYSTEM   ParameterCategory = 9
)

// parameter key PARAMETER_KEY_{category}_{key}
type ParameterKey string // 參數鍵
const (
	PARAMETER_KEY_EXCHANGE_EXCHANGE_FEE ParameterKey = "exchange_fee"
	PARAMETER_KEY_TRANSFER_TRANSFER_FEE ParameterKey = "transfer_fee"
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

type EncryptType int // 加密類型
const (
	ENCRYPT_TYPE_BASE64 EncryptType = 1
	ENCRYPT_TYPE_AES    EncryptType = 2
	ENCRYPT_TYPE_RSA    EncryptType = 3
)
