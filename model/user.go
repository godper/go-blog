package model

// User 用户模型
type User struct {
	Model
	Username       string `json:"username"`
	PasswordDigest string `json:"password_digest,omitempty"`
	Salt           string `json:"salt,omitempty"`
	Status         int    `json:"status"`
	Nickname       string `json:"nickname"`
	Avatar         string `gorm:"size:1000" json:"avatar"`
}
