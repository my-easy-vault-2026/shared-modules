package entities

import "time"

type DeviceVO struct {
	ID         string    `json:"id"`
	Platform   string    `json:"platform,omitempty"`
	DeviceName string    `json:"deviceName,omitempty"`
	DeviceID   string    `json:"deviceID,omitempty"`
	IsMaster   bool      `json:"isMaster,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}

type ListDeviceVO struct {
	Records []*DeviceVO `json:"records"`
}

type VerifyCheck struct {
	IsVerify bool
	DeviceID string
}
