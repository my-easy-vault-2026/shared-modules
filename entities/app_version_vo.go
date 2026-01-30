package entities

// AppVersionVO 表示返回给客户端的版本信息
type AppVersionVO struct {
	Platform     string `json:"platform"`
	Version      string `json:"version"`
	UpdateURL    string `json:"update_url"`
	IsForce      bool   `json:"is_force"`
	ReleaseNotes string `json:"release_notes"`
}
