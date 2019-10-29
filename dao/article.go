package dao

import (
	"blog/model"
	"blog/serialize"
	"blog/util/godper"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

// ExistArticleByID checks if an article exists based on ID
func (d *Dao) ExistArticleByID(id int) (bool, error) {
	var article model.Article
	err := d.DB.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetArticleCount gets the total number of articles based on the constraints
func (d *Dao) GetArticleCount() (int, error) {
	var count int
	if err := d.DB.Model(&model.Article{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

//GetArticleCountWithMaps gets the total number of articles based on the constraints
func (d *Dao) GetArticleCountWithMaps(maps interface{}) (int, error) {
	var count int
	if err := d.DB.Model(&model.Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetArticlesByOffset gets a list of articles based on paging constraints
func (d *Dao) GetArticlesByOffset(offset int, limit int) ([]*model.Article, error) {
	var articles []*model.Article

	err := d.DB.Select("id, title, preview, cover_image_url, created_at, topic_id").
		Preload("Topic", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Select("tags.id, tags.name, article_tags.article_id")
		}).
		Offset(offset).
		Limit(limit).
		Order("id desc").
		Find(&articles).
		Error

	if err != nil {
		return nil, err
	}
	for _, article := range articles {
		godper.Timetranfer(article)
	}
	return articles, nil
}

// GetArticlesByOffsetWithTopicID gets a list of articles based on paging constraints
func (d *Dao) GetArticlesByOffsetWithTopicID(offset int, limit int, TopicID uint) ([]*serialize.ArticleWithTopicID, error) {
	var articles []*serialize.ArticleWithTopicID

	err := d.DB.
		Select("a.id, a.title, a.preview, a.cover_image_url, a.created_at, a.topic_id").
		Table("articles a").
		Joins("left join topics t on a.topic_id = t.id").
		Where("t.deleted_at is NULL AND a.deleted_at is NULL AND t.id = ?", TopicID).
		Offset(offset).
		Limit(limit).
		Order("a.id desc").
		Find(&articles).
		Error

	if err != nil {
		return nil, err
	}
	for _, article := range articles {
		godper.Timetranfer(article)
	}
	return articles, nil
}

// GetArticlesWithTopicIDCount gets the total number of articles based on the constraints
func (d *Dao) GetArticlesWithTopicIDCount(TopicID uint) (int, error) {
	var count int

	err := d.DB.
		Select("a.id, a.title, a.preview, a.cover_image_url, a.created_at, a.topic_id").
		Table("articles a").
		Joins("left join topics t on a.topic_id = t.id").
		Where("t.deleted_at is NULL AND a.deleted_at is NULL AND t.id = ?", TopicID).
		Count(&count).
		Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetArticle Get a single article based on ID
func (d *Dao) GetArticle(id uint) (*model.Article, error) {
	var article model.Article

	err := d.DB.Where("id = ?", id).
		Select("id, title, preview, content, cover_image_url, created_at, topic_id").
		Preload("Topic", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Select("tags.id, tags.name, article_tags.article_id")
		}).
		First(&article).
		Error
	if err != nil {
		return nil, err
	}

	godper.Timetranfer(&article)

	return &article, nil

}

// EditArticle modify a single article
func (d *Dao) EditArticle(article *model.Article) error {
	if err := d.DB.Model(&model.Article{}).Updates(article).Error; err != nil {
		return err
	}
	return nil
}

// AddArticle add a single article
func (d *Dao) AddArticle(article *model.Article) error {
	if err := d.DB.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

// DeleteArticle delete a single article
func (d *Dao) DeleteArticle(id uint) error {
	if err := d.DB.Where("id = ?", id).Delete(model.Article{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllArticle clear all article
func (d *Dao) CleanAllArticle() error {
	if err := d.DB.Unscoped().Where("deleted_at != ? ", nil).Delete(&model.Article{}).Error; err != nil {
		return err
	}
	return nil
}

//AddArticleTagRelate 添加文章标签
func (d *Dao) AddArticleTagRelate(articleID uint, tags []*model.Tag) error {
	articletags := "INSERT INTO article_tags (article_id,tag_id) VALUES "
	for _, tag := range tags {
		articletags += `(` + strconv.Itoa(int(articleID)) + `,` + strconv.Itoa(int(tag.ID)) + `),`
	}
	articletags = strings.Trim(articletags, ",")
	err := d.DB.Exec(articletags).Error
	if err != nil {
		return err
	}
	return nil
}

// GetArticlesbyTopicID gets a list of articles based on paging constraints
func (d *Dao) GetArticlesbyTopicID(TopicID uint) ([]*serialize.ArticlesbyTopicID, error) {
	var articles []*serialize.ArticlesbyTopicID
	err := d.DB.
		Select("id, title").
		Table("articles").
		Where("topic_id = ? AND deleted_at is NULL ", TopicID).
		Order("id desc").
		Find(&articles).
		Error

	if err != nil {
		return nil, err
	}
	return articles, nil
}

// GetArticlesByOffsetWithTagID gets a list of articles based on paging constraints
func (d *Dao) GetArticlesByOffsetWithTagID(offset int, limit int, TagID uint) ([]*serialize.ArticleWithTagID, error) {
	var articles []*serialize.ArticleWithTagID

	err := d.DB.
		Select("a.id, a.title, a.preview, a.cover_image_url, a.created_at, at.tag_id").
		Table("articles a").
		Joins("left join article_tags at on a.id = at.article_id").
		Where("at.deleted_at is NULL AND a.deleted_at is NULL  AND  at.tag_id = ?", TagID).
		Offset(offset).
		Limit(limit).
		Order("a.id desc").
		Find(&articles).
		Error

	if err != nil {
		return nil, err
	}
	for _, article := range articles {
		godper.Timetranfer(article)
	}
	return articles, nil
}

// GetArticlesWithTagIDCount gets the total number of articles based on the constraints
func (d *Dao) GetArticlesWithTagIDCount(TagID uint) (int, error) {
	var count int

	err := d.DB.
		Select("a.id, a.title, a.preview, a.cover_image_url, a.created_at, at.tag_id").
		Table("articles a").
		Joins("left join article_tags at on a.id = at.article_id").
		Where("at.deleted_at is NULL AND a.deleted_at is NULL  AND at.tag_id = ?", TagID).
		Count(&count).
		Error

	if err != nil {
		return 0, err
	}
	return count, nil
}
