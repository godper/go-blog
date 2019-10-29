package serialize

import "blog/model"

// Admin 用户序列化器
type Admin struct {
	ID        uint   `json:"id"`
	Adminname string `json:"adminname"`
	Nickname  string `json:"nickname"`
	Status    int    `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
}

//AdminToken wae
type AdminToken struct {
	Admin Admin  `json:"item"`
	Token string `json:"token"`
}

// BuildAdmin 序列化用户
func BuildAdmin(admin model.Admin) Admin {
	return Admin{
		ID:        admin.ID,
		Adminname: admin.Adminname,
		Nickname:  admin.Nickname,
		Status:    admin.Status,
		Avatar:    admin.Avatar,
		CreatedAt: admin.CreatedAt,
	}
}

//BuildAdminToken 序列化token用户
func BuildAdminToken(admin model.Admin, token string) AdminToken {
	return AdminToken{
		Admin: BuildAdmin(admin),
		Token: token,
	}
}
