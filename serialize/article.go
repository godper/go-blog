package serialize

import (
	"blog/model"
)

//Article 文章模型
type Article struct {
	ID            uint   `json:"id"`
	TopicID       uint   `json:"topic_id" gorm:"index"`
	Title         string `json:"title"`
	Preview       string `json:"preview"`
	Content       string `json:"content"`
	CoverImageURL string `json:"cover_image_url"`
	State         int    `json:"state"`
	CreatedAt     string `json:"created_at"`
}

//ArticleWithTopicID 主题文章
type ArticleWithTopicID struct {
	ID            uint   `json:"id"`
	TopicID       uint   `json:"topic_id" gorm:"index"`
	Title         string `json:"title"`
	Preview       string `json:"preview"`
	Content       string `json:"content"`
	CoverImageURL string `json:"cover_image_url"`
	CreatedAt     string `json:"created_at" godper:"2006-01-02"`
}

//ArticlesbyTopicID 主题文章
type ArticlesbyTopicID struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

//ArticleWithTagID 标签文章
type ArticleWithTagID struct {
	ID            uint   `json:"id"`
	TagID         uint   `json:"tag_id" gorm:"index"`
	Title         string `json:"title"`
	Preview       string `json:"preview"`
	Content       string `json:"content"`
	CoverImageURL string `json:"cover_image_url"`
	CreatedAt     string `json:"created_at" godper:"2006-01-02"`
}

// BuildArticle 序列化文章
func BuildArticle(article *model.Article) *Article {
	return &Article{
		ID:            article.ID,
		TopicID:       article.TopicID,
		Title:         article.Title,
		Preview:       article.Preview,
		Content:       article.Content,
		CoverImageURL: article.CoverImageURL,
		State:         article.State,
		CreatedAt:     article.CreatedAt,
	}
}
