package model

// ArticleTag 文章标签关联模型
type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

// TableName 生成数据库表名称
func (t *ArticleTag) TableName() string {
	return "blog_article_tag"
}
