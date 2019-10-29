package response

import (
	"blog/conf"
	"blog/serialize"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

//R 回复
type R struct {
	c *gin.Context
}

//NewR New R 实例
func NewR(c *gin.Context) *R {
	return &R{
		c: c,
	}
}

const (
	//SuccessStatus 服务处理成功状态
	SuccessStatus = 20001
	//FailedStatus 服务处理状态
	FailedStatus = 40001
	//ErrorStatus 请求参数绑定失败
	ErrorStatus = 40004
	//TokenErrStatus token请求失败
	TokenErrStatus = 60001
)

//Resp 回复
func (r *R) Resp(status int, data interface{}, msg string) {
	var res = serialize.Response{
		Status: status,
		Data:   data,
		Msg:    msg,
	}
	if token, ok := r.c.Get("newtoken"); ok {
		if token, ok := token.(string); ok {
			res := serialize.BuildTokenRespon(&res, token)
			r.c.JSON(http.StatusOK, res)
			return
		}
	}
	r.c.JSON(http.StatusOK, res)
}

//SuccessResponse 请求成功
func (r *R) SuccessResponse(data interface{}, msg string) {
	r.Resp(SuccessStatus, data, msg)
}

//FailedResponse 请求失败
func (r *R) FailedResponse(err error) {
	r.Resp(FailedStatus, nil, err.Error())
}

// ErrorResponse 返回错误消息
func (r *R) ErrorResponse(err error) {
	r.Resp(ErrorStatus, nil, resolveError(err))
}

// TokenErrResponse token无效
func (r *R) TokenErrResponse(err error) {
	r.Resp(TokenErrStatus, nil, err.Error())
}

func resolveError(err error) string {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return fmt.Sprintf("%s%s", field, tag)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return "JSON类型不匹配"

	}
	return "参数错误"
}
