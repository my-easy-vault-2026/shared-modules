package entities

import (
	"shared-modules/common"
	"time"
)

type EtherfiTransactionFrom struct {
	MerchantIds              []uint64                                `json:"merchantId"`
	CardIds                  []string                                `json:"cardIds"`
	TransactionHistoryEvents []common.EtherfiTransactionHistoryEvent `json:"transactionHistoryEvents"`
	From                     *time.Time                              `json:"from"`
	To                       *time.Time                              `json:"to" `
	Duration                 int64                                   `json:"duration"`
	Force                    bool                                    `json:"force"`
	Page                     int                                     `json:"page"`
	Size                     int                                     `json:"size"`
}

type EtherfiTransactionUpdateForm struct {
	TxId string `json:"txId"`
}

type EtherfiSystemAccountCookieRenewForm struct {
	Id uint64 `json:"id,string"`
}

type EtherfiTransactionTForm struct {
	CardIds  []uint64   `json:"cardIds"`
	TxIds    []string   `json:"txIds"`
	From     *time.Time `json:"from"`
	To       *time.Time `json:"to"`
	Duration int64      `json:"duration"`
}

type EtherfiCreatedCardRetryForm struct {
	InitialSec  int `json:"initialSec"`  // 第一次重式的延遲秒數
	MaxRetrySec int `json:"maxRetrySec"` // 最大重試延遲秒數
}

type EtherfiCardSpendingLImitForm struct {
	CardIds  []uint64   `json:"cardIds"`
	From     *time.Time `json:"from"`
	To       *time.Time `json:"to"`
	Duration int64      `json:"duration"`
	Force    bool       `json:"force"`
}

type EtherfiGetAccountAmountForm struct {
	Id uint64 `json:"id,string"`
}

type EtherfiManualCreateCardForm struct {
	EtherfiMail string `json:"etherfiMail"`
}
