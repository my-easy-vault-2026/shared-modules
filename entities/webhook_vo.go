package entities

import "shared-modules/common"

type WebhookNotifyVO struct {
	AuthorizationID string              `json:"authorization_id,omitempty"`
	ResponseCode    common.DeclinedCode `json:"response_code"`
}
