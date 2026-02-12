package common

type MsgOPCode string // 訊息行為碼
const (
	MSG_OPCODE_READ   MsgOPCode = "READ"
	MSG_OPCODE_INFUND MsgOPCode = "INFUND"
)

type Msg struct {
	OP         MsgOPCode `json:"op"`
	MsgID      string    `json:"msgId"`                // 訊息唯一識別碼
	SequenceID string    `json:"sequenceId,omitempty"` // 事件流程中的排序
	NodeID     string    `json:"nodeId,omitempty"`     // 指定由哪個節點處理
	RecordID   string    `json:"recordId,omitempty"`   // 實際業務資料表中的編號
	Subject    string    `json:"subject,omitempty"`
	UserID     uint64    `json:"userId,omitempty"`
	Role       Role      `json:"role,omitempty"`
	Msg        string    `json:"msg"`
}
