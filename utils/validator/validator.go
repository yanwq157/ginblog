package validator

import (
	"fmt"
	"ginblog/utils/errmsg"
	"github.com/go-playground/locales/zh_Hans"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

//结构体是值传递 不能直接入参传，需要传一个空接口，不知道传的是什么类型的需要做断言，目前保证传的是结构体
//验证数据的方法
func Validate(data interface{}) (string, int) {
	validate := validator.New()
	//实例化翻译  转换为中文
	uni := unTrans.New(zh_Hans.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	//注册默认的翻译方法
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	//直接用验证结构体的方法验证传递的data 不知道传的什么需要判断，做断言
	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
