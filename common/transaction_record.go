package common

import "strings"

type TransactionSide int // 交易方
const (
	TRANSACTION_SIDE_FROM TransactionSide = 1
	TRANSACTION_SIDE_TO   TransactionSide = 2
)

type TransactionStatus int // 交易紀錄狀態
const (
	TRANSACTION_STATUS_TRANSFER_PENDING    TransactionStatus = 201
	TRANSACTION_STATUS_TRANSFER_PROCESSING TransactionStatus = 202
	TRANSACTION_STATUS_TRANSFER_SUCCESS    TransactionStatus = 203
	TRANSACTION_STATUS_TRANSFER_FAILED     TransactionStatus = 204

	TRANSACTION_STATUS_EXCHANGE_PENDING    TransactionStatus = 301
	TRANSACTION_STATUS_EXCHANGE_PROCESSING TransactionStatus = 302
	TRANSACTION_STATUS_EXCHANGE_SUCCESS    TransactionStatus = 303
	TRANSACTION_STATUS_EXCHANGE_FAILED     TransactionStatus = 304
)

type TransactionRecordType int // 交易紀錄類型
const (
	TRANSACTION_RECORD_TYPE_TRANSFER TransactionRecordType = 2
	TRANSACTION_RECORD_TYPE_EXCHANGE TransactionRecordType = 3
)

func (trt TransactionRecordType) String() string {
	switch trt {
	case TRANSACTION_RECORD_TYPE_TRANSFER:
		return "transfer"
	case TRANSACTION_RECORD_TYPE_EXCHANGE:
		return "exchange"
	}
	return ""
}

func (TransactionRecordType) FromString(s string) []TransactionRecordType {

	var result []TransactionRecordType
	switch s {
	case "transfer":
		result = append(result, TRANSACTION_RECORD_TYPE_TRANSFER)
	case "exchange":
		result = append(result, TRANSACTION_RECORD_TYPE_EXCHANGE)
	}

	return result
}

func (ts TransactionSide) String() string {
	switch ts {
	case TRANSACTION_SIDE_FROM:
		return "from"
	case TRANSACTION_SIDE_TO:
		return "to"
	}
	return ""
}

func (ts TransactionSide) FromString(s string) TransactionSide {
	switch strings.ToLower(s) {
	case "from":
		return TRANSACTION_SIDE_FROM
	case "to":
		return TRANSACTION_SIDE_TO
	}
	return 0
}

func (ts TransactionStatus) String() string {
	switch ts {
	case TRANSACTION_STATUS_TRANSFER_PENDING:
		return "transfer_pending"
	case TRANSACTION_STATUS_TRANSFER_PROCESSING:
		return "transfer_processing"
	case TRANSACTION_STATUS_TRANSFER_SUCCESS:
		return "transfer_success"
	case TRANSACTION_STATUS_TRANSFER_FAILED:
		return "transfer_failed"

	case TRANSACTION_STATUS_EXCHANGE_PENDING:
		return "exchange_pending"
	case TRANSACTION_STATUS_EXCHANGE_PROCESSING:
		return "exchange_processing"
	case TRANSACTION_STATUS_EXCHANGE_SUCCESS:
		return "exchange_success"
	case TRANSACTION_STATUS_EXCHANGE_FAILED:
		return "exchange_failed"
	}
	return ""
}
