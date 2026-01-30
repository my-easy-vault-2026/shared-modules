package entities

import (
	"shared-modules/utils"

	"github.com/shopspring/decimal"
)

type MerchantExportBalanceChangeVO struct {
	ID                  uint64          `json:"id,string"`
	UserID              uint64          `json:"userId,string"`
	Email               string          `json:"email"`
	AssetID             uint64          `json:"assetId,string"`
	OrderNo             string          `json:"orderNo"`
	Amount              decimal.Decimal `json:"amount"`
	CurrentAmount       decimal.Decimal `json:"currentAmount"`
	CurrentFrozenAmount decimal.Decimal `json:"currentFrozenAmount"`
	Currency            string          `json:"currency"`
	TransactionType     string          `json:"transactionType"`
	BillType            string          `json:"billType"`
	OrderType           string          `json:"orderType"`
	CreatedAt           int64           `json:"createdAt,string"`
	UpdatedAt           int64           `json:"updatedAt,string"`
}

type MerchantBillVO struct {
	ID                  uint64          `json:"id,string" csv:"id"`
	UserID              uint64          `json:"userId,string" csv:"userId"`
	Email               string          `json:"email" csv:"email"`
	AssetID             uint64          `json:"assetId,string" csv:"assetId"`
	OrderNo             string          `json:"orderNo" csv:"orderNo"`
	Amount              decimal.Decimal `json:"amount" csv:"amount"`
	CurrentAmount       decimal.Decimal `json:"currentAmount" csv:"currentAmount"`
	CurrentFrozenAmount decimal.Decimal `json:"currentFrozenAmount" csv:"currentFrozenAmount"`
	Currency            string          `json:"currency" csv:"currency"`
	TransactionType     string          `json:"transactionType" csv:"transactionType"`
	BillType            string          `json:"billType" csv:"billType"`
	OrderType           string          `json:"orderType" csv:"orderType"`
	CreatedAt           int64           `json:"createdAt,string" csv:"createdAt"`
	UpdatedAt           int64           `json:"updatedAt,string" csv:"updatedAt"`
}

type PageMerchantBillVO struct {
	utils.PageData[[]*MerchantBillVO]
}

type MerchantAssetVO struct {
	Amount       decimal.Decimal `json:"amount"`
	FrozenAmount decimal.Decimal `json:"frozenAmount"`
}

type MerchantListAssetsVO struct {
	Assets map[string]*MerchantAssetVO `json:"assets"`
}

type MerchantRequestExportBalanceChangeVO struct {
	ExportToken string `json:"exportToken"`
}
