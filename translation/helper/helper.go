package helper

import (
	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
)

// Add add normal translation and wraps error
func Add(trans ut.Translator, key interface{}, text string, override bool) error {
	return trans.Add(key, text, override)
}

// AddRange adds a range translation and wraps error
func AddRange(trans ut.Translator, key interface{}, text string, rule locales.PluralRule, override bool) error {
	return trans.AddRange(key, text, rule, override)
}

// AddCardinal adds a cardinal translation and wraps error
func AddCardinal(trans ut.Translator, key interface{}, text string, rule locales.PluralRule, override bool) error {
	return trans.AddCardinal(key, text, rule, override)
}

// AddOrdinal adds an ordinal translation and wraps error
func AddOrdinal(trans ut.Translator, key interface{}, text string, rule locales.PluralRule, override bool) error {
	return trans.AddOrdinal(key, text, rule, override)
}
