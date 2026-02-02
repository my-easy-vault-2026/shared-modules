package common

import (
	"time"

	"github.com/shopspring/decimal"
)

type ExchangeRate struct {
	BaseCurrency  Currency        `json:"baseCurrency"`
	QuoteCurrency Currency        `json:"quoteCurrency"`
	Rate          decimal.Decimal `json:"rate"`
	Purpose       RatePurpose     `json:"purpose,omitempty"`
	Timestamp     time.Time       `json:"timestamp"`
}

type RateLimit struct {
	Limit     int
	Remaining int
	Used      int
	Reset     time.Time
}

type Msg struct {
	OP         MsgOPCode `json:"op"`
	MsgID      string    `json:"msgId"`                // 訊息唯一識別碼
	SequenceID string    `json:"sequenceId,omitempty"` // 事件流程中的排序
	NodeID     string    `json:"nodeId,omitempty"`     // 指定由哪個節點處理
	RecordID   uint64    `json:"recordId,omitempty"`   // 實際業務資料表中的id
	Subject    string    `json:"subject,omitempty"`
	UserID     uint64    `json:"userId,omitempty"`
	Role       Role      `json:"role,omitempty"`
	Msg        string    `json:"msg"`
}

type ChannelKey struct{}

type Page struct {
	Current  int `json:"current" validate:"min=1,required"`
	PageSize int `json:"pageSize" validate:"max=3000,required"`
}

type PageData[T interface{}] struct {
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
	Records  T   `json:"records"`
}
