package entities

type SignedForm struct {
	ID       string             `json:"id"`
	RawID    string             `json:"rawId"`
	Type     string             `json:"type"`
	Response VerifyResponseType `json:"response"`
	DeviceID string

	Data string `json:"data"`
	Sign string `json:"sign"`
}

type EncryptedForm struct {
	Data string `json:"data"`
}
