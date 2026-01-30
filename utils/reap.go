package utils

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"time"

	"shared-modules/common"
	"shared-modules/logger"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ResidentialAddress struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	Line3      string `json:"line3,omitempty"` //可能為空
	Line4      string `json:"line4,omitempty"` //可能為空
	Line5      string `json:"line5,omitempty"` //可能為空
	Country    string `json:"country"`
	PostalCode string `json:"postalCode,omitempty"` //可能為空
	City       string `json:"city"`
}

type KYC struct {
	ResidentialAddress ResidentialAddress   `json:"residentialAddress"`
	IDDocumentType     common.ReapIdDocType `json:"idDocumentType"`
	FirstName          string               `json:"firstName"`
	LastName           string               `json:"lastName"`
	DOB                string               `json:"dob"`
	IDDocumentNumber   string               `json:"idDocumentNumber"`
}

type OTPPhoneNumber struct {
	DialCode    int    `json:"dialCode"`
	PhoneNumber string `json:"phoneNumber"`
}

type Meta struct {
	ID             string         `json:"id"`
	Email          string         `json:"email,omitempty"` //可能為空
	OTPPhoneNumber OTPPhoneNumber `json:"otpPhoneNumber"`
}

type TopUpWallet struct {
	Currency string `json:"currency,omitempty"` //可能為空
}

type ReapFreezeCardForm struct {
	Freeze bool `json:"freeze"`
}

type ShowCardPANHTMLForm struct {
	PublicKey     string `json:"publicKey,omitempty"`
	StylesheetURL string `json:"stylesheetUrl,omitempty"`
	CopyPAN       *bool  `json:"copyPan,omitempty"`
}

type CreateCardForm struct {
	CardType          common.ReapCardType     `json:"cardType"`
	CustomerType      common.ReapConsumerType `json:"customerType"`
	KYC               *KYC                    `json:"kyc"`
	Meta              *Meta                   `json:"meta"`
	TopUpWallet       *TopUpWallet            `json:"topUpWallet,omitempty"` //可能為空
	SpendLimit        decimal.Decimal         `json:"spendLimit,omitempty"`  //可能為空
	ExpiryDate        string                  `json:"expiryDate,omitempty"`  //可能為空
	PreferredCardName string                  `json:"preferredCardName"`
	SecondaryCardName string                  `json:"secondaryCardName,omitempty"` //可能為空
	CardDesign        string                  `json:"cardDesign,omitempty"`        //可能為空
}

type RetrieveCardVO struct {
	ID         string `json:"id,omitempty"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
	Parameter  string `json:"parameter,omitempty"`

	// 卡片核心資訊
	CardName            string             `json:"cardName,omitempty"`
	SecondaryCardName   *string            `json:"secondaryCardName,omitempty"`
	Last4               string             `json:"last4,omitempty"`
	AvailableCredit     string             `json:"availableCredit,omitempty"`
	Status              string             `json:"status,omitempty"`
	StatusReason        string             `json:"statusReason,omitempty"`
	CardType            string             `json:"cardType,omitempty"`
	PhysicalCardStatus  string             `json:"physicalCardStatus,omitempty"`
	ShippingAddress     *ShippingAddressVO `json:"shippingAddress,omitempty"`
	SpendControl        *SpendControlVO    `json:"spendControl,omitempty"`
	CardDesign          string             `json:"cardDesign,omitempty"`
	ShippingInformation ShippingInfoVO     `json:"shippingInformation,omitempty"`
	ThreeDSForwarding   bool               `json:"threeDSForwarding,omitempty"`
	Meta                MetaVO             `json:"meta,omitempty"`
}

type SpendControlVO struct {
	SpendControlAmount SpendControlAmountVO `json:"spendControlAmount,omitempty"`
	SpendControlCap    SpendControlCapVO    `json:"spendControlCap,omitempty"`
	ATMControl         ATMControlVO         `json:"atmControl,omitempty"`
}

type SpendControlAmountVO struct {
	DailySpent   decimal.Decimal `json:"dailySpent,omitempty"`
	WeeklySpent  decimal.Decimal `json:"weeklySpent,omitempty"`
	MonthlySpent decimal.Decimal `json:"monthlySpent,omitempty"`
	YearlySpent  decimal.Decimal `json:"yearlySpent,omitempty"`
	AllTimeSpent decimal.Decimal `json:"allTimeSpent,omitempty"`
}

type SpendControlCapVO struct {
	TransactionLimit decimal.Decimal `json:"transactionLimit,omitempty"`
	DailyLimit       decimal.Decimal `json:"dailyLimit,omitempty"`
	WeeklyLimit      decimal.Decimal `json:"weeklyLimit,omitempty"`
	MonthlyLimit     decimal.Decimal `json:"monthlyLimit,omitempty"`
	YearlyLimit      decimal.Decimal `json:"yearlyLimit,omitempty"`
	AllTimeLimit     decimal.Decimal `json:"allTimeLimit,omitempty"`
}

type ATMControlVO struct {
	DailyFrequency    decimal.Decimal `json:"dailyFrequency,omitempty"`
	MonthlyFrequency  decimal.Decimal `json:"monthlyFrequency,omitempty"`
	DailyWithdrawal   decimal.Decimal `json:"dailyWithdrawal,omitempty"`
	MonthlyWithdrawal decimal.Decimal `json:"monthlyWithdrawal,omitempty"`
}

type ShippingInfoVO struct {
	BulkShippingID *string `json:"bulkShippingID,omitempty"`
	SKU            *string `json:"sku,omitempty"`
}

type ShippingAddressVO struct {
	Line1      string `json:"line1,omitempty"`
	Line2      string `json:"line2,omitempty"`
	Zone       string `json:"zone,omitempty"`
	City       string `json:"city,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
	Country    string `json:"country,omitempty"`
}

type MetaVO struct {
	ID             string         `json:"id,omitempty"`
	Email          string         `json:"email,omitempty"`
	OTPPhoneNumber OTPPhoneNumber `json:"otpPhoneNumber,omitempty"`
}

type CreateCardVO struct {
	ID         string `json:"id,omitempty"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
	Parameter  string `json:"parameter,omitempty"`
}

type FreezeCardVO struct {
	Freeze bool `json:"freeze"`
}

type DeleteCardVO struct {
	RemainingCredit decimal.Decimal `json:"remainingCredit"`
	AvailableCredit decimal.Decimal `json:"availableCredit"`
}

type ReapCommonVO struct {
	Code       common.ReapCode `json:"code"`
	Message    string          `json:"message"`
	Parameter  string          `json:"parameter"`
	StatusCode int             `json:"statusCode"`
}

type ShippingAddressForm struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2,omitempty"`
	Zone       string `json:"zone,omitempty"`
	City       string `json:"city"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
}

type ShipCardForm struct {
	ShippingAddress      ShippingAddressForm `json:"shippingAddress"`
	RecipientFirstName   string              `json:"recipientFirstName"`
	RecipientLastName    string              `json:"recipientLastName"`
	RecipientPhoneNumber string              `json:"recipientPhoneNumber"`
	RecipientDialCode    int                 `json:"recipientDialCode"`
	RecipientEmail       string              `json:"recipientEmail"`
	CardDesign           string              `json:"cardDesign"`
}

type ShipCardVO struct {
	Code       string `json:"code,omitempty"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode,omitempty"`
	Parameter  string `json:"parameter,omitempty"`
}

type GetShippingOrderVO struct {
	ID             string                    `json:"id,omitempty"`
	SKU            string                    `json:"sku,omitempty"`
	CourierID      string                    `json:"courierID,omitempty"`
	CourierName    string                    `json:"courierName,omitempty"`
	TrackingNumber string                    `json:"trackingNumber,omitempty"`
	TrackingURL    string                    `json:"trackingUrl,omitempty"`
	Status         common.ReapShippingStatus `json:"status,omitempty"`
	Code           string                    `json:"code,omitempty"`
	Message        string                    `json:"message,omitempty"`
	StatusCode     int                       `json:"statusCode,omitempty"`
	Parameter      string                    `json:"parameter,omitempty"`
}

type BulkShipCardForm struct {
	BulkshipID                 *uuid.UUID                              `json:"bulkshipID,omitempty"`
	BulkShippingAddress        BulkShipCardBulkShippingAddress         `json:"bulkShippingAddress"`
	RecipientInformation       BulkShipCardRecipientInformation        `json:"recipientInformation"`
	DestinationShippingDetails []BulkShipCardDestinationShippingDetail `json:"destinationShippingDetails"`
}

type BulkShipCardBulkShippingAddress struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	Zone       string `json:"zone"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

type BulkShipCardRecipientInformation struct {
	RecipientFirstName   string `json:"recipientFirstName"`
	RecipientLastName    string `json:"recipientLastName"`
	RecipientPhoneNumber string `json:"recipientPhoneNumber"`
	RecipientDialCode    string `json:"recipientDialCode"`
	RecipientEmail       string `json:"recipientEmail"`
}

type BulkShipCardDestinationShippingDetail struct {
	CardID          string                                               `json:"cardID"`
	ShippingAddress BulkShipCardDestinationShippingDetailShippingAddress `json:"shippingAddress"`
	CardDesignID    string                                               `json:"cardDesignID"`
}

type BulkShipCardDestinationShippingDetailShippingAddress struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	Zone       string `json:"zone"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

type BulkShipCardVO struct {
	BulkshipId     uuid.UUID          `json:"bulkshipId"`
	SuccessCount   int                `json:"successCount"`
	FailedCount    int                `json:"failedCount"`
	RemainingSlots int                `json:"remainingSlots"`
	Errors         BulkShipCardErrors `json:"errors"`
	CutoffDate     string             `json:"cutoffDate"`
}

type BulkShipCardErrors []BulkShipCardError

func (b BulkShipCardErrors) Get(cardId string) *BulkShipCardError {
	for _, v := range b {
		if v.CardID == cardId {
			tv := v
			return &tv
		}
	}
	return nil
}

type BulkShipCardError struct {
	CardID  string `json:"cardID"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

type ActivateCardForm struct {
	Code string `json:"code"`
}

type ActivateCardVO struct {
	Code       string `json:"code,omitempty"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode,omitempty"`
	Parameter  string `json:"parameter,omitempty"`
}

type UpdatePINForm struct {
	PIN string `json:"pin"`
}

type UpdatePINVO struct {
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	PIN        string `json:"pin,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
	Parameter  string `json:"parameter,omitempty"`
}

type TransactionResponse struct {
	Items               []TransactionItem   `json:"items"`
	CardTransactionMeta CardTransactionMeta `json:"meta"`
}

type TransactionItem struct {
	ID                  string         `json:"id"`
	CardID              string         `json:"card_id"`
	CreatedAt           time.Time      `json:"created_at"`
	ClearedAt           time.Time      `json:"cleared_at"`
	MerchantData        MerchantData   `json:"merchant_data"`
	Category            string         `json:"category"`
	Fees                Fees           `json:"fees"`
	IsCredit            bool           `json:"is_credit"`
	BillAmount          string         `json:"bill_amount"`
	BillCurrency        string         `json:"bill_currency"`
	TransactionAmount   string         `json:"transaction_amount"`
	TransactionCurrency string         `json:"transaction_currency"`
	ConversionRate      string         `json:"conversion_rate"`
	Status              string         `json:"status"`
	DeclineReason       map[string]any `json:"decline_reason"` // 若永遠是空物件可改為 json.RawMessage 或 struct{}
	Wallet              any            `json:"wallet"`         // 若始終為 null 可用 *Wallet 或 interface{}
	MccPaddingAmount    string         `json:"mcc_padding_amount"`
	Channel             string         `json:"channel"`
	LifecycleEvents     []Lifecycle    `json:"lifecycle_events"`
}

type MerchantData struct {
	MerchantID       string `json:"merchant_id"`
	MerchantName     string `json:"merchant_name"`
	MerchantCity     string `json:"merchant_city"`
	MerchantPostCode string `json:"merchant_post_code"`
	MerchantState    string `json:"merchant_state"`
	MerchantCountry  string `json:"merchant_country"`
	MccCategory      string `json:"mcc_category"`
	MccCode          string `json:"mcc_code"`
}

type Fees struct {
	AtmFees string `json:"atm_fees"`
	FxFees  string `json:"fx_fees"`
}

type Lifecycle struct {
	LifecycleEventID    string               `json:"lifecycle_event_id"`
	BillAmount          decimal.Decimal      `json:"bill_amount"`
	TransactionAmount   decimal.Decimal      `json:"transaction_amount"`
	BillCurrency        string               `json:"bill_currency"`
	TransactionCurrency string               `json:"transaction_currency"`
	CreatedAt           time.Time            `json:"created_at"`
	Type                common.ReapEventName `json:"type"`
	AffectBalance       bool                 `json:"affect_balance"`
	MccPaddingAmount    string               `json:"mcc_padding_amount"`
}

type CardTransactionMeta struct {
	TotalItems   int `json:"totalItems"`
	ItemCount    int `json:"itemCount"`
	ItemsPerPage int `json:"itemsPerPage"`
	TotalPages   int `json:"totalPages"`
	CurrentPage  int `json:"currentPage"`
}

type ReapBillTaskVO struct {
	Message     string `json:"message"`
	ReferenceID string `json:"reference_id"`
}

var (
	urlCreateCard  = func() (string, string) { return "POST", Config.Reap.BaseURL + "cards" }
	urlShowCardPAN = func(cardID string) (string, string) {
		return "POST", Config.Reap.BaseURL + "cards/" + cardID + "/reveal"
	}
	urlShowCardPANHTML = func(cardID string) (string, string) {
		return "POST", Config.Reap.BaseURL + "cards/" + cardID + "/reveal-html"
	}
	urlFreezeCard = func(cardID string) (string, string) {
		return "PUT", Config.Reap.BaseURL + "cards/" + cardID + "/status"
	}
	urlDeleteCard = func(cardID string) (string, string) {
		return "PUT", Config.Reap.BaseURL + "cards/" + cardID
	}
	urlThreedsAnswer = func(cardID string, initiateActionId string) (string, string) {
		return "POST", Config.Reap.BaseURL + "cards/" + cardID + "/3ds-answer/" + initiateActionId
	}
	urlUpdateThreedsMethod = func(cardID string) (string, string) {
		return "PUT", Config.Reap.BaseURL + "cards/" + cardID + "/3ds-forwarding"
	}
	urlShipPhysicalCard = func(cardID string) (string, string) {
		return "POST", Config.Reap.BaseURL + "cards/" + cardID + "/ship"
	}
	urlGetShippingOrder = func(cardID string) (string, string) {
		return "GET", Config.Reap.BaseURL + "cards/" + cardID + "/orders"
	}
	urlActivateCard = func(cardID string) (string, string) {
		return "POST", Config.Reap.BaseURL + "cards/" + cardID + "/activate"
	}
	urlGetAccountAmount = func() (string, string) {
		return "GET", Config.Reap.BaseURL + "account/balance"
	}
	urlUpdatePIN = func(cardID string) (string, string) {
		return "PUT", Config.Reap.BaseURL + "cards/" + cardID + "/pin"
	}
	urlRetrieveCard = func(cardID string) (string, string) {
		return "GET", Config.Reap.BaseURL + "cards/" + cardID
	}
	urlBlockCard = func(cardID string) (string, string) {
		return "PUT", Config.Reap.BaseURL + "cards/" + cardID + "/block"
	}
	urlUnblockCard = func(cardID string) (string, string) {
		return "PUT", Config.Reap.BaseURL + "cards/" + cardID + "/unblock"
	}
	urlGetSingleTransaction = func(transactionID string) (string, string) {
		return "GET", Config.Reap.BaseURL + "transactions/" + transactionID
	}
	urlGetCardTransactions = func(cardID, startDate, endDate string, page string) (string, string) {
		return "GET", Config.Reap.BaseURL + "cards/" + cardID + "/transactions" + "?page=" + page + "&limit=50&toDate=" + endDate + "&fromDate=" + startDate
	}
	urlGetBillReport = func(reportType common.ReapReportRange, date string) (string, string) {
		var fileType string
		switch reportType {
		case common.REAP_REPORT_RANGE_DAILY:
			fileType = "daily"
		case common.REAP_REPORT_RANGE_MONTHLY:
			fileType = "monthly"
		default:
			return "", ""
		}
		url := fmt.Sprintf("%sfiles/%s/?date=%s", Config.Reap.BaseURL, fileType, date)
		return "GET", url
	}
	urlBulkShipPhysicalCard = func() (string, string) {
		return "POST", Config.Reap.BaseURL + "cards/bulk-ship"
	}
)

func ReapRetireveCard(ctx context.Context, issueID string) (res *RetrieveCardVO, rErr error) {

	_, url := urlRetrieveCard(issueID)

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")
	header.Add("x-reap-api-key", Config.Reap.APIKey)
	header.Add("Accept-Version", "v1.0")

	resBody, _, resCode, err := HttpGet(ctx, url, header)
	if err != nil {
		logger.Warn("post failed", err)
		rErr = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}
	logger.Infof("retrieve card response body: %s", string(resBody))

	res = &RetrieveCardVO{}
	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		rErr = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		rErr = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	return
}

func ReapCreateCard(ctx context.Context, form *CreateCardForm) (res *CreateCardVO, rErr error) {

	_, url := urlCreateCard()

	payload, err := json.Marshal(form)
	if err != nil {
		logger.Warn("marshal failed", err)
		rErr = NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		return
	}
	logger.Infof("form: %s", payload)

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")
	header.Add("x-reap-api-key", Config.Reap.APIKey)
	header.Add("Accept-Version", "v1.0")

	resBody, _, resCode, err := HttpPost(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		rErr = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}
	logger.Infof("create card response body: %s", string(resBody))

	res = &CreateCardVO{}
	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		rErr = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	if resCode < 200 || resCode >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, resCode, err)
		rErr = NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		return
	}

	return
}

func ReapFreezeCard(ctx context.Context, cardId string, form *ReapFreezeCardForm) (bool, error) {

	_, url := urlFreezeCard(cardId)

	payload, err := json.Marshal(form)
	if err != nil {
		logger.Warn("marshal failed", err)
		return false, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	logger.Infof("freeze card form: %s", payload)

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")
	header.Add("x-reap-api-key", Config.Reap.APIKey)
	header.Add("Accept-Version", "v1.0")

	resBody, _, code, err := HttpPut(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		return false, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	logger.Infof("freeze card response: %s", string(resBody))
	if code != 200 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, code, err)
		return false, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	res := &FreezeCardVO{}
	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return false, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	return res.Freeze, nil
}

func ReapDeleteCard(ctx context.Context, issueID string) (*DeleteCardVO, error) {

	_, url := urlDeleteCard(issueID)

	logger.Infof("delete card: %s", issueID)

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")
	header.Add("x-reap-api-key", Config.Reap.APIKey)
	header.Add("Accept-Version", "v1.0")

	resBody, _, code, err := HttpDelete(ctx, url, "", header)
	if err != nil {
		logger.Warn("post failed", err)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	logger.Infof("delete card response: %s", string(resBody))
	if code != 200 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	res := &DeleteCardVO{}
	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	return res, nil
}

type ShowCardPANVO struct {
	CardNumber string `json:"pan"`
	CVV        string `json:"cvv"`
	ExpiryDate string `json:"expiryDate"`
}

type ShowCardPANHTMLVO struct {
	AccessURL string `json:"accessUrl"`
}

type ResponseDataVO struct {
	Data string `json:"data"`
}

func ReapShowCardPAN(ctx context.Context, cardID string) (*ShowCardPANVO, error) {

	_, url := urlShowCardPAN(cardID)

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("x-reap-api-key", Config.Reap.APIKey)
	header.Add("Accept-Version", "v1.0")

	resBody, _, code, err := HttpPost(ctx, url, "", header)
	if err != nil {
		logger.Warn("post failed", err)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	if code < 200 || code >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	// logger.Infof("res: %s", resBody)

	res := &ResponseDataVO{}
	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	decodedKey, err := base64.StdEncoding.DecodeString(Config.Reap.PrivateKey)
	if err != nil {
		logger.Warn("decode failed", err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}
	decodedData, err := base64.StdEncoding.DecodeString(res.Data)
	if err != nil {
		logger.Warn("decode failed", err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}
	decryptedData, err := RsaDecryptOAEP(decodedData, decodedKey)
	if err != nil {
		logger.Warn("decrypt failed", err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	cardInfo := &ShowCardPANVO{}
	if err := json.Unmarshal(decryptedData, cardInfo); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, %v", decryptedData, err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	return cardInfo, nil
}

func ReapShowCardPANHTML(ctx context.Context, cardID string) (string, error) {

	_, url := urlShowCardPANHTML(cardID)

	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		logger.Warn("generate failed,", err)
		return "", NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	pubk, err := x509.MarshalPKIXPublicKey(key.Public().(*rsa.PublicKey))
	if err != nil {
		logger.Warn("marshal failed,", err)
		return "", NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	pubkPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pubk,
		},
	)
	pubkb64 := base64.StdEncoding.EncodeToString(pubkPEM)

	payload, err := json.Marshal(&ShowCardPANHTMLForm{
		PublicKey: pubkb64,
		CopyPAN:   Ptr(true),
	})
	if err != nil {
		logger.Warn("marshal failed", err)
		return "", NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	// logger.Infof("get card PAN HTML form: %s", payload)

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")
	header.Add("x-reap-api-key", Config.Reap.APIKey)
	header.Add("Accept-Version", "v1.0")

	resBody, _, code, err := HttpPost(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		return "", NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	if code < 200 || code >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, code, err)
		return "", NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	// logger.Infof("res: %s", resBody)

	res := &ShowCardPANHTMLVO{}
	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return "", NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	encData, err := base64.StdEncoding.DecodeString(res.AccessURL)
	if err != nil {
		logger.Warn("decode failed", err)
		return "", NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	prik, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		logger.Warn("marshal failed,", err)
		return "", NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	prikPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: prik,
		},
	)

	logger.Debugf("public key: [%s]", pubkPEM)
	logger.Debugf("private key: [%s]", prikPEM)

	accessURL, err := key.Decrypt(nil, encData, &rsa.OAEPOptions{
		Hash:    crypto.SHA1,
		MGFHash: crypto.SHA1,
	})
	if err != nil {
		logger.Warn("decrypt failed,", err)
		return "", NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	return string(accessURL[1 : len(accessURL)-1]), nil
}

func ReapShipCard(ctx context.Context, cardID string, form *ShipCardForm) (*ShipCardVO, error) {

	_, url := urlShipPhysicalCard(cardID)

	payload, err := json.Marshal(form)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	logger.Infof("ship card form: %s", payload)

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")
	header.Add("x-reap-api-key", Config.Reap.APIKey)
	header.Add("Accept-Version", "v1.0")

	resBody, _, code, err := HttpPost(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	logger.Infof("res: %s", resBody)

	res := &ShipCardVO{}
	if code < 200 || code >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, code, err)
		if err := json.Unmarshal(resBody, res); err != nil {
			return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		}
		return res, NewBusinessError(ctx, common.CODE_CARD_REAP_SHIP_FAILED)
	}

	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	return res, nil
}

func ReapGetShippingOrder(ctx context.Context, cardID string) (*GetShippingOrderVO, error) {

	_, url := urlGetShippingOrder(cardID)

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")
	header.Add("x-reap-api-key", Config.Reap.APIKey)
	header.Add("Accept-Version", "v1.0")

	resBody, _, code, err := HttpGet(ctx, url, header)
	if err != nil {
		logger.Warn("get failed", err)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	logger.Infof("res: %s", resBody)

	res := &GetShippingOrderVO{}
	if code < 200 || code >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, code, err)
		if err := json.Unmarshal(resBody, res); err != nil {
			return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		}
		return res, NewBusinessError(ctx, common.CODE_CARD_REAP_GET_SHIPPING_ORDER_FAILED)
	}

	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	return res, nil
}

func ReapBulkShipCard(ctx context.Context, form *BulkShipCardForm) (*BulkShipCardVO, error) {
	_, url := urlBulkShipPhysicalCard()
	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")
	header.Add("x-reap-api-key", Config.Reap.APIKey)
	header.Add("Accept-Version", "v1.0")
	payload, err := json.Marshal(form)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}
	logger.Infof("ship card form: %s", payload)
	resBody, _, code, err := HttpPost(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	logger.Infof("res: %s", resBody)
	res := &BulkShipCardVO{}
	if code < 200 || code >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, code, err)
		if err := json.Unmarshal(resBody, res); err != nil {
			return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		}
		return res, NewBusinessError(ctx, common.CODE_CARD_REAP_SHIP_FAILED)
	}
	if err := json.Unmarshal(resBody, res); err != nil {
		logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
		return res, nil
	}
	return res, nil
}

func ReapActivateCard(ctx context.Context, cardID string, form *ActivateCardForm) (*ActivateCardVO, error) {

	_, url := urlActivateCard(cardID)

	payload, err := json.Marshal(form)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	logger.Infof("ship card form: %s", payload)

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")
	header.Add("x-reap-api-key", Config.Reap.APIKey)
	header.Add("Accept-Version", "v1.0")

	resBody, _, code, err := HttpPost(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	logger.Infof("res: %s", resBody)

	res := &ActivateCardVO{}
	if code < 200 || code >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, code, err)
		if err := json.Unmarshal(resBody, res); err != nil {
			return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		}
		return res, NewBusinessError(ctx, common.CODE_CARD_REAP_ACTIVATE_CARD_FAILED)
	}

	return nil, nil
}

func ReapUpdatePIN(ctx context.Context, cardID string, form *UpdatePINForm) (*UpdatePINVO, error) {

	_, url := urlUpdatePIN(cardID)

	payload, err := json.Marshal(form)
	if err != nil {
		logger.Warn("marshal failed", err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	logger.Debugf("update PIN form: %s", payload)

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")
	header.Add("x-reap-api-key", Config.Reap.APIKey)
	header.Add("Accept-Version", "v1.0")

	resBody, _, code, err := HttpPut(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	logger.Debugf("res: %s", resBody)

	res := &UpdatePINVO{}

	if code < 200 || code >= 300 {
		if code == 400 {
			return res, NewBusinessError(ctx, common.CODE_REAP_INVALID_PIN_FORMAT)
		}
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, code, err)
		if err := json.Unmarshal(resBody, res); err != nil {
			return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
		}
		return res, NewBusinessError(ctx, common.CODE_CARD_REAP_UPDATE_PIN_FAILED)
	}

	return res, nil
}

type ReapUpdateThreedsMethodForm struct {
	BiometricEnroll bool `json:"biometricEnroll"`
	SmsEnroll       bool `json:"smsEnroll"`
}

type ReapUpdateThreedsMethodVO map[string]interface{}

func ReapUpdateThreedsMethod(ctx context.Context, issueId string, form *ReapUpdateThreedsMethodForm) error {

	_, url := urlUpdateThreedsMethod(issueId)

	payload, err := json.Marshal(form)
	if err != nil {
		logger.Warn("marshal failed", err)
		return NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}
	logger.Infof("form: %s", payload)

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")
	header.Add("x-reap-api-key", Config.Reap.APIKey)
	header.Add("Accept-Version", "v1.0")

	resBody, _, code, err := HttpPut(ctx, url, string(payload), header)
	if err != nil {
		logger.Warn("post failed", err)
		return NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	if code < 200 || code >= 300 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, code, err)
		return NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	return nil
}

type ReapThreedsAnswerForm struct {
	Approve bool `json:"approve"`
}

type ReapThreedsAnswerVO map[string]interface{}

func ReapThreedsAnswer(ctx context.Context, cardID string, initiateActionId string, approve bool) error {

	_, url := urlThreedsAnswer(cardID, initiateActionId)

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")
	header.Add("x-reap-api-key", Config.Reap.APIKey)
	header.Add("Accept-Version", "v1.0")

	body := map[string]interface{}{
		"approve": approve,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		logger.Warn("marshal failed", err)
		return NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	logger.Infof("answer 3ds url: %s", url)
	logger.Infof("answer 3ds bodyJSON: %s", string(bodyJSON))

	resBody, _, code, err := HttpPost(ctx, url, string(bodyJSON), header)
	if err != nil {
		logger.Warn("post failed", err)
		return NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	if code != 200 {
		logger.Warnf("post failed, resp body: %s, code: %d, %v", resBody, code, err)

		res := ReapThreedsAnswerVO{}
		if err := json.Unmarshal(resBody, &res); err != nil {
			logger.Warnf("unmarshal failed, resp body: %s, code: %d, %v", resBody, code, err)
			return NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
		}
		if res["code"] == string(common.REAP_RESPONSE_CODE_3DS_NOT_FOUND) {
			return NewBusinessError(ctx, common.CODE_REAP_THREEDS_NOTIFICATION_NOT_FOUND)
		}

		return NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	return nil
}

type AccountAmountVO struct {
	AvailableBalance    decimal.Decimal `json:"availableBalance"`
	AvailableToAllocate decimal.Decimal `json:"availableToAllocate"`
}

func ReapAccountAmount(ctx context.Context) (*AccountAmountVO, error) {

	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")
	header.Add("x-reap-api-key", Config.Reap.APIKey)
	header.Add("Accept-Version", "v1.0")

	_, url := urlGetAccountAmount()

	logger.Infof("get reap account amount url: %s", url)

	resBody, _, code, err := HttpGet(ctx, url, header)
	if err != nil {
		logger.Warn("get reap account amount failed", err)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	if code != 200 {
		logger.Warnf("Get failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	result := &AccountAmountVO{}
	if err := json.Unmarshal([]byte(resBody), result); err != nil {
		logger.Infof("Error parsing JSON response: %v", err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	return result, nil
}

func ReapBlockCard(ctx context.Context, issueId string, blockReason string) error {
	_, url := urlBlockCard(issueId)
	header := reapHeader(Config.Reap.APIKey)
	body := map[string]interface{}{"reason": blockReason}
	jsonBody, _ := json.Marshal(body)

	resBody, _, code, err := HttpPut(ctx, url, string(jsonBody), header)
	if err != nil {
		logger.Warn("put block failed", err)
		return NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	if code < 200 || code >= 300 {
		logger.Warnf("put failed, resp body: %s, code: %d, %v", resBody, code, err)
		if len(resBody) != 0 {
			vo := &ReapCommonVO{}
			if err := json.Unmarshal([]byte(resBody), vo); err == nil {
				switch vo.Code {
				case common.REAP_CODE_CARD_ALREADY_BLOCKED:
					return nil
				}
			}
		}
		return NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	return nil
}

func ReapUnblockCard(ctx context.Context, issueId string) error {
	_, url := urlUnblockCard(issueId)
	header := reapHeader(Config.Reap.APIKey)

	resBody, _, code, err := HttpPut(ctx, url, "", header)
	if err != nil {
		logger.Warn("put unblock failed", err)
		return NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	if code < 200 || code >= 300 {
		logger.Warnf("put failed, resp body: %s, code: %d, %v", resBody, code, err)
		return NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	return nil
}

func GetSingleTransaction(ctx context.Context, transactionID string) (*TransactionItem, error) {
	header := reapHeader(Config.Reap.APIKey)
	_, url := urlGetSingleTransaction(transactionID)

	logger.Infof("get reap single transaction url: %s", url)

	resBody, _, code, err := HttpGet(ctx, url, header)
	if err != nil {
		logger.Warn("get reap single transaction failed", err)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	if code != 200 {
		logger.Warnf("Get failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	logger.Infof("get reap single transaction response: %s", resBody)

	result := &TransactionItem{}
	if err := json.Unmarshal([]byte(resBody), result); err != nil {
		logger.Infof("Error parsing JSON response: %v", err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}

	return result, nil
}

func GetCardTransactions(ctx context.Context, cardID, startDate, endDate, page string) (map[string]TransactionItem, int, error) {
	header := reapHeader(Config.Reap.APIKey)
	_, url := urlGetCardTransactions(cardID, startDate, endDate, page)

	logger.Infof("get card transactions url: %s", url)

	resBody, _, code, err := HttpGet(ctx, url, header)
	if err != nil {
		logger.Warn("get card transactions failed", err)
		return nil, 0, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	if code != 200 {
		logger.Warnf("Get failed, resp body: %s, code: %d, %v", resBody, code, err)
		return nil, 0, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	var txResp TransactionResponse
	if err := json.Unmarshal([]byte(resBody), &txResp); err != nil {
		logger.Warn("failed to unmarshal transaction response:", err)
		return nil, 0, NewBusinessError(ctx, common.CODE_MALFORMED_DATA)
	}

	result := make(map[string]TransactionItem)
	for _, item := range txResp.Items {
		result[item.ID] = item
	}

	return result, txResp.CardTransactionMeta.TotalPages, nil
}

func ReapGenerateReport(ctx context.Context, reportType common.ReapReportRange, date string) (*ReapBillTaskVO, error) {
	_, url := urlGetBillReport(reportType, date)
	header := reapHeader(Config.Reap.APIKey)

	resBody, _, code, err := HttpGet(ctx, url, header)
	if err != nil {
		logger.Warn("get report failed", err)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}
	if code < 200 || code >= 300 {
		logger.Warnf("get report failed, resp body: %s, code: %d", resBody, code)
		return nil, NewBusinessError(ctx, common.CODE_EXTERNAL_API_ERROR)
	}

	result := &ReapBillTaskVO{}
	if err := json.Unmarshal([]byte(resBody), result); err != nil {
		logger.Infof("Error parsing JSON response: %v", err)
		return nil, NewBusinessError(ctx, common.CODE_SYSTEM_ERROR)
	}
	return result, nil
}

// ReapAuthDeclineCodeMap 三方回傳的 DeclinedCode -> 系統內部代碼
var ReapAuthDeclineCodeMap = map[common.DeclinedCode]int{
	common.DECLINED_CODE_MERCHANT_NOT_ACCEPT_CARD:         common.CODE_REAP_DECLINE_CODE_MERCHANT_NOT_ACCEPT_CARD,
	common.DECLINED_CODE_DECLINED_BY_CLIENT:               common.CODE_REAP_DECLINE_CODE_DECLINED_BY_CLIENT,
	common.DECLINED_CODE_CARD_STATUS_LOST:                 common.CODE_REAP_DECLINE_CODE_CARD_STATUS_LOST,
	common.DECLINED_CODE_CARD_DESTROYED:                   common.CODE_REAP_DECLINE_CODE_CARD_DESTROYED,
	common.DECLINED_CODE_INSUFFICIENT_FUNDS:               common.CODE_REAP_DECLINE_CODE_INSUFFICIENT_FUNDS,
	common.DECLINED_CODE_INSUFFICIENT_FUNDS_REAL_TIME:     common.CODE_REAP_DECLINE_CODE_INSUFFICIENT_FUNDS,
	common.DECLINED_CODE_CARD_EXPIRED:                     common.CODE_REAP_DECLINE_CODE_CARD_EXPIRED,
	common.DECLINED_CODE_INCORRECT_PIN:                    common.CODE_REAP_DECLINE_CODE_INCORRECT_PIN,
	common.DECLINED_CODE_TRANSACTION_NOT_PERMITTED:        common.CODE_REAP_DECLINE_CODE_TRANSACTION_NOT_PERMITTED,
	common.DECLINED_CODE_POS_TRANSACTION_NOT_ALLOWED:      common.CODE_REAP_DECLINE_CODE_POS_TRANSACTION_NOT_ALLOWED,
	common.DECLINED_CODE_ATM_LIMIT_EXCEEDED:               common.CODE_REAP_DECLINE_CODE_ATM_LIMIT_EXCEEDED,
	common.DECLINED_CODE_WITHDRAWAL_AMOUNT_LIMIT_EXCEEDED: common.CODE_REAP_DECLINE_CODE_WITHDRAWAL_AMOUNT_LIMIT_EXCEEDED,
	common.DECLINED_CODE_WITHDRAWAL_NOT_PERMITTED:         common.CODE_REAP_DECLINE_CODE_WITHDRAWAL_NOT_PERMITTED,
	common.DECLINED_CODE_WITHDRAWAL_FREQUENCY_EXCEEDED:    common.CODE_REAP_DECLINE_CODE_WITHDRAWAL_FREQUENCY_EXCEEDED,
	common.DECLINED_CODE_INVALID_CARD_STATUS:              common.CODE_REAP_DECLINE_CODE_INVALID_CARD_STATUS,
	common.DECLINED_CODE_EXCEEDING_PIN_ATTEMPTS:           common.CODE_REAP_DECLINE_CODE_EXCEEDING_PIN_ATTEMPTS,
	common.DECLINED_CODE_ISSUER_NOT_PARTICIPATE:           common.CODE_REAP_DECLINE_CODE_ISSUER_NOT_PARTICIPATE,
	common.DECLINED_CODE_TRANSACTION_TIMEOUT:              common.CODE_REAP_DECLINE_CODE_TRANSACTION_TIMEOUT,
	common.DECLINED_CODE_PIN_VERIFICATION_FAILED:          common.CODE_REAP_DECLINE_CODE_PIN_VERIFICATION_FAILED,
	common.DECLINED_CODE_ISSUER_INOPERATIVE:               common.CODE_REAP_DECLINE_CODE_ISSUER_INOPERATIVE,
	common.DECLINED_CODE_TRANSACTION_RECORD_DISCREPANCY:   common.CODE_REAP_DECLINE_CODE_TRANSACTION_RECORD_DISCREPANCY,
	common.DECLINED_CODE_UNEXPECTED_ERROR:                 common.CODE_REAP_DECLINE_CODE_UNEXPECTED_ERROR,
	common.DECLINED_CODE_PROCESSING_ERROR:                 common.CODE_REAP_DECLINE_CODE_PROCESSING_ERROR,
	common.DECLINED_CODE_INVALID_CVV:                      common.CODE_REAP_DECLINE_CODE_INVALID_CVV,
	common.DECLINED_CODE_VERIFICATION_FAILED:              common.CODE_REAP_DECLINE_CODE_VERIFICATION_FAILED,
	common.DECLINED_CODE_CARD_FROZEN:                      common.CODE_REAP_DECLINE_CODE_CARD_FROZEN,
	common.DECLINED_CODE_EXCEEDING_CVV_ATTEMPTS:           common.CODE_REAP_DECLINE_CODE_EXCEEDING_CVV_ATTEMPTS,
	common.DECLINED_CODE_EXCEEDING_EXPIRY_DATE_ATTEMPTS:   common.CODE_REAP_DECLINE_CODE_EXCEEDING_EXPIRY_DATE_ATTEMPTS,
	common.DECLINED_CODE_CARD_NOT_ACTIVATED:               common.CODE_REAP_DECLINE_CODE_CARD_NOT_ACTIVATED,
	common.DECLINED_CODE_WITHDRAWAL_NOT_ALLOWED:           common.CODE_REAP_DECLINE_CODE_WITHDRAWAL_NOT_ALLOWED,
	common.DECLINED_CODE_DAILY_SPENDING_LIMIT_EXCEEDED:    common.CODE_REAP_DECLINE_CODE_DAILY_SPENDING_LIMIT_EXCEEDED,
	common.DECLINED_CODE_SPENDING_FREQUENCY_EXCEEDED:      common.CODE_REAP_DECLINE_CODE_SPENDING_FREQUENCY_EXCEEDED,
	common.DECLINED_CODE_CARD_BLOCKED:                     common.CODE_REAP_DECLINE_CODE_CARD_BLOCKED,
	common.DECLINED_CODE_CATEGORY_INSUFFICIENT_FUNDS:      common.CODE_REAP_DECLINE_CODE_CATEGORY_INSUFFICIENT_FUNDS,
	common.DECLINED_CODE_BUDGET_INSUFFICIENT_FUNDS:        common.CODE_REAP_DECLINE_CODE_FUND_ISSUE,
	common.DECLINED_CODE_BUDGET_FROZEN:                    common.CODE_REAP_DECLINE_CODE_BUDGET_FROZEN,
}

// GetReapDeclineCode 檢查三方回傳的 code 是否有定義
// 沒定義則統一回傳 CODE_REAP_DECLINE_CODE_DEFAULT
func GetReapDeclineCode(code string) int {
	authCode := common.DeclinedCode(code)
	if v, ok := ReapAuthDeclineCodeMap[authCode]; ok {
		return v
	}
	return common.CODE_REAP_DECLINE_CODE_DEFAULT
}

func reapHeader(apiKey string) http.Header {
	header := make(http.Header)
	header.Add("accept", "application/json")
	header.Add("content-type", "application/json")
	header.Add("x-reap-api-key", apiKey)
	header.Add("Accept-Version", "v1.0")
	return header
}
