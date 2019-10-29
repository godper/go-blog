package dao

import "blog/model"

//GetTopicIDByArticleID 获取文章标签ID
func (d *Dao) GetTopicIDByArticleID(articleID uint) (*model.Article, error) {
	var article model.Article
	err := d.DB.Where("id = ?", articleID).Select("id, topic_id").First(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

//GetAllTopics 获取所有主題
func (d *Dao) GetAllTopics() ([]*model.Topic, error) {
	var topics []*model.Topic

	if err := d.DB.Select("id,name").Find(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil

}

// DeleteTopic delete a single Topic
func (d *Dao) DeleteTopic(id uint) error {
	if err := d.DB.Where("id = ?", id).Delete(model.Topic{}).Error; err != nil {
		return err
	}
	return nil
}

// EditTopic modify a single topic
func (d *Dao) EditTopic(topic *model.Topic) error {
	if err := d.DB.Model(&model.Topic{}).Updates(topic).Error; err != nil {
		return err
	}
	return nil
}

// AddTopic add a single topic
func (d *Dao) AddTopic(topic *model.Topic) error {
	if err := d.DB.Create(&topic).Error; err != nil {
		return err
	}
	return nil
}
