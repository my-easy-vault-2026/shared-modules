package entities

import (
	"shared-modules/common"
	"shared-modules/utils"
	"time"
)

type AppFeatureParameterCategoriesVO []*AppFeatureParameterCategoryVO

type AppFeatureParameterCategoryVO struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AppFeatureParameterVO struct {
	Id            uint64                    `json:"id"`
	Name          string                    `json:"name"`
	Description   string                    `json:"description"`
	Key           common.ParameterKey       `json:"key"`
	Value         string                    `json:"value"`
	SpecialValue  common.SpecialValue       `json:"specialValue"`
	ValueType     common.ParameterValueType `json:"valueType"`
	Currency      common.Currency           `json:"currency"`
	Unit          common.UnitType           `json:"unit"`
	Encrypt       common.EncryptType        `json:"encrypt"`
	CategoryId    common.ParameterCategory  `json:"categoryId"`
	CategoryName  string                    `json:"categoryName"`
	SecurityLevel common.AdminLevel         `json:"securityLevel"`
	Status        common.ParameterStatus    `json:"status"`
	CreatedAt     time.Time                 `json:"createdAt"`
	UpdatedAt     time.Time                 `json:"updatedAt"`
}

type PageAppFeatureParameterVO struct {
	utils.PageData[[]*AppFeatureParameterVO]
}
