package common

type TransferStatus int // 轉帳狀態
const (
	TRANSFER_STATUS_PENDING TransferStatus = 1
	TRANSFER_STATUS_SUCCESS TransferStatus = 2
	TRANSFER_STATUS_FAILED  TransferStatus = 3
)
