package api

import (
	"blog/http/request"
	"blog/http/response"
	"blog/model"
	s "blog/serialize"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AdminRegister 用户注册接口
func AdminRegister(c *gin.Context) {
	var req request.AdminRegister
	r := response.NewR(c)

	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}

	admin, err := srv.AdminRegister(&req)
	if err != nil {
		r.FailedResponse(err)
	} else {
		data := s.BuildAdmin(admin)
		r.SuccessResponse(data, "注册成功")
	}
}

// AdminLogin 用户登录接口
func AdminLogin(c *gin.Context) {
	var req request.AdminLogin
	r := response.NewR(c)

	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}

	admin, err := srv.AdminLogin(&req)
	if err != nil {
		r.FailedResponse(err)
		return
	}

	tokendata := map[string]interface{}{
		"ID":       admin.ID,
		"Name":     admin.Adminname,
		"Nickname": admin.Nickname,
		"Role":     "admin",
	}
	// 设置token
	token, err := srv.Jwt.GenerateToken(tokendata)
	if err != nil {
		r.FailedResponse(err)
	} else {
		data := s.BuildAdminToken(admin, token)
		r.SuccessResponse(data, "登录成功")
	}
}

//AdminMe 用户详情
func AdminMe(c *gin.Context) {
	r := response.NewR(c)
	CurrentAdmin := CurrentAdmin(c)
	if CurrentAdmin == nil {
		r.FailedResponse(errors.New("获取信息失败"))
	}
	r.SuccessResponse(s.BuildAdmin(*CurrentAdmin), "用户详情")
}

//AdminInfoModify 用户修改资料
func AdminInfoModify(c *gin.Context) {
	var req request.AdminInfoModify
	r := response.NewR(c)

	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}
	cadmin := CurrentAdmin(c)
	updateAdmin := &model.Admin{}
	if cadmin != nil {
		updateAdmin.ID = req.ID
	} else {
		r.FailedResponse(errors.New("修改失败"))
		return
	}
	err := srv.AdminModify(updateAdmin, &req)
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(s.BuildAdmin(*updateAdmin), "修改成功")
	}
}

// // AdminLogout 用户登出
// func AdminLogout(c *gin.Context) {
// 	s := sessions.Default(c)
// 	s.Clear()
// 	s.Save()
// 	c.JSON(200, serializer.Response{
// 		Status: 0,
// 		Msg:    "登出成功",
// 	})
// }

// AdminGetAll 获取所有管理员
func AdminGetAll(c *gin.Context) {
	var req request.Page
	r := response.NewR(c)
	//参数绑定
	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}
	//获取数据
	res, err := srv.GetAdminsWithTotal(req.PageNum, req.PageSize)
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(res, "获取管理员成功")
	}
}

// AdminDelete 删除管理员
func AdminDelete(c *gin.Context) {
	r := response.NewR(c)
	ID, _ := strconv.ParseUint(c.Param("id"), 10, 0)

	ad := *CurrentAdmin(c)
	if ad.Adminname != "admin" || ad.ID == uint(ID) {
		r.FailedResponse(errors.New("禁止删除"))
		return
	}

	if err := srv.AdminDelete(uint(ID)); err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(nil, "删除管理员成功")
	}
}
