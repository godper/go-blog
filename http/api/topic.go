package api

import (
	"blog/http/request"
	"blog/http/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

//TopicsGetAll 获取所有主題
func TopicsGetAll(c *gin.Context) {
	r := response.NewR(c)

	res, err := srv.GetAllTopics()
	if err != nil {
		r.FailedResponse(err)
	}
	r.SuccessResponse(res, "获取所有主題")

}

// TopicDelete 删除主题
func TopicDelete(c *gin.Context) {
	r := response.NewR(c)

	ID, _ := strconv.ParseUint(c.Param("id"), 10, 0)
	if err := srv.TopicDelete(uint(ID)); err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(nil, "删除主题成功")
	}
}

// TopicAdd 添加主题
func TopicAdd(c *gin.Context) {
	var req request.TopicAdd
	r := response.NewR(c)

	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}

	err := srv.TopicAdd(&req)
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(nil, "添加主题成功")
	}
}

// TopicEdit 编辑主题
func TopicEdit(c *gin.Context) {
	var req request.TopicEdit
	r := response.NewR(c)

	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}

	err := srv.TopicEdit(&req)
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(nil, "编辑主题成功")
	}
}
