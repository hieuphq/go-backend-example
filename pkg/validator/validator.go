package validator

import (
	"reflect"
	"strings"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/hieuphq/backend-example/translation"
)

// NewStructValidator make default validator
func NewStructValidator(h translation.Helper) binding.StructValidator {
	return &defaultValidator{
		transHelper: h,
	}
}

type defaultValidator struct {
	once        sync.Once
	validate    *validator.Validate
	transHelper translation.Helper
}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		v.transHelper.InitErrorTranslator(v.validate)

		// add any custom validations etc. here
		v.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			if name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]; name != "-" && name != "" {
				return name
			}
			if form := strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]; form != "-" && form != "" {
				return form
			}
			if xml := strings.SplitN(fld.Tag.Get("xml"), ",", 2)[0]; xml != "-" && xml != "" {
				return xml
			}
			return ""
		})
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
