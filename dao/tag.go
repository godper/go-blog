package dao

import "blog/model"

//DeleteArticleTags 删除文章标签
func (d *Dao) DeleteArticleTags(maps interface{}) error {
	if err := d.DB.Unscoped().Where(maps).Delete(&model.ArticleTag{}).Error; err != nil {
		return err
	}
	return nil
}

//GetTagIDByArticleID 获取文章标签ID
func (d *Dao) GetTagIDByArticleID(articleID uint) ([]*model.ArticleTag, error) {
	var articletags []*model.ArticleTag
	err := d.DB.Where("article_id = ?", articleID).Select("id, article_id, tag_id").Find(&articletags).Error
	if err != nil {
		return nil, err
	}
	return articletags, nil
}

// DeleteTag delete a single tag
func (d *Dao) DeleteTag(id uint) error {
	if err := d.DB.Where("id = ?", id).Delete(model.Tag{}).Error; err != nil {
		return err
	}
	return nil
}

// AddTag add a single tag
func (d *Dao) AddTag(tag *model.Tag) error {
	if err := d.DB.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

// EditTag modify a single tag
func (d *Dao) EditTag(tag *model.Tag) error {
	if err := d.DB.Model(&model.Tag{}).Updates(tag).Error; err != nil {
		return err
	}
	return nil
}

//GetAllTags 获取所有标签
func (d *Dao) GetAllTags() ([]*model.Tag, error) {
	var tags []*model.Tag

	if err := d.DB.Select("id,name").Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil

}
