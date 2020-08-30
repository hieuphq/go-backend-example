package vietnamese

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// RegisterDefaultTranslations registers a set of default translations
// for all built in tag's in validator; you may add your own as desired.
func RegisterDefaultTranslations(v *validator.Validate, trans ut.Translator) (err error) {

	translations := []struct {
		tag             string
		translation     string
		override        bool
		customRegisFunc validator.RegisterTranslationsFunc
		customTransFunc validator.TranslationFunc
	}{
		{
			tag:         "required",
			translation: "{0} là trường bắt buộc",
			override:    false,
		},
		{
			tag: "len",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("len-string", "Độ dài {0} phải là {1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("len-string-character", "{0} ký tự", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("len-string-character", "{0} ký tự", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("len-number", "{0} phải bằng {1}", false); err != nil {
					return
				}

				if err = ut.Add("len-items", "{0} phải chứa {1}", false); err != nil {
					return
				}
				//if err = ut.AddCardinal("len-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("len-items-item", "{0} thành phần con", locales.PluralRuleOther, false); err != nil {
					return
				}

				return

			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("len-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("len-string", translateField(ut, fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("len-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("len-items", translateField(ut, fe), c)

				default:
					t, err = ut.T("len-number", translateField(ut, fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "min",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("min-string", "Độ dài {0} ít nhất phải bằng {1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("min-string-character", "{0} ký tự", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("min-string-character", "{0} ký tự", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("min-number", "{0} nhỏ nhất chỉ có thể là {1}", false); err != nil {
					return
				}

				if err = ut.Add("min-items", "{0} phải chứa ít nhất {1}", false); err != nil {
					return
				}
				//if err = ut.AddCardinal("min-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("min-items-item", "{0} thành phần con", locales.PluralRuleOther, false); err != nil {
					return
				}

				return

			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("min-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-string", translateField(ut, fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("min-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("min-items", translateField(ut, fe), c)

				default:
					t, err = ut.T("min-number", translateField(ut, fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "max",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("max-string", "Độ dài {0} không được vượt quá {1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("max-string-character", "{0} ký tự", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("max-string-character", "{0} ký tự", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("max-number", "{0} phải nhỏ hơn hoặc bằng {1}", false); err != nil {
					return
				}

				if err = ut.Add("max-items", "{0} chỉ có thể chứa tối đa {1}", false); err != nil {
					return
				}
				//if err = ut.AddCardinal("max-items-item", "{0} thành phần con", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("max-items-item", "{0} thành phần con", locales.PluralRuleOther, false); err != nil {
					return
				}

				return

			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string

				var digits uint64
				var kind reflect.Kind

				if idx := strings.Index(fe.Param(), "."); idx != -1 {
					digits = uint64(len(fe.Param()[idx+1:]))
				}

				f64, err := strconv.ParseFloat(fe.Param(), 64)
				if err != nil {
					goto END
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					c, err = ut.C("max-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-string", translateField(ut, fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					c, err = ut.C("max-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("max-items", translateField(ut, fe), c)

				default:
					t, err = ut.T("max-number", translateField(ut, fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eq",
			translation: "{0} không bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ne",
			translation: "{0} không thể bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lt",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("lt-string", "Độ dài {0} phải nhỏ hơn {1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("lt-string-character", "{0} ký tự", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("lt-string-character", "{0} ký tự", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-number", "{0} phải nhỏ hơn {1}", false); err != nil {
					return
				}

				if err = ut.Add("lt-items", "{0} phải chứa ít hơn {1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("lt-items-item", "{0} thành phần con", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("lt-items-item", "{0} thành phần con", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lt-datetime", "{0} phải nhỏ hơn ngày giờ hiện tại", false); err != nil {
					return
				}

				return

			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {

					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-string", translateField(ut, fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lt-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-items", translateField(ut, fe), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
					} else {
						t, err = ut.T("lt-datetime", translateField(ut, fe))
					}

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lt-number", translateField(ut, fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "lte",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("lte-string", "Độ dài {0} không được vượt quá {1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("lte-string-character", "{0} character", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("lte-string-character", "{0} ký tự", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-number", "{0} phải nhỏ hơn hoặc bằng {1}", false); err != nil {
					return
				}

				if err = ut.Add("lte-items", "{0} chỉ có thể chứa tối đa {1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("lte-items-item", "{0} item", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("lte-items-item", "{0} thành phần con", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("lte-datetime", "{0} phải nhỏ hơn hoặc bằng ngày giờ hiện tại", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {

					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-string", translateField(ut, fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("lte-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-items", translateField(ut, fe), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
					} else {
						t, err = ut.T("lte-datetime", translateField(ut, fe))
					}

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("lte-number", translateField(ut, fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gt",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("gt-string", "Độ dài {0} phải lớn hơn {1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("gt-string-character", "{0} ký tự", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("gt-string-character", "{0} ký tự", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-number", "{0} phải lớn hơn {1}", false); err != nil {
					return
				}

				if err = ut.Add("gt-items", "{0} phải lớn hơn {1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("gt-items-item", "{0} thành phần con", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("gt-items-item", "{0} thành phần con", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gt-datetime", "{0} phải lớn hơn ngày giờ hiện tại", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {

					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gt-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-string", translateField(ut, fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gt-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-items", translateField(ut, fe), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
					} else {

						t, err = ut.T("gt-datetime", translateField(ut, fe))
					}

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gt-number", translateField(ut, fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag: "gte",
			customRegisFunc: func(ut ut.Translator) (err error) {

				if err = ut.Add("gte-string", "Độ dài {0} phải ít nhất là {1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("gte-string-character", "{0} ký tự", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("gte-string-character", "{0} ký tự", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-number", "{0} phải lớn hơn hoặc bằng {1}", false); err != nil {
					return
				}

				if err = ut.Add("gte-items", "{0} phải chứa ít nhất {1}", false); err != nil {
					return
				}

				//if err = ut.AddCardinal("gte-items-item", "{0} thành phần con", locales.PluralRuleOne, false); err != nil {
				//	return
				//}

				if err = ut.AddCardinal("gte-items-item", "{0} thành phần con", locales.PluralRuleOther, false); err != nil {
					return
				}

				if err = ut.Add("gte-datetime", "{0} phải lớn hơn hoặc bằng ngày giờ hiện tại", false); err != nil {
					return
				}

				return
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				var err error
				var t string
				var f64 float64
				var digits uint64
				var kind reflect.Kind

				fn := func() (err error) {

					if idx := strings.Index(fe.Param(), "."); idx != -1 {
						digits = uint64(len(fe.Param()[idx+1:]))
					}

					f64, err = strconv.ParseFloat(fe.Param(), 64)

					return
				}

				kind = fe.Kind()
				if kind == reflect.Ptr {
					kind = fe.Type().Elem().Kind()
				}

				switch kind {
				case reflect.String:

					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gte-string-character", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-string", translateField(ut, fe), c)

				case reflect.Slice, reflect.Map, reflect.Array:
					var c string

					err = fn()
					if err != nil {
						goto END
					}

					c, err = ut.C("gte-items-item", f64, digits, ut.FmtNumber(f64, digits))
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-items", translateField(ut, fe), c)

				case reflect.Struct:
					if fe.Type() != reflect.TypeOf(time.Time{}) {
						err = fmt.Errorf("tag '%s' cannot be used on a struct type", fe.Tag())
					} else {
						t, err = ut.T("gte-datetime", translateField(ut, fe))
					}

				default:
					err = fn()
					if err != nil {
						goto END
					}

					t, err = ut.T("gte-number", translateField(ut, fe), ut.FmtNumber(f64, digits))
				}

			END:
				if err != nil {
					fmt.Printf("warning: error translating FieldError: %s", err)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqfield",
			translation: "{0} phải bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "eqcsfield",
			translation: "{0} phải bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "necsfield",
			translation: "{0} không thể bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtcsfield",
			translation: "{0} phải lớn hơn {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtecsfield",
			translation: "{0} phải lớn hơn hoặc bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltcsfield",
			translation: "{0} phải nhỏ hơn {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltecsfield",
			translation: "{0} phải nhỏ hơn hoặc bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "nefield",
			translation: "{0} không thể bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtfield",
			translation: "{0} phải lớn hơn {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "gtefield",
			translation: "{0} phải lớn hơn hoặc bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltfield",
			translation: "{0} phải nhỏ hơn {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "ltefield",
			translation: "{0} phải nhỏ hơn hoặc bằng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "alpha",
			translation: "{0} chỉ có thể chứa các chữ cái",
			override:    false,
		},
		{
			tag:         "alphanum",
			translation: "{0} chỉ có thể chứa các chữ cái và số",
			override:    false,
		},
		{
			tag:         "numeric",
			translation: "{0} phải là số hợp lệ",
			override:    false,
		},
		{
			tag:         "number",
			translation: "{0} phải là số hợp lệ",
			override:    false,
		},
		{
			tag:         "hexadecimal",
			translation: "{0} phải là số hệ 16 hợp lệ",
			override:    false,
		},
		{
			tag:         "hexcolor",
			translation: "{0} phải là màu Hex hợp lệ",
			override:    false,
		},
		{
			tag:         "rgb",
			translation: "{0} phải là màu RGB hợp lệ",
			override:    false,
		},
		{
			tag:         "rgba",
			translation: "{0} phải là màu RGBA hợp lệ",
			override:    false,
		},
		{
			tag:         "hsl",
			translation: "{0} phải là màu HSL hợp lệ",
			override:    false,
		},
		{
			tag:         "hsla",
			translation: "{0} phải là màu HSLA hợp lệ",
			override:    false,
		},
		{
			tag:         "email",
			translation: "{0} phải là địa chỉ email hợp lệ",
			override:    false,
		},
		{
			tag:         "url",
			translation: "{0} phải là một URL hợp lệ",
			override:    false,
		},
		{
			tag:         "uri",
			translation: "{0} phải là một URI hợp lệ",
			override:    false,
		},
		{
			tag:         "base64",
			translation: "{0} phải là một chuỗi Base64 hợp lệ",
			override:    false,
		},
		{
			tag:         "contains",
			translation: "{0} phải chứa '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "containsany",
			translation: "{0} phải chứa ít nhất một trong các ký tự '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludes",
			translation: "{0} không được chứa '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesall",
			translation: "{0} không được chứa bất kỳ ký tự nào trong '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "excludesrune",
			translation: "{0} không được chứa '{1}'",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
		{
			tag:         "isbn",
			translation: "{0} phải là số ISBN hợp lệ",
			override:    false,
		},
		{
			tag:         "isbn10",
			translation: "{0} phải là số ISBN-10 hợp lệ",
			override:    false,
		},
		{
			tag:         "isbn13",
			translation: "{0} phải là số ISBN-13 hợp lệ",
			override:    false,
		},
		{
			tag:         "uuid",
			translation: "{0} phải là UUID",
			override:    false,
		},
		{
			tag:         "uuid3",
			translation: "{0} phải là UUID V3",
			override:    false,
		},
		{
			tag:         "uuid4",
			translation: "{0} phải là UUID V4",
			override:    false,
		},
		{
			tag:         "uuid5",
			translation: "{0} phải là UUID V5",
			override:    false,
		},
		{
			tag:         "ascii",
			translation: "{0} chỉ được chứa các ký tự ascii",
			override:    false,
		},
		{
			tag:         "printascii",
			translation: "{0} chỉ được chứa các ký tự ascii có thể in được",
			override:    false,
		},
		{
			tag:         "multibyte",
			translation: "{0} chỉ chứa các ký tự multibyte",
			override:    false,
		},
		{
			tag:         "datauri",
			translation: "{0} phải chứa một URI dữ liệu hợp lệ",
			override:    false,
		},
		{
			tag:         "latitude",
			translation: "{0} phải chứa tọa độ vĩ độ hợp lệ",
			override:    false,
		},
		{
			tag:         "longitude",
			translation: "{0} phải chứa một tọa độ kinh độ hợp lệ",
			override:    false,
		},
		{
			tag:         "ssn",
			translation: "{0} phải là Số An sinh Xã hội (SSN) hợp lệ",
			override:    false,
		},
		{
			tag:         "ipv4",
			translation: "{0} phải là địa chỉ IPv4 hợp lệ",
			override:    false,
		},
		{
			tag:         "ipv6",
			translation: "{0} phải là địa chỉ IPv6 hợp lệ",
			override:    false,
		},
		{
			tag:         "ip",
			translation: "{0} phải là địa chỉ IP hợp lệ",
			override:    false,
		},
		{
			tag:         "cidr",
			translation: "{0} phải chứa ký tự CIDR hợp lệ",
			override:    false,
		},
		{
			tag:         "cidrv4",
			translation: "{0} phải chứa ký tự CIDR hợp lệ cho địa chỉ IPv4",
			override:    false,
		},
		{
			tag:         "cidrv6",
			translation: "{0} phải chứa ký tự CIDR hợp lệ cho địa chỉ IPv6",
			override:    false,
		},
		{
			tag:         "tcp_addr",
			translation: "{0} phải là một địa chỉ TCP hợp lệ",
			override:    false,
		},
		{
			tag:         "tcp4_addr",
			translation: "{0} phải là địa chỉ IPv4 TCP hợp lệ",
			override:    false,
		},
		{
			tag:         "tcp6_addr",
			translation: "{0} phải là địa chỉ IPv6 TCP hợp lệ",
			override:    false,
		},
		{
			tag:         "udp_addr",
			translation: "{0} phải là địa chỉ UDP hợp lệ",
			override:    false,
		},
		{
			tag:         "udp4_addr",
			translation: "{0} phải là địa chỉ IPv4 UDP hợp lệ",
			override:    false,
		},
		{
			tag:         "udp6_addr",
			translation: "{0} phải là địa chỉ UDP IPv6 hợp lệ",
			override:    false,
		},
		{
			tag:         "ip_addr",
			translation: "{0} phải là địa chỉ IP hợp lệ",
			override:    false,
		},
		{
			tag:         "ip4_addr",
			translation: "{0} phải là địa chỉ IPv4 hợp lệ",
			override:    false,
		},
		{
			tag:         "ip6_addr",
			translation: "{0} phải là địa chỉ IPv6 hợp lệ",
			override:    false,
		},
		{
			tag:         "unix_addr",
			translation: "{0} phải là địa chỉ UNIX hợp lệ",
			override:    false,
		},
		{
			tag:         "mac",
			translation: "{0} phải là địa chỉ MAC hợp lệ",
			override:    false,
		},
		{
			tag:         "iscolor",
			translation: "{0} phải là màu hợp lệ",
			override:    false,
		},
		{
			tag:         "oneof",
			translation: "{0} phải là một trong [{1}]",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				s, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}
				return s
			},
		},
		{
			tag:         "json",
			translation: "{0} phải là chuỗi JSON",
			override:    false,
		},
		{
			tag:         "lowercase",
			translation: "{0} phải là chữ thường",
			override:    false,
		},
		{
			tag:         "uppercase",
			translation: "{0} phải được viết hoa",
			override:    false,
		},
		{
			tag:         "datetime",
			translation: "{0} phải là định dạng {1}",
			override:    false,
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), translateField(ut, fe), fe.Param())
				if err != nil {
					log.Printf("warning: error translating FieldError: %#v", fe)
					return fe.(error).Error()
				}

				return t
			},
		},
	}

	for _, t := range translations {

		if t.customTransFunc != nil && t.customRegisFunc != nil {

			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, t.customTransFunc)

		} else if t.customTransFunc != nil && t.customRegisFunc == nil {

			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), t.customTransFunc)

		} else if t.customTransFunc == nil && t.customRegisFunc != nil {

			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, translateFunc)

		} else {
			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), translateFunc)
		}

		if err != nil {
			return
		}
	}

	return
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {

	return func(ut ut.Translator) (err error) {

		if err = ut.Add(tag, translation, override); err != nil {
			return
		}

		return

	}

}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), translateField(ut, fe))
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}

func translateField(ut ut.Translator, fe validator.FieldError) string {
	f := fe.Field()
	tf, err := ut.T(f)
	if err != nil {
		return f
	}

	return tf
}
