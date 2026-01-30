package utils

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"shared-modules/common"
	"shared-modules/logger"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

const (
	paycryptoSignSeparator = ":"
)

var (
	urlPaycryptoGetInstituteBalance = func() (string, string, string) { return "GET", Config.Paycrypto.BaseURL, "/api/v1/institution/balance" }
	urlPaycryptoGetCardTypes        = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/institution/card/type"
	}
	urlPaycryptoGetInstituteInfo = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/institution/info"
	}
	urlPaycryptoUploadPublicKey = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/institution/publickey"
	}
	urlPaycryptoGetRates = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/institution/rates"
	}
	urlPaycryptoEstimateCurrency = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/institution/estimation/currency"
	}
	urlPaycryptoEstimateCurrencyExternalDeduction = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/institution/estimation/currency/external-deduction"
	}
	urlPaycryptoEstimateCrypto = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/institution/estimation/crypto"
	}
	urlPaycryptoKYCSubmitUserData = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/customers/accounts"
	}
	urlPaycryptoKYCSubmitUserAttachments = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/customers/attachments"
	}
	urlPaycryptoKYCGetAllUserRecords = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/customers/accounts"
	}
	urlPaycryptoKYCGetSpecificUserRecord = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/customers/accounts"
	}
	urlPaycryptoKYCUpdateUserKYCInfo = func() (string, string, string) {
		return "PUT", Config.Paycrypto.BaseURL, "/api/v1/customers/accounts"
	}
	urlPaycryptoKYCUpdateUserAddress = func() (string, string, string) {
		return "PUT", Config.Paycrypto.BaseURL, "/api/v1/customers/accounts/address"
	}
	urlPaycryptoKYCGetUserKYCInfo = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/customers/accounts/kyc"
	}
	urlPaycryptoSubmitCardApplication = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/debit-cards"
	}
	urlPaycryptoSubmitCardActivationAttachment = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/debit-cards/attachment"
	}
	urlPaycryptoActivateCard = func() (string, string, string) {
		return "PUT", Config.Paycrypto.BaseURL, "/api/v1/debit-cards/status"
	}
	urlPaycryptoGetAllCards = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/debit-cards"
	}
	urlPaycryptoGetUserCards = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/debit-cards"
	}
	urlPaycryptoSubmitCardRequest = func() (string, string, string) {
		return "PUT", Config.Paycrypto.BaseURL, "/api/v1/debit-cards/request"
	}
	urlPaycryptoGetCardRequestStatus = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/debit-cards/request"
	}
	urlPaycryptoGetTrackingNumber = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/debit-cards/trackingnumber"
	}
	urlPaycryptoUpgradeCard = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/debit-cards/upgrade"
	}
	urlPaycryptoResendPINEmail = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/debit-cards/resendemail"
	}
	urlPaycryptoDepositTransaction = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/deposit-transactions"
	}
	urlPaycryptoExternalDeductionDeposit = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/deposit-transactions/external-deduction"
	}
	urlPaycryptoFiatAmountDeposit = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/deposit-transactions/fiat-amount"
	}
	urlPaycryptoCurrencyPairPrice = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/deposit-transactions/price"
	}
	urlPaycryptoDepositTransactionStatus = func(txID string) (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, fmt.Sprintf("/api/v1/deposit-transactions/%s/status", txID)
	}
	urlPaycryptoAllDepositTransactions = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/deposit-transactions"
	}
	urlPaycryptoUserDepositTransactions = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/deposit-transactions"
	}
	urlPaycryptoExchangeTransaction = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/exchange-transactions"
	}
	urlPaycryptoExchangeTransactions = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/exchange-transactions"
	}
	urlPaycryptoCheckCardActivation = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/bank/account-status"
	}
	urlPaycryptoGetCardBalance = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/bank/balance"
	}
	urlPaycryptoGetCardTransactions = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/bank/transaction-record"
	}
	urlPaycryptoGetVirtualCardInfo = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/bank/virtualcard"
	}
	urlPaycryptoGetPhysicalCardInfo = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/bank/card"
	}
	urlPaycryptoInitPhysicalCardPassword = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/bank/initpassword"
	}
	urlPaycryptoResetPhysicalCardPassword = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/bank/resetpassword"
	}
	urlPaycryptoResetCardPassword = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/bank/reset-pwd"
	}
	urlPaycryptoGetAuthorizedTransactions = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/bank/authorizedtransaction"
	}
	urlPaycryptoUpdateCardLimit = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/bank/cardlimit"
	}
	urlPaycryptoGetCardLimit = func() (string, string, string) {
		return "GET", Config.Paycrypto.BaseURL, "/api/v1/bank/cardlimit"
	}
	urlPaycryptoFreezeCard = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/bank/freeze"
	}
	urlPaycryptoUnfreezeCard = func() (string, string, string) {
		return "POST", Config.Paycrypto.BaseURL, "/api/v1/bank/unfreeze"
	}
)

// 通用接口定義
type PaycryptoResponseResult interface {
	bool |
		map[string]interface{} |
		[]PaycryptoGetInstituteBalanceVO |
		PaycryptoCardTypeResponse |
		PaycryptoInstituteInfoResponse |
		PaycryptoRatesResponse |
		PaycryptoEstimationResponse |
		PaycryptoKYCUserRecordsResponse |
		PaycryptoKYCUserInfoResponse |
		PaycryptoCardApplicationResponse |
		PaycryptoCardActivationResponse |
		PaycryptoCardRecordsResponse |
		PaycryptoCardRequestStatusResponse |
		[]PaycryptoCardTrackingResponse |
		PaycryptoDepositResponse |
		PaycryptoDepositTransactionResponse |
		PaycryptoCurrencyPairPriceResponse |
		PaycryptoDepositTransactionsResponse |
		PaycryptoExchangeTransactionResponse |
		PaycryptoExchangeTransactionsResponse |
		PaycryptoCardBalanceResponse |
		[]PaycryptoCardStatementResponse |
		PaycryptoCardSensitiveInfoResponse |
		[]PaycryptoAuthorizedTransactionResponse |
		PaycryptoCardLimitResponse
}

type PaycryptoResponse[T PaycryptoResponseResult] struct {
	Code   *int64 `json:"code,omitempty"`
	Msg    string `json:"msg"`
	Result T      `json:"result,omitempty"`
}

type PaycryptoGetInstituteBalanceForm struct{}
type PaycryptoGetInstituteBalanceVO struct {
	Addresses []PaycryptoGetInstituteBalanceVOAddress `json:"addresses"`
	CoinType  string                                  `json:"coin_type"`
}
type PaycryptoGetInstituteBalanceVOAddress struct {
	Address     string          `json:"address"`
	Balance     decimal.Decimal `json:"balance"`
	AddressType string          `json:"address_type"`
}

func PaycryptoGetInstituteBalance(ctx context.Context, form *PaycryptoGetInstituteBalanceForm) (res *PaycryptoResponse[[]PaycryptoGetInstituteBalanceVO], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetInstituteBalance()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("paycrypto response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[[]PaycryptoGetInstituteBalanceVO]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// 支持的卡种类查询 API
type PaycryptoCardTypeResponse struct {
	Total   int64               `json:"total"`
	Records []PaycryptoCardType `json:"records"`
}

type PaycryptoCardType struct {
	CardTypeID   string `json:"card_type_id"`
	CurrencyType string `json:"currency_type"`
	BankID       string `json:"bank_id"`
	Description  string `json:"description"`
	CardNetwork  string `json:"card_network"`
	CardTitle    string `json:"card_title"`
	VirtualCard  bool   `json:"virtual_card"`
}

type PaycryptoGetCardTypesForm struct{}

func PaycryptoGetCardTypes(ctx context.Context, form *PaycryptoGetCardTypesForm) (res *PaycryptoResponse[PaycryptoCardTypeResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetCardTypes()
	requestQueryStr := ""

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)
	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath

	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("paycrypto response body: %s", string(resBody))
	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	res = &PaycryptoResponse[PaycryptoCardTypeResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// 公钥查询 API
type PaycryptoInstituteInfoResponse struct {
	Publickey string `json:"publickey"`
}

type PaycryptoGetInstituteInfoForm struct{}

func PaycryptoGetInstituteInfo(ctx context.Context, form *PaycryptoGetInstituteInfoForm) (res *PaycryptoResponse[PaycryptoInstituteInfoResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetInstituteInfo()
	requestQueryStr := ""

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)
	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath

	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	res = &PaycryptoResponse[PaycryptoInstituteInfoResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// 上传公钥 API
type PaycryptoUploadPublicKeyForm struct {
	PublicKey string `json:"public_key"`
}

func PaycryptoUploadPublicKey(ctx context.Context, form *PaycryptoUploadPublicKeyForm) (res *PaycryptoResponse[bool], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoUploadPublicKey()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))
	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath

	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	res = &PaycryptoResponse[bool]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// 卡费率查询 API
type PaycryptoStepRate struct {
	Min               decimal.Decimal `json:"min"`
	Max               decimal.Decimal `json:"max"`
	PaycryptoStepRate decimal.Decimal `json:"step_rate"`
}

type PaycryptoRatesResponse struct {
	ExchangeRate        decimal.Decimal     `json:"exchange_rate"`
	CardApplicationFee  decimal.Decimal     `json:"card_application_fee"`
	MinDeposit          decimal.Decimal     `json:"min_deposit"`
	MaxDeposit          decimal.Decimal     `json:"max_deposit"`
	LoadingRate         []PaycryptoStepRate `json:"loading_rate"`
	BankAtmFee          decimal.Decimal     `json:"bank_atm_fee"`
	BankAtmRate         decimal.Decimal     `json:"bank_atm_rate"`
	BankTransactionRate decimal.Decimal     `json:"bank_transaction_rate"`
}

type PaycryptoGetRatesForm struct {
	CardTypeID string `json:"card_type_id"`
}

func PaycryptoGetRates(ctx context.Context, form *PaycryptoGetRatesForm) (res *PaycryptoResponse[PaycryptoRatesResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetRates()
	requestQueryStr := "card_type_id=" + form.CardTypeID

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)
	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath + "?" + requestQueryStr

	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	res = &PaycryptoResponse[PaycryptoRatesResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// 估算将到账的法币金额 API
type PaycryptoEstimationCurrencyForm struct {
	CoinAmount decimal.Decimal `json:"coin_amount"`
	CoinType   string          `json:"coin_type"`
	CardTypeID string          `json:"card_type_id"`
}

type PaycryptoEstimationResponse struct {
	CoinType         string          `json:"coin_type"`
	CoinAmount       decimal.Decimal `json:"coin_amount"`
	CurrencyType     string          `json:"currency_type"`
	CurrencyAmount   decimal.Decimal `json:"currency_amount"`
	ExchangeRate     decimal.Decimal `json:"exchange_rate"`
	FiatExchangeRate decimal.Decimal `json:"fiat_exchange_rate"`
	ExchangeFee      decimal.Decimal `json:"exchange_fee"`
	ExchangeFeeRate  decimal.Decimal `json:"exchange_fee_rate"`
	LoadingFee       decimal.Decimal `json:"loading_fee"`
	CoinPrice        decimal.Decimal `json:"coin_price"`
}

func PaycryptoEstimateCurrency(ctx context.Context, form *PaycryptoEstimationCurrencyForm) (res *PaycryptoResponse[PaycryptoEstimationResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoEstimateCurrency()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))
	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath

	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	res = &PaycryptoResponse[PaycryptoEstimationResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// 估算将到账的法币金额（外扣）API
func PaycryptoEstimateCurrencyExternalDeduction(ctx context.Context, form *PaycryptoEstimationCurrencyForm) (res *PaycryptoResponse[PaycryptoEstimationResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoEstimateCurrencyExternalDeduction()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))
	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath

	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	res = &PaycryptoResponse[PaycryptoEstimationResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// 估算需要充值多少数字货币 API
type PaycryptoEstimationCryptoForm struct {
	CurrencyAmount string `json:"currency_amount"`
	CardTypeID     string `json:"card_type_id"`
	CoinType       string `json:"coin_type"`
}

func PaycryptoEstimateCrypto(ctx context.Context, form *PaycryptoEstimationCryptoForm) (res *PaycryptoResponse[PaycryptoEstimationResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoEstimateCrypto()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))
	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath

	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	res = &PaycryptoResponse[PaycryptoEstimationResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

func PaycryptoHmacSign(timestamp, action, apiSecret string, body []byte) string {
	// 構建簽名字符串
	var message strings.Builder
	message.WriteString(timestamp)
	message.WriteString(strings.ToUpper(action))

	// 如果有body，添加body
	if body != nil && len(body) > 0 {
		message.Write(body)
	}

	// 使用HMAC-SHA256算法計算簽名
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(message.String()))

	// 返回Base64編碼的簽名
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, queryString, apiKey, apiSecret string, body []byte) string {
	// 構建簽名字符串
	var message strings.Builder
	message.WriteString(timestamp)
	message.WriteString(strings.ToUpper(method))
	message.WriteString(apiKey)
	message.WriteString(requestPath)
	if queryString != "" {
		message.WriteString("?")
		message.WriteString(queryString)
	}

	// 如果有body，添加body
	if body != nil && len(body) > 0 {
		message.Write(body)
	}

	// 使用HMAC-SHA256算法計算簽名
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(message.String()))

	// 返回Base64編碼的簽名
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func PaycryptoConvertRequestBodyToSignString(jsonBody []byte) (string, error) {
	// 解析JSON字符串為map[string]json.RawMessage
	var bodyMap map[string]json.RawMessage
	err := json.Unmarshal(jsonBody, &bodyMap)
	if err != nil {
		return "", fmt.Errorf("解析JSON失敗: %v", err)
	}

	// 獲取所有鍵並按ASCII碼排序
	keys := make([]string, 0, len(bodyMap))
	for k := range bodyMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 構建參數字符串
	var pairs []string
	for _, key := range keys {
		rawValue := bodyMap[key]
		var valueStr string

		// 根據原始JSON值的格式直接處理
		rawStr := string(rawValue)

		// 檢查是否為null
		if rawStr == "null" {
			valueStr = ""
		} else if len(rawStr) >= 2 && rawStr[0] == '"' && rawStr[len(rawStr)-1] == '"' {
			// 字符串類型，去除引號
			valueStr = rawStr[1 : len(rawStr)-1]
		} else if rawStr == "true" || rawStr == "false" {
			// 布爾類型
			valueStr = rawStr
		} else {
			// 數字類型或其他
			// 嘗試解析為數字以確保格式正確
			if bytes.ContainsAny(rawValue, ".eE") {
				// 可能是浮點數
				var floatVal float64
				if err := json.Unmarshal(rawValue, &floatVal); err == nil {
					valueStr = strconv.FormatFloat(floatVal, 'f', -1, 64)
				} else {
					valueStr = rawStr
				}
			} else {
				// 可能是整數
				var intVal int64
				if err := json.Unmarshal(rawValue, &intVal); err == nil {
					valueStr = strconv.FormatInt(intVal, 10)
				} else {
					valueStr = rawStr
				}
			}
		}

		pairs = append(pairs, fmt.Sprintf("%s=%s", key, valueStr))
	}

	// 用&連接所有參數對
	return strings.Join(pairs, "&"), nil
}

// Submit User KYC Data form
type PaycryptoKYCSubmitUserDataForm struct {
	AcctNo               string   `json:"acct_no"`
	AcctName             string   `json:"acct_name"`
	FirstName            string   `json:"first_name"`
	LastName             string   `json:"last_name"`
	Gender               string   `json:"gender"`
	Birthday             string   `json:"birthday"`
	City                 string   `json:"city"`
	State                string   `json:"state"`
	Country              string   `json:"country"`
	Nationality          string   `json:"nationality"`
	DocNo                string   `json:"doc_no"`
	DocType              string   `json:"doc_type"`
	FrontDoc             string   `json:"front_doc"`
	BackDoc              string   `json:"back_doc,omitempty"`
	MixDoc               string   `json:"mix_doc"`
	CountryCode          string   `json:"country_code"`
	Mobile               string   `json:"mobile"`
	Mail                 string   `json:"mail"`
	Address              string   `json:"address"`
	ZipCode              string   `json:"zipcode"`
	MaidenName           string   `json:"maiden_name"`
	CardTypeID           string   `json:"card_type_id"`
	KYCInfo              string   `json:"kyc_info,omitempty"`
	MailVerificationCode string   `json:"mail_verification_code,omitempty"`
	MailToken            string   `json:"mail_token,omitempty"`
	CustTxID             string   `json:"cust_tx_id,omitempty"`
	SignImg              string   `json:"sign_img,omitempty"`
	PoaDoc               []string `json:"poa_doc,omitempty"`
	CardNumber           string   `json:"card_number,omitempty"`
}

// Submit User KYC Data
func PaycryptoKYCSubmitUserData(ctx context.Context, form *PaycryptoKYCSubmitUserDataForm) (res *PaycryptoResponse[bool], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoKYCSubmitUserData()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	dataCopy := form
	if len(dataCopy.FrontDoc) > 10 {
		dataCopy.FrontDoc = dataCopy.FrontDoc[:10] + "..."
	}
	if len(dataCopy.BackDoc) > 10 {
		dataCopy.BackDoc = dataCopy.BackDoc[:10] + "..."
	}
	if len(dataCopy.MixDoc) > 10 {
		dataCopy.MixDoc = dataCopy.MixDoc[:10] + "..."
	}

	logger.Infof("path: [%s], req: [%s]", url, dataCopy)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("kyc response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[bool]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Submit User KYC Attachments form
type PaycryptoKYCSubmitUserAttachmentsForm struct {
	AcctNo     string   `json:"acct_no"`
	CardTypeID string   `json:"card_type_id"`
	Docs       []string `json:"docs"`
}

// Submit User KYC Attachments
func PaycryptoKYCSubmitUserAttachments(ctx context.Context, form *PaycryptoKYCSubmitUserAttachmentsForm) (res *PaycryptoResponse[bool], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoKYCSubmitUserAttachments()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("kyc response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[bool]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// KYC User Info Response
type PaycryptoKYCUserInfoResponse struct {
	Total   int64                  `json:"total"`
	Records []PaycryptoKYCUserInfo `json:"records"`
}

type PaycryptoKYCUserInfo struct {
	AcctNo      string                     `json:"acct_no"`
	AcctName    string                     `json:"acct_name"`
	FirstName   string                     `json:"first_name"`
	LastName    string                     `json:"last_name"`
	Gender      string                     `json:"gender"`
	Birthday    string                     `json:"birthday"`
	City        string                     `json:"city"`
	State       string                     `json:"state"`
	Country     string                     `json:"country"`
	Nationality string                     `json:"nationality"`
	DocNo       string                     `json:"doc_no"`
	DocType     string                     `json:"doc_type"`
	CountryCode string                     `json:"country_code"`
	Mobile      string                     `json:"mobile"`
	Mail        string                     `json:"mail"`
	Address     string                     `json:"address"`
	ZipCode     string                     `json:"zipcode"`
	MaidenName  string                     `json:"maiden_name"`
	CardTypeID  string                     `json:"card_type_id"`
	KYCInfo     string                     `json:"kyc_info"`
	CustTxID    string                     `json:"cust_tx_id"`
	Status      *common.PaycryptoKYCStatus `json:"status"` // 状态码: 0 已提交， 1 认证通过(开卡成功)， 2 认证未通过， 3 认证中， 4 提交信息处理中 6 已退款
}

// Get All User KYC Records query params
type PaycryptoKYCGetAllUserRecordsParams struct {
	PageNum    int    `json:"page_num,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	FormerTime int64  `json:"former_time,omitempty"`
	LatterTime int64  `json:"latter_time,omitempty"`
	TimeSort   string `json:"time_sort,omitempty"`
}

// KYC User Records Response
type PaycryptoKYCUserRecordsResponse struct {
	Total   int64                    `json:"total"`
	Records []PaycryptoKYCUserRecord `json:"records"`
}

type PaycryptoKYCUserRecord struct {
	AcctNo             string                     `json:"acct_no"`
	CardTypeID         string                     `json:"card_type_id"`
	Status             *common.PaycryptoKYCStatus `json:"status"`
	FaceRecognitionURL string                     `json:"face_recognition_url,omitempty"`
	Reason             string                     `json:"reason,omitempty"`
	CreateTime         int64                      `json:"create_time"`
}

// Get All User KYC Records
func PaycryptoKYCGetAllUserRecords(ctx context.Context, params *PaycryptoKYCGetAllUserRecordsParams) (res *PaycryptoResponse[PaycryptoKYCUserRecordsResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoKYCGetAllUserRecords()

	// Build query string
	queryValues := url.Values{}
	if params.PageNum > 0 {
		queryValues.Add("page_num", strconv.Itoa(params.PageNum))
	}
	if params.PageSize > 0 {
		queryValues.Add("page_size", strconv.Itoa(params.PageSize))
	}
	if params.FormerTime > 0 {
		queryValues.Add("former_time", strconv.FormatInt(params.FormerTime, 10))
	}
	if params.LatterTime > 0 {
		queryValues.Add("latter_time", strconv.FormatInt(params.LatterTime, 10))
	}
	if params.TimeSort != "" {
		queryValues.Add("time_sort", params.TimeSort)
	}
	requestQueryStr := queryValues.Encode()

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("kyc response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoKYCUserRecordsResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Get Specific User KYC Record query params
type PaycryptoKYCGetSpecificUserRecordParams struct {
	AcctNo     string `json:"acct_no"`
	PageNum    int    `json:"page_num,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	FormerTime int64  `json:"former_time,omitempty"`
	LatterTime int64  `json:"latter_time,omitempty"`
	TimeSort   string `json:"time_sort,omitempty"`
}

// Get Specific User KYC Record
func PaycryptoKYCGetSpecificUserRecord(ctx context.Context, params *PaycryptoKYCGetSpecificUserRecordParams) (res *PaycryptoResponse[PaycryptoKYCUserRecordsResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoKYCGetSpecificUserRecord()

	// Build query string
	queryValues := url.Values{}
	queryValues.Add("acct_no", params.AcctNo)
	if params.PageNum > 0 {
		queryValues.Add("page_num", strconv.Itoa(params.PageNum))
	}
	if params.PageSize > 0 {
		queryValues.Add("page_size", strconv.Itoa(params.PageSize))
	}
	if params.FormerTime > 0 {
		queryValues.Add("former_time", strconv.FormatInt(params.FormerTime, 10))
	}
	if params.LatterTime > 0 {
		queryValues.Add("latter_time", strconv.FormatInt(params.LatterTime, 10))
	}
	if params.TimeSort != "" {
		queryValues.Add("time_sort", params.TimeSort)
	}
	requestQueryStr := queryValues.Encode()

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("kyc response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoKYCUserRecordsResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Update User KYC Info form
type PaycryptoKYCUpdateUserKYCInfoForm struct {
	AcctNo  string `json:"acct_no"`
	KYCInfo string `json:"kyc_info"`
}

// Update User KYC Info
func PaycryptoKYCUpdateUserKYCInfo(ctx context.Context, form *PaycryptoKYCUpdateUserKYCInfoForm) (res *PaycryptoResponse[bool], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoKYCUpdateUserKYCInfo()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("kyc response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("put failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[bool]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Update User Address form
type PaycryptoKYCUpdateUserAddressForm struct {
	AcctNo      string `json:"acct_no"`
	Mobile      string `json:"mobile"`
	CountryCode string `json:"country_code"`
	Country     string `json:"country"`
	State       string `json:"state"`
	City        string `json:"city"`
	Address     string `json:"address"`
	ZipCode     string `json:"zipcode"`
}

// Update User Address
func PaycryptoKYCUpdateUserAddress(ctx context.Context, form *PaycryptoKYCUpdateUserAddressForm) (res *PaycryptoResponse[bool], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoKYCUpdateUserAddress()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("kyc response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("put failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[bool]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Get User KYC Info query params
type PaycryptoKYCGetUserKYCInfoParams struct {
	AcctNo string `json:"acct_no"`
}

// Get User KYC Info
func PaycryptoKYCGetUserKYCInfo(ctx context.Context, params *PaycryptoKYCGetUserKYCInfoParams) (res *PaycryptoResponse[PaycryptoKYCUserInfoResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoKYCGetUserKYCInfo()

	// Build query string
	queryValues := url.Values{}
	queryValues.Add("acct_no", params.AcctNo)
	requestQueryStr := queryValues.Encode()

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("kyc response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoKYCUserInfoResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Card Activation Response
type PaycryptoCardActivationResponse struct {
	CardTypeID string                              `json:"card_type_id"`
	CardNo     string                              `json:"card_no"`
	AcctNo     string                              `json:"acct_no"`
	Status     *common.PaycryptoCardActivateStatus `json:"status"` // 状态码: 0 冻结， 1 激活成功， 2未激活， 3. 激活待审核， 4. 激活审核失败, 5. 申请失败(卡片正在制作中，请过会再申请) 6.已注销
}

// Card Records Response
type PaycryptoCardRecordsResponse struct {
	Total   int64                 `json:"total"`
	Records []PaycryptoCardRecord `json:"records"`
}

type PaycryptoCardRecord struct {
	AcctNo     string                              `json:"acct_no"`
	CardTypeID string                              `json:"card_type_id"`
	CardNo     string                              `json:"card_no"`
	Status     *common.PaycryptoCardActivateStatus `json:"status"`
	Reason     string                              `json:"reason,omitempty"`
	CreateTime int64                               `json:"create_time"`
}

// Card Request Status Response
type PaycryptoCardRequestStatusResponse struct {
	AcctNo     string                             `json:"acct_no"`
	CardNo     string                             `json:"card_no"`
	Status     *common.PaycryptoCardRequestStatus `json:"status"` //状态码: 0.处理中 1.成功 2.失败
	CreateTime int64                              `json:"create_time"`
}

// Card Tracking Response
type PaycryptoCardTrackingResponse struct {
	AcctNo         string `json:"acct_no"`
	CardNo         string `json:"card_no"`
	Website        string `json:"website"`
	TrackingNumber string `json:"tracking_number"`
	CreateTime     int64  `json:"create_time"`
}

// Card Application Form
type PaycryptoSubmitCardApplicationForm struct {
	AcctNo     string `json:"acct_no"`
	CardTypeID string `json:"card_type_id"`
	CardNumber string `json:"card_number,omitempty"`
	URN        string `json:"urn,omitempty"`
}

// Card Application Response
type PaycryptoCardApplicationResponse struct {
	CardNo     string                                `json:"card_no"`
	CardNumber string                                `json:"card_number"`
	Status     common.PaycryptoCardApplicationStatus `json:"status"` // 状态：2. 开卡申请成功(未激活)， 5. 申请失败(卡片正在制作中，请过会再申请)
}

// Submit Card Application
func PaycryptoSubmitCardApplication(ctx context.Context, form *PaycryptoSubmitCardApplicationForm) (res *PaycryptoResponse[PaycryptoCardApplicationResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoSubmitCardApplication()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoCardApplicationResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Card Activation Attachment Form
type PaycryptoSubmitCardActivationAttachmentForm struct {
	CardNo    string   `json:"card_no"`
	PoaDoc    []string `json:"poa_doc,omitempty"`
	ActiveDoc string   `json:"active_doc,omitempty"`
}

// Submit Card Activation Attachment
func PaycryptoSubmitCardActivationAttachment(ctx context.Context, form *PaycryptoSubmitCardActivationAttachmentForm) (res *PaycryptoResponse[bool], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoSubmitCardActivationAttachment()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[bool]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Card Activation Form
type PaycryptoActivateCardForm struct {
	CardNo string `json:"card_no"`
	AcctNo string `json:"acct_no"`
	CVV    string `json:"cvv,omitempty"`
}

// Activate Card
func PaycryptoActivateCard(ctx context.Context, form *PaycryptoActivateCardForm) (res *PaycryptoResponse[PaycryptoCardActivationResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoActivateCard()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("put failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoCardActivationResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Get All Cards Query Params
type PaycryptoGetAllCardsParams struct {
	PageNum    int    `json:"page_num,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	FormerTime int64  `json:"former_time,omitempty"`
	LatterTime int64  `json:"latter_time,omitempty"`
	TimeSort   string `json:"time_sort,omitempty"`
}

// Get All Cards
func PaycryptoGetAllCards(ctx context.Context, params *PaycryptoGetAllCardsParams) (res *PaycryptoResponse[PaycryptoCardRecordsResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetAllCards()

	// Build query string
	queryValues := url.Values{}
	if params.PageNum > 0 {
		queryValues.Add("page_num", strconv.Itoa(params.PageNum))
	}
	if params.PageSize > 0 {
		queryValues.Add("page_size", strconv.Itoa(params.PageSize))
	}
	if params.FormerTime > 0 {
		queryValues.Add("former_time", strconv.FormatInt(params.FormerTime, 10))
	}
	if params.LatterTime > 0 {
		queryValues.Add("latter_time", strconv.FormatInt(params.LatterTime, 10))
	}
	if params.TimeSort != "" {
		queryValues.Add("time_sort", params.TimeSort)
	}
	requestQueryStr := queryValues.Encode()

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoCardRecordsResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Get User Cards Query Params
type PaycryptoGetUserCardsParams struct {
	AcctNo     string `json:"acct_no"`
	PageNum    int    `json:"page_num,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	FormerTime int64  `json:"former_time,omitempty"`
	LatterTime int64  `json:"latter_time,omitempty"`
	TimeSort   string `json:"time_sort,omitempty"`
}

// Get User Cards
func PaycryptoGetUserCards(ctx context.Context, params *PaycryptoGetUserCardsParams) (res *PaycryptoResponse[PaycryptoCardRecordsResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetUserCards()

	// Build query string
	queryValues := url.Values{}
	queryValues.Add("acct_no", params.AcctNo)
	if params.PageNum > 0 {
		queryValues.Add("page_num", strconv.Itoa(params.PageNum))
	}
	if params.PageSize > 0 {
		queryValues.Add("page_size", strconv.Itoa(params.PageSize))
	}
	if params.FormerTime > 0 {
		queryValues.Add("former_time", strconv.FormatInt(params.FormerTime, 10))
	}
	if params.LatterTime > 0 {
		queryValues.Add("latter_time", strconv.FormatInt(params.LatterTime, 10))
	}
	if params.TimeSort != "" {
		queryValues.Add("time_sort", params.TimeSort)
	}
	requestQueryStr := queryValues.Encode()

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoCardRecordsResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Card Request Form
type PaycryptoSubmitCardRequestForm struct {
	CardNo       string `json:"card_no"`
	AcctNo       string `json:"acct_no"`
	RequestType  int    `json:"request_type"`
	SignatureDoc string `json:"signature_doc"`
	Address      string `json:"address,omitempty"`
	Phone        string `json:"phone,omitempty"`
}

// Submit Card Request
func PaycryptoSubmitCardRequest(ctx context.Context, form *PaycryptoSubmitCardRequestForm) (res *PaycryptoResponse[bool], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoSubmitCardRequest()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("put failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[bool]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Card Request Status Query Params
type PaycryptoGetCardRequestStatusParams struct {
	CardNo      string                          `json:"card_no"`
	AcctNo      string                          `json:"acct_no"`
	RequestType common.PaycryptoCardRequestType `json:"request_type"` //1.冻结 2.解冻 3.挂失 4.重置密码 5.补卡 6.重发PIN
}

// Get Card Request Status
func PaycryptoGetCardRequestStatus(ctx context.Context, params *PaycryptoGetCardRequestStatusParams) (res *PaycryptoResponse[PaycryptoCardRequestStatusResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetCardRequestStatus()

	// Build query string
	queryValues := url.Values{}
	queryValues.Add("card_no", params.CardNo)
	queryValues.Add("acct_no", params.AcctNo)
	queryValues.Add("request_type", strconv.Itoa(int(params.RequestType)))
	requestQueryStr := queryValues.Encode()

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoCardRequestStatusResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Tracking Number Query Params - by month_year
type PaycryptoGetTrackingNumberByMonthParams struct {
	MonthYear string `json:"month_year"`
}

// Get Tracking Number By Month
func PaycryptoGetTrackingNumberByMonth(ctx context.Context, params *PaycryptoGetTrackingNumberByMonthParams) (res *PaycryptoResponse[[]PaycryptoCardTrackingResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetTrackingNumber()

	// Build query string
	queryValues := url.Values{}
	queryValues.Add("month_year", params.MonthYear)
	requestQueryStr := queryValues.Encode()

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[[]PaycryptoCardTrackingResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Tracking Number Query Params - by card
type PaycryptoGetTrackingNumberByCardParams struct {
	CardNo string `json:"card_no"`
	AcctNo string `json:"acct_no"`
}

func PaycryptoGetTrackingNumberByCard(ctx context.Context, params *PaycryptoGetTrackingNumberByCardParams) (res *PaycryptoResponse[[]PaycryptoCardTrackingResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetTrackingNumber()

	// Build query string
	queryValues := url.Values{}
	queryValues.Add("card_no", params.CardNo)
	queryValues.Add("acct_no", params.AcctNo)
	requestQueryStr := queryValues.Encode()

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[[]PaycryptoCardTrackingResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Card Upgrade Form
type PaycryptoUpgradeCardForm struct {
	CardNo string `json:"card_no"`
	AcctNo string `json:"acct_no"`
}

// Upgrade Card
func PaycryptoUpgradeCard(ctx context.Context, form *PaycryptoUpgradeCardForm) (res *PaycryptoResponse[bool], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoUpgradeCard()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[bool]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Resend PIN Email Form
type PaycryptoResendPINEmailForm struct {
	CardNo string `json:"card_no"`
	AcctNo string `json:"acct_no"`
}

// Resend PIN Email
func PaycryptoResendPINEmail(ctx context.Context, form *PaycryptoResendPINEmailForm) (res *PaycryptoResponse[bool], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoResendPINEmail()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[bool]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Deposit Transaction Form (internal deduction)
type PaycryptoDepositTransactionForm struct {
	CardNo       string `json:"card_no"`
	AcctNo       string `json:"acct_no"`
	Amount       string `json:"amount"`
	CoinType     string `json:"coin_type"`
	CustTxID     string `json:"cust_tx_id"`
	Remark       string `json:"remark,omitempty"`
	CardCurrency string `json:"card_currency,omitempty"`
}

// Deposit Transaction Response
type PaycryptoDepositResponse struct {
	TxID             string          `json:"tx_id"`
	CoinType         string          `json:"coin_type"`
	TxAmount         decimal.Decimal `json:"tx_amount"`
	ExchangeFeeRate  decimal.Decimal `json:"exchange_fee_rate"`
	ExchangeFee      decimal.Decimal `json:"exchange_fee"`
	LoadingFee       decimal.Decimal `json:"loading_fee"`
	DepositUsdt      decimal.Decimal `json:"deposit_usdt"`
	CurrencyType     string          `json:"currency_type"`
	CurrencyAmount   decimal.Decimal `json:"currency_amount"`
	ExchangeRate     decimal.Decimal `json:"exchange_rate"`
	FiatExchangeRate decimal.Decimal `json:"fiat_exchange_rate"`
}

// Deposit Transaction (internal deduction)
func PaycryptoDepositTransaction(ctx context.Context, form *PaycryptoDepositTransactionForm) (res *PaycryptoResponse[PaycryptoDepositResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoDepositTransaction()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("deposit response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoDepositResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Deposit Transaction Form (external deduction)
type PaycryptoExternalDeductionDepositForm struct {
	CardNo       string `json:"card_no"`
	AcctNo       string `json:"acct_no"`
	Amount       string `json:"amount"`
	CoinType     string `json:"coin_type"`
	CustTxID     string `json:"cust_tx_id"`
	Remark       string `json:"remark,omitempty"`
	CardCurrency string `json:"card_currency,omitempty"`
}

// Deposit Transaction (external deduction)
func PaycryptoExternalDeductionDeposit(ctx context.Context, form *PaycryptoExternalDeductionDepositForm) (res *PaycryptoResponse[PaycryptoDepositResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoExternalDeductionDeposit()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("deposit response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoDepositResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Fiat Amount Deposit Form
type PaycryptoFiatAmountDepositForm struct {
	CardNo         string `json:"card_no"`
	AcctNo         string `json:"acct_no"`
	CreditedAmount string `json:"credited_amount"`
	CoinType       string `json:"coin_type"`
	CustTxID       string `json:"cust_tx_id"`
	Remark         string `json:"remark,omitempty"`
	CardCurrency   string `json:"card_currency,omitempty"`
}

// Fiat Amount Deposit
func PaycryptoFiatAmountDeposit(ctx context.Context, form *PaycryptoFiatAmountDepositForm) (res *PaycryptoResponse[PaycryptoDepositResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoFiatAmountDeposit()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("deposit response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoDepositResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Currency Pair Price Response
type PaycryptoCurrencyPairPriceResponse struct {
	Total   int64                        `json:"total"`
	Records []PaycryptoCurrencyPairPrice `json:"records"`
}

type PaycryptoCurrencyPairPrice struct {
	Symbol     string          `json:"symbol"`
	Price      decimal.Decimal `json:"price"`
	UpdateTime string          `json:"update_time"`
}

// Get Currency Pair Prices
func PaycryptoGetCurrencyPairPrices(ctx context.Context) (res *PaycryptoResponse[PaycryptoCurrencyPairPriceResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoCurrencyPairPrice()
	requestQueryStr := ""

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("price response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoCurrencyPairPriceResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Deposit Transaction Status Response
type PaycryptoDepositTransactionResponse struct {
	CustTxID         string                    `json:"cust_tx_id"`
	AcctNo           string                    `json:"acct_no"`
	CardNo           string                    `json:"card_no"`
	CustTxTime       int64                     `json:"cust_tx_time"`
	TxID             string                    `json:"tx_id"`
	CoinType         string                    `json:"coin_type"`
	TxAmount         decimal.Decimal           `json:"tx_amount"`
	ExchangeFee      decimal.Decimal           `json:"exchange_fee"`
	LoadingFee       decimal.Decimal           `json:"loading_fee"`
	CurrencyType     string                    `json:"currency_type"`
	CurrencyAmount   decimal.Decimal           `json:"currency_amount"`
	ExchangeRate     decimal.Decimal           `json:"exchange_rate"`
	FiatExchangeRate decimal.Decimal           `json:"fiat_exchange_rate"`
	TxStatus         *common.PaycryptoTXStatus `json:"tx_status"` // 交易状态。0、3、4:待处理中，1:充值成功，2充值失败，5:充值失败
	CoinPrice        decimal.Decimal           `json:"coin_price"`
}

// Get Deposit Transaction Status
func PaycryptoGetDepositTransactionStatus(ctx context.Context, txID string) (res *PaycryptoResponse[PaycryptoDepositTransactionResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoDepositTransactionStatus(txID)
	requestQueryStr := ""

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("deposit status response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoDepositTransactionResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Get All Deposit Transactions Params
type PaycryptoGetAllDepositTransactionsParams struct {
	PageNum    int    `json:"page_num,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	FormerTime int64  `json:"former_time,omitempty"`
	LatterTime int64  `json:"latter_time,omitempty"`
	TimeSort   string `json:"time_sort,omitempty"`
}

// Deposit Transactions Response
type PaycryptoDepositTransactionsResponse struct {
	Total   int64                                 `json:"total"`
	Records []PaycryptoDepositTransactionResponse `json:"records"`
}

// Get All Deposit Transactions
func PaycryptoGetAllDepositTransactions(ctx context.Context, params *PaycryptoGetAllDepositTransactionsParams) (res *PaycryptoResponse[PaycryptoDepositTransactionsResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoAllDepositTransactions()

	// Build query string
	queryValues := url.Values{}
	if params.PageNum > 0 {
		queryValues.Add("page_num", strconv.Itoa(params.PageNum))
	}
	if params.PageSize > 0 {
		queryValues.Add("page_size", strconv.Itoa(params.PageSize))
	}
	if params.FormerTime > 0 {
		queryValues.Add("former_time", strconv.FormatInt(params.FormerTime, 10))
	}
	if params.LatterTime > 0 {
		queryValues.Add("latter_time", strconv.FormatInt(params.LatterTime, 10))
	}
	if params.TimeSort != "" {
		queryValues.Add("time_sort", params.TimeSort)
	}
	requestQueryStr := queryValues.Encode()

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("deposit transactions response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoDepositTransactionsResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Get User Deposit Transactions Params
type PaycryptoGetUserDepositTransactionsParams struct {
	AcctNo     string `json:"acct_no"`
	PageNum    int    `json:"page_num,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	FormerTime int64  `json:"former_time,omitempty"`
	LatterTime int64  `json:"latter_time,omitempty"`
	TimeSort   string `json:"time_sort,omitempty"`
}

// Get User Deposit Transactions
func PaycryptoGetUserDepositTransactions(ctx context.Context, params *PaycryptoGetUserDepositTransactionsParams) (res *PaycryptoResponse[PaycryptoDepositTransactionsResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoUserDepositTransactions()

	// Build query string
	queryValues := url.Values{}
	queryValues.Add("acct_no", params.AcctNo)
	if params.PageNum > 0 {
		queryValues.Add("page_num", strconv.Itoa(params.PageNum))
	}
	if params.PageSize > 0 {
		queryValues.Add("page_size", strconv.Itoa(params.PageSize))
	}
	if params.FormerTime > 0 {
		queryValues.Add("former_time", strconv.FormatInt(params.FormerTime, 10))
	}
	if params.LatterTime > 0 {
		queryValues.Add("latter_time", strconv.FormatInt(params.LatterTime, 10))
	}
	if params.TimeSort != "" {
		queryValues.Add("time_sort", params.TimeSort)
	}
	requestQueryStr := queryValues.Encode()

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("user deposit transactions response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoDepositTransactionsResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

type PaycryptoExchangeTransactionForm struct {
	CardNo         string `json:"card_no"`
	AcctNo         string `json:"acct_no"`
	CurrencyAmount string `json:"currency_amount"`
	CustTxID       string `json:"cust_tx_id"`
	Remark         string `json:"remark,omitempty"`
	CardCurrency   string `json:"card_currency,omitempty"`
}

// Exchange Transaction Response
type PaycryptoExchangeTransactionResponse struct {
	CustTxID       string                    `json:"cust_tx_id"`
	AcctNo         string                    `json:"acct_no"`
	CardNo         string                    `json:"card_no"`
	CustTxTime     int64                     `json:"cust_tx_time"`
	TxID           string                    `json:"tx_id"`
	CoinType       string                    `json:"coin_type"`
	CoinAmount     decimal.Decimal           `json:"coin_amount"`
	ExchangeFee    decimal.Decimal           `json:"exchange_fee"`
	CurrencyType   string                    `json:"currency_type"`
	CurrencyAmount decimal.Decimal           `json:"currency_amount"`
	ExchangeRate   decimal.Decimal           `json:"exchange_rate"`
	TxStatus       *common.PaycryptoTXStatus `json:"tx_status"`
	Reason         string                    `json:"reason,omitempty"`
}

// Exchange Transaction
func PaycryptoExchangeTransaction(ctx context.Context, form *PaycryptoExchangeTransactionForm) (res *PaycryptoResponse[PaycryptoExchangeTransactionResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoExchangeTransaction()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("exchange response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoExchangeTransactionResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Get Exchange Transactions Params
type PaycryptoGetExchangeTransactionsParams struct {
	AcctNo     string `json:"acct_no"`
	CardNo     string `json:"card_no"`
	TxID       string `json:"tx_id,omitempty"`
	PageNum    int    `json:"page_num,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	FormerTime int64  `json:"former_time,omitempty"`
	LatterTime int64  `json:"latter_time,omitempty"`
	TimeSort   string `json:"time_sort,omitempty"`
}

// Exchange Transactions Response
type PaycryptoExchangeTransactionsResponse struct {
	Total   int64                                `json:"total"`
	Records []PaycryptoExchangeTransactionRecord `json:"records"`
}

type PaycryptoExchangeTransactionRecord struct {
	CustTxID       string                    `json:"cust_tx_id"`
	AcctNo         string                    `json:"acct_no"`
	CardNo         string                    `json:"card_no"`
	CustTxTime     int64                     `json:"cust_tx_time"`
	TxID           string                    `json:"tx_id"`
	CoinType       string                    `json:"coin_type"`
	CoinAmount     decimal.Decimal           `json:"coin_amount"`
	ExchangeFee    string                    `json:"exchange_fee"`
	CurrencyType   string                    `json:"currency_type"`
	CurrencyAmount decimal.Decimal           `json:"currency_amount"`
	ExchangeRate   decimal.Decimal           `json:"exchange_rate"`
	TxStatus       *common.PaycryptoTXStatus `json:"tx_status"`
	Reason         string                    `json:"reason,omitempty"`
}

// Get Exchange Transactions
func PaycryptoGetExchangeTransactions(ctx context.Context, params *PaycryptoGetExchangeTransactionsParams) (res *PaycryptoResponse[PaycryptoExchangeTransactionsResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoExchangeTransactions()

	// Build query string
	queryValues := url.Values{}
	queryValues.Add("acct_no", params.AcctNo)
	queryValues.Add("card_no", params.CardNo)
	if params.TxID != "" {
		queryValues.Add("tx_id", params.TxID)
	}
	if params.PageNum > 0 {
		queryValues.Add("page_num", strconv.Itoa(params.PageNum))
	}
	if params.PageSize > 0 {
		queryValues.Add("page_size", strconv.Itoa(params.PageSize))
	}
	if params.FormerTime > 0 {
		queryValues.Add("former_time", strconv.FormatInt(params.FormerTime, 10))
	}
	if params.LatterTime > 0 {
		queryValues.Add("latter_time", strconv.FormatInt(params.LatterTime, 10))
	}
	if params.TimeSort != "" {
		queryValues.Add("time_sort", params.TimeSort)
	}
	requestQueryStr := queryValues.Encode()

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("exchange transactions response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoExchangeTransactionsResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Check Card Activation Form
type PaycryptoCheckCardActivationForm struct {
	CardNo string `json:"card_no"`
}

// Check Card Activation
func PaycryptoCheckCardActivation(ctx context.Context, form *PaycryptoCheckCardActivationForm) (res *PaycryptoResponse[bool], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoCheckCardActivation()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card activation response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[bool]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Get Card Balance Form
type PaycryptoGetCardBalanceForm struct {
	CardNo       string `json:"card_no"`
	CardCurrency string `json:"card_currency,omitempty"`
}

// Card Balance Response
type PaycryptoCardBalanceResponse struct {
	CardNumber          string          `json:"card_number"`
	CardType            string          `json:"card_type"`
	CurrentBalance      decimal.Decimal `json:"current_balance"`
	CurrentBalanceUSD   decimal.Decimal `json:"current_balance_usd"`
	AvailableBalance    decimal.Decimal `json:"available_balance"`
	AvailableBalanceUSD decimal.Decimal `json:"available_balance_usd"`
}

// Get Card Balance
func PaycryptoGetCardBalance(ctx context.Context, form *PaycryptoGetCardBalanceForm) (res *PaycryptoResponse[PaycryptoCardBalanceResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetCardBalance()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card balance response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoCardBalanceResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

type PaycryptoMonthYear time.Time

func (d PaycryptoMonthYear) MarshalJSON() ([]byte, error) {
	// Step 1. Format the time as a Go string.
	t := time.Time(d)
	formatted := t.In(time.UTC).Format("012006")

	// Step 2. Convert our formatted time to a JSON string.
	jsonStr := "\"" + formatted + "\""
	return []byte(jsonStr), nil
}

func (d *PaycryptoMonthYear) UnmarshalJSON(b []byte) error {
	if len(b) < 2 || b[0] != '"' || b[len(b)-1] != '"' {
		return errors.New("not a json string")
	}

	// 1. Strip the double quotes from the JSON string.
	b = b[1 : len(b)-1]

	// 2. Parse the result using our desired format.
	t, err := time.Parse("012006", string(b))
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}
	t = TimeFromDB(t)

	// finally, assign t to *d
	*d = PaycryptoMonthYear(t)

	return nil
}

// Get Card Transactions Form
type PaycryptoGetCardTransactionsForm struct {
	CardNo          string             `json:"card_no"`
	TXID            string             `json:"tx_id"`
	CardCurrency    string             `json:"card_currency,omitempty"`
	FormerMonthYear PaycryptoMonthYear `json:"former_month_year"`
	LatterMonthYear PaycryptoMonthYear `json:"latter_month_year"`
}

type PaycryptoDate time.Time

func (d *PaycryptoDate) UnmarshalJSON(b []byte) error {
	if len(b) < 2 || b[0] != '"' || b[len(b)-1] != '"' {
		return errors.New("not a json string")
	}

	// 1. Strip the double quotes from the JSON string.
	b = b[1 : len(b)-1]

	// 2. Parse the result using our desired format.
	t, err := time.Parse("02/01/2006", string(b))
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}

	// finally, assign t to *d
	*d = PaycryptoDate(t)

	return nil
}

// Card Statement Response
type PaycryptoCardStatementResponse struct {
	MonthYear           PaycryptoMonthYear               `json:"month_year"`
	StatementCycleDate  PaycryptoDate                    `json:"statement_cycle_date"`
	OpeningBalance      decimal.Decimal                  `json:"opening_balance"`
	OpeningBalanceUSD   decimal.Decimal                  `json:"opening_balance_usd"`
	ClosingBalance      decimal.Decimal                  `json:"closing_balance"`
	ClosingBalanceUSD   decimal.Decimal                  `json:"closing_balance_usd"`
	AvailableBalance    decimal.Decimal                  `json:"available_balance"`
	AvailableBalanceUSD decimal.Decimal                  `json:"available_balance_usd"`
	BankTxList          []PaycryptoCardTransactionRecord `json:"bank_tx_list"`
}

// Card Transaction Record
type PaycryptoCardTransactionRecord struct {
	TransactionDate     PaycryptoDate                `json:"transaction_date"`
	PostingDate         PaycryptoDate                `json:"posting_date"`
	TxID                string                       `json:"tx_id"`
	Description         string                       `json:"description"`
	Debit               decimal.Decimal              `json:"debit"`
	DebitUSD            decimal.Decimal              `json:"debit_usd"`
	DebitUSDT           decimal.Decimal              `json:"debit_usdt"`
	Credit              decimal.Decimal              `json:"credit"`
	CreditUSD           decimal.Decimal              `json:"credit_usd"`
	CreditUSDT          decimal.Decimal              `json:"credit_usdt"`
	Fee                 decimal.Decimal              `json:"fee"`
	EndBal              decimal.Decimal              `json:"end_bal,omitempty"`               // 假的
	OriginTransactionID string                       `json:"origin_transaction_id,omitempty"` // 會放0
	Type                common.PaycryptoTxRecordType `json:"type,omitempty"`
	Reason              string                       `json:"reason,omitempty"`
	TxCurrency          string                       `json:"tx_currency,omitempty"` // TWD
	TxAmount            *decimal.Decimal             `json:"tx_amount,omitempty"`
	TransactionTime     int64                        `json:"transaction_time"`
}

// Get Card Transactions
func PaycryptoGetCardTransactions(ctx context.Context, form *PaycryptoGetCardTransactionsForm) (res *PaycryptoResponse[[]PaycryptoCardStatementResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetCardTransactions()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card transactions response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)
	filter := make(map[string]interface{})
	if err = json.Unmarshal(resBody, &filter); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	paycryptoIterDelete(filter, "N/A")

	filteredJSON, err := json.Marshal(filter)
	if err != nil {
		logger.Warnf("marshal failed, resp body: %s, code: %d, %v", filteredJSON, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	logger.Infof("filtered res: [%s]", filteredJSON)

	res = &PaycryptoResponse[[]PaycryptoCardStatementResponse]{}
	if err = json.Unmarshal(filteredJSON, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Get Card Sensitive Info Form
type PaycryptoGetCardSensitiveInfoForm struct {
	CardNo    string `json:"card_no"`
	PublicKey string `json:"publickey,omitempty"`
}

// Card Sensitive Info Response
type PaycryptoCardSensitiveInfoResponse struct {
	EncryptData []string `json:"encrypt_data"`
	PublicKey   string   `json:"public_key"`
}

// Card Sensitive Info Response
type PaycryptoCardSensitiveInfoDecryptedResponse struct {
	CardNumber string `json:"card_number"`
	Expire     string `json:"expire"`
	CVV        string `json:"cvv"` // 格式05/27
}

// Get Virtual Card Sensitive Info
func PaycryptoGetVirtualCardInfo(ctx context.Context, form *PaycryptoGetCardSensitiveInfoForm) (res *PaycryptoResponse[PaycryptoCardSensitiveInfoResponse], dec PaycryptoCardSensitiveInfoDecryptedResponse, err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetVirtualCardInfo()

	// Build query string
	queryValues := url.Values{}
	queryValues.Add("card_no", form.CardNo)
	requestQueryStr := queryValues.Encode()

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, requestQueryStr)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Debugf("virtual card info response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Debugf("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoCardSensitiveInfoResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	decodedKey, err := base64.StdEncoding.DecodeString(Config.Paycrypto.PrivateKey)
	if err != nil {
		logger.Warn("decode failed", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	if len(res.Result.EncryptData) == 0 {
		return
	}

	decodedData, err := base64.StdEncoding.DecodeString(res.Result.EncryptData[0])
	if err != nil {
		logger.Warn("decode failed", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	decryptedData, err := RsaDecrypt(decodedData, decodedKey)
	if err != nil {
		logger.Warn("decrypt failed", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	dec = PaycryptoCardSensitiveInfoDecryptedResponse{}
	err = json.Unmarshal(decryptedData, &dec)
	if err != nil {
		logger.Warn("unmarshal failed", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Get Physical Card Sensitive Info
func PaycryptoGetPhysicalCardInfo(ctx context.Context, form *PaycryptoGetCardSensitiveInfoForm) (res *PaycryptoResponse[PaycryptoCardSensitiveInfoResponse], dec PaycryptoCardSensitiveInfoDecryptedResponse, err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetPhysicalCardInfo()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("physical card info response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoCardSensitiveInfoResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	decodedKey, err := base64.StdEncoding.DecodeString(Config.Paycrypto.PrivateKey)
	if err != nil {
		logger.Warn("decode failed", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	if len(res.Result.EncryptData) == 0 {
		return
	}

	decodedData, err := base64.StdEncoding.DecodeString(res.Result.EncryptData[0])
	if err != nil {
		logger.Warn("decode failed", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	decryptedData, err := RsaDecrypt(decodedData, decodedKey)
	if err != nil {
		logger.Warn("decrypt failed", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	dec = PaycryptoCardSensitiveInfoDecryptedResponse{}
	err = json.Unmarshal(decryptedData, &dec)
	if err != nil {
		logger.Warn("unmarshal failed", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Get Authorized Transactions Params
type PaycryptoGetAuthorizedTransactionsParams struct {
	CardNo string `json:"card_no"`
}

// Authorized Transaction Response - similar to Card Transaction Record
type PaycryptoAuthorizedTransactionResponse struct {
	TransactionDate PaycryptoDate   `json:"transaction_date"`
	PostingDate     PaycryptoDate   `json:"posting_date"`
	TxID            string          `json:"tx_id"`
	Description     string          `json:"description"`
	Debit           decimal.Decimal `json:"debit"`
	DebitUSD        decimal.Decimal `json:"debit_usd"`
	Credit          decimal.Decimal `json:"credit"`
	CreditUSD       decimal.Decimal `json:"credit_usd"`
	Fee             decimal.Decimal `json:"fee"`
	TxCurrency      string          `json:"tx_currency,omitempty"`
	TxAmount        decimal.Decimal `json:"tx_amount,omitempty"`
}

// Get Authorized Transactions
func PaycryptoGetAuthorizedTransactions(ctx context.Context, params *PaycryptoGetAuthorizedTransactionsParams) (res *PaycryptoResponse[[]PaycryptoAuthorizedTransactionResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetAuthorizedTransactions()

	// Build query string
	queryValues := url.Values{}
	queryValues.Add("card_no", params.CardNo)
	requestQueryStr := queryValues.Encode()

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("authorized transactions response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[[]PaycryptoAuthorizedTransactionResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Update Card Limit Form
type PaycryptoUpdateCardLimitForm struct {
	CardNo           string `json:"card_no"`
	MaxAmountSingle  string `json:"max_amount_single"`
	MaxAmountDaily   string `json:"max_amount_daily"`
	MaxAmountMonthly string `json:"max_amount_monthly"`
}

// Card Limit Response
type PaycryptoCardLimitResponse struct {
	CardNo           string          `json:"card_no"`
	CardType         string          `json:"card_type"`
	MaxAmountSingle  decimal.Decimal `json:"max_amount_single"`
	MaxAmountDaily   decimal.Decimal `json:"max_amount_daily"`
	MaxAmountMonthly decimal.Decimal `json:"max_amount_monthly"`
	AvailableBalance decimal.Decimal `json:"available_balance"`
}

// Update Card Limit
func PaycryptoUpdateCardLimit(ctx context.Context, form *PaycryptoUpdateCardLimitForm) (res *PaycryptoResponse[PaycryptoCardLimitResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoUpdateCardLimit()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("update card limit response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoCardLimitResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Get Card Limit Params
type PaycryptoGetCardLimitParams struct {
	CardNo string `json:"card_no"`
}

func PaycryptoGetCardLimit(ctx context.Context, params *PaycryptoGetCardLimitParams) (res *PaycryptoResponse[PaycryptoCardLimitResponse], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoGetCardLimit()

	// Build query string
	queryValues := url.Values{}
	queryValues.Add("card_no", params.CardNo)
	requestQueryStr := queryValues.Encode()

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, nil)

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s]", url)
	resBody, _, resCode, err := HttpDo(ctx, method, url, "", header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("card limit response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("get failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[PaycryptoCardLimitResponse]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Freeze Card Form
type PaycryptoFreezeCardForm struct {
	CardNo string `json:"card_no"`
}

// Freeze Card
func PaycryptoFreezeCard(ctx context.Context, form *PaycryptoFreezeCardForm) (res *PaycryptoResponse[map[string]interface{}], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoFreezeCard()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("freeze card response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[map[string]interface{}]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Unfreeze Card Form
type PaycryptoUnfreezeCardForm struct {
	CardNo string `json:"card_no"`
}

// Unfreeze Card
func PaycryptoUnfreezeCard(ctx context.Context, form *PaycryptoUnfreezeCardForm) (res *PaycryptoResponse[map[string]interface{}], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoUnfreezeCard()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("unfreeze card response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[map[string]interface{}]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Init Physical Card Password Form
type PaycryptoInitPhysicalCardPasswordForm struct {
	CardNo     string `json:"card_no"`
	Pin        string `json:"pin"`
	ConfirmPin string `json:"confirm_pin"`
}

// Init Physical Card Password
func PaycryptoInitPhysicalCardPassword(ctx context.Context, form *PaycryptoInitPhysicalCardPasswordForm) (res *PaycryptoResponse[bool], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoInitPhysicalCardPassword()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("init physical card password response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[bool]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

// Reset Physical Card Password Form
type PaycryptoResetPhysicalCardPasswordForm struct {
	CardNo     string `json:"card_no"`
	Pin        string `json:"pin"`
	ConfirmPin string `json:"confirm_pin"`
}

// Reset Physical Card Password
func PaycryptoResetPhysicalCardPassword(ctx context.Context, form *PaycryptoResetPhysicalCardPasswordForm) (res *PaycryptoResponse[bool], err error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	method, host, requestPath := urlPaycryptoResetPhysicalCardPassword()
	requestQueryStr := ""

	data, err := json.Marshal(form)
	if err != nil {
		logger.Errorf("marshal failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	dataForSign, err := PaycryptoConvertRequestBodyToSignString(data)
	if err != nil {
		logger.Errorf("convert failed: %v", err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	sign := PaycryptoHmacSHA256Base64Sign(timestamp, method, requestPath, requestQueryStr, Config.Paycrypto.APIKey, Config.Paycrypto.APISecret, []byte(dataForSign))

	authorizationStr := strings.Join([]string{
		"Railone",
		Config.Paycrypto.APIKey,
		timestamp,
		sign,
	}, paycryptoSignSeparator)

	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", authorizationStr)
	header.Add("Access-Passphrase", Config.Paycrypto.APIPassphrase)

	url := host + requestPath
	if requestQueryStr != "" {
		url += "?" + requestQueryStr
	}
	logger.Infof("path: [%s], req: [%s]", url, data)
	resBody, _, resCode, err := HttpDo(ctx, method, url, string(data), header)
	if err != nil {
		logger.Warn("http failed, ", err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("reset physical card password response body: %s", string(resBody))

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	logger.Infof("res: [%s]", resBody)

	res = &PaycryptoResponse[bool]{}
	if err = json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		err = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}

	return
}

func paycryptoIterDelete(m map[string]interface{}, key string) {
	for k, v := range m {
		switch v.(type) {
		case string:
			if v.(string) == "N/A" {
				delete(m, k)
			}
		case map[string]interface{}:
			paycryptoIterDelete(v.(map[string]interface{}), key)
		case []interface{}:
			for _, vv := range v.([]interface{}) {
				switch vv.(type) {
				case map[string]interface{}:
					paycryptoIterDelete(vv.(map[string]interface{}), key)
				}
			}
		}
	}
	return
}
