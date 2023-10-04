package routers

import (
	_ "blog-service/docs"
	"blog-service/internal/middleware"
	v1 "blog-service/internal/routers/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter 创建 gin 实例
func NewRouter() *gin.Engine {
	// 使用自定义的方式创建一个 gin 实例
	r := gin.New()

	// 将 gin 的日志中间件和异常中间件注册进 gin 实例中
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.Translations(), // 国际化语言翻译中间件
	)

	// 注册 Swagger API
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 使用路由分组保证良好的分层结构，便于管理
	apiv1 := r.Group("/api/v1")
	{
		// TAG API
		tag := v1.NewTag()
		apiv1.POST("/tags", tag.Create)            // 新增标签
		apiv1.DELETE("/tags/:id", tag.Delete)      // 删除指定标签
		apiv1.PUT("/tags/:id", tag.Update)         // 更新指定标签
		apiv1.PATCH("/tags/:id/state", tag.Update) // 修改标签状态
		apiv1.GET("/tags/:id", tag.Get)            // 获取指定标签
		apiv1.GET("/tags", tag.List)               // 获取全部标签

		// Article API
		article := v1.NewArticle()
		apiv1.POST("/articles", article.Create)            // 新增文章
		apiv1.DELETE("/articles/:id", article.Delete)      // 删除指定文章
		apiv1.PUT("/articles/:id", article.Update)         // 更新指定文章
		apiv1.PATCH("/articles/:id/state", article.Update) // 更新文章状态
		apiv1.GET("/articles/:id", article.Get)            // 获取指定文章
		apiv1.GET("/articles", article.List)               // 获取全部文章
	}

	// 返回 gin 实例
	return r
}
