package model

import (
	"blog-service/pkg/app"
)

// Tag 标签模型
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// TagSwagger 标签 Swagger 模型
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// TableName 生成数据库表名称
func (t *Tag) TableName() string {
	return "blog_tag"
}
