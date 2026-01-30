package entities

import (
	"shared-modules/common"

	"github.com/shopspring/decimal"
)

type ChippayExpressRequestForm struct {
	ToCardID      uint64          `json:"toCardId,string"`
	FromAmount    decimal.Decimal `json:"fromAmount"`
	FromCurrency  string          `json:"fromCurrency"`
	PaymentMethod string          `json:"paymentMethod"` // bank, alipay, wechat_pay
	Link          string          `json:"link"`
}

type ChippayExpressWebhookForm struct {
	CompanyOrderNUM string                         `json:"companyOrderNum"`          // 商户订单号
	OTCOrderNum     string                         `json:"otcOrderNum"`              // ChipPay平台订单号
	CoinAmount      string                         `json:"coinAmount"`               // 订单币种数量，精度最多至小数点后4位
	CoinSign        string                         `json:"coinSign"`                 // 数字货币标识：usdt
	OrderType       common.CPExpressDirection      `json:"orderType,string"`         // 快捷买卖方向(1.Express buy, 2.Express sell)
	TradeStatus     *common.CPExpressWebhookStatus `json:"tradeStatus,string"`       // 交易状态(0:交易失败，1:交易成功，2：快捷批量卖单生成失败)
	CancelReason    string                         `json:"cancelReason,omitempty"`   // 卖单取消原因(当 orderType=2, tradeStatus=0 时才会返回)
	TradeOrderTime  string                         `json:"tradeOrderTime,omitempty"` // 订单交易时间(北京时间)
	UnitPrice       *decimal.Decimal               `json:"unitPrice,omitempty"`      // 数字货币单价
	Total           *decimal.Decimal               `json:"total,omitempty"`          // 用户付款的法币实际到账金额
	SuccessAmount   *decimal.Decimal               `json:"successAmount,omitempty"`  // 数字货币到账数量
	Sign            string                         `json:"sign"`                     // 参数签名
}

type ChippayExpressGetProcessingOrderForm struct{}
