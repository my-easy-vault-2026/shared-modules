package entities

import (
	"shared-modules/common"
	"shared-modules/utils"
)

type AppFeatureParameterCategoriesForm struct {
	//
}

type AppFeatureParameterPageForm struct {
	Id            uint64                   `json:"id"`
	Key           common.ParameterKey      `json:"key"`
	CategoryID    common.ParameterCategory `json:"categoryId"`
	SecurityLevel common.AdminLevel        `json:"securityLevel"`
	Status        common.ParameterStatus   `json:"status"`
	utils.Page
}

type AppFeatureParameterGetForm struct {
	Id uint64 `json:"id" validate:"required"`
}

type AppFeatureParameterActivateForm struct {
	Id uint64 `json:"id" validate:"required"`
}

type AppFeatureParameterDeactivateForm struct {
	Id uint64 `json:"id" validate:"required"`
}

type AppFeatureParameterEditForm struct {
	Id          uint64 `json:"id" validate:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       string `json:"value" validate:"required"`
}

type AppFeatureParameterValidateForm struct {
	Value     string                    `json:"value" validate:"required"`
	ValueType common.ParameterValueType `json:"valueType" validate:"required"`
}
