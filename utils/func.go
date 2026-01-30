package utils

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"shared-modules/common"
	"shared-modules/logger"

	"github.com/shopspring/decimal"
)

func InArrayInt(v int64, a []int64) bool {

	for i := 0; i < len(a); i++ {
		if v == a[i] {
			return true
		}
	}
	return false
}

func InArrayString(v string, a []string) bool {

	for i := 0; i < len(a); i++ {
		if v == a[i] {
			return true
		}
	}
	return false
}

func MakeDirByMd5(name string) (string, error) {

	path := Config.System.UploadPath + "/" + name[0:2] + "/" + name[2:4]
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}
	return path, err
}

func GetRandString(n int) string {
	chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	charsLen := len(chars)
	if n > 10 {
		n = 10
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randStr := ""
	for i := 0; i < n; i++ {
		randIndex := r.Intn(charsLen)
		randStr += chars[randIndex : randIndex+1]
	}
	return randStr
}

func ToStr(value interface{}) string {
	var str string
	if value == nil {
		return str
	}
	switch v := value.(type) {
	case float64:
		str = strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		str = strconv.FormatFloat(float64(v), 'f', -1, 64)
	case int:
		str = strconv.Itoa(v)
	case uint:
		str = strconv.Itoa(int(v))
	case int8:
		str = strconv.Itoa(int(v))
	case uint8:
		str = strconv.Itoa(int(v))
	case int16:
		str = strconv.Itoa(int(v))
	case uint16:
		str = strconv.Itoa(int(v))
	case int32:
		str = strconv.Itoa(int(v))
	case uint32:
		str = strconv.Itoa(int(v))
	case int64:
		str = strconv.FormatInt(v, 10)
	case uint64:
		str = strconv.FormatUint(v, 10)
	case string:
		str = v
	case []byte:
		str = string(v)
	case time.Time:
		str = v.Format("2006-01-02 15:04:05")
	case decimal.Decimal:
		str = v.String()
	default:
		newValue, _ := json.Marshal(value)
		str = string(newValue)
	}
	return str
}

func GetTodayLeftSeconds() int64 {

	timeZone, _ := time.LoadLocation("Asia/Shanghai") //上海
	timeTemplate := "2006-01-02 15:04:05"
	timeStr := time.Now().Format("2006-01-02")
	nowTimeStr := time.Now().In(timeZone).Format(timeTemplate)
	endTimeStr := timeStr + " 23:59:59"

	time1, _ := time.Parse(timeTemplate, nowTimeStr)
	time2, _ := time.Parse(timeTemplate, endTimeStr)

	t1 := time1.Unix()
	t2 := time2.Unix()

	return t2 - t1
}

func Ptr[T any](t T) *T {
	return &t
}

func DecPtrToStr(dec *decimal.Decimal, places ...int) string {
	if dec != nil {
		if len(places) > 0 {
			return dec.StringFixed(int32(places[0]))
		}
		return dec.String()
	}
	return ""
}

func IsSuspectedFraud(status common.CardBlockReason) bool {
	switch status {
	case common.CARD_STATUS_EUSD_USER_BLOCK,
		common.CARD_STATUS_EUSD_USER_ALL_BLOCK,
		common.CARD_STATUS_EUSD_CUSTOMER_REQUIRE_BLOCK,
		common.CARD_STATUS_EUSD_REAP_KYC,
		common.CARD_STATUS_SUSPECTED_FRAUD_CVV,
		common.CARD_STATUS_SUSPECTED_FRAUD_PIN,
		common.CARD_STATUS_SUSPECTED_FRAUD_EXPIRY,
		common.CARD_STATUS_FUND_INSUFFICIENT_BALANCE,
		common.CARD_STATUS_FUND_DAILY_WITHDRAW_EXCEEDED,
		common.CARD_STATUS_FUND_DAILY_TX_EXCEEDED,
		common.CARD_STATUS_FUND_TX_AMOUNT_EXCEEDED,
		common.CARD_STATUS_FUND_EXCEED_MAX_CONSECUTIVE_FAILURES,
		common.CARD_STATUS_SECURITY_RISKY_TX,
		common.CARD_STATUS_SECURITY_INCORRECT_PIN,
		common.CARD_STATUS_SECURITY_UNUSUAL_LOCATION,
		common.CARD_STATUS_SECURITY_FOREIGN_TX,
		common.CARD_STATUS_SECURITY_CARD_LOST,
		common.CARD_STATUS_SECURITY_CARD_STOLEN,
		common.CARD_STATUS_SECURITY_RISKY_PAYEE,
		common.CARD_STATUS_TECH_PAYEE_AUTH_FAIL,
		common.CARD_STATUS_TECH_DEVICE_STRICT,
		common.CARD_STATUS_TECH_PAYEE_STRICT,
		common.CARD_STATUS_LEGAL_SUSPICIOUS_TX,
		common.CARD_STATUS_LEGAL_FOREIGN_COMPLIANCE,
		common.CARD_STATUS_LEGAL_RESTRICTED_CATEGORY,
		common.CARD_STATUS_LEGAL_FRAUD_RISK,
		common.CARD_STATUS_LEGAL_AGE_LIMIT,
		common.CARD_STATUS_ACCOUNT_IDENTITY_INCORRECT,
		common.CARD_STATUS_ACCOUNT_RESTRICTED_BY_PROVIDER,
		common.CARD_STATUS_ACCOUNT_RESTRICTED_BY_HOLDER,
		common.CARD_STATUS_ACCOUNT_MAIN_ACCOUNT,
		common.CARD_STATUS_ACCOUNT_AUTH_FAIL,
		common.CARD_STATUS_ACCOUNT_ANNUAL_FEE_UNPAID,
		common.CARD_STATUS_ACCOUNT_PAST_DUE,
		common.CARD_STATUS_ACCOUNT_NOT_IN_USE,
		common.CARD_STATUS_ACCOUNT_INACTIVE,
		common.CARD_STATUS_ACCOUNT_HOLDER_DECEASED,
		common.CARD_STATUS_ACCOUNT_HOLDER_BANKRUPT,
		common.CARD_STATUS_ACCOUNT_INVALID_KYC:
		return true
	default:
		return false
	}
}

type PromotuionUrl struct {
	Url string `json:"url"`
}

func GetPromotionLink(ctx context.Context, promotionCode string) (string, error) {
	url := Config.System.PromotionLink

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")

	body := map[string]interface{}{
		"branch_key": Config.System.PromotionLinkKey,
		"channel":    promotionCode,
		"feature":    "promo",
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		logger.Warn("marshal failed", err)
		return "", NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	logger.Infof("promotion Link  bodyJSON: %s", string(bodyJSON))

	resBody, _, code, err := HttpPost(ctx, url, string(bodyJSON), header)
	if err != nil {
		logger.Warn("post failed", err)
		return "", NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	if code != 200 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, code, err)
		return "", NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	res := &PromotuionUrl{}
	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return "", NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}
	return res.Url, nil
}
