package translation

import (
	english "github.com/hieuphq/backend-example/translation/en"
	vietnamese "github.com/hieuphq/backend-example/translation/vi"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/vi"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// Helper app translator helper
type Helper interface {
	InitErrorTranslator(validate *validator.Validate)
	GetTranslator(locale string) ut.Translator
}

// NewTranslatorHelper make a translator helper
func NewTranslatorHelper() Helper {
	// initialize translator
	en := en.New()
	vi := vi.New()
	uni := ut.New(en, en, vi)
	defTr, _ := uni.GetTranslator("en")
	return &helper{
		uni:           uni,
		defTranslator: defTr,
	}
}

type helper struct {
	uni           *ut.UniversalTranslator
	defTranslator ut.Translator
}

func (h *helper) InitErrorTranslator(validate *validator.Validate) {
	// initialize translations
	vietnamese.Init(h.uni, validate)
	english.Init(h.uni, validate)

	defTrans, _ := h.uni.FindTranslator("en")
	h.defTranslator = defTrans
}

func (h *helper) GetTranslator(locale string) ut.Translator {
	if locale == "" {
		return h.defTranslator
	}
	t, found := h.uni.GetTranslator(locale)
	if !found {
		return h.defTranslator
	}
	return t
}
