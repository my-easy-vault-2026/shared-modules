package common

type TransactionType int // 交易動作
const (
	TRANSACTION_TYPE_ASSET_ADD           TransactionType = 1
	TRANSACTION_TYPE_ASSET_DEDUCT        TransactionType = 2
	TRANSACTION_TYPE_ASSET_FREEZE        TransactionType = 3
	TRANSACTION_TYPE_ASSET_UNFREEZE      TransactionType = 4
	TRANSACTION_TYPE_FROZEN_ASSET_ADD    TransactionType = 5
	TRANSACTION_TYPE_FROZEN_ASSET_DEDUCT TransactionType = 6
)

func (tt TransactionType) String() string {
	switch tt {
	case TRANSACTION_TYPE_ASSET_ADD:
		return "asset_add"
	case TRANSACTION_TYPE_ASSET_DEDUCT:
		return "asset_deduct"
	case TRANSACTION_TYPE_ASSET_FREEZE:
		return "asset_freeze"
	case TRANSACTION_TYPE_ASSET_UNFREEZE:
		return "asset_unfreeze"
	case TRANSACTION_TYPE_FROZEN_ASSET_ADD:
		return "frozen_asset_add"
	case TRANSACTION_TYPE_FROZEN_ASSET_DEDUCT:
		return "frozon_asset_deduct"
	}
	return ""
}
