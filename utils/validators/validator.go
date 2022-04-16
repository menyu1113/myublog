package validators

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"myublog/global/myerrors"
	"reflect"
)

func Verification(data interface{}) (string, int) {
	validate := validator.New()  //生成验证器
	uni := unTrans.New(zh.New()) //生成一个带中文的翻译集
	//trans, _ := uni.GetTranslator(zh.New().Locale())					//返回给定语言环境的指定翻译器，
	trans, _ := uni.GetTranslator("zh")

	err := zhTrans.RegisterDefaultTranslations(validate, trans) //注册 为所有内置标签的验证器注册一组默认翻译
	if err != nil {
		fmt.Println("err:", err)
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), myerrors.ERRORCODE
		}
	}
	return "", myerrors.SUCCSECODE
}
