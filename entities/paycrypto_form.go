package entities

import (
	"shared-modules/common"
)

type PaycryptoWebhookEventParams interface {
	PaycryptoKYCStatusEventParams | PaycryptoCardApplicationReadyEventParams | PaycryptoCardLockStatusEventParams |
		PaycryptoTxStatusEventParams | PaycryptoCardBalanceUpdatedEventParams | PaycryptoCard3dsOtpEventParams
}
type PaycryptoWebhookForm struct {
	Action common.PaycryptoWebhookAction `json:"action"`
	Events []string                      `json:"events"`
}

// Base event structure that all events share
type PaycryptoWebhookEventBase struct {
	ID         string `json:"id"`
	CreateTime int64  `json:"create_time"`
}

type PaycryptoWebhookEvent[T PaycryptoWebhookEventParams] struct {
	PaycryptoWebhookEventBase
	Params T `json:"params"`
}

type PaycryptoKYCStatusEventParams struct {
	AcctNo             string                           `json:"acct_no"`
	CardTypeID         string                           `json:"card_type_id"`
	Status             common.PaycryptoWebhookKYCStatus `json:"status"`
	FaceRecognitionURL string                           `json:"face_recognition_url,omitempty"`
}

type PaycryptoCardApplicationReadyEventParams struct {
	AcctNo     string `json:"acct_no"`
	CardTypeID string `json:"card_type_id"`
}

type PaycryptoCardLockStatusEventParams struct {
	CardNO      string                                 `json:"card_no"`
	RequestType common.PaycryptoWebhookCardRequestType `json:"requestType"`
	Status      common.PaycryptoWebhookCardStatus      `json:"status"`
}

type PaycryptoTxStatusEventParams struct {
	TXID   string                          `json:"tx_id"`
	Status common.PaycryptoWebhookTXStatus `json:"status"`
}

type PaycryptoCardBalanceUpdatedEventParams struct {
	TXID   string                       `json:"tx_id"`
	CardNO string                       `json:"card_no"`
	Type   common.PaycryptoTxRecordType `json:"type,omitempty"` // Optional field in the example
}
type PaycryptoCard3dsOtpEventParams struct {
	OTP                 string `json:"otp"`
	CardNO              string `json:"card_no"`
	TransactionCurrency string `json:"transaction_currency"`
	TransactionAmount   string `json:"transaction_amount"`
	MerchantName        string `json:"merchant_name"`
}
