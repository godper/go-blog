package serialize

import "blog/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Status    int    `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
}

//UserToken wae
type UserToken struct {
	User  User   `json:"item"`
	Token string `json:"token"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
	}
}

//BuildUserToken 序列化token用户
func BuildUserToken(user model.User, token string) UserToken {
	return UserToken{
		User:  BuildUser(user),
		Token: token,
	}
}
