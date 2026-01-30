package utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"shared-modules/common"
	"shared-modules/logger"

	"github.com/shopspring/decimal"
)

type MerchantCallbackForm struct {
	Data string `json:"data"`
	Sign string `json:"sign"`
}

type FeesForm struct {
	FXFees  decimal.Decimal `json:"fxFees"`
	ATMFees decimal.Decimal `json:"atmFees"`
}

type MerchantDataForm struct {
	MCCCode          string  `json:"mccCode"`
	MerchantID       string  `json:"merchantId"`
	MCCCategory      string  `json:"mccCategory"`
	MerchantCity     string  `json:"merchantCity"`
	MerchantName     string  `json:"merchantName"`
	MerchantState    *string `json:"merchantState,omitempty"`
	MerchantCountry  string  `json:"merchantCountry"`
	MerchantPostCode *string `json:"merchantPostCode,omitempty"`
}

type DeclineReasonForm struct {
	Description  string `json:"description"`
	ResponseCode string `json:"responseCode"`
}

type EventDataType interface {
	TransactionEventForm | AuthorizationEventForm | ForwardForm | WalletOTPForm | CardStatusForm | ExportEventForm | KYCEventForm
}

type AuthorizationEventForm struct {
	ID                  uint64           `json:"id"`
	LifecycleEventID    string           `json:"lifecycleEventId"`
	Fees                FeesForm         `json:"fees,omitempty"`
	Wallet              string           `json:"wallet"`
	AuthorizationID     uint64           `json:"authorizationId"`
	OrderNO             string           `json:"orderNo"`
	CardID              uint64           `json:"cardId"`
	Channel             string           `json:"channel"`
	CreatedAt           int64            `json:"createdAt"`
	BillAmount          decimal.Decimal  `json:"billAmount"`
	BillCurrency        string           `json:"billCurrency"`
	TransactionAmount   decimal.Decimal  `json:"transactionAmount"`
	TransactionCurrency string           `json:"transactionCurrency"`
	ExchangeRate        decimal.Decimal  `json:"exchangeRate"`
	MerchantData        MerchantDataForm `json:"merchantData"`
}

type TransactionEventForm struct {
	ID                  uint64            `json:"id"`
	LifecycleEventID    string            `json:"lifecycleEventId"`
	Fees                FeesForm          `json:"fees"`
	Wallet              string            `json:"wallet"`
	OrderNO             string            `json:"orderNo"`
	ParentOrderNO       string            `json:"parentOrderNo,omitempty"`
	CardID              uint64            `json:"cardId"`
	Channel             string            `json:"channel"`
	ClearedAt           int64             `json:"clearedAt,omitempty"`
	AuthorizedAt        int64             `json:"authorizedAt,omitempty"`
	ReversalAt          int64             `json:"reversalAt,omitempty"`
	RefundedAt          int64             `json:"refundedAt,omitempty"`
	RejectedAt          int64             `json:"rejectedAt,omitempty"`
	BillAmount          decimal.Decimal   `json:"billAmount"`
	BillCurrency        string            `json:"billCurrency"`
	TransactionAmount   decimal.Decimal   `json:"transactionAmount"`
	TransactionCurrency string            `json:"transactionCurrency"`
	ExchangeRate        decimal.Decimal   `json:"exchangeRate"`
	MerchantData        MerchantDataForm  `json:"merchantData"`
	DeclineReason       DeclineReasonForm `json:"declineReason"`
}

type ExportEventForm struct {
	ID          uint64              `json:"id"`
	FileID      uint64              `json:"fileId,string,omitempty"`
	ExportToken string              `json:"exportToken"`
	DownloadURL string              `json:"downloadUrl,omitempty"`
	CreatedAt   int64               `json:"createdAt,omitempty"`
	ExpiredAt   int64               `json:"expiredAt,omitempty"`
	Status      common.ExportStatus `json:"status"`
}

type KYCEventForm struct {
	UserID uint64 `json:"userId,string"`
	Email  string `json:"email"`
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
}

type EventForm[T EventDataType] struct {
	EventName common.MerchantEventName `json:"eventName"`
	EventType common.MerchantEventType `json:"eventType"`
	Data      T                        `json:"data"`
}

type AuthorizationEventVO struct {
	Data string `json:"data"`
	Sign string `json:"sign"`
}

type AuthorizationEventDataVO struct {
	AuthorizationID uint64              `json:"authorizationId,omitempty"`
	ResponseCode    common.DeclinedCode `json:"responseCode"`
}

type TransactionEventVO struct {
	Response[interface{}]
}

type CardStatusForm struct {
	ID             uint64                 `json:"id"`
	CardID         uint64                 `json:"cardId"`
	Status         string                 `json:"status,omitempty"`
	FreezeStatus   string                 `json:"freezeStatus,omitempty"`
	FreezeReason   string                 `json:"freezeReason,omitempty"`
	BlockReason    common.CardBlockReason `json:"blockReason,omitempty"`
	DeliveryStatus string                 `json:"deliveryStatus,omitempty"`
	TrackingLink   string                 `json:"trackingLink,omitempty"`
}

type CardStatusVO struct {
	Response[interface{}]
}

type ForwardForm struct {
	ID                   uint64            `json:"id"`
	CardID               uint64            `json:"cardId" gorm:"default:null"`
	MerchantName         string            `json:"merchantName"`
	TransactionAmount    decimal.Decimal   `json:"transactionAmount"`
	MerchantCountryCode  common.NationCode `json:"merchantCountryCode"`
	TransactionCurrency  string            `json:"transactionCurrency"`
	TransactionTimestamp time.Time         `json:"transactionTimestamp"`
}

type ForwardVO struct {
	Response[interface{}]
}

type WalletOTPForm struct {
	OTP         string `json:"otp"`
	Wallet      string `json:"wallet"`
	CardID      uint64 `json:"cardId,string"`
	DialCode    int    `json:"dialCode"`
	PhoneNumber string `json:"phoneNumber"`
}

type WalletOTPVO struct {
	Response[interface{}]
}

func MerchantPayRequest(ctx context.Context, form *EventForm[AuthorizationEventForm], pubKey string, url string) (
	resData *AuthorizationEventDataVO, oriReq string, encReq string, oriRes string, code int, status common.CallbackStatus, err error) {

	jf, err := json.Marshal(form)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}
	logger.Infof("form: %s", jf)

	decodedKey, err := base64.StdEncoding.DecodeString(Config.Reap.MerchantPrivateKey)
	if err != nil {
		logger.Warn("decode failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	sign, err := RsaSignWithSha256(jf, decodedKey)
	if err != nil {
		logger.Warn("sign failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}
	logger.Infof("sign: %s", base64.StdEncoding.EncodeToString(sign))

	pf := &MerchantCallbackForm{
		Data: string(jf),
		Sign: base64.StdEncoding.EncodeToString(sign),
	}

	payload, err := json.Marshal(pf)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")

	resBody, _, code, err := HttpPost(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		return nil, string(jf), string(payload), "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR, err.Error())
	}
	logger.Infof("merchant response body: %s", string(resBody))
	if code < 200 || code >= 300 {
		logger.Warnf("post failed, resp body: [%s], code: [%d]", string(resBody), code)
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	res := &Response[*AuthorizationEventDataVO]{}

	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR, err.Error())
	}

	if res.Code != common.CODE_OK {
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	if res.Data == nil {
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	return res.Data, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_SUCCESS, nil
}

func MerchantPayTX(ctx context.Context, form *EventForm[TransactionEventForm], pubKey string, url string) (res *TransactionEventVO, oriReq string, encReq string, oriRes string, code int, status common.CallbackStatus, err error) {

	jf, err := json.Marshal(form)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, "", "", "", 0, 0, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}
	logger.Infof("form: %s", jf)

	decodedKey, err := base64.StdEncoding.DecodeString(Config.Reap.MerchantPrivateKey)
	if err != nil {
		logger.Warn("decode failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	sign, err := RsaSignWithSha256(jf, decodedKey)
	if err != nil {
		logger.Warn("sign failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	pf := &MerchantCallbackForm{
		Data: string(jf),
		Sign: base64.StdEncoding.EncodeToString(sign),
	}

	payload, err := json.Marshal(pf)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")

	resBody, _, code, err := HttpPost(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		return nil, string(jf), string(payload), "", 0, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR, err.Error())
	}
	logger.Infof("merchant response body: %s", string(resBody))
	if code < 200 || code >= 300 {
		logger.Warnf("post failed, resp body: [%s], code: [%d]", string(resBody), code)
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	res = &TransactionEventVO{}
	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR, err.Error())
	}

	if res.Code != common.CODE_OK {
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	return res, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_SUCCESS, nil
}

func MerchantForward(ctx context.Context, form *EventForm[ForwardForm], pubKey string, url string) (res *ForwardVO, oriReq string, encReq string, oriRes string, code int, status common.CallbackStatus, err error) {

	jf, err := json.Marshal(form)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, "", "", "", 0, 0, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}
	logger.Infof("form: %s", jf)

	decodedKey, err := base64.StdEncoding.DecodeString(Config.Reap.MerchantPrivateKey)
	if err != nil {
		logger.Warn("decode failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	sign, err := RsaSignWithSha256(jf, decodedKey)
	if err != nil {
		logger.Warn("sign failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	pf := &MerchantCallbackForm{
		Data: string(jf),
		Sign: base64.StdEncoding.EncodeToString(sign),
	}

	payload, err := json.Marshal(pf)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")

	resBody, _, code, err := HttpPost(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		return nil, string(jf), string(payload), "", 0, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR, err.Error())
	}
	logger.Infof("merchant response body: %s", string(resBody))
	if code < 200 || code >= 300 {
		logger.Warnf("post failed, resp body: [%s], code: [%d]", string(resBody), code)
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	res = &ForwardVO{}
	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR, err.Error())
	}

	if res.Code != common.CODE_OK {
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	return res, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_SUCCESS, nil
}

func MerchantWalletOTP(ctx context.Context, form *EventForm[WalletOTPForm], pubKey string, url string) (res *WalletOTPVO, oriReq string, encReq string, oriRes string, code int, status common.CallbackStatus, err error) {

	jf, err := json.Marshal(form)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, "", "", "", 0, 0, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}
	logger.Infof("form: %s", jf)

	decodedKey, err := base64.StdEncoding.DecodeString(Config.Reap.MerchantPrivateKey)
	if err != nil {
		logger.Warn("decode failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	sign, err := RsaSignWithSha256(jf, decodedKey)
	if err != nil {
		logger.Warn("sign failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	pf := &MerchantCallbackForm{
		Data: string(jf),
		Sign: base64.StdEncoding.EncodeToString(sign),
	}

	payload, err := json.Marshal(pf)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")

	resBody, _, code, err := HttpPost(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		return nil, string(jf), string(payload), "", 0, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR, err.Error())
	}
	logger.Infof("merchant response body: %s", string(resBody))
	if code < 200 || code >= 300 {
		logger.Warnf("post failed, resp body: [%s], code: [%d]", string(resBody), code)
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	res = &WalletOTPVO{}
	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR, err.Error())
	}

	if res.Code != common.CODE_OK {
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	return res, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_SUCCESS, nil
}

func MerchantCardStatus(ctx context.Context, form *EventForm[CardStatusForm], pubKey string, url string) (res *CardStatusVO, oriReq string, encReq string, oriRes string, code int, status common.CallbackStatus, err error) {

	jf, err := json.Marshal(form)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, "", "", "", 0, 0, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}
	logger.Infof("form: %s", jf)

	decodedKey, err := base64.StdEncoding.DecodeString(Config.Reap.MerchantPrivateKey)
	if err != nil {
		logger.Warn("decode failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	sign, err := RsaSignWithSha256(jf, decodedKey)
	if err != nil {
		logger.Warn("sign failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	pf := &MerchantCallbackForm{
		Data: string(jf),
		Sign: base64.StdEncoding.EncodeToString(sign),
	}

	payload, err := json.Marshal(pf)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")

	resBody, _, code, err := HttpPost(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		return nil, string(jf), string(payload), "", 0, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR, err.Error())
	}
	logger.Infof("merchant response body: %s", string(resBody))
	if code < 200 || code >= 300 {
		logger.Warnf("post failed, resp body: [%s], code: [%d]", string(resBody), code)
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	res = &CardStatusVO{}
	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR, err.Error())
	}

	if res.Code != common.CODE_OK {
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	return res, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_SUCCESS, nil
}

func MerchantExport(ctx context.Context, form *EventForm[ExportEventForm], pubKey string, url string) (res *TransactionEventVO, oriReq string, encReq string, oriRes string, code int, status common.CallbackStatus, err error) {

	jf, err := json.Marshal(form)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, "", "", "", 0, 0, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}
	logger.Infof("form: %s", jf)

	decodedKey, err := base64.StdEncoding.DecodeString(Config.Reap.MerchantPrivateKey)
	if err != nil {
		logger.Warn("decode failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	sign, err := RsaSignWithSha256(jf, decodedKey)
	if err != nil {
		logger.Warn("sign failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	pf := &MerchantCallbackForm{
		Data: string(jf),
		Sign: base64.StdEncoding.EncodeToString(sign),
	}

	payload, err := json.Marshal(pf)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")

	resBody, _, code, err := HttpPost(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		return nil, string(jf), string(payload), "", 0, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR, err.Error())
	}
	logger.Infof("merchant response body: %s", string(resBody))
	if code < 200 || code >= 300 {
		logger.Warnf("post failed, resp body: [%s], code: [%d]", string(resBody), code)
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	res = &TransactionEventVO{}
	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR, err.Error())
	}

	if res.Code != common.CODE_OK {
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	return res, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_SUCCESS, nil
}

func MerchantKYC(ctx context.Context, form *EventForm[KYCEventForm], pubKey string, url string) (res *TransactionEventVO, oriReq string, encReq string, oriRes string, code int, status common.CallbackStatus, err error) {

	jf, err := json.Marshal(form)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, "", "", "", 0, 0, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}
	logger.Infof("form: %s", jf)

	decodedKey, err := base64.StdEncoding.DecodeString(Config.Reap.MerchantPrivateKey)
	if err != nil {
		logger.Warn("decode failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	sign, err := RsaSignWithSha256(jf, decodedKey)
	if err != nil {
		logger.Warn("sign failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	pf := &MerchantCallbackForm{
		Data: string(jf),
		Sign: base64.StdEncoding.EncodeToString(sign),
	}

	payload, err := json.Marshal(pf)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, "", "", "", 0, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")

	resBody, _, code, err := HttpPost(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		return nil, string(jf), string(payload), "", 0, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR, err.Error())
	}
	logger.Infof("merchant response body: %s", string(resBody))
	if code < 200 || code >= 300 {
		logger.Warnf("post failed, resp body: [%s], code: [%d]", string(resBody), code)
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	res = &TransactionEventVO{}
	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_PENDING_RETRY, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR, err.Error())
	}

	if res.Code != common.CODE_OK {
		return nil, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_FAILED, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	return res, string(jf), string(payload), string(resBody), code, common.CALLBACK_STATUS_SUCCESS, nil
}
