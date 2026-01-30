package entities

import (
	"shared-modules/common"
	"shared-modules/utils"
	"time"

	"github.com/shopspring/decimal"
)

type CreateAdminUserVO struct {
}

type AdminVO struct {
	AdminID       uint64                 `json:"adminID,string"`
	Name          string                 `json:"name"`
	Role          common.Role            `json:"role"`
	MerchantID    uint64                 `json:"merchantID,string"`
	GroupID       uint64                 `json:"groupID,string"`
	GroupIDs      []uint64               `json:"groupIDs,string"`
	Level         common.AdminLevel      `json:"level"`
	Status        common.AdminUserStatus `json:"status"`
	PromotionCode string                 `json:"promotionCode"`
	PromotionLink string                 `json:"promotionLink"`
	CreatedBy     uint64                 `json:"createBy,string"`
	UpatedBy      uint64                 `json:"updateBy,string"`
	CreatedAt     int64                  `json:"createdAt,string"`
	UpdatedAt     int64                  `json:"updatedAt,string"`
}

type PageAdminVO struct {
	utils.PageData[[]*AdminVO]
}

type PageAdminTransactionRecordsVO struct {
	utils.PageData[[]*AdminTransactionRecordVO]
}

type PromotuionUrl struct {
	Url string `json:"url"`
}

type AdminUserKycFileVo struct {
	Kyc2FacePhotoToken    string `json:"kyc2FacePhotoToken"`
	Kyc3FacePhoto1Token   string `json:"kyc3FacePhoto1Token"`
	Kyc3FacePhoto2Token   string `json:"kyc3FacePhoto2Token"`
	DocumentFrontToken    string `json:"documentFrontToken"`
	DocumentBackToken     string `json:"documentBackToken"`
	DocumentHandheldToken string `json:"documentHandheldToken"`
}

type AdminTransactionRecordVO struct {
	ID                        uint64                       `json:"id,string"` //可能為空
	Title                     string                       `json:"title"`
	Type                      string                       `json:"type"`
	TransactionNO             string                       `json:"transactionNo"` //可能為空
	TransferTransactionNO     string                       `json:"transferTransactionNo,omitempty"`
	UserID                    uint64                       `json:"userId,string"`
	MerchantID                uint64                       `json:"merchantId,string,omitempty"`
	CardID                    uint64                       `json:"cardId,string,omitempty"` //可能為空
	IssueID                   string                       `json:"issueId,omitempty"`
	CentralCardID             uint64                       `json:"centralCardId,string,omitempty"`
	DelegatedID               uint64                       `json:"delegatedId,string,omitempty"`
	Income                    string                       `json:"income"`
	IncomeCategoryID          uint64                       `json:"incomeCategoryId,string"`
	IncomeCurrency            string                       `json:"incomeCurrency"`
	AgainstIncome             string                       `json:"againstIncome,omitempty"`                  //可能為空
	AgainstIncomeCategoryID   uint64                       `json:"againstIncomeCategoryId,string,omitempty"` //可能為空
	AgainstIncomeCurrency     string                       `json:"againstIncomeCurrency,omitempty"`          //可能為空
	Side                      string                       `json:"side,omitempty"`                           //可能為空
	FromCardID                uint64                       `json:"fromCardId,string,omitempty"`              //可能為空
	FromCategoryID            uint64                       `json:"fromCategoryId,string,omitempty"`          //可能為空
	FromCategory              string                       `json:"fromCategory,omitempty"`                   //可能為空
	FromCurrency              string                       `json:"fromCurrency,omitempty"`                   //可能為空
	FromAmount                string                       `json:"fromAmount,omitempty"`                     //可能為空
	FromDiscount              string                       `json:"fromDiscount,omitempty"`                   //可能為空
	FromUserID                uint64                       `json:"fromUserId,string,omitempty"`              //可能為空
	FromEmail                 string                       `json:"fromEmail,omitempty"`                      //可能為空
	FromAddress               string                       `json:"fromAddress,omitempty"`                    //可能為空
	FromMerchantID            uint64                       `json:"fromMerchantId,string,omitempty"`          //可能為空
	FromBalanceType           string                       `json:"fromBalanceType,omitempty"`                //可能為空
	FromRole                  string                       `json:"fromRole,omitempty"`                       //可能為空
	FromAlias                 string                       `json:"fromAlias,omitempty"`                      //可能為空
	FromPANNumber             string                       `json:"fromPanNumber,omitempty"`                  //可能為空
	ToCardID                  uint64                       `json:"toCardId,string,omitempty"`                //可能為空
	ToCategoryID              uint64                       `json:"toCategoryId,string,omitempty"`            //可能為空
	ToCategory                string                       `json:"toCategory,omitempty"`                     //可能為空
	ToCurrency                string                       `json:"toCurrency,omitempty"`                     //可能為空
	ToAmount                  string                       `json:"toAmount,omitempty"`                       //可能為空
	ToBonus                   string                       `json:"toBonus,omitempty"`                        //可能為空
	ToUserID                  uint64                       `json:"toUserId,string,omitempty"`                //可能為空
	ToEmail                   string                       `json:"toEmail,omitempty"`                        //可能為空
	ToAddress                 string                       `json:"toAddress,omitempty"`                      //可能為空
	ToMerchantID              uint64                       `json:"toMerchantId,string,omitempty"`            //可能為空
	ToBalanceType             string                       `json:"toBalanceType,omitempty"`                  //可能為空
	ToRole                    string                       `json:"toRole,omitempty"`                         //可能為空
	ToAlias                   string                       `json:"toAlias,omitempty"`                        //可能為空
	ToPANNumber               string                       `json:"toPanNumber,omitempty"`                    //可能為空
	ToProductID               uint64                       `json:"toProductID,string,omitempty"`             //可能為空
	ToProduct                 string                       `json:"toProduct,omitempty"`                      //可能為空
	ToVendorIconURL           string                       `json:"toVendorIconUrl,omitempty"`                //可能為空
	TransferToType            string                       `json:"transferToType,omitempty"`                 //可能為空
	TransferToCardID          uint64                       `json:"transferToCardId,string,omitempty"`        //可能為空
	TransferToCategoryID      uint64                       `json:"transferToCategoryId,string,omitempty"`    //可能為空
	TransferToCurrency        string                       `json:"transferToCurrency,omitempty"`             //可能為空
	TransferToUserID          uint64                       `json:"transferToUserId,string,omitempty"`        //可能為空
	TransferToAlias           string                       `json:"transferToAlias,omitempty"`                //可能為空
	TransferToPANNumber       string                       `json:"transferToPanNumber,omitempty"`            //可能為空
	TXHash                    string                       `json:"txHash,omitempty"`                         //可能為空
	Mainnet                   string                       `json:"mainnet,omitempty"`                        //可能為空
	MainnetName               string                       `json:"mainnetName,omitempty"`                    //可能為空
	Protocol                  string                       `json:"protocol,omitempty"`                       //可能為空
	TransferChannel           string                       `json:"transferChannel,omitempty"`                //可能為空
	RedirectLink              string                       `json:"redirectLink,omitempty"`                   //可能為空
	ReapMerchantID            string                       `json:"reapMerchantId,omitempty"`                 //可能為空
	MerchantMCCCode           string                       `json:"merchantMccCode,omitempty"`                //可能為空
	MerchantMCCCategory       string                       `json:"merchantMccCategory,omitempty"`            //可能為空
	MerchantCity              string                       `json:"merchantCity,omitempty"`                   //可能為空
	MerchantName              string                       `json:"merchantName,omitempty"`                   //可能為空
	MerchantState             string                       `json:"merchantState,omitempty"`                  //可能為空
	MerchantCountry           string                       `json:"merchantCountry,omitempty"`                //可能為空
	MerchantPostCode          string                       `json:"merchantPostCode,omitempty"`               //可能為空
	ReapTransactionID         string                       `json:"reapTransactionId,omitempty"`              //可能為空
	ParentTransactionID       string                       `json:"parentTransactionId,omitempty"`            //可能為空
	ReapChannel               string                       `json:"reapChannel,omitempty"`                    //可能為空
	ClearAmount               string                       `json:"clearAmount,omitempty"`                    //可能為空
	ReversalAmount            string                       `json:"reversalAmount,omitempty"`                 //可能為空
	ReversalTransactionAmount string                       `json:"reversalTransactionAmount,omitempty"`      //可能為空
	ReversalExchangeRate      *decimal.Decimal             `json:"reversalExchangeRate,string,omitempty"`    //可能為空
	RefundAmount              string                       `json:"refundAmount,omitempty"`                   //可能為空
	RefundTransactionAmount   string                       `json:"refundTransactionAmount,omitempty"`        //可能為空
	RefundExchangeRate        *decimal.Decimal             `json:"refundExchangeRate,string,omitempty"`      //可能為空
	DisplayRate               *decimal.Decimal             `json:"displayRate,string,omitempty"`             //可能為空
	ExchangeRate              *decimal.Decimal             `json:"exchangeRate,string,omitempty"`            //可能為空
	ExchangeFee               string                       `json:"exchangeFee,omitempty"`                    //可能為空
	ExchangeFeeCurrency       string                       `json:"exchangeFeeCurrency,omitempty"`            //可能為空
	DepositFee                string                       `json:"depositFee,omitempty"`                     //可能為空
	DepositFeeCurrency        string                       `json:"depositFeeCurrency,omitempty"`             //可能為空
	TransferFee               string                       `json:"transferFee,omitempty"`                    //可能為空
	TransferFeeCurrency       string                       `json:"transferFeeCurrency,omitempty"`            //可能為空
	WithdrawFee               string                       `json:"withdrawFee,omitempty"`                    //可能為空
	WithdrawFeeCurrency       string                       `json:"withdrawFeeCurrency,omitempty"`            //可能為空
	CardFee                   string                       `json:"cardFee,omitempty"`                        //可能為空
	CardFeeCurrency           string                       `json:"cardFeeCurrency,omitempty"`                //可能為空
	PhysicalCardFee           string                       `json:"physicalCardFee,omitempty"`                //可能為空
	PhysicalCardFeeCurrency   string                       `json:"physicalCardFeeCurrency,omitempty"`        //可能為空
	DeliveryFee               string                       `json:"deliveryFee,omitempty"`                    //可能為空
	DeliveryFeeCurrency       string                       `json:"deliveryFeeCurrency,omitempty"`            //可能為空
	TopUpFee                  string                       `json:"topUpFee,omitempty"`                       //可能為空
	TopUpFeeCurrency          string                       `json:"topUpFeeCurrency,omitempty"`               //可能為空
	TopDownFee                string                       `json:"topDownFee,omitempty"`                     //可能為空
	TopDownFeeCurrency        string                       `json:"topDownFeeCurrency,omitempty"`             //可能為空
	CardToCardFee             string                       `json:"cardToCardFee,omitempty"`                  //可能為空
	CardToCardFeeCurrency     string                       `json:"cardToCardFeeCurrency,omitempty"`          //可能為空
	ATMFee                    string                       `json:"atmFee,omitempty"`                         //可能為空
	ATMFeeCurrency            string                       `json:"atmFeeCurrency,omitempty"`                 //可能為空
	FXFee                     string                       `json:"fxFee,omitempty"`                          //可能為空
	FXFeeCurrency             string                       `json:"fxFeeCurrency,omitempty"`                  //可能為空
	CPExpressFee              string                       `json:"cpExpressFee,omitempty"`                   //可能為空
	CPExpressFeeCurrency      string                       `json:"cpExpressFeeCurrency,omitempty"`           //可能為空
	PromotionCode             string                       `json:"promotionCode,omitempty"`                  //可能為空
	PromotionType             string                       `json:"promotionType,omitempty"`                  //可能為空
	RewardStatus              string                       `json:"rewardStatus,omitempty"`                   //可能為空
	Status                    string                       `json:"status"`
	FailReason                string                       `json:"failReason,omitempty"`              //可能為空
	ReapDataLoss              string                       `json:"reapDataLoss,omitempty"`            //可能為空
	ReversalAt                int64                        `json:"reversalAt,string,omitempty"`       //可能為空
	ClearedAt                 int64                        `json:"clearedAt,string,omitempty"`        //可能為空
	RefundedAt                int64                        `json:"refundedAt,string,omitempty"`       //可能為空
	CoinsdoRepliedAt          int64                        `json:"coinsdoRepliedAt,string,omitempty"` //可能為空
	CreatedAt                 int64                        `json:"createdAt,string"`
	UpdatedAt                 int64                        `json:"updatedAt,string"`
	Remark                    *string                      `json:"remark,omitempty"`
	TypeNum                   common.TransactionRecordType `json:"typeNum"`
	StatusNum                 common.TransactionStatus     `json:"statusNum"`
	TxRawData                 string                       `json:"txRawData,omitempty"`
	VendorTxId                string                       `json:"vendorTxId,omitempty"`
	CurrentBalance            *decimal.Decimal             `json:"currentBalance,string,omitempty"` //可能為空
}

type AdminAccountVO struct {
	Id        uint64                 `json:"id,string"`
	Name      string                 `json:"name"`
	Role      common.Role            `json:"role"`
	GroupIDs  []uint64               `json:"groupIds,string"`
	Level     common.AdminLevel      `json:"level"`
	Status    common.AdminUserStatus `json:"status"`
	CreatedBy uint64                 `json:"createBy,string"`
	UpatedBy  uint64                 `json:"updateBy,string"`
	CreatedAt int64                  `json:"createdAt,string"`
	UpdatedAt int64                  `json:"updatedAt,string"`
}

type PageAdminAccountVO struct {
	utils.PageData[[]*AdminAccountVO]
}

type AgentAccountVO struct {
	Id            uint64                 `json:"id,string"`
	Name          string                 `json:"name"`
	Role          common.Role            `json:"role"`
	GroupIDs      []uint64               `json:"groupIDs,string"`
	Level         common.AdminLevel      `json:"level"`
	Status        common.AdminUserStatus `json:"status"`
	Discount      decimal.Decimal        `json:"discount,string"`
	DiscountType  *common.DiscountType   `json:"discountType"`
	PromotionCode string                 `json:"promotionCode"`
	PromotionLink string                 `json:"promotionLink"`
	CreatedBy     uint64                 `json:"createBy,string"`
	UpdatedBy     uint64                 `json:"updateBy,string"`
	CreatedAt     *time.Time             `json:"createdAt,string"`
	UpdatedAt     *time.Time             `json:"updatedAt,string"`
}

type PageAgentAccountVO struct {
	utils.PageData[[]*AgentAccountVO]
}
