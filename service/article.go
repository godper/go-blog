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

//ArticleAdd 添加文章服务
func (s *Service) ArticleAdd(req *request.ArticleAdd) error {
	article := model.Article{
		Title:         req.Title,
		Preview:       req.Preview,
		Content:       req.Content,
		CoverImageURL: req.CoverImageURL,
		State:         req.State,
	}
	if req.Topic != nil {
		//添加主题ID
		article.TopicID = req.Topic.ID
	}
	if err := s.Dao.AddArticle(&article); err != nil {
		//存入数据
		return err
	}
	//清理缓存
	s.Cache.DeleteLike("blog_articles*")

	if len(req.Tags) != 0 {
		//tags 关系
		if err := s.Dao.AddArticleTagRelate(article.ID, req.Tags); err != nil {
			return err
		}
	}
	return nil
}

//ArticleEdit 编辑文章服务
func (s *Service) ArticleEdit(req *request.ArticleEdit) error {
	article := &model.Article{
		Title:         req.Title,
		Preview:       req.Preview,
		Content:       req.Content,
		CoverImageURL: req.CoverImageURL,
		State:         req.State,
	}
	article.ID = req.ID
	//主题
	if req.Topic != nil {
		if req.Topic.ID != req.TopicID {
			article.TopicID = req.Topic.ID
		}
	}

	if err := s.Dao.EditArticle(article); err != nil {
		return errors.New("文章修改失败")
	}
	//清理缓存
	s.Cache.DeleteLike("blog_articles*")

	//标签
	articletags, _ := s.Dao.GetTagIDByArticleID(article.ID)
	if len(articletags) > 0 {
		if err := s.Dao.DeleteArticleTags(map[string]interface{}{"article_id": article.ID}); err != nil {
			return errors.New("标签更新失败")
		}
	}
	if req.Tags != nil {
		if err := s.Dao.AddArticleTagRelate(article.ID, req.Tags); err != nil {
			return errors.New("标签更新失败")
		}
	}

	return nil
}

//ArticleGet 获取单个文章服务
func (s *Service) ArticleGet(ID uint) (*model.Article, error) {
	var article *model.Article

	cachekey := cache.GenArticleKey(ID)
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &article)
		return article, nil
	}

	article, err = s.Dao.GetArticle(ID)
	if err != nil {
		return nil, err
	}

	s.Cache.Set(cachekey, article, 5*time.Hour)
	return article, nil
}

//getArticlesByOffset 分页文章列表
func (s *Service) getArticlesByOffset(PageNum int, PageSize int) ([]*model.Article, error) {
	var articles []*model.Article
	offset, limit := s.Page.OffsetLimit(PageNum, PageSize)

	//获取Redis缓存
	cachekey := cache.GenArticlesByOffset(offset, limit)
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &articles)
		return articles, nil
	}

	//数据库查询数据
	articles, err = s.Dao.GetArticlesByOffset(offset, limit)
	if err != nil {
		return nil, err
	}

	//Redis缓存数据
	s.Cache.Set(cachekey, articles, 5*time.Hour)
	return articles, nil
}

//ArticleExistByID 查询文章byID
func (s *Service) ArticleExistByID(ID int) (bool, error) {
	return s.Dao.ExistArticleByID(ID)
}

//ArticleCount 获取文章数量
func (s *Service) getArticleCount() (int, error) {
	var count int
	//获取Redis缓存
	cachekey := cache.GenArticleCountKey()
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &count)
		return count, nil
	}
	//数据库查询数据
	if count, err = s.Dao.GetArticleCount(); err != nil {
		return 0, nil
	}
	//Redis缓存数据
	s.Cache.Set(cachekey, count, 5*time.Hour)
	return count, nil
}

//getArticleCountWithMaps 获取文章数量 带条件查询
func (s *Service) getArticleCountWithMaps(maps interface{}) (int, error) {
	return s.Dao.GetArticleCountWithMaps(maps)
}

//ArticleDelete 删除文章
func (s *Service) ArticleDelete(ID uint) error {
	//清理缓存
	s.Cache.DeleteLike("blog_articles*")

	return s.Dao.DeleteArticle(ID)
}

//UserGetArticles 用户获取文章列表服务
func (s *Service) UserGetArticles(PageNum int, PageSize int) (map[string]interface{}, error) {

	articles, err := s.getArticlesByOffset(PageNum, PageSize)
	if err != nil {
		return nil, err
	}
	counts, err := s.getArticleCount()
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"items": articles,
		"total": counts,
	}
	return res, nil
}

//UserGetArticlesWithTopicID 用户获取主题文章列表服务
func (s *Service) UserGetArticlesWithTopicID(PageNum int, PageSize int, TopicID uint) (map[string]interface{}, error) {

	articles, err := s.getArticlesByOffsetWithTopicID(PageNum, PageSize, TopicID)
	if err != nil {
		return nil, err
	}
	counts, err := s.getArticleCountWithTopicID(TopicID)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"items": articles,
		"total": counts,
	}
	return res, nil
}

func (s *Service) getArticlesByOffsetWithTopicID(PageNum int, PageSize int, TopicID uint) ([]*serialize.ArticleWithTopicID, error) {
	var articles []*serialize.ArticleWithTopicID
	offset, limit := s.Page.OffsetLimit(PageNum, PageSize)

	//获取Redis缓存
	cachekey := cache.GenArticlesByOffsetWithTopicID(offset, limit, TopicID)
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &articles)
		return articles, nil
	}

	//数据库查询数据
	articles, err = s.Dao.GetArticlesByOffsetWithTopicID(offset, limit, TopicID)
	if err != nil {
		return nil, err
	}

	//Redis缓存数据
	s.Cache.Set(cachekey, articles, 5*time.Hour)
	return articles, nil
}

//ArticleCount 获取主题文章数量
func (s *Service) getArticleCountWithTopicID(TopicID uint) (int, error) {
	var count int
	//获取Redis缓存
	cachekey := cache.GenArticleCountWithTopicIDKey(TopicID)
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &count)
		return count, nil
	}
	//数据库查询数据
	if count, err = s.Dao.GetArticlesWithTopicIDCount(TopicID); err != nil {
		return 0, nil
	}
	//Redis缓存数据
	s.Cache.Set(cachekey, count, 5*time.Hour)
	return count, nil
}

//GetArticlesbyArticleID 获取同主题文章列表
func (s *Service) GetArticlesbyArticleID(ArticleID uint) ([]*serialize.ArticlesbyTopicID, error) {
	TopicID, err := s.getTopicIDByArticleID(ArticleID)
	if err != nil {
		return nil, err
	}
	if TopicID == 0 {
		return nil, nil
	}

	var articles []*serialize.ArticlesbyTopicID
	//获取Redis缓存
	cachekey := cache.GengetArticlesbyArticleID(TopicID)
	val, err := s.Cache.Get(cachekey)
	if err == nil {
		json.Unmarshal(val, &articles)
		return articles, nil
	}

	articles, err = s.Dao.GetArticlesbyTopicID(TopicID)
	if err != nil {
		return nil, err
	}

	s.Cache.Set(cachekey, articles, 5*time.Hour)
	return articles, nil
}
