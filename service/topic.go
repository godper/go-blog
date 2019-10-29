package service

import (
	"blog/cache"
	"blog/http/request"
	"blog/model"
	"encoding/json"
	"errors"
	"time"
)

//getTopicIDByArticleID 获取文章主题ID
func (s *Service) getTopicIDByArticleID(ArticleID uint) (uint, error) {
	var TopicID uint
	//获取Redis缓存
	cachekey := cache.GengetTopicIDByArticleID(ArticleID)
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &TopicID)
		return TopicID, nil
	}
	//数据库查询数据
	article, err := s.Dao.GetTopicIDByArticleID(ArticleID)
	if err != nil {
		return 0, nil
	}
	//Redis缓存数据
	s.Cache.Set(cachekey, article.TopicID, 5*time.Hour)
	return article.TopicID, nil
}

//TopicEdit 编辑主题服务
func (s *Service) TopicEdit(req *request.TopicEdit) error {
	topic := &model.Topic{
		Name: req.Name,
	}
	topic.ID = req.ID

	if err := s.Dao.EditTopic(topic); err != nil {
		return errors.New("主题修改失败")
	}
	//清理缓存
	s.Cache.Delete(cache.GenTopicsKey())
	return nil
}

//TopicAdd 添加标签
func (s *Service) TopicAdd(req *request.TopicAdd) error {
	var topic model.Topic
	topic.Name = req.Name
	if err := s.Dao.AddTopic(&topic); err != nil {
		//存入数据
		return err
	}
	s.Cache.Delete(cache.GenTopicsKey())
	return nil
}

//GetAllTopics 获取所有主題
func (s *Service) GetAllTopics() ([]*model.Topic, error) {
	var topics []*model.Topic
	//获取Redis缓存
	cachekey := cache.GenTopicsKey()
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &topics)
		return topics, nil
	}

	topics, err = s.Dao.GetAllTopics()
	if err != nil {
		return nil, err
	}
	//Redis缓存数据
	s.Cache.Set(cachekey, topics, 5*time.Hour)
	return topics, err
}

//TopicDelete 删除主题
func (s *Service) TopicDelete(ID uint) error {
	//清理缓存
	s.Cache.Delete(cache.GenTopicsKey())
	return s.Dao.DeleteTopic(ID)
}
