package model

import (
	"blog-service/pkg/app"
)

// Article 文章模型
type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

// ArticleSwagger 标签 Swagger 模型
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

// TableName 生成数据库表名称
func (t *Article) TableName() string {
	return "blog_article"
}
