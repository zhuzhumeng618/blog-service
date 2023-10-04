// Package service Tag 参数绑定和参数校验
package service

// CountTagRequest 验证 Tag 状态模型
type CountTagRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

// GetTagRequest 查询 Tag 行为模型
type GetTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

// CreateTagRequest 新增 Tag 行为模型
type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

// UpdateTagRequest 修改 Tag 行为模型
type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

// DeleteTagRequest 删除 Tag 行为模型
type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
