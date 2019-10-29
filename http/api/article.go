package api

import (
	"blog/http/request"
	"blog/http/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ArticleAdd 添加文章
func ArticleAdd(c *gin.Context) {
	var req request.ArticleAdd
	r := response.NewR(c)

	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}

	err := srv.ArticleAdd(&req)
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(nil, "添加文章成功")
	}
}

// ArticleEdit 编辑文章
func ArticleEdit(c *gin.Context) {
	var req request.ArticleEdit
	r := response.NewR(c)

	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}

	err := srv.ArticleEdit(&req)
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(nil, "编辑文章成功")
	}
}

// ArticleGet 获取文章
func ArticleGet(c *gin.Context) {
	r := response.NewR(c)

	ID, _ := strconv.ParseUint(c.Param("id"), 10, 0)

	article, err := srv.ArticleGet(uint(ID))
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(article, "获取文章成功")
	}
}

// ArticleGetAll 获取所有文章
func ArticleGetAll(c *gin.Context) {
	var req request.Page
	r := response.NewR(c)
	//参数绑定
	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}
	//获取数据
	res, err := srv.UserGetArticles(req.PageNum, req.PageSize)
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(res, "获取文章成功")
	}
}

// ArticleDelete 删除文章
func ArticleDelete(c *gin.Context) {
	r := response.NewR(c)

	ID, _ := strconv.ParseUint(c.Param("id"), 10, 0)
	if err := srv.ArticleDelete(uint(ID)); err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(nil, "删除文章成功")
	}
}

// ArticlesGetWithTagID 获取所有文章
func ArticlesGetWithTagID(c *gin.Context) {
	var req request.ArticlesGetWithTagID
	r := response.NewR(c)
	//参数绑定
	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}
	//获取数据
	res, err := srv.UserGetArticlesWithTagID(req.PageNum, req.PageSize, req.TagID)
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(res, "获取文章成功")
	}
}

// ArticlesGetWithTopicID 获取所有文章
func ArticlesGetWithTopicID(c *gin.Context) {
	var req request.ArticlesGetWithTopicID
	r := response.NewR(c)
	//参数绑定
	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}
	//获取数据
	res, err := srv.UserGetArticlesWithTopicID(req.PageNum, req.PageSize, req.TopicID)
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(res, "获取文章成功")
	}
}

// GetArticlesbyArticleID 获取所有文章
func GetArticlesbyArticleID(c *gin.Context) {
	var req request.GetArticlesbyArticleID
	r := response.NewR(c)
	//参数绑定
	if err := c.ShouldBind(&req); err != nil {
		r.ErrorResponse(err)
		return
	}
	//获取数据
	res, err := srv.GetArticlesbyArticleID(req.ID)
	if err != nil {
		r.FailedResponse(err)
	} else {
		r.SuccessResponse(res, "获取文章成功")
	}
}
