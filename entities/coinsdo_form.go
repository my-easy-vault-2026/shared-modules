package entities

import "shared-modules/common"

type OuterRequest struct {
	Data string `json:"data"`
	Sign string `json:"sign"`
}

type CoinsdoNotifyForm struct {
	TXHash           string `json:"txHash"`
	TXIndex          string `json:"txIndex"`
	FromAddress      string `json:"fromAddress"`
	ToAddress        string `json:"toAddress"`
	Amount           string `json:"amount"`
	TXFee            string `json:"txFee"`
	TXTime           string `json:"txTime"`
	DeviceUUID       string `json:"deviceUuid"`
	Currency         string `json:"currency"`
	Flag             string `json:"flag"`
	FeeSymbol        string `json:"feeSymbol"`
	BlockNumber      string `json:"blockNumber"`
	IsReachedConfirm string `json:"isReachedConfirm"`
	BlockConfirm     string `json:"blockConfirm"`
	Mainnet          string `json:"mainnet"`
	CoinName         string `json:"coinName"`
	CoinsDoId        string `json:"coinsDoId"`
	Protocol         string `json:"protocol"`
	TokenAddress     string `json:"tokenAddress"`
	CoinType         string `json:"coinType"`
	RecordId         string `json:"recordId"`
	ManualCallback   string `json:"manualCallback"`
}

type CoinsdoNotifyConvertForm struct {
	TXHash           string `json:"txHash"`
	TXIndex          string `json:"txIndex"`
	FromAddress      string `json:"fromAddress"`
	ToAddress        string `json:"toAddress"`
	Amount           string `json:"amount"`
	TXFee            string `json:"txFee"`
	TXTime           int64  `json:"txTime,string"`
	DeviceUUID       string `json:"deviceUuid"`
	Currency         string `json:"currency"`
	Flag             string `json:"flag"`
	FeeSymbol        string `json:"feeSymbol"`
	BlockNumber      uint64 `json:"blockNumber,string"`
	IsReachedConfirm string `json:"isReachedConfirm"`
	BlockConfirm     string `json:"blockConfirm"`
	Mainnet          string `json:"mainnet"`
	CoinName         string `json:"coinName"`
	CoinsDoId        uint64 `json:"coinsDoId,string"`
	Protocol         string `json:"protocol"`
	TokenAddress     string `json:"tokenAddress"`
	CoinType         string `json:"coinType"`
	RecordId         string `json:"recordId"`
	ManualCallback   string `json:"manualCallback"`
}

type AddSignIdentityForm struct {
	Name       string `json:"name" validate:"required"`
	DeviceID   string `json:"deviceId" validate:"required"`
	PublicKey  string `json:"publicKey" validate:"required"`
	BindingURL string `json:"bindingUrl"`
	Remark     string `json:"remark"`
}

type WithdrawNotifyForm struct {
	BusinessID     string                       `json:"businessId"`
	Mainnet        string                       `json:"mainnet"`
	CoinName       string                       `json:"coinName"`
	TXFee          string                       `json:"txFee"`
	FeeSymbol      string                       `json:"feeSymbol"`
	TXHash         string                       `json:"txHash"`
	BlockNumber    string                       `json:"blockNumber"`
	TXTime         string                       `json:"txTime"`
	TXMemo         string                       `json:"txMemo"`
	WithdrawStatus common.CoinsdoWithdrawStatus `json:"withdrawStatus"`
}

type ReconcileDepositsForm struct {
	FromID    uint64 `json:"fromId" default:"0"`
	ToID      uint64 `json:"toId" default:"99999999999"`
	BatchSize uint64 `json:"batchSize" default:"10"`
}
