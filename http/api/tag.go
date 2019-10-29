package api

import (
	"blog/http/request"
	"blog/http/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

//TagsGetAll 获取所有标签
func TagsGetAll(c *gin.Context) {
	r := response.NewR(c)

	res, err := srv.GetAllTags()
	if err != nil {
		r.FailedResponse(err)
	}
	r.SuccessResponse(res, "获取所有标签")

}

// TagDelete 删除标签
func TagDelete(c *gin.Context) {
	r := response.NewR(c)

	ID, _ := strconv.ParseUint(c.Param("id"), 10, 0)
	if err := srv.TagDelete(uint(ID)); err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(nil, "删除标签成功")
	}
}

// TagAdd 添加标签
func TagAdd(c *gin.Context) {
	var req request.TagAdd
	r := response.NewR(c)

	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}

	err := srv.TagAdd(&req)
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(nil, "添加标签成功")
	}
}

// TagEdit 编辑主题
func TagEdit(c *gin.Context) {
	var req request.TagEdit
	r := response.NewR(c)

	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}

	err := srv.TagEdit(&req)
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(nil, "编辑标签成功")
	}
}
