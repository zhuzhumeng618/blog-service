package v1

import (
	"github.com/gin-gonic/gin"
)

// Tag 标签方法接受者
type Tag struct {
}

// NewTag 创建标签实例
func NewTag() *Tag {
	return &Tag{}
}

// Get 获取指定标签
func (t *Tag) Get(ctx *gin.Context) {}

// List 获取所有标签
// @Summary 获取所有标签
// @Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0,1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t *Tag) List(ctx *gin.Context) {}

// Create 创建标签
// @Summary 创建标签
// @Produce json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t *Tag) Create(ctx *gin.Context) {}

// Update 更新标签
// @Summary 更新标签
// @Produce json
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" maxlength(100)
// @Param state body int false "状态" Enums(0,1) default(1)
// @Param modified_by body string false "修改者" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t *Tag) Update(ctx *gin.Context) {}

// Delete 删除标签
// @Summary 删除标签
// @Produce json
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t *Tag) Delete(ctx *gin.Context) {}
