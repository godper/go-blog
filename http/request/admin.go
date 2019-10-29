package request

// AdminRegister 管理用户注册表单
type AdminRegister struct {
	Adminname       string `form:"adminname" json:"adminname" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// AdminLogin 管理用户登录表单
type AdminLogin struct {
	Adminname string `form:"adminname" json:"adminname" binding:"required,min=5,max=30"`
	Password  string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// AdminInfoModify 管理用户修改信息
type AdminInfoModify struct {
	ID              uint   `form:"id"              json:"id"              binding:"required"`
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"omitempty,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"omitempty,min=8,max=40"`
}
