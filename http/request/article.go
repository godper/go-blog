package request

import "blog/model"

// ArticleAdd 添加文章表单
type ArticleAdd struct {
	Topic         *model.Topic `form:"topic"           json:"topic"`
	Tags          []*model.Tag `form:"tags"           json:"tags"`
	Title         string       `form:"title"           json:"title"           binding:"required,min=5,max=50"`
	Preview       string       `form:"preview"            json:"preview"      binding:"required,min=5,max=200"`
	Content       string       `form:"content"         json:"content"         binding:"required,min=20"`
	CoverImageURL string       `form:"cover_image_url" json:"cover_image_url" binding:"omitempty"`
	State         int          `form:"state"           json:"state"           binding:"omitempty"`
}

// ArticleEdit 编辑文章表单
type ArticleEdit struct {
	ID            uint         `form:"id"              json:"id"              binding:"required"`
	Topic         *model.Topic `form:"topic"           json:"topic"`
	Tags          []*model.Tag `form:"tags"           json:"tags"`
	TopicID       uint         `form:"topic_id"        json:"topic_id" `
	Title         string       `form:"title"           json:"title"           binding:"required,min=5,max=50"`
	Preview       string       `form:"preview"            json:"preview"      binding:"required,min=5,max=100"`
	Content       string       `form:"content"         json:"content"         binding:"required,min=20"`
	CoverImageURL string       `form:"cover_image_url" json:"cover_image_url" binding:"omitempty"`
	State         int          `form:"state"           json:"state"           binding:"omitempty"`
}

// ArticleDelete 删除文章表单
type ArticleDelete struct {
	ID uint `form:"id"              json:"id"              binding:"required"`
}

// ArticleGet 获取文章表单
type ArticleGet struct {
	ID uint `form:"id"              json:"id"              binding:"required"`
}

// Page 页码
type Page struct {
	PageNum  int `form:"page_num"              json:"page_num" `
	PageSize int `form:"page_size"              json:"page_size"`
}

// TagAdd 添加标签表单
type TagAdd struct {
	Name string `form:"name"            json:"name"      binding:"required,min=1"`
}

// TagEdit 编辑标签表单
type TagEdit struct {
	ID   uint   `form:"id"              json:"id"              binding:"required"`
	Name string `form:"name"            json:"name"      binding:"required,min=1"`
}

// TopicAdd 添加标签表单
type TopicAdd struct {
	Name string `form:"name"            json:"name"      binding:"required,min=1"`
}

// TopicEdit 编辑标签表单
type TopicEdit struct {
	ID   uint   `form:"id"              json:"id"              binding:"required"`
	Name string `form:"name"            json:"name"      binding:"required,min=1"`
}

// ArticlesGetWithTopicID 页码
type ArticlesGetWithTopicID struct {
	PageNum  int  `form:"page_num"              json:"page_num" `
	PageSize int  `form:"page_size"              json:"page_size"`
	TopicID  uint `form:"topic_id"        json:"topic_id" binding:"required,min=1"`
}

// GetArticlesbyArticleID we
type GetArticlesbyArticleID struct {
	ID uint `form:"id"              json:"id"              binding:"required"`
}

// ArticlesGetWithTagID 页码
type ArticlesGetWithTagID struct {
	PageNum  int  `form:"page_num"              json:"page_num" `
	PageSize int  `form:"page_size"              json:"page_size"`
	TagID    uint `form:"tag_id"        json:"tag_id" binding:"required,min=1"`
}
