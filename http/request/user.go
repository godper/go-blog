package request

// UserRegister 管理用户注册表单
type UserRegister struct {
	CaptchaID       string `form:"captcha_id" json:"captcha_id" binding:"required"`
	Captcha         string `form:"captcha" json:"captcha" binding:"required,min=6,max=6"`
	Username        string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// UserLogin 管理用户登录表单
type UserLogin struct {
	Username string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// UserInfoModify 管理用户修改信息
type UserInfoModify struct {
	ID              uint   `form:"id"              json:"id"              binding:"required"`
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"omitempty,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"omitempty,min=8,max=40"`
}
