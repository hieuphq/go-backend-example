package vietnamese

import (
	"errors"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

const (
	locale = "vi"
)

// Init initializes the english locale translations
func Init(uni *ut.UniversalTranslator, validate *validator.Validate) error {
	vi, found := uni.GetTranslator(locale)
	if !found {
		return errors.New("Translation not found")
	}

	vi.Add("username", "tên đăng nhập", false)
	vi.Add("password", "mật khẩu", false)
	vi.Add("amount", "số lượng sản phẩm", false)
	vi.Add("bad request", "không thể thực hiện yêu cầu", false)

	// validator translations & Overrides
	err := RegisterDefaultTranslations(validate, vi)
	if err != nil {
		return errors.New("Error adding default translations: " + err.Error())
	}

	if err := vi.VerifyTranslations(); err != nil {
		return errors.New("Missing Translations: " + err.Error())
	}
	return nil
}
