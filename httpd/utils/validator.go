package utils

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

// 全局翻译器T
var Trans ut.Translator

func InitTrans(locale string) (err error){
	//修改gin框架中Validator引擎属性，实现自定制
	if v,ok := binding.Validator.Engine().(*validator.Validate);ok{
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"),",",2)[0]
			if name == "-"{
				return ""
			}
			return name
		})
		zhT := zh.New()//中文翻译器
		enT := en.New()//英文翻译器
		// 第一个参数是备用(fallback)语言环境
		// 后面参数是应该支持语言环境(可支持多个)
		uni := ut.New(enT,zhT,enT)
		// locale通常取决于http请求'Accept-language'
		var ok bool
		Trans,ok = uni.GetTranslator(locale)
		if !ok{
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		// 注册翻译器
		switch locale{
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		}
		return
	}
	return
}

func RemoveStructName(fields map[string]string) map[string]string {
	result := map[string]string{}

	for field, err := range fields {
		result[field[strings.Index(field, ".")+1:]] = err
	}
	return result
}