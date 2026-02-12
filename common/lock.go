package common

type LockPurpose int // 鎖用途
const (
	LOCK_PURPOSE_CREATE_WALLET    LockPurpose = 1
	LOCK_PURPOSE_EXCHANGE_CONFIRM LockPurpose = 2
	LOCK_PURPOSE_TRANSFER_CONFIRM LockPurpose = 3
)

func (lp LockPurpose) String() string {
	switch lp {
	case LOCK_PURPOSE_EXCHANGE_CONFIRM:
		return "exchange_confirm"
	case LOCK_PURPOSE_TRANSFER_CONFIRM:
		return "transfer_confirm"
	case LOCK_PURPOSE_CREATE_WALLET:
		return "create_wallet"
	}
	return ""
}
