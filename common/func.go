package common

import (
	"math"
	"strconv"
	"strings"
)

func IsSystemAccount(userID uint64) bool {
	return userID < 10
}

func (al AdminLevel) String() string {
	switch al {
	case ADMIN_LEVEL_ANALYTICS:
		return "analytics"
	case ADMIN_LEVEL_CUSTOMER_SERVICE:
		return "customer_service"
	case ADMIN_LEVEL_FINANCE:
		return "finance"
	case ADMIN_LEVEL_ADMIN:
		return "admin"
	}
	return ""
}

func (al AdminLevel) FromString(s string) AdminLevel {
	switch s {
	case "analytics":
		return ADMIN_LEVEL_ANALYTICS
	case "customer_service":
		return ADMIN_LEVEL_CUSTOMER_SERVICE
	case "finance":
		return ADMIN_LEVEL_FINANCE
	case "admin":
		return ADMIN_LEVEL_ADMIN
	}
	return 0
}

func (a AssetType) String() string {
	switch a {
	case ASSET_TYPE_CRYPTO:
		return "crypto"
	case ASSET_TYPE_FIAT:
		return "fiat"
	case ASSET_TYPE_CARD_PRODUCT:
		return "card_product"
	case ASSET_TYPE_POINT:
		return "point"
	case ASSET_TYPE_AUTO_YIELD:
		return "auto_yield"
	}
	return ""
}

func GetAssetType(categoryID uint64) AssetType {
	switch true {
	case categoryID >= 100 && categoryID < 200:
		return ASSET_TYPE_CRYPTO
	case categoryID >= 200 && categoryID < 300:
		return ASSET_TYPE_FIAT
	case categoryID >= 300 && categoryID < 1000:
		return ASSET_TYPE_CARD_PRODUCT
	case categoryID >= 1000 && categoryID < 2000:
		return ASSET_TYPE_POINT
	}
	return ASSET_TYPE_CRYPTO
}

func (bt BalanceType) String() string {
	switch bt {
	case BALANCE_TYPE_CENTRALIZED:
		return "centralized"
	case BALANCE_TYPE_DECENTRALIZED:
		return "decentralized"
	case BALANCE_TYPE_MERCHANT_CENTRALIZED:
		return "merchant_centralized"
	case BALANCE_TYPE_MERCHANT_DECENTRALIZED:
		return "merchant_decentralized"
	}
	return ""
}

func (at ATMToggle) String() string {
	switch at {
	case ATM_TOGGLE_ENABLED:
		return "enabled"
	case ATM_TOGGLE_DISABLED:
		return "disabled"
	}
	return ""
}

func (a3 Auto3DSStatus) String() string {
	switch a3 {
	case AUTO_3DS_STATUS_ENABLED:
		return "enabled"
	case AUTO_3DS_STATUS_DISABLED:
		return "disabled"
	}
	return ""
}

func (at AutoTopUpStatus) String() string {
	switch at {
	case AUTO_TOP_UP_STATUS_ENABLED:
		return "enabled"
	case AUTO_TOP_UP_STATUS_DISABLED:
		return "disabled"
	}
	return ""
}

func (bt BalanceType) FromString(s string) BalanceType {
	switch strings.ToLower(s) {
	case "centralized":
		return BALANCE_TYPE_CENTRALIZED
	case "decentralized":
		return BALANCE_TYPE_DECENTRALIZED
	case "merchant_centralized":
		return BALANCE_TYPE_MERCHANT_CENTRALIZED
	case "merchant_decentralized":
		return BALANCE_TYPE_MERCHANT_DECENTRALIZED
	}
	return 0
}

func (bt BillType) String() string {
	switch bt {
	case BILL_TYPE_APPLY_FEE_DEDUCT:
		return "apply_fee_deduct"
	case BILL_TYPE_APPLY_FEE_ADD:
		return "apply_fee_add"
	case BILL_TYPE_APPLY_CARD_FEE_DEDUCT:
		return "apply_card_fee_deduct"
	case BILL_TYPE_APPLY_CARD_FEE_ADD:
		return "apply_card_fee_add"
	case BILL_TYPE_APPLY_AUTO_REACTIVATE_FEE_DEDUCT:
		return "apply_auto_reactivate_fee_deduct"
	case BILL_TYPE_APPLY_AUTO_REACTIVATE_FEE_ADD:
		return "apply_auto_reactivate_fee_add"
	case BILL_TYPE_APPLY_AUTO_REACTIVATE_CARD_FEE_DEDUCT:
		return "apply_auto_reactivate_card_fee_deduct"
	case BILL_TYPE_APPLY_AUTO_REACTIVATE_CARD_FEE_ADD:
		return "apply_auto_reactivate_card_fee_add"
	case BILL_TYPE_APPLY_MERCHANT_FEE_DEDUCT:
		return "apply_merchant_fee_deduct"
	case BILL_TYPE_APPLY_MERCHANT_FEE_ADD:
		return "apply_merchant_fee_add"
	case BILL_TYPE_APPLY_MERCHANT_CARD_FEE_DEDUCT:
		return "apply_merchant_card_fee_deduct"
	case BILL_TYPE_APPLY_MERCHANT_CARD_FEE_ADD:
		return "apply_merchant_card_fee_add"
	case BILL_TYPE_APPLY_WHALE_FEE_DEDUCT:
		return "apply_whale_fee_deduct"
	case BILL_TYPE_APPLY_WHALE_FEE_ADD:
		return "apply_whale_fee_add"
	case BILL_TYPE_APPLY_WHALE_CARD_FEE_DEDUCT:
		return "apply_whale_card_fee_deduct"
	case BILL_TYPE_APPLY_WHALE_CARD_FEE_ADD:
		return "apply_whale_card_fee_add"
	case BILL_TYPE_APPLY_WHALE_RECEIVE_ADD:
		return "apply_whale_receive_add"
	case BILL_TYPE_APPLY_WHALE_RECEIVE_DEDUCT:
		return "apply_whale_receive_deduct"
	case BILL_TYPE_APPLY_PAYCRYPTO_FEE_DEDUCT:
		return "apply_paycrypto_fee_deduct"
	case BILL_TYPE_APPLY_PAYCRYPTO_FEE_ADD:
		return "apply_paycrypto_fee_add"
	case BILL_TYPE_APPLY_PAYCRYPTO_CARD_FEE_DEDUCT:
		return "apply_paycrypto_card_fee_deduct"
	case BILL_TYPE_APPLY_PAYCRYPTO_CARD_FEE_ADD:
		return "apply_paycrypto_card_fee_add"
	case BILL_TYPE_APPLY_PAYCRYPTO_RECEIVE_ADD:
		return "apply_paycrypto_receive_add"
	case BILL_TYPE_APPLY_PAYCRYPTO_RECEIVE_DEDUCT:
		return "apply_paycrypto_receive_deduct"
	case BILL_TYPE_TRANSFER_SEND_DEDUCT:
		return "transfer_send_deduct"
	case BILL_TYPE_TRANSFER_SEND_FEE_DEDUCT:
		return "transfer_send_fee_deduct"
	case BILL_TYPE_TRANSFER_SEND_ADD:
		return "transfer_send_add"
	case BILL_TYPE_TRANSFER_SEND_FEE_ADD:
		return "transfer_send_fee_add"
	case BILL_TYPE_TRANSFER_RECEIVE_ADD:
		return "transfer_receive_add"
	case BILL_TYPE_TRANSFER_RECEIVE_DEDUCT:
		return "transfer_receive_deduct"
	case BILL_TYPE_TRANSFER_SEND_EXCHANGE_FEE_DEDUCT:
		return "transfer_send_exchange_fee_deduct"
	case BILL_TYPE_TRANSFER_SEND_EXCHANGE_FEE_ADD:
		return "transfer_send_exchange_fee_add"
	case BILL_TYPE_TRANSFER_MERCHANT_SEND_DEDUCT:
		return "transfer_merchant_send_deduct"
	case BILL_TYPE_TRANSFER_MERCHANT_SEND_FEE_DEDUCT:
		return "transfer_merchant_send_fee_deduct"
	case BILL_TYPE_TRANSFER_MERCHANT_SEND_ADD:
		return "transfer_merchant_send_add"
	case BILL_TYPE_TRANSFER_MERCHANT_SEND_FEE_ADD:
		return "transfer_merchant_send_fee_add"
	case BILL_TYPE_TRANSFER_MERCHANT_RECEIVE_ADD:
		return "transfer_merchant_receive_add"
	case BILL_TYPE_TRANSFER_MERCHANT_RECEIVE_DEDUCT:
		return "transfer_merchant_receive_deduct"
	case BILL_TYPE_TRANSFER_MERCHANT_SEND_EXCHANGE_FEE_DEDUCT:
		return "transfer_merchant_send_exchange_fee_deduct"
	case BILL_TYPE_TRANSFER_MERCHANT_SEND_EXCHANGE_FEE_ADD:
		return "transfer_merchant_send_exchange_fee_add"
	case BILL_TYPE_EXCHANGE_SEND_DEDUCT:
		return "exchange_send_deduct"
	case BILL_TYPE_EXCHANGE_SEND_FEE_DEDUCT:
		return "exchange_send_fee_deduct"
	case BILL_TYPE_EXCHANGE_SEND_ADD:
		return "exchange_send_add"
	case BILL_TYPE_EXCHANGE_SEND_FEE_ADD:
		return "exchange_send_fee_add"
	case BILL_TYPE_EXCHANGE_RECEIVE_ADD:
		return "exchange_receive_add"
	case BILL_TYPE_EXCHANGE_RECEIVE_DEDUCT:
		return "exchange_receive_deduct"
	case BILL_TYPE_EXCHANGE_MERCHANT_SEND_DEDUCT:
		return "exchange_merchant_send_deduct"
	case BILL_TYPE_EXCHANGE_MERCHANT_SEND_FEE_DEDUCT:
		return "exchange_merchant_send_fee_deduct"
	case BILL_TYPE_EXCHANGE_MERCHANT_SEND_ADD:
		return "exchange_merchant_send_add"
	case BILL_TYPE_EXCHANGE_MERCHANT_SEND_FEE_ADD:
		return "exchange_merchant_send_fee_add"
	case BILL_TYPE_EXCHANGE_MERCHANT_RECEIVE_ADD:
		return "exchange_merchant_receive_add"
	case BILL_TYPE_EXCHANGE_MERCHANT_RECEIVE_DEDUCT:
		return "exchange_merchant_receive_deduct"
	case BILL_TYPE_WITHDRAW_SEND_DEDUCT:
		return "withdraw_send_deduct"
	case BILL_TYPE_WITHDRAW_SEND_FEE_DEDUCT:
		return "withdraw_send_fee_deduct"
	case BILL_TYPE_WITHDRAW_SEND_ADD:
		return "withdraw_send_add"
	case BILL_TYPE_WITHDRAW_SEND_FEE_ADD:
		return "withdraw_send_fee_add"
	case BILL_TYPE_WITHDRAW_REJECT_DEDUCT:
		return "withdraw_reject_deduct"
	case BILL_TYPE_WITHDRAW_REJECT_FEE_DEDUCT:
		return "withdraw_reject_fee_deduct"
	case BILL_TYPE_WITHDRAW_REJECT_ADD:
		return "withdraw_reject_add"
	case BILL_TYPE_WITHDRAW_REJECT_FEE_ADD:
		return "withdraw_reject_fee_add"
	case BILL_TYPE_WITHDRAW_CANCEL_DEDUCT:
		return "withdraw_cancel_deduct"
	case BILL_TYPE_WITHDRAW_CANCEL_ADD:
		return "withdraw_cancel_add"
	case BILL_TYPE_WITHDRAW_CANCEL_FEE_DEDUCT:
		return "withdraw_cancel_fee_deduct"
	case BILL_TYPE_WITHDRAW_CANCEL_FEE_ADD:
		return "withdraw_cancel_fee_add"
	case BILL_TYPE_WITHDRAW_SUCCESS_DEDUCT:
		return "withdraw_success_deduct"
	case BILL_TYPE_WITHDRAW_SUCCESS_ADD:
		return "withdraw_success_add"
	case BILL_TYPE_WITHDRAW_TRANSFER_SEND_DEDUCT:
		return "withdraw_transfer_send_deduct"
	case BILL_TYPE_WITHDRAW_TRANSFER_SEND_FEE_DEDUCT:
		return "withdraw_transfer_send_fee_deduct"
	case BILL_TYPE_WITHDRAW_TRANSFER_SEND_ADD:
		return "withdraw_transfer_send_add"
	case BILL_TYPE_WITHDRAW_TRANSFER_SEND_FEE_ADD:
		return "withdraw_transfer_send_fee_add"
	case BILL_TYPE_WITHDRAW_TRANSFER_FAILED_SEND_DEDUCT:
		return "withdraw_transfer_failed_send_deduct"
	case BILL_TYPE_WITHDRAW_TRANSFER_FAILED_SEND_FEE_DEDUCT:
		return "withdraw_transfer_failed_send_fee_deduct"
	case BILL_TYPE_WITHDRAW_TRANSFER_FAILED_SEND_ADD:
		return "withdraw_transfer_failed_send_add"
	case BILL_TYPE_WITHDRAW_TRANSFER_FAILED_SEND_FEE_ADD:
		return "withdraw_transfer_failed_send_fee_add"
	case BILL_TYPE_PAY_REAP_AUTHORIZATION_FREEZE:
		return "pay_reap_authorization_freeze"
	case BILL_TYPE_PAY_REAP_TRANSACTION_DEDUCT_FREEZED:
		return "pay_reap_transaction_deduct_freezed"
	case BILL_TYPE_PAY_REAP_TRANSACTION_ADD:
		return "pay_reap_transaction_add"
	case BILL_TYPE_PAY_REAP_TRANSACTION_REVERSAL_UNFREEZE:
		return "pay_reap_transaction_reversal_unfreeze"
	case BILL_TYPE_PAY_REAP_DECLINE_UNFREEZE:
		return "pay_reap_decline_unfreeze"
	case BILL_TYPE_PAY_REAP_TRANSACTION_DIRECT_DEDUCT:
		return "pay_reap_transaction_direct_deduct"
	case BILL_TYPE_PAY_REAP_TRANSACTION_DIRECT_ADD:
		return "pay_reap_transaction_direct_add"
	case BILL_TYPE_PAY_REAP_TRANSACTION_OFFLINE_DEDUCT:
		return "pay_reap_transaction_offline_deduct"
	case BILL_TYPE_PAY_REAP_TRANSACTION_OFFLINE_ADD:
		return "pay_reap_transaction_offline_add"
	case BILL_TYPE_PAY_REAP_TRANSACTION_REFUND_DEDUCT:
		return "pay_reap_transaction_refund_deduct"
	case BILL_TYPE_PAY_REAP_TRANSACTION_REFUND_ADD:
		return "pay_reap_transaction_refund_add"
	case BILL_TYPE_PAY_REAP_TRANSACTION_DRIECT_AFTER_DECLINE_ADD:
		return "pay_reap_transaction_driect_after_decline_add"
	case BILL_TYPE_PAY_REAP_TRANSACTION_DRIECT_AFTER_DECLINE_DEDUCT:
		return "pay_reap_transaction_driect_after_decline_deduct"
	case BILL_TYPE_PAY_REAP_MERCHANT_AUTHORIZATION_FREEZE:
		return "pay_reap_merchant_authorization_freeze"
	case BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_DEDUCT_FREEZED:
		return "pay_reap_merchant_transaction_deduct_freezed"
	case BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_ADD:
		return "pay_reap_merchant_transaction_add"
	case BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_REVERSAL_UNFREEZE:
		return "pay_reap_merchant_transaction_reversal_unfreeze"
	case BILL_TYPE_PAY_REAP_MERCHANT_DECLINE_UNFREEZE:
		return "pay_reap_merchant_decline_unfreeze"
	case BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_DIRECT_DEDUCT:
		return "pay_reap_merchant_transaction_direct_deduct"
	case BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_DIRECT_ADD:
		return "pay_reap_merchant_transaction_direct_add"
	case BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_OFFLINE_DEDUCT:
		return "pay_reap_merchant_transaction_offline_deduct"
	case BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_OFFLINE_ADD:
		return "pay_reap_merchant_transaction_offline_add"
	case BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_REFUND_DEDUCT:
		return "pay_reap_merchant_transaction_refund_deduct"
	case BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_REFUND_ADD:
		return "pay_reap_merchant_transaction_refund_add"
	case BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_DRIECT_AFTER_DECLINE_ADD:
		return "pay_reap_merchant_transaction_driect_after_decline_add"
	case BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_DRIECT_AFTER_DECLINE_DEDUCT:
		return "pay_reap_merchant_transaction_driect_after_decline_deduct"
	case BILL_TYPE_PAY_REAP_TRANSACTION_REVERSAL_PARTIAL_UNFREEZE:
		return "pay_reap_transaction_reversal_partial_unfreeze"
	case BILL_TYPE_PAY_REAP_MERCHANT_TRANSACTION_REVERSAL_PARTIAL_UNFREEZE:
		return "pay_reap_merchant_transaction_reversal_partial_unfreeze"
	case BILL_TYPE_DEPOSIT_RECEIVE_ADD:
		return "deposit_receive_add"
	case BILL_TYPE_DEPOSIT_RECEIVE_DEDUCT:
		return "deposit_receive_deduct"
	case BILL_TYPE_DEPOSIT_FEE_DEDUCT:
		return "deposit_fee_deduct"
	case BILL_TYPE_DEPOSIT_FEE_ADD:
		return "deposit_fee_add"
	case BILL_TYPE_TOP_UP_SEND_DEDUCT:
		return "top_up_send_deduct"
	case BILL_TYPE_TOP_UP_SEND_FEE_DEDUCT:
		return "top_up_send_fee_deduct"
	case BILL_TYPE_TOP_UP_SEND_ADD:
		return "top_up_send_add"
	case BILL_TYPE_TOP_UP_SEND_FEE_ADD:
		return "top_up_send_fee_add"
	case BILL_TYPE_TOP_UP_RECEIVE_ADD:
		return "top_up_receive_add"
	case BILL_TYPE_TOP_UP_RECEIVE_DEDUCT:
		return "top_up_receive_deduct"
	case BILL_TYPE_TOP_UP_SEND_EXCHANGE_FEE_DEDUCT:
		return "top_up_send_exchange_fee_deduct"
	case BILL_TYPE_TOP_UP_SEND_EXCHANGE_FEE_ADD:
		return "top_up_send_exchange_fee_add"
	case BILL_TYPE_TOP_UP_AUTO_SEND_DEDUCT:
		return "top_up_auto_send_deduct"
	case BILL_TYPE_TOP_UP_AUTO_SEND_FEE_DEDUCT:
		return "top_up_auto_send_fee_deduct"
	case BILL_TYPE_TOP_UP_AUTO_SEND_ADD:
		return "top_up_auto_send_add"
	case BILL_TYPE_TOP_UP_AUTO_SEND_FEE_ADD:
		return "top_up_auto_send_fee_add"
	case BILL_TYPE_TOP_UP_AUTO_RECEIVE_ADD:
		return "top_up_auto_receive_add"
	case BILL_TYPE_TOP_UP_AUTO_RECEIVE_DEDUCT:
		return "top_up_auto_receive_deduct"
	case BILL_TYPE_TOP_UP_AUTO_SEND_EXCHANGE_FEE_DEDUCT:
		return "top_up_auto_send_exchange_fee_deduct"
	case BILL_TYPE_TOP_UP_AUTO_SEND_EXCHANGE_FEE_ADD:
		return "top_up_auto_send_exchange_fee_add"
	case BILL_TYPE_TOP_UP_WHALE_SEND_FREEZE:
		return "top_up_whale_send_freeze"
	case BILL_TYPE_TOP_UP_WHALE_SEND_FEE_FREEZE:
		return "top_up_whale_send_fee_freeze"
	case BILL_TYPE_TOP_UP_WHALE_SEND_FROZEN_DEDUCT:
		return "top_up_whale_send_frozen_deduct"
	case BILL_TYPE_TOP_UP_WHALE_SEND_FEE_FROZEN_DEDUCT:
		return "top_up_whale_send_fee_frozen_deduct"
	case BILL_TYPE_TOP_UP_WHALE_SEND_ADD:
		return "top_up_whale_send_add"
	case BILL_TYPE_TOP_UP_WHALE_SEND_FEE_ADD:
		return "top_up_whale_send_fee_add"
	case BILL_TYPE_TOP_UP_WHALE_RECEIVE_ADD:
		return "top_up_whale_receive_add"
	case BILL_TYPE_TOP_UP_WHALE_RECEIVE_DEDUCT:
		return "top_up_whale_receive_deduct"
	case BILL_TYPE_TOP_UP_WHALE_SEND_EXCHANGE_FEE_FREEZE:
		return "top_up_whale_send_exchange_fee_freeze"
	case BILL_TYPE_TOP_UP_WHALE_SEND_EXCHANGE_FEE_FROZEN_DEDUCT:
		return "top_up_whale_send_exchange_fee_frozen_deduct"
	case BILL_TYPE_TOP_UP_WHALE_SEND_EXCHANGE_FEE_ADD:
		return "top_up_whale_send_exchange_fee_add"
	case BILL_TYPE_TOP_UP_WHALE_SEND_UNFREEZE:
		return "top_up_whale_send_unfreeze"
	case BILL_TYPE_TOP_UP_WHALE_SEND_FEE_UNFREEZE:
		return "top_up_whale_send_fee_unfreeze"
	case BILL_TYPE_TOP_UP_WHALE_SEND_EXCHANGE_FEE_UNFREEZE:
		return "top_up_whale_send_exchange_fee_unfreeze"
	case BILL_TYPE_TOP_UP_PAYCRYPTO_SEND_DEDUCT:
		return "top_up_paycrypto_send_deduct"
	case BILL_TYPE_TOP_UP_PAYCRYPTO_SEND_FEE_DEDUCT:
		return "top_up_paycrypto_send_fee_deduct"
	case BILL_TYPE_TOP_UP_PAYCRYPTO_SEND_ADD:
		return "top_up_paycrypto_send_add"
	case BILL_TYPE_TOP_UP_PAYCRYPTO_SEND_FEE_ADD:
		return "top_up_paycrypto_send_fee_add"
	case BILL_TYPE_TOP_UP_PAYCRYPTO_RECEIVE_ADD:
		return "top_up_paycrypto_receive_add"
	case BILL_TYPE_TOP_UP_PAYCRYPTO_RECEIVE_DEDUCT:
		return "top_up_paycrypto_receive_deduct"
	case BILL_TYPE_TOP_UP_PAYCRYPTO_SEND_EXCHANGE_FEE_DEDUCT:
		return "top_up_paycrypto_send_exchange_fee_deduct"
	case BILL_TYPE_TOP_UP_PAYCRYPTO_SEND_EXCHANGE_FEE_ADD:
		return "top_up_paycrypto_send_exchange_fee_add"
	case BILL_TYPE_TOP_DOWN_SEND_DEDUCT:
		return "top_down_send_deduct"
	case BILL_TYPE_TOP_DOWN_SEND_FEE_DEDUCT:
		return "top_down_send_fee_deduct"
	case BILL_TYPE_TOP_DOWN_SEND_ADD:
		return "top_down_send_add"
	case BILL_TYPE_TOP_DOWN_SEND_FEE_ADD:
		return "top_down_send_fee_add"
	case BILL_TYPE_TOP_DOWN_RECEIVE_ADD:
		return "top_down_receive_add"
	case BILL_TYPE_TOP_DOWN_RECEIVE_DEDUCT:
		return "top_down_receive_deduct"
	case BILL_TYPE_TOP_DOWN_SEND_EXCHANGE_FEE_ADD:
		return "top_down_send_exchange_fee_add"
	case BILL_TYPE_TOP_DOWN_SEND_EXCHANGE_FEE_DEDUCT:
		return "top_down_send_exchange_fee_deduct"
	case BILL_TYPE_TOP_DOWN_WHALE_SEND_FREEZE:
		return "top_down_whale_send_freeze"
	case BILL_TYPE_TOP_DOWN_WHALE_SEND_FROZEN_DEDUCT:
		return "top_down_whale_send_frozen_deduct"
	case BILL_TYPE_TOP_DOWN_WHALE_SEND_FEE_FREEZE:
		return "top_down_whale_send_fee_freeze"
	case BILL_TYPE_TOP_DOWN_WHALE_SEND_FEE_FROZEN_DEDUCT:
		return "top_down_whale_send_fee_frozen_deduct"
	case BILL_TYPE_TOP_DOWN_WHALE_SEND_ADD:
		return "top_down_whale_send_add"
	case BILL_TYPE_TOP_DOWN_WHALE_SEND_FEE_ADD:
		return "top_down_whale_send_fee_add"
	case BILL_TYPE_TOP_DOWN_WHALE_RECEIVE_ADD:
		return "top_down_whale_receive_add"
	case BILL_TYPE_TOP_DOWN_WHALE_RECEIVE_DEDUCT:
		return "top_down_whale_receive_deduct"
	case BILL_TYPE_TOP_DOWN_WHALE_SEND_EXCHANGE_FEE_ADD:
		return "top_down_whale_send_exchange_fee_add"
	case BILL_TYPE_TOP_DOWN_WHALE_SEND_EXCHANGE_FEE_FREEZE:
		return "top_down_whale_send_exchange_fee_freeze"
	case BILL_TYPE_TOP_DOWN_WHALE_SEND_EXCHANGE_FEE_FROZEN_DEDUCT:
		return "top_down_whale_send_exchange_fee_frozen_deduct"
	case BILL_TYPE_TOP_DOWN_WHALE_SEND_UNFREEZE:
		return "top_down_whale_send_unfreeze"
	case BILL_TYPE_TOP_DOWN_WHALE_SEND_FEE_UNFREEZE:
		return "top_down_whale_send_fee_unfreeze"
	case BILL_TYPE_TOP_DOWN_WHALE_SEND_EXCHANGE_FEE_UNFREEZE:
		return "top_down_whale_send_exchange_fee_unfreeze"
	case BILL_TYPE_REFUND_REAP_TRANSACTION_REFUND_DEDUCT:
		return "refund_reap_transaction_refund_deduct"
	case BILL_TYPE_REFUND_REAP_TRANSACTION_REFUND_ADD:
		return "refund_reap_transaction_refund_add"
	case BILL_TYPE_CARD_TO_CARD_SEND_DEDUCT:
		return "card_to_card_send_deduct"
	case BILL_TYPE_CARD_TO_CARD_SEND_FEE_DEDUCT:
		return "card_to_card_send_fee_deduct"
	case BILL_TYPE_CARD_TO_CARD_SEND_ADD:
		return "card_to_card_send_add"
	case BILL_TYPE_CARD_TO_CARD_SEND_FEE_ADD:
		return "card_to_card_send_fee_add"
	case BILL_TYPE_CARD_TO_CARD_RECEIVE_ADD:
		return "card_to_card_receive_add"
	case BILL_TYPE_CARD_TO_CARD_RECEIVE_DEDUCT:
		return "card_to_card_receive_deduct"
	case BILL_TYPE_CARD_TO_CARD_SEND_EXCHANGE_FEE_DEDUCT:
		return "card_to_card_send_exchange_fee_deduct"
	case BILL_TYPE_CARD_TO_CARD_SEND_EXCHANGE_FEE_ADD:
		return "card_to_card_send_exchange_fee_add"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FREEZE:
		return "card_to_card_whale_send_freeze"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FROZEN_DEDUCT:
		return "card_to_card_whale_send_frozen_deduct"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FEE_FREEZE:
		return "card_to_card_whale_send_fee_freeze"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FEE_FROZEN_DEDUCT:
		return "card_to_card_whale_send_fee_frozen_deduct"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_ADD:
		return "card_to_card_whale_send_add"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FEE_ADD:
		return "card_to_card_whale_send_fee_add"
	case BILL_TYPE_CARD_TO_CARD_WHALE_RECEIVE_ADD:
		return "card_to_card_whale_receive_add"
	case BILL_TYPE_CARD_TO_CARD_WHALE_RECEIVE_DEDUCT:
		return "card_to_card_whale_receive_deduct"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_EXCHANGE_FEE_FREEZE:
		return "card_to_card_whale_send_exchange_fee_freeze"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_EXCHANGE_FEE_FROZEN_DEDUCT:
		return "card_to_card_whale_send_exchange_fee_frozen_deduct"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_EXCHANGE_FEE_ADD:
		return "card_to_card_whale_send_exchange_fee_add"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_UNFREEZE:
		return "card_to_card_whale_send_unfreeze"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FEE_UNFREEZE:
		return "card_to_card_whale_send_fee_unfreeze"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_EXCHANGE_FEE_UNFREEZE:
		return "card_to_card_whale_send_exchange_fee_unfreeze"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SELF_SEND_DEDUCT:
		return "card_to_card_whale_self_send_deduct"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SELF_SEND_FEE_DEDUCT:
		return "card_to_card_whale_self_send_fee_deduct"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SELF_SEND_ADD:
		return "card_to_card_whale_self_send_add"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SELF_SEND_FEE_ADD:
		return "card_to_card_whale_self_send_fee_add"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SELF_RECEIVE_ADD:
		return "card_to_card_whale_self_receive_add"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SELF_RECEIVE_DEDUCT:
		return "card_to_card_whale_self_receive_deduct"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SELF_SEND_EXCHANGE_FEE_DEDUCT:
		return "card_to_card_whale_self_send_exchange_fee_deduct"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SELF_SEND_EXCHANGE_FEE_ADD:
		return "card_to_card_whale_self_send_exchange_fee_add"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FAILED_HALFWAY_FROZEN_DEDUCT:
		return "card_to_card_whale_send_failed_halfway_frozen_deduct"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FAILED_HALFWAY_ADD:
		return "card_to_card_whale_send_failed_halfway_add"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FAILED_HALFWAY_FEE_FROZEN_DEDUCT:
		return "card_to_card_whale_send_failed_halfway_fee_frozen_deduct"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FAILED_HALFWAY_FEE_ADD:
		return "card_to_card_whale_send_failed_halfway_fee_add"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FAILED_HALFWAY_EXCHANGE_FEE_FROZEN_DEDUCT:
		return "card_to_card_whale_send_failed_halfway_exchange_fee_frozen_deduct"
	case BILL_TYPE_CARD_TO_CARD_WHALE_SEND_FAILED_HALFWAY_EXCHANGE_FEE_ADD:
		return "card_to_card_whale_send_failed_halfway_exchange_fee_add"
	case BILL_TYPE_CARD_TO_CARD_ADMIN_DELETE_SEND_DEDUCT:
		return "card_to_card_admin_delete_send_deduct"
	case BILL_TYPE_CARD_TO_CARD_ADMIN_DELETE_SEND_ADD:
		return "card_to_card_admin_delete_send_add"
	case BILL_TYPE_MANUAL_ADD_ADD:
		return "manual_add_add"
	case BILL_TYPE_MANUAL_ADD_DEDUCT:
		return "manual_add_deduct"
	case BILL_TYPE_MANUAL_DEDUCT_DEDUCT:
		return "manual_deduct_deduct"
	case BILL_TYPE_MANUAL_DEDUCT_ADD:
		return "manual_deduct_add"
	case BILL_TYPE_MANUAL_FREEZE_FREEZE:
		return "manual_freeze_freeze"
	case BILL_TYPE_MANUAL_UNFREEZE_UNFREEZE:
		return "manual_unfreeze_unfreeze"
	case BILL_TYPE_MANUAL_ADD_FREEZE_ADD_FREEZE:
		return "manual_add_freeze_add_freeze"
	case BILL_TYPE_MANUAL_ADD_FREEZE_DEDUCT:
		return "manual_add_freeze_deduct"
	case BILL_TYPE_MANUAL_DEDUCT_FREEZE_DEDUCT_FREEZE:
		return "manual_deduct_freeze_deduct_freeze"
	case BILL_TYPE_MANUAL_DEDUCT_FREEZE_ADD:
		return "manual_deduct_freeze_add"
	case BILL_TYPE_MERCHANT_AUTO_EXCHANGE_TIMED_SEND_DEDUCT:
		return "merchant_auto_exchange_timed_send_deduct"
	case BILL_TYPE_MERCHANT_AUTO_EXCHANGE_TIMED_SEND_FEE_DEDUCT:
		return "merchant_auto_exchange_timed_send_fee_deduct"
	case BILL_TYPE_MERCHANT_AUTO_EXCHANGE_TIMED_SEND_ADD:
		return "merchant_auto_exchange_timed_send_add"
	case BILL_TYPE_MERCHANT_AUTO_EXCHANGE_TIMED_SEND_FEE_ADD:
		return "merchant_auto_exchange_timed_send_fee_add"
	case BILL_TYPE_MERCHANT_AUTO_EXCHANGE_TIMED_RECEIVE_ADD:
		return "merchant_auto_exchange_timed_receive_add"
	case BILL_TYPE_MERCHANT_AUTO_EXCHANGE_TIMED_RECEIVE_DEDUCT:
		return "merchant_auto_exchange_timed_receive_deduct"

	case BILL_TYPE_CHIPPAY_EXPRESS_RECEIVE_ADD:
		return "chippay_express_receive_add"
	case BILL_TYPE_CHIPPAY_EXPRESS_RECEIVE_DEDUCT:
		return "chippay_express_receive_deduct"
	case BILL_TYPE_CHIPPAY_EXPRESS_FEE_DEDUCT:
		return "chippay_express_fee_deduct"
	case BILL_TYPE_CHIPPAY_EXPRESS_FEE_ADD:
		return "chippay_express_fee_add"
	case BILL_TYPE_CHIPPAY_EXPRESS_PAYCRYPTO_RECEIVE_ADD:
		return "chippay_express_paycrypto_receive_add"
	case BILL_TYPE_CHIPPAY_EXPRESS_PAYCRYPTO_RECEIVE_DEDUCT:
		return "chippay_express_paycrypto_receive_deduct"
	case BILL_TYPE_CHIPPAY_EXPRESS_PAYCRYPTO_FEE_DEDUCT:
		return "chippay_express_paycrypto_fee_deduct"
	case BILL_TYPE_CHIPPAY_EXPRESS_PAYCRYPTO_FEE_ADD:
		return "chippay_express_paycrypto_fee_add"

	case BILL_TYPE_MERCHANT_AUTO_EXCHANGE_ON_DEPOSIT_SEND_DEDUCT:
		return "merchant_auto_exchange_on_deposit_send_deduct"
	case BILL_TYPE_MERCHANT_AUTO_EXCHANGE_ON_DEPOSIT_SEND_FEE_DEDUCT:
		return "merchant_auto_exchange_on_deposit_send_fee_deduct"
	case BILL_TYPE_MERCHANT_AUTO_EXCHANGE_ON_DEPOSIT_SEND_ADD:
		return "merchant_auto_exchange_on_deposit_send_add"
	case BILL_TYPE_MERCHANT_AUTO_EXCHANGE_ON_DEPOSIT_SEND_FEE_ADD:
		return "merchant_auto_exchange_on_deposit_send_fee_add"
	case BILL_TYPE_MERCHANT_AUTO_EXCHANGE_ON_DEPOSIT_RECEIVE_ADD:
		return "merchant_auto_exchange_on_deposit_receive_add"
	case BILL_TYPE_MERCHANT_AUTO_EXCHANGE_ON_DEPOSIT_RECEIVE_DEDUCT:
		return "merchant_auto_exchange_on_deposit_receive_deduct"

	case BILL_TYPE_MERCHANT_ADJUST_BALANCE_POSITIVE_SEND_DEDUCT:
		return "merchant_adjust_balance_positive_send_deduct"
	case BILL_TYPE_MERCHANT_ADJUST_BALANCE_POSITIVE_SEND_FEE_DEDUCT:
		return "merchant_adjust_balance_positive_send_fee_deduct"
	case BILL_TYPE_MERCHANT_ADJUST_BALANCE_POSITIVE_SEND_ADD:
		return "merchant_adjust_balance_positive_send_add"
	case BILL_TYPE_MERCHANT_ADJUST_BALANCE_POSITIVE_SEND_FEE_ADD:
		return "merchant_adjust_balance_positive_send_fee_add"
	case BILL_TYPE_MERCHANT_ADJUST_BALANCE_POSITIVE_RECEIVE_ADD:
		return "merchant_adjust_balance_positive_receive_add"
	case BILL_TYPE_MERCHANT_ADJUST_BALANCE_POSITIVE_RECEIVE_DEDUCT:
		return "merchant_adjust_balance_positive_receive_deduct"

	case BILL_TYPE_MERCHANT_ADJUST_BALANCE_NEGATIVE_RECEIVE_ADD:
		return "merchant_adjust_balance_negative_receive_add"
	case BILL_TYPE_MERCHANT_ADJUST_BALANCE_NEGATIVE_RECEIVE_FEE_DEDUCT:
		return "merchant_adjust_balance_negative_receive_fee_deduct"
	case BILL_TYPE_MERCHANT_ADJUST_BALANCE_NEGATIVE_RECEIVE_DEDUCT:
		return "merchant_adjust_balance_negative_receive_deduct"
	case BILL_TYPE_MERCHANT_ADJUST_BALANCE_NEGATIVE_RECEIVE_FEE_ADD:
		return "merchant_adjust_balance_negative_receive_fee_add"
	case BILL_TYPE_MERCHANT_ADJUST_BALANCE_NEGATIVE_SEND_DEDUCT:
		return "merchant_adjust_balance_negative_send_deduct"
	case BILL_TYPE_MERCHANT_ADJUST_BALANCE_NEGATIVE_SEND_ADD:
		return "merchant_adjust_balance_negative_send_add"

	case BILL_TYPE_POINT_ACCURAL_POINT_ADD:
		return "point_accural_point_add"
	case BILL_TYPE_POINT_ACCURAL_POINT_DEDUCT:
		return "point_accural_point_deduct"
	case BILL_TYPE_POINT_ACCURAL_WHALE_POINT_ADD:
		return "point_accural_whale_point_add"
	case BILL_TYPE_POINT_ACCURAL_WHALE_POINT_DEDUCT:
		return "point_accural_whale_point_deduct"
	case BILL_TYPE_POINT_ACCURAL_CHIPPAY_EXPRESS_POINT_ADD:
		return "point_accural_chippay_express_point_add"
	case BILL_TYPE_POINT_ACCURAL_CHIPPAY_EXPRESS_POINT_DEDUCT:
		return "point_accural_chippay_express_point_deduct"
	case BILL_TYPE_POINT_ACCURAL_PAYCRYPTO_POINT_ADD:
		return "point_accural_paycrypto_point_add"
	case BILL_TYPE_POINT_ACCURAL_PAYCRYPTO_POINT_DEDUCT:
		return "point_accural_paycrypto_point_deduct"

	case BILL_TYPE_POINT_REDEEM_POINT_DEDUCT:
		return "point_redeem_point_deduct"
	case BILL_TYPE_POINT_REDEEM_POINT_ADD:
		return "point_redeem_point_add"

	case BILL_TYPE_INTEREST_ADD:
		return "interest_add"
	case BILL_TYPE_INTEREST_DEDUCT:
		return "interest_deduct"

	case BILL_TYPE_FINANCIAL_TRANSFER_SEND_DEDUCT:
		return "financial_transfer_send_deduct"
	case BILL_TYPE_FINANCIAL_TRANSFER_RECEIVE_ADD:
		return "financial_transfer_receive_add"

	case BILL_TYPE_WHALE_PAY_CONSUMPTION_PENDING_FREEZE:
		return "whale_pay_consumption_pending_freeze"
	case BILL_TYPE_WHALE_PAY_CONSUMPTION_CLOSED_DEDUCT_FROZEN:
		return "whale_pay_consumption_closed_deduct_frozen"
	case BILL_TYPE_WHALE_PAY_CONSUMPTION_CLOSED_ADD:
		return "whale_pay_consumption_closed_add"
	case BILL_TYPE_WHALE_PAY_CONSUMPTION_CLOSED_UNFREEZE:
		return "whale_pay_consumption_closed_unfreeze"
	case BILL_TYPE_WHALE_PAY_REVERSAL_DEDUCT:
		return "whale_pay_reversal_deduct"
	case BILL_TYPE_WHALE_PAY_REVERSAL_ADD:
		return "whale_pay_reversal_add"
	case BILL_TYPE_WHALE_PAY_REVERSAL_FEE_DEDUCT:
		return "whale_pay_reversal_fee_deduct"
	case BILL_TYPE_WHALE_PAY_REVERSAL_FEE_ADD:
		return "whale_pay_reversal_fee_add"
	case BILL_TYPE_WHALE_PAY_CREDIT_DEDUCT:
		return "whale_pay_credit_deduct"
	case BILL_TYPE_WHALE_PAY_CREDIT_ADD:
		return "whale_pay_credit_add"
	case BILL_TYPE_WHALE_PAY_CREDIT_FEE_DEDUCT:
		return "whale_pay_credit_fee_deduct"
	case BILL_TYPE_WHALE_PAY_CREDIT_FEE_ADD:
		return "whale_pay_credit_fee_add"
	case BILL_TYPE_WHALE_PAY_CONSUMPTION_PENDING_FEE_FREEZE:
		return "whale_pay_consumption_pending_fee_freeze"
	case BILL_TYPE_WHALE_PAY_CONSUMPTION_CLOSED_FEE_DEDUCT_FROZEN:
		return "whale_pay_consumption_closed_fee_deduct_frozen"
	case BILL_TYPE_WHALE_PAY_CONSUMPTION_CLOSED_FEE_ADD:
		return "whale_pay_consumption_closed_fee_add"
	case BILL_TYPE_WHALE_PAY_CONSUMPTION_DIRECT_CLOSED_DEDUCT:
		return "whale_pay_consumption_direct_closed_deduct"
	case BILL_TYPE_WHALE_PAY_CONSUMPTION_DIRECT_CLOSED_ADD:
		return "whale_pay_consumption_direct_closed_add"
	case BILL_TYPE_WHALE_PAY_CONSUMPTION_DIRECT_CLOSED_FEE_DEDUCT:
		return "whale_pay_consumption_direct_closed_fee_deduct"
	case BILL_TYPE_WHALE_PAY_CONSUMPTION_DIRECT_CLOSED_FEE_ADD:
		return "whale_pay_consumption_direct_closed_fee_add"
	case BILL_TYPE_WHALE_PAY_CONSUMPTION_FAIL_UNFREEZE:
		return "whale_pay_consumption_fail_unfreeze"
	case BILL_TYPE_WHALE_PAY_CONSUMPTION_FAIL_FEE_UNFREEZE:
		return "whale_pay_consumption_fail_fee_unfreeze"

	case BILL_TYPE_DISTRIBUTE_DISCOUNT_DEDUCT:
		return "distribute_discount_deduct"
	case BILL_TYPE_DISTRIBUTE_DISCOUNT_ADD:
		return "distribute_discount_add"

	case BILL_TYPE_PAYCRYPTO_PAY_PAY_DEDUCT:
		return "paycrypto_pay_pay_deduct"
	case BILL_TYPE_PAYCRYPTO_PAY_PAY_ADD:
		return "paycrypto_pay_pay_add"
	case BILL_TYPE_PAYCRYPTO_PAY_REVERSAL_DEDUCT:
		return "paycrypto_pay_reversal_deduct"
	case BILL_TYPE_PAYCRYPTO_PAY_REVERSAL_ADD:
		return "paycrypto_pay_reversal_add"
	case BILL_TYPE_PAYCRYPTO_PAY_REFUND_DEDUCT:
		return "paycrypto_pay_refund_deduct"
	case BILL_TYPE_PAYCRYPTO_PAY_REFUND_ADD:
		return "paycrypto_pay_refund_add"
	case BILL_TYPE_PAYCRYPTO_PAY_FAIL_DEDUCT:
		return "paycrypto_pay_fail_deduct"
	case BILL_TYPE_PAYCRYPTO_PAY_FAIL_ADD:
		return "paycrypto_pay_fail_add"
	case BILL_TYPE_PAYCRYPTO_PAY_CLOSE_DEDUCT:
		return "paycrypto_pay_close_deduct"
	case BILL_TYPE_PAYCRYPTO_PAY_CLOSE_ADD:
		return "paycrypto_pay_close_add"

	case BILL_TYPE_BINANCE_PAY_DEPOSIT_RECEIVE_DEDUCT:
		return "binance_pay_deposit_receive_deduct"
	case BILL_TYPE_BINANCE_PAY_DEPOSIT_FEE_DEDUCT:
		return "binance_pay_deposit_fee_deduct"
	case BILL_TYPE_BINANCE_PAY_DEPOSIT_FEE_ADD:
		return "binance_pay_deposit_fee_add"
	case BILL_TYPE_BINANCE_PAY_DEPOSIT_RECEIVE_ADD:
		return "binance_pay_deposit_receive_add"

	case BILL_TYPE_ADMIN_USER_DELETE_ADD:
		return "admin_user_delete_add"
	case BILL_TYPE_ADMIN_USER_DELETE_DEDUT:
		return "admin_user_delete_deduct"
	case BILL_TYPE_ADMIN_USER_DELETE_FREEZE_ADD:
		return "admin_user_delete_freeze_add"
	case BILL_TYPE_ADMIN_USER_DELETE_FREEZE_DEDUT:
		return "admin_user_delete_freeze_deduct"
	}

	return strconv.FormatInt(int64(bt), 10)
}

func (t CallbackType) String() string {
	switch t {
	case CALLBACK_TYPE_MERCHANT_PAY:
		return "merchant_pay"
	case CALLBACK_TYPE_MERCHANT_3DS:
		return "merchant_3ds"
	case CALLBACK_TYPE_MERCHANT_WALLET_OTP:
		return "merchant_wallet_otp"
	case CALLBACK_TYPE_MERCHANT_CARD_STATUS:
		return "merchant_card_status"
	case CALLBACK_TYPE_MERCHANT_EXPORT:
		return "merchant_export"
	case CALLBACK_TYPE_MERCHANT_KYC_RESULT:
		return "merchant_kyc_result"
	}
	return ""
}

func (s CardStatus) String() string {
	switch s {
	case CARD_STATUS_NOT_CREATED:
		return "not_created"
	case CARD_STATUS_CREATED:
		return "created"
	case CARD_STATUS_NOT_ACTIVATED:
		return "not_activated"
	case CARD_STATUS_ACTIVATED:
		return "activated"
	case CARD_STATUS_DEACTIVATED:
		return "deactivated"
	case CARD_STATUS_DELETED:
		return "deleted"
	case CARD_STATUS_BLOCKED:
		return "blocked"
	}
	return ""
}

func (f CardFormat) String() string {
	switch f {
	case CARD_FORMAT_PHYSICAL:
		return "physical"
	case CARD_FORMAT_VIRTUAL:
		return "virtual"
	case CARD_FORMAT_GENERAL:
		return "general"
	}
	return ""
}

func (s CardEarningStatus) String() string {
	switch s {
	case CARD_EARNING_STATUS_ENABLED:
		return "enabled"
	case CARD_EARNING_STATUS_DISABLED:
		return "disabled"
	}
	return ""
}

func (CardEarningStatus) FromString(s string) CardEarningStatus {
	switch s {
	case "enabled":
		return CARD_EARNING_STATUS_ENABLED
	case "disabled":
		return CARD_EARNING_STATUS_DISABLED
	}
	return 0
}

func (s CardFreezeStatus) String() string {
	switch s {
	case CARD_FREEZE_STATUS_UNFROZEN:
		return "unfrozen"
	case CARD_FREEZE_STATUS_FROZEN:
		return "frozen"
	}
	return ""
}

func (s CardFreezeReason) String() string {
	switch s {
	case CARD_FREEZE_REASON_REAP_MANUAL:
		return "reap_manual"
	case CARD_FREEZE_REASON_REAP_NEGATIVE_BALANCE:
		return "reap_negative_balance"
	case CARD_FREEZE_REASON_WHALE_MANUAL:
		return "whale_manual"
	case CARD_FREEZE_REASON_WHALE_BY_SYSTEM:
		return "whale_by_system"
	case CARD_FREEZE_REASON_EUSD_BLOCK:
		return "eusd_block"
	case CARD_FREEZE_REASON_PAYCRYPTO_MANUAL:
		return "paycrypto_manual"
	case CARD_FREEZE_REASON_PAYCRYPTO_DELETED:
		return "paycrypto_deleted"
	}
	return ""
}

func (s CardOrganization) String() string {
	switch s {
	case CARD_ORGANIZATION_VISA:
		return "visa"
	case CARD_ORGANIZATION_MASTERCARD:
		return "mastercard"
	case CARD_ORGANIZATION_AMEX:
		return "amex"
	case CARD_ORGANIZATION_DISCOVER:
		return "discover"
	case CARD_ORGANIZATION_JCB:
		return "jcb"
	case CARD_ORGANIZATION_DINERS:
		return "diners"
	case CARD_ORGANIZATION_UNIONPAY:
		return "unionpay"
	}
	return ""
}

func (cfu CategoryFrontendUsage) String() string {
	// 通常是小寫+底線 但是這個已經跟前端串好了小寫大寫 沒辦法改回去
	switch cfu {
	case CATEGORY_FRONT_USAGE_ENABLE_APPLE_PAY:
		return "enableApplePay"
	case CATEGORY_FRONT_USAGE_ENABLE_DELETE:
		return "enableDelete"
	case CATEGORY_FRONT_USAGE_ENABLE_SPENDING_LIMIT:
		return "enableSpendingLimit"
	case CATEGORY_FRONT_USAGE_ENABLE_FINANCE:
		return "enableFinance"
	case CATEGORY_FRONT_USAGE_ENABLE_QR_CODE:
		return "enableQrCode"
	case CATEGORY_FRONT_USAGE_ENABLE_CONVERT:
		return "enableConvert"
	case CATEGORY_FRONT_USAGE_ENABLE_CHANGELLY:
		return "enableChangelly"
	case CATEGORY_FRONT_USAGE_ENABLE_CHIPPAY:
		return "enableChippay"
	case CATEGORY_FRONT_USAGE_ENABLE_APPLE_DEPOSIT:
		return "enableAppleDeposit"
	case CATEGORY_FRONT_USAGE_ENABLE_SEND:
		return "enableSend"
	case CATEGORY_FRONT_USAGE_ENABLE_TO_CONVERT:
		return "enableToConvert"
	case CATEGORY_FRONT_USAGE_ENABLE_BINANCE_PAY:
		return "enableBinance"
	case CATEGORY_FRONT_USAGE_ENABLE_GOOGLE_PAY:
		return "enableGooglePay"
	case CATEGORY_FRONT_USAGE_ENABLE_PAYPAL:
		return "enablePaypal"
	case CATEGORY_FRONT_USAGE_ENABLE_DEPOSIT:
		return "enableDeposit"
	case CATEGORY_FRONT_USAGE_ENABLE_ACTIVATE:
		return "enableActivate"
	case CATEGORY_FRONT_USAGE_ENABLE_FROM_CONVERSION:
		return "enableFromConversion"
	case CATEGORY_FRONT_USAGE_ENABLE_TO_CONVERSION:
		return "enableToConversion"
	case CATEGORY_FRONT_USAGE_ENABLE_APPLY_SELECT:
		return "enableApplySelect"
	case CATEGORY_FRONT_USAGE_ENABLE_ALI_PAY:
		return "enableAliPay"
	case CATEGORy_FRONT_USAGE_ENABLE_WECHAT_PAY:
		return "enableWechatPay"
	}

	return ""
}

func (c CPPaymentMethod) FromString(s string) CPPaymentMethod {
	switch s {
	case "bank":
		return CPPAYMENT_METHOD_BANK
	case "alipay":
		return CPPAYMENT_METHOD_ALIPAY
	case "wechat_pay":
		return CPPAYMENT_METHOD_WECHAT_PAY
	}
	return 0
}

func (c CPPaymentMethod) String() string {
	switch c {
	case CPPAYMENT_METHOD_BANK:
		return "bank"
	case CPPAYMENT_METHOD_ALIPAY:
		return "alipay"
	case CPPAYMENT_METHOD_WECHAT_PAY:
		return "wechat_pay"
	}
	return ""
}

func (c Currency) String() string {
	switch c {
	// crypto
	case CURRENCY_USDT:
		return "usdt"
	case CURRENCY_ETH:
		return "eth"
	case CURRENCY_USD:
		return "usd"
	case CURRENCY_BTC:
		return "btc"
	case CURRENCY_USDC:
		return "usdc"
	case CURRENCY_DAI:
		return "dai"
	case CURRENCY_WBTC:
		return "wbtc"
	case CURRENCY_TRX:
		return "trx"
	case CURRENCY_ADA:
		return "ada"
	case CURRENCY_BCH:
		return "bch"
	case CURRENCY_DOGE:
		return "doge"
	case CURRENCY_LTC:
		return "ltc"
	case CURRENCY_XRP:
		return "xrp"
	case CURRENCY_SOL:
		return "sol"
	case CURRENCY_BNB:
		return "bnb"
	case CURRENCY_ETC:
		return "etc"
	case CURRENCY_MATIC:
		return "matic"
	case CURRENCY_TON:
		return "ton"
	case CURRENCY_AVAX:
		return "avax"
	// fiat
	case CURRENCY_EUR:
		return "eur"
	case CURRENCY_JPY:
		return "jpy"
	case CURRENCY_GBP:
		return "gbp"
	case CURRENCY_CNY:
		return "cny"
	case CURRENCY_CAD:
		return "cad"
	case CURRENCY_AUD:
		return "aud"
	case CURRENCY_CHF:
		return "chf"
	case CURRENCY_HKD:
		return "hkd"
	case CURRENCY_SEK:
		return "sek"
	case CURRENCY_NZD:
		return "nzd"
	case CURRENCY_MXN:
		return "mxn"
	case CURRENCY_SGD:
		return "sgd"
	case CURRENCY_NOK:
		return "nok"
	case CURRENCY_KRW:
		return "krw"
	case CURRENCY_TRY:
		return "try"
	case CURRENCY_INR:
		return "inr"
	case CURRENCY_BRL:
		return "brl"
	case CURRENCY_ZAR:
		return "zar"
	case CURRENCY_RUB:
		return "rub"
	case CURRENCY_TWD:
		return "twd"
	case CURRENCY_DKK:
		return "dkk"
	case CURRENCY_PLN:
		return "pln"
	case CURRENCY_ILS:
		return "ils"
	case CURRENCY_AED:
		return "aed"
	case CURRENCY_ARS:
		return "ars"
	case CURRENCY_MYR:
		return "myr"
	case CURRENCY_THB:
		return "thb"
	case CURRENCY_SAR:
		return "sar"
	case CURRENCY_CZK:
		return "czk"
	case CURRENCY_UAH:
		return "uah"
	case CURRENCY_PHP:
		return "php"
	case CURRENCY_MOP:
		return "mop"
	case CURRENCY_EGP:
		return "egp"
	case CURRENCY_MAD:
		return "mad"
	case CURRENCY_KES:
		return "kes"
	case CURRENCY_RON:
		return "ron"
	case CURRENCY_BGN:
		return "bgn"
	case CURRENCY_VND:
		return "vnd"
	case CURRENCY_COP:
		return "cop"
	case CURRENCY_OTHER_FIAT:
		return "other_fiat"
	// point
	case CURRENCY_EPOINT:
		return "epoint"
	}
	return ""
}

func (c Currency) IsValid() bool {
	switch c {
	case
		// crypto
		CURRENCY_USDT,
		CURRENCY_ETH,
		CURRENCY_BTC,
		CURRENCY_USDC,
		CURRENCY_DAI,
		CURRENCY_WBTC,
		CURRENCY_TRX,
		CURRENCY_ADA,
		CURRENCY_BCH,
		CURRENCY_DOGE,
		CURRENCY_LTC,
		CURRENCY_XRP,
		CURRENCY_SOL,
		CURRENCY_BNB,
		CURRENCY_ETC,
		CURRENCY_MATIC,
		// fiat
		CURRENCY_USD,
		CURRENCY_EUR,
		CURRENCY_JPY,
		CURRENCY_GBP,
		CURRENCY_CNY,
		CURRENCY_CAD,
		CURRENCY_AUD,
		CURRENCY_CHF,
		CURRENCY_HKD,
		CURRENCY_SEK,
		CURRENCY_NZD,
		CURRENCY_MXN,
		CURRENCY_SGD,
		CURRENCY_NOK,
		CURRENCY_KRW,
		CURRENCY_TRY,
		CURRENCY_INR,
		CURRENCY_BRL,
		CURRENCY_ZAR,
		CURRENCY_RUB,
		CURRENCY_TWD,
		CURRENCY_DKK,
		CURRENCY_PLN,
		CURRENCY_ILS,
		CURRENCY_AED,
		CURRENCY_ARS,
		CURRENCY_MYR,
		CURRENCY_THB,
		CURRENCY_SAR,
		CURRENCY_CZK,
		CURRENCY_UAH,
		CURRENCY_PHP,
		CURRENCY_MOP,
		CURRENCY_EGP,
		CURRENCY_MAD,
		CURRENCY_KES,
		CURRENCY_RON,
		CURRENCY_BGN,
		CURRENCY_VND,
		CURRENCY_COP,
		CURRENCY_OTHER_FIAT,
		// point
		CURRENCY_EPOINT:
		return true
	}
	return false
}

func (Currency) FromString(s string) Currency {
	sLower := strings.ToLower(s)
	switch sLower {
	// crypto
	case "usdt":
		return CURRENCY_USDT
	case "eth":
		return CURRENCY_ETH
	case "btc":
		return CURRENCY_BTC
	case "usdc":
		return CURRENCY_USDC
	case "dai":
		return CURRENCY_DAI
	case "wbtc":
		return CURRENCY_WBTC
	case "trx":
		return CURRENCY_TRX
	case "ada":
		return CURRENCY_ADA
	case "bch":
		return CURRENCY_BCH
	case "doge":
		return CURRENCY_DOGE
	case "ltc":
		return CURRENCY_LTC
	case "xrp":
		return CURRENCY_XRP
	case "sol":
		return CURRENCY_SOL
	case "bnb":
		return CURRENCY_BNB
	case "etc":
		return CURRENCY_ETC
	case "matic":
		return CURRENCY_MATIC
	case "ton":
		return CURRENCY_TON
	case "avax":
		return CURRENCY_AVAX

	// fiat
	case "usd":
		return CURRENCY_USD
	case "eur":
		return CURRENCY_EUR
	case "jpy":
		return CURRENCY_JPY
	case "gbp":
		return CURRENCY_GBP
	case "cny":
		return CURRENCY_CNY
	case "cad":
		return CURRENCY_CAD
	case "aud":
		return CURRENCY_AUD
	case "chf":
		return CURRENCY_CHF
	case "hkd":
		return CURRENCY_HKD
	case "sek":
		return CURRENCY_SEK
	case "nzd":
		return CURRENCY_NZD
	case "mxn":
		return CURRENCY_MXN
	case "sgd":
		return CURRENCY_SGD
	case "nok":
		return CURRENCY_NOK
	case "krw":
		return CURRENCY_KRW
	case "try":
		return CURRENCY_TRY
	case "inr":
		return CURRENCY_INR
	case "brl":
		return CURRENCY_BRL
	case "zar":
		return CURRENCY_ZAR
	case "rub":
		return CURRENCY_RUB
	case "twd":
		return CURRENCY_TWD
	case "dkk":
		return CURRENCY_DKK
	case "pln":
		return CURRENCY_PLN
	case "ils":
		return CURRENCY_ILS
	case "aed":
		return CURRENCY_AED
	case "ars":
		return CURRENCY_ARS
	case "myr":
		return CURRENCY_MYR
	case "thb":
		return CURRENCY_THB
	case "sar":
		return CURRENCY_SAR
	case "czk":
		return CURRENCY_CZK
	case "uah":
		return CURRENCY_UAH
	case "php":
		return CURRENCY_PHP
	case "mop":
		return CURRENCY_MOP
	case "egp":
		return CURRENCY_EGP
	case "mad":
		return CURRENCY_MAD
	case "kes":
		return CURRENCY_KES
	case "ron":
		return CURRENCY_RON
	case "bgn":
		return CURRENCY_BGN
	case "vnd":
		return CURRENCY_VND
	case "cop":
		return CURRENCY_COP
	case "other_fiat":
		return CURRENCY_OTHER_FIAT

	// point
	case "epoint":
		return CURRENCY_EPOINT
	}
	return 0
}

func (Currency) FromISO4217(s string) Currency {
	// https://en.wikipedia.org/wiki/ISO_4217
	switch s {
	case "840":
		return CURRENCY_USD
	case "978":
		return CURRENCY_EUR
	case "392":
		return CURRENCY_JPY
	case "826":
		return CURRENCY_GBP
	case "156":
		return CURRENCY_CNY
	case "124":
		return CURRENCY_CAD
	case "036":
		return CURRENCY_AUD
	case "756":
		return CURRENCY_CHF
	case "344":
		return CURRENCY_HKD
	case "752":
		return CURRENCY_SEK
	case "554":
		return CURRENCY_NZD
	case "484":
		return CURRENCY_MXN
	case "702":
		return CURRENCY_SGD
	case "578":
		return CURRENCY_NOK
	case "410":
		return CURRENCY_KRW
	case "949":
		return CURRENCY_TRY
	case "356":
		return CURRENCY_INR
	case "986":
		return CURRENCY_BRL
	case "710":
		return CURRENCY_ZAR
	case "643":
		return CURRENCY_RUB
	case "901":
		return CURRENCY_TWD
	case "208":
		return CURRENCY_DKK
	case "985":
		return CURRENCY_PLN
	case "376":
		return CURRENCY_ILS
	case "784":
		return CURRENCY_AED
	case "032":
		return CURRENCY_ARS
	case "458":
		return CURRENCY_MYR
	case "764":
		return CURRENCY_THB
	case "682":
		return CURRENCY_SAR
	case "203":
		return CURRENCY_CZK
	case "980":
		return CURRENCY_UAH
	case "608":
		return CURRENCY_PHP
	case "446":
		return CURRENCY_MOP
	case "818":
		return CURRENCY_EGP
	case "504":
		return CURRENCY_MAD
	case "404":
		return CURRENCY_KES
	case "946":
		return CURRENCY_RON
	case "975":
		return CURRENCY_BGN
	case "704":
		return CURRENCY_VND
	case "170":
		return CURRENCY_COP
	}
	return CURRENCY_OTHER_FIAT
}

func ISO4217CurrencyString(s string) string {
	switch s {
	case "840": // United States
		return "usd"
	case "156": // China
		return "cny"
	case "392": // Japan
		return "jpy"
	case "978": // European Union
		return "eur"
	case "276": // Germany
		return "eur"
	case "250": // France
		return "eur"
	case "380": // Italy
		return "eur"
	case "724": // Spain
		return "eur"
	case "528": // Netherlands
		return "eur"
	case "056": // Belgium
		return "eur"
	case "040": // Austria
		return "eur"
	case "620": // Portugal
		return "eur"
	case "300": // Greece
		return "eur"
	case "246": // Finland
		return "eur"
	case "372": // Ireland
		return "eur"
	case "442": // Luxembourg
		return "eur"
	case "705": // Slovenia
		return "eur"
	case "196": // Cyprus
		return "eur"
	case "470": // Malta
		return "eur"
	case "703": // Slovakia
		return "eur"
	case "233": // Estonia
		return "eur"
	case "428": // Latvia
		return "eur"
	case "440": // Lithuania
		return "eur"
	case "191": // Croatia
		return "eur"
	case "826": // United Kingdom
		return "gbp"
	case "124": // Canada
		return "cad"
	case "036": // Australia
		return "aud"
	case "756": // Switzerland
		return "chf"
	case "752": // Sweden
		return "sek"
	case "578": // Norway
		return "nok"
	case "208": // Denmark
		return "dkk"
	case "702": // Singapore
		return "sgd"
	case "344": // Hong Kong
		return "hkd"
	case "158": // Taiwan
		return "twd"
	case "410": // South Korea
		return "krw"
	case "764": // Thailand
		return "thb"
	case "458": // Malaysia
		return "myr"
	case "360": // Indonesia
		return "idr"
	case "608": // Philippines
		return "php"
	case "704": // Vietnam
		return "vnd"
	case "356": // India
		return "inr"
	case "643": // Russia
		return "rub"
	case "076": // Brazil
		return "brl"
	case "484": // Mexico
		return "mxn"
	case "710": // South Africa
		return "zar"
	case "784": // United Arab Emirates
		return "aed"
	case "682": // Saudi Arabia
		return "sar"
	case "376": // Israel
		return "ils"
	case "792": // Turkey
		return "try"
	case "032": // Argentina
		return "ars"
	case "152": // Chile
		return "clp"
	case "170": // Colombia
		return "cop"
	case "604": // Peru
		return "pen"
	case "858": // Uruguay
		return "uyu"
	case "818": // Egypt
		return "egp"
	case "504": // Morocco
		return "mad"
	case "012": // Algeria
		return "dzd"
	case "404": // Kenya
		return "kes"
	case "566": // Nigeria
		return "ngn"
	case "050": // Bangladesh
		return "bdt"
	case "586": // Pakistan
		return "pkr"
	case "144": // Sri Lanka
		return "lkr"
	case "096": // Brunei
		return "bnd"
	case "554": // New Zealand
		return "nzd"
	case "242": // Fiji
		return "fjd"
	case "980": // Ukraine
		return "uah"
	case "446": // Macau
		return "mop"
	case "946": // Romania
		return "ron"
	case "975": // Bulgaria
		return "bgn"
	case "971":
		return "afn"
	case "008":
		return "all"
	case "051":
		return "amd"
	case "973":
		return "aoa"
	case "533":
		return "awg"
	case "944":
		return "azn"
	case "977":
		return "bam"
	case "052":
		return "bbd"
	case "048":
		return "bhd"
	case "108":
		return "bif"
	case "060":
		return "bmd"
	case "068":
		return "bob"
	case "984":
		return "bov"
	case "986":
		return "brl"
	case "044":
		return "bsd"
	case "064":
		return "btn"
	case "072":
		return "bwp"
	case "933":
		return "byn"
	case "084":
		return "bzd"
	case "976":
		return "cdf"
	case "947":
		return "che"
	case "948":
		return "chw"
	case "990":
		return "clf"
	case "970":
		return "cou"
	case "188":
		return "crc"
	case "192":
		return "cup"
	case "132":
		return "cve"
	case "203":
		return "czk"
	case "262":
		return "djf"
	case "214":
		return "dop"
	case "232":
		return "ern"
	case "230":
		return "etb"
	case "238":
		return "fkp"
	case "981":
		return "gel"
	case "936":
		return "ghs"
	case "292":
		return "gip"
	case "270":
		return "gmd"
	case "324":
		return "gnf"
	case "320":
		return "gtq"
	case "328":
		return "gyd"
	case "340":
		return "hnl"
	case "332":
		return "htg"
	case "348":
		return "huf"
	case "368":
		return "iqd"
	case "364":
		return "irr"
	case "352":
		return "isk"
	case "388":
		return "jmd"
	case "400":
		return "jod"
	case "417":
		return "kgs"
	case "116":
		return "khr"
	case "174":
		return "kmf"
	case "408":
		return "kpw"
	case "414":
		return "kwd"
	case "136":
		return "kyd"
	case "398":
		return "kzt"
	case "418":
		return "lak"
	case "422":
		return "lbp"
	case "430":
		return "lrd"
	case "426":
		return "lsl"
	case "434":
		return "lyd"
	case "498":
		return "mdl"
	case "969":
		return "mga"
	case "807":
		return "mkd"
	case "104":
		return "mmk"
	case "496":
		return "mnt"
	case "929":
		return "mru"
	case "480":
		return "mur"
	case "462":
		return "mvr"
	case "454":
		return "mwk"
	case "979":
		return "mxv"
	case "943":
		return "mzn"
	case "516":
		return "nad"
	case "558":
		return "nio"
	case "524":
		return "npr"
	case "512":
		return "omr"
	case "590":
		return "pab"
	case "598":
		return "pgk"
	case "985":
		return "pln"
	case "600":
		return "pyg"
	case "634":
		return "qar"
	case "941":
		return "rsd"
	case "646":
		return "rwf"
	case "090":
		return "sbd"
	case "690":
		return "scr"
	case "938":
		return "sdg"
	case "654":
		return "shp"
	case "925":
		return "sle"
	case "706":
		return "sos"
	case "968":
		return "srd"
	case "728":
		return "ssp"
	case "930":
		return "stn"
	case "222":
		return "svc"
	case "760":
		return "syp"
	case "748":
		return "szl"
	case "972":
		return "tjs"
	case "934":
		return "tmt"
	case "788":
		return "tnd"
	case "776":
		return "top"
	case "949":
		return "try"
	case "780":
		return "ttd"
	case "901":
		return "twd"
	case "834":
		return "tzs"
	case "800":
		return "ugx"
	case "997":
		return "usn"
	case "940":
		return "uyi"
	case "927":
		return "uyw"
	case "860":
		return "uzs"
	case "926":
		return "ved"
	case "928":
		return "ves"
	case "548":
		return "vuv"
	case "882":
		return "wst"
	case "396":
		return "xad"
	case "950":
		return "xaf"
	case "961":
		return "xag"
	case "959":
		return "xau"
	case "955":
		return "xba"
	case "956":
		return "xbb"
	case "957":
		return "xbc"
	case "958":
		return "xbd"
	case "951":
		return "xcd"
	case "532":
		return "xcg"
	case "960":
		return "xdr"
	case "952":
		return "xof"
	case "964":
		return "xpd"
	case "953":
		return "xpf"
	case "962":
		return "xpt"
	case "994":
		return "xsu"
	case "963":
		return "xts"
	case "965":
		return "xua"
	case "999":
		return "xxx"
	case "886":
		return "yer"
	case "967":
		return "zmw"

	default:
		return "other_fiat"
	}
}

func (c Currency) ISO4217() string {
	// https://en.wikipedia.org/wiki/ISO_4217
	switch c {
	case CURRENCY_USD:
		return "840"
	case CURRENCY_EUR:
		return "978"
	case CURRENCY_JPY:
		return "392"
	case CURRENCY_GBP:
		return "826"
	case CURRENCY_CNY:
		return "156"
	case CURRENCY_CAD:
		return "124"
	case CURRENCY_AUD:
		return "036"
	case CURRENCY_CHF:
		return "756"
	case CURRENCY_HKD:
		return "344"
	case CURRENCY_SEK:
		return "752"
	case CURRENCY_NZD:
		return "554"
	case CURRENCY_MXN:
		return "484"
	case CURRENCY_SGD:
		return "702"
	case CURRENCY_NOK:
		return "578"
	case CURRENCY_KRW:
		return "410"
	case CURRENCY_TRY:
		return "949"
	case CURRENCY_INR:
		return "356"
	case CURRENCY_BRL:
		return "986"
	case CURRENCY_ZAR:
		return "710"
	case CURRENCY_RUB:
		return "643"
	case CURRENCY_TWD:
		return "901"
	case CURRENCY_DKK:
		return "208"
	case CURRENCY_PLN:
		return "985"
	case CURRENCY_ILS:
		return "376"
	case CURRENCY_AED:
		return "784"
	case CURRENCY_ARS:
		return "032"
	case CURRENCY_MYR:
		return "458"
	case CURRENCY_THB:
		return "764"
	case CURRENCY_SAR:
		return "682"
	case CURRENCY_CZK:
		return "203"
	case CURRENCY_UAH:
		return "980"
	case CURRENCY_PHP:
		return "608"
	case CURRENCY_MOP:
		return "446"
	case CURRENCY_EGP:
		return "818"
	case CURRENCY_MAD:
		return "504"
	case CURRENCY_KES:
		return "404"
	case CURRENCY_RON:
		return "946"
	case CURRENCY_BGN:
		return "975"
	case CURRENCY_VND:
		return "704"
	case CURRENCY_COP:
		return "170"
	}
	return ""
}

func (c Currency) Type() CurrencyType {
	switch true {
	case c >= 100 && c < 200:
		return CURRENCY_TYPE_CRYPTO
	case c >= 200 && c < 300:
		return CURRENCY_TYPE_FIAT
	case c >= 1000 && c < 2000:
		return CURRENCY_TYPE_POINT
	}
	return 0
}

func (c Currency) AssetType() AssetType {
	switch {
	case c >= 100 && c < 200:
		return ASSET_TYPE_CRYPTO
	case c >= 200 && c < 300:
		return ASSET_TYPE_FIAT
	case c >= 1000 && c < 2000:
		return ASSET_TYPE_POINT
	case c > 2000 && c < 3000:
		return ASSET_TYPE_AUTO_YIELD
	}
	return 0
}

func (ct CurrencyType) String() string {
	switch ct {
	case CURRENCY_TYPE_CRYPTO:
		return "crypto"
	case CURRENCY_TYPE_FIAT:
		return "fiat"
	case CURRENCY_TYPE_POINT:
		return "point"
	}
	return ""
}

func (el EPointLevel) String() string {
	switch el {
	case EPOINT_LEVEL_BRONZE:
		return "bronze"
	case EPOINT_LEVEL_SILVER:
		return "silver"
	case EPOINT_LEVEL_GOLD:
		return "gold"
	}
	return ""
}

func (es ExternalRequestLogStage) String() string {
	switch es {
	case EXTERNAL_REQUEST_LOG_STAGE_PRELOG:
		return "prelog"
	case EXTERNAL_REQUEST_LOG_STAGE_UPDATE_1:
		return "update_1"
	case EXTERNAL_REQUEST_LOG_STAGE_UPDATE_2:
		return "update_2"
	case EXTERNAL_REQUEST_LOG_STAGE_UPDATE_3:
		return "update_3"
	case EXTERNAL_REQUEST_LOG_STAGE_POSTLOG:
		return "postlog"
	}
	return ""
}

func (f FilePurpose) String() string {
	switch f {
	case FILE_PURPOSE_KYC2_FACE_PHOTO1:
		return "kyc2_face_photo1"
	case FILE_PURPOSE_KYC3_FACE_PHOTO1:
		return "kyc3_face_photo1"
	case FILE_PURPOSE_KYC3_FACE_PHOTO2:
		return "kyc3_face_photo2"
	case FILE_PURPOSE_KYC3_DOCUMENT_PHOTO:
		return "kyc3_document_photo"
	case FILE_PURPOSE_AGENT_REPORT:
		return "agent_report"
	case FILE_PURPOSE_MERCHANT_BILL:
		return "merchant_bill"
	}
	return ""
}

func (f FinancialProductStatus) String() string {
	switch f {
	case FINANCIAL_PRODUCT_STATUS_ACTIVE:
		return "active"
	case FINANCIAL_PRODUCT_STATUS_INACTIVE:
		return "inactive"
	}
	return ""
}

func (ft ForwardType) String() string {
	switch ft {
	case FORWARD_TYPE_REAP:
		return "reap"
	case FORWARD_TYPE_EMAIL:
		return "email"
	case FORWARD_TYPE_SMS:
		return "sms"
	case FORWARD_TYPE_APP:
		return "app"
	case FORWARD_TYPE_MERCHANT:
		return "merchant"
	}
	return ""
}

func (i IdentityDocumentType) String() string {
	switch i {
	case IDENTITY_DOCUMENT_TYPE_PASSPORT:
		return "Passport"
	case IDENTITY_DOCUMENT_TYPE_HEALTH:
		return "Health"
	case IDENTITY_DOCUMENT_TYPE_NATIONAL_ID:
		return "NationalID"
	case IDENTITY_DOCUMENT_TYPE_TAX_ID_NUMBER:
		return "TaxIDNumber"
	case IDENTITY_DOCUMENT_TYPE_SOCIAL_SERVICE:
		return "SocialService"
	case IDENTITY_DOCUMENT_TYPE_DRIVER:
		return "DriversLicense"
	}
	return ""
}

func (m Mainnet) String() string {
	switch m {
	case MAINNET_BTC:
		return "btc"
	case MAINNET_ETH:
		return "eth"
	case MAINNET_TRX:
		return "trx"
	case MAINNET_ADA:
		return "ada"
	case MAINNET_BCH:
		return "bch"
	case MAINNET_DOGE:
		return "doge"
	case MAINNET_LTC:
		return "ltc"
	case MAINNET_XRP:
		return "xrp"
	case MAINNET_SOL:
		return "sol"
	case MAINNET_BSC:
		return "bsc"
	case MAINNET_ETC:
		return "etc"
	case MAINNET_MATIC:
		return "matic"
	case MAINNET_TON:
		return "ton"
	case MAINNET_AVAXC:
		return "avaxc"
	}
	return ""
}

func (Mainnet) FromString(s string) Mainnet {
	sLower := strings.ToLower(s)
	switch sLower {
	case "btc":
		return MAINNET_BTC
	case "eth":
		return MAINNET_ETH
	case "trx":
		return MAINNET_TRX
	case "ada":
		return MAINNET_ADA
	case "bch":
		return MAINNET_BCH
	case "doge":
		return MAINNET_DOGE
	case "ltc":
		return MAINNET_LTC
	case "xrp":
		return MAINNET_XRP
	case "sol":
		return MAINNET_SOL
	case "bsc":
		return MAINNET_BSC
	case "etc":
		return MAINNET_ETC
	case "matic":
		return MAINNET_MATIC
	case "ton":
		return MAINNET_TON
	case "avaxc":
		return MAINNET_AVAXC
	}
	return 0
}

func (m Mainnet) IsValid() bool {
	switch m {
	case
		MAINNET_BTC,
		MAINNET_ETH,
		MAINNET_TRX,
		MAINNET_ADA,
		MAINNET_BCH,
		MAINNET_DOGE,
		MAINNET_LTC,
		MAINNET_XRP,
		MAINNET_SOL,
		MAINNET_BSC,
		MAINNET_ETC,
		MAINNET_MATIC,
		MAINNET_TON,
		MAINNET_AVAXC:
		return true
	}
	return false
}

func (ms MerchantStatus) String() string {
	switch ms {
	case MERCHANT_STATUS_ACTIVE:
		return "active"
	case MERCHANT_STATUS_DISABLED:
		return "disabled"
	case MERCHANT_STATUS_FROZEN:
		return "frozen"
	case MERCHANT_STATUS_DELETED:
		return "deleted"
	}
	return ""
}

func (mp MessagePurpose) String() string {
	switch mp {
	case MESSAGE_PURPOSE_USER_LOGIN_OR_REGISTER:
		return "user_login_or_register"
	case MESSAGE_PURPOSE_VERIFIED:
		return "user_verify"
	case MESSAGE_PURPOSE_FORGET_PIN_CODE:
		return "forget_pin_code"
	}
	return ""
}

func (nc NationCode) IsValid() bool {
	switch nc {
	case NATION_CODE_USA,
		NATION_CODE_JPN,
		NATION_CODE_AUT,
		NATION_CODE_BEL,
		NATION_CODE_BGR,
		NATION_CODE_CYP,
		NATION_CODE_CZE,
		NATION_CODE_DEU,
		NATION_CODE_DNK,
		NATION_CODE_ESP,
		NATION_CODE_EST,
		NATION_CODE_FIN,
		NATION_CODE_FRA,
		NATION_CODE_GRC,
		NATION_CODE_HUN,
		NATION_CODE_IRL,
		NATION_CODE_ITA,
		NATION_CODE_LAT,
		NATION_CODE_LIT,
		NATION_CODE_LUX,
		NATION_CODE_MLT,
		NATION_CODE_NLD,
		NATION_CODE_POL,
		NATION_CODE_PRT,
		NATION_CODE_ROU,
		NATION_CODE_SVK,
		NATION_CODE_SVN,
		NATION_CODE_SWE,
		NATION_CODE_GBR,
		NATION_CODE_CHN,
		NATION_CODE_CAN,
		NATION_CODE_AUS,
		NATION_CODE_CHE,
		NATION_CODE_HKG,
		NATION_CODE_NZL,
		NATION_CODE_MEX,
		NATION_CODE_SGP,
		NATION_CODE_NOR,
		NATION_CODE_KOR,
		NATION_CODE_TUR,
		NATION_CODE_IND,
		NATION_CODE_BRA,
		NATION_CODE_ZAF,
		NATION_CODE_RUS,
		NATION_CODE_TWN,
		NATION_CODE_ISR,
		NATION_CODE_ARE,
		NATION_CODE_ARG,
		NATION_CODE_MYS,
		NATION_CODE_THA,
		NATION_CODE_SAU,
		NATION_CODE_IDN,
		NATION_CODE_PAK,
		NATION_CODE_NGA,
		NATION_CODE_EGY,
		NATION_CODE_VNM,
		NATION_CODE_PHL,
		NATION_CODE_COL,
		NATION_CODE_CHL,
		NATION_CODE_PER,
		NATION_CODE_DZA,
		NATION_CODE_MAR,
		NATION_CODE_KEN,
		NATION_CODE_UKR,
		NATION_CODE_AFG,
		NATION_CODE_ETH,
		NATION_CODE_COD,
		NATION_CODE_TZA,
		NATION_CODE_UZB,
		NATION_CODE_ALA,
		NATION_CODE_ALB,
		NATION_CODE_ASM,
		NATION_CODE_AND,
		NATION_CODE_AGO,
		NATION_CODE_AIA,
		NATION_CODE_ATA,
		NATION_CODE_ATG,
		NATION_CODE_ARM,
		NATION_CODE_ABW,
		NATION_CODE_AZE,
		NATION_CODE_BHS,
		NATION_CODE_BHR,
		NATION_CODE_BGD,
		NATION_CODE_BRB,
		NATION_CODE_BLR,
		NATION_CODE_BLZ,
		NATION_CODE_BEN,
		NATION_CODE_BMU,
		NATION_CODE_BTN,
		NATION_CODE_BOL,
		NATION_CODE_BES,
		NATION_CODE_BIH,
		NATION_CODE_BWA,
		NATION_CODE_BVT,
		NATION_CODE_IOT,
		NATION_CODE_BRN,
		NATION_CODE_BFA,
		NATION_CODE_BDI,
		NATION_CODE_CPV,
		NATION_CODE_KHM,
		NATION_CODE_CMR,
		NATION_CODE_CYM,
		NATION_CODE_CAF,
		NATION_CODE_TCD,
		NATION_CODE_CXR,
		NATION_CODE_CCK,
		NATION_CODE_COM,
		NATION_CODE_COG,
		NATION_CODE_COK,
		NATION_CODE_CRI,
		NATION_CODE_CIV,
		NATION_CODE_HRV,
		NATION_CODE_CUB,
		NATION_CODE_CUW,
		NATION_CODE_DJI,
		NATION_CODE_DMA,
		NATION_CODE_DOM,
		NATION_CODE_ECU,
		NATION_CODE_SLV,
		NATION_CODE_GNQ,
		NATION_CODE_ERI,
		NATION_CODE_SWZ,
		NATION_CODE_FLK,
		NATION_CODE_FRO,
		NATION_CODE_FJI,
		NATION_CODE_GUF,
		NATION_CODE_PYF,
		NATION_CODE_ATF,
		NATION_CODE_GAB,
		NATION_CODE_GMB,
		NATION_CODE_GEO,
		NATION_CODE_GHA,
		NATION_CODE_GIB,
		NATION_CODE_GRL,
		NATION_CODE_GRD,
		NATION_CODE_GLP,
		NATION_CODE_GUM,
		NATION_CODE_GTM,
		NATION_CODE_GGY,
		NATION_CODE_GIN,
		NATION_CODE_GNB,
		NATION_CODE_GUY,
		NATION_CODE_HTI,
		NATION_CODE_HMD,
		NATION_CODE_VAT,
		NATION_CODE_HND,
		NATION_CODE_ISL,
		NATION_CODE_IRN,
		NATION_CODE_IRQ,
		NATION_CODE_IMN,
		NATION_CODE_JAM,
		NATION_CODE_JEY,
		NATION_CODE_JOR,
		NATION_CODE_KAZ,
		NATION_CODE_KIR,
		NATION_CODE_PRK,
		NATION_CODE_KWT,
		NATION_CODE_KGZ,
		NATION_CODE_LAO,
		NATION_CODE_LVA,
		NATION_CODE_LBN,
		NATION_CODE_LSO,
		NATION_CODE_LBR,
		NATION_CODE_LBY,
		NATION_CODE_LIE,
		NATION_CODE_LTU,
		NATION_CODE_MAC,
		NATION_CODE_MDG,
		NATION_CODE_MWI,
		NATION_CODE_MDV,
		NATION_CODE_MLI,
		NATION_CODE_MHL,
		NATION_CODE_MTQ,
		NATION_CODE_MRT,
		NATION_CODE_MUS,
		NATION_CODE_MYT,
		NATION_CODE_FSM,
		NATION_CODE_MDA,
		NATION_CODE_MCO,
		NATION_CODE_MNG,
		NATION_CODE_MNE,
		NATION_CODE_MSR,
		NATION_CODE_MOZ,
		NATION_CODE_MMR,
		NATION_CODE_NAM,
		NATION_CODE_NRU,
		NATION_CODE_NPL,
		NATION_CODE_NCL,
		NATION_CODE_NIC,
		NATION_CODE_NER,
		NATION_CODE_NIU,
		NATION_CODE_NFK,
		NATION_CODE_MKD,
		NATION_CODE_MNP,
		NATION_CODE_OMN,
		NATION_CODE_PLW,
		NATION_CODE_PSE,
		NATION_CODE_PAN,
		NATION_CODE_PNG,
		NATION_CODE_PRY,
		NATION_CODE_PCN,
		NATION_CODE_PRI,
		NATION_CODE_QAT,
		NATION_CODE_REU,
		NATION_CODE_RWA,
		NATION_CODE_BLM,
		NATION_CODE_SHN,
		NATION_CODE_KNA,
		NATION_CODE_LCA,
		NATION_CODE_MAF,
		NATION_CODE_SPM,
		NATION_CODE_VCT,
		NATION_CODE_WSM,
		NATION_CODE_SMR,
		NATION_CODE_STP,
		NATION_CODE_SEN,
		NATION_CODE_SRB,
		NATION_CODE_SYC,
		NATION_CODE_SLE,
		NATION_CODE_SXM,
		NATION_CODE_SLB,
		NATION_CODE_SOM,
		NATION_CODE_SGS,
		NATION_CODE_SSD,
		NATION_CODE_LKA,
		NATION_CODE_SDN,
		NATION_CODE_SUR,
		NATION_CODE_SJM,
		NATION_CODE_SYR,
		NATION_CODE_TJK,
		NATION_CODE_TLS,
		NATION_CODE_TGO,
		NATION_CODE_TKL,
		NATION_CODE_TON,
		NATION_CODE_TTO,
		NATION_CODE_TUN,
		NATION_CODE_TKM,
		NATION_CODE_TCA,
		NATION_CODE_TUV,
		NATION_CODE_UGA,
		NATION_CODE_UMI,
		NATION_CODE_URY,
		NATION_CODE_VUT,
		NATION_CODE_VEN,
		NATION_CODE_VGB,
		NATION_CODE_VIR,
		NATION_CODE_WLF,
		NATION_CODE_ESH,
		NATION_CODE_YEM,
		NATION_CODE_ZMB,
		NATION_CODE_ZWE:
		return true
	default:
		return false
	}
}

func (nc NationCode) NationCodeFromISO4217(s string) NationCode {
	switch s {
	case "840":
		return NATION_CODE_USA
	case "392":
		return NATION_CODE_JPN
	case "040":
		return NATION_CODE_AUT
	case "056":
		return NATION_CODE_BEL
	case "100":
		return NATION_CODE_BGR
	case "196":
		return NATION_CODE_CYP
	case "203":
		return NATION_CODE_CZE
	case "276":
		return NATION_CODE_DEU
	case "208":
		return NATION_CODE_DNK
	case "220":
		return NATION_CODE_ESP
	case "233":
		return NATION_CODE_EST
	case "246":
		return NATION_CODE_FIN
	case "250":
		return NATION_CODE_FRA
	case "300":
		return NATION_CODE_GRC
	case "348":
		return NATION_CODE_HUN
	case "372":
		return NATION_CODE_IRL
	case "380":
		return NATION_CODE_ITA
	case "428":
		return NATION_CODE_LAT
	case "440":
		return NATION_CODE_LIT
	case "442":
		return NATION_CODE_LUX
	case "470":
		return NATION_CODE_MLT
	case "528":
		return NATION_CODE_NLD
	case "616":
		return NATION_CODE_POL
	case "620":
		return NATION_CODE_PRT
	case "642":
		return NATION_CODE_ROU
	case "703":
		return NATION_CODE_SVK
	case "705":
		return NATION_CODE_SVN
	case "752":
		return NATION_CODE_SWE
	case "826":
		return NATION_CODE_GBR
	case "156":
		return NATION_CODE_CHN
	case "124":
		return NATION_CODE_CAN
	case "036":
		return NATION_CODE_AUS
	case "756":
		return NATION_CODE_CHE
	case "344":
		return NATION_CODE_HKG
	case "554":
		return NATION_CODE_NZL
	case "484":
		return NATION_CODE_MEX
	case "702":
		return NATION_CODE_SGP
	case "578":
		return NATION_CODE_NOR
	case "410":
		return NATION_CODE_KOR
	case "949":
		return NATION_CODE_TUR
	case "356":
		return NATION_CODE_IND
	case "986":
		return NATION_CODE_BRA
	case "710":
		return NATION_CODE_ZAF
	case "643":
		return NATION_CODE_RUS
	case "158":
		return NATION_CODE_TWN
	case "376":
		return NATION_CODE_ISR
	case "784":
		return NATION_CODE_ARE
	case "032":
		return NATION_CODE_ARG
	case "458":
		return NATION_CODE_MYS
	case "764":
		return NATION_CODE_THA
	case "682":
		return NATION_CODE_SAU
	case "360": // 印度尼西亞 (Indonesia)
		return NATION_CODE_IDN
	case "586": // 巴基斯坦 (Pakistan)
		return NATION_CODE_PAK
	case "566": // 奈及利亞 (Nigeria)
		return NATION_CODE_NGA
	case "818": // 埃及 (Egypt)
		return NATION_CODE_EGY
	case "704": // 越南 (Vietnam)
		return NATION_CODE_VNM
	case "608": // 菲律賓 (Philippines)
		return NATION_CODE_PHL
	case "170": // 哥倫比亞 (Colombia)
		return NATION_CODE_COL
	case "152": // 智利 (Chile)
		return NATION_CODE_CHL
	case "604": // 秘魯 (Peru)
		return NATION_CODE_PER
	case "012": // 阿爾及利亞 (Algeria)
		return NATION_CODE_DZA
	case "504": // 摩洛哥 (Morocco)
		return NATION_CODE_MAR
	case "404": // 肯亞 (Kenya)
		return NATION_CODE_KEN
	case "804": // 烏克蘭 (Ukraine)
		return NATION_CODE_UKR
	case "004": // 阿富汗 (Afghanistan)
		return NATION_CODE_AFG
	case "231": // 埃塞俄比亞 (Ethiopia)
		return NATION_CODE_ETH
	case "180": // 剛果民主共和國 (Democratic Republic of the Congo)
		return NATION_CODE_COD
	case "834": // 坦尚尼亞 (Tanzania)
		return NATION_CODE_TZA
	case "860": // 烏茲別克 (Uzbekistan)
		return NATION_CODE_UZB

	case "248":
		return NATION_CODE_ALA
	case "008":
		return NATION_CODE_ALB
	case "016":
		return NATION_CODE_ASM
	case "020":
		return NATION_CODE_AND
	case "024":
		return NATION_CODE_AGO
	case "660":
		return NATION_CODE_AIA
	case "010":
		return NATION_CODE_ATA
	case "028":
		return NATION_CODE_ATG
	case "051":
		return NATION_CODE_ARM
	case "533":
		return NATION_CODE_ABW
	case "031":
		return NATION_CODE_AZE
	case "044":
		return NATION_CODE_BHS
	case "048":
		return NATION_CODE_BHR
	case "050":
		return NATION_CODE_BGD
	case "052":
		return NATION_CODE_BRB
	case "112":
		return NATION_CODE_BLR
	case "084":
		return NATION_CODE_BLZ
	case "204":
		return NATION_CODE_BEN
	case "060":
		return NATION_CODE_BMU
	case "064":
		return NATION_CODE_BTN
	case "068":
		return NATION_CODE_BOL
	case "535":
		return NATION_CODE_BES
	case "070":
		return NATION_CODE_BIH
	case "072":
		return NATION_CODE_BWA
	case "074":
		return NATION_CODE_BVT
	case "076":
		return NATION_CODE_BRA
	case "086":
		return NATION_CODE_IOT
	case "096":
		return NATION_CODE_BRN
	case "854":
		return NATION_CODE_BFA
	case "108":
		return NATION_CODE_BDI
	case "132":
		return NATION_CODE_CPV
	case "116":
		return NATION_CODE_KHM
	case "120":
		return NATION_CODE_CMR
	case "136":
		return NATION_CODE_CYM
	case "140":
		return NATION_CODE_CAF
	case "148":
		return NATION_CODE_TCD
	case "162":
		return NATION_CODE_CXR
	case "166":
		return NATION_CODE_CCK
	case "174":
		return NATION_CODE_COM
	case "178":
		return NATION_CODE_COG
	case "184":
		return NATION_CODE_COK
	case "188":
		return NATION_CODE_CRI
	case "384":
		return NATION_CODE_CIV
	case "191":
		return NATION_CODE_HRV
	case "192":
		return NATION_CODE_CUB
	case "531":
		return NATION_CODE_CUW
	case "262":
		return NATION_CODE_DJI
	case "212":
		return NATION_CODE_DMA
	case "214":
		return NATION_CODE_DOM
	case "218":
		return NATION_CODE_ECU
	case "222":
		return NATION_CODE_SLV
	case "226":
		return NATION_CODE_GNQ
	case "232":
		return NATION_CODE_ERI
	case "748":
		return NATION_CODE_SWZ
	case "238":
		return NATION_CODE_FLK
	case "234":
		return NATION_CODE_FRO
	case "242":
		return NATION_CODE_FJI
	case "254":
		return NATION_CODE_GUF
	case "258":
		return NATION_CODE_PYF
	case "260":
		return NATION_CODE_ATF
	case "266":
		return NATION_CODE_GAB
	case "270":
		return NATION_CODE_GMB
	case "268":
		return NATION_CODE_GEO
	case "288":
		return NATION_CODE_GHA
	case "292":
		return NATION_CODE_GIB
	case "304":
		return NATION_CODE_GRL
	case "308":
		return NATION_CODE_GRD
	case "312":
		return NATION_CODE_GLP
	case "316":
		return NATION_CODE_GUM
	case "320":
		return NATION_CODE_GTM
	case "831":
		return NATION_CODE_GGY
	case "324":
		return NATION_CODE_GIN
	case "624":
		return NATION_CODE_GNB
	case "328":
		return NATION_CODE_GUY
	case "332":
		return NATION_CODE_HTI
	case "334":
		return NATION_CODE_HMD
	case "336":
		return NATION_CODE_VAT
	case "340":
		return NATION_CODE_HND
	case "352":
		return NATION_CODE_ISL
	case "364":
		return NATION_CODE_IRN
	case "368":
		return NATION_CODE_IRQ
	case "833":
		return NATION_CODE_IMN
	case "388":
		return NATION_CODE_JAM
	case "832":
		return NATION_CODE_JEY
	case "400":
		return NATION_CODE_JOR
	case "398":
		return NATION_CODE_KAZ
	case "296":
		return NATION_CODE_KIR
	case "408":
		return NATION_CODE_PRK
	case "414":
		return NATION_CODE_KWT
	case "417":
		return NATION_CODE_KGZ
	case "418":
		return NATION_CODE_LAO
	case "422":
		return NATION_CODE_LBN
	case "426":
		return NATION_CODE_LSO
	case "430":
		return NATION_CODE_LBR
	case "434":
		return NATION_CODE_LBY
	case "438":
		return NATION_CODE_LIE
	case "446":
		return NATION_CODE_MAC
	case "450":
		return NATION_CODE_MDG
	case "454":
		return NATION_CODE_MWI
	case "462":
		return NATION_CODE_MDV
	case "466":
		return NATION_CODE_MLI
	case "584":
		return NATION_CODE_MHL
	case "474":
		return NATION_CODE_MTQ
	case "478":
		return NATION_CODE_MRT
	case "480":
		return NATION_CODE_MUS
	case "175":
		return NATION_CODE_MYT
	case "583":
		return NATION_CODE_FSM
	case "498":
		return NATION_CODE_MDA
	case "492":
		return NATION_CODE_MCO
	case "496":
		return NATION_CODE_MNG
	case "499":
		return NATION_CODE_MNE
	case "500":
		return NATION_CODE_MSR
	case "508":
		return NATION_CODE_MOZ
	case "104":
		return NATION_CODE_MMR
	case "516":
		return NATION_CODE_NAM
	case "520":
		return NATION_CODE_NRU
	case "524":
		return NATION_CODE_NPL
	case "540":
		return NATION_CODE_NCL
	case "558":
		return NATION_CODE_NIC
	case "562":
		return NATION_CODE_NER
	case "570":
		return NATION_CODE_NIU
	case "574":
		return NATION_CODE_NFK
	case "807":
		return NATION_CODE_MKD
	case "580":
		return NATION_CODE_MNP
	case "512":
		return NATION_CODE_OMN
	case "585":
		return NATION_CODE_PLW
	case "275":
		return NATION_CODE_PSE
	case "591":
		return NATION_CODE_PAN
	case "598":
		return NATION_CODE_PNG
	case "600":
		return NATION_CODE_PRY
	case "612":
		return NATION_CODE_PCN
	case "630":
		return NATION_CODE_PRI
	case "634":
		return NATION_CODE_QAT
	case "638":
		return NATION_CODE_REU
	case "646":
		return NATION_CODE_RWA
	case "652":
		return NATION_CODE_BLM
	case "654":
		return NATION_CODE_SHN
	case "659":
		return NATION_CODE_KNA
	case "662":
		return NATION_CODE_LCA
	case "663":
		return NATION_CODE_MAF
	case "666":
		return NATION_CODE_SPM
	case "670":
		return NATION_CODE_VCT
	case "882":
		return NATION_CODE_WSM
	case "674":
		return NATION_CODE_SMR
	case "678":
		return NATION_CODE_STP
	case "686":
		return NATION_CODE_SEN
	case "688":
		return NATION_CODE_SRB
	case "690":
		return NATION_CODE_SYC
	case "694":
		return NATION_CODE_SLE
	case "534":
		return NATION_CODE_SXM
	case "090":
		return NATION_CODE_SLB
	case "706":
		return NATION_CODE_SOM
	case "239":
		return NATION_CODE_SGS
	case "728":
		return NATION_CODE_SSD
	case "724":
		return NATION_CODE_ESP
	case "144":
		return NATION_CODE_LKA
	case "729":
		return NATION_CODE_SDN
	case "740":
		return NATION_CODE_SUR
	case "744":
		return NATION_CODE_SJM
	case "760":
		return NATION_CODE_SYR
	case "762":
		return NATION_CODE_TJK
	case "626":
		return NATION_CODE_TLS
	case "768":
		return NATION_CODE_TGO
	case "772":
		return NATION_CODE_TKL
	case "776":
		return NATION_CODE_TON
	case "780":
		return NATION_CODE_TTO
	case "788":
		return NATION_CODE_TUN
	case "792":
		return NATION_CODE_TUR
	case "795":
		return NATION_CODE_TKM
	case "796":
		return NATION_CODE_TCA
	case "798":
		return NATION_CODE_TUV
	case "800":
		return NATION_CODE_UGA
	case "581":
		return NATION_CODE_UMI
	case "858":
		return NATION_CODE_URY
	case "548":
		return NATION_CODE_VUT
	case "862":
		return NATION_CODE_VEN
	case "092":
		return NATION_CODE_VGB
	case "850":
		return NATION_CODE_VIR
	case "876":
		return NATION_CODE_WLF
	case "732":
		return NATION_CODE_ESH
	case "887":
		return NATION_CODE_YEM
	case "894":
		return NATION_CODE_ZMB
	case "716":
		return NATION_CODE_ZWE
	default:
		return ""
	}
}

func (nc NationCode) ToTwoLetter() string {
	var converter = map[NationCode]string{
		NATION_CODE_USA: "US",
		NATION_CODE_JPN: "JP",
		NATION_CODE_AUT: "AT",
		NATION_CODE_BEL: "BE",
		NATION_CODE_BGR: "BG",
		NATION_CODE_CYP: "CY",
		NATION_CODE_CZE: "CZ",
		NATION_CODE_DEU: "DE",
		NATION_CODE_DNK: "DK",
		NATION_CODE_ESP: "ES",
		NATION_CODE_EST: "EE",
		NATION_CODE_FIN: "FI",
		NATION_CODE_FRA: "FR",
		NATION_CODE_GRC: "GR",
		NATION_CODE_HUN: "HU",
		NATION_CODE_IRL: "IE",
		NATION_CODE_ITA: "IT",
		NATION_CODE_LAT: "LV",
		NATION_CODE_LIT: "LT",
		NATION_CODE_LUX: "LU",
		NATION_CODE_MLT: "MT",
		NATION_CODE_NLD: "NL",
		NATION_CODE_POL: "PL",
		NATION_CODE_PRT: "PT",
		NATION_CODE_ROU: "RO",
		NATION_CODE_SVK: "SK",
		NATION_CODE_SVN: "SI",
		NATION_CODE_SWE: "SE",
		NATION_CODE_GBR: "GB",
		NATION_CODE_CHN: "CN",
		NATION_CODE_CAN: "CA",
		NATION_CODE_AUS: "AU",
		NATION_CODE_CHE: "CH",
		NATION_CODE_HKG: "HK",
		NATION_CODE_NZL: "NZ",
		NATION_CODE_MEX: "MX",
		NATION_CODE_SGP: "SG",
		NATION_CODE_NOR: "NO",
		NATION_CODE_KOR: "KR",
		NATION_CODE_TUR: "TR",
		NATION_CODE_IND: "IN",
		NATION_CODE_BRA: "BR",
		NATION_CODE_ZAF: "ZA",
		NATION_CODE_RUS: "RU",
		NATION_CODE_TWN: "TW",
		NATION_CODE_ISR: "IL",
		NATION_CODE_ARE: "AE",
		NATION_CODE_ARG: "AR",
		NATION_CODE_MYS: "MY",
		NATION_CODE_THA: "TH",
		NATION_CODE_SAU: "SA",
		NATION_CODE_IDN: "ID",
		NATION_CODE_PAK: "PK",
		NATION_CODE_NGA: "NG",
		NATION_CODE_EGY: "EG",
		NATION_CODE_VNM: "VN",
		NATION_CODE_PHL: "PH",
		NATION_CODE_COL: "CO",
		NATION_CODE_CHL: "CL",
		NATION_CODE_PER: "PE",
		NATION_CODE_DZA: "DZ",
		NATION_CODE_MAR: "MA",
		NATION_CODE_KEN: "KE",
		NATION_CODE_UKR: "UA",
		NATION_CODE_AFG: "AF",
		NATION_CODE_ETH: "ET",
		NATION_CODE_COD: "CD",
		NATION_CODE_TZA: "TZ",
		NATION_CODE_UZB: "UZ",
		NATION_CODE_ALA: "AX", // Åland Islands
		NATION_CODE_ALB: "AL", // Albania
		NATION_CODE_ASM: "AS", // American Samoa
		NATION_CODE_AND: "AD", // Andorra
		NATION_CODE_AGO: "AO", // Angola
		NATION_CODE_AIA: "AI", // Anguilla
		NATION_CODE_ATA: "AQ", // Antarctica
		NATION_CODE_ATG: "AG", // Antigua and Barbuda
		NATION_CODE_ARM: "AM", // Armenia
		NATION_CODE_ABW: "AW", // Aruba
		NATION_CODE_AZE: "AZ", // Azerbaijan
		NATION_CODE_BHS: "BS", // Bahamas
		NATION_CODE_BHR: "BH", // Bahrain
		NATION_CODE_BGD: "BD", // Bangladesh
		NATION_CODE_BRB: "BB", // Barbados
		NATION_CODE_BLR: "BY", // Belarus
		NATION_CODE_BLZ: "BZ", // Belize
		NATION_CODE_BEN: "BJ", // Benin
		NATION_CODE_BMU: "BM", // Bermuda
		NATION_CODE_BTN: "BT", // Bhutan
		NATION_CODE_BOL: "BO", // Bolivia (Plurinational State of)
		NATION_CODE_BES: "BQ", // Bonaire, Sint Eustatius and Saba
		NATION_CODE_BIH: "BA", // Bosnia and Herzegovina
		NATION_CODE_BWA: "BW", // Botswana
		NATION_CODE_BVT: "BV", // Bouvet Island
		NATION_CODE_IOT: "IO", // British Indian Ocean Territory
		NATION_CODE_BRN: "BN", // Brunei Darussalam
		NATION_CODE_BFA: "BF", // Burkina Faso
		NATION_CODE_BDI: "BI", // Burundi
		NATION_CODE_CPV: "CV", // Cabo Verde
		NATION_CODE_KHM: "KH", // Cambodia
		NATION_CODE_CMR: "CM", // Cameroon
		NATION_CODE_CYM: "KY", // Cayman Islands
		NATION_CODE_CAF: "CF", // Central African Republic
		NATION_CODE_TCD: "TD", // Chad
		NATION_CODE_CXR: "CX", // Christmas Island
		NATION_CODE_CCK: "CC", // Cocos (Keeling) Islands
		NATION_CODE_COM: "KM", // Comoros
		NATION_CODE_COG: "CG", // Congo
		NATION_CODE_COK: "CK", // Cook Islands
		NATION_CODE_CRI: "CR", // Costa Rica
		NATION_CODE_CIV: "CI", // Côte d'Ivoire
		NATION_CODE_HRV: "HR", // Croatia
		NATION_CODE_CUB: "CU", // Cuba
		NATION_CODE_CUW: "CW", // Curaçao
		NATION_CODE_DJI: "DJ", // Djibouti
		NATION_CODE_DMA: "DM", // Dominica
		NATION_CODE_DOM: "DO", // Dominican Republic
		NATION_CODE_ECU: "EC", // Ecuador
		NATION_CODE_SLV: "SV", // El Salvador
		NATION_CODE_GNQ: "GQ", // Equatorial Guinea
		NATION_CODE_ERI: "ER", // Eritrea
		NATION_CODE_SWZ: "SZ", // Eswatini
		NATION_CODE_FLK: "FK", // Falkland Islands (Malvinas)
		NATION_CODE_FRO: "FO", // Faroe Islands
		NATION_CODE_FJI: "FJ", // Fiji
		NATION_CODE_GUF: "GF", // French Guiana
		NATION_CODE_PYF: "PF", // French Polynesia
		NATION_CODE_ATF: "TF", // French Southern Territories
		NATION_CODE_GAB: "GA", // Gabon
		NATION_CODE_GMB: "GM", // Gambia
		NATION_CODE_GEO: "GE", // Georgia
		NATION_CODE_GHA: "GH", // Ghana
		NATION_CODE_GIB: "GI", // Gibraltar
		NATION_CODE_GRL: "GL", // Greenland
		NATION_CODE_GRD: "GD", // Grenada
		NATION_CODE_GLP: "GP", // Guadeloupe
		NATION_CODE_GUM: "GU", // Guam
		NATION_CODE_GTM: "GT", // Guatemala
		NATION_CODE_GGY: "GG", // Guernsey
		NATION_CODE_GIN: "GN", // Guinea
		NATION_CODE_GNB: "GW", // Guinea-Bissau
		NATION_CODE_GUY: "GY", // Guyana
		NATION_CODE_HTI: "HT", // Haiti
		NATION_CODE_HMD: "HM", // Heard Island and McDonald Islands
		NATION_CODE_VAT: "VA", // Holy See
		NATION_CODE_HND: "HN", // Honduras
		NATION_CODE_ISL: "IS", // Iceland
		NATION_CODE_IRN: "IR", // Iran (Islamic Republic of)
		NATION_CODE_IRQ: "IQ", // Iraq
		NATION_CODE_IMN: "IM", // Isle of Man
		NATION_CODE_JAM: "JM", // Jamaica
		NATION_CODE_JEY: "JE", // Jersey
		NATION_CODE_JOR: "JO", // Jordan
		NATION_CODE_KAZ: "KZ", // Kazakhstan
		NATION_CODE_KIR: "KI", // Kiribati
		NATION_CODE_PRK: "KP", // Korea (Democratic People's Republic of)
		NATION_CODE_KWT: "KW", // Kuwait
		NATION_CODE_KGZ: "KG", // Kyrgyzstan
		NATION_CODE_LAO: "LA", // Lao People's Democratic Republic
		NATION_CODE_LBN: "LB", // Lebanon
		NATION_CODE_LSO: "LS", // Lesotho
		NATION_CODE_LBR: "LR", // Liberia
		NATION_CODE_LBY: "LY", // Libya
		NATION_CODE_LIE: "LI", // Liechtenstein
		NATION_CODE_MAC: "MO", // Macao
		NATION_CODE_MDG: "MG", // Madagascar
		NATION_CODE_MWI: "MW", // Malawi
		NATION_CODE_MDV: "MV", // Maldives
		NATION_CODE_MLI: "ML", // Mali
		NATION_CODE_MHL: "MH", // Marshall Islands
		NATION_CODE_MTQ: "MQ", // Martinique
		NATION_CODE_MRT: "MR", // Mauritania
		NATION_CODE_MUS: "MU", // Mauritius
		NATION_CODE_MYT: "YT", // Mayotte
		NATION_CODE_FSM: "FM", // Micronesia (Federated States of)
		NATION_CODE_MDA: "MD", // Moldova, Republic of
		NATION_CODE_MCO: "MC", // Monaco
		NATION_CODE_MNG: "MN", // Mongolia
		NATION_CODE_MNE: "ME", // Montenegro
		NATION_CODE_MSR: "MS", // Montserrat
		NATION_CODE_MOZ: "MZ", // Mozambique
		NATION_CODE_MMR: "MM", // Myanmar
		NATION_CODE_NAM: "NA", // Namibia
		NATION_CODE_NRU: "NR", // Nauru
		NATION_CODE_NPL: "NP", // Nepal
		NATION_CODE_NCL: "NC", // New Caledonia
		NATION_CODE_NIC: "NI", // Nicaragua
		NATION_CODE_NER: "NE", // Niger
		NATION_CODE_NIU: "NU", // Niue
		NATION_CODE_NFK: "NF", // Norfolk Island
		NATION_CODE_MKD: "MK", // North Macedonia
		NATION_CODE_MNP: "MP", // Northern Mariana Islands
		NATION_CODE_OMN: "OM", // Oman
		NATION_CODE_PLW: "PW", // Palau
		NATION_CODE_PSE: "PS", // Palestine, State of
		NATION_CODE_PAN: "PA", // Panama
		NATION_CODE_PNG: "PG", // Papua New Guinea
		NATION_CODE_PRY: "PY", // Paraguay
		NATION_CODE_PCN: "PN", // Pitcairn
		NATION_CODE_PRI: "PR", // Puerto Rico
		NATION_CODE_QAT: "QA", // Qatar
		NATION_CODE_REU: "RE", // Réunion
		NATION_CODE_RWA: "RW", // Rwanda
		NATION_CODE_BLM: "BL", // Saint Barthélemy
		NATION_CODE_SHN: "SH", // Saint Helena, Ascension and Tristan da Cunha
		NATION_CODE_KNA: "KN", // Saint Kitts and Nevis
		NATION_CODE_LCA: "LC", // Saint Lucia
		NATION_CODE_MAF: "MF", // Saint Martin (French part)
		NATION_CODE_SPM: "PM", // Saint Pierre and Miquelon
		NATION_CODE_VCT: "VC", // Saint Vincent and the Grenadines
		NATION_CODE_WSM: "WS", // Samoa
		NATION_CODE_SMR: "SM", // San Marino
		NATION_CODE_STP: "ST", // Sao Tome and Principe
		NATION_CODE_SEN: "SN", // Senegal
		NATION_CODE_SRB: "RS", // Serbia
		NATION_CODE_SYC: "SC", // Seychelles
		NATION_CODE_SLE: "SL", // Sierra Leone
		NATION_CODE_SXM: "SX", // Sint Maarten (Dutch part)
		NATION_CODE_SLB: "SB", // Solomon Islands
		NATION_CODE_SOM: "SO", // Somalia
		NATION_CODE_SGS: "GS", // South Georgia and the South Sandwich Islands
		NATION_CODE_SSD: "SS", // South Sudan
		NATION_CODE_LKA: "LK", // Sri Lanka
		NATION_CODE_SDN: "SD", // Sudan
		NATION_CODE_SUR: "SR", // Suriname
		NATION_CODE_SJM: "SJ", // Svalbard and Jan Mayen
		NATION_CODE_SYR: "SY", // Syrian Arab Republic
		NATION_CODE_TJK: "TJ", // Tajikistan
		NATION_CODE_TLS: "TL", // Timor-Leste
		NATION_CODE_TGO: "TG", // Togo
		NATION_CODE_TKL: "TK", // Tokelau
		NATION_CODE_TON: "TO", // Tonga
		NATION_CODE_TTO: "TT", // Trinidad and Tobago
		NATION_CODE_TUN: "TN", // Tunisia
		NATION_CODE_TKM: "TM", // Turkmenistan
		NATION_CODE_TCA: "TC", // Turks and Caicos Islands
		NATION_CODE_TUV: "TV", // Tuvalu
		NATION_CODE_UGA: "UG", // Uganda
		NATION_CODE_UMI: "UM", // United States Minor Outlying Islands
		NATION_CODE_URY: "UY", // Uruguay
		NATION_CODE_VUT: "VU", // Vanuatu
		NATION_CODE_VEN: "VE", // Venezuela (Bolivarian Republic of)
		NATION_CODE_VGB: "VG", // Virgin Islands (British)
		NATION_CODE_VIR: "VI", // Virgin Islands (U.S.)
		NATION_CODE_WLF: "WF", // Wallis and Futuna
		NATION_CODE_ESH: "EH", // Western Sahara
		NATION_CODE_YEM: "YE", // Yemen
		NATION_CODE_ZMB: "ZM", // Zambia
		NATION_CODE_ZWE: "ZW", // Zimbabwe
	}

	return converter[nc]
}

func (nt NotifyType) String() string {
	switch nt {
	case NOTIFY_TYPE_EMAIL:
		return "email"
	case NOTIFY_TYPE_SMS:
		return "sms"
	case NOTIFY_TYPE_APP:
		return "app"
	case NOTIFY_TYPE_MESSAGE:
		return "message"
	}
	return ""
}

func (od OrderDirection) String() string {
	switch od {
	case ORDER_DIRECTION_ASC:
		return "asc"
	case ORDER_DIRECTION_DESC:
		return "desc"
	}
	return ""
}

func (pp PreviewPurpose) String() string {
	switch pp {
	case PREVIEW_PURPOSE_APPLY:
		return "apply"
	case PREVIEW_PURPOSE_TRANSFER:
		return "transfer"
	case PREVIEW_PURPOSE_EXCHANGE:
		return "exchange"
	case PREVIEW_PURPOSE_WITHDRAW:
		return "withdraw"
	case PREVIEW_PURPOSE_TOP_UP:
		return "top_up"
	case PREVIEW_PURPOSE_TOP_DOWN:
		return "top_down"
	case PREVIEW_PURPOSE_CARD_TO_CARD:
		return "card_to_card"
	case PREVIEW_PURPOSE_UNBLOCK:
		return "unblock"

	}
	return ""
}

func (pt PromotionType) String() string {
	switch pt {
	case PROMOTIOM_TYPE_APPLY:
		return "apply"
	case PROMOTIOM_TYPE_EXCHANGE:
		return "exchange"
	case PROMOTIOM_TYPE_TRANSFER:
		return "transfer"
	case PROMOTIOM_TYPE_WITHDRAWAL:
		return "withdrawal"
	}
	return ""
}

func (p Protocol) String() string {
	switch p {
	case PROTOCOL_OMNI:
		return "omni"
	case PROTOCOL_ERC20:
		return "erc-20"
	case PROTOCOL_TRC20:
		return "trc-20"
	case PROTOCOL_DEP20:
		return "dep-20"
	case PROTOCOL_SPL:
		return "spl"
	case PROTOCOL_BEP20:
		return "bep-20"
	case PROTOCOL_MATIC_ERC20:
		return "matic-erc-20"
	}
	return ""
}
func (Protocol) FromString(s string) Protocol {
	sLower := strings.ToLower(s)
	switch sLower {
	case "omni":
		return PROTOCOL_OMNI
	case "erc-20":
		return PROTOCOL_ERC20
	case "trc-20":
		return PROTOCOL_TRC20
	case "dep-20":
		return PROTOCOL_DEP20
	case "spl":
		return PROTOCOL_SPL
	case "bep-20":
		return PROTOCOL_BEP20
	case "matic-erc-20":
		return PROTOCOL_MATIC_ERC20
	}
	return 0
}

func (p Protocol) IsValid() bool {
	switch p {
	case
		PROTOCOL_OMNI,
		PROTOCOL_ERC20,
		PROTOCOL_TRC20,
		PROTOCOL_SPL,
		PROTOCOL_BEP20,
		PROTOCOL_MATIC_ERC20:
		return true
	}
	return false
}

func (r RatePurpose) String() string {
	switch r {
	case RATE_PURPOSE_APPLY:
		return "apply"
	case RATE_PURPOSE_EXCHANGE:
		return "exchange"
	case RATE_PURPOSE_TRANSFER:
		return "transfer"
	case RATE_PURPOSE_CARD_TO_CARD:
		return "card_to_card"
	case RATE_PURPOSE_TOP_UP:
		return "top_up"
	case RATE_PURPOSE_TOP_DOWN:
		return "top_down"
	case RATE_PURPOSE_FEE:
		return "fee"
	case RATE_PURPOSE_LIMIT:
		return "limit"
	case RATE_PURPOSE_MERCHANT_EXCHANGE:
		return "merchant_exchange"
	case RATE_PURPOSE_MERCHANT_TRANSFER:
		return "merchant_transfer"
	case RATE_PURPOSE_DISCOUNT:
		return "discount"
	case RATE_PURPOSE_DISPLAY:
		return "display"
	case RATE_PURPOSE_CHIPPAY_EXPRESS_BUY:
		return "chippay_express_buy"
	}
	return ""
}

func (rc ReapChannel) String() string {
	switch rc {
	case REAP_CHANNEL_ATM:
		return "atm"
	case REAP_CHANNEL_POS:
		return "pos"
	case REAP_CHANNEL_ECOMMERCE:
		return "ecommerce"
	case REAP_CHANNEL_VISA_DIRECT:
		return "visa_direct"
	}
	return ""
}

func (rc ReapChannel) FromString(s string) ReapChannel {
	sLower := strings.ToLower(s)
	switch sLower {
	case "atm":
		return REAP_CHANNEL_ATM
	case "pos":
		return REAP_CHANNEL_POS
	case "ecommerce":
		return REAP_CHANNEL_ECOMMERCE
	case "visa_direct":
		return REAP_CHANNEL_VISA_DIRECT
	}
	return 0
}

func (rdl ReapDataLoss) String() string {
	switch rdl {
	case REAP_DATA_LOSS_OK:
		return "ok"
	case REAP_DATA_LOSS_DATA_LOSS:
		return "data_loss"
	case REAP_DATA_LOSS_NOT_FREEZED:
		return "not_freezed"
	case REAP_DATA_LOSS_RECOVERED:
		return "recovered"

	}
	return ""
}

func (rds DeliveryStatus) String() string {
	switch rds {
	case DELIVERY_STATUS_CARD_NOT_ORDERED:
		return "CARD_NOT_ORDERED"
	case DELIVERY_STATUS_CARD_IN_PRODUCTION:
		return "CARD_IN_PRODUCTION"
	case DELIVERY_STATUS_CARD_SENT_TO_USER:
		return "CARD_SENT_TO_USER"
	case DELIVERY_STATUS_CARD_ACTIVATED:
		return "CARD_ACTIVATED"
	case DELIVERY_STATUS_NOT_PHYSICAL_CARD:
		return "NOT_PHYSICAL_CARD"
	}
	return ""
}

func (rds DeliveryStatus) FromString(s string) DeliveryStatus {
	switch strings.ToLower(s) {
	case "card_not_ordered":
		return DELIVERY_STATUS_CARD_NOT_ORDERED
	case "card_in_production":
		return DELIVERY_STATUS_CARD_IN_PRODUCTION
	case "card_sent_to_user":
		return DELIVERY_STATUS_CARD_SENT_TO_USER
	case "card_activated":
		return DELIVERY_STATUS_CARD_ACTIVATED
	case "not_physical_card":
		return DELIVERY_STATUS_NOT_PHYSICAL_CARD
	}
	return 0
}

func (rpt RewardProductType) String() string {
	switch rpt {
	case REWARD_PRODUCT_TYPE_VOUCHER:
		return "voucher"
	case REWARD_PRODUCT_TYPE_PHYSICAL:
		return "physical"
	case REWARD_PRODUCT_TYPE_GAME:
		return "game"
	case REWARD_PRODUCT_TYPE_WAIVER:
		return "waiver"
	}
	return ""
}

func (r RewardProductRestriction) String() string {
	switch r {
	case REWARD_PRODUCT_RESTRICTION_UNRESTRICTED:
		return "unrestricted"
	case REWARD_PRODUCT_RESTRICTION_PARTIALLY_ALLOWED:
		return "partially_allowed"
	case REWARD_PRODUCT_RESTRICTION_PARTIALLY_RESTRICTED:
		return "partially_restricted"
	}
	return ""
}

func (r RewardProductStatus) String() string {
	switch r {
	case REWARD_PRODUCT_STATUS_ACTIVE:
		return "active"
	case REWARD_PRODUCT_STATUS_INACTIVE:
		return "inactive"
	}
	return ""
}

func (r RewardOrderStatus) String() string {
	switch r {
	case REWARD_ORDER_STATUS_PENDING:
		return "pending"
	case REWARD_ORDER_STATUS_PROCESSING:
		return "processing"
	case REWARD_ORDER_STATUS_SUCCESS:
		return "success"
	case REWARD_ORDER_STATUS_FAILED:
		return "failed"
	}
	return ""
}

func (r Role) FromString(s string) Role {
	switch s {
	case "guest":
		return ROLE_GUEST
	case "user":
		return ROLE_USER
	case "admin":
		return ROLE_ADMIN
	case "merchant":
		return ROLE_MERCHANT
	case "merchant_admin":
		return ROLE_MERCHANT_ADMIN
	case "merchant_user":
		return ROLE_MERCHANT_USER
	case "admin_agent_user":
		return ROLE_ADMIN_AGENT_USER
	case "system":
		return ROLE_SYSTEM
	}
	return 0
}

func (r Role) String() string {
	switch r {
	case ROLE_GUEST:
		return "guest"
	case ROLE_USER:
		return "user"
	case ROLE_ADMIN:
		return "admin"
	case ROLE_MERCHANT:
		return "merchant"
	case ROLE_MERCHANT_ADMIN:
		return "merchant_admin"
	case ROLE_MERCHANT_USER:
		return "merchant_user"
	case ROLE_ADMIN_AGENT_USER:
		return "agent_admin_user"
	case ROLE_SYSTEM:
		return "system"
	}
	return ""
}

func (lp LockPurpose) String() string {
	switch lp {
	case LOCK_PURPOSE_APPLY_CONFIRM:
		return "apply_confirm"
	case LOCK_PURPOSE_REAP_TRANSACTION:
		return "reap_transaction"
	case LOCK_PURPOSE_EXCHANGE_CONFIRM:
		return "exchange_confirm"
	case LOCK_PURPOSE_TRANSFER_CONFIRM:
		return "transfer_confirm"
	case LOCK_PURPOSE_TOP_UP_CONFIRM:
		return "top_up_confirm"
	case LOCK_PURPOSE_TOP_DOWN_CONFIRM:
		return "top_down_confirm"
	case LOCK_PURPOSE_THREEDS_DAO:
		return "threeds_dao"
	case LOCK_PURPOSE_CARD_TO_CARD_CONFIRM:
		return "card_to_card_confirm"
	case LOCK_PURPOSE_WITHDRAW_CONFIRM:
		return "withdraw_confirm"
	case LOCK_PURPOSE_MERCHANT_APPLY:
		return "merchant_apply"
	case LOCK_PURPOSE_CALLBACK:
		return "callback"
	case LOCK_PURPOSE_MANUAL:
		return "manual"
	case LOCK_PURPOSE_MERCHANT_THREEDS_DAO:
		return "merchant_threeds_dao"
	case LOCK_PURPOSE_CREATE_POINT:
		return "create_point"
	case LOCK_PURPOSE_CREATE_WALLET:
		return "create_wallet"
	case LOCK_PURPOSE_TASK:
		return "task"
	case LOCK_PURPOSE_CONSUME_POINT:
		return "consume_point"
	case LOCK_PURPOSE_CREATE_AUTO_YIELD:
		return "create_auto_yield"
	case LOCK_PURPOSE_DISTRIBUTE_AUTO_YIELD:
		return "distribute_auto_yield"
	case LOCK_PURPOSE_GET_TEMP_ADDRESS:
		return "get_temp_address"
	case LOCK_PURPOSE_WHALE_WEBHOOK_BALANCE_EDIT:
		return "whale_webhook_balance_edit"
	case LOCK_PURPOSE_WHALE_WEBHOOK_TRANSFER:
		return "whale_webhook_transfer"
	case LOCK_PURPOSE_WHALE_WEBHOOK_CARD_TRANSACTION:
		return "whale_webhook_card_transaction"
	case LOCK_PURPOSE_WHALE_THREEDS_DAO:
		return "whale_threeds_dao"
	case LOCK_PURPOSE_CARD_STATUS_FROZEN:
		return "card_status_frozen"
	case LOCK_PURPOSE_WHALE_STATUS_DELETED:
		return "whale_status_deleted"
	case LOCK_PURPOSE_WHALE_CHECK_ADJUST:
		return "whale_check_adjust"
	case LOCK_PURPOSE_PAYCRYPTO_KYC_SUBMIT:
		return "paycrypto_kyc_submit"
	case LOCK_PURPOSE_PAYCRYPTO_ACTIVATE:
		return "paycrypto_activate"
	case LOCK_PURPOSE_PAYCRYPTO_WEBHOOK:
		return "paycrypto_webhook"
	case LOCK_PURPOSE_ADMIN_APPLY_CONFIRM:
		return "admin_apply_confirm"
	case LOCK_PURPOSE_ADMIN_DELETE_USER:
		return "admin_delete_user"
	case LOCK_PURPOSE_ADMIN_DELETE_CARD:
		return "admin_delete_card"
	case LOCK_PURPOSE_ETHEFI_SYNC_TRANSACTION:
		return "etherfi_sync_transaction"
	case LOCK_PURPOSE_ETHEFI_SYNC_CARD_TRANSACTION:
		return "etherfi_sync_card_transaction"
	case LOCK_PURPOSE_ETHEFI_SYNC_CARD_EVENT:
		return "etherfi_sync_card_event"
	case LOCK_PURPOSE_PIN_UNLOCK:
		return "pin_unlock"
	case LOCK_PURPOSE_ETHEFI_SYNC_SPENDING_LIMIT:
		return "etherfi_sync_pending_limit"
	case LOCK_PURPOSE_PHYSICAL_CARD_APPLY:
		return "physical_card_apply"
	case LOCK_PURPOSE_FINANCE_REPORT:
		return "finance_report"
	case LOCK_PURPOSE_ETHERFI_APPLY_CONFIRM:
		return "etherfi_apply_confirm"
	case LOCK_PURPOSE_ETHERFI_REPORT_DOWNLOAD:
		return "etherfi_report_download"
	}
	return ""
}

func (s ShippingStatus) String() string {
	switch s {
	case SHIPPING_STATUS_NOT_REQUESTED:
		return "not_requested"
	case SHIPPING_STATUS_IN_PROGRESS:
		return "in_progress"
	case SHIPPING_STATUS_IMCOMPLETE:
		return "incomplete" // 修正拼寫
	case SHIPPING_STATUS_PENDING_APPROVAL:
		return "pending_approval"
	case SHIPPING_STATUS_PENDING_FULFILMENT:
		return "pending_fulfilment"
	case SHIPPING_STATUS_FULFILLED:
		return "fulfilled"
	case SHIPPING_STATUS_CANCELLED:
		return "cancelled"
	case SHIPPING_STATUS_ORDER_IN_PRODUCTION:
		return "order_in_production"
	case SHIPPING_STATUS_REJECTED:
		return "rejected"
	}
	return ""
}

func (ts ThreedsStatus) String() string {
	switch ts {
	case THREEDS_STATUS_PENDING:
		return "pending"
	case THREEDS_STATUS_AUTHORIZING:
		return "authorizing"
	case THREEDS_STATUS_AUTHORIZED:
		return "authorized"
	case THREEDS_STATUS_DECLINING:
		return "declining"
	case THREEDS_STATUS_DECLINED:
		return "declined"
	case THREEDS_STATUS_TIMEOUT:
		return "timeout"

	}
	return ""
}

func (trt TransactionRecordType) String() string {
	switch trt {
	case TRANSACTION_RECORD_TYPE_APPLY:
		return "apply"
	case TRANSACTION_RECORD_TYPE_TRANSFER:
		return "transfer"
	case TRANSACTION_RECORD_TYPE_EXCHANGE:
		return "exchange"
	case TRANSACTION_RECORD_TYPE_WITHDRAW:
		return "withdraw"
	case TRANSACTION_RECORD_TYPE_PAY:
		return "pay"
	case TRANSACTION_RECORD_TYPE_DEPOSIT:
		return "deposit"
	case TRANSACTION_RECORD_TYPE_TOP_UP:
		return "top_up"
	case TRANSACTION_RECORD_TYPE_TOP_DOWN:
		return "top_down"
	case TRANSACTION_RECORD_TYPE_REFUND:
		return "refund"
	case TRANSACTION_RECORD_TYPE_VERIFY:
		return "verify"
	case TRANSACTION_RECORD_TYPE_CARD_TO_CARD:
		return "card_to_card"
	case TRANSACTION_RECORD_TYPE_MANUAL:
		return "manual"
	case TRANSACTION_RECORD_TYPE_DELETE_CARD:
		return "delete_card"
	case TRANSACTION_RECORD_TYPE_MERCHANT_AUTO_EXCHANGE:
		return "merchant_auto_exchange"
	case TRANSACTION_RECORD_TYPE_CHIPPAY_EXPRESS:
		return "chippay_express"
	case TRANSACTION_RECORD_TYPE_MERCHANT_ADJUST_BALANCE:
		return "merchant_adjust_balance"
	case TRANSACTION_RECORD_TYPE_POINT_ACCURAL:
		return "point_accural"
	case TRANSACTION_RECORD_TYPE_POINT_REDEEM:
		return "point_redeem"
	case TRANSACTION_RECORD_TYPE_INTEREST:
		return "interest"
	case TRANSACTION_RECORD_TYPE_FINANCIAL_TRANSFER:
		return "financial_transfer"
	case TRANSACTION_RECORD_TYPE_WHALE_PAY:
		return "pay"
	case TRANSACTION_RECORD_TYPE_WHALE_REFUND:
		return "refund"
	case TRANSACTION_RECORD_TYPE_DISTRIBUTE_APPLY:
		return "distribute_apply"
	case TRANSACTION_RECORD_TYPE_PAYCRYPTO_PAY:
		return "pay"
	case TRANSACTION_RECORD_TYPE_PAYCRYPTO_REFUND:
		return "refund"
	case TRANSACTION_RECORD_TYPE_PAYCRYPTO_VERIFY:
		return "verify"
	case TRANSACTION_RECORD_TYPE_BINANCE_PAY:
		return "binance_pay"
	case TRANSACTION_RECORD_TYPE_SYSTEM_CHARGE:
		return "system_charge"
	case TRANSACTION_RECORD_TYPE_ETHERFI:
		return "pay"
	case TRANSACTION_RECORD_TYPE_ETHERFI_REFUND:
		return "refund"
	}
	return ""
}

func (TransactionRecordType) FromString(s string) []TransactionRecordType {

	var result []TransactionRecordType
	switch s {
	case "apply":
		result = append(result, TRANSACTION_RECORD_TYPE_APPLY)
	case "transfer":
		result = append(result, TRANSACTION_RECORD_TYPE_TRANSFER)
	case "exchange":
		result = append(result, TRANSACTION_RECORD_TYPE_EXCHANGE)
	case "withdraw":
		result = append(result, TRANSACTION_RECORD_TYPE_WITHDRAW)
	case "pay":
		result = append(result, TRANSACTION_RECORD_TYPE_PAY, TRANSACTION_RECORD_TYPE_WHALE_PAY, TRANSACTION_RECORD_TYPE_PAYCRYPTO_PAY)
	case "deposit":
		result = append(result, TRANSACTION_RECORD_TYPE_DEPOSIT)
	case "top_up":
		result = append(result, TRANSACTION_RECORD_TYPE_TOP_UP)
	case "top_down":
		result = append(result, TRANSACTION_RECORD_TYPE_TOP_DOWN)
	case "refund":
		result = append(result, TRANSACTION_RECORD_TYPE_REFUND, TRANSACTION_RECORD_TYPE_WHALE_REFUND, TRANSACTION_RECORD_TYPE_PAYCRYPTO_REFUND)
	case "verify":
		result = append(result, TRANSACTION_RECORD_TYPE_VERIFY, TRANSACTION_RECORD_TYPE_PAYCRYPTO_VERIFY)
	case "card_to_card":
		result = append(result, TRANSACTION_RECORD_TYPE_CARD_TO_CARD)
	case "chippay_express":
		result = append(result, TRANSACTION_RECORD_TYPE_CHIPPAY_EXPRESS)
	case "point_accural":
		result = append(result, TRANSACTION_RECORD_TYPE_POINT_ACCURAL)
	case "point_redeem":
		result = append(result, TRANSACTION_RECORD_TYPE_POINT_REDEEM)
	case "interest":
		result = append(result, TRANSACTION_RECORD_TYPE_INTEREST)
	case "financial_transfer":
		result = append(result, TRANSACTION_RECORD_TYPE_FINANCIAL_TRANSFER)
	case "distribute_apply":
		result = append(result, TRANSACTION_RECORD_TYPE_DISTRIBUTE_APPLY)
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
	case TRANSACTION_STATUS_APPLY_PENDING:
		return "apply_pending"
	case TRANSACTION_STATUS_APPLY_CREATED:
		return "apply_created"
	case TRANSACTION_STATUS_APPLY_DEPOSITED:
		return "apply_deposited"
	case TRANSACTION_STATUS_APPLY_FAILED:
		return "apply_failed"

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

	case TRANSACTION_STATUS_WITHDRAW_PENDING:
		return "withdraw_pending"
	case TRANSACTION_STATUS_WITHDRAW_SUCCESS:
		return "withdraw_success"
	case TRANSACTION_STATUS_WITHDRAW_FAILED:
		return "withdraw_failed"
	case TRANSACTION_STATUS_WITHDRAW_FREEZED:
		return "withdraw_freezed"
	case TRANSACTION_STATUS_WITHDRAW_COINSDO_PENDING:
		return "withdraw_coinsdo_pending"
	case TRANSACTION_STATUS_WITHDRAW_COINSDO_PASSED:
		return "withdraw_coinsdo_passed"
	case TRANSACTION_STATUS_WITHDRAW_COINSDO_REJECTED:
		return "withdraw_coinsdo_rejected"
	case TRANSACTION_STATUS_WITHDRAW_COINSDO_CANCELLED:
		return "withdraw_coinsdo_cancelled"
	case TRANSACTION_STATUS_WITHDRAW_COINSDO_FAILED:
		return "withdraw_coinsdo_failed"

	case TRANSACTION_STATUS_PAY_REQUESTING:
		return "pay_requesting"
	case TRANSACTION_STATUS_PAY_AUTHORIZING:
		return "pay_authorizing"
	case TRANSACTION_STATUS_PAY_PENDING:
		return "pay_pending"
	case TRANSACTION_STATUS_PAY_REVERSAL:
		return "pay_reversal"
	case TRANSACTION_STATUS_PAY_CLEARED:
		return "pay_cleared"
	case TRANSACTION_STATUS_PAY_DECLINED:
		return "pay_declined"
	case TRANSACTION_STATUS_PAY_VOID:
		return "pay_void"
	case TRANSACTION_STATUS_PAY_REFUNDED:
		return "pay_refunded"

	case TRANSACTION_STATUS_DEPOSIT_PENDING:
		return "deposit_pending"
	case TRANSACTION_STATUS_DEPOSIT_CONFIRMED:
		return "deposit_confirmed"
	case TRANSACTION_STATUS_DEPOSIT_DEPOSITED:
		return "deposit_deposited"
	case TRANSACTION_STATUS_DEPOSIT_FAILED:
		return "deposit_failed"

	case TRANSACTION_STATUS_TOP_UP_PENDING:
		return "top_up_pending"
	case TRANSACTION_STATUS_TOP_UP_PROCESSING:
		return "top_up_pending"
	case TRANSACTION_STATUS_TOP_UP_SUCCESS:
		return "top_up_success"
	case TRANSACTION_STATUS_TOP_UP_FAILED:
		return "top_up_failed"

	case TRANSACTION_STATUS_TOP_DOWN_PENDING:
		return "top_down_pending"
	case TRANSACTION_STATUS_TOP_DOWN_PROCESSING:
		return "top_down_pending"
	case TRANSACTION_STATUS_TOP_DOWN_SUCCESS:
		return "top_down_success"
	case TRANSACTION_STATUS_TOP_DOWN_FAILED:
		return "top_down_failed"

	case TRANSACTION_STATUS_REFUND_PENDING:
		return "refund_pending"
	case TRANSACTION_STATUS_REFUND_SUCCESS:
		return "refund_success"
	case TRANSACTION_STATUS_REFUND_FAILED:
		return "refund_failed"

	case TRANSACTION_STATUS_VERIFY_PENDING:
		return "verify_pending"
	case TRANSACTION_STATUS_VERIFY_SUCCESS:
		return "verify_success"
	case TRANSACTION_STATUS_VERIFY_FAILED:
		return "verify_failed"

	case TRANSACTION_STATUS_CARD_TO_CARD_PENDING:
		return "card_to_card_pending"
	case TRANSACTION_STATUS_CARD_TO_CARD_PROCESSING:
		return "card_to_card_processing"
	case TRANSACTION_STATUS_CARD_TO_CARD_SUCCESS:
		return "card_to_card_success"
	case TRANSACTION_STATUS_CARD_TO_CARD_FAILED:
		return "card_to_card_failed"

	case TRANSACTION_STATUS_CHIPPAY_EXPRESS_PENDING:
		return "cp_express_pending"
	case TRANSACTION_STATUS_CHIPPAY_EXPRESS_PROCESSING:
		return "cp_express_processing"
	case TRANSACTION_STATUS_CHIPPAY_EXPRESS_SUCCESS:
		return "cp_express_success"
	case TRANSACTION_STATUS_CHIPPAY_EXPRESS_FAILED:
		return "cp_express_failed"

	case TRANSACTION_STATUS_MERCHANT_ADJUST_BALANCE_PENDING:
		return "merchant_adjust_balance_pending"
	case TRANSACTION_STATUS_MERCHANT_ADJUST_BALANCE_SUCCESS:
		return "merchant_adjust_balance_success"
	case TRANSACTION_STATUS_MERCHANT_ADJUST_BALANCE_FAILED:
		return "merchant_adjust_balance_failed"

	case TRANSACTION_STATUS_MANUAL_PENDING:
		return "manual_pending"
	case TRANSACTION_STATUS_MANUAL_PENDING_AUDIT:
		return "manual_pending_audit"
	case TRANSACTION_STATUS_MANUAL_AUDITED:
		return "manual_audited"
	case TRANSACTION_STATUS_MANUAL_SUCCESS:
		return "manual_success"
	case TRANSACTION_STATUS_MANUAL_AUDIT_FAILED:
		return "manual_audit_failed"
	case TRANSACTION_STATUS_MANUAL_FAILED:
		return "manual_failed"

	case TRANSACTION_STATUS_POINT_ACCURAL_PENDING:
		return "point_accural_pending"
	case TRANSACTION_STATUS_POINT_ACCURAL_SUCCESS:
		return "point_accural_success"
	case TRANSACTION_STATUS_POINT_ACCURAL_FAILED:
		return "point_accural_failed"

	case TRANSACTION_STATUS_POINT_REDEEM_PENDING:
		return "point_redeem_pending"
	case TRANSACTION_STATUS_POINT_REDEEM_PROCESSING:
		return "point_redeem_processing"
	case TRANSACTION_STATUS_POINT_REDEEM_SUCCESS:
		return "point_redeem_success"
	case TRANSACTION_STATUS_POINT_REDEEM_FAILED:
		return "point_redeem_failed"

	case TRANSACTION_STATUS_INTEREST_PENDING:
		return "interest_pending"
	case TRANSACTION_STATUS_INTEREST_SUCCESS:
		return "interest_success"
	case TRANSACTION_STATUS_INTEREST_FAILED:
		return "interest_failed"

	case TRANSACTION_STATUS_FINANCIAL_TRANSFER_PENDING:
		return "financial_transfer_pending"
	case TRANSACTION_STATUS_FINANCIAL_TRANSFER_SUCCESS:
		return "financial_transfer_success"
	case TRANSACTION_STATUS_FINANCIAL_TRANSFER_FAILED:
		return "financial_transfer_failed"

	case TRANSACTION_STATUS_WHALE_PAY_PENDING:
		return "pay_pending"
	case TRANSACTION_STATUS_WHALE_PAY_CLOSED:
		return "pay_cleared"
	case TRANSACTION_STATUS_WHALE_PAY_FAILED:
		return "pay_declined"
	case TRANSACTION_STATUS_WHALE_PAY_CREDIT:
		return "pay_refunded"
	case TRANSACTION_STATUS_WHALE_PAY_REVERSAL:
		return "pay_void"

	case TRANSACTION_STATUS_DISTRIBUTE_DISCOUNT_PENDING:
		return "distribute_pending"
	case TRANSACTION_STATUS_DISTRIBUTE_DISCOUNT_PROCESSING:
		return "distribute_process"
	case TRANSACTION_STATUS_DISTRIBUTE_DISCOUNT_SUCCESS:
		return "distribute_success"
	case TRANSACTION_STATUS_DISTRIBUTE_DISCOUNT_FAILED:
		return "distribute_fail"

	case TRANSACTION_STATUS_PAYCRYPTO_PAY_PENDING:
		return "pay_cleared"
	case TRANSACTION_STATUS_PAYCRYPTO_PAY_CLOSED:
		return "pay_cleared"
	case TRANSACTION_STATUS_PAYCRYPTO_PAY_FAILED:
		return "pay_declined"
	case TRANSACTION_STATUS_PAYCRYPTO_PAY_REFUNDED:
		return "pay_refunded"
	case TRANSACTION_STATUS_PAYCRYPTO_PAY_REVERSAL:
		return "pay_refunded"

	case TRANSACTION_STATUS_BINANCE_PAY_PENDING:
		return "pay_pending"
	case TRANSACTION_STATUS_BINANCE_PAY_SUCCESS:
		return "pay_success"
	case TRANSACTION_STATUS_BINANCE_PAY_FAILED:
		return "pay_failed"
	case TRANSACTION_STATUS_BINANCE_PAY_PROCESSING:
		return "pay_processing"
	case TRANSACTION_STATUS_BINANCE_PAY_CLOSED:
		return "pay_closed"

	case TRANSACTION_STATUS_SYSTEM_CHARGE_PENDING:
		return "system_recharge_pending"
	case TRANSACTION_STATUS_SYSTEM_CHARGE_SUCCESS:
		return "system_recharge_success"
	case TRANSACTION_STATUS_SYSTEM_CHARGE_FAILED:
		return "system_recharge_failed"

	case TRANSACTION_STATUS_ETHERFI_PENDING:
		return "pay_pending"
	case TRANSACTION_STATUS_ETHERFI_CLEARED:
		return "pay_cleared"
	case TRANSACTION_STATUS_ETHERFI_CANCELLED:
		return "pay_declined"
	case TRANSACTION_STATUS_ETHERFI_DECLINED:
		return "pay_declined"
	case TRANSACTION_STATUS_ETHERFI_REFUND:
		return "pay_refunded"
	case TRANSACTION_STATUS_ETHERFI_SUCCESS:
		return "pay_success"
	case TRANSACTION_STATUS_ETHERFI_FAILED:
		return "pay_failed"
	}
	return ""
}

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

func (tc TransferChannel) String() string {
	switch tc {
	case TRANSFER_CHANNEL_USER_ID:
		return "user_id"
	case TRANSFER_CHANNEL_EMAIL:
		return "email"
	case TRANSFER_CHANNEL_CARD_ID:
		return "card_id"
	case TRANSFER_CHANNEL_ADDRESS:
		return "address"
	case TRANSFER_CHANNEL_WITHDRAW:
		return "withdraw"
	}
	return ""
}

func (v CardProductVendor) String() string {
	switch v {
	case CARD_PRODUCT_VENDOR_REAP:
		return "reap"
	case CARD_PRODUCT_VENDOR_WHALE:
		return "whale"
	case CARD_PRODUCT_VENDOR_PAYCRYPTO:
		return "paycrypto"
	case CARD_PRODUCT_VENDOR_ETHERFI:
		return "etherfi"
	}
	return ""
}

func (wct WhaleCardType) String() string {
	switch wct {
	case WHALE_CARD_TYPE_VISA:
		return "VISA"
	case WHALE_CARD_TYPE_VISA_PREMIUM:
		return "VISA_PREMIUM"
	case WHALE_CARD_TYPE_VISA_PREMIUM_USD:
		return "VISA_PREMIUM_USD"
	case WHALE_CARD_TYPE_VISA_PREMIUM_HKD:
		return "VISA_PREMIUM_HKD"
	case WHALE_CARD_TYPE_MASTER:
		return "MASTER"
	}
	return ""
}

func (ws WithdrawScene) String() string {
	switch ws {
	case WITHDRAW_SCENE_NO_REVIEW:
		return "user_no_review"
	case WITHDRAW_SCENE_ADMIN_REVIEWED:
		return "admin_reviewed"
	case WITHDRAW_SCENE_ADMIN_APPROVED:
		return "admin_approved"
	case WITHDRAW_SCENE_AUTO_REVIEWED:
		return "auto_reviewed"
	case WITHDRAW_SCENE_AUTO_APPROVED:
		return "auto_approved"
	}
	return ""
}

func (ws WalletStatus) String() string {
	switch ws {
	case WALLET_STATUS_ENABLED:
		return "enabled"
	case WALLET_STATUS_DISABLED:
		return "disabled"
	}
	return ""
}

// 定義幣種組別
var currencyGroups = map[Currency]int{
	CURRENCY_USD:    1,
	CURRENCY_USDT:   1,
	CURRENCY_USDC:   1,
	CURRENCY_DAI:    1,
	CURRENCY_EPOINT: 1,
	// 將來可以在這裡新增其他幣種組別
}

// 判斷兩個幣種是否屬於同一組
func CurrencyIsSameGroup(currency1, currency2 Currency) bool {
	group1, keyExist1 := currencyGroups[currency1]
	group2, keyExist2 := currencyGroups[currency2]

	// 只有當兩個幣種都存在於同一組時才返回 true
	return keyExist1 && keyExist2 && group1 == group2
}

func (PasskeyVerifyStatus) PasskeyVerifyStatusFromInt(s int) PasskeyVerifyStatus {
	switch s {
	case 1:
		return PASSKEY_VERIFY_STATUS_PENDING
	case 2:
		return PASSKEY_VERIFY_STATUS_SUCCESS
	case 3:
		return PASSKEY_VERIFY_STATUS_FAILED
	}
	return 0
}

func (pvs PasskeyVerifyStatus) Integer() int {
	switch pvs {
	case PASSKEY_VERIFY_STATUS_PENDING:
		return 1
	case PASSKEY_VERIFY_STATUS_SUCCESS:
		return 2
	case PASSKEY_VERIFY_STATUS_FAILED:
		return 3
	}
	return 0
}

func (Language) FromString(s string) Language {
	sLower := strings.ToLower(s)
	switch sLower {
	case "zh":
		return LANGUAGE_CHINESE
	case "zh_hans":
		return LANGUAGE_SIMPLIFIED_CHINESE
	case "zh_hant":
		return LANGUAGE_TRADITIONAL_CHINESE_TW
	case "zh_hk":
		return LANGUAGE_TRADITIONAL_CHINESE_HK
	case "zh_sg":
		return LANGUAGE_SIMPLIFIED_CHINESE_SG
	case "en":
		return LANGUAGE_ENGLISH
	case "en_us":
		return LANGUAGE_ENGLISH_US
	case "en_gb":
		return LANGUAGE_ENGLISH_GB
	case "en_au":
		return LANGUAGE_ENGLISH_AU
	case "en_ca":
		return LANGUAGE_ENGLISH_CA
	case "de":
		return LANGUAGE_GERMAN
	case "de_de":
		return LANGUAGE_GERMAN_DE
	case "de_at":
		return LANGUAGE_GERMAN_AT
	case "de_ch":
		return LANGUAGE_GERMAN_CH
	case "fr":
		return LANGUAGE_FRENCH
	case "fr_fr":
		return LANGUAGE_FRENCH_FR
	case "fr_ca":
		return LANGUAGE_FRENCH_CA
	case "es":
		return LANGUAGE_SPANISH
	case "es_es":
		return LANGUAGE_SPANISH_ES
	case "es_mx":
		return LANGUAGE_SPANISH_MX
	case "it":
		return LANGUAGE_ITALIAN
	case "pt":
		return LANGUAGE_PORTUGUESE
	case "pt_br":
		return LANGUAGE_PORTUGUESE_BR
	case "nl":
		return LANGUAGE_DUTCH
	case "ru":
		return LANGUAGE_RUSSIAN
	case "pl":
		return LANGUAGE_POLISH
	case "tr":
		return LANGUAGE_TURKISH
	case "ja":
		return LANGUAGE_JAPANESE
	case "ko":
		return LANGUAGE_KOREAN
	case "th":
		return LANGUAGE_THAI
	case "vi":
		return LANGUAGE_VIETNAMESE
	case "hi":
		return LANGUAGE_HINDI
	case "ar":
		return LANGUAGE_ARABIC
	case "af":
		return LANGUAGE_AFRIKAANS
	case "bg":
		return LANGUAGE_BULGARIAN
	case "ca":
		return LANGUAGE_CATALAN
	case "cs":
		return LANGUAGE_CZECH
	case "da":
		return LANGUAGE_DANISH
	case "el":
		return LANGUAGE_GREEK
	case "fi":
		return LANGUAGE_FINNISH
	case "he":
		return LANGUAGE_HEBREW
	case "hu":
		return LANGUAGE_HUNGARIAN
	case "id":
		return LANGUAGE_INDONESIAN
	case "no":
		return LANGUAGE_NORWEGIAN
	case "ro":
		return LANGUAGE_ROMANIAN
	case "sk":
		return LANGUAGE_SLOVAK
	case "sv":
		return LANGUAGE_SWEDISH
	case "uk":
		return LANGUAGE_UKRAINIAN
	}
	return ""
}

func (lang Language) String() string {
	switch lang {
	case LANGUAGE_CHINESE:
		return "zh"
	case LANGUAGE_SIMPLIFIED_CHINESE:
		return "zh_hans"
	case LANGUAGE_TRADITIONAL_CHINESE_TW:
		return "zh_hant"
	case LANGUAGE_TRADITIONAL_CHINESE_HK:
		return "zh_hk"
	case LANGUAGE_SIMPLIFIED_CHINESE_SG:
		return "zh_sg"
	case LANGUAGE_ENGLISH:
		return "en"
	case LANGUAGE_ENGLISH_US:
		return "en_us"
	case LANGUAGE_ENGLISH_GB:
		return "en_gb"
	case LANGUAGE_ENGLISH_AU:
		return "en_au"
	case LANGUAGE_ENGLISH_CA:
		return "en_ca"
	case LANGUAGE_GERMAN:
		return "de"
	case LANGUAGE_GERMAN_DE:
		return "de_de"
	case LANGUAGE_GERMAN_AT:
		return "de_at"
	case LANGUAGE_GERMAN_CH:
		return "de_ch"
	case LANGUAGE_FRENCH:
		return "fr"
	case LANGUAGE_FRENCH_FR:
		return "fr_fr"
	case LANGUAGE_FRENCH_CA:
		return "fr_ca"
	case LANGUAGE_SPANISH:
		return "es"
	case LANGUAGE_SPANISH_ES:
		return "es_es"
	case LANGUAGE_SPANISH_MX:
		return "es_mx"
	case LANGUAGE_ITALIAN:
		return "it"
	case LANGUAGE_PORTUGUESE:
		return "pt"
	case LANGUAGE_PORTUGUESE_BR:
		return "pt_br"
	case LANGUAGE_DUTCH:
		return "nl"
	case LANGUAGE_RUSSIAN:
		return "ru"
	case LANGUAGE_POLISH:
		return "pl"
	case LANGUAGE_TURKISH:
		return "tr"
	case LANGUAGE_JAPANESE:
		return "ja"
	case LANGUAGE_KOREAN:
		return "ko"
	case LANGUAGE_THAI:
		return "th"
	case LANGUAGE_VIETNAMESE:
		return "vi"
	case LANGUAGE_HINDI:
		return "hi"
	case LANGUAGE_ARABIC:
		return "ar"
	case LANGUAGE_AFRIKAANS:
		return "af"
	case LANGUAGE_BULGARIAN:
		return "bg"
	case LANGUAGE_CATALAN:
		return "ca"
	case LANGUAGE_CZECH:
		return "cs"
	case LANGUAGE_DANISH:
		return "da"
	case LANGUAGE_GREEK:
		return "el"
	case LANGUAGE_FINNISH:
		return "fi"
	case LANGUAGE_HEBREW:
		return "he"
	case LANGUAGE_HUNGARIAN:
		return "hu"
	case LANGUAGE_INDONESIAN:
		return "id"
	case LANGUAGE_NORWEGIAN:
		return "no"
	case LANGUAGE_ROMANIAN:
		return "ro"
	case LANGUAGE_SLOVAK:
		return "sk"
	case LANGUAGE_SWEDISH:
		return "sv"
	case LANGUAGE_UKRAINIAN:
		return "uk"
	}
	return ""
}

func (l Language) IsValid() bool {
	switch l {
	case
		LANGUAGE_ENGLISH,
		LANGUAGE_SIMPLIFIED_CHINESE,
		LANGUAGE_TRADITIONAL_CHINESE_TW,
		LANGUAGE_KOREAN,
		LANGUAGE_JAPANESE,
		LANGUAGE_FRENCH,
		LANGUAGE_GERMAN,
		LANGUAGE_SPANISH,
		LANGUAGE_PORTUGUESE,
		LANGUAGE_RUSSIAN,
		LANGUAGE_ARABIC,
		LANGUAGE_VIETNAMESE:
		return true
	}
	return false
}

func (MsgOPCode) FromString(s string) MsgOPCode {
	stringUpper := strings.ToUpper(s)
	switch stringUpper {
	case "READ":
		return MSG_OPCODE_READ
	case "3DS_BALANCED":
		return MSG_OPCODE_FORWARD_3DS_BALANCED
	case "OTP":
		return MSG_OPCODE_OTP
	case "LOGIN_NOTIFY":
		return MSG_OPCODE_LOGIN_NOTIFY
	case "CARD_TRANSACTION":
		return MSG_OPCODE_CARD_TRANSACTION
	case "RECHARGE_SUCCESSFUL":
		return MSG_OPCODE_RECHARGE_SUCCESSFUL
	case "CARD_FROZEN":
		return MSG_OPCODE_CARD_FROZEN
	case "CARD_UNFROZEN":
		return MSG_OPCODE_CARD_UNFROZEN
	case "CARD_DELETE":
		return MSG_OPCODE_CARD_DELETE
	case "KYC_PASSED":
		return MSG_OPCODE_KYC_PASSED
	case "STATEMENT":
		return MSG_OPCODE_STATEMENT
	case "WITHDRAW_REQUEST":
		return MSG_OPCODE_WITHDRAW_REQUEST
	case "WITHDRAW_SUCCESS":
		return MSG_OPCODE_WITHDRAW_SUCCESS
	case "REQUEST_3DS":
		return MSG_OPCODE_REQUEST_3DS
	case "REAP_WALLET_BIND_OTP":
		return MSG_OPCODE_REAP_WALLET_BIND_OTP
	case "REAP_TRANSACTION_FAIL":
		return MSG_OPCODE_REAP_TRANSACTION_FAIL
	case "EXTRENAL_REQUEST_LOG":
		return MSG_OPCODE_EXTRENAL_REQUEST_LOG
	case "EXTRENAL_REQUEST_LOG_PARTITIONED":
		return MSG_OPCODE_EXTRENAL_REQUEST_LOG_PARTITIONED
	case "DAPP_APPLY_CARD_SUCESS":
		return MSG_OPCODE_DAPP_APPLY_CARD_SUCCESS
	case "ADMIN_NOTICE":
		return MSG_OPCODE_ADMIN_NOTICE
	case "REQUEST_WHALE_3DS":
		return MSG_OPCODE_REQUEST_WHALE_3DS
	case "WHALE_CARD_RISKY":
		return MSG_OPCODE_WHALE_CARD_RISKY
	case "KYC_NOTIFY":
		return MSG_OPCODE_KYC_NOTIFY
	case "USER_BLOCKED":
		return MSG_OPCODE_USER_BLOCKED
	case "USER_UNBLOCKED":
		return MSG_OPCODE_USER_UNBLOCKED
	case "KYC_NOTIFY_FINAL":
		return MSG_OPCODE_KYC_NOTIFY_FINAL
	case "CARD_BLOCKED":
		return MSG_OPCODE_CARD_BLOCKED
	case "CARD_UNBLOCKED":
		return MSG_OPCODE_CARD_UNBLOCKED
	case "USER_DELETE":
		return MSG_OPCODE_USER_DELETE
	case "CARD_TRANSACTION_OTP":
		return MSG_OPCODE_CARD_TRANSACTION_OTP
	case "APPLY_CARD_SUCCESS":
		return MSG_OPCODE_APPLY_CARD_SUCCESS
	case "APPLY_CARD_FAIL":
		return MSG_OPCODE_APPLY_CARD_FAIL
	case "CRYPTO_CARD_BLOCKED":
		return MSG_OPCODE_CRYPTO_CARD_BLOCKED
	case "CRYPTO_CARD_UNBLOCKED":
		return MSG_OPCODE_CRYPTO_CARD_UNBLOCKED
	case "CARD_AMOUNT_LOW_WARNING":
		return MSG_OPCODE_CARD_AMOUNT_LOW_WARNING
	default:
		// 如果找不到對應值，可以根據需要返回一個預設值或錯誤
		return MsgOPCode("")
	}
}

func (moc MsgOPCode) String() string {
	switch moc {
	case MSG_OPCODE_READ:
		return "READ"
	case MSG_OPCODE_FORWARD_3DS:
		return "3DS"
	case MSG_OPCODE_FORWARD_3DS_BALANCED:
		return "3DS_BALANCED"
	case MSG_OPCODE_OTP:
		return "OTP"
	case MSG_OPCODE_LOGIN_NOTIFY:
		return "LOGIN_NOTIFY"
	case MSG_OPCODE_CARD_TRANSACTION:
		return "CARD_TRANSACTION"
	case MSG_OPCODE_RECHARGE_SUCCESSFUL:
		return "RECHARGE_SUCCESSFUL"
	case MSG_OPCODE_CARD_FROZEN:
		return "CARD_FROZEN"
	case MSG_OPCODE_CARD_UNFROZEN:
		return "CARD_UNFROZEN"
	case MSG_OPCODE_CARD_DELETE:
		return "CARD_DELETE"
	case MSG_OPCODE_KYC_PASSED:
		return "KYC_PASSED"
	case MSG_OPCODE_STATEMENT:
		return "STATEMENT"
	case MSG_OPCODE_WITHDRAW_REQUEST:
		return "WITHDRAW_REQUEST"
	case MSG_OPCODE_WITHDRAW_SUCCESS:
		return "WITHDRAW_SUCCESS"
	case MSG_OPCODE_REQUEST_3DS:
		return "REQUEST_3DS"
	case MSG_OPCODE_REAP_WALLET_BIND_OTP:
		return "REAP_WALLET_BIND_OTP"
	case MSG_OPCODE_REAP_TRANSACTION_FAIL:
		return "REAP_TRANSACTION_FAIL"
	case MSG_OPCODE_EXTRENAL_REQUEST_LOG:
		return "EXTRENAL_REQUEST_LOG"
	case MSG_OPCODE_EXTRENAL_REQUEST_LOG_PARTITIONED:
		return "EXTRENAL_REQUEST_LOG_PARTITIONED"
	case MSG_OPCODE_DAPP_APPLY_CARD_SUCCESS:
		return "DAPP_APPLY_CARD_SUCESS"
	case MSG_OPCODE_ADMIN_NOTICE:
		return "ADMIN_NOTICE"
	case MSG_OPCODE_REQUEST_WHALE_3DS:
		return "REQUEST_WHALE_3DS"
	case MSG_OPCODE_WHALE_CARD_RISKY:
		return "WHALE_CARD_RISKY"
	case MSG_OPCODE_KYC_NOTIFY:
		return "KYC_NOTIFY"
	case MSG_OPCODE_USER_BLOCKED:
		return "USER_BLOCKED"
	case MSG_OPCODE_USER_UNBLOCKED:
		return "USER_UNBLOCKED"
	case MSG_OPCODE_KYC_NOTIFY_FINAL:
		return "KYC_NOTIFY_FINAL"
	case MSG_OPCODE_CARD_BLOCKED:
		return "CARD_BLOCKED"
	case MSG_OPCODE_CARD_UNBLOCKED:
		return "CARD_UNBLOCKED"
	case MSG_OPCODE_USER_DELETE:
		return "USER_DELETE"
	case MSG_OPCODE_CUSTOMISED:
		return "CUSTOMISED"
	case MSG_OPCODE_CARD_BLOCKED_CONSECUTIVE_FAILURE:
		return "CARD_BLOCKED_CONSECUTIVE_FAILURE"
	case MSG_OPCODE_CARD_REINSTATE_AFTER_BLOCKED:
		return "CARD_REINSTATE_AFTER_BLOCKED"
	case MSG_OPCODE_DELIVERY_CARD_IN_PROGRESS:
		return "DELIVERY_CARD_IN_PROGRESS"
	case MSG_OPCODE_DELIVERY_CARD_SENT_TO_USER:
		return "DELIVERY_CARD_SENT_TO_USER"
	case MSG_OPCODE_CARD_TRANSACTION_OTP:
		return "CARD_TRANSACTION_OTP"
	case MSG_OPCODE_APPLY_CARD_SUCCESS:
		return "APPLY_CARD_SUCCESS"
	case MSG_OPCODE_APPLY_CARD_FAIL:
		return "APPLY_CARD_FAIL"
	case MSG_OPCODE_CRYPTO_CARD_BLOCKED:
		return "CRYPTO_CARD_BLOCKED"
	case MSG_OPCODE_CRYPTO_CARD_UNBLOCKED:
		return "CRYPTO_CARD_UNBLOCKED"
	case MSG_OPCODE_CARD_AMOUNT_LOW_WARNING:
		return "CARD_AMOUNT_LOW_WARNING"
	default:
		return ""
	}
}

func (wp WorkerPool) String() string {
	switch wp {
	case WORKER_POOL_DEFAULT:
		return "default"
	case WORKER_POOL_OPEN_REQUEST_LOG:
		return "open_request_log"
	default:
		return ""
	}
}

func (ThirdPartyDepositStatus) ChangellyFromString(s string) ThirdPartyDepositStatus {
	sLower := strings.ToLower(s)
	switch sLower {
	// crypto
	case "created":
		return TRANSACTION_STATUS_CREATED
	case "pending":
		return TRANSACTION_STATUS_PENDING
	case "hold":
		return TRANSACTION_STATUS_HOLD
	case "refunded":
		return TRANSACTION_STATUS_REFUNDED
	case "expired":
		return TRANSACTION_STATUS_EXPIRED
	case "failed":
		return TRANSACTION_STATUS_FAILED
	case "complete":
		return TRANSACTION_STATUS_COMPLETE
	}
	return 0
}

func (uta UserTempAddressUsage) String() string {
	switch uta {
	case ADDRESS_USAGE_ORIGINAL:
		return "ORIGINAL"
	case ADDRESS_USAGE_DAPP_CARD_APPLY:
		return "DAPP_CARD_APPLY"
	case ADDRESS_USAGE_DAPP_DEPOSIT:
		return "DAPP_DEPOSIT"
	case ADDRESS_USAGE_CHANGELLY:
		return "CHANGELLY"
	case ADDRESS_USAGE_BINANCEPAY:
		return "BINANCEPAY"
	default:
		return ""
	}
}

func (ct CardType) String() string {
	switch ct {
	case CARD_TYPE_REAP:
		return "reap"
	case CARD_TYPE_WHALE:
		return "whale"
	case CARD_TYPE_WHALE_AI:
		return "whale_ai"
	case CARD_TYPE_PAYCRYPTO:
		return "paycrypto"
	}
	return ""
}

func (g Gender) String() string {
	switch g {
	case GENDER_MALE:
		return "male"
	case GENDER_FEMALE:
		return "female"
	default:
		return ""
	}
}

func (b UserBlockStatus) String() string {
	switch b {
	case USER_BLOCK_STATUS_NORMAL:
		return "NORMAL"
	case USER_BLOCK_STATUS_BLOCK:
		return "BLOCK"
	default:
		return ""
	}
}

func (r UserBlockReason) String() string {
	switch r {
	case USER_STATUS_EUSD_USER_BLOCK:
		return "EUSD_USER_BLOCK"
	case USER_STATUS_EUSD_CUSTOMER_REQUIRE_BLOCK:
		return "EUSD_CUSTOMER_REQUIRE_BLOCK"
	case USER_STATUS_EUSD_REAP_KYC:
		return "EUSD_REAP_KYC"
	default:
		return ""
	}
}

func (r UserBlockReason) IsValid() bool {
	switch r {
	case USER_STATUS_EUSD_USER_BLOCK:
		return true
	case USER_STATUS_EUSD_CUSTOMER_REQUIRE_BLOCK:
		return true
	case USER_STATUS_EUSD_REAP_KYC:
		return true
	default:
		return false
	}
}

func (r UserBlockReason) ToCode(code string) UserBlockReason {
	switch code {
	case "EUSD_USER_BLOCK":
		return 1
	case "EUSD_CUSTOMER_REQUIRE_BLOCK":
		return 2
	case "EUSD_REAP_KYC":
		return 3
	default:
		return 0
	}
}

func (ws WithdrawStatus) String() string {
	switch ws {
	case WITHDRAW_STATUS_PENDING:
		return "pending"
	case WITHDRAW_STATUS_FAILED:
		return "failed"
	case WITHDRAW_STATUS_SUCCESS:
		return "success"
	case WITHDRAW_STATUS_FREEZED:
		return "freezed"
	case WITHDRAW_STATUS_COINSDO_PENDING:
		return "coinsdo_pending"
	case WITHDRAW_STATUS_COINSDO_PASSED:
		return "coinsdo_passed"
	case WITHDRAW_STATUS_COINSDO_REJECTED:
		return "coinsdo_rejected"
	case WITHDRAW_STATUS_COINSDO_CANCELLED:
		return "coinsdo_cancelled"
	case WITHDRAW_STATUS_COINSDO_FAILED:
		return "coinsdo_failed"
	}
	return ""
}

func (s KYCStatus) String() string {
	switch s {
	case KYC_STATUS_PENDING:
		return "pending"
	case KYC_STATUS_PASS:
		return "pass"
	case KYC_STATUS_REJECT:
		return "reject"
	case KYC_STATUS_FAILED:
		return "failed"
	default:
		return ""
	}

}

func (rs ReapTransactionStatus) String() string {
	switch rs {
	case REAP_TRANSACTION_STATUS_REQUESTING:
		return "requesting"
	case REAP_TRANSACTION_STATUS_AUTHORIZING:
		return "authorizing"
	case REAP_TRANSACTION_STATUS_PENDING:
		return "pending"
	case REAP_TRANSACTION_STATUS_REVERSAL:
		return "reversal"
	case REAP_TRANSACTION_STATUS_CLEARED:
		return "cleared"
	case REAP_TRANSACTION_STATUS_DECLINED:
		return "declined"
	case REAP_TRANSACTION_STATUS_VOID:
		return "void"
	case REAP_TRANSACTION_STATUS_REFUNDED:
		return "refunded"
	case REAP_TRANSACTION_STATUS_UNRELATED_PENDING:
		return "unrelated_pending"
	case REAP_TRANSACTION_STATUS_UNRELATED_REFUNDED:
		return "unrelated_refunded"
	case REAP_TRANSACTION_STATUS_UNRELATED_FAILED:
		return "unrelated_failed"
	case REAP_TRANSACTION_STATUS_VERIFY_PENDING:
		return "verify_pending"
	case REAP_TRANSACTION_STATUS_VERIFY_SUCCESS:
		return "verify_success"
	case REAP_TRANSACTION_STATUS_VERIFY_FAILED:
		return "verify_failed"
	default:
		return ""
	}
}

func (e EtherfiTransactionLifecycleEventAction) Index() int {
	switch e {
	case ETHERFI_REQUESTED:
		return 1
	case ETHERFI_CREATED:
		return 2
	case ETHERFI_UPDATED:
		return 3
	case ETHERFI_COMPLETED:
		return 4
	}
	return math.MaxInt
}

func (v PhysicalCardRequestItemStatus) String() string {
	switch v {
	case PHYSICAL_CARD_REQUEST_ITEM_STATUS_PENDING:
		return "pending"
	case PHYSICAL_CARD_REQUEST_ITEM_STATUS_SUCCESS:
		return "success"
	case PHYSICAL_CARD_REQUEST_ITEM_STATUS_FAIIED:
		return "failed"
	}
	return ""
}

func (v PhysicalCardRequestOrderMerchant) String() string {
	switch v {
	case PHYSICAL_CARD_REQUEST_ORDER_MERCHANT_REAP:
		return "reap"
	case PHYSICAL_CARD_REQUEST_ORDER_MERCHANT_MONETA:
		return "moneta"
	case PHYSICAL_CARD_REQUEST_ORDER_MERCHANT_JINBEL:
		return "jinbel"
	}
	return ""
}

func (v EtherfiTransactionStatus) String() string {
	switch v {
	case ETHERFI_PENDING:
		return "PENDING"
	case ETHERFI_CLEARED:
		return "CLEARED"
	case ETHERFI_DECLINED:
		return "DECLINED"
	case ETHERFI_CANCELLED:
		return "CANCELLED"
	case ETHERFI_REFUND:
		return "REFUND"
	}
	return ""
}
