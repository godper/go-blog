package api

import (
	"blog/model"
	"blog/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

var srv *service.Service

//Init 初始化api
func Init() {
	srv = service.Srv
}

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

//CurrentUser 当前用户
func CurrentUser(c *gin.Context) *model.User {
	userID, ok := c.Get("userID")
	if !ok {
		return nil
	}
	userID, ok = userID.(uint)
	if !ok {
		return nil
	}
	if user, err := srv.Dao.GetUserByID(userID); err == nil {
		return &user
	}
	return nil

}

//CurrentAdmin 当前用户
func CurrentAdmin(c *gin.Context) *model.Admin {
	adminID, ok := c.Get("adminID")
	if !ok {
		return nil
	}
	fmt.Printf("%v", adminID)
	adminID, ok = adminID.(uint)
	if !ok {
		return nil
	}
	fmt.Printf("%v", adminID)
	if admin, err := srv.Dao.GetAdminByID(adminID); err == nil {

		return &admin
	}
	return nil
}
