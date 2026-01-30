package entities

type RegisterFinishForm struct {
	ID         string               `json:"id"`
	RawID      string               `json:"rawId"`
	Type       string               `json:"type"`
	Response   RegisterResponseType `json:"response"`
	Platform   string
	DeviceName string
	DeviceID   string
}

type VerifyFinishForm struct {
	ID       string             `json:"id"`
	RawID    string             `json:"rawId"`
	Type     string             `json:"type"`
	Response VerifyResponseType `json:"response"`
	DeviceID string
}

// VerifyStartForm represents a form structure used for verifying start and finish processes with an email field.
type VerifyStartForm struct {
	Email string `json:"email"`
}

type RegisterResponseType struct {
	AttestationObject string `json:"attestationObject"`
	ClientDataJSON    string `json:"clientDataJSON"`
}

type VerifyResponseType struct {
	AuthenticatorData string `json:"authenticatorData"`
	ClientDataJSON    string `json:"clientDataJSON"`
	Signature         string `json:"signature"`
	UserHandle        string `json:"userHandle"`
}

type RegisterStartForm struct {
	PinCode string `json:"pinCode" validate:"required"`
	OTPCode string `json:"otpCode" validate:"required"`
}

type RemoveDeviceForm struct {
	ID       string `json:"id" validate:"required"`
	DeviceID string `json:"-"`
}

type RemovePasskeyForm struct {
	DeviceID string `json:"-"`
	OTPCode  string `json:"otpCode" validate:"required,numeric"`
	PinCode  string `json:"pinCode" validate:"required,numeric"`
}

type VerifyDeviceForm struct {
	DeviceID string `json:"-"`
}
