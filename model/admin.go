package model

// Admin 用户模型
type Admin struct {
	Model
	Adminname      string `json:"adminname"`
	PasswordDigest string `json:"password_digest,omitempty"`
	Salt           string `json:"salt,omitempty"`
	Status         int    `json:"status"`
	Nickname       string `json:"nickname"`
	Avatar         string `gorm:"size:1000" json:"avatar"`
}
