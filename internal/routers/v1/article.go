package v1

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// Article 文章方法接受者
type Article struct {
}

// NewArticle 创建文章实例
func NewArticle() *Article {
	return &Article{}
}

// Get 获取指定文章
func (a *Article) Get(ctx *gin.Context) {
	app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
	return
}

// List 获取所有文章
func (a *Article) List(ctx *gin.Context) {}

// Create 创建文章
func (a *Article) Create(ctx *gin.Context) {}

// Update 修改文章
func (a *Article) Update(ctx *gin.Context) {}

// Delete 删除文章
func (a *Article) Delete(ctx *gin.Context) {}
