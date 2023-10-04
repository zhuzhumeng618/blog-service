// Package app 全局响应模型
package app

import (
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 统一响应模型
type Response struct {
	Ctx *gin.Context
}

// Pager 分页模型
type Pager struct {
	Page      int `json:"page"`       // 当前页
	PageSize  int `json:"page_size"`  // 每页记录条数
	TotalRows int `json:"total_rows"` // 总计记录行数
}

// NewResponse 创建统一响应实例
func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

// ToResponse 发送响应，普通形式
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	// 使用 gin 的 JSON 形式返回数据
	r.Ctx.JSON(
		http.StatusOK, // 状态码
		data,          // 数据
	)
}

// ToResponseList 发送响应，列表形式
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(
		http.StatusOK,
		gin.H{
			"list": list, // 当前页的数据
			"pager": &Pager{ // 返回分页实例
				Page:      GetPage(r.Ctx),     // 当前页
				PageSize:  GetPageSize(r.Ctx), // 当前页记录条数
				TotalRows: totalRows,          // 总记录条数
			},
		},
	)
}

// ToErrorResponse 发送响应，错误形式
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code": err.GetCode(),    // 错误码，自定义
		"msg":  err.GetMessage(), // 错误原因
	}

	details := err.GetDetails() // 错误详情
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(
		err.StatusCode(), // 错误码对应的 HTTP 状态码
		response,         // 错误完整信息
	)
}
