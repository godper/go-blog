package cache

import "strconv"

//GenArticleKey 生成单个文章缓存key
func GenArticleKey(ID uint) string {
	return "blog_articles_" + strconv.FormatUint(uint64(ID), 10)
}

//GenArticleCountKey 生成文章数量缓存key
func GenArticleCountKey() string {
	return "blog_articles_count"
}

//GenArticlesByOffset 生成文章列表缓存key
func GenArticlesByOffset(offset int, limit int) string {
	return "blog_articles_by_offset" + strconv.Itoa(offset) + "_limit" + strconv.Itoa(limit)
}

//GenTagsKey 获取tagskey
func GenTagsKey() string {
	return "blog_tags"
}

//GenTopicsKey 获取topicskey
func GenTopicsKey() string {
	return "blog_topics"
}

//GenAdminsByOffset 生成管理员列表缓存key
func GenAdminsByOffset(offset int, limit int) string {
	return "blog_admins_by_offset" + strconv.Itoa(offset) + "_limit" + strconv.Itoa(limit)
}

//GenAdminCountKey 生成管理员数量缓存key
func GenAdminCountKey() string {
	return "blog_admin_count"
}

//GenUsersByOffset 生成用户列表缓存key
func GenUsersByOffset(offset int, limit int) string {
	return "blog_users_by_offset" + strconv.Itoa(offset) + "_limit" + strconv.Itoa(limit)
}

//GenUserCountKey 生成用户数量缓存key
func GenUserCountKey() string {
	return "blog_user_count"
}

//GenArticlesByOffsetWithTopicID 生成主题文章列表缓存key
func GenArticlesByOffsetWithTopicID(offset int, limit int, TopicID uint) string {
	return "blog_articles_by_offset" + strconv.Itoa(offset) +
		"_limit" + strconv.Itoa(limit) +
		"_topicid" + strconv.FormatUint(uint64(TopicID), 10)
}

//GenArticleCountWithTopicIDKey 生成主题文章数量缓存key
func GenArticleCountWithTopicIDKey(TopicID uint) string {
	return "blog_articles_count_with_topicid_" + strconv.FormatUint(uint64(TopicID), 10)
}

//GenArticlesByOffsetWithTagID 生成标签文章列表缓存key
func GenArticlesByOffsetWithTagID(offset int, limit int, TagID uint) string {
	return "blog_articles_by_offset" + strconv.Itoa(offset) +
		"_limit" + strconv.Itoa(limit) +
		"_tagid" + strconv.FormatUint(uint64(TagID), 10)
}

//GenArticleCountWithTagIDKey 生成标签文章数量缓存key
func GenArticleCountWithTagIDKey(TagID uint) string {
	return "blog_articles_count_with_tagid_" + strconv.FormatUint(uint64(TagID), 10)
}

//GengetTopicIDByArticleID rest
func GengetTopicIDByArticleID(ArticleID uint) string {
	return "blog_articles_TopicID_ByArticleID_" + strconv.FormatUint(uint64(ArticleID), 10)
}

//GengetArticlesbyArticleID erwf
func GengetArticlesbyArticleID(TopicID uint) string {
	return "blog_articles_ByArticleID_" + strconv.FormatUint(uint64(TopicID), 10)
}
