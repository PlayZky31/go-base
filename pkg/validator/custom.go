package validator

import (
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

const (
	ErrInvalidDateMSG = "the date should be in the format 'Y-M-D'"
)

func (v ValidationEngine) addTranslation(tag, errMessage string) {
	registerFn := func(ut ut.Translator) error {
		return ut.Add(tag, errMessage, false)
	}

	transFn := func(ut ut.Translator, fe validator.FieldError) string {
		param := fe.Param()
		feTag := fe.Tag()

		t, err := ut.T(feTag, fe.Field(), param)
		if err != nil {
			return fe.(error).Error()
		}
		return t
	}

	_ = v.Validator.RegisterTranslation(tag, v.ENTranslator, registerFn, transFn)
}

func validateCustomDate(fl validator.FieldLevel) bool {
	// 2023-12-31
	layout := "2006-01-02"
	_, err := time.Parse(layout, fl.Field().String())

	return err == nil
}
