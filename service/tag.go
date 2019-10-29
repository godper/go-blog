package service

import (
	"blog/cache"
	"blog/http/request"
	"blog/model"
	"blog/serialize"
	"encoding/json"
	"errors"
	"time"
)

//UserGetArticlesWithTagID 用户获取标签文章列表服务
func (s *Service) UserGetArticlesWithTagID(PageNum int, PageSize int, TagID uint) (map[string]interface{}, error) {

	articles, err := s.getArticlesByOffsetWithTagID(PageNum, PageSize, TagID)
	if err != nil {
		return nil, err
	}
	counts, err := s.getArticleCountWithTagID(TagID)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"items": articles,
		"total": counts,
	}
	return res, nil
}

func (s *Service) getArticlesByOffsetWithTagID(PageNum int, PageSize int, TagID uint) ([]*serialize.ArticleWithTagID, error) {
	var articles []*serialize.ArticleWithTagID
	offset, limit := s.Page.OffsetLimit(PageNum, PageSize)

	//获取Redis缓存
	cachekey := cache.GenArticlesByOffsetWithTagID(offset, limit, TagID)
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &articles)
		return articles, nil
	}

	//数据库查询数据
	articles, err = s.Dao.GetArticlesByOffsetWithTagID(offset, limit, TagID)
	if err != nil {
		return nil, err
	}

	//Redis缓存数据
	s.Cache.Set(cachekey, articles, 5*time.Hour)
	return articles, nil
}

//ArticleCount 获取标签文章数量
func (s *Service) getArticleCountWithTagID(TagID uint) (int, error) {
	var count int
	//获取Redis缓存
	cachekey := cache.GenArticleCountWithTagIDKey(TagID)
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &count)
		return count, nil
	}
	//数据库查询数据
	if count, err = s.Dao.GetArticlesWithTagIDCount(TagID); err != nil {
		return 0, nil
	}
	//Redis缓存数据
	s.Cache.Set(cachekey, count, 5*time.Hour)
	return count, nil
}

//TagEdit 编辑标签服务
func (s *Service) TagEdit(req *request.TagEdit) error {
	tag := &model.Tag{
		Name: req.Name,
	}
	tag.ID = req.ID

	if err := s.Dao.EditTag(tag); err != nil {
		return errors.New("标签修改失败")
	}
	//清理缓存
	s.Cache.Delete(cache.GenTagsKey())
	return nil
}

//TagAdd 添加标签
func (s *Service) TagAdd(req *request.TagAdd) error {
	var tag model.Tag
	tag.Name = req.Name
	if err := s.Dao.AddTag(&tag); err != nil {
		//存入数据
		return err
	}
	s.Cache.Delete(cache.GenTagsKey())
	return nil
}

//TagDelete 删除标签
func (s *Service) TagDelete(ID uint) error {
	//清理缓存
	s.Cache.Delete(cache.GenTagsKey())
	return s.Dao.DeleteTag(ID)
}

//GetAllTags 获取所有标签
func (s *Service) GetAllTags() ([]*model.Tag, error) {
	var tags []*model.Tag
	//获取Redis缓存
	cachekey := cache.GenTagsKey()
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &tags)
		return tags, nil
	}

	tags, err = s.Dao.GetAllTags()
	if err != nil {
		return nil, err
	}
	//Redis缓存数据
	s.Cache.Set(cachekey, tags, 5*time.Hour)
	return tags, err
}
