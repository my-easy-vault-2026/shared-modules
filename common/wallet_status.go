package common

type WalletStatus int // 錢包狀態
const (
	WALLET_STATUS_ENABLED  WalletStatus = 1
	WALLET_STATUS_DISABLED WalletStatus = 2
)

func (ws WalletStatus) String() string {
	switch ws {
	case WALLET_STATUS_ENABLED:
		return "enabled"
	case WALLET_STATUS_DISABLED:
		return "disabled"
	}
	return ""
}
