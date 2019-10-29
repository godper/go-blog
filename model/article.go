package model

//Article 文章模型
type Article struct {
	Model
	Title         string `json:"title"`
	Preview       string `json:"preview,omitempty"`
	Content       string `json:"content,omitempty"`
	CoverImageURL string `json:"cover_image_url"`
	State         int    `json:"state"`
	TopicID       uint   `json:"topic_id" gorm:"index"`
	Topic         Topic  `json:"topic"`
	Tags          []Tag  `json:"tags" gorm:"many2many:article_tags;"`
}

//Topic 文章主题
type Topic struct {
	Model
	Articles []Article `json:"articles,omitempty"`
	Name     string    `json:"name,omitempty"`
	State    int       `json:"state,omitempty"`
}

// Tag 标签
type Tag struct {
	Model
	Name  string `json:"name,omitempty"`
	State int    `json:"state,omitempty"`
}

//ArticleTag 文章标签映射
type ArticleTag struct {
	Model
	ArticleID uint `gorm:"index" json:"article_id,omitempty"`
	TagID     uint `gorm:"index" json:"tag_id,omitempty"`
}
