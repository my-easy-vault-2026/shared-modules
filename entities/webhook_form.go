package entities

import (
	"time"

	"shared-modules/common"

	"github.com/shopspring/decimal"
)

type FeesForm struct {
	FXFees  decimal.Decimal `json:"fx_fees"`
	ATMFees decimal.Decimal `json:"atm_fees"`
}

type MerchantDataForm struct {
	MCCCode          string  `json:"mcc_code"`
	MerchantID       string  `json:"merchant_id"`
	MCCCategory      string  `json:"mcc_category"`
	MerchantCity     string  `json:"merchant_city"`
	MerchantName     string  `json:"merchant_name"`
	MerchantState    *string `json:"merchant_state,omitempty"`
	MerchantCountry  string  `json:"merchant_country"`
	MerchantPostCode *string `json:"merchant_post_code,omitempty"`
}

type DeclineReasonForm struct {
	Description  string `json:"description"`
	ResponseCode string `json:"response_code"`
}

type EventDataType interface {
	map[string]interface{} | TransactionEventForm | AuthorizationEventForm | *AuthorizationEventForm | NotificationEventForm | CardEventForm | ReapBillReportEventForm | ReapShippingOrderConfirmationEventForm
}

type AuthorizationEventForm struct {
	MCC                 string            `json:"mcc"`
	Fees                FeesForm          `json:"fees"`
	Status              string            `json:"status"`
	Wallet              string            `json:"wallet"`
	AuthorizationID     string            `json:"authorization_id"`
	TransactionID       string            `json:"transaction_id"`
	CardID              string            `json:"card_id"`
	Channel             string            `json:"channel"`
	ClearedAt           string            `json:"cleared_at"`
	CreatedAt           time.Time         `json:"created_at"`
	BillAmount          decimal.Decimal   `json:"bill_amount"`
	ClearedDate         time.Time         `json:"cleared_date"`
	BillCurrency        string            `json:"bill_currency"`
	MerchantData        MerchantDataForm  `json:"merchant_data"`
	DeclineReason       DeclineReasonForm `json:"decline_reason"`
	ExchangeRate        decimal.Decimal   `json:"exchange_rate"`
	ConversionRate      decimal.Decimal   `json:"conversion_rate"`
	LifecycleEventID    string            `json:"lifecycle_event_id"`
	TransactionAmount   decimal.Decimal   `json:"transaction_amount"`
	TransactionCurrency string            `json:"transaction_currency"`
}

type TransactionEventForm struct {
	ID                  string            `json:"id"`
	Fees                FeesForm          `json:"fees"`
	Status              string            `json:"status"`
	Wallet              string            `json:"wallet"`
	CardID              string            `json:"card_id"`
	Channel             string            `json:"channel"`
	ClearedAt           string            `json:"cleared_at"`
	CreatedAt           time.Time         `json:"created_at"`
	BillAmount          decimal.Decimal   `json:"bill_amount"`
	ClearedDate         time.Time         `json:"cleared_date"`
	BillCurrency        string            `json:"bill_currency"`
	MerchantData        MerchantDataForm  `json:"merchant_data"`
	DeclineReason       DeclineReasonForm `json:"decline_reason"`
	ConversionRate      decimal.Decimal   `json:"conversion_rate"`
	ExchangeRate        decimal.Decimal   `json:"exchange_rate"`
	LifecycleEventID    string            `json:"lifecycle_event_id"`
	TransactionAmount   decimal.Decimal   `json:"transaction_amount"`
	TransactionCurrency string            `json:"transaction_currency"`
}
type NotificationEventForm struct {
	CardID               string         `json:"cardId"`
	Status               string         `json:"status"`
	MerchantName         string         `json:"merchantName"`
	InitiateActionID     string         `json:"initiateActionId"`
	TransactionAmount    string         `json:"transactionAmount"`
	MerchantCountryCode  string         `json:"merchantCountryCode"`
	TransactionCurrency  string         `json:"transactionCurrency"`
	TransactionTimestamp time.Time      `json:"transactionTimestamp"`
	Otp                  string         `json:"otp"`
	Wallet               string         `json:"wallet"`
	OtpCardID            string         `json:"card_id"`
	OtpPhoneNumber       OtpPhoneNumber `json:"otp_phone_number"`
}

type CardEventForm struct {
	IssueID string                 `json:"id"`
	Status  common.CardBlockReason `json:"status"`
}

type ReapShippingOrderConfirmationEventForm struct {
	Shipment []ReapShipment `json:"shipment"`
}

type ReapShipment struct {
	SKU    string `json:"sku"`
	CardID string `json:"cardId"`
}

type ReapBillReportEventForm struct {
	ExpireAt    uint64 `json:"expireAt"`
	Link        string `json:"link"`
	ReferenceID string `json:"referenceId"`
}

type EventForm[T EventDataType] struct {
	EventName common.ReapEventName `json:"eventName"`
	EventType common.ReapEventType `json:"eventType"`
	Data      T                    `json:"data"`
}

type OtpPhoneNumber struct {
	DialCode    int    `json:"dial_code"`    // 國碼，例如 86
	PhoneNumber string `json:"phone_number"` // 電話號碼，例如 18616660505
}

type WhaleBasicForm struct {
	BusinessType common.WhaleBusinessType `json:"businessType"`
}

type WhaleCreateCardForm struct {
	WhaleBasicForm
	CardID     uint64    `json:"cardId,string"`
	ID         uint64    `json:"id,string"`
	Status     string    `json:"status"`
	Currency   string    `json:"currency"`
	UserName   string    `json:"userName"`
	CreateTime time.Time `json:"createTime"`
	UID        uint64    `json:"uid,string"`
}

type WhaleCardStatusForm struct {
	WhaleBasicForm
	CardID     uint64                 `json:"cardId,string"`
	ID         uint64                 `json:"id,string"`
	Status     common.WhaleCardStatus `json:"status"`
	CreateTime time.Time              `json:"createTime"`
	UID        uint64                 `json:"uid,string"`
}

type WhaleCardTransactionForm struct {
	WhaleBasicForm
	CardID          uint64                            `json:"cardId,string"`
	ID              string                            `json:"id"`
	Currency        string                            `json:"currency"`
	Amount          decimal.Decimal                   `json:"amount"`
	Fee             decimal.Decimal                   `json:"fee"`
	Type            common.WhaleCardTransactionType   `json:"type"`   // Consumption Credit Reversal Frozen UnFrozen
	Status          common.WhaleCardTransactionStatus `json:"status"` // Pending Closed Fail
	Detail          string                            `json:"detail"`
	FailReason      string                            `json:"failure_reason"`
	TransactionTime time.Time                         `json:"transactionTime"`
	UID             uint64                            `json:"uid,string"`
}

type WhaleBalanceDepositTransactionForm struct {
	WhaleBasicForm
	BalanceDepositID       uint64          `json:"balanceDepositId,string"`
	TXHash                 string          `json:"txHash"`
	Symbol                 string          `json:"symbol"`  // USDT USDC
	Network                string          `json:"network"` // POLYGON TRON
	Amount                 decimal.Decimal `json:"amount"`
	DepositTransactionTime time.Time       `json:"depositTransactionTime"`
	UID                    uint64          `json:"uid,string"`
}

type WhaleBalanceEditForm struct {
	WhaleBasicForm
	UID              uint64          `json:"uid,string"`
	Amount           decimal.Decimal `json:"amount"`
	BalanceDepositID uint64          `json:"balanceDepositId,string"`
	CreateTime       time.Time       `json:"createTime"`
	CardID           string          `json:"cardId"`
	Remark           string          `json:"remark"`
	ID               uint64          `json:"id,string"`
	Status           int             `json:"status,string"` // 1成功
}

type WhaleCard3dsOtpForm struct {
	WhaleBasicForm
	UID        uint64    `json:"uid,string"`
	CardID     uint64    `json:"cardId,string"`
	OTP        string    `json:"otp"`
	ID         uint64    `json:"id,string"`
	CreateTime time.Time `json:"createTime"`
	Status     string    `json:"status"` // active
}
