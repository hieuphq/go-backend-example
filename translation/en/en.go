package english

import (
	"errors"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	valtrans "github.com/go-playground/validator/v10/translations/en"
)

const (
	locale = "en"
)

// Init initializes the english locale translations
func Init(uni *ut.UniversalTranslator, validate *validator.Validate) error {
	en, found := uni.GetTranslator(locale)
	if !found {
		return errors.New("Translation not found")
	}

	// validator translations & Overrides
	err := valtrans.RegisterDefaultTranslations(validate, en)
	if err != nil {
		return errors.New("Error adding default translations: " + err.Error())
	}

	if err := en.VerifyTranslations(); err != nil {
		return errors.New("Missing Translations: " + err.Error())
	}
	return nil
}
