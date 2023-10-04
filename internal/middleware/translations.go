// Package middleware 信息国际化处理
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ut 为通用翻译器
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := ctx.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				// 将验证器和对应语言类型的 Translator 注册进来
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
			}

			// 将 Translator 存储到全局上下文中，方便后续翻译时使用
			ctx.Set("trans", trans)
		}

		// 继续下个处理函数
		ctx.Next()
	}
}
