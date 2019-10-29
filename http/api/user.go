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

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var req request.UserRegister
	r := response.NewR(c)

	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}

	if ok := captchaVerify(req.CaptchaID, req.Captcha); !ok {
		r.FailedResponse(errors.New("验证码错误"))
	}

	user, err := srv.Register(&req)
	if err != nil {
		r.FailedResponse(err)
	} else {
		data := s.BuildUser(user)
		r.SuccessResponse(data, "注册成功")
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var req request.UserLogin
	r := response.NewR(c)

	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}

	user, err := srv.Login(&req)
	if err != nil {
		r.FailedResponse(err)
		return
	}

	tokendata := map[string]interface{}{
		"ID":       user.ID,
		"Name":     user.Username,
		"Nickname": user.Nickname,
		"Role":     "user",
	}
	// 设置token
	token, err := srv.Jwt.GenerateToken(tokendata)
	if err != nil {
		r.FailedResponse(err)
	} else {
		data := s.BuildUserToken(user, token)
		r.SuccessResponse(data, "登录成功")
	}
}

//UserMe 用户详情
func UserMe(c *gin.Context) {
	r := response.NewR(c)
	r.SuccessResponse(s.BuildUser(*CurrentUser(c)), "用户详情")
}

//UserInfoModify 用户修改资料
func UserInfoModify(c *gin.Context) {
	var req request.UserInfoModify
	r := response.NewR(c)

	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}
	cuser := CurrentUser(c)
	if cuser == nil {
		cuser = &model.User{}
		cadmin := CurrentAdmin(c)
		if cadmin != nil {
			cuser.ID = req.ID
		}
	}
	err := srv.UserModify(cuser, &req)
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(s.BuildUser(*cuser), "修改成功")
	}
}

// // UserLogout 用户登出
// func UserLogout(c *gin.Context) {
// 	s := sessions.Default(c)
// 	s.Clear()
// 	s.Save()
// 	c.JSON(200, serializer.Response{
// 		Status: 0,
// 		Msg:    "登出成功",
// 	})
// }

// GetDataByTime 一个需要token认证的测试接口
func GetDataByTime(c *gin.Context) {
	r := response.NewR(c)

	userInfo := CurrentUser(c)
	if userInfo != nil {
		r.SuccessResponse(userInfo, "登录成功")
	}
}

// UserGetAll 获取所有用户
func UserGetAll(c *gin.Context) {
	var req request.Page
	r := response.NewR(c)
	//参数绑定
	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}
	//获取数据
	res, err := srv.GetUsersWithTotal(req.PageNum, req.PageSize)
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(res, "获取用户成功")
	}
}

// UserDelete 删除文章
func UserDelete(c *gin.Context) {
	r := response.NewR(c)

	ID, _ := strconv.ParseUint(c.Param("id"), 10, 0)
	if err := srv.UserDelete(uint(ID)); err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(nil, "删除文章成功")
	}
}
