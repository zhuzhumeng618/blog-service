package app

import (
	"blog-service/global"
	"blog-service/pkg/convert"
	"github.com/gin-gonic/gin"
)

// GetPage 获取当前页
func GetPage(ctx *gin.Context) int {
	// 获取 querystring 参数，并通过 convert 包函数将数据转换为 Int 类型数据
	page := convert.StrTo(ctx.Query("page")).MustInt()
	// 当 page 为负值，则自动修正为第一页，也就是起始页
	if page <= 0 {
		return 1
	}

	// 数据没问题，则直接返回本身
	return page
}

// GetPageSize 获取每页记录条数
func GetPageSize(ctx *gin.Context) int {
	pageSize := convert.StrTo(ctx.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.App.DefaultPageSize // 100 条
	}

	return pageSize
}

// GetPageOffset 获取偏移量
func GetPageOffset(page, pageSize int) int {
	// 初始化 result
	result := 0
	/*
		每页十条
								page pageSize result
		第 1 页：limit 0, 10      1     10     (1-1)*10
		第 2 页：limit 10, 10     2     10     (2-1)*10
		第 3 页：limit 20, 10     3     10     (3-1)*10
		第 n 页：limit n, 10      n     10     (n-1)*10
	*/
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
